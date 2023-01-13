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
}

var alertCounter = prometheus.NewCounter(
   prometheus.CounterOpts{
       Name: "alert_counter",
       Help: "Should alert something",
   },
)

func alert(w http.ResponseWriter, req *http.Request) {
   alertCounter.Inc()
   fmt.Fprintf(w, "alert inc()")
}

func pprint(w http.ResponseWriter, req *http.Request) {
   b, err := io.ReadAll(req.Body)
   if err != nil {
       panic(err)
   }

   fmt.Printf("%s", b)
}

func main() {
   prometheus.MustRegister(pingCounter)
   prometheus.MustRegister(alertCounter)

   http.HandleFunc("/ping", ping)
   http.HandleFunc("/alert", alert)
   http.HandleFunc("/print", pprint)
   http.Handle("/metrics", promhttp.Handler())
   http.ListenAndServe(":8090", nil)
}