package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/koni-kuliner/entity"
	"github.com/koni-kuliner/models"
	"github.com/koni-kuliner/resource/request"
	"github.com/koni-kuliner/utility"
	"gopkg.in/go-playground/validator.v9"
)

func (mysql *Mysql) GetCustomers(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var filteredArgs []interface{}

	// filter query params
	filter := utility.Filter(r, []string{"id", "name", "address", "offset", "limit"})

	// build query
	query := "SELECT * FROM customers WHERE 1=1"
	query, filteredArgs = utility.AppendQuery(query, filter)

	// run query
	var model []models.Customer
	mysql.db.Raw(query, filteredArgs...).Scan(&model)
	result := utility.CustomerResponse(model)
	utility.SendSuccessResponseWithLimitAndOffset(w, result, http.StatusOK, filter, countCustomer(mysql))
}

func (mysql *Mysql) GetCustomerDetails(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	customerID, _ := strconv.ParseInt(params.ByName("ID"), 10, 64)

	// run query
	var model models.Customer

	// return not found is record not exist
	if mysql.db.First(&model, customerID).RecordNotFound() {
		utility.SendErrorResponse(w, entity.CustomerNotFoundError)
		return
	}

	result := utility.CustomerDetailResponse(model)
	utility.SendSuccessResponse(w, result, http.StatusOK)
}

func (mysql *Mysql) CreateCustomer(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// decode params
	var customerRequest request.CustomerCreateRequest

	err := json.NewDecoder(r.Body).Decode(&customerRequest)
	if err != nil {
		utility.SendErrorResponse(w, entity.FailedDecodeJSONError)
		return
	}

	// validate body params
	v := validator.New()
	err = v.Struct(customerRequest)

	if err != nil {
		println("error: " + err.Error())
		utility.SendErrorResponse(w, entity.UnprocessableEntityError)
		return
	}

	// assign body params
	model := models.Customer{
		Name:      customerRequest.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mysql.db.Create(&model)
	result := utility.CustomerDetailResponse(model)
	utility.SendSuccessResponse(w, result, http.StatusCreated)
}

func (mysql *Mysql) UpdateCustomer(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	customerID, _ := strconv.ParseInt(params.ByName("ID"), 10, 64)

	// run query
	var model models.Customer

	// return not found if record not exist
	if mysql.db.First(&model, customerID).RecordNotFound() {
		utility.SendErrorResponse(w, entity.CustomerNotFoundError)
		return
	}

	var customerRequest request.CustomerUpdateRequest

	err := json.NewDecoder(r.Body).Decode(&customerRequest)
	if err != nil {
		utility.SendErrorResponse(w, entity.FailedDecodeJSONError)
		return
	}

	// validate body params
	v := validator.New()
	err = v.Struct(customerRequest)

	if err != nil {
		println("error: " + err.Error())
		utility.SendErrorResponse(w, entity.UnprocessableEntityError)
		return
	}

	mysql.db.Model(&model).Updates(customerRequest)
	result := utility.CustomerDetailResponse(model)
	utility.SendSuccessResponse(w, result, http.StatusOK)
}

func (mysql *Mysql) DeleteCustomer(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	customerID, _ := strconv.ParseInt(params.ByName("ID"), 10, 64)

	// run query
	var model models.Customer

	// return not found if record not exist
	if mysql.db.First(&model, customerID).RecordNotFound() {
		utility.SendErrorResponse(w, entity.CustomerNotFoundError)
		return
	}
	mysql.db.Delete(&model)
	w.WriteHeader(http.StatusOK)
}

// private func

func countCustomer(mysql *Mysql) int {
	var count int
	mysql.db.Table("customers").Count(&count)
	return count
}
