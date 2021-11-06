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
	http.HandleFunc("/data", data)
	http.ListenAndServe(":80", nil)
}

func data(w http.ResponseWriter, req *http.Request) {

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

	var currentHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.StringToHttpClientHeaderType(req.FormValue(WasteLibrary.HTTP_HEADER))
	if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_RFID {
		resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_SERIAL_RFID_DEVICE, currentHttpHeader.DeviceNo)
		if resultVal.Result == WasteLibrary.RESULT_FAIL {
			resultVal = createDevice(currentHttpHeader)
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_HTTP_POST
				w.Write(resultVal.ToByte())

				return
			}
			resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_SERIAL_RFID_DEVICE, currentHttpHeader.DeviceNo)
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_DEVICE_NOTFOUND
				w.Write(resultVal.ToByte())

				return
			}
		}
		var deviceIdStr = resultVal.Retval.(string)
		var currentDevice WasteLibrary.RfidDeviceType
		currentDevice.New()
		currentDevice.DeviceId = WasteLibrary.StringIdToFloat64(deviceIdStr)
		currentDevice.GetAll()
		currentHttpHeader.CustomerId = currentDevice.DeviceMain.CustomerId
		currentHttpHeader.DeviceId = currentDevice.DeviceId
	} else if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_ULT {
		resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_SERIAL_ULT_DEVICE, currentHttpHeader.DeviceNo)
		if resultVal.Result == WasteLibrary.RESULT_FAIL {
			resultVal = createDevice(currentHttpHeader)
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_HTTP_POST
				w.Write(resultVal.ToByte())

				return
			}
			resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_SERIAL_ULT_DEVICE, currentHttpHeader.DeviceNo)
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_DEVICE_NOTFOUND
				w.Write(resultVal.ToByte())

				return
			}
		}
		var deviceIdStr = resultVal.Retval.(string)
		var currentDevice WasteLibrary.UltDeviceType
		currentDevice.New()
		currentDevice.DeviceId = WasteLibrary.StringIdToFloat64(deviceIdStr)
		currentDevice.GetAll()
		currentHttpHeader.CustomerId = currentDevice.DeviceMain.CustomerId
		currentHttpHeader.DeviceId = currentDevice.DeviceId
	} else if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_RECY {
		resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_SERIAL_RECY_DEVICE, currentHttpHeader.DeviceNo)
		if resultVal.Result == WasteLibrary.RESULT_FAIL {
			resultVal = createDevice(currentHttpHeader)
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_HTTP_POST
				w.Write(resultVal.ToByte())

				return
			}
			resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_SERIAL_RECY_DEVICE, currentHttpHeader.DeviceNo)
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_DEVICE_NOTFOUND
				w.Write(resultVal.ToByte())

				return
			}
		}
		var deviceIdStr = resultVal.Retval.(string)
		var currentDevice WasteLibrary.RecyDeviceType
		currentDevice.New()
		currentDevice.DeviceId = WasteLibrary.StringIdToFloat64(deviceIdStr)
		currentDevice.GetAll()
		currentHttpHeader.CustomerId = currentDevice.DeviceMain.CustomerId
		currentHttpHeader.DeviceId = currentDevice.DeviceId
	} else {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_DEVICE_NOTFOUND
		w.Write(resultVal.ToByte())

		return
	}

	var serviceClusterIp string = ""
	if currentHttpHeader.AppType == WasteLibrary.APPTYPE_RFID {

		if currentHttpHeader.ReaderType == WasteLibrary.READERTYPE_RF {
			serviceClusterIp = "waste-rfreader-cluster-ip"

		} else if currentHttpHeader.ReaderType == WasteLibrary.READERTYPE_GPS {
			serviceClusterIp = "waste-gpsreader-cluster-ip"

		} else if currentHttpHeader.ReaderType == WasteLibrary.READERTYPE_STATUS {
			serviceClusterIp = "waste-statusreader-cluster-ip"

		} else if currentHttpHeader.ReaderType == WasteLibrary.READERTYPE_THERM {
			serviceClusterIp = "waste-thermreader-cluster-ip"

		} else if currentHttpHeader.ReaderType == WasteLibrary.READERTYPE_CAM {
			serviceClusterIp = "waste-camreader-cluster-ip"
		} else {
			resultVal.Result = WasteLibrary.RESULT_FAIL
		}
	} else if currentHttpHeader.AppType == WasteLibrary.APPTYPE_ULT {
		if currentHttpHeader.ReaderType == WasteLibrary.READERTYPE_ULT {

			//TO DO
			//part reader type by values

		} else {
			resultVal.Result = WasteLibrary.RESULT_FAIL
		}
	} else if currentHttpHeader.AppType == WasteLibrary.APPTYPE_RECY {
		resultVal.Result = WasteLibrary.RESULT_OK
	} else {
		resultVal.Result = WasteLibrary.RESULT_FAIL
	}

	data := url.Values{
		WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
		WasteLibrary.HTTP_DATA:   {req.FormValue(WasteLibrary.HTTP_DATA)},
	}

	resultVal = WasteLibrary.SaveBulkDbMainForStoreApi(data)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
		w.Write(resultVal.ToByte())

		return
	}
	if serviceClusterIp != "" {
		resultVal = WasteLibrary.HttpPostReq("http://"+serviceClusterIp+"/reader", data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_HTTP_POST
			w.Write(resultVal.ToByte())

			return
		}
	}
	w.Write(resultVal.ToByte())

}

func createDevice(currentHttpHeader WasteLibrary.HttpClientHeaderType) WasteLibrary.ResultType {
	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	var createHttpHeader WasteLibrary.HttpClientHeaderType
	createHttpHeader.New()
	createHttpHeader.AppType = WasteLibrary.APPTYPE_LISTENER
	createHttpHeader.DeviceType = currentHttpHeader.DeviceType
	createHttpHeader.DeviceNo = currentHttpHeader.DeviceNo
	data := url.Values{
		WasteLibrary.HTTP_HEADER: {createHttpHeader.ToString()},
	}

	resultVal = WasteLibrary.HttpPostReq("http://waste-enhcapi-cluster-ip/createDevice", data)
	return resultVal
}
