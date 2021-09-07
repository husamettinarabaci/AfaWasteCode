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
	http.HandleFunc("/reader", reader)
	http.ListenAndServe(":80", nil)
}

func reader(w http.ResponseWriter, req *http.Request) {

	var resultVal devafatekresult.ResultType
	resultVal.Result = "FAIL"
	if err := req.ParseForm(); err != nil {
		WasteLibrary.LogErr(err)
		return
	}
	var currentHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.StringToHttpClientHeaderType(req.FormValue("HEADER"))

	if currentHttpHeader.Repeat == "0" {
		var currentData WasteLibrary.DeviceType = WasteLibrary.StringToDeviceType(req.FormValue("DATA"))
		WasteLibrary.LogStr(currentHttpHeader.ToString() + " - " + currentData.ToString())
		currentData.StatusTime = currentHttpHeader.Time
		data := url.Values{
			"HEADER": {currentHttpHeader.ToString()},
			"DATA":   {currentData.ToString()},
		}
		resultVal = WasteLibrary.SaveStaticDbMainForStoreApi(data)

		if resultVal.Result == "OK" {
			currentData.DeviceId = WasteLibrary.StringIdToFloat64(resultVal.Retval.(string))

			if currentData.ReaderAppStatus == "1" {
				currentData.ReaderAppLastOkTime = currentHttpHeader.Time
			}
			if currentData.ReaderConnStatus == "1" {
				currentData.ReaderConnLastOkTime = currentHttpHeader.Time
			}
			if currentData.ReaderStatus == "1" {
				currentData.ReaderLastOkTime = currentHttpHeader.Time
			}

			if currentData.CamAppStatus == "1" {
				currentData.CamAppLastOkTime = currentHttpHeader.Time
			}
			if currentData.CamConnStatus == "1" {
				currentData.CamConnLastOkTime = currentHttpHeader.Time
			}
			if currentData.CamStatus == "1" {
				currentData.CamLastOkTime = currentHttpHeader.Time
			}

			if currentData.GpsAppStatus == "1" {
				currentData.GpsAppLastOkTime = currentHttpHeader.Time
			}
			if currentData.GpsConnStatus == "1" {
				currentData.GpsConnLastOkTime = currentHttpHeader.Time
			}
			if currentData.GpsStatus == "1" {
				currentData.GpsLastOkTime = currentHttpHeader.Time
			}

			if currentData.ThermAppStatus == "1" {
				currentData.ThermAppLastOkTime = currentHttpHeader.Time
			}
			if currentData.TransferAppStatus == "1" {
				currentData.TransferAppLastOkTime = currentHttpHeader.Time
			}
			if currentData.AliveStatus == "1" {
				currentData.AliveLastOkTime = currentHttpHeader.Time
			}
			if currentData.ContactStatus == "1" {
				currentData.ContactLastOkTime = currentHttpHeader.Time
			}

			data := url.Values{
				"HEADER": {currentHttpHeader.ToString()},
				"DATA":   {currentData.ToString()},
			}
			var currentDevice WasteLibrary.DeviceType = WasteLibrary.StringToDeviceType(WasteLibrary.GetStaticDbMainForStoreApi(data).Retval.(string))
			resultVal = WasteLibrary.SaveRedisForStoreApi("devices", currentDevice.ToIdString(), currentDevice.ToString())

			var newCurrentHttpHeader WasteLibrary.HttpClientHeaderType
			newCurrentHttpHeader.AppType = "RFID"
			newCurrentHttpHeader.OpType = "DEVICE"
			data = url.Values{
				"HEADER": {newCurrentHttpHeader.ToString()},
				"DATA":   {currentDevice.ToString()},
			}
			WasteLibrary.SaveReaderDbMainForStoreApi(data)
		}

	} else {
		resultVal.Result = "OK"
	}
	w.Write(resultVal.ToByte())
}
