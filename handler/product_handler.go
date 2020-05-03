package handler

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
	"github.com/koni-kuliner/models"
	"github.com/koni-kuliner/utility"
)

type Mysql struct {
	db *gorm.DB
}

func NewProductHandler(db *gorm.DB) *Mysql {
	return &Mysql{
		db: db,
	}
}

func (mysql *Mysql) GetProducts(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var filteredArgs []interface{}

	// filter query params
	filter := utility.Filter(r, []string{"offset", "limit"})

	// build query
	query := "SELECT * FROM products WHERE 1=1"
	query, filteredArgs = utility.AppendQuery(query, filter)

	// run query
	var model []models.Product
	mysql.db.Raw(query, filteredArgs...).Scan(&model)
	result := utility.ProductResponse(model)
	utility.SendSuccessResponseWithLimitAndOffset(w, result, http.StatusOK, filter, Count(mysql))
}

func Count(mysql *Mysql) int {
	var count int
	mysql.db.Table("products").Count(&count)
	return count
}
