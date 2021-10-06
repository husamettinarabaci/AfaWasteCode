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
	http.ListenAndServe(":80", nil)
}

func getUser(w http.ResponseWriter, req *http.Request) {
	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	if err := req.ParseForm(); err != nil {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_HTTP_PARSE
		w.Write(resultVal.ToByte())
		WasteLibrary.LogErr(err)
		return
	}

	resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_LINK, req.Host)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_NOTFOUND
		w.Write(resultVal.ToByte())
		return
	}
	var customerId string = resultVal.Retval.(string)

	resultVal = checkAuth(req.Form, customerId)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		w.Write(resultVal.ToByte())
		return
	}

	var currentData WasteLibrary.UserType = WasteLibrary.StringToUserType(req.FormValue(WasteLibrary.HTTP_DATA))
	resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_USERS, currentData.ToIdString())
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_USER_NOTFOUND
		w.Write(resultVal.ToByte())
		return
	}
	currentData = WasteLibrary.StringToUserType(resultVal.Retval.(string))
	if currentData.ToCustomerIdString() == customerId {
		w.Write(resultVal.ToByte())
	} else {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_USER_NOTFOUND
		w.Write(resultVal.ToByte())
	}
}

func setUser(w http.ResponseWriter, req *http.Request) {
	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	if err := req.ParseForm(); err != nil {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_HTTP_PARSE
		w.Write(resultVal.ToByte())
		WasteLibrary.LogErr(err)
		return
	}

	resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_LINK, req.Host)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_NOTFOUND
		w.Write(resultVal.ToByte())
		return
	}
	var customerId string = resultVal.Retval.(string)

	resultVal = checkAuth(req.Form, customerId)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		w.Write(resultVal.ToByte())
		return
	}

	var currentData WasteLibrary.UserType = WasteLibrary.StringToUserType(req.FormValue(WasteLibrary.HTTP_DATA))
	if currentData.ToCustomerIdString() != customerId {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_USER_NOTFOUND
		w.Write(resultVal.ToByte())
		return
	}

	var httpHeaderForUserInsert WasteLibrary.HttpClientHeaderType = WasteLibrary.HttpClientHeaderType{
		AppType:      WasteLibrary.APPTYPE_ADMIN,
		DeviceNo:     "",
		OpType:       WasteLibrary.OPTYPE_USER,
		Time:         WasteLibrary.GetTime(),
		Repeat:       WasteLibrary.STATU_PASSIVE,
		DeviceId:     0,
		CustomerId:   WasteLibrary.StringIdToFloat64(customerId),
		BaseDataType: WasteLibrary.BASETYPE_USER,
	}
	data := url.Values{
		WasteLibrary.HTTP_HEADER: {httpHeaderForUserInsert.ToString()},
		WasteLibrary.HTTP_DATA:   {currentData.ToString()},
	}
	resultVal = WasteLibrary.SaveConfigDbMainForStoreApi(data)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
		w.Write(resultVal.ToByte())
		return
	}
	currentData = WasteLibrary.StringToUserType(resultVal.Retval.(string))
	resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_USERS, currentData.ToIdString(), currentData.ToString())
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
		w.Write(resultVal.ToByte())
		return
	}
	w.Write(resultVal.ToByte())

}

func setConfig(w http.ResponseWriter, req *http.Request) {

	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	if err := req.ParseForm(); err != nil {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_HTTP_PARSE
		w.Write(resultVal.ToByte())
		WasteLibrary.LogErr(err)
		return
	}
	resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_LINK, req.Host)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_NOTFOUND
		w.Write(resultVal.ToByte())
		return
	}
	var customerId string = resultVal.Retval.(string)

	resultVal = checkAuth(req.Form, customerId)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		w.Write(resultVal.ToByte())
		return
	}

	var currentHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.StringToHttpClientHeaderType(req.FormValue(WasteLibrary.HTTP_HEADER))
	if currentHttpHeader.OpType == WasteLibrary.OPTYPE_CUSTOMERCONFIG {
		var currentData WasteLibrary.CustomerConfigType = WasteLibrary.StringToCustomerConfigType(req.FormValue(WasteLibrary.HTTP_DATA))
		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_CUSTOMERCONFIG, customerId, currentData.ToString())
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())
			return
		}
	} else if currentHttpHeader.OpType == WasteLibrary.OPTYPE_ADMINCONFIG {
		var currentData WasteLibrary.AdminConfigType = WasteLibrary.StringToAdminConfigType(req.FormValue(WasteLibrary.HTTP_DATA))
		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_ADMINCONFIG, customerId, currentData.ToString())
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())
			return
		}
	} else if currentHttpHeader.OpType == WasteLibrary.OPTYPE_LOCALCONFIG {
		var currentData WasteLibrary.LocalConfigType = WasteLibrary.StringToLocalConfigType(req.FormValue(WasteLibrary.HTTP_DATA))
		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_LOCALCONFIG, customerId, currentData.ToString())
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())
			return
		}
	} else {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_OPTYPE
		w.Write(resultVal.ToByte())
		return
	}

	w.Write(resultVal.ToByte())
}

func getCustomer(w http.ResponseWriter, req *http.Request) {

	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	if err := req.ParseForm(); err != nil {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_HTTP_PARSE
		w.Write(resultVal.ToByte())
		WasteLibrary.LogErr(err)
		return
	}
	resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_LINK, req.Host)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_NOTFOUND
		w.Write(resultVal.ToByte())
		return
	}
	var customerId string = resultVal.Retval.(string)

	resultVal = checkAuth(req.Form, customerId)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		w.Write(resultVal.ToByte())
		return
	}

	resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMERS, customerId)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_GET
		w.Write(resultVal.ToByte())
		return
	}
	w.Write(resultVal.ToByte())
}

func getConfig(w http.ResponseWriter, req *http.Request) {

	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	if err := req.ParseForm(); err != nil {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_HTTP_PARSE
		w.Write(resultVal.ToByte())
		WasteLibrary.LogErr(err)
		return
	}
	resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_LINK, req.Host)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_NOTFOUND
		w.Write(resultVal.ToByte())
		return
	}
	var customerId string = resultVal.Retval.(string)

	resultVal = checkAuth(req.Form, customerId)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		w.Write(resultVal.ToByte())
		return
	}

	var currentHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.StringToHttpClientHeaderType(req.FormValue(WasteLibrary.HTTP_HEADER))
	if currentHttpHeader.OpType == WasteLibrary.OPTYPE_CUSTOMERCONFIG {
		resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_CUSTOMERCONFIG, customerId)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_GET
			w.Write(resultVal.ToByte())
			return
		}
	} else if currentHttpHeader.OpType == WasteLibrary.OPTYPE_ADMINCONFIG {
		resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_ADMINCONFIG, customerId)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_GET
			w.Write(resultVal.ToByte())
			return
		}
	} else if currentHttpHeader.OpType == WasteLibrary.OPTYPE_LOCALCONFIG {
		resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_LOCALCONFIG, customerId)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_GET
			w.Write(resultVal.ToByte())
			return
		}
	} else {
		resultVal.Result = WasteLibrary.RESULT_FAIL
	}

	w.Write(resultVal.ToByte())
}

func getUsers(w http.ResponseWriter, req *http.Request) {

	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	if err := req.ParseForm(); err != nil {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_HTTP_PARSE
		w.Write(resultVal.ToByte())
		WasteLibrary.LogErr(err)
		return
	}
	resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_LINK, req.Host)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_NOTFOUND
		w.Write(resultVal.ToByte())
		return
	}
	var customerId string = resultVal.Retval.(string)

	resultVal = checkAuth(req.Form, customerId)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		w.Write(resultVal.ToByte())
		return
	}

	resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_USERS, customerId)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_GET
		w.Write(resultVal.ToByte())
		return
	}

	var customerUsers WasteLibrary.CustomerUsersType = WasteLibrary.StringToCustomerUsersType(resultVal.Retval.(string))
	var customerUsersList WasteLibrary.CustomerUsersListType = WasteLibrary.CustomerUsersListType{
		CustomerId: WasteLibrary.StringIdToFloat64(customerId),
		Users:      make(map[string]WasteLibrary.UserType),
	}
	for _, userId := range customerUsers.Users {

		if userId != 0 {
			resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_USERS, WasteLibrary.Float64IdToString(userId))
			if resultVal.Result == WasteLibrary.RESULT_OK {
				var currentUser WasteLibrary.UserType = WasteLibrary.StringToUserType(resultVal.Retval.(string))
				customerUsersList.Users[currentUser.ToIdString()] = currentUser
			}

		}
	}
	resultVal.Result = WasteLibrary.RESULT_OK
	resultVal.Retval = customerUsersList.ToString()

	w.Write(resultVal.ToByte())
}

func checkAuth(data url.Values, customerId string) WasteLibrary.ResultType {
	return WasteLibrary.CheckAuth(data, customerId, "ADMIN")

}
