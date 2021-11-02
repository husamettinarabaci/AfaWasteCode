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
	http.HandleFunc("/setCustomer", setCustomer)
	http.HandleFunc("/getCustomer", getCustomer)
	http.HandleFunc("/getCustomers", getCustomers)
	http.HandleFunc("/setDevice", setDevice)
	http.HandleFunc("/getDevice", getDevice)
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
		currentHttpHeader.AppType = WasteLibrary.APPTYPE_AFATEK
		currentHttpHeader.ReaderType = WasteLibrary.OPTYPE_CUSTOMER
		currentHttpHeader.BaseDataType = WasteLibrary.BASETYPE_CUSTOMER
		var currentData WasteLibrary.CustomerType
		currentData.CustomerName = "AFATEK"
		currentData.CustomerLink = "afatek.aws.afatek.com.tr"
		currentData.RfIdApp = WasteLibrary.STATU_ACTIVE
		currentData.UltApp = WasteLibrary.STATU_ACTIVE
		currentData.RecyApp = WasteLibrary.STATU_ACTIVE
		currentData.Active = WasteLibrary.STATU_ACTIVE
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
		data = url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentData.ToString()},
		}
		resultVal = WasteLibrary.GetStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_GET
			w.Write(resultVal.ToByte())
			return
		}
		var currentCustomer WasteLibrary.CustomerType = WasteLibrary.StringToCustomerType(resultVal.Retval.(string))

		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_CUSTOMERS, currentCustomer.ToIdString(), currentCustomer.ToString())
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())
			return
		}
		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_LINK, currentCustomer.CustomerLink, currentCustomer.ToIdString())
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())
			return
		}
		var currentCustomerRfidDevices WasteLibrary.CustomerRfidDevicesType
		currentCustomerRfidDevices.New()
		currentCustomerRfidDevices.CustomerId = currentCustomer.CustomerId
		currentCustomerRfidDevices.Devices["0"] = 0

		var currentCustomerRecyDevices WasteLibrary.CustomerRecyDevicesType
		currentCustomerRecyDevices.New()
		currentCustomerRecyDevices.CustomerId = currentCustomer.CustomerId
		currentCustomerRecyDevices.Devices["0"] = 0

		var currentCustomerUltDevices WasteLibrary.CustomerUltDevicesType
		currentCustomerUltDevices.New()
		currentCustomerUltDevices.CustomerId = currentCustomer.CustomerId
		currentCustomerUltDevices.Devices["0"] = 0

		var currentCustomerUsers WasteLibrary.CustomerUsersType
		currentCustomerUsers.New()
		currentCustomerUsers.CustomerId = currentCustomer.CustomerId
		currentCustomerUsers.Users["0"] = 0

		var currentCustomerTags WasteLibrary.CustomerTagsType
		currentCustomerTags.New()
		currentCustomerTags.CustomerId = currentCustomer.CustomerId
		currentCustomerTags.Tags["0"] = 0

		var currentCustomerConfig WasteLibrary.CustomerConfigType
		currentCustomerConfig.New()
		currentCustomerConfig.CustomerId = currentCustomer.CustomerId

		var currentAdminConfig WasteLibrary.AdminConfigType
		currentAdminConfig.New()
		currentAdminConfig.CustomerId = currentCustomer.CustomerId

		var currentLocalConfig WasteLibrary.LocalConfigType
		currentLocalConfig.New()
		currentLocalConfig.CustomerId = currentCustomer.CustomerId

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
			currentCustomers.Customers[currentCustomer.ToIdString()] = currentCustomer.CustomerId
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

func setCustomer(w http.ResponseWriter, req *http.Request) {
	if WasteLibrary.AllowCors {

		w.Header().Set("Access-Control-Allow-Origin", "*")
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
	currentHttpHeader.AppType = WasteLibrary.APPTYPE_AFATEK
	currentHttpHeader.ReaderType = WasteLibrary.OPTYPE_CUSTOMER
	currentHttpHeader.BaseDataType = WasteLibrary.BASETYPE_CUSTOMER
	var currentData WasteLibrary.CustomerType = WasteLibrary.StringToCustomerType(req.FormValue(WasteLibrary.HTTP_DATA))
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
	data = url.Values{
		WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
		WasteLibrary.HTTP_DATA:   {currentData.ToString()},
	}
	resultVal = WasteLibrary.GetStaticDbMainForStoreApi(data)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_GET
		w.Write(resultVal.ToByte())
		return
	}
	var currentCustomer WasteLibrary.CustomerType = WasteLibrary.StringToCustomerType(resultVal.Retval.(string))

	resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_CUSTOMERS, currentCustomer.ToIdString(), currentCustomer.ToString())
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
		w.Write(resultVal.ToByte())
		return
	}
	resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_LINK, currentCustomer.CustomerLink, currentCustomer.ToIdString())
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
		w.Write(resultVal.ToByte())
		return
	}
	if !isCustomerExist {
		var currentCustomerRfidDevices WasteLibrary.CustomerRfidDevicesType
		currentCustomerRfidDevices.New()
		currentCustomerRfidDevices.CustomerId = currentCustomer.CustomerId
		currentCustomerRfidDevices.Devices["0"] = 0

		var currentCustomerRecyDevices WasteLibrary.CustomerRecyDevicesType
		currentCustomerRecyDevices.New()
		currentCustomerRecyDevices.CustomerId = currentCustomer.CustomerId
		currentCustomerRecyDevices.Devices["0"] = 0

		var currentCustomerUltDevices WasteLibrary.CustomerUltDevicesType
		currentCustomerUltDevices.New()
		currentCustomerUltDevices.CustomerId = currentCustomer.CustomerId
		currentCustomerUltDevices.Devices["0"] = 0

		var currentCustomerUsers WasteLibrary.CustomerUsersType
		currentCustomerUsers.New()
		currentCustomerUsers.CustomerId = currentCustomer.CustomerId
		currentCustomerUsers.Users["0"] = 0

		var currentCustomerTags WasteLibrary.CustomerTagsType
		currentCustomerTags.New()
		currentCustomerTags.CustomerId = currentCustomer.CustomerId
		currentCustomerTags.Tags["0"] = 0

		var currentCustomerConfig WasteLibrary.CustomerConfigType
		currentCustomerConfig.New()
		currentCustomerConfig.CustomerId = currentCustomer.CustomerId

		var currentAdminConfig WasteLibrary.AdminConfigType
		currentAdminConfig.New()
		currentAdminConfig.CustomerId = currentCustomer.CustomerId

		var currentLocalConfig WasteLibrary.LocalConfigType
		currentLocalConfig.New()
		currentLocalConfig.CustomerId = currentCustomer.CustomerId

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
			currentCustomers.Customers[currentCustomer.ToIdString()] = currentCustomer.CustomerId
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

func getCustomers(w http.ResponseWriter, req *http.Request) {
	if WasteLibrary.AllowCors {

		w.Header().Set("Access-Control-Allow-Origin", "*")
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
	currentHttpHeader.AppType = WasteLibrary.APPTYPE_AFATEK
	currentHttpHeader.ReaderType = WasteLibrary.OPTYPE_DEVICE
	currentHttpHeader.BaseDataType = WasteLibrary.BASETYPE_DEVICE
	if currentHttpHeader.DeviceType == WasteLibrary.DEVICE_TYPE_RFID {
		var currentData WasteLibrary.RfidDeviceType = WasteLibrary.StringToRfidDeviceType(req.FormValue(WasteLibrary.HTTP_DATA))
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

		currentData.DeviceId = WasteLibrary.StringIdToFloat64(resultVal.Retval.(string))
		data = url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentData.ToString()},
		}
		resultVal = WasteLibrary.GetStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_GET
			w.Write(resultVal.ToByte())
			return
		}
		var currentDevice WasteLibrary.RfidDeviceType = WasteLibrary.StringToRfidDeviceType(resultVal.Retval.(string))

		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_RFID_DEVICES, currentDevice.ToIdString(), currentDevice.ToString())
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())
			return
		}

		resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_RFID_DEVICES, currentDevice.ToCustomerIdString())
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_GET
			w.Write(resultVal.ToByte())
			return
		}

		var customerDevices WasteLibrary.CustomerRfidDevicesType = WasteLibrary.StringToCustomerRfidDevicesType(resultVal.Retval.(string))
		customerDevices.Devices[currentDevice.ToIdString()] = currentDevice.DeviceId
		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_RFID_DEVICES, customerDevices.ToIdString(), customerDevices.ToString())
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())
			return
		}
	} else if currentHttpHeader.DeviceType == WasteLibrary.DEVICE_TYPE_ULT {
		var currentData WasteLibrary.UltDeviceType = WasteLibrary.StringToUltDeviceType(req.FormValue(WasteLibrary.HTTP_DATA))
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

		currentData.DeviceId = WasteLibrary.StringIdToFloat64(resultVal.Retval.(string))
		data = url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentData.ToString()},
		}
		resultVal = WasteLibrary.GetStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_GET
			w.Write(resultVal.ToByte())
			return
		}
		var currentDevice WasteLibrary.UltDeviceType = WasteLibrary.StringToUltDeviceType(resultVal.Retval.(string))

		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_ULT_DEVICES, currentDevice.ToIdString(), currentDevice.ToString())
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())
			return
		}

		resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_ULT_DEVICES, currentDevice.ToCustomerIdString())
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_GET
			w.Write(resultVal.ToByte())
			return
		}

		var customerDevices WasteLibrary.CustomerUltDevicesType = WasteLibrary.StringToCustomerUltDevicesType(resultVal.Retval.(string))
		customerDevices.Devices[currentDevice.ToIdString()] = currentDevice.DeviceId
		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_ULT_DEVICES, customerDevices.ToIdString(), customerDevices.ToString())
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())
			return
		}
	} else if currentHttpHeader.DeviceType == WasteLibrary.DEVICE_TYPE_RECY {
		var currentData WasteLibrary.RecyDeviceType = WasteLibrary.StringToRecyDeviceType(req.FormValue(WasteLibrary.HTTP_DATA))
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

		currentData.DeviceId = WasteLibrary.StringIdToFloat64(resultVal.Retval.(string))
		data = url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentData.ToString()},
		}
		resultVal = WasteLibrary.GetStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_GET
			w.Write(resultVal.ToByte())
			return
		}
		var currentDevice WasteLibrary.RecyDeviceType = WasteLibrary.StringToRecyDeviceType(resultVal.Retval.(string))

		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_RECY_DEVICES, currentDevice.ToIdString(), currentDevice.ToString())
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())
			return
		}

		resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_RECY_DEVICES, currentDevice.ToCustomerIdString())
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_GET
			w.Write(resultVal.ToByte())
			return
		}

		var customerDevices WasteLibrary.CustomerRecyDevicesType = WasteLibrary.StringToCustomerRecyDevicesType(resultVal.Retval.(string))
		customerDevices.Devices[currentDevice.ToIdString()] = currentDevice.DeviceId
		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_RECY_DEVICES, customerDevices.ToIdString(), customerDevices.ToString())
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())
			return
		}
	} else {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_DEVICE_NOTFOUND
	}

	w.Write(resultVal.ToByte())
}

func getDevices(w http.ResponseWriter, req *http.Request) {
	if WasteLibrary.AllowCors {

		w.Header().Set("Access-Control-Allow-Origin", "*")
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
	if currentHttpHeader.DeviceType == WasteLibrary.DEVICE_TYPE_RFID {
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
				resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_RFID_DEVICES, WasteLibrary.Float64IdToString(deviceId))
				if resultVal.Result == WasteLibrary.RESULT_OK {
					var currentDevice WasteLibrary.RfidDeviceType = WasteLibrary.StringToRfidDeviceType(resultVal.Retval.(string))
					customerDevicesList.Devices[currentDevice.ToIdString()] = currentDevice
				}
			}
		}

		resultVal.Result = WasteLibrary.RESULT_OK
		resultVal.Retval = customerDevicesList.ToString()
	} else if currentHttpHeader.DeviceType == WasteLibrary.DEVICE_TYPE_ULT {
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
				resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_ULT_DEVICES, WasteLibrary.Float64IdToString(deviceId))
				if resultVal.Result == WasteLibrary.RESULT_OK {
					var currentDevice WasteLibrary.UltDeviceType = WasteLibrary.StringToUltDeviceType(resultVal.Retval.(string))
					customerDevicesList.Devices[currentDevice.ToIdString()] = currentDevice
				}
			}
		}

		resultVal.Result = WasteLibrary.RESULT_OK
		resultVal.Retval = customerDevicesList.ToString()
	} else if currentHttpHeader.DeviceType == WasteLibrary.DEVICE_TYPE_RECY {
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
				resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_RECY_DEVICES, WasteLibrary.Float64IdToString(deviceId))
				if resultVal.Result == WasteLibrary.RESULT_OK {
					var currentDevice WasteLibrary.RecyDeviceType = WasteLibrary.StringToRecyDeviceType(resultVal.Retval.(string))
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
	if currentHttpHeader.DeviceType == WasteLibrary.DEVICE_TYPE_RFID {
		var currentDevice WasteLibrary.RfidDeviceType = WasteLibrary.StringToRfidDeviceType(req.FormValue(WasteLibrary.HTTP_DATA))
		resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_RFID_DEVICES, currentDevice.ToIdString())
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DEVICES_NOTFOUND
			w.Write(resultVal.ToByte())
			return
		}

		currentDevice = WasteLibrary.StringToRfidDeviceType(resultVal.Retval.(string))

		resultVal.Result = WasteLibrary.RESULT_OK
		resultVal.Retval = currentDevice.ToString()
	} else if currentHttpHeader.DeviceType == WasteLibrary.DEVICE_TYPE_ULT {
		var currentDevice WasteLibrary.UltDeviceType = WasteLibrary.StringToUltDeviceType(req.FormValue(WasteLibrary.HTTP_DATA))
		resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_ULT_DEVICES, currentDevice.ToIdString())
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DEVICES_NOTFOUND
			w.Write(resultVal.ToByte())
			return
		}

		currentDevice = WasteLibrary.StringToUltDeviceType(resultVal.Retval.(string))

		resultVal.Result = WasteLibrary.RESULT_OK
		resultVal.Retval = currentDevice.ToString()
	} else if currentHttpHeader.DeviceType == WasteLibrary.DEVICE_TYPE_RECY {
		var currentDevice WasteLibrary.RecyDeviceType = WasteLibrary.StringToRecyDeviceType(req.FormValue(WasteLibrary.HTTP_DATA))
		resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_RECY_DEVICES, currentDevice.ToIdString())
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DEVICES_NOTFOUND
			w.Write(resultVal.ToByte())
			return
		}

		currentDevice = WasteLibrary.StringToRecyDeviceType(resultVal.Retval.(string))

		resultVal.Result = WasteLibrary.RESULT_OK
		resultVal.Retval = currentDevice.ToString()
	} else {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_DEVICES_NOTFOUND
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
