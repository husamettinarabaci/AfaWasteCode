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

	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.FAIL
	if err := req.ParseForm(); err != nil {
		WasteLibrary.LogErr(err)
		return
	}

	var currentHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.StringToHttpClientHeaderType(req.FormValue(WasteLibrary.HEADER))
	resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_SERIAL_DEVICE, currentHttpHeader.DeviceNo)
	if resultVal.Result == WasteLibrary.FAIL {

		var createHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.HttpClientHeaderType{
			AppType:      WasteLibrary.ADMIN,
			DeviceNo:     "",
			OpType:       WasteLibrary.DEVICE,
			Time:         WasteLibrary.GetTime(),
			Repeat:       WasteLibrary.PASSIVE,
			DeviceId:     0,
			CustomerId:   0,
			BaseDataType: WasteLibrary.DEVICE,
		}

		var createDevice WasteLibrary.DeviceType = WasteLibrary.DeviceType{
			DeviceId:     0,
			CustomerId:   -1,
			SerialNumber: currentHttpHeader.DeviceNo,
		}
		WasteLibrary.LogStr(createDevice.ToString())
		data := url.Values{
			WasteLibrary.HEADER: {createHttpHeader.ToString()},
			WasteLibrary.DATA:   {createDevice.ToString()},
		}

		resultVal = WasteLibrary.HttpPostReq("http://waste-afatekapi-cluster-ip/setDevice", data)
		resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_SERIAL_DEVICE, currentHttpHeader.DeviceNo)
	}
	var deviceIdStr = resultVal.Retval.(string)
	var currentDevice WasteLibrary.DeviceType = WasteLibrary.StringToDeviceType(WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_DEVICES, deviceIdStr).Retval.(string))
	currentHttpHeader.CustomerId = currentDevice.CustomerId
	currentHttpHeader.DeviceId = currentDevice.DeviceId

	var serviceClusterIp string = ""
	if currentHttpHeader.AppType == WasteLibrary.RFID {

		if currentHttpHeader.OpType == WasteLibrary.RF {

			currentHttpHeader.BaseDataType = WasteLibrary.TAG
			serviceClusterIp = "waste-rfreader-cluster-ip"

		} else if currentHttpHeader.OpType == WasteLibrary.GPS {

			currentHttpHeader.BaseDataType = WasteLibrary.DEVICE
			serviceClusterIp = "waste-gpsreader-cluster-ip"

		} else if currentHttpHeader.OpType == WasteLibrary.STATUS {

			currentHttpHeader.BaseDataType = WasteLibrary.DEVICE
			serviceClusterIp = "waste-statusreader-cluster-ip"

		} else if currentHttpHeader.OpType == WasteLibrary.THERM {

			currentHttpHeader.BaseDataType = WasteLibrary.DEVICE
			serviceClusterIp = "waste-thermreader-cluster-ip"

		} else if currentHttpHeader.OpType == WasteLibrary.CAM {

			currentHttpHeader.BaseDataType = WasteLibrary.TAG
			serviceClusterIp = "waste-camreader-cluster-ip"

		} else {
			resultVal.Result = WasteLibrary.FAIL
		}
	} else if currentHttpHeader.AppType == WasteLibrary.ULT {
		if currentHttpHeader.OpType == WasteLibrary.SENS {

			currentHttpHeader.BaseDataType = WasteLibrary.DEVICE
			serviceClusterIp = "waste-ultreader-cluster-ip"

		} else if currentHttpHeader.OpType == WasteLibrary.ATMP {

			currentHttpHeader.BaseDataType = WasteLibrary.DEVICE
			serviceClusterIp = "waste-alarmreader-cluster-ip"

		} else if currentHttpHeader.OpType == WasteLibrary.AGPS {

			currentHttpHeader.BaseDataType = WasteLibrary.DEVICE
			serviceClusterIp = "waste-alarmreader-cluster-ip"

		} else {
			resultVal.Result = WasteLibrary.FAIL
		}
	} else if currentHttpHeader.AppType == WasteLibrary.RECY {
		resultVal.Result = WasteLibrary.OK
	} else {
		resultVal.Result = WasteLibrary.FAIL
	}

	data := url.Values{
		WasteLibrary.HEADER: {currentHttpHeader.ToString()},
		WasteLibrary.DATA:   {req.FormValue(WasteLibrary.DATA)},
	}

	resultVal = WasteLibrary.SaveBulkDbMainForStoreApi(data)

	if serviceClusterIp != "" {
		resultVal = WasteLibrary.HttpPostReq("http://"+serviceClusterIp+"/reader", data)
	}
	w.Write(resultVal.ToByte())
}
