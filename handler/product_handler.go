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

func (mysql *Mysql) GetProducts(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var filteredArgs []interface{}

	// filter query params
	filter := utility.Filter(r, []string{"id", "name", "offset", "limit"})

	// build query
	query := "SELECT * FROM products WHERE 1=1"
	query, filteredArgs = utility.AppendQuery(query, filter)

	// run query
	var model []models.Product
	mysql.db.Raw(query, filteredArgs...).Scan(&model)
	result := utility.ProductResponse(model)
	utility.SendSuccessResponseWithLimitAndOffset(w, result, http.StatusOK, filter, CountProduct(mysql))
}

func (mysql *Mysql) GetProductDetails(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	productID, _ := strconv.ParseInt(params.ByName("ID"), 10, 64)

	// run query
	var model models.Product

	if mysql.db.First(&model, productID).RecordNotFound() {
		utility.SendErrorResponse(w, entity.ProductNotFoundError)
		return
	}

	result := utility.ProductDetailResponse(model)
	utility.SendSuccessResponse(w, result, http.StatusOK)
}

func (mysql *Mysql) CreateProduct(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// assign params
	var productRequest request.ProductCreateRequest

	// decode params
	err := json.NewDecoder(r.Body).Decode(&productRequest)
	if err != nil {
		utility.SendErrorResponse(w, entity.FailedDecodeJSONError)
		return
	}

	// validate body params
	v := validator.New()
	err = v.Struct(productRequest)

	if err != nil {
		println("error: " + err.Error())
		utility.SendErrorResponse(w, entity.UnprocessableEntityError)
		return
	}

	// assign body params
	model := models.Product{
		Name:      productRequest.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mysql.db.Create(&model)
	result := utility.ProductDetailResponse(model)
	utility.SendSuccessResponse(w, result, http.StatusCreated)
	// return err
}

func CountProduct(mysql *Mysql) int {
	var count int
	mysql.db.Table("products").Count(&count)
	return count
}
