package handler

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
	"github.com/koni-kuliner/models"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Mysql struct {
	db *gorm.DB
}

func NewHandler(db *gorm.DB) *Mysql {
	return &Mysql{
		db: db,
	}
}

func HealthzHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("OK"))
}

func MetricHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	promhttp.Handler().ServeHTTP(w, r)
}

func GetSingleDetailRelationOutletProduct(mysql *Mysql, outletProduct *models.OutletProduct) {
	var product models.Product
	mysql.db.First(&product, outletProduct.ProductID)

	var outlet models.Outlet
	mysql.db.First(&outlet, outletProduct.OutletID)

	outletProduct.Product = product
	outletProduct.Outlet = outlet
}

func GetSingleDetailRelationProduct(mysql *Mysql, product *models.Product) {
	var brand models.Brand
	mysql.db.First(&brand, product.BrandID)

	product.Brand = brand
}
