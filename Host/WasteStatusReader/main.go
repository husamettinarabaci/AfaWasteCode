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
	http.HandleFunc("/reader", reader)
	http.ListenAndServe(":80", nil)
}

func reader(w http.ResponseWriter, req *http.Request) {

	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.FAIL
	if err := req.ParseForm(); err != nil {
		WasteLibrary.LogErr(err)
		return
	}
	var currentHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.StringToHttpClientHeaderType(req.FormValue(WasteLibrary.HEADER))

	if currentHttpHeader.Repeat == WasteLibrary.PASSIVE {
		var currentData WasteLibrary.DeviceType = WasteLibrary.StringToDeviceType(req.FormValue(WasteLibrary.DATA))
		WasteLibrary.LogStr("Header : " + currentHttpHeader.ToString())
		WasteLibrary.LogStr("Data : " + currentData.ToString())
		currentData.StatusTime = currentHttpHeader.Time
		data := url.Values{
			WasteLibrary.HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.DATA:   {currentData.ToString()},
		}
		resultVal = WasteLibrary.SaveStaticDbMainForStoreApi(data)

		if resultVal.Result == WasteLibrary.OK {
			currentData.DeviceId = WasteLibrary.StringIdToFloat64(resultVal.Retval.(string))

			if currentData.ReaderAppStatus == WasteLibrary.ACTIVE {
				currentData.ReaderAppLastOkTime = currentHttpHeader.Time
			}
			if currentData.ReaderConnStatus == WasteLibrary.ACTIVE {
				currentData.ReaderConnLastOkTime = currentHttpHeader.Time
			}
			if currentData.ReaderStatus == WasteLibrary.ACTIVE {
				currentData.ReaderLastOkTime = currentHttpHeader.Time
			}

			if currentData.CamAppStatus == WasteLibrary.ACTIVE {
				currentData.CamAppLastOkTime = currentHttpHeader.Time
			}
			if currentData.CamConnStatus == WasteLibrary.ACTIVE {
				currentData.CamConnLastOkTime = currentHttpHeader.Time
			}
			if currentData.CamStatus == WasteLibrary.ACTIVE {
				currentData.CamLastOkTime = currentHttpHeader.Time
			}

			if currentData.GpsAppStatus == WasteLibrary.ACTIVE {
				currentData.GpsAppLastOkTime = currentHttpHeader.Time
			}
			if currentData.GpsConnStatus == WasteLibrary.ACTIVE {
				currentData.GpsConnLastOkTime = currentHttpHeader.Time
			}
			if currentData.GpsStatus == WasteLibrary.ACTIVE {
				currentData.GpsLastOkTime = currentHttpHeader.Time
			}

			if currentData.ThermAppStatus == WasteLibrary.ACTIVE {
				currentData.ThermAppLastOkTime = currentHttpHeader.Time
			}
			if currentData.TransferAppStatus == WasteLibrary.ACTIVE {
				currentData.TransferAppLastOkTime = currentHttpHeader.Time
			}
			if currentData.AliveStatus == WasteLibrary.ACTIVE {
				currentData.AliveLastOkTime = currentHttpHeader.Time
			}
			if currentData.ContactStatus == WasteLibrary.ACTIVE {
				currentData.ContactLastOkTime = currentHttpHeader.Time
			}

			data := url.Values{
				WasteLibrary.HEADER: {currentHttpHeader.ToString()},
				WasteLibrary.DATA:   {currentData.ToString()},
			}
			var currentDevice WasteLibrary.DeviceType = WasteLibrary.StringToDeviceType(WasteLibrary.GetStaticDbMainForStoreApi(data).Retval.(string))
			resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_DEVICES, currentDevice.ToIdString(), currentDevice.ToString())

			var newCurrentHttpHeader WasteLibrary.HttpClientHeaderType
			newCurrentHttpHeader.AppType = WasteLibrary.RFID
			newCurrentHttpHeader.OpType = WasteLibrary.DEVICE
			data = url.Values{
				WasteLibrary.HEADER: {newCurrentHttpHeader.ToString()},
				WasteLibrary.DATA:   {currentDevice.ToString()},
			}
			WasteLibrary.SaveReaderDbMainForStoreApi(data)
		}

	} else {
		resultVal.Result = WasteLibrary.OK
	}
	w.Write(resultVal.ToByte())
}
