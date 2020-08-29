package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var ERROR_RATE int

var (
	requestCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "service_b_requests_total",
			Help: "The total number of requests received by Service B.",
		},
		[]string{"status"},
	)
)

func handler(w http.ResponseWriter, r *http.Request) {
	if rand.Intn(100) >= ERROR_RATE {
		fmt.Fprintln(w, "Service B: Yay! nounce", rand.Uint32())
		requestCounter.WithLabelValues("2xx").Inc()
	} else {
		http.Error(w, fmt.Sprintf("Service B: Ooops... nounce %v", rand.Uint32()),
			http.StatusInternalServerError)
		requestCounter.WithLabelValues("5xx").Inc()
	}
}

func main() {
	n, err := strconv.Atoi(os.Getenv("ERROR_RATE"))
	if err != nil {
		log.Fatal("Can not parse ERROR_RATE env")
	}
	ERROR_RATE = n

	prometheus.MustRegister(requestCounter)
	http.Handle("/metrics", promhttp.Handler())
	go http.ListenAndServe(":"+os.Getenv("METRICS_PORT"), nil)

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("SERVICE_PORT"), nil))
}
