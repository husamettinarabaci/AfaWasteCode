package main

import (
	"net/http"

	"github.com/AfatekDevelopers/result_lib_go/devafatekresult"
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
	http.HandleFunc("/setCustomer", setCustomer)
	http.HandleFunc("/getCustomer", getCustomer)
	http.HandleFunc("/setDevice", setDevice)
	http.HandleFunc("/getDevice", getDevice)
	http.ListenAndServe(":80", nil)
}

func getCustomer(w http.ResponseWriter, req *http.Request) {
	var resultVal devafatekresult.ResultType
	resultVal.Result = "FAIL"
	if err := req.ParseForm(); err != nil {
		WasteLibrary.LogErr(err)
		return
	}

	resultVal = WasteLibrary.HttpPostReq("http://waste-adminapi-cluster-ip/getCustomer", req.Form)
	w.Write(resultVal.ToByte())
}

func setCustomer(w http.ResponseWriter, req *http.Request) {
	var resultVal devafatekresult.ResultType
	resultVal.Result = "FAIL"
	if err := req.ParseForm(); err != nil {
		WasteLibrary.LogErr(err)
		return
	}

	resultVal = WasteLibrary.HttpPostReq("http://waste-adminapi-cluster-ip/setCustomer", req.Form)
	w.Write(resultVal.ToByte())
}

func setDevice(w http.ResponseWriter, req *http.Request) {
	var resultVal devafatekresult.ResultType
	resultVal.Result = "FAIL"
	if err := req.ParseForm(); err != nil {
		WasteLibrary.LogErr(err)
		return
	}

	resultVal = WasteLibrary.HttpPostReq("http://waste-adminapi-cluster-ip/setDevice", req.Form)
	w.Write(resultVal.ToByte())
}

func getDevice(w http.ResponseWriter, req *http.Request) {
	var resultVal devafatekresult.ResultType
	resultVal.Result = "FAIL"
	if err := req.ParseForm(); err != nil {
		WasteLibrary.LogErr(err)
		return
	}

	resultVal = WasteLibrary.HttpPostReq("http://waste-adminapi-cluster-ip/getDevice", req.Form)
	w.Write(resultVal.ToByte())
}
