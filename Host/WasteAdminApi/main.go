package main

import (
	"net/http"
	"net/url"

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

	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL

	var linkVal string = req.Host
	WasteLibrary.LogStr("Get Link : " + linkVal)
	w.Write(resultVal.ToByte())
}

func setConfig(w http.ResponseWriter, req *http.Request) {

	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	if err := req.ParseForm(); err != nil {
		WasteLibrary.LogErr(err)
		return
	}
	resultVal = checkAuth(req.Form)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		w.Write(resultVal.ToByte())
		return
	}

	resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_LINK, req.Host)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		w.Write(resultVal.ToByte())
		return
	}
	var customerId string = resultVal.Retval.(string)

	var currentHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.StringToHttpClientHeaderType(req.FormValue(WasteLibrary.HTTP_HEADER))
	if currentHttpHeader.OpType == WasteLibrary.OPTYPE_CUSTOMERCONFIG {
		var currentData WasteLibrary.CustomerConfigType = WasteLibrary.StringToCustomerConfigType(req.FormValue(WasteLibrary.HTTP_DATA))
		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_CUSTOMERCONFIG, customerId, currentData.ToString())
	} else if currentHttpHeader.OpType == WasteLibrary.OPTYPE_ADMINCONFIG {
		var currentData WasteLibrary.AdminConfigType = WasteLibrary.StringToAdminConfigType(req.FormValue(WasteLibrary.HTTP_DATA))
		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_ADMINCONFIG, customerId, currentData.ToString())
	} else if currentHttpHeader.OpType == WasteLibrary.OPTYPE_LOCALCONFIG {
		var currentData WasteLibrary.LocalConfigType = WasteLibrary.StringToLocalConfigType(req.FormValue(WasteLibrary.HTTP_DATA))
		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_LOCALCONFIG, customerId, currentData.ToString())
	} else {
		resultVal.Result = WasteLibrary.RESULT_FAIL
	}

	w.Write(resultVal.ToByte())
}

func getCustomer(w http.ResponseWriter, req *http.Request) {

	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	if err := req.ParseForm(); err != nil {
		WasteLibrary.LogErr(err)
		return
	}
	resultVal = checkAuth(req.Form)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		w.Write(resultVal.ToByte())
		return
	}

	resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_LINK, req.Host)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		w.Write(resultVal.ToByte())
		return
	}
	var customerId string = resultVal.Retval.(string)

	resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMERS, customerId)

	w.Write(resultVal.ToByte())
}

func getDevice(w http.ResponseWriter, req *http.Request) {

	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	if err := req.ParseForm(); err != nil {
		WasteLibrary.LogErr(err)
		return
	}
	resultVal = checkAuth(req.Form)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		w.Write(resultVal.ToByte())
		return
	}

	resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_LINK, req.Host)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		w.Write(resultVal.ToByte())
		return
	}
	var customerId string = resultVal.Retval.(string)

	var currentData WasteLibrary.DeviceType = WasteLibrary.StringToDeviceType(req.FormValue(WasteLibrary.HTTP_DATA))
	resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_DEVICES, currentData.ToIdString())

	currentData = WasteLibrary.StringToDeviceType(resultVal.Retval.(string))
	if currentData.ToCustomerIdString() == customerId {
		w.Write(resultVal.ToByte())
	} else {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = ""
		w.Write(resultVal.ToByte())
	}
}

func getDevices(w http.ResponseWriter, req *http.Request) {

	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	if err := req.ParseForm(); err != nil {
		WasteLibrary.LogErr(err)
		return
	}
	resultVal = checkAuth(req.Form)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		w.Write(resultVal.ToByte())
		return
	}

	resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_LINK, req.Host)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		w.Write(resultVal.ToByte())
		return
	}
	var customerId string = resultVal.Retval.(string)

	resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_DEVICES, customerId)

	w.Write(resultVal.ToByte())
}

func getConfig(w http.ResponseWriter, req *http.Request) {

	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	if err := req.ParseForm(); err != nil {
		WasteLibrary.LogErr(err)
		return
	}
	resultVal = checkAuth(req.Form)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		w.Write(resultVal.ToByte())
		return
	}

	resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_LINK, req.Host)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		w.Write(resultVal.ToByte())
		return
	}
	var customerId string = resultVal.Retval.(string)

	var currentHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.StringToHttpClientHeaderType(req.FormValue(WasteLibrary.HTTP_HEADER))
	if currentHttpHeader.OpType == WasteLibrary.OPTYPE_CUSTOMERCONFIG {
		resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_CUSTOMERCONFIG, customerId)
	} else if currentHttpHeader.OpType == WasteLibrary.OPTYPE_ADMINCONFIG {
		resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_ADMINCONFIG, customerId)
	} else if currentHttpHeader.OpType == WasteLibrary.OPTYPE_LOCALCONFIG {
		resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_LOCALCONFIG, customerId)
	} else {
		resultVal.Result = WasteLibrary.RESULT_FAIL
	}

	w.Write(resultVal.ToByte())
}

func getTags(w http.ResponseWriter, req *http.Request) {

	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	if err := req.ParseForm(); err != nil {
		WasteLibrary.LogErr(err)
		return
	}
	resultVal = checkAuth(req.Form)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		w.Write(resultVal.ToByte())
		return
	}

	resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_LINK, req.Host)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		w.Write(resultVal.ToByte())
		return
	}
	var customerId string = resultVal.Retval.(string)

	resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_TAGS, customerId)

	w.Write(resultVal.ToByte())
}

func getTag(w http.ResponseWriter, req *http.Request) {

	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	if err := req.ParseForm(); err != nil {
		WasteLibrary.LogErr(err)
		return
	}
	resultVal = checkAuth(req.Form)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		w.Write(resultVal.ToByte())
		return
	}

	resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_LINK, req.Host)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		w.Write(resultVal.ToByte())
		return
	}
	var customerId string = resultVal.Retval.(string)

	var currentData WasteLibrary.TagType = WasteLibrary.StringToTagType(req.FormValue(WasteLibrary.HTTP_DATA))

	resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_TAGS, customerId)
	if resultVal.Result == WasteLibrary.RESULT_OK {
		var currentCustomerTags WasteLibrary.CustomerTagsType = WasteLibrary.StringToCustomerTagsType(resultVal.Retval.(string))
		currentData = currentCustomerTags.Tags[currentData.ToIdString()]
		resultVal.Retval = currentData.ToString()
	}
	w.Write(resultVal.ToByte())
}

func getUser(w http.ResponseWriter, req *http.Request) {

	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	if err := req.ParseForm(); err != nil {
		WasteLibrary.LogErr(err)
		return
	}
	resultVal = checkAuth(req.Form)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		w.Write(resultVal.ToByte())
		return
	}

	resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_LINK, req.Host)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		w.Write(resultVal.ToByte())
		return
	}
	var customerId string = resultVal.Retval.(string)

	var currentData WasteLibrary.UserType = WasteLibrary.StringToUserType(req.FormValue(WasteLibrary.HTTP_DATA))
	resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_USERS, currentData.ToIdString())

	currentData = WasteLibrary.StringToUserType(resultVal.Retval.(string))
	if currentData.ToCustomerIdString() == customerId {
		w.Write(resultVal.ToByte())
	} else {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = ""
		w.Write(resultVal.ToByte())
	}
}

func setUser(w http.ResponseWriter, req *http.Request) {

	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	if err := req.ParseForm(); err != nil {
		WasteLibrary.LogErr(err)
		return
	}
	resultVal = checkAuth(req.Form)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		w.Write(resultVal.ToByte())
		return
	}

	resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_LINK, req.Host)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		w.Write(resultVal.ToByte())
		return
	}
	var customerId string = resultVal.Retval.(string)

	var currentHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.HttpClientHeaderType{
		AppType:      WasteLibrary.APPTYPE_ADMIN,
		DeviceNo:     "",
		OpType:       WasteLibrary.OPTYPE_USER,
		Time:         WasteLibrary.GetTime(),
		Repeat:       WasteLibrary.STATU_PASSIVE,
		DeviceId:     0,
		CustomerId:   WasteLibrary.StringIdToFloat64(customerId),
		BaseDataType: WasteLibrary.BASETYPE_CUSTOMER,
	}
	var currentData WasteLibrary.UserType = WasteLibrary.StringToUserType(req.FormValue(WasteLibrary.HTTP_DATA))
	currentData.CustomerId = WasteLibrary.StringIdToFloat64(customerId)
	WasteLibrary.LogStr("Data : " + currentData.ToString())
	data := url.Values{
		WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
		WasteLibrary.HTTP_DATA:   {currentData.ToString()},
	}
	resultVal = WasteLibrary.SaveStaticDbMainForStoreApi(data)
	if resultVal.Result == WasteLibrary.RESULT_OK {

		var currentData WasteLibrary.UserType
		currentData.UserId = WasteLibrary.StringIdToFloat64(resultVal.Retval.(string))
		data := url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentData.ToString()},
		}
		var currentUser WasteLibrary.UserType = WasteLibrary.StringToUserType(WasteLibrary.GetStaticDbMainForStoreApi(data).Retval.(string))

		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_USERS, currentUser.ToIdString(), currentUser.ToString())
		resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_USERS, currentUser.ToCustomerIdString())
		if resultVal.Result == WasteLibrary.RESULT_OK {
			WasteLibrary.LogStr("Customer-Users : " + resultVal.Retval.(string))
			var currentCustomerUsers WasteLibrary.CustomerUsersType = WasteLibrary.StringToCustomerUsersType(resultVal.Retval.(string))
			currentCustomerUsers.Users[currentUser.ToIdString()] = currentUser.UserId
			WasteLibrary.LogStr("New Customer-Users : " + currentCustomerUsers.ToString())
			resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_USERS, currentCustomerUsers.ToIdString(), currentCustomerUsers.ToString())
		}
	}
	w.Write(resultVal.ToByte())
}

func getUsers(w http.ResponseWriter, req *http.Request) {

	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	if err := req.ParseForm(); err != nil {
		WasteLibrary.LogErr(err)
		return
	}
	resultVal = checkAuth(req.Form)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		w.Write(resultVal.ToByte())
		return
	}

	resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_LINK, req.Host)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		w.Write(resultVal.ToByte())
		return
	}
	var customerId string = resultVal.Retval.(string)

	resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_USERS, customerId)

	w.Write(resultVal.ToByte())
}

func checkAuth(data url.Values) WasteLibrary.ResultType {
	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	resultVal.Result = WasteLibrary.RESULT_OK
	return resultVal
}
