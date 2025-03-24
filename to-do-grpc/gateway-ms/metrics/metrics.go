package metrics

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	RequestCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests received",
		},
		[]string{"method", "path", "status_code"},
	)

	RequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Histogram of HTTP request durations",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "path", "status_code"},
	)
)

type statusCodeResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *statusCodeResponseWriter) WriterHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}

func GeneralMetricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/metrics" {
            next.ServeHTTP(w, r)
            return
        }

		scrw := statusCodeResponseWriter{ResponseWriter: w}

		start := time.Now()
		next.ServeHTTP(scrw, r)
		duration := time.Since(start).Seconds()

		RequestCounter.WithLabelValues(r.Method, r.URL.Path, http.StatusText(scrw.statusCode)).Inc()
		RequestDuration.WithLabelValues(r.Method, r.URL.Path, http.StatusText(scrw.statusCode)).Observe(duration)
	})
}

func init() {
	prometheus.MustRegister(RequestCounter)
	prometheus.MustRegister(RequestDuration)
}