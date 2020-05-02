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

func (conn *Mysql) GetProducts(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	model := []models.Product{}
	conn.db.Find(&model)
	result := utility.ProductResponse(model)
	utility.SendSuccessResponse(w, result, http.StatusOK)
}
