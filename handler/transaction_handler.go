package handler

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/koni-kuliner/entity"
	"github.com/koni-kuliner/models"
	"github.com/koni-kuliner/utility"
)

func (mysql *Mysql) GetTransactions(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var filteredArgs []interface{}

	// filter query params
	filter := utility.Filter(r, []string{"id", "customer_id", "offset", "limit"})

	// build query
	query := "SELECT * FROM transactions WHERE 1=1"
	query, filteredArgs = utility.AppendQuery(query, filter)

	// run query
	var model []models.Transaction
	mysql.db.Raw(query, filteredArgs...).Scan(&model)

	GetAllDetailRelationTransaction(mysql, &model)

	result := utility.TransactionResponse(model)
	utility.SendSuccessResponseWithLimitAndOffset(w, result, http.StatusOK, filter, CountQuery(mysql, query, filteredArgs))
}

func (mysql *Mysql) GetTransactionDetails(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	outletPoductID, _ := strconv.ParseInt(params.ByName("ID"), 10, 64)

	// run query
	var model models.Transaction

	// return not found is record not exist
	if mysql.db.First(&model, outletPoductID).RecordNotFound() {
		utility.SendErrorResponse(w, entity.TransactionNotFoundError)
		return
	}
	GetSingleDetailRelationTransaction(mysql, &model)
	result := utility.TransactionDetailResponse(model)
	utility.SendSuccessResponse(w, result, http.StatusOK)
}

// func (mysql *Mysql) CreateTransaction(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
// 	// decode params
// 	var outletPoductRequest request.TransactionCreateRequest

// 	err := json.NewDecoder(r.Body).Decode(&outletPoductRequest)
// 	if err != nil {
// 		utility.SendErrorResponse(w, entity.FailedDecodeJSONError)
// 		return
// 	}

// 	// validate body params
// 	v := validator.New()
// 	err = v.Struct(outletPoductRequest)

// 	if err != nil {
// 		println("error: " + err.Error())
// 		utility.SendErrorResponse(w, entity.UnprocessableEntityError)
// 		return
// 	}

// 	// TO DO: validate product & outlet
// 	// var outlet models.Outlet
// 	// var product models.Product
// 	// if mysql.db.First(product, outletPoductRequest.ProductID).RecordNotFound() {
// 	// 	utility.SendErrorResponse(w, entity.ProductNotValidError)
// 	// 	return
// 	// }

// 	// if mysql.db.First(outlet, outletPoductRequest.OutletID).RecordNotFound() {
// 	// 	utility.SendErrorResponse(w, entity.OutletNotValidError)
// 	// 	return
// 	// }

// 	// assign body params
// 	model := models.Transaction{
// 		ProductID:  outletPoductRequest.ProductID,
// 		OutletID:   outletPoductRequest.OutletID,
// 		State:      models.TransactionStateActive,
// 		Price:      outletPoductRequest.Price,
// 		OrderPrice: outletPoductRequest.OrderPrice,
// 	}

// 	mysql.db.Create(&model)
// 	GetSingleDetailRelationTransaction(mysql, &model)
// 	result := utility.TransactionDetailResponse(model)
// 	utility.SendSuccessResponse(w, result, http.StatusCreated)
// }

// func (mysql *Mysql) UpdateTransaction(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
// 	outletPoductID, _ := strconv.ParseInt(params.ByName("ID"), 10, 64)

// 	// run query
// 	var model models.Transaction

// 	// return not found if record not exist
// 	if mysql.db.First(&model, outletPoductID).RecordNotFound() {
// 		utility.SendErrorResponse(w, entity.TransactionNotFoundError)
// 		return
// 	}

// 	var outletPoductRequest request.TransactionUpdateRequest

// 	err := json.NewDecoder(r.Body).Decode(&outletPoductRequest)
// 	if err != nil {
// 		utility.SendErrorResponse(w, entity.FailedDecodeJSONError)
// 		return
// 	}

// 	// validate body params
// 	v := validator.New()
// 	err = v.Struct(outletPoductRequest)

// 	if err != nil {
// 		println("error: " + err.Error())
// 		utility.SendErrorResponse(w, entity.UnprocessableEntityError)
// 		return
// 	}

// 	mysql.db.Model(&model).Updates(
// 		models.Transaction{
// 			State:      models.ToTransactionStateType(outletPoductRequest.State),
// 			Price:      outletPoductRequest.Price,
// 			OrderPrice: outletPoductRequest.OrderPrice,
// 		},
// 	)
// 	GetSingleDetailRelationTransaction(mysql, &model)
// 	result := utility.TransactionDetailResponse(model)
// 	utility.SendSuccessResponse(w, result, http.StatusOK)
// }

func (mysql *Mysql) DeleteTransaction(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	outletPoductID, _ := strconv.ParseInt(params.ByName("ID"), 10, 64)

	// run query
	var model models.Transaction

	// return not found if record not exist
	if mysql.db.First(&model, outletPoductID).RecordNotFound() {
		utility.SendErrorResponse(w, entity.TransactionNotFoundError)
		return
	}
	mysql.db.Delete(&model)
	w.WriteHeader(http.StatusOK)
}

// private func
