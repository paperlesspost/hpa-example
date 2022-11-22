package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"strconv"
)

var addr = flag.String("listen-address", ":8000", "The address to listen on for HTTP requests.")

// health check fn
func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

// helper types to decode json from settings svc
type ValueContainer struct {
	Value int `json:"value"`
}

type ExampleMetric struct {
	ExampleMetric ValueContainer `json:"example_metric"`
}

type ExampleMetricResponse struct {
	Data ExampleMetric `json:"data"`
}

func getMetric(w http.ResponseWriter, req *http.Request) {
	// pull metric from Settings service
	resp, err := http.Get("http://earth.ppstaging.net/flyer/api/settings/hpa-example")
	// resp, err := http.Get("http://settings/settings/hpa-example")

	if err != nil {
		fmt.Println("error fetching from settings service")
		return
	}

	defer resp.Body.Close()

	rawJson, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("error reading response")
		return
	}

	var metricResponse ExampleMetricResponse
	err = json.Unmarshal(rawJson, &metricResponse)
	if err != nil {
		fmt.Println("error unmarshaling response")
		return
	}
	
	metricsText := 
`# HELP example_metric Custom arbitrary value to test example HPA
# TYPE example_metric gauge
example_metric ` + strconv.Itoa(metricResponse.Data.ExampleMetric.Value) + "\n"

	fmt.Fprintf(w, metricsText)
}

func main() {
	flag.Parse()

  	http.HandleFunc("/hello", hello)
	http.HandleFunc("/metrics", getMetric)

	fmt.Println("example-hpa app running on port " + *addr)

	log.Fatal(http.ListenAndServe(*addr, nil))
}
