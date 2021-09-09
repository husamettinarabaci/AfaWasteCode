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
	http.HandleFunc("/setDevice", setDevice)
	http.HandleFunc("/setCustomerConfig", setCustomerConfig)
	http.HandleFunc("/setAdminConfig", setAdminConfig)
	http.HandleFunc("/setLocalConfig", setLocalConfig)
	http.HandleFunc("/getCustomer", getCustomer)
	http.HandleFunc("/getCustomers", getCustomers)
	http.HandleFunc("/getDevice", getDevice)
	http.HandleFunc("/getDevices", getDevices)
	http.HandleFunc("/getCustomerConfig", getCustomerConfig)
	http.HandleFunc("/getAdminConfig", getAdminConfig)
	http.HandleFunc("/getLocalConfig", getLocalConfig)
	http.HandleFunc("/getTags", getTags)
	http.HandleFunc("/getTag", getTag)
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

func getCustomers(w http.ResponseWriter, req *http.Request) {
	var resultVal devafatekresult.ResultType
	resultVal.Result = "FAIL"
	if err := req.ParseForm(); err != nil {
		WasteLibrary.LogErr(err)
		return
	}

	resultVal = WasteLibrary.HttpPostReq("http://waste-adminapi-cluster-ip/getCustomers", req.Form)
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

func setConfig(w http.ResponseWriter, req *http.Request) {
	var resultVal devafatekresult.ResultType
	resultVal.Result = "FAIL"
	if err := req.ParseForm(); err != nil {
		WasteLibrary.LogErr(err)
		return
	}

	resultVal = WasteLibrary.HttpPostReq("http://waste-adminapi-cluster-ip/setConfig", req.Form)
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

func getDevices(w http.ResponseWriter, req *http.Request) {
	var resultVal devafatekresult.ResultType
	resultVal.Result = "FAIL"
	if err := req.ParseForm(); err != nil {
		WasteLibrary.LogErr(err)
		return
	}

	resultVal = WasteLibrary.HttpPostReq("http://waste-adminapi-cluster-ip/getDevices", req.Form)
	w.Write(resultVal.ToByte())
}

func getTag(w http.ResponseWriter, req *http.Request) {
	var resultVal devafatekresult.ResultType
	resultVal.Result = "FAIL"
	if err := req.ParseForm(); err != nil {
		WasteLibrary.LogErr(err)
		return
	}

	resultVal = WasteLibrary.HttpPostReq("http://waste-adminapi-cluster-ip/getTag", req.Form)
	w.Write(resultVal.ToByte())
}

func getTags(w http.ResponseWriter, req *http.Request) {
	var resultVal devafatekresult.ResultType
	resultVal.Result = "FAIL"
	if err := req.ParseForm(); err != nil {
		WasteLibrary.LogErr(err)
		return
	}

	resultVal = WasteLibrary.HttpPostReq("http://waste-adminapi-cluster-ip/getTags", req.Form)
	w.Write(resultVal.ToByte())
}

func getCustomerConfig(w http.ResponseWriter, req *http.Request) {
	var resultVal devafatekresult.ResultType
	resultVal.Result = "FAIL"
	if err := req.ParseForm(); err != nil {
		WasteLibrary.LogErr(err)
		return
	}

	resultVal = WasteLibrary.HttpPostReq("http://waste-adminapi-cluster-ip/getCustomerConfig", req.Form)
	w.Write(resultVal.ToByte())
}

func getAdminConfig(w http.ResponseWriter, req *http.Request) {
	var resultVal devafatekresult.ResultType
	resultVal.Result = "FAIL"
	if err := req.ParseForm(); err != nil {
		WasteLibrary.LogErr(err)
		return
	}

	resultVal = WasteLibrary.HttpPostReq("http://waste-adminapi-cluster-ip/getAdminConfig", req.Form)
	w.Write(resultVal.ToByte())
}

func getLocalConfig(w http.ResponseWriter, req *http.Request) {
	var resultVal devafatekresult.ResultType
	resultVal.Result = "FAIL"
	if err := req.ParseForm(); err != nil {
		WasteLibrary.LogErr(err)
		return
	}

	resultVal = WasteLibrary.HttpPostReq("http://waste-adminapi-cluster-ip/getLocalConfig", req.Form)
	w.Write(resultVal.ToByte())
}

func setCustomerConfig(w http.ResponseWriter, req *http.Request) {
	var resultVal devafatekresult.ResultType
	resultVal.Result = "FAIL"
	if err := req.ParseForm(); err != nil {
		WasteLibrary.LogErr(err)
		return
	}

	resultVal = WasteLibrary.HttpPostReq("http://waste-adminapi-cluster-ip/setCustomerConfig", req.Form)
	w.Write(resultVal.ToByte())
}

func setAdminConfig(w http.ResponseWriter, req *http.Request) {
	var resultVal devafatekresult.ResultType
	resultVal.Result = "FAIL"
	if err := req.ParseForm(); err != nil {
		WasteLibrary.LogErr(err)
		return
	}

	resultVal = WasteLibrary.HttpPostReq("http://waste-adminapi-cluster-ip/setAdminConfig", req.Form)
	w.Write(resultVal.ToByte())
}

func setLocalConfig(w http.ResponseWriter, req *http.Request) {
	var resultVal devafatekresult.ResultType
	resultVal.Result = "FAIL"
	if err := req.ParseForm(); err != nil {
		WasteLibrary.LogErr(err)
		return
	}

	resultVal = WasteLibrary.HttpPostReq("http://waste-adminapi-cluster-ip/setLocalConfig", req.Form)
	w.Write(resultVal.ToByte())
}
