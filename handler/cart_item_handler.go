package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/koni-kuliner/entity"
	"github.com/koni-kuliner/models"
	"github.com/koni-kuliner/resource/request"
	"github.com/koni-kuliner/utility"
	"gopkg.in/go-playground/validator.v9"
)

func (mysql *Mysql) GetCartItems(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var filteredArgs []interface{}

	// filter query params
	filter := utility.Filter(r, []string{"id", "customer_id", "offset", "limit"})

	// build query
	query := "SELECT * FROM cart_items WHERE 1=1"
	query, filteredArgs = utility.AppendQuery(query, filter)

	// run query
	var model []models.CartItem
	mysql.db.Raw(query, filteredArgs...).Scan(&model)

	GetAllDetailRelationCartItem(mysql, &model)

	result := utility.CartItemResponse(model)
	utility.SendSuccessResponseWithLimitAndOffset(w, result, http.StatusOK, filter, CountQuery(mysql, query, filteredArgs))
}

func (mysql *Mysql) CreateCartItem(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// decode params
	var cartItemRequest request.CartItemCreateRequest

	err := json.NewDecoder(r.Body).Decode(&cartItemRequest)
	if err != nil {
		utility.SendErrorResponse(w, entity.FailedDecodeJSONError)
		return
	}

	// validate body params
	v := validator.New()
	err = v.Struct(cartItemRequest)

	if err != nil {
		println("error: " + err.Error())
		utility.SendErrorResponse(w, entity.UnprocessableEntityError)
		return
	}

	// TO DO: validate outlet product

	// get outlet_product
	var outletProduct models.OutletProduct
	mysql.db.First(&outletProduct, cartItemRequest.OutletProductID)

	// assign body params
	model := models.CartItem{
		OutletProductID: cartItemRequest.OutletProductID,
		CustomerID:      cartItemRequest.CustomerID,
		Quantity:        cartItemRequest.Quantity,
		Price:           outletProduct.Price,
		OrderPrice:      outletProduct.OrderPrice,
	}

	mysql.db.Create(&model)
	GetSingleDetailRelationCartItem(mysql, &model)
	result := utility.CartItemDetailResponse(model)
	utility.SendSuccessResponse(w, result, http.StatusCreated)
}

func (mysql *Mysql) UpdateCartItem(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	cartItemID, _ := strconv.ParseInt(params.ByName("ID"), 10, 64)

	// run query
	var model models.CartItem

	// return not found if record not exist
	if mysql.db.First(&model, cartItemID).RecordNotFound() {
		utility.SendErrorResponse(w, entity.CartItemNotFoundError)
		return
	}

	var cartItemRequest request.CartItemUpdateRequest

	err := json.NewDecoder(r.Body).Decode(&cartItemRequest)
	if err != nil {
		utility.SendErrorResponse(w, entity.FailedDecodeJSONError)
		return
	}

	// validate body params
	v := validator.New()
	err = v.Struct(cartItemRequest)

	if err != nil {
		println("error: " + err.Error())
		utility.SendErrorResponse(w, entity.UnprocessableEntityError)
		return
	}

	mysql.db.Model(&model).Updates(
		models.CartItem{
			Quantity: cartItemRequest.Quantity,
		},
	)
	GetSingleDetailRelationCartItem(mysql, &model)
	result := utility.CartItemDetailResponse(model)
	utility.SendSuccessResponse(w, result, http.StatusOK)
}

func (mysql *Mysql) DeleteCartItem(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	cartItemID, _ := strconv.ParseInt(params.ByName("ID"), 10, 64)

	// run query
	var model models.CartItem

	// return not found if record not exist
	if mysql.db.First(&model, cartItemID).RecordNotFound() {
		utility.SendErrorResponse(w, entity.CartItemNotFoundError)
		return
	}
	mysql.db.Delete(&model)
	w.WriteHeader(http.StatusOK)
}

// private func
