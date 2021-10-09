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
	resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_SERIAL_DEVICE, currentHttpHeader.DeviceNo)
	if resultVal.Result == WasteLibrary.RESULT_FAIL {
		//TO DO
		// insert new device
		var createHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.HttpClientHeaderType{
			AppType:      WasteLibrary.APPTYPE_ADMIN,
			DeviceNo:     "",
			OpType:       WasteLibrary.OPTYPE_DEVICE,
			Time:         WasteLibrary.GetTime(),
			Repeat:       WasteLibrary.STATU_PASSIVE,
			DeviceId:     0,
			CustomerId:   0,
			BaseDataType: WasteLibrary.BASETYPE_DEVICE,
		}

		var createDevice WasteLibrary.DeviceType = WasteLibrary.DeviceType{
			DeviceId:     0,
			CustomerId:   -1,
			SerialNumber: currentHttpHeader.DeviceNo,
		}
		WasteLibrary.LogStr(createDevice.ToString())
		data := url.Values{
			WasteLibrary.HTTP_HEADER: {createHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {createDevice.ToString()},
		}

		resultVal = WasteLibrary.HttpPostReq("http://waste-afatekapi-cluster-ip/setDevice", data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_HTTP_POST
			w.Write(resultVal.ToByte())
			return
		}
		resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_SERIAL_DEVICE, currentHttpHeader.DeviceNo)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DEVICE_NOTFOUND
			w.Write(resultVal.ToByte())
			return
		}
	}
	var deviceIdStr = resultVal.Retval.(string)
	resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_DEVICES, deviceIdStr)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_DEVICE_NOTFOUND
		w.Write(resultVal.ToByte())
		return
	}
	var currentDevice WasteLibrary.DeviceType = WasteLibrary.StringToDeviceType(resultVal.Retval.(string))
	currentHttpHeader.CustomerId = currentDevice.CustomerId
	currentHttpHeader.DeviceId = currentDevice.DeviceId

	var serviceClusterIp string = ""
	if currentHttpHeader.AppType == WasteLibrary.APPTYPE_RFID {

		if currentHttpHeader.OpType == WasteLibrary.OPTYPE_RF {

			currentHttpHeader.BaseDataType = WasteLibrary.BASETYPE_TAG
			serviceClusterIp = "waste-rfreader-cluster-ip"

		} else if currentHttpHeader.OpType == WasteLibrary.OPTYPE_GPS {

			currentHttpHeader.BaseDataType = WasteLibrary.BASETYPE_DEVICE
			serviceClusterIp = "waste-gpsreader-cluster-ip"

		} else if currentHttpHeader.OpType == WasteLibrary.OPTYPE_STATUS {

			currentHttpHeader.BaseDataType = WasteLibrary.BASETYPE_DEVICE
			serviceClusterIp = "waste-statusreader-cluster-ip"

		} else if currentHttpHeader.OpType == WasteLibrary.OPTYPE_THERM {

			currentHttpHeader.BaseDataType = WasteLibrary.BASETYPE_DEVICE
			serviceClusterIp = "waste-thermreader-cluster-ip"

		} else if currentHttpHeader.OpType == WasteLibrary.OPTYPE_CAM {

			currentHttpHeader.BaseDataType = WasteLibrary.BASETYPE_TAG
			serviceClusterIp = "waste-camreader-cluster-ip"

		} else {
			resultVal.Result = WasteLibrary.RESULT_FAIL
		}
	} else if currentHttpHeader.AppType == WasteLibrary.APPTYPE_ULT {
		if currentHttpHeader.OpType == WasteLibrary.OPTYPE_SENS {

			currentHttpHeader.BaseDataType = WasteLibrary.BASETYPE_DEVICE
			serviceClusterIp = "waste-ultreader-cluster-ip"

		} else if currentHttpHeader.OpType == WasteLibrary.OPTYPE_ATMP {

			currentHttpHeader.BaseDataType = WasteLibrary.BASETYPE_DEVICE
			serviceClusterIp = "waste-alarmreader-cluster-ip"

		} else if currentHttpHeader.OpType == WasteLibrary.OPTYPE_AGPS {

			currentHttpHeader.BaseDataType = WasteLibrary.BASETYPE_DEVICE
			serviceClusterIp = "waste-alarmreader-cluster-ip"

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
