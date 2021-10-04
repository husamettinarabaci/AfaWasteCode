package main

import (
	"fmt"
	"net/http"
	"time"

	_ "github.com/lib/pq"

	"github.com/devafatek/WasteLibrary"
)

func main() {

	http.HandleFunc("/health", WasteLibrary.HealthHandler)
	http.HandleFunc("/readiness", WasteLibrary.ReadinessHandler)
	http.HandleFunc("/status", status)
	http.HandleFunc("/log", log)
	http.ListenAndServe(":80", nil)
}

func status(w http.ResponseWriter, req *http.Request) {
	var resultVal WasteLibrary.ResultType
	if err := req.ParseForm(); err != nil {
		logErr(err)
		return
	}
	opType := req.FormValue("OPTYPE")
	logStr(opType)
	resultVal.Result = "FAIL"
	if opType == "TYPE" {
		resultVal.Result = "WasteLogServer"
	} else if opType == "APP" {
		resultVal.Result = "OK"
	} else {
		resultVal.Result = "FAIL"
	}
	w.Write(resultVal.ToByte())
}

func log(w http.ResponseWriter, req *http.Request) {
	var resultVal WasteLibrary.ResultType
	resultVal.Result = "OK"
	if err := req.ParseForm(); err != nil {
		logErr(err)
		return
	}
	container := req.FormValue("CONTAINER")
	logType := req.FormValue("LOGTYPE")
	funcVal := req.FormValue("FUNC")
	logVal := req.FormValue("LOG")
	logStr("Time : " + time.Now().String() + " - Container : " + container + " - LogType : " + logType + " - Func : " + funcVal + " - Log : " + logVal + " - IP : " + req.RemoteAddr)
	w.Write(resultVal.ToByte())
}

func logErr(err error) {
	if err != nil {
		logStr(err.Error())
	}
}

func logStr(value string) {
	fmt.Println(value)
}
