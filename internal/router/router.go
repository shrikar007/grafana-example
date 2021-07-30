package router

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"grafana-example/internal/middlerware"
	"net/http"
)

func Init() *chi.Mux {
	var resp middlerware.Response
	router := chi.NewRouter()
	prometheus.Register(middlerware.TotalRequests)
	prometheus.Register(middlerware.ResponseStatus)
	prometheus.Register(middlerware.HttpDuration)

	router.Use(middlerware.PrometheusMiddleware)
	router.Handle("/metrics", promhttp.Handler())
	router.Get("/check", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("content-type", "application/json")
		w := json.NewEncoder(writer)
		resp = middlerware.Response{
			StatusCode: http.StatusOK,
			Message:    "Success",
		}
		w.Encode(resp)
	})
	return router
}
