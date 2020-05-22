package handler

import (
	"strings"

	"github.com/koni-kuliner/models"
)

func CountQuery(mysql *Mysql, query string, filteredArgs []interface{}) int {
	var count int
	query = strings.TrimSuffix(query, " LIMIT ? OFFSET ?")
	query = strings.Replace(query, "*", "COUNT(*)", 1)
	filteredArgs = filteredArgs[:len(filteredArgs)-2]
	mysql.db.Raw(query, filteredArgs...).Count(&count)
	return count
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

func GetSingleDetailRelationTransaction(mysql *Mysql, transaction *models.Transaction) {
	var customer models.Customer
	mysql.db.First(&customer, transaction.CustomerID)

	transaction.Customer = customer
}
