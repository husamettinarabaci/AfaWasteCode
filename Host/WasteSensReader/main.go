package main

import (
	"net/http"
	"net/url"

	"github.com/devafatek/WasteLibrary"
)

func initStart() {

	WasteLibrary.LogStr("Successfully connected!")
	go WasteLibrary.InitLog()
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
	if currentHttpHeader.Repeat == WasteLibrary.STATU_PASSIVE {
		if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_ULT {
			var currentData WasteLibrary.UltDeviceType = WasteLibrary.StringToUltDeviceType(req.FormValue(WasteLibrary.HTTP_DATA))
			currentData.DeviceId = currentHttpHeader.DeviceId
			currentData.DeviceSens.DeviceId = currentData.DeviceId
			WasteLibrary.LogStr("Header : " + currentHttpHeader.ToString())
			WasteLibrary.LogStr("Data : " + currentData.ToString())
			currentHttpHeader.DataType = WasteLibrary.DATATYPE_ULT_SENS_DEVICE
			currentData.DeviceSens.UltTime = currentHttpHeader.Time

			var oldData WasteLibrary.UltDeviceType
			oldData.DeviceId = currentData.DeviceId
			oldData.GetByRedis()

			if currentData.DeviceSens.UltCount < 24 {
				currentData.DeviceSens.UltRange24 = oldData.DeviceSens.UltRange24
			}
			if currentData.DeviceSens.UltCount < 23 {
				currentData.DeviceSens.UltRange23 = oldData.DeviceSens.UltRange23
			}
			if currentData.DeviceSens.UltCount < 22 {
				currentData.DeviceSens.UltRange22 = oldData.DeviceSens.UltRange22
			}
			if currentData.DeviceSens.UltCount < 21 {
				currentData.DeviceSens.UltRange21 = oldData.DeviceSens.UltRange21
			}
			if currentData.DeviceSens.UltCount < 20 {
				currentData.DeviceSens.UltRange20 = oldData.DeviceSens.UltRange20
			}
			if currentData.DeviceSens.UltCount < 19 {
				currentData.DeviceSens.UltRange19 = oldData.DeviceSens.UltRange19
			}
			if currentData.DeviceSens.UltCount < 18 {
				currentData.DeviceSens.UltRange18 = oldData.DeviceSens.UltRange18
			}
			if currentData.DeviceSens.UltCount < 17 {
				currentData.DeviceSens.UltRange17 = oldData.DeviceSens.UltRange17
			}
			if currentData.DeviceSens.UltCount < 16 {
				currentData.DeviceSens.UltRange16 = oldData.DeviceSens.UltRange16
			}
			if currentData.DeviceSens.UltCount < 15 {
				currentData.DeviceSens.UltRange15 = oldData.DeviceSens.UltRange15
			}
			if currentData.DeviceSens.UltCount < 14 {
				currentData.DeviceSens.UltRange14 = oldData.DeviceSens.UltRange14
			}
			if currentData.DeviceSens.UltCount < 13 {
				currentData.DeviceSens.UltRange13 = oldData.DeviceSens.UltRange13
			}
			if currentData.DeviceSens.UltCount < 12 {
				currentData.DeviceSens.UltRange12 = oldData.DeviceSens.UltRange12
			}
			if currentData.DeviceSens.UltCount < 11 {
				currentData.DeviceSens.UltRange11 = oldData.DeviceSens.UltRange11
			}
			if currentData.DeviceSens.UltCount < 10 {
				currentData.DeviceSens.UltRange10 = oldData.DeviceSens.UltRange10
			}
			if currentData.DeviceSens.UltCount < 9 {
				currentData.DeviceSens.UltRange9 = oldData.DeviceSens.UltRange9
			}
			if currentData.DeviceSens.UltCount < 8 {
				currentData.DeviceSens.UltRange8 = oldData.DeviceSens.UltRange8
			}
			if currentData.DeviceSens.UltCount < 7 {
				currentData.DeviceSens.UltRange7 = oldData.DeviceSens.UltRange7
			}
			if currentData.DeviceSens.UltCount < 6 {
				currentData.DeviceSens.UltRange6 = oldData.DeviceSens.UltRange6
			}
			if currentData.DeviceSens.UltCount < 5 {
				currentData.DeviceSens.UltRange5 = oldData.DeviceSens.UltRange5
			}
			if currentData.DeviceSens.UltCount < 4 {
				currentData.DeviceSens.UltRange4 = oldData.DeviceSens.UltRange4
			}
			if currentData.DeviceSens.UltCount < 3 {
				currentData.DeviceSens.UltRange3 = oldData.DeviceSens.UltRange3
			}
			if currentData.DeviceSens.UltCount < 2 {
				currentData.DeviceSens.UltRange2 = oldData.DeviceSens.UltRange2
			}

			data := url.Values{
				WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
				WasteLibrary.HTTP_DATA:   {currentData.DeviceSens.ToString()},
			}
			resultVal = WasteLibrary.SaveStaticDbMainForStoreApi(data)
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
				w.Write(resultVal.ToByte())

				return
			}

			currentData.DeviceSens.DeviceId = WasteLibrary.StringIdToFloat64(resultVal.Retval.(string))

			resultVal = currentData.DeviceSens.SaveToRedis()
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
				w.Write(resultVal.ToByte())

				return
			}
			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
				WasteLibrary.HTTP_DATA:   {currentData.DeviceSens.ToString()},
			}
			resultVal = WasteLibrary.SaveReaderDbMainForStoreApi(data)
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
				w.Write(resultVal.ToByte())

				return
			}

			WasteLibrary.PublishRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_CHANNEL+currentHttpHeader.ToCustomerIdString(), WasteLibrary.DATATYPE_ULT_SENS_DEVICE, currentData.DeviceSens.ToString())
		}
	} else {
		resultVal.Result = WasteLibrary.RESULT_OK
	}
	w.Write(resultVal.ToByte())

}
