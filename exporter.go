package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/metrics/", metricsPage)
	http.HandleFunc("/metrics/observatory/", getObsMetrics)
	http.HandleFunc("/metrics/24x7/", get24x7Metrics)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`<html>
<head><title>Exporter</title></head>
			 <body>
			 <h1>Exporter</h1>
			 <p><a href='/metrics/'>Metrics</a></p>
			 </body>
			 </html>`))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
