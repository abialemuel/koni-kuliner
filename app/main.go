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
	// models.DBMigrate(db)
	productHandler := handler.NewProductHandler(db)

	router := httprouter.New()
	router.GET("/healthz", handler.HealthzHandler)
	router.GET("/metrics", handler.MetricHandler)

	// products resource
	router.GET("/products", utility.BasicAuth(productHandler.GetProducts))

	fmt.Println("Connected to port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
