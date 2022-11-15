package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/http"
//	"time"

	"github.com/prometheus/client_golang/prometheus"
//	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func hello(w http.ResponseWriter, req *http.Request) {
	for i := 1; i < 100; i++ {
		fmt.Println(math.Sqrt(rand.Float64()))
	}
	fmt.Fprintf(w, "hello\n")
}

var addr = flag.String("listen-address", ":8000", "The address to listen on for HTTP requests.")

func main() {
	flag.Parse()

	// non-global registry.
	reg := prometheus.NewRegistry()

  http.HandleFunc("/hello", hello)
	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg}))
	log.Fatal(http.ListenAndServe(*addr, nil))
}
