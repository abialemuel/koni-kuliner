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

func (mysql *Mysql) GetSellers(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var filteredArgs []interface{}

	// filter query params
	filter := utility.Filter(r, []string{"id", "name", "address", "offset", "limit"})

	// build query
	query := "SELECT * FROM sellers WHERE 1=1"
	query, filteredArgs = utility.AppendQuery(query, filter)

	// run query
	var model []models.Seller
	mysql.db.Raw(query, filteredArgs...).Scan(&model)
	result := utility.SellerResponse(model)
	utility.SendSuccessResponseWithLimitAndOffset(w, result, http.StatusOK, filter, CountQuery(mysql, query, filteredArgs))
}

func (mysql *Mysql) GetSellerDetails(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	sellerID, _ := strconv.ParseInt(params.ByName("ID"), 10, 64)

	// run query
	var model models.Seller

	// return not found is record not exist
	if mysql.db.First(&model, sellerID).RecordNotFound() {
		utility.SendErrorResponse(w, entity.SellerNotFoundError)
		return
	}

	result := utility.SellerDetailResponse(model)
	utility.SendSuccessResponse(w, result, http.StatusOK)
}

func (mysql *Mysql) CreateSeller(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// decode params
	var sellerRequest request.SellerCreateRequest

	err := json.NewDecoder(r.Body).Decode(&sellerRequest)
	if err != nil {
		utility.SendErrorResponse(w, entity.FailedDecodeJSONError)
		return
	}

	// validate body params
	v := validator.New()
	err = v.Struct(sellerRequest)

	if err != nil {
		println("error: " + err.Error())
		utility.SendErrorResponse(w, entity.UnprocessableEntityError)
		return
	}

	// assign body params
	model := models.Seller{
		Name:      sellerRequest.Name,
		Address:   sellerRequest.Address,
		Phone:     sellerRequest.Phone,
		Username:  sellerRequest.Username,
		Password:  sellerRequest.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mysql.db.Create(&model)
	result := utility.SellerDetailResponse(model)
	utility.SendSuccessResponse(w, result, http.StatusCreated)
}

func (mysql *Mysql) UpdateSeller(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	sellerID, _ := strconv.ParseInt(params.ByName("ID"), 10, 64)

	// run query
	var model models.Seller

	// return not found if record not exist
	if mysql.db.First(&model, sellerID).RecordNotFound() {
		utility.SendErrorResponse(w, entity.SellerNotFoundError)
		return
	}

	var sellerRequest request.SellerUpdateRequest

	err := json.NewDecoder(r.Body).Decode(&sellerRequest)
	if err != nil {
		utility.SendErrorResponse(w, entity.FailedDecodeJSONError)
		return
	}

	// validate body params
	v := validator.New()
	err = v.Struct(sellerRequest)

	if err != nil {
		println("error: " + err.Error())
		utility.SendErrorResponse(w, entity.UnprocessableEntityError)
		return
	}

	mysql.db.Model(&model).Updates(sellerRequest)
	result := utility.SellerDetailResponse(model)
	utility.SendSuccessResponse(w, result, http.StatusOK)
}

func (mysql *Mysql) DeleteSeller(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	sellerID, _ := strconv.ParseInt(params.ByName("ID"), 10, 64)

	// run query
	var model models.Seller

	// return not found if record not exist
	if mysql.db.First(&model, sellerID).RecordNotFound() {
		utility.SendErrorResponse(w, entity.SellerNotFoundError)
		return
	}
	mysql.db.Delete(&model)
	w.WriteHeader(http.StatusOK)
}

// private func
