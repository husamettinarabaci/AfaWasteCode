package main

import (
	"net/http"
	"net/url"

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
	http.HandleFunc("/getCustomers", getCustomers)
	http.HandleFunc("/setDevice", setDevice)
	http.HandleFunc("/getDevice", getDevice)
	http.HandleFunc("/getDevices", getDevices)
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
	var currentHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.StringToHttpClientHeaderType(req.FormValue("HEADER"))
	var currentData WasteLibrary.CustomerType = WasteLibrary.StringToCustomerType(req.FormValue("DATA"))
	data := url.Values{
		"HEADER": {currentHttpHeader.ToString()},
		"DATA":   {currentData.ToString()},
	}
	resultVal = WasteLibrary.SaveStaticDbMainForStoreApi(data)
	if resultVal.Result == "OK" {

		var isCustomerExist = false
		if currentData.CustomerId != 0 {
			isCustomerExist = true
		}
		currentData.CustomerId = WasteLibrary.StringIdToFloat64(resultVal.Retval.(string))
		data := url.Values{
			"HEADER": {currentHttpHeader.ToString()},
			"DATA":   {currentData.ToString()},
		}
		var currentCustomer WasteLibrary.CustomerType = WasteLibrary.StringToCustomerType(WasteLibrary.GetStaticDbMainForStoreApi(data).Retval.(string))

		resultVal = WasteLibrary.SaveRedisForStoreApi("customers", currentCustomer.ToIdString(), currentCustomer.ToString())
		resultVal = WasteLibrary.SaveRedisForStoreApi("customer-weblink", currentCustomer.WebLink, currentCustomer.ToIdString())
		resultVal = WasteLibrary.SaveRedisForStoreApi("customer-reportlink", currentCustomer.ReportLink, currentCustomer.ToIdString())
		resultVal = WasteLibrary.SaveRedisForStoreApi("customer-adminlink", currentCustomer.AdminLink, currentCustomer.ToIdString())
		if !isCustomerExist {
			var currentCustomerDevices WasteLibrary.CustomerDevicesType = WasteLibrary.CustomerDevicesType{
				CustomerId: currentCustomer.CustomerId,
				Devices:    make(map[float64]float64),
			}
			currentCustomerDevices.Devices[0] = 0

			var currentCustomerTags WasteLibrary.CustomerTagsType
			currentCustomerTags.CustomerId = currentCustomer.CustomerId
			currentCustomerTags.Tags = make(map[float64]WasteLibrary.TagType)
			currentCustomerTags.Tags[0] = WasteLibrary.TagType{TagID: 0}

			var currentCustomerConfig WasteLibrary.CustomerConfigType = WasteLibrary.CustomerConfigType{
				CustomerId: currentCustomer.CustomerId,
			}
			var currentAdminConfig WasteLibrary.AdminConfigType = WasteLibrary.AdminConfigType{
				CustomerId: currentCustomer.CustomerId,
			}
			var currentLocalConfig WasteLibrary.LocalConfigType = WasteLibrary.LocalConfigType{
				CustomerId: currentCustomer.CustomerId,
			}

			WasteLibrary.LogStr("CustomerDevices : " + currentCustomerDevices.ToString())
			WasteLibrary.LogStr("CustomerTags : " + currentCustomerTags.ToString())
			if resultVal.Result == "OK" {
				resultVal = WasteLibrary.GetRedisForStoreApi("customers", "customers")
				var currentCustomers WasteLibrary.CustomersType
				if resultVal.Result == "OK" {
					currentCustomers = WasteLibrary.StringToCustomersType(resultVal.Retval.(string))

				} else {
					currentCustomers = WasteLibrary.CustomersType{
						Customers: make(map[float64]float64),
					}
				}
				currentCustomers.Customers[currentCustomer.CustomerId] = currentCustomer.CustomerId
				resultVal = WasteLibrary.SaveRedisForStoreApi("customers", "customers", currentCustomers.ToString())
			}
			resultVal = WasteLibrary.SaveRedisForStoreApi("customer-devices", currentCustomerDevices.ToIdString(), currentCustomerDevices.ToString())
			resultVal = WasteLibrary.SaveRedisForStoreApi("customer-tags", currentCustomerTags.ToIdString(), currentCustomerTags.ToString())
			resultVal = WasteLibrary.SaveRedisForStoreApi("customer-customerconfig", currentCustomerConfig.ToIdString(), currentCustomerConfig.ToString())
			resultVal = WasteLibrary.SaveRedisForStoreApi("customer-adminconfig", currentAdminConfig.ToIdString(), currentAdminConfig.ToString())
			resultVal = WasteLibrary.SaveRedisForStoreApi("customer-localconfig", currentLocalConfig.ToIdString(), currentLocalConfig.ToString())
		}

	}
	w.Write(resultVal.ToByte())
}

func setConfig(w http.ResponseWriter, req *http.Request) {

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

	var currentHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.StringToHttpClientHeaderType(req.FormValue("HEADER"))
	if currentHttpHeader.OpType == "CUSTOMER" {
		var currentData WasteLibrary.CustomerConfigType = WasteLibrary.StringToCustomerConfigType(req.FormValue("DATA"))
		resultVal = WasteLibrary.SaveRedisForStoreApi("customer-customerconfig", currentData.ToIdString(), currentData.ToString())
	} else if currentHttpHeader.OpType == "ADMIN" {
		var currentData WasteLibrary.AdminConfigType = WasteLibrary.StringToAdminConfigType(req.FormValue("DATA"))
		resultVal = WasteLibrary.SaveRedisForStoreApi("customer-adminconfig", currentData.ToIdString(), currentData.ToString())
	} else if currentHttpHeader.OpType == "LOCAL" {
		var currentData WasteLibrary.LocalConfigType = WasteLibrary.StringToLocalConfigType(req.FormValue("DATA"))
		resultVal = WasteLibrary.SaveRedisForStoreApi("customer-localconfig", currentData.ToIdString(), currentData.ToString())
	} else {
		resultVal.Result = "FAIL"
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

	var currentHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.StringToHttpClientHeaderType(req.FormValue("HEADER"))
	var currentData WasteLibrary.DeviceType = WasteLibrary.StringToDeviceType(req.FormValue("DATA"))
	WasteLibrary.LogStr("Data : " + currentData.ToString())
	data := url.Values{
		"HEADER": {currentHttpHeader.ToString()},
		"DATA":   {currentData.ToString()},
	}
	resultVal = WasteLibrary.SaveStaticDbMainForStoreApi(data)
	if resultVal.Result == "OK" {

		var currentData WasteLibrary.DeviceType
		currentData.DeviceId = WasteLibrary.StringIdToFloat64(resultVal.Retval.(string))
		data := url.Values{
			"HEADER": {currentHttpHeader.ToString()},
			"DATA":   {currentData.ToString()},
		}
		var currentDevice WasteLibrary.DeviceType = WasteLibrary.StringToDeviceType(WasteLibrary.GetStaticDbMainForStoreApi(data).Retval.(string))

		resultVal = WasteLibrary.SaveRedisForStoreApi("devices", currentDevice.ToIdString(), currentDevice.ToString())
		resultVal = WasteLibrary.SaveRedisForStoreApi("serial-device", currentDevice.SerialNumber, currentDevice.ToIdString())
		resultVal = WasteLibrary.GetRedisForStoreApi("customer-devices", currentDevice.ToCustomerIdString())
		if resultVal.Result == "OK" {
			WasteLibrary.LogStr("Customer-Devices : " + resultVal.Retval.(string))
			var currentCustomerDevices WasteLibrary.CustomerDevicesType = WasteLibrary.StringToCustomerDevicesType(resultVal.Retval.(string))
			currentCustomerDevices.Devices[currentDevice.DeviceId] = currentDevice.DeviceId
			WasteLibrary.LogStr("New Customer-Devices : " + currentCustomerDevices.ToString())
			resultVal = WasteLibrary.SaveRedisForStoreApi("customer-devices", currentCustomerDevices.ToIdString(), currentCustomerDevices.ToString())
		}
	}
	w.Write(resultVal.ToByte())
}

func getCustomer(w http.ResponseWriter, req *http.Request) {

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
	//var currentHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.StringToHttpClientHeaderType(req.FormValue("HEADER"))
	var currentData WasteLibrary.CustomerType = WasteLibrary.StringToCustomerType(req.FormValue("DATA"))
	resultVal = WasteLibrary.GetRedisForStoreApi("customers", currentData.ToIdString())

	w.Write(resultVal.ToByte())
}

func getCustomers(w http.ResponseWriter, req *http.Request) {

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
	//var currentHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.StringToHttpClientHeaderType(req.FormValue("HEADER"))
	//var currentData WasteLibrary.CustomersType = WasteLibrary.StringToCustomersType(req.FormValue("DATA"))
	resultVal = WasteLibrary.GetRedisForStoreApi("customers", "customers")

	w.Write(resultVal.ToByte())
}

func getDevice(w http.ResponseWriter, req *http.Request) {

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

	//var currentHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.StringToHttpClientHeaderType(req.FormValue("HEADER"))
	var currentData WasteLibrary.DeviceType = WasteLibrary.StringToDeviceType(req.FormValue("DATA"))
	resultVal = WasteLibrary.GetRedisForStoreApi("devices", currentData.ToIdString())

	w.Write(resultVal.ToByte())
}

func getDevices(w http.ResponseWriter, req *http.Request) {

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

	//var currentHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.StringToHttpClientHeaderType(req.FormValue("HEADER"))
	var currentData WasteLibrary.CustomerType = WasteLibrary.StringToCustomerType(req.FormValue("DATA"))
	resultVal = WasteLibrary.GetRedisForStoreApi("customer-devices", currentData.ToIdString())

	w.Write(resultVal.ToByte())
}

func getConfig(w http.ResponseWriter, req *http.Request) {

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

	var currentHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.StringToHttpClientHeaderType(req.FormValue("HEADER"))
	if currentHttpHeader.OpType == "CUSTOMER" {
		var currentData WasteLibrary.CustomerConfigType = WasteLibrary.StringToCustomerConfigType(req.FormValue("DATA"))
		resultVal = WasteLibrary.GetRedisForStoreApi("customer-customerconfig", currentData.ToIdString())
	} else if currentHttpHeader.OpType == "ADMIN" {
		var currentData WasteLibrary.AdminConfigType = WasteLibrary.StringToAdminConfigType(req.FormValue("DATA"))
		resultVal = WasteLibrary.GetRedisForStoreApi("customer-adminconfig", currentData.ToIdString())
	} else if currentHttpHeader.OpType == "LOCAL" {
		var currentData WasteLibrary.LocalConfigType = WasteLibrary.StringToLocalConfigType(req.FormValue("DATA"))
		resultVal = WasteLibrary.GetRedisForStoreApi("customer-localconfig", currentData.ToIdString())
	} else {
		resultVal.Result = "FAIL"
	}

	w.Write(resultVal.ToByte())
}

func getTags(w http.ResponseWriter, req *http.Request) {

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

	//var currentHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.StringToHttpClientHeaderType(req.FormValue("HEADER"))
	var currentData WasteLibrary.CustomerType = WasteLibrary.StringToCustomerType(req.FormValue("DATA"))
	resultVal = WasteLibrary.GetRedisForStoreApi("customer-tags", currentData.ToIdString())

	w.Write(resultVal.ToByte())
}

func getTag(w http.ResponseWriter, req *http.Request) {

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

	//var currentHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.StringToHttpClientHeaderType(req.FormValue("HEADER"))
	var currentData WasteLibrary.TagType = WasteLibrary.StringToTagType(req.FormValue("DATA"))
	resultVal = WasteLibrary.GetRedisForStoreApi("tags", currentData.ToIdString())

	w.Write(resultVal.ToByte())
}

func checkAuth(data url.Values) devafatekresult.ResultType {
	var resultVal devafatekresult.ResultType
	resultVal.Result = "FAIL"
	resultVal.Result = "OK"
	return resultVal
}
