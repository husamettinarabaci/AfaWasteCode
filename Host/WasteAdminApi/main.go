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
	http.HandleFunc("/setConfig", setConfig)
	http.HandleFunc("/getConfig", getConfig)
	http.HandleFunc("/getCustomer", getCustomer)
	http.HandleFunc("/setUser", setUser)
	http.HandleFunc("/getUser", getUser)
	http.HandleFunc("/getUsers", getUsers)
	http.HandleFunc("/getLink", getLink)
	http.ListenAndServe(":80", nil)
}

func getLink(w http.ResponseWriter, req *http.Request) {

	var resultVal devafatekresult.ResultType
	resultVal.Result = "FAIL"

	var linkVal string = req.Host
	WasteLibrary.LogStr("Get Link : " + linkVal)
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

	resultVal = WasteLibrary.GetRedisForStoreApi("customer-link", req.Host)
	if resultVal.Result != "OK" {
		w.Write(resultVal.ToByte())
		return
	}
	var customerId string = resultVal.Retval.(string)

	var currentHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.StringToHttpClientHeaderType(req.FormValue("HEADER"))
	if currentHttpHeader.OpType == "CUSTOMER" {
		var currentData WasteLibrary.CustomerConfigType = WasteLibrary.StringToCustomerConfigType(req.FormValue("DATA"))
		resultVal = WasteLibrary.SaveRedisForStoreApi("customer-customerconfig", customerId, currentData.ToString())
	} else if currentHttpHeader.OpType == "ADMIN" {
		var currentData WasteLibrary.AdminConfigType = WasteLibrary.StringToAdminConfigType(req.FormValue("DATA"))
		resultVal = WasteLibrary.SaveRedisForStoreApi("customer-adminconfig", customerId, currentData.ToString())
	} else if currentHttpHeader.OpType == "LOCAL" {
		var currentData WasteLibrary.LocalConfigType = WasteLibrary.StringToLocalConfigType(req.FormValue("DATA"))
		resultVal = WasteLibrary.SaveRedisForStoreApi("customer-localconfig", customerId, currentData.ToString())
	} else {
		resultVal.Result = "FAIL"
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

	resultVal = WasteLibrary.GetRedisForStoreApi("customer-link", req.Host)
	if resultVal.Result != "OK" {
		w.Write(resultVal.ToByte())
		return
	}
	var customerId string = resultVal.Retval.(string)

	resultVal = WasteLibrary.GetRedisForStoreApi("customers", customerId)

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

	resultVal = WasteLibrary.GetRedisForStoreApi("customer-link", req.Host)
	if resultVal.Result != "OK" {
		w.Write(resultVal.ToByte())
		return
	}
	var customerId string = resultVal.Retval.(string)

	var currentData WasteLibrary.DeviceType = WasteLibrary.StringToDeviceType(req.FormValue("DATA"))
	resultVal = WasteLibrary.GetRedisForStoreApi("devices", currentData.ToIdString())

	currentData = WasteLibrary.StringToDeviceType(resultVal.Retval.(string))
	if currentData.ToCustomerIdString() == customerId {
		w.Write(resultVal.ToByte())
	} else {
		resultVal.Result = "FAIL"
		resultVal.Retval = ""
		w.Write(resultVal.ToByte())
	}
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

	resultVal = WasteLibrary.GetRedisForStoreApi("customer-link", req.Host)
	if resultVal.Result != "OK" {
		w.Write(resultVal.ToByte())
		return
	}
	var customerId string = resultVal.Retval.(string)

	resultVal = WasteLibrary.GetRedisForStoreApi("customer-devices", customerId)

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

	resultVal = WasteLibrary.GetRedisForStoreApi("customer-link", req.Host)
	if resultVal.Result != "OK" {
		w.Write(resultVal.ToByte())
		return
	}
	var customerId string = resultVal.Retval.(string)

	var currentHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.StringToHttpClientHeaderType(req.FormValue("HEADER"))
	if currentHttpHeader.OpType == "CUSTOMER" {
		resultVal = WasteLibrary.GetRedisForStoreApi("customer-customerconfig", customerId)
	} else if currentHttpHeader.OpType == "ADMIN" {
		resultVal = WasteLibrary.GetRedisForStoreApi("customer-adminconfig", customerId)
	} else if currentHttpHeader.OpType == "LOCAL" {
		resultVal = WasteLibrary.GetRedisForStoreApi("customer-localconfig", customerId)
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

	resultVal = WasteLibrary.GetRedisForStoreApi("customer-link", req.Host)
	if resultVal.Result != "OK" {
		w.Write(resultVal.ToByte())
		return
	}
	var customerId string = resultVal.Retval.(string)

	resultVal = WasteLibrary.GetRedisForStoreApi("customer-tags", customerId)

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

	resultVal = WasteLibrary.GetRedisForStoreApi("customer-link", req.Host)
	if resultVal.Result != "OK" {
		w.Write(resultVal.ToByte())
		return
	}
	var customerId string = resultVal.Retval.(string)

	var currentData WasteLibrary.TagType = WasteLibrary.StringToTagType(req.FormValue("DATA"))

	resultVal = WasteLibrary.GetRedisForStoreApi("customer-tags", customerId)
	if resultVal.Result == "OK" {
		var currentCustomerTags WasteLibrary.CustomerTagsType = WasteLibrary.StringToCustomerTagsType(resultVal.Retval.(string))
		currentData = currentCustomerTags.Tags[currentData.ToIdString()]
		resultVal.Retval = currentData.ToString()
	}
	w.Write(resultVal.ToByte())
}

func getUser(w http.ResponseWriter, req *http.Request) {

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

	resultVal = WasteLibrary.GetRedisForStoreApi("customer-link", req.Host)
	if resultVal.Result != "OK" {
		w.Write(resultVal.ToByte())
		return
	}
	var customerId string = resultVal.Retval.(string)

	var currentData WasteLibrary.UserType = WasteLibrary.StringToUserType(req.FormValue("DATA"))
	resultVal = WasteLibrary.GetRedisForStoreApi("users", currentData.ToIdString())

	currentData = WasteLibrary.StringToUserType(resultVal.Retval.(string))
	if currentData.ToCustomerIdString() == customerId {
		w.Write(resultVal.ToByte())
	} else {
		resultVal.Result = "FAIL"
		resultVal.Retval = ""
		w.Write(resultVal.ToByte())
	}
}

func setUser(w http.ResponseWriter, req *http.Request) {

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

	resultVal = WasteLibrary.GetRedisForStoreApi("customer-link", req.Host)
	if resultVal.Result != "OK" {
		w.Write(resultVal.ToByte())
		return
	}
	var customerId string = resultVal.Retval.(string)

	var currentHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.HttpClientHeaderType{
		AppType:      "ADMIN",
		DeviceNo:     "",
		OpType:       "USER",
		Time:         WasteLibrary.GetTime(),
		Repeat:       "0",
		DeviceId:     0,
		CustomerId:   WasteLibrary.StringIdToFloat64(customerId),
		BaseDataType: "CUSTOMER",
	}
	var currentData WasteLibrary.UserType = WasteLibrary.StringToUserType(req.FormValue("DATA"))
	currentData.CustomerId = WasteLibrary.StringIdToFloat64(customerId)
	WasteLibrary.LogStr("Data : " + currentData.ToString())
	data := url.Values{
		"HEADER": {currentHttpHeader.ToString()},
		"DATA":   {currentData.ToString()},
	}
	resultVal = WasteLibrary.SaveStaticDbMainForStoreApi(data)
	if resultVal.Result == "OK" {

		var currentData WasteLibrary.UserType
		currentData.UserId = WasteLibrary.StringIdToFloat64(resultVal.Retval.(string))
		data := url.Values{
			"HEADER": {currentHttpHeader.ToString()},
			"DATA":   {currentData.ToString()},
		}
		var currentUser WasteLibrary.UserType = WasteLibrary.StringToUserType(WasteLibrary.GetStaticDbMainForStoreApi(data).Retval.(string))

		resultVal = WasteLibrary.SaveRedisForStoreApi("users", currentUser.ToIdString(), currentUser.ToString())
		resultVal = WasteLibrary.GetRedisForStoreApi("customer-users", currentUser.ToCustomerIdString())
		if resultVal.Result == "OK" {
			WasteLibrary.LogStr("Customer-Users : " + resultVal.Retval.(string))
			var currentCustomerUsers WasteLibrary.CustomerUsersType = WasteLibrary.StringToCustomerUsersType(resultVal.Retval.(string))
			currentCustomerUsers.Users[currentUser.ToIdString()] = currentUser.UserId
			WasteLibrary.LogStr("New Customer-Users : " + currentCustomerUsers.ToString())
			resultVal = WasteLibrary.SaveRedisForStoreApi("customer-users", currentCustomerUsers.ToIdString(), currentCustomerUsers.ToString())
		}
	}
	w.Write(resultVal.ToByte())
}

func getUsers(w http.ResponseWriter, req *http.Request) {

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

	resultVal = WasteLibrary.GetRedisForStoreApi("customer-link", req.Host)
	if resultVal.Result != "OK" {
		w.Write(resultVal.ToByte())
		return
	}
	var customerId string = resultVal.Retval.(string)

	resultVal = WasteLibrary.GetRedisForStoreApi("customer-users", customerId)

	w.Write(resultVal.ToByte())
}

func checkAuth(data url.Values) devafatekresult.ResultType {
	var resultVal devafatekresult.ResultType
	resultVal.Result = "FAIL"
	resultVal.Result = "OK"
	return resultVal
}
