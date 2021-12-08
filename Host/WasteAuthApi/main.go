package main

import (
	"net/http"
	"time"

	"github.com/devafatek/WasteLibrary"
)

func initStart() {

	WasteLibrary.LogStr("Successfully connected!")
	go WasteLibrary.InitLog()
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

	if WasteLibrary.AllowCors {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,access-control-allow-origin, access-control-allow-headers")
	}

	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL

	if err := req.ParseForm(); err != nil {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_HTTP_PARSE
		w.Write(resultVal.ToByte())

		WasteLibrary.LogErr(err)
		return
	}

	var linkCustomer WasteLibrary.CustomerType
	resultVal = linkCustomer.GetByRedisByLink(req.Host)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_NOTFOUND
		w.Write(resultVal.ToByte())

		return
	}

	var currentUser WasteLibrary.UserType = WasteLibrary.StringToUserType(req.FormValue(WasteLibrary.HTTP_DATA))
	var currentCustomerUsers WasteLibrary.CustomerUsersType
	currentCustomerUsers.CustomerId = linkCustomer.CustomerId
	resultVal = currentCustomerUsers.GetByRedis()
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_NOTFOUND
		w.Write(resultVal.ToByte())

		return
	}

	for _, userId := range currentCustomerUsers.Users {
		if userId != 0 {
			var inRedisUser WasteLibrary.UserType
			inRedisUser.UserId = userId
			resultVal = inRedisUser.GetByRedis()
			if resultVal.Result == WasteLibrary.RESULT_OK {

				if inRedisUser.Email == currentUser.Email && inRedisUser.Email != "" {
					resultVal.Result = WasteLibrary.RESULT_FAIL
					resultVal.Retval = WasteLibrary.RESULT_ERROR_USER_AUTH
					w.Write(resultVal.ToByte())

					return
				}
			}
		}
	}
	var userRole string = WasteLibrary.USER_ROLE_GUEST
	if len(currentCustomerUsers.Users) == 1 {
		userRole = WasteLibrary.USER_ROLE_ADMIN
	}

	if currentUser.Email == "" {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_USER_AUTH
		w.Write(resultVal.ToByte())

		return
	}

	if currentUser.Password == "" {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_USER_AUTH
		w.Write(resultVal.ToByte())

		return
	}

	currentUser.UserRole = userRole
	currentUser.CustomerId = linkCustomer.CustomerId
	currentUser.Password = WasteLibrary.GetMD5Hash(currentUser.Password)
	currentUser.Active = WasteLibrary.STATU_ACTIVE
	currentUser.CreateTime = WasteLibrary.GetTime()
	resultVal = currentUser.SaveToDb()
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
		w.Write(resultVal.ToByte())

		return
	}

	resultVal = currentUser.SaveToRedis()
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
		w.Write(resultVal.ToByte())

		return
	}
	currentCustomerUsers.Users[currentUser.ToIdString()] = currentUser.UserId
	resultVal = currentCustomerUsers.SaveToRedis()
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
		w.Write(resultVal.ToByte())

		return
	}
	w.Write(resultVal.ToByte())

}

func login(w http.ResponseWriter, req *http.Request) {

	if WasteLibrary.AllowCors {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,access-control-allow-origin, access-control-allow-headers")
	}

	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL

	if err := req.ParseForm(); err != nil {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_HTTP_PARSE
		w.Write(resultVal.ToByte())

		WasteLibrary.LogErr(err)
		return
	}

	var linkCustomer WasteLibrary.CustomerType
	resultVal = linkCustomer.GetByRedisByLink(req.Host)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_NOTFOUND
		w.Write(resultVal.ToByte())

		return
	}

	var currentUser WasteLibrary.UserType = WasteLibrary.StringToUserType(req.FormValue(WasteLibrary.HTTP_DATA))
	var currentCustomerUsers WasteLibrary.CustomerUsersType
	currentCustomerUsers.CustomerId = linkCustomer.CustomerId
	resultVal = currentCustomerUsers.GetByRedis()
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_NOTFOUND
		w.Write(resultVal.ToByte())

		return
	}

	var userExist bool = false
	for _, userId := range currentCustomerUsers.Users {
		if userId != 0 {
			var inRedisUser WasteLibrary.UserType
			inRedisUser.UserId = userId
			resultVal = inRedisUser.GetByRedis()
			if resultVal.Result == WasteLibrary.RESULT_OK {
				if inRedisUser.Email == currentUser.Email {
					userExist = true

					if WasteLibrary.GetMD5Hash(currentUser.Password) != inRedisUser.Password {
						resultVal.Result = WasteLibrary.RESULT_FAIL
						resultVal.Retval = WasteLibrary.RESULT_ERROR_USER_AUTH
						w.Write(resultVal.ToByte())

						return
					}
					currentUser = inRedisUser
					break
				}
			}
		}
	}
	if !userExist {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_USER_AUTH
		w.Write(resultVal.ToByte())

		return
	}

	var token string = WasteLibrary.GenerateToken(currentUser.Email+currentUser.Password+currentUser.Email+WasteLibrary.GetTime(), currentUser.ToIdString())
	newDate := WasteLibrary.GetTimePlus(time.Hour * 1)
	resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_USER_TOKEN, currentUser.ToIdString(), token)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
		w.Write(resultVal.ToByte())

		return
	}
	resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_USER_TOKENENDDATE, currentUser.ToIdString(), newDate)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
		w.Write(resultVal.ToByte())

		return
	}
	currentUser.Token = token
	currentUser.Password = ""
	resultVal.Result = WasteLibrary.RESULT_OK
	resultVal.Retval = currentUser.ToString()

	w.Write(resultVal.ToByte())

}

func checkAuth(w http.ResponseWriter, req *http.Request) {

	if WasteLibrary.AllowCors {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,access-control-allow-origin, access-control-allow-headers")
	}

	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL

	if err := req.ParseForm(); err != nil {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_HTTP_PARSE
		w.Write(resultVal.ToByte())

		WasteLibrary.LogErr(err)
		return
	}

	var token string = req.FormValue(WasteLibrary.HTTP_TOKEN)
	var userIdByToken string = WasteLibrary.GetUserIdByToken(token)
	resultVal = WasteLibrary.GetRedisForStoreApi("0", WasteLibrary.REDIS_USER_TOKEN, userIdByToken)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_USER_AUTH
		w.Write(resultVal.ToByte())

		return
	}
	if token != resultVal.Retval.(string) {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_USER_AUTH
		w.Write(resultVal.ToByte())

		return
	}
	resultVal = WasteLibrary.GetRedisForStoreApi("0", WasteLibrary.REDIS_USER_TOKENENDDATE, userIdByToken)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_USER_AUTH
		w.Write(resultVal.ToByte())

		return
	}

	endDate := WasteLibrary.StringToTime(resultVal.Retval.(string))
	if time.Since(endDate).Seconds() < -1 {
		newDate := WasteLibrary.GetTimePlus(time.Hour * 1)
		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_USER_TOKENENDDATE, userIdByToken, newDate)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
	} else {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_USER_AUTH
		w.Write(resultVal.ToByte())

		return
	}

	var currentUser WasteLibrary.UserType
	currentUser.UserId = WasteLibrary.StringIdToFloat64(userIdByToken)
	resultVal = currentUser.GetByRedis()
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_USER_AUTH
		w.Write(resultVal.ToByte())

		return
	}

	customerId := WasteLibrary.StringIdToFloat64(req.FormValue(WasteLibrary.HTTP_CUSTOMERID))
	if currentUser.CustomerId != customerId {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_USER_AUTH
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
		resultVal.Retval = WasteLibrary.RESULT_ERROR_USER_AUTH
		w.Write(resultVal.ToByte())

		return
	}

}
