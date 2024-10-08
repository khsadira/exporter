package exporter_observatory

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"
)

type Metrics struct {
	Target   string `json:"target"`
	Grade    string `json:"grade"`
	TestPass int    `json:"tests_passed"`
	Score    int    `json:"score"`
}

const (
	helpGrade = "# HELP exporter_grade Grade representation of score, A+=12, A=11, A-=10, B+=9, B=8, B-=7, C+=6, C=5, C-=4, D+=3, D=2, D-=1, F=0"
	typeGrade = "# TYPE exporter_grade gauge"
	helpScore = "# HELP exporter_score Http score"
	typeScore = "# TYPE exporter_score gauge"
	helpTest  = "# HELP exporter_test_passed Number of test passed"
	typeTest  = "# TYPE exporter_test_passed gauge"
)

func GetMetrics(w http.ResponseWriter, r *http.Request) []byte {
	var resp string
	var metrics []Metrics
	hosts, ok := r.URL.Query()["host"]

	if !ok || len(hosts[0]) < 0 {
		log.Println("Url Param 'host' is missing")
		resp = "No host query found"
	} else {
		targets := strings.Split(hosts[0], ",")
		targets = reworkURL(targets)
		for _, target := range targets {
			metrics = append(metrics, newMetrics(target))
		}
		resp = retAnswer(metrics)
	}
	return ([]byte(resp))
}

func retAnswer(metrics []Metrics) string {
	var grade string
	var score string
	var test string

	for _, metric := range metrics {
		grade += fmt.Sprintf("exporter_grade{host=%s} %d\n", metric.Target, gradeLetterToInt(metric.Grade))
		score += fmt.Sprintf("exporter_score{host=%s} %d\n", metric.Target, metric.Score)
		test += fmt.Sprintf("exporter_test{host=%s} %d\n", metric.Target, metric.TestPass)
	}

	resp := fmt.Sprintf("%s\n%s\n%s%s\n%s\n%s%s\n%s\n%s", helpGrade, typeGrade, grade, helpScore, typeScore, score, helpTest, typeTest, test)
	return resp
}

func gradeLetterToInt(str string) int {
	mapping := map[string]int{
		"A+": 12,
		"A":  11,
		"A-": 10,
		"B+": 9,
		"B":  8,
		"B-": 7,
		"C+": 6,
		"C":  5,
		"C-": 4,
		"D+": 3,
		"D":  2,
		"D-": 1,
		"F":  0,
	}

	str = strings.ToUpper(str)
	l, ok := mapping[str]
	if !ok {
		l = 0
	}
	return l
}

func reworkURL(str []string) []string {
	var ret []string
	for _, targetURL := range str {
		targetURL = strings.TrimPrefix(targetURL, "https://")
		targetURL = strings.TrimPrefix(targetURL, "http://")
		targetURL = "https://" + targetURL
		val, _ := url.Parse(targetURL)
		targetURL = val.Hostname()
		_, err := net.Dial("tcp", targetURL+":http")
		if err != nil {
			log.Println("unreachable, error:", err)
		} else {
			ret = append(ret, targetURL)
		}
	}
	return ret
}
