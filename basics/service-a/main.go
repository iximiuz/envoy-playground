package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var requestCounter = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "service_a_requests_total",
		Help: "The total number of requests received by Service A.",
	},
	[]string{"status"},
)

var UPSTREAM_URL = "http://b.service"

func httpGet(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("HTTP %v - %v", resp.StatusCode, string(body))
	}

	return string(body), nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	resp, err := httpGet(UPSTREAM_URL)
	if err == nil {
		fmt.Fprintln(w, "Service A: upstream responded with:", resp)
		requestCounter.WithLabelValues("2xx").Inc()
	} else {
		http.Error(w, fmt.Sprintf("Service A: upstream failed with: %v", err.Error()),
			http.StatusInternalServerError)
		requestCounter.WithLabelValues("5xx").Inc()
	}
}

func main() {
	prometheus.MustRegister(requestCounter)
	http.Handle("/metrics", promhttp.Handler())
	go http.ListenAndServe(":"+os.Getenv("PORT_METRICS"), nil)

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT_API"), nil))
}
