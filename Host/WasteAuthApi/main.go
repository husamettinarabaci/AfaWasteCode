package main

import (
	"net/http"
	"net/url"
	"time"

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
	http.HandleFunc("/register", register)
	http.HandleFunc("/login", login)
	http.HandleFunc("/checkAuth", checkAuth)
	http.ListenAndServe(":80", nil)
}

func register(w http.ResponseWriter, req *http.Request) {

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

	var currentUser WasteLibrary.UserType = WasteLibrary.StringToUserType(req.FormValue(WasteLibrary.HTTP_DATA))
	resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_USERS, customerId)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_NOTFOUND
		w.Write(resultVal.ToByte())
		return
	}
	var currentCustomerUsers WasteLibrary.CustomerUsersType = WasteLibrary.StringToCustomerUsersType(resultVal.Retval.(string))

	for _, userId := range currentCustomerUsers.Users {
		resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_USERS, WasteLibrary.Float64IdToString(userId))
		var inRedisUser WasteLibrary.UserType = WasteLibrary.StringToUserType(resultVal.Retval.(string))
		if inRedisUser.UserName == currentUser.UserName {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_USER_USERNAMEEXIST
			w.Write(resultVal.ToByte())
			return
		}
	}
	var userRole string = WasteLibrary.USER_ROLE_GUEST
	if len(currentCustomerUsers.Users) == 1 {
		userRole = WasteLibrary.USER_ROLE_ADMIN
	}

	if currentUser.Password == "" {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_USER_PASSWORDEMPTY
		w.Write(resultVal.ToByte())
		return
	}

	currentUser.UserRole = userRole
	currentUser.CustomerId = WasteLibrary.StringIdToFloat64(customerId)
	currentUser.Password = WasteLibrary.GetMD5Hash(currentUser.Password)

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
		WasteLibrary.HTTP_DATA:   {currentUser.ToString()},
	}
	resultVal = WasteLibrary.SaveConfigDbMainForStoreApi(data)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
		w.Write(resultVal.ToByte())
		return
	}
	currentUser = WasteLibrary.StringToUserType(resultVal.Retval.(string))
	resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_USERS, currentUser.ToIdString(), currentUser.ToString())
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
		w.Write(resultVal.ToByte())
		return
	}
	currentCustomerUsers.Users[currentUser.ToIdString()] = currentUser.UserId
	resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_USERS, currentCustomerUsers.ToCustomerIdString(), currentCustomerUsers.ToString())
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
		w.Write(resultVal.ToByte())
		return
	}
	w.Write(resultVal.ToByte())
}

func login(w http.ResponseWriter, req *http.Request) {

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

	var currentUser WasteLibrary.UserType = WasteLibrary.StringToUserType(req.FormValue(WasteLibrary.HTTP_DATA))
	resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_USERS, customerId)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_NOTFOUND
		w.Write(resultVal.ToByte())
		return
	}
	var currentCustomerUsers WasteLibrary.CustomerUsersType = WasteLibrary.StringToCustomerUsersType(resultVal.Retval.(string))

	var userExist bool = false
	for _, userId := range currentCustomerUsers.Users {
		resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_USERS, WasteLibrary.Float64IdToString(userId))
		var inRedisUser WasteLibrary.UserType = WasteLibrary.StringToUserType(resultVal.Retval.(string))
		if inRedisUser.UserName == currentUser.UserName {
			userExist = true

			if WasteLibrary.GetMD5Hash(currentUser.Password) != inRedisUser.Password {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_USER_INVALIDPASSWORD
				w.Write(resultVal.ToByte())
				return
			}
			break
		}
	}
	if !userExist {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_USER_INVALIDUSER
		w.Write(resultVal.ToByte())
		return
	}

	var token string = WasteLibrary.GenerateToken(currentUser.UserName + currentUser.Password + currentUser.Email)
	newDate := WasteLibrary.GetTimePlus(time.Hour * 1)
	resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_TOKEN, token, newDate)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
		w.Write(resultVal.ToByte())
		return
	}
	resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_TOKEN_USER, token, currentUser.ToIdString())
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
		w.Write(resultVal.ToByte())
		return
	}
	resultVal.Result = WasteLibrary.RESULT_OK
	resultVal.Retval = token

	w.Write(resultVal.ToByte())
}

func checkAuth(w http.ResponseWriter, req *http.Request) {

	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	if err := req.ParseForm(); err != nil {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_HTTP_PARSE
		w.Write(resultVal.ToByte())
		WasteLibrary.LogErr(err)
		return
	}

	var currentHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.StringToHttpClientHeaderType(req.FormValue(WasteLibrary.HTTP_HEADER))
	resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_TOKEN, currentHttpHeader.Token)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_USER_INVALIDTOKEN
		w.Write(resultVal.ToByte())
		return
	}

	endDate := WasteLibrary.StringToTime(resultVal.Retval.(string))

	if time.Since(endDate).Seconds() > -1 {
		newDate := WasteLibrary.GetTimePlus(time.Hour * 1)
		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_TOKEN, currentHttpHeader.Token, newDate)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())
			return
		}
	} else {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_USER_ENDTOKEN
		w.Write(resultVal.ToByte())
		return
	}

	resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_TOKEN_USER, currentHttpHeader.Token)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_USER_INVALIDTOKEN
		w.Write(resultVal.ToByte())
		return
	}

	resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_USERS, resultVal.Retval.(string))
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_USER_INVALIDUSER
		w.Write(resultVal.ToByte())
		return
	}

	var currentUser WasteLibrary.UserType = WasteLibrary.StringToUserType(resultVal.Retval.(string))
	customerId := WasteLibrary.StringIdToFloat64(req.FormValue(WasteLibrary.HTTP_CUSTOMERID))
	if currentUser.CustomerId != customerId {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_USER_INVALIDUSER
		w.Write(resultVal.ToByte())
		return
	}

	reqRole := req.FormValue(WasteLibrary.HTTP_USERROLE)

	if currentUser.UserRole == WasteLibrary.USER_ROLE_ADMIN || (reqRole == WasteLibrary.USER_ROLE_REPORT && currentUser.UserRole == WasteLibrary.USER_ROLE_REPORT) {
		resultVal.Result = WasteLibrary.RESULT_OK
		resultVal.Retval = ""
		w.Write(resultVal.ToByte())
		return
	} else {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_USER_INVALIDUSERROLE
		w.Write(resultVal.ToByte())
		return
	}

}
