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

func (mysql *Mysql) GetBrands(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var filteredArgs []interface{}

	// filter query params
	filter := utility.Filter(r, []string{"id", "name", "offset", "limit"})

	// build query
	query := "SELECT * FROM brands WHERE 1=1"
	query, filteredArgs = utility.AppendQuery(query, filter)

	// run query
	var model []models.Brand
	mysql.db.Raw(query, filteredArgs...).Scan(&model)
	result := utility.BrandResponse(model)
	utility.SendSuccessResponseWithLimitAndOffset(w, result, http.StatusOK, filter, countBrand(mysql))
}

func (mysql *Mysql) GetBrandDetails(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	brandID, _ := strconv.ParseInt(params.ByName("ID"), 10, 64)

	// run query
	var model models.Brand

	// return not found if record not exist
	if mysql.db.First(&model, brandID).RecordNotFound() {
		utility.SendErrorResponse(w, entity.BrandNotFoundError)
		return
	}
	mysql.db.First(&model, brandID)
	result := utility.BrandDetailResponse(model)
	utility.SendSuccessResponse(w, result, http.StatusOK)
}

func (mysql *Mysql) CreateBrand(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// decode params
	var brandRequest request.BrandCreateRequest

	err := json.NewDecoder(r.Body).Decode(&brandRequest)
	if err != nil {
		utility.SendErrorResponse(w, entity.FailedDecodeJSONError)
		return
	}

	// validate body params
	v := validator.New()
	err = v.Struct(brandRequest)

	if err != nil {
		println("error: " + err.Error())
		utility.SendErrorResponse(w, entity.UnprocessableEntityError)
		return
	}

	// assign body params
	model := models.Brand{
		Name:      brandRequest.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mysql.db.Create(&model)
	result := utility.BrandDetailResponse(model)
	utility.SendSuccessResponse(w, result, http.StatusCreated)
}

func (mysql *Mysql) UpdateBrand(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	brandID, _ := strconv.ParseInt(params.ByName("ID"), 10, 64)

	// run query
	var model models.Brand

	// return not found if record not exist
	if mysql.db.First(&model, brandID).RecordNotFound() {
		utility.SendErrorResponse(w, entity.BrandNotFoundError)
		return
	}

	var brandRequest request.BrandUpdateRequest

	err := json.NewDecoder(r.Body).Decode(&brandRequest)
	if err != nil {
		utility.SendErrorResponse(w, entity.FailedDecodeJSONError)
		return
	}

	// validate body params
	v := validator.New()
	err = v.Struct(brandRequest)

	if err != nil {
		println("error: " + err.Error())
		utility.SendErrorResponse(w, entity.UnprocessableEntityError)
		return
	}

	mysql.db.Model(&model).Updates(brandRequest)
	result := utility.BrandDetailResponse(model)
	utility.SendSuccessResponse(w, result, http.StatusOK)
}

func (mysql *Mysql) DeleteBrand(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	brandID, _ := strconv.ParseInt(params.ByName("ID"), 10, 64)

	// run query
	var model models.Brand

	// return not found if record not exist
	if mysql.db.First(&model, brandID).RecordNotFound() {
		utility.SendErrorResponse(w, entity.BrandNotFoundError)
		return
	}
	mysql.db.Delete(&model)
	w.WriteHeader(http.StatusOK)
}

// private func

func countBrand(mysql *Mysql) int {
	var count int
	mysql.db.Table("brands").Count(&count)
	return count
}
