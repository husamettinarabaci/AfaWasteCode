package main

import (
	"net/http"
	"net/url"
	"strconv"

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
	http.ListenAndServe(":80", nil)
}

func setCustomer(w http.ResponseWriter, req *http.Request) {

	var resultVal devafatekresult.ResultType
	resultVal.Result = "FAIL"
	if err := req.ParseForm(); err != nil {
		WasteLibrary.LogErr(err)
		return
	}
	resultVal = checkAuth(req.Form)
	if resultVal.Result != "OK" {
		w.Write(resultVal.ToByte())
		return
	}
	data := url.Values{
		"APPTYPE": {"ADMIN"},
		"OPTYPE":  {"CUSTOMER"},
	}
	customerId, _ := strconv.Atoi(req.FormValue("CUSTOMERID"))
	var currentCustomer WasteLibrary.CustomerType = WasteLibrary.CustomerType{
		CustomerId:   float64(customerId),
		CustomerName: req.FormValue("CUSTOMERNAME"),
		Domain:       req.FormValue("DOMAIN"),
		RfIdApp:      req.FormValue("RFIDAPP"),
		UltApp:       req.FormValue("ULTAPP"),
		RecyApp:      req.FormValue("RECYAPP"),
	}
	data.Add("DATA", currentCustomer.ToString())

	resultVal = WasteLibrary.SaveStaticDbMainForStoreApi(data)
	if resultVal.Result == "OK" {

		customerTypeVal := WasteLibrary.StringToCustomerType(resultVal.Retval.(string))
		WasteLibrary.LogStr("Customer : " + customerTypeVal.ToString())
		resultVal = WasteLibrary.SaveRedisForStoreApi("customers", customerTypeVal.ToIdString(), customerTypeVal.ToString())
	}
	w.Write(resultVal.ToByte())
}

func setDevice(w http.ResponseWriter, req *http.Request) {

	var resultVal devafatekresult.ResultType
	resultVal.Result = "FAIL"
	if err := req.ParseForm(); err != nil {
		WasteLibrary.LogErr(err)
		return
	}
	resultVal = checkAuth(req.Form)
	if resultVal.Result != "OK" {
		w.Write(resultVal.ToByte())
		return
	}
	data := url.Values{
		"APPTYPE": {"ADMIN"},
		"OPTYPE":  {"DEVICE"},
	}
	deviceId, _ := strconv.Atoi(req.FormValue("DEVICEID"))
	customerId, _ := strconv.Atoi(req.FormValue("CUSTOMERID"))
	var currentDevice WasteLibrary.DeviceType = WasteLibrary.DeviceType{
		DeviceId:     float64(deviceId),
		CustomerId:   float64(customerId),
		DeviceName:   req.FormValue("DEVICENAME"),
		DeviceType:   req.FormValue("DEVICETYPE"),
		SerialNumber: req.FormValue("SERIALNUMBER"),
	}
	data.Add("DATA", currentDevice.ToString())

	resultVal = WasteLibrary.SaveStaticDbMainForStoreApi(data)
	if resultVal.Result == "OK" {

		deviceTypeVal := WasteLibrary.StringToCustomerType(resultVal.Retval.(string))
		WasteLibrary.LogStr("Device : " + deviceTypeVal.ToString())
		resultVal = WasteLibrary.SaveRedisForStoreApi("devices", deviceTypeVal.ToIdString(), deviceTypeVal.ToString())
	}
	w.Write(resultVal.ToByte())
}

func checkAuth(data url.Values) devafatekresult.ResultType {
	var resultVal devafatekresult.ResultType
	resultVal.Result = "FAIL"
	resultVal.Result = "OK"
	return resultVal
}
