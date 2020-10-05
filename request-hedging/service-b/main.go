package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	requestDuration = prometheus.NewSummary(prometheus.SummaryOpts{
		Name:       "service_b_request_request_duration",
		Help:       "Summary for the duration of Service B requests.",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	})
)

func handler(w http.ResponseWriter, r *http.Request) {
	timer := prometheus.NewTimer(requestDuration)
	defer timer.ObserveDuration()

	chance := rand.Intn(1000)
	if chance < 100 {
		time.Sleep(10 * time.Millisecond)
	}

	fmt.Fprintln(w, "Service B: Yay! nounce", rand.Uint32())
	log.Println("HTTP 200", r.Method, r.URL, r.RemoteAddr)
}

func main() {
	prometheus.MustRegister(requestDuration)
	http.Handle("/metrics", promhttp.Handler())
	go http.ListenAndServe(":"+os.Getenv("METRICS_PORT"), nil)

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("SERVICE_PORT"), nil))
}
