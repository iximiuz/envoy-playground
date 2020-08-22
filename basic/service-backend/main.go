package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var FAIL_PCT int

var (
	requestCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "backend_requests_total",
			Help: "The total number of requests received by the backend service",
		},
		[]string{"status"},
	)
)

func handler(w http.ResponseWriter, r *http.Request) {
	ts := time.Now().Format(time.RFC850)
	if rand.Intn(100) >= FAIL_PCT {
		fmt.Fprintf(w, "Greetings from backend! Time is %v", ts)
		requestCounter.WithLabelValues("2xx").Inc()
	} else {
		http.Error(w, fmt.Sprintf("Ooops... Random error. Time is %v", ts),
			http.StatusInternalServerError)
		requestCounter.WithLabelValues("5xx").Inc()
	}
}

func main() {
	n, err := strconv.Atoi(os.Getenv("FAIL_PCT"))
	if err != nil {
		log.Fatal("Can not parse FAIL_PCT env")
	}
	FAIL_PCT = n

	prometheus.MustRegister(requestCounter)
	http.Handle("/metrics", promhttp.Handler())
	go http.ListenAndServe(":"+os.Getenv("PORT_METRICS"), nil)

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT_API"), nil))
}
