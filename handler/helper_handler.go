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

// outletProducts

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

// transactions

func GetAllDetailRelationTransaction(mysql *Mysql, transaction *[]models.Transaction) {
	for i, m := range *transaction {
		var customer models.Customer
		mysql.db.First(&customer, m.CustomerID)

		var cartItems []models.CartItem
		mysql.db.Where("transaction_id = ?", m.ID).Find(&cartItems)

		(*transaction)[i].Customer = customer
		(*transaction)[i].CartItems = cartItems
	}
}

func GetSingleDetailRelationTransaction(mysql *Mysql, transaction *models.Transaction) {
	var customer models.Customer
	mysql.db.First(&customer, transaction.CustomerID)

	var cartItems []models.CartItem
	mysql.db.Where("transaction_id = ?", transaction.ID).Find(&cartItems)

	GetAllDetailRelationCartItem(mysql, &cartItems)
	transaction.Customer = customer
	transaction.CartItems = cartItems
}

// cartItems

func GetAllDetailRelationCartItem(mysql *Mysql, cartItem *[]models.CartItem) {
	for i, m := range *cartItem {
		var outletProduct models.OutletProduct
		mysql.db.First(&outletProduct, m.OutletProductID)

		GetSingleDetailRelationOutletProduct(mysql, &outletProduct)
		(*cartItem)[i].OutletProduct = outletProduct
	}
}

func GetSingleDetailRelationCartItem(mysql *Mysql, cartItem *models.CartItem) {
	var outletProduct models.OutletProduct
	mysql.db.First(&outletProduct, cartItem.OutletProductID)

	GetSingleDetailRelationOutletProduct(mysql, &outletProduct)
	cartItem.OutletProduct = outletProduct
}
