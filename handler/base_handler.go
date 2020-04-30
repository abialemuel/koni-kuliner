package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func HealthzHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("OK"))
}

func MetricHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	promhttp.Handler().ServeHTTP(w, r)
}
