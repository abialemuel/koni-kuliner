package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/koni-kuliner/database"
	"github.com/koni-kuliner/handler"
	"github.com/koni-kuliner/utility"
)

func main() {
	db := database.InitMysql()

	newHandler := handler.NewHandler(db)

	router := httprouter.New()
	router.GET("/healthz", handler.HealthzHandler)
	router.GET("/metrics", handler.MetricHandler)

	// brands resource
	router.GET("/brands", utility.BasicAuth(newHandler.GetBrands))
	router.GET("/brands/:ID", utility.BasicAuth(newHandler.GetBrandDetails))
	router.PATCH("/brands/:ID", utility.BasicAuth(newHandler.UpdateBrand))
	router.POST("/brands", utility.BasicAuth(newHandler.CreateBrand))
	router.DELETE("/brands/:ID", utility.BasicAuth(newHandler.DeleteBrand))

	// customers resource
	router.GET("/customers", utility.BasicAuth(newHandler.GetCustomers))
	router.GET("/customers/:ID", utility.BasicAuth(newHandler.GetCustomerDetails))
	router.PATCH("/customers/:ID", utility.BasicAuth(newHandler.UpdateCustomer))
	router.POST("/customers", utility.BasicAuth(newHandler.CreateCustomer))
	router.DELETE("/customers/:ID", utility.BasicAuth(newHandler.DeleteCustomer))

	// outlets resource
	router.GET("/outlets", utility.BasicAuth(newHandler.GetOutlets))
	router.GET("/outlets/:ID", utility.BasicAuth(newHandler.GetOutletDetails))
	router.PATCH("/outlets/:ID", utility.BasicAuth(newHandler.UpdateOutlet))
	router.POST("/outlets", utility.BasicAuth(newHandler.CreateOutlet))
	router.DELETE("/outlets/:ID", utility.BasicAuth(newHandler.DeleteOutlet))

	// outlet products resource
	router.GET("/outlet-products", utility.BasicAuth(newHandler.GetOutletProducts))
	router.GET("/outlet-products/:ID", utility.BasicAuth(newHandler.GetOutletProductDetails))
	router.PATCH("/outlet-products/:ID", utility.BasicAuth(newHandler.UpdateOutletProduct))
	router.POST("/outlet-products", utility.BasicAuth(newHandler.CreateOutletProduct))
	router.DELETE("/outlet-products/:ID", utility.BasicAuth(newHandler.DeleteOutletProduct))

	// products resource
	router.GET("/products", utility.BasicAuth(newHandler.GetProducts))
	router.GET("/products/:ID", utility.BasicAuth(newHandler.GetProductDetails))
	router.PATCH("/products/:ID", utility.BasicAuth(newHandler.UpdateProduct))
	router.POST("/products", utility.BasicAuth(newHandler.CreateProduct))
	router.DELETE("/products/:ID", utility.BasicAuth(newHandler.DeleteProduct))

	// transactions resource
	router.GET("/transactions", utility.BasicAuth(newHandler.GetTransactions))
	router.GET("/transactions/:ID", utility.BasicAuth(newHandler.GetTransactionDetails))
	router.PATCH("/transactions/:ID", utility.BasicAuth(newHandler.UpdateTransaction))
	router.POST("/transactions", utility.BasicAuth(newHandler.CreateTransaction))

	// cart-itemss resource
	router.GET("/cart-items", utility.BasicAuth(newHandler.GetCartItems))
	router.PATCH("/cart-items/:ID", utility.BasicAuth(newHandler.UpdateCartItem))
	router.POST("/cart-items", utility.BasicAuth(newHandler.CreateCartItem))
	router.DELETE("/cart-items/:ID", utility.BasicAuth(newHandler.DeleteCartItem))

	// sellers resource
	router.GET("/sellers", utility.BasicAuth(newHandler.GetSellers))
	router.GET("/sellers/:ID", utility.BasicAuth(newHandler.GetSellerDetails))
	router.PATCH("/sellers/:ID", utility.BasicAuth(newHandler.UpdateSeller))
	router.POST("/sellers", utility.BasicAuth(newHandler.CreateSeller))
	router.DELETE("/sellers/:ID", utility.BasicAuth(newHandler.DeleteSeller))

	fmt.Println("Connected to port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
