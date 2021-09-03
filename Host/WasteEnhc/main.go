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
	http.HandleFunc("/data", data)
	http.ListenAndServe(":80", nil)
}

func data(w http.ResponseWriter, req *http.Request) {

	var resultVal devafatekresult.ResultType
	resultVal.Result = "FAIL"
	if err := req.ParseForm(); err != nil {
		WasteLibrary.LogErr(err)
		return
	}
	appTypeVal := req.FormValue("APPTYPE")
	didVal := req.FormValue("DID")
	dataTypeVal := req.FormValue("DATATYPE")
	customerIdVal := WasteLibrary.GetRedisForStoreApi("devices-customer", didVal).Retval.(string)
	req.Form.Add("CUSTOMERID", customerIdVal)

	resultVal = WasteLibrary.HttpPostReq("http://waste-storeapi-cluster-ip/saveBulkDbMain", req.Form)

	var serviceClusterIp string = ""
	if appTypeVal == "RFID" {

		if dataTypeVal == "RF" {

			serviceClusterIp = "waste-rfreader-cluster-ip"

		} else if dataTypeVal == "GPS" {

			serviceClusterIp = "waste-gpsreader-cluster-ip"

		} else if dataTypeVal == "STATUS" {

			serviceClusterIp = "waste-statusreader-cluster-ip"

		} else if dataTypeVal == "THERM" {

			serviceClusterIp = "waste-thermreader-cluster-ip"

		} else if dataTypeVal == "CAM" {

			serviceClusterIp = "waste-camreader-cluster-ip"

		} else {
			resultVal.Result = "FAIL"
		}
	} else if appTypeVal == "ULT" {
		resultVal.Result = "OK"
	} else if appTypeVal == "RECY" {
		resultVal.Result = "OK"
	} else {
		resultVal.Result = "FAIL"
	}
	if serviceClusterIp != "" {
		resultVal = WasteLibrary.HttpPostReq("http://"+serviceClusterIp+"/reader", req.Form)
	}
	w.Write(resultVal.ToByte())
}
