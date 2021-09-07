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
		currentData.GpsTime = currentHttpHeader.Time
		if currentData.Longitude != 0 && currentData.Latitude != 0 {
			data := url.Values{
				"HEADER": {currentHttpHeader.ToString()},
				"DATA":   {currentData.ToString()},
			}
			resultVal = WasteLibrary.SaveStaticDbMainForStoreApi(data)

			if resultVal.Result == "OK" {
				currentData.DeviceId = WasteLibrary.StringIdToFloat64(resultVal.Retval.(string))
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
	} else {
		resultVal.Result = "OK"
	}
	w.Write(resultVal.ToByte())
}
