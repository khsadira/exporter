package exporter_24x7

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func GetMetrics() []byte {
	token, err := getToken()
	if err != nil {
		log.Println(err)
		return ([]byte(""))
	}
	js, err := getJSON(token)
	if err != nil {
		log.Println(err)
		return ([]byte(""))
	}

	var struct24x7 metrics24x7

	json.Unmarshal(js, &struct24x7)

	metrics := toCorrectFormat(struct24x7)
	ret := fromMetricsToString(metrics)

	return ret
}

func getJSON(token string) ([]byte, error) {
	type Scan struct {
		Msg string `json:"message"`
	}

	var scan Scan

	req, err := http.NewRequest("GET", "https://www.site24x7.com/api/reports/performance/195124000074307003?unit_of_time=1&locations=113770000000073133&period=18&report_attribute=RESPONSETIME", nil)
	if err != nil {
		return []byte(""), err
	}
	//	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("Accept", "application/json; version=2.0")
	req.Header.Set("Authorization", "Zoho-oauthtoken "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return []byte(""), err
	}
	defer resp.Body.Close()
	buf, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(buf, &scan)
	if scan.Msg != "success" {
		return []byte(""), errors.New("Json was not well generated")
	}
	return buf, nil
}

func fromMetricsToString(metrics [8]chartData) []byte {
	var ret string

	exp := "24x7exporter"
	HELP := "# HELP " + exp
	TYPE := "# TYPE " + exp

	ret += HELP + "_label of the id\n"
	ret += TYPE + "_label\n"
	for i := 0; i < 8; i++ {
		ret += fmt.Sprintf("%s_label{%d} %s\n", exp, metrics[i].nbStruct, metrics[i].Label)
	}

	ret += HELP + "_max of the id\n"
	ret += TYPE + "_max\n"
	for i := 0; i < 8; i++ {
		ret += fmt.Sprintf("%s_max{%d} %v\n", exp, metrics[i].nbStruct, metrics[i].Max)
	}

	ret += HELP + "_min of the id\n"
	ret += TYPE + "_min\n"
	for i := 0; i < 8; i++ {
		ret += fmt.Sprintf("%s_min{%d} %v\n", exp, metrics[i].nbStruct, metrics[i].Min)
	}

	ret += HELP + "_95_percentile of the id\n"
	ret += TYPE + "_95_percentile\n"
	for i := 0; i < 8; i++ {
		ret += fmt.Sprintf("%s_percentile{%d} %v\n", exp, metrics[i].nbStruct, metrics[i].Percentile)
	}

	ret += HELP + "_average of the id\n"
	ret += TYPE + "_average\n"
	for i := 0; i < 8; i++ {
		ret += fmt.Sprintf("%s_average{%d} %v\n", exp, metrics[i].nbStruct, metrics[i].Average)
	}
	return []byte(ret)
}
