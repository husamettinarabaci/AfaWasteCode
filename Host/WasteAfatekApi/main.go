package main

import (
	"net/http"
	"net/url"

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
	http.HandleFunc("/logApp", logApp)
	http.HandleFunc("/setCustomer", setCustomer)
	http.HandleFunc("/getCustomer", getCustomer)
	http.HandleFunc("/getCustomers", getCustomers)
	http.HandleFunc("/setDevice", setDevice)
	http.HandleFunc("/getDevice", getDevice)
	http.HandleFunc("/getConfig", getConfig)
	http.HandleFunc("/setConfig", setConfig)
	http.HandleFunc("/getDevices", getDevices)
	http.HandleFunc("/startSystem", startSystem)
	http.ListenAndServe(":80", nil)
}

func startSystem(w http.ResponseWriter, req *http.Request) {

	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL

	if err := req.ParseForm(); err != nil {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_HTTP_PARSE
		w.Write(resultVal.ToByte())

		WasteLibrary.LogErr(err)
		return
	}

	resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMERS, "1")
	if resultVal.Result == WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_APP_STARTED
		w.Write(resultVal.ToByte())

		return
	} else {
		WasteLibrary.LogStr("AfatekApi Start System Add Customer AFATEK")
		var currentHttpHeader WasteLibrary.HttpClientHeaderType
		currentHttpHeader.New()
		currentHttpHeader.DataType = WasteLibrary.DATATYPE_CUSTOMER
		var currentData WasteLibrary.CustomerType
		currentData.CustomerName = "AFATEK"
		currentData.CustomerLink = "afatek.aws.afatek.com.tr"
		currentData.RfIdApp = WasteLibrary.STATU_ACTIVE
		currentData.UltApp = WasteLibrary.STATU_ACTIVE
		currentData.RecyApp = WasteLibrary.STATU_ACTIVE
		currentData.Active = WasteLibrary.STATU_ACTIVE
		currentData.CreateTime = WasteLibrary.GetTime()
		data := url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentData.ToString()},
		}
		WasteLibrary.LogStr("AfatekApi Send Header : " + currentHttpHeader.ToString())
		WasteLibrary.LogStr("AfatekApi Send Data : " + currentData.ToString())
		resultVal = WasteLibrary.SaveStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		currentData.CustomerId = WasteLibrary.StringIdToFloat64(resultVal.Retval.(string))

		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_CUSTOMERS, currentData.ToIdString(), currentData.ToString())
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_LINK, currentData.CustomerLink, currentData.ToIdString())
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		var currentCustomerRfidDevices WasteLibrary.CustomerRfidDevicesType
		currentCustomerRfidDevices.New()
		currentCustomerRfidDevices.CustomerId = currentData.CustomerId
		currentCustomerRfidDevices.Devices["0"] = 0

		var currentCustomerRecyDevices WasteLibrary.CustomerRecyDevicesType
		currentCustomerRecyDevices.New()
		currentCustomerRecyDevices.CustomerId = currentData.CustomerId
		currentCustomerRecyDevices.Devices["0"] = 0

		var currentCustomerUltDevices WasteLibrary.CustomerUltDevicesType
		currentCustomerUltDevices.New()
		currentCustomerUltDevices.CustomerId = currentData.CustomerId
		currentCustomerUltDevices.Devices["0"] = 0

		var currentCustomerUsers WasteLibrary.CustomerUsersType
		currentCustomerUsers.New()
		currentCustomerUsers.CustomerId = currentData.CustomerId
		currentCustomerUsers.Users["0"] = 0

		var currentCustomerTags WasteLibrary.CustomerTagsType
		currentCustomerTags.New()
		currentCustomerTags.CustomerId = currentData.CustomerId
		currentCustomerTags.Tags["0"] = 0

		var currentCustomerConfig WasteLibrary.CustomerConfigType
		currentCustomerConfig.New()
		currentCustomerConfig.CustomerId = currentData.CustomerId

		var currentAdminConfig WasteLibrary.AdminConfigType
		currentAdminConfig.New()
		currentAdminConfig.CustomerId = currentData.CustomerId

		var currentLocalConfig WasteLibrary.LocalConfigType
		currentLocalConfig.New()
		currentLocalConfig.CustomerId = currentData.CustomerId

		WasteLibrary.LogStr("CustomerRfidDevices : " + currentCustomerRfidDevices.ToString())
		WasteLibrary.LogStr("CustomerRecyDevices : " + currentCustomerRecyDevices.ToString())
		WasteLibrary.LogStr("CustomerUltDevices : " + currentCustomerUltDevices.ToString())
		WasteLibrary.LogStr("CustomerUsers : " + currentCustomerUsers.ToString())
		WasteLibrary.LogStr("CustomerTags : " + currentCustomerTags.ToString())
		if resultVal.Result == WasteLibrary.RESULT_OK {
			resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMERS, WasteLibrary.REDIS_CUSTOMERS)
			var currentCustomers WasteLibrary.CustomersType
			if resultVal.Result == WasteLibrary.RESULT_OK {
				currentCustomers = WasteLibrary.StringToCustomersType(resultVal.Retval.(string))

			} else {
				currentCustomers = WasteLibrary.CustomersType{
					Customers: make(map[string]float64),
				}
			}
			currentCustomers.Customers[currentData.ToIdString()] = currentData.CustomerId
			resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_CUSTOMERS, WasteLibrary.REDIS_CUSTOMERS, currentCustomers.ToString())
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
				w.Write(resultVal.ToByte())

				return
			}
		}
		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_RFID_DEVICES, currentCustomerRfidDevices.ToIdString(), currentCustomerRfidDevices.ToString())
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_RECY_DEVICES, currentCustomerRecyDevices.ToIdString(), currentCustomerRecyDevices.ToString())
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_ULT_DEVICES, currentCustomerUltDevices.ToIdString(), currentCustomerUltDevices.ToString())
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_USERS, currentCustomerUsers.ToIdString(), currentCustomerUsers.ToString())
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_TAGS, currentCustomerTags.ToIdString(), currentCustomerTags.ToString())
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_CUSTOMERCONFIG, currentCustomerConfig.ToIdString(), currentCustomerConfig.ToString())
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_ADMINCONFIG, currentAdminConfig.ToIdString(), currentAdminConfig.ToString())
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_LOCALCONFIG, currentLocalConfig.ToIdString(), currentLocalConfig.ToString())
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}

	}
	w.Write(resultVal.ToByte())

}

func logApp(w http.ResponseWriter, req *http.Request) {

	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL

	if err := req.ParseForm(); err != nil {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_HTTP_PARSE
		w.Write(resultVal.ToByte())

		WasteLibrary.LogErr(err)
		return
	}

	appTypes := req.URL.Query()["appType"]
	opTypes := req.URL.Query()["opType"]
	//data := url.Values{}
	//resultVal = WasteLibrary.HttpPostReq("http://waste-"+appTypes[0]+"-cluster-ip/"+opTypes[0], data)
	resultVal = WasteLibrary.PublishRedisForStoreApi(WasteLibrary.REDIS_APP_LOG_CHANNEL, appTypes[0], opTypes[0])

	w.Write(resultVal.ToByte())

}

func setCustomer(w http.ResponseWriter, req *http.Request) {

	if WasteLibrary.AllowCors {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
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
	currentHttpHeader.DataType = WasteLibrary.DATATYPE_CUSTOMER
	var currentData WasteLibrary.CustomerType = WasteLibrary.StringToCustomerType(req.FormValue(WasteLibrary.HTTP_DATA))
	var oldCustomer WasteLibrary.CustomerType
	if currentData.CustomerId != 0 {
		resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMERS, currentData.ToIdString())
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_NOTFOUND
			w.Write(resultVal.ToByte())

			return
		}
		oldCustomer = WasteLibrary.StringToCustomerType(resultVal.Retval.(string))
	}

	data := url.Values{
		WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
		WasteLibrary.HTTP_DATA:   {currentData.ToString()},
	}
	resultVal = WasteLibrary.SaveStaticDbMainForStoreApi(data)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
		w.Write(resultVal.ToByte())

		return
	}

	var isCustomerExist = false
	if currentData.CustomerId != 0 {
		isCustomerExist = true
	}
	currentData.CustomerId = WasteLibrary.StringIdToFloat64(resultVal.Retval.(string))

	resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_CUSTOMERS, currentData.ToIdString(), currentData.ToString())
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
		w.Write(resultVal.ToByte())

		return
	}
	resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_LINK, currentData.CustomerLink, currentData.ToIdString())
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
		w.Write(resultVal.ToByte())

		return
	}
	if isCustomerExist {
		if oldCustomer.CustomerLink != currentData.CustomerLink {
			resultVal = WasteLibrary.DeleteRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_LINK, oldCustomer.CustomerLink)
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
				w.Write(resultVal.ToByte())

				return
			}
		}
	} else {
		var currentCustomerRfidDevices WasteLibrary.CustomerRfidDevicesType
		currentCustomerRfidDevices.New()
		currentCustomerRfidDevices.CustomerId = currentData.CustomerId
		currentCustomerRfidDevices.Devices["0"] = 0

		var currentCustomerRecyDevices WasteLibrary.CustomerRecyDevicesType
		currentCustomerRecyDevices.New()
		currentCustomerRecyDevices.CustomerId = currentData.CustomerId
		currentCustomerRecyDevices.Devices["0"] = 0

		var currentCustomerUltDevices WasteLibrary.CustomerUltDevicesType
		currentCustomerUltDevices.New()
		currentCustomerUltDevices.CustomerId = currentData.CustomerId
		currentCustomerUltDevices.Devices["0"] = 0

		var currentCustomerUsers WasteLibrary.CustomerUsersType
		currentCustomerUsers.New()
		currentCustomerUsers.CustomerId = currentData.CustomerId
		currentCustomerUsers.Users["0"] = 0

		var currentCustomerTags WasteLibrary.CustomerTagsType
		currentCustomerTags.New()
		currentCustomerTags.CustomerId = currentData.CustomerId
		currentCustomerTags.Tags["0"] = 0

		var currentCustomerConfig WasteLibrary.CustomerConfigType
		currentCustomerConfig.New()
		currentCustomerConfig.CustomerId = currentData.CustomerId

		var currentAdminConfig WasteLibrary.AdminConfigType
		currentAdminConfig.New()
		currentAdminConfig.CustomerId = currentData.CustomerId

		var currentLocalConfig WasteLibrary.LocalConfigType
		currentLocalConfig.New()
		currentLocalConfig.CustomerId = currentData.CustomerId

		WasteLibrary.LogStr("CustomerRfidDevices : " + currentCustomerRfidDevices.ToString())
		WasteLibrary.LogStr("CustomerRecyDevices : " + currentCustomerRecyDevices.ToString())
		WasteLibrary.LogStr("CustomerUltDevices : " + currentCustomerUltDevices.ToString())
		WasteLibrary.LogStr("CustomerUsers : " + currentCustomerUsers.ToString())
		WasteLibrary.LogStr("CustomerTags : " + currentCustomerTags.ToString())
		if resultVal.Result == WasteLibrary.RESULT_OK {
			resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMERS, WasteLibrary.REDIS_CUSTOMERS)
			var currentCustomers WasteLibrary.CustomersType
			if resultVal.Result == WasteLibrary.RESULT_OK {
				currentCustomers = WasteLibrary.StringToCustomersType(resultVal.Retval.(string))

			} else {
				currentCustomers = WasteLibrary.CustomersType{
					Customers: make(map[string]float64),
				}
			}
			currentCustomers.Customers[currentData.ToIdString()] = currentData.CustomerId
			resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_CUSTOMERS, WasteLibrary.REDIS_CUSTOMERS, currentCustomers.ToString())
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
				w.Write(resultVal.ToByte())

				return
			}
		}
		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_RFID_DEVICES, currentCustomerRfidDevices.ToIdString(), currentCustomerRfidDevices.ToString())
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_RECY_DEVICES, currentCustomerRecyDevices.ToIdString(), currentCustomerRecyDevices.ToString())
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_ULT_DEVICES, currentCustomerUltDevices.ToIdString(), currentCustomerUltDevices.ToString())
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_USERS, currentCustomerUsers.ToIdString(), currentCustomerUsers.ToString())
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_TAGS, currentCustomerTags.ToIdString(), currentCustomerTags.ToString())
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_CUSTOMERCONFIG, currentCustomerConfig.ToIdString(), currentCustomerConfig.ToString())
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_ADMINCONFIG, currentAdminConfig.ToIdString(), currentAdminConfig.ToString())
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_LOCALCONFIG, currentLocalConfig.ToIdString(), currentLocalConfig.ToString())
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
	}

	resultVal = WasteLibrary.PublishRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_CHANNEL+currentData.ToIdString(), WasteLibrary.DATATYPE_CUSTOMER, currentData.ToString())

	w.Write(resultVal.ToByte())

}

func getCustomers(w http.ResponseWriter, req *http.Request) {

	if WasteLibrary.AllowCors {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
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

	resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMERS, WasteLibrary.REDIS_CUSTOMERS)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMERS_NOTFOUND
		w.Write(resultVal.ToByte())

		return
	}

	var customers WasteLibrary.CustomersType = WasteLibrary.StringToCustomersType(resultVal.Retval.(string))
	var customersList WasteLibrary.CustomersListType
	customersList.New()
	for _, customerId := range customers.Customers {

		if customerId != 0 {
			resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMERS, WasteLibrary.Float64IdToString(customerId))
			if resultVal.Result == WasteLibrary.RESULT_OK {

				var currentCustomer WasteLibrary.CustomerType = WasteLibrary.StringToCustomerType(resultVal.Retval.(string))
				customersList.Customers[currentCustomer.ToIdString()] = currentCustomer
			}
		}
	}

	resultVal.Result = WasteLibrary.RESULT_OK
	resultVal.Retval = customersList.ToString()
	w.Write(resultVal.ToByte())

}

func getCustomer(w http.ResponseWriter, req *http.Request) {

	if WasteLibrary.AllowCors {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
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

	var currentData WasteLibrary.CustomerType = WasteLibrary.StringToCustomerType(req.FormValue(WasteLibrary.HTTP_DATA))

	resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMERS, currentData.ToIdString())
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_NOTFOUND
		w.Write(resultVal.ToByte())

		return
	}

	w.Write(resultVal.ToByte())

}

func setDevice(w http.ResponseWriter, req *http.Request) {

	if WasteLibrary.AllowCors {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
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
	if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_RFID {
		currentHttpHeader.DataType = WasteLibrary.DATATYPE_RFID_MAIN_DEVICE
		var currentData WasteLibrary.RfidDeviceType = WasteLibrary.StringToRfidDeviceType(req.FormValue(WasteLibrary.HTTP_DATA))
		var currentOldData WasteLibrary.RfidDeviceType
		currentOldData.DeviceId = currentData.DeviceId
		currentOldData.GetAll()
		data := url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentData.DeviceMain.ToString()},
		}
		resultVal = WasteLibrary.SaveStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		currentData.DeviceMain.DeviceId = WasteLibrary.StringIdToFloat64(resultVal.Retval.(string))

		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_RFID_MAIN_DEVICES, currentData.DeviceMain.ToIdString(), currentData.DeviceMain.ToString())
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		if currentData.DeviceMain.CustomerId != currentOldData.DeviceMain.CustomerId {
			resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_RFID_DEVICES, currentData.DeviceMain.ToCustomerIdString())
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_GET
				w.Write(resultVal.ToByte())

				return
			}

			var customerDevices WasteLibrary.CustomerRfidDevicesType = WasteLibrary.StringToCustomerRfidDevicesType(resultVal.Retval.(string))
			customerDevices.Devices[currentData.DeviceMain.ToIdString()] = currentData.DeviceMain.DeviceId
			resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_RFID_DEVICES, customerDevices.ToIdString(), customerDevices.ToString())
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
				w.Write(resultVal.ToByte())

				return
			}

			resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_RFID_DEVICES, currentOldData.DeviceMain.ToCustomerIdString())
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_GET
				w.Write(resultVal.ToByte())

				return
			}

			var oldCustomerDevices WasteLibrary.CustomerRfidDevicesType = WasteLibrary.StringToCustomerRfidDevicesType(resultVal.Retval.(string))
			delete(oldCustomerDevices.Devices, currentData.DeviceMain.ToIdString())
			resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_RFID_DEVICES, oldCustomerDevices.ToIdString(), oldCustomerDevices.ToString())
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
				w.Write(resultVal.ToByte())

				return
			}
		}

		resultVal = WasteLibrary.PublishRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_CHANNEL+currentData.DeviceMain.ToCustomerIdString(), WasteLibrary.DATATYPE_RFID_MAIN_DEVICE, currentData.DeviceMain.ToString())

	} else if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_ULT {
		currentHttpHeader.DataType = WasteLibrary.DATATYPE_ULT_MAIN_DEVICE
		var currentData WasteLibrary.UltDeviceType = WasteLibrary.StringToUltDeviceType(req.FormValue(WasteLibrary.HTTP_DATA))
		var currentOldData WasteLibrary.UltDeviceType
		currentOldData.DeviceId = currentData.DeviceId
		currentOldData.GetAll()
		data := url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentData.DeviceMain.ToString()},
		}
		resultVal = WasteLibrary.SaveStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		currentData.DeviceMain.DeviceId = WasteLibrary.StringIdToFloat64(resultVal.Retval.(string))

		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_ULT_MAIN_DEVICES, currentData.DeviceMain.ToIdString(), currentData.DeviceMain.ToString())
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		if currentData.DeviceMain.CustomerId != currentOldData.DeviceMain.CustomerId {
			resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_ULT_DEVICES, currentData.DeviceMain.ToCustomerIdString())
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_GET
				w.Write(resultVal.ToByte())

				return
			}

			var customerDevices WasteLibrary.CustomerUltDevicesType = WasteLibrary.StringToCustomerUltDevicesType(resultVal.Retval.(string))
			customerDevices.Devices[currentData.DeviceMain.ToIdString()] = currentData.DeviceMain.DeviceId
			resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_ULT_DEVICES, customerDevices.ToIdString(), customerDevices.ToString())
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
				w.Write(resultVal.ToByte())

				return
			}

			resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_ULT_DEVICES, currentOldData.DeviceMain.ToCustomerIdString())
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_GET
				w.Write(resultVal.ToByte())

				return
			}

			var oldCustomerDevices WasteLibrary.CustomerUltDevicesType = WasteLibrary.StringToCustomerUltDevicesType(resultVal.Retval.(string))
			delete(oldCustomerDevices.Devices, currentData.DeviceMain.ToIdString())
			resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_ULT_DEVICES, oldCustomerDevices.ToIdString(), oldCustomerDevices.ToString())
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
				w.Write(resultVal.ToByte())

				return
			}
		}
		resultVal = WasteLibrary.PublishRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_CHANNEL+currentData.DeviceMain.ToCustomerIdString(), WasteLibrary.DATATYPE_ULT_MAIN_DEVICE, currentData.DeviceMain.ToString())

	} else if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_RECY {
		currentHttpHeader.DataType = WasteLibrary.DATATYPE_RECY_MAIN_DEVICE
		var currentData WasteLibrary.RecyDeviceType = WasteLibrary.StringToRecyDeviceType(req.FormValue(WasteLibrary.HTTP_DATA))
		var currentOldData WasteLibrary.RecyDeviceType
		currentOldData.DeviceId = currentData.DeviceId
		currentOldData.GetAll()
		data := url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentData.DeviceMain.ToString()},
		}
		resultVal = WasteLibrary.SaveStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		currentData.DeviceMain.DeviceId = WasteLibrary.StringIdToFloat64(resultVal.Retval.(string))

		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_RECY_MAIN_DEVICES, currentData.DeviceMain.ToIdString(), currentData.DeviceMain.ToString())
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		if currentData.DeviceMain.CustomerId != currentOldData.DeviceMain.CustomerId {
			resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_RECY_DEVICES, currentData.DeviceMain.ToCustomerIdString())
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_GET
				w.Write(resultVal.ToByte())

				return
			}

			var customerDevices WasteLibrary.CustomerRecyDevicesType = WasteLibrary.StringToCustomerRecyDevicesType(resultVal.Retval.(string))
			customerDevices.Devices[currentData.DeviceMain.ToIdString()] = currentData.DeviceMain.DeviceId
			resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_RECY_DEVICES, customerDevices.ToIdString(), customerDevices.ToString())
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
				w.Write(resultVal.ToByte())

				return
			}

			resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_RECY_DEVICES, currentOldData.DeviceMain.ToCustomerIdString())
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_GET
				w.Write(resultVal.ToByte())

				return
			}

			var oldCustomerDevices WasteLibrary.CustomerRecyDevicesType = WasteLibrary.StringToCustomerRecyDevicesType(resultVal.Retval.(string))
			delete(oldCustomerDevices.Devices, currentData.DeviceMain.ToIdString())
			resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_RECY_DEVICES, oldCustomerDevices.ToIdString(), oldCustomerDevices.ToString())
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
				w.Write(resultVal.ToByte())

				return
			}
		}
		resultVal = WasteLibrary.PublishRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_CHANNEL+currentData.DeviceMain.ToCustomerIdString(), WasteLibrary.DATATYPE_RECY_MAIN_DEVICE, currentData.DeviceMain.ToString())

	} else {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_DEVICE_NOTFOUND
	}

	w.Write(resultVal.ToByte())

}

func getDevices(w http.ResponseWriter, req *http.Request) {

	if WasteLibrary.AllowCors {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
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
	if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_RFID {
		var customerDevicesList WasteLibrary.CustomerRfidDevicesListType = WasteLibrary.StringToCustomerRfidDevicesListType(req.FormValue(WasteLibrary.HTTP_DATA))
		resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_RFID_DEVICES, customerDevicesList.ToIdString())
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMERS_DEVICES_NOTFOUND
			w.Write(resultVal.ToByte())

			return
		}

		var customerDevices WasteLibrary.CustomerRfidDevicesType = WasteLibrary.StringToCustomerRfidDevicesType(resultVal.Retval.(string))
		customerDevicesList.Devices = make(map[string]WasteLibrary.RfidDeviceType)
		for _, deviceId := range customerDevices.Devices {

			if deviceId != 0 {
				var currentDevice WasteLibrary.RfidDeviceType
				currentDevice.New()
				currentDevice.DeviceId = deviceId
				resultVal = currentDevice.GetAll()
				if resultVal.Result == WasteLibrary.RESULT_OK {
					customerDevicesList.Devices[currentDevice.ToIdString()] = currentDevice
				}
			}
		}

		resultVal.Result = WasteLibrary.RESULT_OK
		resultVal.Retval = customerDevicesList.ToString()
	} else if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_ULT {
		var customerDevicesList WasteLibrary.CustomerUltDevicesListType = WasteLibrary.StringToCustomerUltDevicesListType(req.FormValue(WasteLibrary.HTTP_DATA))
		resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_ULT_DEVICES, customerDevicesList.ToIdString())
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMERS_DEVICES_NOTFOUND
			w.Write(resultVal.ToByte())

			return
		}

		var customerDevices WasteLibrary.CustomerUltDevicesType = WasteLibrary.StringToCustomerUltDevicesType(resultVal.Retval.(string))
		customerDevicesList.Devices = make(map[string]WasteLibrary.UltDeviceType)
		for _, deviceId := range customerDevices.Devices {

			if deviceId != 0 {
				var currentDevice WasteLibrary.UltDeviceType
				currentDevice.New()
				currentDevice.DeviceId = deviceId
				resultVal = currentDevice.GetAll()
				if resultVal.Result == WasteLibrary.RESULT_OK {
					customerDevicesList.Devices[currentDevice.ToIdString()] = currentDevice
				}
			}
		}

		resultVal.Result = WasteLibrary.RESULT_OK
		resultVal.Retval = customerDevicesList.ToString()
	} else if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_RECY {
		var customerDevicesList WasteLibrary.CustomerRecyDevicesListType = WasteLibrary.StringToCustomerRecyDevicesListType(req.FormValue(WasteLibrary.HTTP_DATA))
		resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_RECY_DEVICES, customerDevicesList.ToIdString())
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMERS_DEVICES_NOTFOUND
			w.Write(resultVal.ToByte())

			return
		}

		var customerDevices WasteLibrary.CustomerRecyDevicesType = WasteLibrary.StringToCustomerRecyDevicesType(resultVal.Retval.(string))
		customerDevicesList.Devices = make(map[string]WasteLibrary.RecyDeviceType)
		for _, deviceId := range customerDevices.Devices {

			if deviceId != 0 {
				var currentDevice WasteLibrary.RecyDeviceType
				currentDevice.New()
				currentDevice.DeviceId = deviceId
				resultVal = currentDevice.GetAll()
				if resultVal.Result == WasteLibrary.RESULT_OK {
					customerDevicesList.Devices[currentDevice.ToIdString()] = currentDevice
				}
			}
		}

		resultVal.Result = WasteLibrary.RESULT_OK
		resultVal.Retval = customerDevicesList.ToString()
	} else {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMERS_DEVICES_NOTFOUND
	}
	w.Write(resultVal.ToByte())

}

func getDevice(w http.ResponseWriter, req *http.Request) {

	if WasteLibrary.AllowCors {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
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
	if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_RFID {
		var currentDevice WasteLibrary.RfidDeviceType = WasteLibrary.StringToRfidDeviceType(req.FormValue(WasteLibrary.HTTP_DATA))
		resultVal = currentDevice.GetAll()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DEVICES_NOTFOUND
			w.Write(resultVal.ToByte())

			return
		}

		resultVal.Result = WasteLibrary.RESULT_OK
		resultVal.Retval = currentDevice.ToString()
	} else if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_ULT {
		var currentDevice WasteLibrary.UltDeviceType = WasteLibrary.StringToUltDeviceType(req.FormValue(WasteLibrary.HTTP_DATA))
		resultVal = currentDevice.GetAll()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DEVICES_NOTFOUND
			w.Write(resultVal.ToByte())

			return
		}

		resultVal.Result = WasteLibrary.RESULT_OK
		resultVal.Retval = currentDevice.ToString()
	} else if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_RECY {
		var currentDevice WasteLibrary.RecyDeviceType = WasteLibrary.StringToRecyDeviceType(req.FormValue(WasteLibrary.HTTP_DATA))
		resultVal = currentDevice.GetAll()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DEVICES_NOTFOUND
			w.Write(resultVal.ToByte())

			return
		}

		resultVal.Result = WasteLibrary.RESULT_OK
		resultVal.Retval = currentDevice.ToString()
	} else {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_DEVICES_NOTFOUND
	}
	w.Write(resultVal.ToByte())

}

func getConfig(w http.ResponseWriter, req *http.Request) {

	if WasteLibrary.AllowCors {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
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
	if currentHttpHeader.DataType == WasteLibrary.DATATYPE_CUSTOMERCONFIG {
		var currentConfig WasteLibrary.CustomerConfigType = WasteLibrary.StringToCustomerConfigType(req.FormValue(WasteLibrary.HTTP_DATA))
		resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_CUSTOMERCONFIG, currentConfig.ToIdString())
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_CUSTOMERCONFIG_NOTFOUND
			w.Write(resultVal.ToByte())

			return
		}
	} else {
		resultVal.Result = WasteLibrary.RESULT_FAIL
	}

	w.Write(resultVal.ToByte())

}

func setConfig(w http.ResponseWriter, req *http.Request) {

	if WasteLibrary.AllowCors {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
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
	if currentHttpHeader.DataType == WasteLibrary.DATATYPE_CUSTOMERCONFIG {
		var currentData WasteLibrary.CustomerConfigType = WasteLibrary.StringToCustomerConfigType(req.FormValue(WasteLibrary.HTTP_DATA))
		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_CUSTOMERCONFIG, currentData.ToIdString(), currentData.ToString())
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
	} else {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_DATATYPE
		w.Write(resultVal.ToByte())

		return
	}

	w.Write(resultVal.ToByte())

}

func checkAuth(data url.Values, customerId string) WasteLibrary.ResultType {
	var resultVal WasteLibrary.ResultType
	resultVal = WasteLibrary.CheckAuth(data, customerId, "ADMIN")

	if resultVal.Result == WasteLibrary.RESULT_OK {
		resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMERS, customerId)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_NOTFOUND
			return resultVal
		}
		var currentCustomer WasteLibrary.CustomerType = WasteLibrary.StringToCustomerType(resultVal.Retval.(string))
		if currentCustomer.CustomerName != "AFATEK" {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_INVALID
			return resultVal
		}
	}
	return resultVal

}
