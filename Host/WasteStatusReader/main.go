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
		if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_RFID {
			var currentData WasteLibrary.RfidDeviceType = WasteLibrary.StringToRfidDeviceType(req.FormValue(WasteLibrary.HTTP_DATA))
			currentData.DeviceId = currentHttpHeader.DeviceId
			currentData.DeviceStatu.DeviceId = currentData.DeviceId
			WasteLibrary.LogStr("Header : " + currentHttpHeader.ToString())
			WasteLibrary.LogStr("Data : " + currentData.ToString())
			currentHttpHeader.DataType = WasteLibrary.DATATYPE_RFID_STATU_DEVICE
			currentData.DeviceStatu.StatusTime = currentHttpHeader.Time

			if currentData.DeviceStatu.ReaderAppStatus == WasteLibrary.STATU_ACTIVE {
				currentData.DeviceStatu.ReaderAppLastOkTime = currentHttpHeader.Time
			}
			if currentData.DeviceStatu.ReaderConnStatus == WasteLibrary.STATU_ACTIVE {
				currentData.DeviceStatu.ReaderConnLastOkTime = currentHttpHeader.Time
			}
			if currentData.DeviceStatu.ReaderStatus == WasteLibrary.STATU_ACTIVE {
				currentData.DeviceStatu.ReaderLastOkTime = currentHttpHeader.Time
			}

			if currentData.DeviceStatu.CamAppStatus == WasteLibrary.STATU_ACTIVE {
				currentData.DeviceStatu.CamAppLastOkTime = currentHttpHeader.Time
			}
			if currentData.DeviceStatu.CamConnStatus == WasteLibrary.STATU_ACTIVE {
				currentData.DeviceStatu.CamConnLastOkTime = currentHttpHeader.Time
			}
			if currentData.DeviceStatu.CamStatus == WasteLibrary.STATU_ACTIVE {
				currentData.DeviceStatu.CamLastOkTime = currentHttpHeader.Time
			}

			if currentData.DeviceStatu.GpsAppStatus == WasteLibrary.STATU_ACTIVE {
				currentData.DeviceStatu.GpsAppLastOkTime = currentHttpHeader.Time
			}
			if currentData.DeviceStatu.GpsConnStatus == WasteLibrary.STATU_ACTIVE {
				currentData.DeviceStatu.GpsConnLastOkTime = currentHttpHeader.Time
			}
			if currentData.DeviceStatu.GpsStatus == WasteLibrary.STATU_ACTIVE {
				currentData.DeviceStatu.GpsLastOkTime = currentHttpHeader.Time
			}

			if currentData.DeviceStatu.ThermAppStatus == WasteLibrary.STATU_ACTIVE {
				currentData.DeviceStatu.ThermAppLastOkTime = currentHttpHeader.Time
			}
			if currentData.DeviceStatu.TransferAppStatus == WasteLibrary.STATU_ACTIVE {
				currentData.DeviceStatu.TransferAppLastOkTime = currentHttpHeader.Time
			}
			if currentData.DeviceStatu.SystemAppStatus == WasteLibrary.STATU_ACTIVE {
				currentData.DeviceStatu.SystemAppLastOkTime = currentHttpHeader.Time
			}
			if currentData.DeviceStatu.UpdaterAppStatus == WasteLibrary.STATU_ACTIVE {
				currentData.DeviceStatu.UpdaterAppLastOkTime = currentHttpHeader.Time
			}
			if currentData.DeviceStatu.AliveStatus == WasteLibrary.STATU_ACTIVE {
				currentData.DeviceStatu.AliveLastOkTime = currentHttpHeader.Time
			}
			if currentData.DeviceStatu.ContactStatus == WasteLibrary.STATU_ACTIVE {
				currentData.DeviceStatu.ContactLastOkTime = currentHttpHeader.Time
			}

			WasteLibrary.LogReport("Device Statu Status Reader 1: " + currentData.DeviceStatu.ToString())
			data := url.Values{
				WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
				WasteLibrary.HTTP_DATA:   {currentData.DeviceStatu.ToString()},
			}
			resultVal = WasteLibrary.SaveStaticDbMainForStoreApi(data)
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
				w.Write(resultVal.ToByte())

				return
			}

			currentData.DeviceStatu.DeviceId = WasteLibrary.StringIdToFloat64(resultVal.Retval.(string))

			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
				WasteLibrary.HTTP_DATA:   {currentData.DeviceStatu.ToString()},
			}
			resultVal = WasteLibrary.GetStaticDbMainForStoreApi(data)
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_GET
				w.Write(resultVal.ToByte())

				return
			}
			WasteLibrary.LogReport("Device Statu Status Db Reader : " + resultVal.Retval.(string))
			currentData.DeviceStatu = WasteLibrary.StringToRfidDeviceStatuType(resultVal.Retval.(string))
			WasteLibrary.LogReport("Device Statu Status Reader 2: " + currentData.DeviceStatu.ToString())
			resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_RFID_STATU_DEVICES, currentData.DeviceStatu.ToIdString(), currentData.DeviceStatu.ToString())
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
				w.Write(resultVal.ToByte())

				return
			}
			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
				WasteLibrary.HTTP_DATA:   {currentData.DeviceStatu.ToString()},
			}
			resultVal = WasteLibrary.SaveReaderDbMainForStoreApi(data)
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
				w.Write(resultVal.ToByte())

				return
			}
			//TO DO
			//send data web socket
		}

	} else {
		resultVal.Result = WasteLibrary.RESULT_OK
	}
	w.Write(resultVal.ToByte())

}
