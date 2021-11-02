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

	if currentHttpHeader.Repeat == WasteLibrary.STATU_PASSIVE {
		var currentData WasteLibrary.RfidDeviceType = WasteLibrary.StringToRfidDeviceType(req.FormValue(WasteLibrary.HTTP_DATA))
		currentData.DeviceId = currentHttpHeader.DeviceId
		currentData.CustomerId = currentHttpHeader.CustomerId
		WasteLibrary.LogStr("Header : " + currentHttpHeader.ToString())
		WasteLibrary.LogStr("Data : " + currentData.ToString())
		currentData.StatusTime = currentHttpHeader.Time
		data := url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentData.ToString()},
		}
		resultVal = WasteLibrary.SaveStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())
			return
		}

		currentData.DeviceId = WasteLibrary.StringIdToFloat64(resultVal.Retval.(string))

		if currentData.ReaderAppStatus == WasteLibrary.STATU_ACTIVE {
			currentData.ReaderAppLastOkTime = currentHttpHeader.Time
		}
		if currentData.ReaderConnStatus == WasteLibrary.STATU_ACTIVE {
			currentData.ReaderConnLastOkTime = currentHttpHeader.Time
		}
		if currentData.ReaderStatus == WasteLibrary.STATU_ACTIVE {
			currentData.ReaderLastOkTime = currentHttpHeader.Time
		}

		if currentData.CamAppStatus == WasteLibrary.STATU_ACTIVE {
			currentData.CamAppLastOkTime = currentHttpHeader.Time
		}
		if currentData.CamConnStatus == WasteLibrary.STATU_ACTIVE {
			currentData.CamConnLastOkTime = currentHttpHeader.Time
		}
		if currentData.CamStatus == WasteLibrary.STATU_ACTIVE {
			currentData.CamLastOkTime = currentHttpHeader.Time
		}

		if currentData.GpsAppStatus == WasteLibrary.STATU_ACTIVE {
			currentData.GpsAppLastOkTime = currentHttpHeader.Time
		}
		if currentData.GpsConnStatus == WasteLibrary.STATU_ACTIVE {
			currentData.GpsConnLastOkTime = currentHttpHeader.Time
		}
		if currentData.GpsStatus == WasteLibrary.STATU_ACTIVE {
			currentData.GpsLastOkTime = currentHttpHeader.Time
		}

		if currentData.ThermAppStatus == WasteLibrary.STATU_ACTIVE {
			currentData.ThermAppLastOkTime = currentHttpHeader.Time
		}
		if currentData.TransferAppStatus == WasteLibrary.STATU_ACTIVE {
			currentData.TransferAppLastOkTime = currentHttpHeader.Time
		}
		if currentData.SystemAppStatus == WasteLibrary.STATU_ACTIVE {
			currentData.SystemAppLastOkTime = currentHttpHeader.Time
		}
		if currentData.UpdaterAppStatus == WasteLibrary.STATU_ACTIVE {
			currentData.UpdaterAppLastOkTime = currentHttpHeader.Time
		}
		if currentData.AliveStatus == WasteLibrary.STATU_ACTIVE {
			currentData.AliveLastOkTime = currentHttpHeader.Time
		}
		if currentData.ContactStatus == WasteLibrary.STATU_ACTIVE {
			currentData.ContactLastOkTime = currentHttpHeader.Time
		}

		data = url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentData.ToString()},
		}
		resultVal = WasteLibrary.GetStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_GET
			w.Write(resultVal.ToByte())
			return
		}
		var currentDevice WasteLibrary.RfidDeviceType = WasteLibrary.StringToRfidDeviceType(resultVal.Retval.(string))
		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_RFID_DEVICES, currentDevice.ToIdString(), currentDevice.ToString())
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())
			return
		}
		var newCurrentHttpHeader WasteLibrary.HttpClientHeaderType
		newCurrentHttpHeader.AppType = WasteLibrary.APPTYPE_RFID
		newCurrentHttpHeader.ReaderType = WasteLibrary.OPTYPE_DEVICE
		data = url.Values{
			WasteLibrary.HTTP_HEADER: {newCurrentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentDevice.ToString()},
		}
		resultVal = WasteLibrary.SaveReaderDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())
			return
		}

	} else {
		resultVal.Result = WasteLibrary.RESULT_OK
	}
	w.Write(resultVal.ToByte())
}
