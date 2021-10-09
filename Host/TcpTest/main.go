package main

import (
	"net/http"

	"github.com/devafatek/WasteLibrary"
)

func initStart() {

	WasteLibrary.LogStr("Successfully connected!")
}
func main() {

	initStart()

	http.HandleFunc("/health", WasteLibrary.HealthHandler)
	http.HandleFunc("/readiness", WasteLibrary.ReadinessHandler)
	http.HandleFunc("/status", WasteLibrary.StatusHandler)
	http.HandleFunc("/status1", statusHandler1)
	http.HandleFunc("/status2", statusHandler2)
	http.ListenAndServe(":80", nil)
}

func statusHandler1(w http.ResponseWriter, req *http.Request) {
	if WasteLibrary.AllowCors {

		w.Header().Set("Access-Control-Allow-Origin", "*")
	}
	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_OK
	WasteLibrary.LogStr("status 1")
	w.Write(resultVal.ToByte())
}

func statusHandler2(w http.ResponseWriter, req *http.Request) {
	if WasteLibrary.AllowCors {

		w.Header().Set("Access-Control-Allow-Origin", "*")
	}
	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_OK
	WasteLibrary.LogStr("status 1")
	w.Write(resultVal.ToByte())
}
