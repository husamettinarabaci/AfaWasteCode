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

	var customerBase WasteLibrary.CustomerType
	customerBase.CustomerId = 1
	resultVal = customerBase.GetByRedis()
	if resultVal.Result == WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_APP_STARTED
		w.Write(resultVal.ToByte())

		return
	} else {
		WasteLibrary.LogStr("AfatekApi Start System Add Customer AFATEK")
		var currentData WasteLibrary.CustomerType
		currentData.CustomerName = "AFATEK"
		currentData.CustomerLink = "afatek.aws.afatek.com.tr"
		currentData.RfIdApp = WasteLibrary.STATU_ACTIVE
		currentData.UltApp = WasteLibrary.STATU_ACTIVE
		currentData.RecyApp = WasteLibrary.STATU_ACTIVE
		currentData.Active = WasteLibrary.STATU_ACTIVE
		currentData.CreateTime = WasteLibrary.GetTime()

		resultVal = currentData.SaveToDb()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		resultVal = currentData.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		resultVal = currentData.SaveToRedisLink()
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

		if resultVal.Result == WasteLibrary.RESULT_OK {
			var currentCustomers WasteLibrary.CustomersType
			resultVal = currentCustomers.GetByRedis()
			if resultVal.Result == WasteLibrary.RESULT_OK {

			} else {
				currentCustomers = WasteLibrary.CustomersType{
					Customers: make(map[string]float64),
				}
			}
			currentCustomers.Customers[currentData.ToIdString()] = currentData.CustomerId
			resultVal = currentCustomers.SaveToRedis()
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
				w.Write(resultVal.ToByte())

				return
			}
		}
		resultVal = currentCustomerRfidDevices.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		resultVal = currentCustomerRecyDevices.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		resultVal = currentCustomerUltDevices.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		resultVal = currentCustomerUsers.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		resultVal = currentCustomerTags.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		resultVal = currentCustomerConfig.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		resultVal = currentAdminConfig.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		resultVal = currentLocalConfig.SaveToRedis()
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

	var linkCustomer WasteLibrary.CustomerType
	resultVal = linkCustomer.GetByRedisByLink(req.Host)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_NOTFOUND
		w.Write(resultVal.ToByte())

		return
	}

	resultVal = checkAuth(req.Form, linkCustomer.ToIdString())

	if resultVal.Result != WasteLibrary.RESULT_OK {
		w.Write(resultVal.ToByte())

		return
	}

	var currentData WasteLibrary.CustomerType = WasteLibrary.StringToCustomerType(req.FormValue(WasteLibrary.HTTP_DATA))
	var oldCustomer WasteLibrary.CustomerType
	var isCustomerExist = false
	if currentData.CustomerId != 0 {
		isCustomerExist = true
	}

	if currentData.CustomerId != 0 {
		oldCustomer.CustomerId = currentData.CustomerId
		resultVal = oldCustomer.GetByRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_NOTFOUND
			w.Write(resultVal.ToByte())
			return
		}
	}

	resultVal = currentData.SaveToDb()
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
		w.Write(resultVal.ToByte())

		return
	}

	resultVal = currentData.SaveToRedis()
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
		w.Write(resultVal.ToByte())

		return
	}
	resultVal = currentData.SaveToRedisLink()
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

		if resultVal.Result == WasteLibrary.RESULT_OK {
			var currentCustomers WasteLibrary.CustomersType
			resultVal = currentCustomers.GetByRedis()

			if resultVal.Result == WasteLibrary.RESULT_OK {

			} else {
				currentCustomers = WasteLibrary.CustomersType{
					Customers: make(map[string]float64),
				}
			}
			currentCustomers.Customers[currentData.ToIdString()] = currentData.CustomerId
			resultVal = currentCustomers.SaveToRedis()
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
				w.Write(resultVal.ToByte())

				return
			}
		}
		resultVal = currentCustomerRfidDevices.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		resultVal = currentCustomerRecyDevices.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		resultVal = currentCustomerUltDevices.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		resultVal = currentCustomerUsers.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		resultVal = currentCustomerTags.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		resultVal = currentCustomerConfig.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		resultVal = currentAdminConfig.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		resultVal = currentLocalConfig.SaveToRedis()
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

	var linkCustomer WasteLibrary.CustomerType
	resultVal = linkCustomer.GetByRedisByLink(req.Host)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_NOTFOUND
		w.Write(resultVal.ToByte())

		return
	}

	resultVal = checkAuth(req.Form, linkCustomer.ToIdString())

	if resultVal.Result != WasteLibrary.RESULT_OK {
		w.Write(resultVal.ToByte())

		return
	}

	var customers WasteLibrary.CustomersType
	resultVal = customers.GetByRedis()
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMERS_NOTFOUND
		w.Write(resultVal.ToByte())

		return
	}

	var customersList WasteLibrary.CustomersListType
	customersList.New()
	for _, customerId := range customers.Customers {

		if customerId != 0 {
			var currentCustomer WasteLibrary.CustomerType
			currentCustomer.CustomerId = customerId
			resultVal = currentCustomer.GetByRedis()
			if resultVal.Result == WasteLibrary.RESULT_OK {
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

	var linkCustomer WasteLibrary.CustomerType
	resultVal = linkCustomer.GetByRedisByLink(req.Host)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_NOTFOUND
		w.Write(resultVal.ToByte())

		return
	}

	resultVal = checkAuth(req.Form, linkCustomer.ToIdString())

	if resultVal.Result != WasteLibrary.RESULT_OK {
		w.Write(resultVal.ToByte())

		return
	}

	var currentData WasteLibrary.CustomerType = WasteLibrary.StringToCustomerType(req.FormValue(WasteLibrary.HTTP_DATA))

	resultVal = currentData.GetByRedis()
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

	var linkCustomer WasteLibrary.CustomerType
	resultVal = linkCustomer.GetByRedisByLink(req.Host)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_NOTFOUND
		w.Write(resultVal.ToByte())

		return
	}

	resultVal = checkAuth(req.Form, linkCustomer.ToIdString())

	if resultVal.Result != WasteLibrary.RESULT_OK {
		w.Write(resultVal.ToByte())

		return
	}

	var currentHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.StringToHttpClientHeaderType(req.FormValue(WasteLibrary.HTTP_HEADER))
	if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_RFID {
		var currentData WasteLibrary.RfidDeviceType = WasteLibrary.StringToRfidDeviceType(req.FormValue(WasteLibrary.HTTP_DATA))
		var currentOldData WasteLibrary.RfidDeviceType
		currentOldData.DeviceId = currentData.DeviceId
		currentOldData.GetByRedis()
		currentData.DeviceMain.DeviceId = currentData.DeviceId
		resultVal = currentData.DeviceMain.SaveToDb()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		resultVal = currentData.DeviceMain.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		if currentData.DeviceMain.CustomerId != currentOldData.DeviceMain.CustomerId {
			var customerDevices WasteLibrary.CustomerRfidDevicesType
			customerDevices.CustomerId = currentData.DeviceMain.CustomerId
			resultVal = customerDevices.GetByRedis()
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_GET
				w.Write(resultVal.ToByte())

				return
			}

			customerDevices.Devices[currentData.DeviceMain.ToIdString()] = currentData.DeviceMain.DeviceId
			resultVal = customerDevices.SaveToRedis()
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
				w.Write(resultVal.ToByte())

				return
			}

			var oldCustomerDevices WasteLibrary.CustomerRfidDevicesType
			oldCustomerDevices.CustomerId = currentOldData.DeviceMain.CustomerId
			resultVal = oldCustomerDevices.GetByRedis()
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_GET
				w.Write(resultVal.ToByte())

				return
			}

			delete(oldCustomerDevices.Devices, currentData.DeviceMain.ToIdString())
			resultVal = oldCustomerDevices.SaveToRedis()
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
				w.Write(resultVal.ToByte())

				return
			}
		}

		resultVal = WasteLibrary.PublishRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_CHANNEL+currentData.DeviceMain.ToCustomerIdString(), WasteLibrary.DATATYPE_RFID_MAIN_DEVICE, currentData.DeviceMain.ToString())

	} else if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_ULT {
		var currentData WasteLibrary.UltDeviceType = WasteLibrary.StringToUltDeviceType(req.FormValue(WasteLibrary.HTTP_DATA))
		var currentOldData WasteLibrary.UltDeviceType
		currentOldData.DeviceId = currentData.DeviceId
		currentOldData.GetByRedis()
		currentData.DeviceMain.DeviceId = currentData.DeviceId
		resultVal = currentData.DeviceMain.SaveToDb()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		resultVal = currentData.DeviceMain.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		if currentData.DeviceMain.CustomerId != currentOldData.DeviceMain.CustomerId {
			var customerDevices WasteLibrary.CustomerUltDevicesType
			customerDevices.CustomerId = currentData.DeviceMain.CustomerId
			resultVal = customerDevices.GetByRedis()
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_GET
				w.Write(resultVal.ToByte())

				return
			}

			customerDevices.Devices[currentData.DeviceMain.ToIdString()] = currentData.DeviceMain.DeviceId
			resultVal = customerDevices.SaveToRedis()
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
				w.Write(resultVal.ToByte())

				return
			}
			var oldCustomerDevices WasteLibrary.CustomerUltDevicesType
			oldCustomerDevices.CustomerId = currentOldData.DeviceMain.CustomerId
			resultVal = oldCustomerDevices.GetByRedis()
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_GET
				w.Write(resultVal.ToByte())

				return
			}

			delete(oldCustomerDevices.Devices, currentData.DeviceMain.ToIdString())
			resultVal = oldCustomerDevices.SaveToRedis()
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
				w.Write(resultVal.ToByte())

				return
			}
		}
		resultVal = WasteLibrary.PublishRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_CHANNEL+currentData.DeviceMain.ToCustomerIdString(), WasteLibrary.DATATYPE_ULT_MAIN_DEVICE, currentData.DeviceMain.ToString())

	} else if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_RECY {
		var currentData WasteLibrary.RecyDeviceType = WasteLibrary.StringToRecyDeviceType(req.FormValue(WasteLibrary.HTTP_DATA))
		var currentOldData WasteLibrary.RecyDeviceType
		currentOldData.DeviceId = currentData.DeviceId
		currentOldData.GetByRedis()
		currentData.DeviceMain.DeviceId = currentData.DeviceId
		resultVal = currentData.DeviceMain.SaveToDb()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		resultVal = currentData.DeviceMain.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		if currentData.DeviceMain.CustomerId != currentOldData.DeviceMain.CustomerId {
			var customerDevices WasteLibrary.CustomerRecyDevicesType
			customerDevices.CustomerId = currentData.DeviceMain.CustomerId
			resultVal = customerDevices.GetByRedis()
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_GET
				w.Write(resultVal.ToByte())

				return
			}

			customerDevices.Devices[currentData.DeviceMain.ToIdString()] = currentData.DeviceMain.DeviceId
			resultVal = customerDevices.SaveToRedis()
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
				w.Write(resultVal.ToByte())

				return
			}

			var oldCustomerDevices WasteLibrary.CustomerRecyDevicesType
			oldCustomerDevices.CustomerId = currentOldData.DeviceMain.CustomerId
			resultVal = oldCustomerDevices.GetByRedis()
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_GET
				w.Write(resultVal.ToByte())

				return
			}

			delete(oldCustomerDevices.Devices, currentData.DeviceMain.ToIdString())
			resultVal = oldCustomerDevices.SaveToRedis()
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

	var linkCustomer WasteLibrary.CustomerType
	resultVal = linkCustomer.GetByRedisByLink(req.Host)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_NOTFOUND
		w.Write(resultVal.ToByte())

		return
	}

	resultVal = checkAuth(req.Form, linkCustomer.ToIdString())

	if resultVal.Result != WasteLibrary.RESULT_OK {
		w.Write(resultVal.ToByte())

		return
	}
	var currentHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.StringToHttpClientHeaderType(req.FormValue(WasteLibrary.HTTP_HEADER))
	if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_RFID {
		var customerDevicesList WasteLibrary.CustomerRfidDevicesListType = WasteLibrary.StringToCustomerRfidDevicesListType(req.FormValue(WasteLibrary.HTTP_DATA))

		var customerDevices WasteLibrary.CustomerRfidDevicesType
		customerDevices.CustomerId = customerDevicesList.CustomerId
		resultVal = customerDevices.GetByRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMERS_DEVICES_NOTFOUND
			w.Write(resultVal.ToByte())

			return
		}

		customerDevicesList.Devices = make(map[string]WasteLibrary.RfidDeviceType)
		for _, deviceId := range customerDevices.Devices {

			if deviceId != 0 {
				var currentDevice WasteLibrary.RfidDeviceType
				currentDevice.New()
				currentDevice.DeviceId = deviceId
				resultVal = currentDevice.GetByRedis()
				if resultVal.Result == WasteLibrary.RESULT_OK {
					customerDevicesList.Devices[currentDevice.ToIdString()] = currentDevice
				}
			}
		}

		resultVal.Result = WasteLibrary.RESULT_OK
		resultVal.Retval = customerDevicesList.ToString()
	} else if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_ULT {
		var customerDevicesList WasteLibrary.CustomerUltDevicesListType = WasteLibrary.StringToCustomerUltDevicesListType(req.FormValue(WasteLibrary.HTTP_DATA))
		var customerDevices WasteLibrary.CustomerUltDevicesType
		customerDevices.CustomerId = customerDevicesList.CustomerId
		resultVal = customerDevices.GetByRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMERS_DEVICES_NOTFOUND
			w.Write(resultVal.ToByte())

			return
		}

		customerDevicesList.Devices = make(map[string]WasteLibrary.UltDeviceType)
		for _, deviceId := range customerDevices.Devices {

			if deviceId != 0 {
				var currentDevice WasteLibrary.UltDeviceType
				currentDevice.New()
				currentDevice.DeviceId = deviceId
				resultVal = currentDevice.GetByRedis()
				if resultVal.Result == WasteLibrary.RESULT_OK {
					customerDevicesList.Devices[currentDevice.ToIdString()] = currentDevice
				}
			}
		}

		resultVal.Result = WasteLibrary.RESULT_OK
		resultVal.Retval = customerDevicesList.ToString()
	} else if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_RECY {
		var customerDevicesList WasteLibrary.CustomerRecyDevicesListType = WasteLibrary.StringToCustomerRecyDevicesListType(req.FormValue(WasteLibrary.HTTP_DATA))
		var customerDevices WasteLibrary.CustomerRecyDevicesType
		customerDevices.CustomerId = customerDevicesList.CustomerId
		resultVal = customerDevices.GetByRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMERS_DEVICES_NOTFOUND
			w.Write(resultVal.ToByte())

			return
		}

		customerDevicesList.Devices = make(map[string]WasteLibrary.RecyDeviceType)
		for _, deviceId := range customerDevices.Devices {

			if deviceId != 0 {
				var currentDevice WasteLibrary.RecyDeviceType
				currentDevice.New()
				currentDevice.DeviceId = deviceId
				resultVal = currentDevice.GetByRedis()
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

	var linkCustomer WasteLibrary.CustomerType
	resultVal = linkCustomer.GetByRedisByLink(req.Host)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_NOTFOUND
		w.Write(resultVal.ToByte())

		return
	}

	resultVal = checkAuth(req.Form, linkCustomer.ToIdString())

	if resultVal.Result != WasteLibrary.RESULT_OK {
		w.Write(resultVal.ToByte())

		return
	}

	var currentHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.StringToHttpClientHeaderType(req.FormValue(WasteLibrary.HTTP_HEADER))
	if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_RFID {
		var currentDevice WasteLibrary.RfidDeviceType = WasteLibrary.StringToRfidDeviceType(req.FormValue(WasteLibrary.HTTP_DATA))
		resultVal = currentDevice.GetByRedis()
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
		resultVal = currentDevice.GetByRedis()
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
		resultVal = currentDevice.GetByRedis()
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
	var linkCustomer WasteLibrary.CustomerType
	resultVal = linkCustomer.GetByRedisByLink(req.Host)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_NOTFOUND
		w.Write(resultVal.ToByte())

		return
	}

	resultVal = checkAuth(req.Form, linkCustomer.ToIdString())

	if resultVal.Result != WasteLibrary.RESULT_OK {
		w.Write(resultVal.ToByte())

		return
	}

	var currentHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.StringToHttpClientHeaderType(req.FormValue(WasteLibrary.HTTP_HEADER))
	if currentHttpHeader.DataType == WasteLibrary.DATATYPE_CUSTOMERCONFIG {
		var currentConfig WasteLibrary.CustomerConfigType = WasteLibrary.StringToCustomerConfigType(req.FormValue(WasteLibrary.HTTP_DATA))
		resultVal = currentConfig.GetByRedis()
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
	var linkCustomer WasteLibrary.CustomerType
	resultVal = linkCustomer.GetByRedisByLink(req.Host)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_NOTFOUND
		w.Write(resultVal.ToByte())

		return
	}

	resultVal = checkAuth(req.Form, linkCustomer.ToIdString())

	if resultVal.Result != WasteLibrary.RESULT_OK {
		w.Write(resultVal.ToByte())

		return
	}

	var currentHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.StringToHttpClientHeaderType(req.FormValue(WasteLibrary.HTTP_HEADER))
	if currentHttpHeader.DataType == WasteLibrary.DATATYPE_CUSTOMERCONFIG {
		var currentData WasteLibrary.CustomerConfigType = WasteLibrary.StringToCustomerConfigType(req.FormValue(WasteLibrary.HTTP_DATA))
		resultVal = currentData.SaveToRedis()
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
		var currentCustomer WasteLibrary.CustomerType
		currentCustomer.CustomerId = WasteLibrary.StringIdToFloat64(customerId)
		resultVal = currentCustomer.GetByRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_NOTFOUND
			return resultVal
		}
		if currentCustomer.CustomerName != "AFATEK" {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_INVALID
			return resultVal
		}
	}
	return resultVal

}
