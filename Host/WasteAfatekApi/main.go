package main

import (
	"net/http"

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
	/*http.HandleFunc("/setCustomer", setCustomer)
	http.HandleFunc("/getCustomer", getCustomer)
	http.HandleFunc("/getCustomers", getCustomers)
	http.HandleFunc("/setDevice", setDevice)
	http.HandleFunc("/getDevice", getDevice)
	http.HandleFunc("/getDevices", getDevices)*/
	http.ListenAndServe(":80", nil)
}

/*
func setCustomer(w http.ResponseWriter, req *http.Request) {

	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	if err := req.ParseForm(); err != nil {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_HTTP_PARSE
		w.Write(resultVal.ToByte())
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
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_NOTFOUND
		w.Write(resultVal.ToByte())
		return
	}
	var customerId string = resultVal.Retval.(string)

	if customerId != "1" {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = ""
		w.Write(resultVal.ToByte())
		return
	}

	var currentHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.StringToHttpClientHeaderType(req.FormValue(WasteLibrary.HTTP_HEADER))
	var currentData WasteLibrary.CustomerType = WasteLibrary.StringToCustomerType(req.FormValue(WasteLibrary.HTTP_DATA))
	data := url.Values{
		WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
		WasteLibrary.HTTP_DATA:   {currentData.ToString()},
	}
	resultVal = WasteLibrary.SaveStaticDbMainForStoreApi(data)
	if resultVal.Result == WasteLibrary.RESULT_OK {

		var isCustomerExist = false
		if currentData.CustomerId != 0 {
			isCustomerExist = true
		}
		currentData.CustomerId = WasteLibrary.StringIdToFloat64(resultVal.Retval.(string))
		data := url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentData.ToString()},
		}
		var currentCustomer WasteLibrary.CustomerType = WasteLibrary.StringToCustomerType(WasteLibrary.GetStaticDbMainForStoreApi(data).Retval.(string))

		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_CUSTOMERS, currentCustomer.ToIdString(), currentCustomer.ToString())
		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_LINK, currentCustomer.CustomerLink, currentCustomer.ToIdString())
		if !isCustomerExist {
			var currentCustomerDevices WasteLibrary.CustomerDevicesType = WasteLibrary.CustomerDevicesType{
				CustomerId: currentCustomer.CustomerId,
				Devices:    make(map[string]float64),
			}
			currentCustomerDevices.Devices["0"] = 0

			var currentCustomerTags WasteLibrary.CustomerTagsType
			currentCustomerTags.CustomerId = currentCustomer.CustomerId
			currentCustomerTags.Tags = make(map[string]WasteLibrary.TagType)
			currentCustomerTags.Tags["0"] = WasteLibrary.TagType{TagID: 0}

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
			}
			resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_DEVICES, currentCustomerDevices.ToIdString(), currentCustomerDevices.ToString())
			resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_TAGS, currentCustomerTags.ToIdString(), currentCustomerTags.ToString())
			resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_CUSTOMERCONFIG, currentCustomerConfig.ToIdString(), currentCustomerConfig.ToString())
			resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_ADMINCONFIG, currentAdminConfig.ToIdString(), currentAdminConfig.ToString())
			resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_LOCALCONFIG, currentLocalConfig.ToIdString(), currentLocalConfig.ToString())
		}

	}
	w.Write(resultVal.ToByte())
}
*/
/*
func setDevice(w http.ResponseWriter, req *http.Request) {

	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	if err := req.ParseForm(); err != nil {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_HTTP_PARSE
		w.Write(resultVal.ToByte())
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
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_NOTFOUND
		w.Write(resultVal.ToByte())
		return
	}
	var customerId string = resultVal.Retval.(string)

	if customerId != "1" {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = ""
		w.Write(resultVal.ToByte())
		return
	}

	var currentHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.StringToHttpClientHeaderType(req.FormValue(WasteLibrary.HTTP_HEADER))
	var currentData WasteLibrary.DeviceType = WasteLibrary.StringToDeviceType(req.FormValue(WasteLibrary.HTTP_DATA))
	WasteLibrary.LogStr("Data : " + currentData.ToString())
	data := url.Values{
		WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
		WasteLibrary.HTTP_DATA:   {currentData.ToString()},
	}
	resultVal = WasteLibrary.SaveStaticDbMainForStoreApi(data)
	if resultVal.Result == WasteLibrary.RESULT_OK {

		var currentData WasteLibrary.DeviceType
		currentData.DeviceId = WasteLibrary.StringIdToFloat64(resultVal.Retval.(string))
		data := url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentData.ToString()},
		}
		var currentDevice WasteLibrary.DeviceType = WasteLibrary.StringToDeviceType(WasteLibrary.GetStaticDbMainForStoreApi(data).Retval.(string))

		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_DEVICES, currentDevice.ToIdString(), currentDevice.ToString())
		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_SERIAL_DEVICE, currentDevice.SerialNumber, currentDevice.ToIdString())
		resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_DEVICES, currentDevice.ToCustomerIdString())
		if resultVal.Result == WasteLibrary.RESULT_OK {
			WasteLibrary.LogStr("Customer-Devices : " + resultVal.Retval.(string))
			var currentCustomerDevices WasteLibrary.CustomerDevicesType = WasteLibrary.StringToCustomerDevicesType(resultVal.Retval.(string))
			currentCustomerDevices.Devices[currentDevice.ToIdString()] = currentDevice.DeviceId
			WasteLibrary.LogStr("New Customer-Devices : " + currentCustomerDevices.ToString())
			resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_DEVICES, currentCustomerDevices.ToIdString(), currentCustomerDevices.ToString())
		}
	}
	w.Write(resultVal.ToByte())
}
*/
