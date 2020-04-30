package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"github.com/koni-kuliner/handler"
)

func main() {

	router := httprouter.New()
	router.GET("/healthz", handler.HealthzHandler)
	router.GET("/metrics", handler.MetricHandler)
	fmt.Println("Connected to port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
