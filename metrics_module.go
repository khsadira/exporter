package main

import (
	"github.com/khsadira/exporter/exporter_24x7"
	"github.com/khsadira/exporter/exporter_observatory"
	"net/http"
)

func metricsPage(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte(`<html>
<head><title>Exporter</title></head>
<body>
<h1>Observatory Exporter</h1>
<p><a href='/metrics/observatory/'>Observatory</a></p>
<p><a href='/metrics/24x7/'>24x7</a></p>
</body>
</html>`))
}

func getObsMetrics(w http.ResponseWriter, r *http.Request) {
	metrics := exporter_observatory.GetMetrics(w, r)
	w.Write(metrics)
}

func get24x7Metrics(w http.ResponseWriter, r *http.Request) {
	metrics := exporter_24x7.GetMetrics()
	w.Write(metrics)
}
