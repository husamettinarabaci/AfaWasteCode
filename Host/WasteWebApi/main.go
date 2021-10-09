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
	http.HandleFunc("/getCustomer", getCustomer)
	http.HandleFunc("/getDevice", getDevice)
	http.HandleFunc("/getDevices", getDevices)
	http.HandleFunc("/getConfig", getConfig)
	http.HandleFunc("/getTags", getTags)
	http.HandleFunc("/getTag", getTag)
	http.ListenAndServe(":80", nil)
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

	resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMERS, customerId)

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

	var currentData WasteLibrary.DeviceType = WasteLibrary.StringToDeviceType(req.FormValue(WasteLibrary.HTTP_DATA))
	resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_DEVICES, currentData.ToIdString())
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_DEVICE_NOTFOUND
		w.Write(resultVal.ToByte())
		return
	}
	currentData = WasteLibrary.StringToDeviceType(resultVal.Retval.(string))
	if currentData.ToCustomerIdString() == customerId {
		resultVal.Result = WasteLibrary.RESULT_OK
		resultVal.Retval = currentData.ToString()
		w.Write(resultVal.ToByte())
	} else {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_DEVICE_NOTFOUND
		w.Write(resultVal.ToByte())
	}
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

	resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_DEVICES, customerId)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_NOTFOUND
		w.Write(resultVal.ToByte())
		return
	}

	var customerDevices WasteLibrary.CustomerDevicesType = WasteLibrary.StringToCustomerDevicesType(resultVal.Retval.(string))
	var customerDevicesList WasteLibrary.CustomerDevicesListType = WasteLibrary.CustomerDevicesListType{
		CustomerId: WasteLibrary.StringIdToFloat64(customerId),
		Devices:    make(map[string]WasteLibrary.DeviceType),
	}
	for _, deviceId := range customerDevices.Devices {

		if deviceId != 0 {
			resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_DEVICES, WasteLibrary.Float64IdToString(deviceId))
			if resultVal.Result == WasteLibrary.RESULT_OK {
				var currentDevice WasteLibrary.DeviceType = WasteLibrary.StringToDeviceType(resultVal.Retval.(string))
				customerDevicesList.Devices[currentDevice.ToIdString()] = currentDevice
			}

		}
	}
	resultVal.Result = WasteLibrary.RESULT_OK
	resultVal.Retval = customerDevicesList.ToString()

	w.Write(resultVal.ToByte())
}

func getConfig(w http.ResponseWriter, req *http.Request) {
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

	var currentHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.StringToHttpClientHeaderType(req.FormValue(WasteLibrary.HTTP_HEADER))
	if currentHttpHeader.OpType == WasteLibrary.OPTYPE_CUSTOMERCONFIG {
		resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_CUSTOMERCONFIG, customerId)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_CUSTOMERCONFIG_NOTFOUND
			w.Write(resultVal.ToByte())
			return
		}
	} else if currentHttpHeader.OpType == WasteLibrary.OPTYPE_ADMINCONFIG {
		resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_ADMINCONFIG, customerId)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_ADMINCONFIG_NOTFOUND
			w.Write(resultVal.ToByte())
			return
		}
	} else if currentHttpHeader.OpType == WasteLibrary.OPTYPE_LOCALCONFIG {
		resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_LOCALCONFIG, customerId)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_LOCALCONFIG_NOTFOUND
			w.Write(resultVal.ToByte())
			return
		}
	} else {
		resultVal.Result = WasteLibrary.RESULT_FAIL
	}

	w.Write(resultVal.ToByte())
}

func getTags(w http.ResponseWriter, req *http.Request) {
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

	resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_TAGS, customerId)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_TAGS_NOTFOUND
		w.Write(resultVal.ToByte())
		return
	}
	w.Write(resultVal.ToByte())
}

func getTag(w http.ResponseWriter, req *http.Request) {
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

	var currentData WasteLibrary.TagType = WasteLibrary.StringToTagType(req.FormValue(WasteLibrary.HTTP_DATA))

	resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_TAGS, customerId)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_NOTFOUND
		w.Write(resultVal.ToByte())
		return
	}

	var currentCustomerTags WasteLibrary.CustomerTagsType = WasteLibrary.StringToCustomerTagsType(resultVal.Retval.(string))
	currentData = currentCustomerTags.Tags[currentData.ToIdString()]
	resultVal.Retval = currentData.ToString()

	w.Write(resultVal.ToByte())
}
