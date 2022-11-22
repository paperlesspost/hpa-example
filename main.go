package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/http"
	"runtime"
	"strconv"
	"time"
	"github.com/prometheus/client_golang/prometheus"
  "github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	addr = flag.String("listen-address", ":8000", "The address to listen on for HTTP requests.")

	totalRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Number of get requests.",
		},
		[]string{"path"},
	)

	responseStatus = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "response_status",
			Help: "Status of HTTP response",
		},
		[]string{"status"},
	)

	httpDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name: "http_response_time_seconds",
		Help: "Duration of HTTP requests.",
	}, []string{"path"})


)


func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func cpuStress(w http.ResponseWriter, req *http.Request) {
	var x = float64(1)
	for i := float64(1); i <= 10000; i++ {
    x = i * rand.Float64()
		fmt.Println(math.Sqrt(float64(x)))
		responseStatus.WithLabelValues(strconv.Itoa(statusCode)).Inc()
		totalRequests.WithLabelValues(path).Inc()
  }
  fmt.Println("cpu stress test complete")
}

func metr (){
	prometheus.Register(totalRequests)
	prometheus.Register(responseStatus)
	prometheus.Register(httpDuration)

	prometheus.DefaultGatherer
	promhttp.HandlerOpts{}
}


func main() {
	runtime.GOMAXPROCS(1)
	flag.Parse()



  http.HandleFunc("/hello", hello)
	http.Handle("/metrics", metr)
	http.HandleFunc("/load", cpuStress)


  log.Fatal(http.ListenAndServe(*addr, nil))
	time.Sleep(1 * time.Nanosecond) // Wait for goroutines to finish
}
