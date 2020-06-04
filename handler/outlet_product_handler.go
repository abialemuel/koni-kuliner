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

func (mysql *Mysql) GetOutletProducts(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var filteredArgs []interface{}

	// filter query params
	filter := utility.Filter(r, []string{"id", "outlet_id", "offset", "limit"})

	// build query
	query := "SELECT * FROM outlet_products WHERE 1=1"
	query, filteredArgs = utility.AppendQuery(query, filter)

	// run query
	var model []models.OutletProduct
	mysql.db.Raw(query, filteredArgs...).Scan(&model)

	getAllDetailRelationOutletProduct(mysql, &model)

	result := utility.OutletProductResponse(model)
	utility.SendSuccessResponseWithLimitAndOffset(w, result, http.StatusOK, filter, CountQuery(mysql, query, filteredArgs))
}

func (mysql *Mysql) GetOutletProductDetails(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	outletProductID, _ := strconv.ParseInt(params.ByName("ID"), 10, 64)

	// run query
	var model models.OutletProduct

	// return not found is record not exist
	if mysql.db.First(&model, outletProductID).RecordNotFound() {
		utility.SendErrorResponse(w, entity.OutletProductNotFoundError)
		return
	}
	GetSingleDetailRelationOutletProduct(mysql, &model)
	result := utility.OutletProductDetailResponse(model)
	utility.SendSuccessResponse(w, result, http.StatusOK)
}

func (mysql *Mysql) CreateOutletProduct(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// decode params
	var outletProductRequest request.OutletProductCreateRequest

	err := json.NewDecoder(r.Body).Decode(&outletProductRequest)
	if err != nil {
		utility.SendErrorResponse(w, entity.FailedDecodeJSONError)
		return
	}

	// validate body params
	v := validator.New()
	err = v.Struct(outletProductRequest)

	if err != nil {
		println("error: " + err.Error())
		utility.SendErrorResponse(w, entity.UnprocessableEntityError)
		return
	}

	// TO DO: validate product & outlet
	// var outlet models.Outlet
	// var product models.Product
	// if mysql.db.First(product, outletProductRequest.ProductID).RecordNotFound() {
	// 	utility.SendErrorResponse(w, entity.ProductNotValidError)
	// 	return
	// }

	// if mysql.db.First(outlet, outletProductRequest.OutletID).RecordNotFound() {
	// 	utility.SendErrorResponse(w, entity.OutletNotValidError)
	// 	return
	// }

	// assign body params
	model := models.OutletProduct{
		ProductID:  outletProductRequest.ProductID,
		OutletID:   outletProductRequest.OutletID,
		State:      models.OutletProductStateActive,
		Price:      outletProductRequest.Price,
		OrderPrice: outletProductRequest.OrderPrice,
	}

	mysql.db.Create(&model)
	GetSingleDetailRelationOutletProduct(mysql, &model)
	result := utility.OutletProductDetailResponse(model)
	utility.SendSuccessResponse(w, result, http.StatusCreated)
}

func (mysql *Mysql) UpdateOutletProduct(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	outletProductID, _ := strconv.ParseInt(params.ByName("ID"), 10, 64)

	// run query
	var model models.OutletProduct

	// return not found if record not exist
	if mysql.db.First(&model, outletProductID).RecordNotFound() {
		utility.SendErrorResponse(w, entity.OutletProductNotFoundError)
		return
	}

	var outletProductRequest request.OutletProductUpdateRequest

	err := json.NewDecoder(r.Body).Decode(&outletProductRequest)
	if err != nil {
		utility.SendErrorResponse(w, entity.FailedDecodeJSONError)
		return
	}

	// validate body params
	v := validator.New()
	err = v.Struct(outletProductRequest)

	if err != nil {
		println("error: " + err.Error())
		utility.SendErrorResponse(w, entity.UnprocessableEntityError)
		return
	}

	mysql.db.Model(&model).Updates(
		models.OutletProduct{
			State:      models.ToOutletProductStateType(outletProductRequest.State),
			Price:      outletProductRequest.Price,
			OrderPrice: outletProductRequest.OrderPrice,
		},
	)
	GetSingleDetailRelationOutletProduct(mysql, &model)
	result := utility.OutletProductDetailResponse(model)
	utility.SendSuccessResponse(w, result, http.StatusOK)
}

func (mysql *Mysql) DeleteOutletProduct(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	outletProductID, _ := strconv.ParseInt(params.ByName("ID"), 10, 64)

	// run query
	var model models.OutletProduct

	// return not found if record not exist
	if mysql.db.First(&model, outletProductID).RecordNotFound() {
		utility.SendErrorResponse(w, entity.OutletProductNotFoundError)
		return
	}
	mysql.db.Delete(&model)
	w.WriteHeader(http.StatusOK)
}

// private func

func getAllDetailRelationOutletProduct(mysql *Mysql, outletProduct *[]models.OutletProduct) {
	for i, m := range *outletProduct {
		var product models.Product
		mysql.db.First(&product, m.ProductID)

		var outlet models.Outlet
		mysql.db.First(&outlet, m.OutletID)

		(*outletProduct)[i].Product = product
		(*outletProduct)[i].Outlet = outlet
	}
}
