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
	http.HandleFunc("/openLog", WasteLibrary.OpenLogHandler)
	http.HandleFunc("/closeLog", WasteLibrary.CloseLogHandler)
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

			var oldSensData WasteLibrary.UltDeviceSensType
			resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_ULT_SENS_DEVICES, currentData.ToIdString())
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_NOTFOUND
				w.Write(resultVal.ToByte())

				return
			}
			oldSensData = WasteLibrary.StringToUltDeviceSensType(resultVal.Retval.(string))

			if currentData.DeviceSens.UltCount < 24 {
				currentData.DeviceSens.UltRange24 = oldSensData.UltRange24
			}
			if currentData.DeviceSens.UltCount < 23 {
				currentData.DeviceSens.UltRange23 = oldSensData.UltRange23
			}
			if currentData.DeviceSens.UltCount < 22 {
				currentData.DeviceSens.UltRange22 = oldSensData.UltRange22
			}
			if currentData.DeviceSens.UltCount < 21 {
				currentData.DeviceSens.UltRange21 = oldSensData.UltRange21
			}
			if currentData.DeviceSens.UltCount < 20 {
				currentData.DeviceSens.UltRange20 = oldSensData.UltRange20
			}
			if currentData.DeviceSens.UltCount < 19 {
				currentData.DeviceSens.UltRange19 = oldSensData.UltRange19
			}
			if currentData.DeviceSens.UltCount < 18 {
				currentData.DeviceSens.UltRange18 = oldSensData.UltRange18
			}
			if currentData.DeviceSens.UltCount < 17 {
				currentData.DeviceSens.UltRange17 = oldSensData.UltRange17
			}
			if currentData.DeviceSens.UltCount < 16 {
				currentData.DeviceSens.UltRange16 = oldSensData.UltRange16
			}
			if currentData.DeviceSens.UltCount < 15 {
				currentData.DeviceSens.UltRange15 = oldSensData.UltRange15
			}
			if currentData.DeviceSens.UltCount < 14 {
				currentData.DeviceSens.UltRange14 = oldSensData.UltRange14
			}
			if currentData.DeviceSens.UltCount < 13 {
				currentData.DeviceSens.UltRange13 = oldSensData.UltRange13
			}
			if currentData.DeviceSens.UltCount < 12 {
				currentData.DeviceSens.UltRange12 = oldSensData.UltRange12
			}
			if currentData.DeviceSens.UltCount < 11 {
				currentData.DeviceSens.UltRange11 = oldSensData.UltRange11
			}
			if currentData.DeviceSens.UltCount < 10 {
				currentData.DeviceSens.UltRange10 = oldSensData.UltRange10
			}
			if currentData.DeviceSens.UltCount < 9 {
				currentData.DeviceSens.UltRange9 = oldSensData.UltRange9
			}
			if currentData.DeviceSens.UltCount < 8 {
				currentData.DeviceSens.UltRange8 = oldSensData.UltRange8
			}
			if currentData.DeviceSens.UltCount < 7 {
				currentData.DeviceSens.UltRange7 = oldSensData.UltRange7
			}
			if currentData.DeviceSens.UltCount < 6 {
				currentData.DeviceSens.UltRange6 = oldSensData.UltRange6
			}
			if currentData.DeviceSens.UltCount < 5 {
				currentData.DeviceSens.UltRange5 = oldSensData.UltRange5
			}
			if currentData.DeviceSens.UltCount < 4 {
				currentData.DeviceSens.UltRange4 = oldSensData.UltRange4
			}
			if currentData.DeviceSens.UltCount < 3 {
				currentData.DeviceSens.UltRange3 = oldSensData.UltRange3
			}
			if currentData.DeviceSens.UltCount < 2 {
				currentData.DeviceSens.UltRange2 = oldSensData.UltRange2
			}

			//TO DO
			//calculate ult status by container type
			ultCm := currentData.DeviceSens.UltRange1 * 173 / 10000
			if ultCm < 50 {
				currentData.DeviceSens.UltStatus = WasteLibrary.CONTAINER_FULLNESS_STATU_FULL
			}
			if ultCm >= 50 && ultCm < 100 {
				currentData.DeviceSens.UltStatus = WasteLibrary.CONTAINER_FULLNESS_STATU_MEDIUM
			}

			if ultCm >= 100 && ultCm < 150 {
				currentData.DeviceSens.UltStatus = WasteLibrary.CONTAINER_FULLNESS_STATU_LITTLE
			}

			if ultCm >= 150 {
				currentData.DeviceSens.UltStatus = WasteLibrary.CONTAINER_FULLNESS_STATU_EMPTY
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
			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
				WasteLibrary.HTTP_DATA:   {currentData.DeviceSens.ToString()},
			}
			resultVal = WasteLibrary.GetStaticDbMainForStoreApi(data)
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_GET
				w.Write(resultVal.ToByte())

				return
			}
			currentData.DeviceSens = WasteLibrary.StringToUltDeviceSensType(resultVal.Retval.(string))
			resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_ULT_SENS_DEVICES, currentData.DeviceSens.ToIdString(), currentData.DeviceSens.ToString())
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
