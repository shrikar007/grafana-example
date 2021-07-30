package middlerware

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"net/http"
	"strconv"
)

type Response struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func NewResponseWriter(w http.ResponseWriter) *Response {
	return &Response{http.StatusOK, ""}
}

var TotalRequests = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total_custom",
		Help: "Number of get requests.",
	},
	[]string{"path"},
)

var ResponseStatus = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "response_status_custom",
		Help: "Status of HTTP response",
	},
	[]string{"status"},
)

var HttpDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
	Name: "http_response_time_seconds_custom",
	Help: "Duration of HTTP requests.",
}, []string{"path"})

func PrometheusMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path:=r.URL.Path
		timer := prometheus.NewTimer(HttpDuration.WithLabelValues(path))
		rw := NewResponseWriter(w)
		next.ServeHTTP(w, r)

		statusCode := rw.StatusCode

		ResponseStatus.WithLabelValues(strconv.Itoa(statusCode)).Inc()
		TotalRequests.WithLabelValues(path).Inc()
		timer.ObserveDuration()
	})
}
