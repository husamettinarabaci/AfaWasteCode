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
	resultVal.Result = WasteLibrary.RESULT_FAIL
	if err := req.ParseForm(); err != nil {
		WasteLibrary.LogErr(err)
		return
	}
	var currentHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.StringToHttpClientHeaderType(req.FormValue(WasteLibrary.HTTP_HEADER))

	if currentHttpHeader.Repeat == WasteLibrary.STATU_PASSIVE {
		var currentData WasteLibrary.DeviceType = WasteLibrary.StringToDeviceType(req.FormValue(WasteLibrary.HTTP_DATA))
		WasteLibrary.LogStr("Header : " + currentHttpHeader.ToString())
		WasteLibrary.LogStr("Data : " + currentData.ToString())
		currentData.GpsTime = currentHttpHeader.Time
		if currentData.Longitude != 0 && currentData.Latitude != 0 {
			data := url.Values{
				WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
				WasteLibrary.HTTP_DATA:   {currentData.ToString()},
			}
			resultVal = WasteLibrary.SaveStaticDbMainForStoreApi(data)

			if resultVal.Result == WasteLibrary.RESULT_OK {
				currentData.DeviceId = WasteLibrary.StringIdToFloat64(resultVal.Retval.(string))
				data := url.Values{
					WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
					WasteLibrary.HTTP_DATA:   {currentData.ToString()},
				}
				var currentDevice WasteLibrary.DeviceType = WasteLibrary.StringToDeviceType(WasteLibrary.GetStaticDbMainForStoreApi(data).Retval.(string))
				resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_DEVICES, currentDevice.ToIdString(), currentDevice.ToString())

				var newCurrentHttpHeader WasteLibrary.HttpClientHeaderType
				newCurrentHttpHeader.AppType = WasteLibrary.APPTYPE_RFID
				newCurrentHttpHeader.OpType = WasteLibrary.OPTYPE_DEVICE
				data = url.Values{
					WasteLibrary.HTTP_HEADER: {newCurrentHttpHeader.ToString()},
					WasteLibrary.HTTP_DATA:   {currentDevice.ToString()},
				}
				WasteLibrary.SaveReaderDbMainForStoreApi(data)
			}

		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}
	} else {
		resultVal.Result = WasteLibrary.RESULT_OK
	}
	w.Write(resultVal.ToByte())
}
