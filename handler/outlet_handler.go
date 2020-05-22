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

func (mysql *Mysql) GetOutlets(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var filteredArgs []interface{}

	// filter query params
	filter := utility.Filter(r, []string{"id", "name", "offset", "limit"})

	// build query
	query := "SELECT * FROM outlets WHERE 1=1"
	query, filteredArgs = utility.AppendQuery(query, filter)

	// run query
	var model []models.Outlet
	mysql.db.Raw(query, filteredArgs...).Scan(&model)
	result := utility.OutletResponse(model)
	utility.SendSuccessResponseWithLimitAndOffset(w, result, http.StatusOK, filter, CountQuery(mysql, query, filteredArgs))
}

func (mysql *Mysql) GetOutletDetails(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	outletID, _ := strconv.ParseInt(params.ByName("ID"), 10, 64)

	// run query
	var model models.Outlet

	// return not found if record not exist
	if mysql.db.First(&model, outletID).RecordNotFound() {
		utility.SendErrorResponse(w, entity.OutletNotFoundError)
		return
	}
	mysql.db.First(&model, outletID)
	result := utility.OutletDetailResponse(model)
	utility.SendSuccessResponse(w, result, http.StatusOK)
}

func (mysql *Mysql) CreateOutlet(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// decode params
	var outletRequest request.OutletCreateRequest

	err := json.NewDecoder(r.Body).Decode(&outletRequest)
	if err != nil {
		utility.SendErrorResponse(w, entity.FailedDecodeJSONError)
		return
	}

	// validate body params
	v := validator.New()
	err = v.Struct(outletRequest)

	if err != nil {
		println("error: " + err.Error())
		utility.SendErrorResponse(w, entity.UnprocessableEntityError)
		return
	}

	// assign body params
	model := models.Outlet{
		Name:      outletRequest.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mysql.db.Create(&model)
	result := utility.OutletDetailResponse(model)
	utility.SendSuccessResponse(w, result, http.StatusCreated)
}

func (mysql *Mysql) UpdateOutlet(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	outletID, _ := strconv.ParseInt(params.ByName("ID"), 10, 64)

	// run query
	var model models.Outlet

	// return not found if record not exist
	if mysql.db.First(&model, outletID).RecordNotFound() {
		utility.SendErrorResponse(w, entity.OutletNotFoundError)
		return
	}

	var outletRequest request.OutletUpdateRequest

	err := json.NewDecoder(r.Body).Decode(&outletRequest)
	if err != nil {
		utility.SendErrorResponse(w, entity.FailedDecodeJSONError)
		return
	}

	// validate body params
	v := validator.New()
	err = v.Struct(outletRequest)

	if err != nil {
		println("error: " + err.Error())
		utility.SendErrorResponse(w, entity.UnprocessableEntityError)
		return
	}

	mysql.db.Model(&model).Updates(outletRequest)
	result := utility.OutletDetailResponse(model)
	utility.SendSuccessResponse(w, result, http.StatusOK)
}

func (mysql *Mysql) DeleteOutlet(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	outletID, _ := strconv.ParseInt(params.ByName("ID"), 10, 64)

	// run query
	var model models.Outlet

	// return not found if record not exist
	if mysql.db.First(&model, outletID).RecordNotFound() {
		utility.SendErrorResponse(w, entity.OutletNotFoundError)
		return
	}
	mysql.db.Delete(&model)
	w.WriteHeader(http.StatusOK)
}

// private func
