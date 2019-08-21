package exporter_observatory

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

func newMetrics(target string) Metrics {
	var metrics Metrics

	json, err := getJSONID(target, 0)
	if err != nil {
		return metrics
	}

	metrics = createMetrics(json)
	metrics.Target = target
	return metrics
}

func createMetrics(js []byte) Metrics {
	var metrics Metrics

	json.Unmarshal(js, &metrics)
	return metrics
}

type scan struct {
	State string `json:"state"`
}

func getJSONID(target string, loop int) ([]byte, error) {
	var scan scan
	resp, err := http.Post("https://http-observatory.security.mozilla.org/api/v1/analyze?host="+target+"&rescan=true", "", nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(http.StatusText(resp.StatusCode))
	}
	buf, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(buf, &scan)
	if scan.State != "FINISHED" {
		if loop > 30 {
			return []byte(""), errors.New("Loop error while touching the server")
		}
		time.Sleep(time.Second / 2)
		return getJSONID(target, loop+1)
	}
	return buf, nil
}
