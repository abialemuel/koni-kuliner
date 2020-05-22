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
	filter := utility.Filter(r, []string{"id", "customer_id", "state", "offset", "limit"})

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
	transactionID, _ := strconv.ParseInt(params.ByName("ID"), 10, 64)

	// run query
	var model models.Transaction

	// return not found is record not exist
	if mysql.db.First(&model, transactionID).RecordNotFound() {
		utility.SendErrorResponse(w, entity.TransactionNotFoundError)
		return
	}
	GetSingleDetailRelationTransaction(mysql, &model)
	result := utility.TransactionDetailResponse(model)
	utility.SendSuccessResponse(w, result, http.StatusOK)
}

// func (mysql *Mysql) CreateTransaction(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
// 	// decode params
// 	var transactionRequest request.TransactionCreateRequest

// 	err := json.NewDecoder(r.Body).Decode(&transactionRequest)
// 	if err != nil {
// 		utility.SendErrorResponse(w, entity.FailedDecodeJSONError)
// 		return
// 	}

// 	// validate body params
// 	v := validator.New()
// 	err = v.Struct(transactionRequest)

// 	if err != nil {
// 		println("error: " + err.Error())
// 		utility.SendErrorResponse(w, entity.UnprocessableEntityError)
// 		return
// 	}

// 	// assign body params
// 	model := models.Transaction{
// 		ProductID:  transactionRequest.ProductID,
// 		OutletID:   transactionRequest.OutletID,
// 		State:      models.TransactionStateActive,
// 		Price:      transactionRequest.Price,
// 		OrderPrice: transactionRequest.OrderPrice,
// 	}

// 	mysql.db.Create(&model)
// 	GetSingleDetailRelationTransaction(mysql, &model)
// 	result := utility.TransactionDetailResponse(model)
// 	utility.SendSuccessResponse(w, result, http.StatusCreated)
// }

// func (mysql *Mysql) UpdateTransaction(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
// 	transactionID, _ := strconv.ParseInt(params.ByName("ID"), 10, 64)

// 	// run query
// 	var model models.Transaction

// 	// return not found if record not exist
// 	if mysql.db.First(&model, transactionID).RecordNotFound() {
// 		utility.SendErrorResponse(w, entity.TransactionNotFoundError)
// 		return
// 	}

// 	var transactionRequest request.TransactionUpdateRequest

// 	err := json.NewDecoder(r.Body).Decode(&transactionRequest)
// 	if err != nil {
// 		utility.SendErrorResponse(w, entity.FailedDecodeJSONError)
// 		return
// 	}

// 	// validate body params
// 	v := validator.New()
// 	err = v.Struct(transactionRequest)

// 	if err != nil {
// 		println("error: " + err.Error())
// 		utility.SendErrorResponse(w, entity.UnprocessableEntityError)
// 		return
// 	}

// 	mysql.db.Model(&model).Updates(
// 		models.Transaction{
// 			State:      models.ToTransactionStateType(transactionRequest.State),
// 			Price:      transactionRequest.Price,
// 			OrderPrice: transactionRequest.OrderPrice,
// 		},
// 	)
// 	GetSingleDetailRelationTransaction(mysql, &model)
// 	result := utility.TransactionDetailResponse(model)
// 	utility.SendSuccessResponse(w, result, http.StatusOK)
// }

func (mysql *Mysql) DeleteTransaction(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	transactionID, _ := strconv.ParseInt(params.ByName("ID"), 10, 64)

	// run query
	var model models.Transaction

	// return not found if record not exist
	if mysql.db.First(&model, transactionID).RecordNotFound() {
		utility.SendErrorResponse(w, entity.TransactionNotFoundError)
		return
	}
	mysql.db.Delete(&model)
	w.WriteHeader(http.StatusOK)
}

// private func
