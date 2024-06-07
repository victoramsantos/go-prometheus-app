package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var pingCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "ping_request_count",
		Help: "No of request handled by Ping handler",
	},
)

func ping(w http.ResponseWriter, req *http.Request) {
	pingCounter.Inc()
	fmt.Fprintf(w, "pong")
	fmt.Println("pong")
}

var alertGauge = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Name: "alert_counter",
		Help: "Should alert something",
	},
)

var randomGauge = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Name: "random_gauge",
		Help: "Return a random value",
	},
)

var staticCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "static_counter",
		Help: "Returns a static value",
	},
)

var requestDuration = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name:    "http_request_duration_seconds",
		Help:    "A histogram of latencies for HTTP requests.",
		Buckets: []float64{.5, .7, .9, 1, 1.5, 2, 2.5, 5},
	},
	[]string{"id"},
)

func alert(w http.ResponseWriter, req *http.Request) {
	alertGauge.Inc()
	fmt.Fprintf(w, "alert inc()")
	fmt.Println("alert inc()")
}

func resetAlert(w http.ResponseWriter, req *http.Request) {
	alertGauge.Set(0)
	fmt.Fprintf(w, "alert reset")
	fmt.Println("alert reset")
}

func pprint(w http.ResponseWriter, req *http.Request) {
	b, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "%s", b)
	fmt.Printf("%s\n", b)
}

func version(w http.ResponseWriter, req *http.Request) {
	version := "1.0.3"
	fmt.Fprintf(w, "%s", version)
	fmt.Println(version)
}

func histogram(w http.ResponseWriter, req *http.Request) {
	var id = rand.Intn(5)
	timer := prometheus.NewTimer(requestDuration.WithLabelValues(fmt.Sprint(id)))
	defer timer.ObserveDuration()

	var sleep = 0.1 + (rand.Float64() * 4.9)
	time.Sleep(time.Duration(sleep * float64(time.Second)))

	fmt.Fprintf(w, "id=%d, sleep=%f", id, sleep)
	fmt.Printf("id=%d, sleep=%f\n", id, sleep)
}

func metricsHandle() http.Handler {
	randomGauge.Set(1 + rand.Float64()*(10-1))
	return promhttp.Handler()
}

func initRegisters() {
	prometheus.MustRegister(pingCounter)
	prometheus.MustRegister(alertGauge)
	prometheus.MustRegister(staticCounter)
	prometheus.MustRegister(randomGauge)
	prometheus.MustRegister(requestDuration)
}

func initStatics() {
	staticCounter.Inc()
}

func initHandlers() {
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/alert", alert)
	http.HandleFunc("/resetalert", resetAlert)
	http.HandleFunc("/print", pprint)
	http.HandleFunc("/version", version)
	http.HandleFunc("/histogram", histogram)

	metrics_path := os.Getenv("METRICS_PATH")
	if metrics_path == "" {
		metrics_path = "/metrics"
	}
	http.Handle(metrics_path, metricsHandle())
}

func main() {
	initRegisters()
	initStatics()
	initHandlers()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8090"
	}

	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
