package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/http"
	"runtime"
	"time"

	"github.com/prometheus/client_golang/prometheus"
//	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var addr = flag.String("listen-address", ":8000", "The address to listen on for HTTP requests.")

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func cpuStress(w http.ResponseWriter, req *http.Request) {
	var x = float64(1)
	for i := float64(1); i <= 10000000; i++ {
    x = i * rand.Float64()
		fmt.Println(math.Sqrt(float64(x)))
  }
  fmt.Println("cpu load done")
}


func main() {
	runtime.GOMAXPROCS(1)

	flag.Parse()

	// non-global registry.
	reg := prometheus.NewRegistry()

  http.HandleFunc("/hello", hello)
	http.HandleFunc("/load", cpuStress)
	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg}))
	log.Fatal(http.ListenAndServe(*addr, nil))

	time.Sleep(1 * time.Nanosecond) // Wait for goroutines to finish
}
