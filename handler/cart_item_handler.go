package handler

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/koni-kuliner/entity"
	"github.com/koni-kuliner/models"
	"github.com/koni-kuliner/utility"
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

// func (mysql *Mysql) GetCartItemDetails(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
// 	cartItemID, _ := strconv.ParseInt(params.ByName("ID"), 10, 64)

// 	// run query
// 	var model models.CartItem

// 	// return not found is record not exist
// 	if mysql.db.First(&model, cartItemID).RecordNotFound() {
// 		utility.SendErrorResponse(w, entity.CartItemNotFoundError)
// 		return
// 	}
// 	GetSingleDetailRelationCartItem(mysql, &model)
// 	result := utility.CartItemDetailResponse(model)
// 	utility.SendSuccessResponse(w, result, http.StatusOK)
// }

// func (mysql *Mysql) CreateCartItem(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
// 	// decode params
// 	var cartItemRequest request.CartItemCreateRequest

// 	err := json.NewDecoder(r.Body).Decode(&cartItemRequest)
// 	if err != nil {
// 		utility.SendErrorResponse(w, entity.FailedDecodeJSONError)
// 		return
// 	}

// 	// validate body params
// 	v := validator.New()
// 	err = v.Struct(cartItemRequest)

// 	if err != nil {
// 		println("error: " + err.Error())
// 		utility.SendErrorResponse(w, entity.UnprocessableEntityError)
// 		return
// 	}

// 	// TO DO: validate product & outlet
// 	// var outlet models.Outlet
// 	// var product models.Product
// 	// if mysql.db.First(product, cartItemRequest.ProductID).RecordNotFound() {
// 	// 	utility.SendErrorResponse(w, entity.ProductNotValidError)
// 	// 	return
// 	// }

// 	// if mysql.db.First(outlet, cartItemRequest.OutletID).RecordNotFound() {
// 	// 	utility.SendErrorResponse(w, entity.OutletNotValidError)
// 	// 	return
// 	// }

// 	// assign body params
// 	model := models.CartItem{
// 		ProductID:  cartItemRequest.ProductID,
// 		OutletID:   cartItemRequest.OutletID,
// 		State:      models.CartItemStateActive,
// 		Price:      cartItemRequest.Price,
// 		OrderPrice: cartItemRequest.OrderPrice,
// 	}

// 	mysql.db.Create(&model)
// 	GetSingleDetailRelationCartItem(mysql, &model)
// 	result := utility.CartItemDetailResponse(model)
// 	utility.SendSuccessResponse(w, result, http.StatusCreated)
// }

// func (mysql *Mysql) UpdateCartItem(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
// 	cartItemID, _ := strconv.ParseInt(params.ByName("ID"), 10, 64)

// 	// run query
// 	var model models.CartItem

// 	// return not found if record not exist
// 	if mysql.db.First(&model, cartItemID).RecordNotFound() {
// 		utility.SendErrorResponse(w, entity.CartItemNotFoundError)
// 		return
// 	}

// 	var cartItemRequest request.CartItemUpdateRequest

// 	err := json.NewDecoder(r.Body).Decode(&cartItemRequest)
// 	if err != nil {
// 		utility.SendErrorResponse(w, entity.FailedDecodeJSONError)
// 		return
// 	}

// 	// validate body params
// 	v := validator.New()
// 	err = v.Struct(cartItemRequest)

// 	if err != nil {
// 		println("error: " + err.Error())
// 		utility.SendErrorResponse(w, entity.UnprocessableEntityError)
// 		return
// 	}

// 	mysql.db.Model(&model).Updates(
// 		models.CartItem{
// 			State:      models.ToCartItemStateType(cartItemRequest.State),
// 			Price:      cartItemRequest.Price,
// 			OrderPrice: cartItemRequest.OrderPrice,
// 		},
// 	)
// 	GetSingleDetailRelationCartItem(mysql, &model)
// 	result := utility.CartItemDetailResponse(model)
// 	utility.SendSuccessResponse(w, result, http.StatusOK)
// }

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
