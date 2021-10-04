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
	opType := req.FormValue(WasteLibrary.HTTP_OPTYPE)
	logStr(opType)
	resultVal.Result = WasteLibrary.RESULT_FAIL
	if opType == WasteLibrary.OPTYPE_TYPE {
		resultVal.Result = "WasteLogServer"
	} else if opType == WasteLibrary.OPTYPE_APP {
		resultVal.Result = WasteLibrary.RESULT_OK
	} else {
		resultVal.Result = WasteLibrary.RESULT_FAIL
	}
	w.Write(resultVal.ToByte())
}

func log(w http.ResponseWriter, req *http.Request) {
	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_OK
	if err := req.ParseForm(); err != nil {
		logErr(err)
		return
	}
	container := req.FormValue(WasteLibrary.LOGGER_CONTAINER)
	logType := req.FormValue(WasteLibrary.LOGGER_LOGTYPE)
	funcVal := req.FormValue(WasteLibrary.LOGGER_FUNC)
	logVal := req.FormValue(WasteLibrary.LOGGER_LOG)
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
