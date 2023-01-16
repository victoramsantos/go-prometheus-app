package main

import (
   "fmt"
   "net/http"
   "io"

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
   fmt.Println(fmt.Sprintf("%s", b))
}

func version(w http.ResponseWriter, req *http.Request) {
  version := "1.0.3"
  fmt.Fprintf(w, "%s", version)
  fmt.Println(version)
}

func main() {
   prometheus.MustRegister(pingCounter)
   prometheus.MustRegister(alertGauge)

   http.HandleFunc("/ping", ping)
   http.HandleFunc("/alert", alert)
   http.HandleFunc("/resetalert", resetAlert)
   http.HandleFunc("/print", pprint)
   http.HandleFunc("/version", version)
   http.Handle("/metrics", promhttp.Handler())
   http.ListenAndServe(":8090", nil)
}