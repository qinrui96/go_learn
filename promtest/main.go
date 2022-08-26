package main

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func main() {
	http.Handle("/metrics", promhttp.Handler())
	_ = http.ListenAndServe("localhost:8080", nil)
}
