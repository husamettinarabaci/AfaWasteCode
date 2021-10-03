package main

import (
	"net/http"
	"net/url"

	"github.com/AfatekDevelopers/result_lib_go/devafatekresult"
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

	var resultVal devafatekresult.ResultType
	resultVal.Result = "FAIL"
	if err := req.ParseForm(); err != nil {
		WasteLibrary.LogErr(err)
		return
	}

	var currentHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.StringToHttpClientHeaderType(req.FormValue("HEADER"))
	resultVal = WasteLibrary.GetRedisForStoreApi("serial-device", currentHttpHeader.DeviceNo)
	if resultVal.Result == "FAIL" {

		var createHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.HttpClientHeaderType{
			AppType:      "ADMIN",
			DeviceNo:     "",
			OpType:       "DEVICE",
			Time:         WasteLibrary.GetTime(),
			Repeat:       "0",
			DeviceId:     0,
			CustomerId:   0,
			BaseDataType: "DEVICE",
		}

		var createDevice WasteLibrary.DeviceType = WasteLibrary.DeviceType{
			DeviceId:     0,
			CustomerId:   -1,
			SerialNumber: currentHttpHeader.DeviceNo,
		}
		WasteLibrary.LogStr(createDevice.ToString())
		data := url.Values{
			"HEADER": {createHttpHeader.ToString()},
			"DATA":   {createDevice.ToString()},
		}

		resultVal = WasteLibrary.HttpPostReq("http://waste-afatekapi-cluster-ip/setDevice", data)
		resultVal = WasteLibrary.GetRedisForStoreApi("serial-device", currentHttpHeader.DeviceNo)
	}
	var deviceIdStr = resultVal.Retval.(string)
	var currentDevice WasteLibrary.DeviceType = WasteLibrary.StringToDeviceType(WasteLibrary.GetRedisForStoreApi("devices", deviceIdStr).Retval.(string))
	currentHttpHeader.CustomerId = currentDevice.CustomerId
	currentHttpHeader.DeviceId = currentDevice.DeviceId

	var serviceClusterIp string = ""
	if currentHttpHeader.AppType == "RFID" {

		if currentHttpHeader.OpType == "RF" {

			currentHttpHeader.BaseDataType = "TAG"
			serviceClusterIp = "waste-rfreader-cluster-ip"

		} else if currentHttpHeader.OpType == "GPS" {

			currentHttpHeader.BaseDataType = "DEVICE"
			serviceClusterIp = "waste-gpsreader-cluster-ip"

		} else if currentHttpHeader.OpType == "STATUS" {

			currentHttpHeader.BaseDataType = "DEVICE"
			serviceClusterIp = "waste-statusreader-cluster-ip"

		} else if currentHttpHeader.OpType == "THERM" {

			currentHttpHeader.BaseDataType = "DEVICE"
			serviceClusterIp = "waste-thermreader-cluster-ip"

		} else if currentHttpHeader.OpType == "CAM" {

			currentHttpHeader.BaseDataType = "TAG"
			serviceClusterIp = "waste-camreader-cluster-ip"

		} else {
			resultVal.Result = "FAIL"
		}
	} else if currentHttpHeader.AppType == "ULT" {
		if currentHttpHeader.OpType == "SENS" {

			currentHttpHeader.BaseDataType = "DEVICE"
			serviceClusterIp = "waste-ultreader-cluster-ip"

		} else if currentHttpHeader.OpType == "ATMP" {

			currentHttpHeader.BaseDataType = "DEVICE"
			serviceClusterIp = "waste-alarmreader-cluster-ip"

		} else if currentHttpHeader.OpType == "AGPS" {

			currentHttpHeader.BaseDataType = "DEVICE"
			serviceClusterIp = "waste-alarmreader-cluster-ip"

		} else {
			resultVal.Result = "FAIL"
		}
	} else if currentHttpHeader.AppType == "RECY" {
		resultVal.Result = "OK"
	} else {
		resultVal.Result = "FAIL"
	}

	data := url.Values{
		"HEADER": {currentHttpHeader.ToString()},
		"DATA":   {req.FormValue("DATA")},
	}

	resultVal = WasteLibrary.SaveBulkDbMainForStoreApi(data)

	if serviceClusterIp != "" {
		resultVal = WasteLibrary.HttpPostReq("http://"+serviceClusterIp+"/reader", data)
	}
	w.Write(resultVal.ToByte())
}
