package main

import (
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Define custom Prometheus metrics
var (
	reqCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests received",
		},
	)

	durationHistogram = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Histogram of request processing time",
			Buckets: prometheus.DefBuckets,
		},
	)
)

func init() {
	// Register metrics with Prometheus
	prometheus.MustRegister(reqCounter)
	prometheus.MustRegister(durationHistogram)
}

func handler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	reqCounter.Inc() // Increment request counter

	time.Sleep(time.Millisecond * 500) // Simulate request processing delay

	duration := time.Since(start).Seconds()
	durationHistogram.Observe(duration) // Observe request duration

	w.Write([]byte("Hello, Prometheus!"))
}

func main() {
	http.Handle("/metrics", promhttp.Handler()) // Expose Prometheus metrics endpoint
	http.HandleFunc("/", handler)

	log.Println("Server is running on :9090")
	log.Fatal(http.ListenAndServe(":9090", nil))
}
