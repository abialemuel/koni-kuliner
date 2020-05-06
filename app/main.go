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

	// products resource
	router.GET("/products", utility.BasicAuth(newHandler.GetProducts))
	router.GET("/products/:ID", utility.BasicAuth(newHandler.GetProductDetails))
	router.PATCH("/products/:ID", utility.BasicAuth(newHandler.UpdateProduct))
	router.POST("/products", utility.BasicAuth(newHandler.CreateProduct))
	router.DELETE("/products/:ID", utility.BasicAuth(newHandler.DeleteProduct))

	// outlets resource
	router.GET("/outlets", utility.BasicAuth(newHandler.GetOutlets))
	router.GET("/outlets/:ID", utility.BasicAuth(newHandler.GetOutletDetails))
	router.PATCH("/outlets/:ID", utility.BasicAuth(newHandler.UpdateOutlet))
	router.POST("/outlets", utility.BasicAuth(newHandler.CreateOutlet))
	router.DELETE("/outlets/:ID", utility.BasicAuth(newHandler.DeleteOutlet))

	fmt.Println("Connected to port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
