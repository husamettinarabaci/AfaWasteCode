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
	http.HandleFunc("/openLog", WasteLibrary.OpenLogHandler)
	http.HandleFunc("/closeLog", WasteLibrary.CloseLogHandler)
	http.HandleFunc("/data", data)
	http.ListenAndServe(":80", nil)
}

func data(w http.ResponseWriter, req *http.Request) {

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

	var currentHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.StringToHttpClientHeaderType(req.FormValue(WasteLibrary.HTTP_HEADER))
	if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_RFID {
		resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_SERIAL_RFID_DEVICE, currentHttpHeader.DeviceNo)
		if resultVal.Result == WasteLibrary.RESULT_FAIL {
			resultVal = createDevice(currentHttpHeader, req.FormValue(WasteLibrary.HTTP_DATA))
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
			resultVal = createDevice(currentHttpHeader, req.FormValue(WasteLibrary.HTTP_DATA))
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
			resultVal = createDevice(currentHttpHeader, req.FormValue(WasteLibrary.HTTP_DATA))
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
			resultVal = sendReader(serviceClusterIp, currentHttpHeader.ToString(), req.FormValue(WasteLibrary.HTTP_DATA))
		} else if currentHttpHeader.ReaderType == WasteLibrary.READERTYPE_GPS {
			serviceClusterIp = "waste-gpsreader-cluster-ip"
			resultVal = sendReader(serviceClusterIp, currentHttpHeader.ToString(), req.FormValue(WasteLibrary.HTTP_DATA))
		} else if currentHttpHeader.ReaderType == WasteLibrary.READERTYPE_STATUS {
			serviceClusterIp = "waste-statusreader-cluster-ip"
			resultVal = sendReader(serviceClusterIp, currentHttpHeader.ToString(), req.FormValue(WasteLibrary.HTTP_DATA))
		} else if currentHttpHeader.ReaderType == WasteLibrary.READERTYPE_THERM {
			serviceClusterIp = "waste-thermreader-cluster-ip"
			resultVal = sendReader(serviceClusterIp, currentHttpHeader.ToString(), req.FormValue(WasteLibrary.HTTP_DATA))
		} else if currentHttpHeader.ReaderType == WasteLibrary.READERTYPE_CAM {
			serviceClusterIp = "waste-camreader-cluster-ip"
			resultVal = sendReader(serviceClusterIp, currentHttpHeader.ToString(), req.FormValue(WasteLibrary.HTTP_DATA))
		} else {
			resultVal.Result = WasteLibrary.RESULT_FAIL
		}
	} else if currentHttpHeader.AppType == WasteLibrary.APPTYPE_ULT {
		if currentHttpHeader.ReaderType == WasteLibrary.READERTYPE_ULT {

			//serviceClusterIp = "waste-gpsreader-cluster-ip"
			//resultVal = sendReader(serviceClusterIp, currentHttpHeader.ToString(), req.FormValue(WasteLibrary.HTTP_DATA))

			serviceClusterIp = "waste-statusreader-cluster-ip"
			resultVal = sendReader(serviceClusterIp, currentHttpHeader.ToString(), req.FormValue(WasteLibrary.HTTP_DATA))

			serviceClusterIp = "waste-thermreader-cluster-ip"
			resultVal = sendReader(serviceClusterIp, currentHttpHeader.ToString(), req.FormValue(WasteLibrary.HTTP_DATA))

			serviceClusterIp = "waste-batteryreader-cluster-ip"
			resultVal = sendReader(serviceClusterIp, currentHttpHeader.ToString(), req.FormValue(WasteLibrary.HTTP_DATA))

			serviceClusterIp = "waste-sensreader-cluster-ip"
			resultVal = sendReader(serviceClusterIp, currentHttpHeader.ToString(), req.FormValue(WasteLibrary.HTTP_DATA))

		} else {
			resultVal.Result = WasteLibrary.RESULT_FAIL
		}
	} else if currentHttpHeader.AppType == WasteLibrary.APPTYPE_RECY {
		resultVal.Result = WasteLibrary.RESULT_OK
	} else {
		resultVal.Result = WasteLibrary.RESULT_FAIL
	}
	w.Write(resultVal.ToByte())

}

func sendReader(serviceClusterIp string, httpHeader string, httpData string) WasteLibrary.ResultType {
	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	data := url.Values{
		WasteLibrary.HTTP_HEADER: {httpHeader},
		WasteLibrary.HTTP_DATA:   {httpData},
	}

	resultVal = WasteLibrary.SaveBulkDbMainForStoreApi(data)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE

		return resultVal
	}
	if serviceClusterIp != "" {
		resultVal = WasteLibrary.HttpPostReq("http://"+serviceClusterIp+"/reader", data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_HTTP_POST

			return resultVal
		}
	}
	return resultVal
}

func createDevice(currentHttpHeader WasteLibrary.HttpClientHeaderType, currentData string) WasteLibrary.ResultType {
	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	var createHttpHeader WasteLibrary.HttpClientHeaderType
	createHttpHeader.New()
	createHttpHeader.AppType = WasteLibrary.APPTYPE_LISTENER
	createHttpHeader.DeviceType = currentHttpHeader.DeviceType
	createHttpHeader.DeviceNo = currentHttpHeader.DeviceNo
	data := url.Values{
		WasteLibrary.HTTP_HEADER: {createHttpHeader.ToString()},
		WasteLibrary.HTTP_DATA:   {currentData},
	}

	resultVal = WasteLibrary.HttpPostReq("http://waste-enhcapi-cluster-ip/createDevice", data)
	return resultVal
}
