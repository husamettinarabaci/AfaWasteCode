package main

import (
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/AfatekDevelopers/gps_lib_go/devafatekresult"
)

var debug bool = os.Getenv("DEBUG") == "1"
var appStatus string = "1"

func initStart() {

	logStr("Successfully connected!")
}
func main() {

	initStart()

	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/readiness", readinessHandler)
	http.HandleFunc("/status", status)
	http.ListenAndServe(":80", nil)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func readinessHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func status(w http.ResponseWriter, req *http.Request) {
	var resultVal devafatekresult.ResultType
	if err := req.ParseForm(); err != nil {
		logErr(err)
		return
	}
	opType := req.FormValue("OPTYPE")
	logStr(opType)

	resultVal.Result = "FAIL"
	if opType == "TYPE" {
		resultVal.Result = "TcpTest"
	} else if opType == "APP" {
		if appStatus == "1" {
			resultVal.Result = "OK"
		} else {
			resultVal.Result = "FAIL"
		}
	} else {
		resultVal.Result = "FAIL"
	}
	w.Write(resultVal.ToByte())
}

func logErr(err error) {
	if err != nil {
		sendLogServer("ERR", err.Error())
	}
}

func logStr(value string) {
	if debug {
		sendLogServer("INFO", value)
	}
}

var container string = os.Getenv("CONTAINER_TYPE")

func sendLogServer(logType string, logVal string) {
	data := url.Values{
		"CONTAINER": {container},
		"LOGTYPE":   {logType},
		"LOG":       {logVal},
	}
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	client.PostForm("http://waste-logserver-cluster-ip/log", data)
}
