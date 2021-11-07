package main

import (
	"fmt"
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

	resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMERS, customerId)

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
	var currentHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.StringToHttpClientHeaderType(req.FormValue(WasteLibrary.HTTP_HEADER))
	if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_RFID {
		var currentData WasteLibrary.RfidDeviceType = WasteLibrary.StringToRfidDeviceType(req.FormValue(WasteLibrary.HTTP_DATA))
		currentData.GetAll()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DEVICE_NOTFOUND
			w.Write(resultVal.ToByte())

			return
		}
		if currentData.DeviceMain.ToCustomerIdString() == customerId {
			resultVal.Result = WasteLibrary.RESULT_OK
			resultVal.Retval = currentData.ToString()
			w.Write(resultVal.ToByte())

		} else {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DEVICE_NOTFOUND
			w.Write(resultVal.ToByte())

		}
	} else if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_ULT {
		var currentData WasteLibrary.UltDeviceType = WasteLibrary.StringToUltDeviceType(req.FormValue(WasteLibrary.HTTP_DATA))
		currentData.GetAll()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DEVICE_NOTFOUND
			w.Write(resultVal.ToByte())

			return
		}
		if currentData.DeviceMain.ToCustomerIdString() == customerId {
			resultVal.Result = WasteLibrary.RESULT_OK
			resultVal.Retval = currentData.ToString()
			w.Write(resultVal.ToByte())

		} else {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DEVICE_NOTFOUND
			w.Write(resultVal.ToByte())

		}
	} else if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_RECY {
		var currentData WasteLibrary.RecyDeviceType = WasteLibrary.StringToRecyDeviceType(req.FormValue(WasteLibrary.HTTP_DATA))
		currentData.GetAll()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DEVICE_NOTFOUND
			w.Write(resultVal.ToByte())

			return
		}
		if currentData.DeviceMain.ToCustomerIdString() == customerId {
			resultVal.Result = WasteLibrary.RESULT_OK
			resultVal.Retval = currentData.ToString()
			w.Write(resultVal.ToByte())

		} else {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DEVICE_NOTFOUND
			w.Write(resultVal.ToByte())

		}
	} else {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_DEVICE_NOTFOUND
		w.Write(resultVal.ToByte())

	}

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

	fmt.Println(req.Form)

	resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_LINK, req.Host)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_NOTFOUND
		w.Write(resultVal.ToByte())

		return
	}
	var customerId string = resultVal.Retval.(string)
	var currentHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.StringToHttpClientHeaderType(req.FormValue(WasteLibrary.HTTP_HEADER))

	if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_RFID {
		resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_RFID_DEVICES, customerId)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_NOTFOUND
			w.Write(resultVal.ToByte())

			return
		}

		var customerDevices WasteLibrary.CustomerRfidDevicesType = WasteLibrary.StringToCustomerRfidDevicesType(resultVal.Retval.(string))
		var customerDevicesList WasteLibrary.CustomerRfidDevicesListType
		customerDevicesList.New()
		customerDevicesList.CustomerId = WasteLibrary.StringIdToFloat64(customerId)
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
		resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_ULT_DEVICES, customerId)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_NOTFOUND
			w.Write(resultVal.ToByte())

			return
		}

		var customerDevices WasteLibrary.CustomerUltDevicesType = WasteLibrary.StringToCustomerUltDevicesType(resultVal.Retval.(string))
		var customerDevicesList WasteLibrary.CustomerUltDevicesListType
		customerDevicesList.New()
		customerDevicesList.CustomerId = WasteLibrary.StringIdToFloat64(customerId)
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
		resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_RECY_DEVICES, customerId)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_NOTFOUND
			w.Write(resultVal.ToByte())

			return
		}

		var customerDevices WasteLibrary.CustomerRecyDevicesType = WasteLibrary.StringToCustomerRecyDevicesType(resultVal.Retval.(string))
		var customerDevicesList WasteLibrary.CustomerRecyDevicesListType
		customerDevicesList.New()
		customerDevicesList.CustomerId = WasteLibrary.StringIdToFloat64(customerId)

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
		resultVal.Retval = WasteLibrary.RESULT_ERROR_DEVICE_NOTFOUND
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

	var currentHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.StringToHttpClientHeaderType(req.FormValue(WasteLibrary.HTTP_HEADER))
	if currentHttpHeader.DataType == WasteLibrary.DATATYPE_CUSTOMERCONFIG {
		resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_CUSTOMERCONFIG, customerId)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_CUSTOMERCONFIG_NOTFOUND
			w.Write(resultVal.ToByte())

			return
		}
	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_ADMINCONFIG {
		resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_ADMINCONFIG, customerId)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_ADMINCONFIG_NOTFOUND
			w.Write(resultVal.ToByte())

			return
		}
	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_LOCALCONFIG {
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

	resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_TAGS, customerId)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_NOTFOUND
		w.Write(resultVal.ToByte())

		return
	}

	var customerTags WasteLibrary.CustomerTagsType = WasteLibrary.StringToCustomerTagsType(resultVal.Retval.(string))
	var customerTagsList WasteLibrary.CustomerTagsListType
	customerTagsList.New()
	customerTagsList.CustomerId = WasteLibrary.StringIdToFloat64(customerId)
	for _, tagId := range customerTags.Tags {

		if tagId != 0 {

			var currentTag WasteLibrary.TagType
			currentTag.New()
			currentTag.TagId = tagId
			resultVal = currentTag.GetAll()
			if resultVal.Result == WasteLibrary.RESULT_OK {
				customerTagsList.Tags[currentTag.ToIdString()] = currentTag
			}

		}
	}
	resultVal.Result = WasteLibrary.RESULT_OK
	resultVal.Retval = customerTagsList.ToString()
	w.Write(resultVal.ToByte())

}

func getTag(w http.ResponseWriter, req *http.Request) {

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

	var currentData WasteLibrary.TagType = WasteLibrary.StringToTagType(req.FormValue(WasteLibrary.HTTP_DATA))

	resultVal = currentData.GetAll()
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_TAG_NOTFOUND
		w.Write(resultVal.ToByte())

		return
	}
	currentData = WasteLibrary.StringToTagType(resultVal.Retval.(string))
	currentData.GetAll()
	if currentData.TagMain.ToCustomerIdString() == customerId {
		resultVal.Result = WasteLibrary.RESULT_OK
		resultVal.Retval = currentData.ToString()
		w.Write(resultVal.ToByte())

	} else {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_TAG_NOTFOUND
		w.Write(resultVal.ToByte())

	}
}
