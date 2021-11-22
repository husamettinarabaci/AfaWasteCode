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
			currentData.DeviceBattery.DeviceId = currentData.DeviceId
			if WasteLibrary.StringIdToFloat64(currentData.DeviceBattery.Battery) > 3300 {
				currentData.DeviceBattery.BatteryStatus = WasteLibrary.BATTERYSTATU_NORMAL
			} else {
				currentData.DeviceBattery.BatteryStatus = WasteLibrary.BATTERYSTATU_LOW
			}
			currentHttpHeader.DataType = WasteLibrary.DATATYPE_ULT_BATTERY_DEVICE
			currentData.DeviceBattery.BatteryTime = currentHttpHeader.Time
			data := url.Values{
				WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
				WasteLibrary.HTTP_DATA:   {currentData.DeviceBattery.ToString()},
			}
			resultVal = WasteLibrary.SaveStaticDbMainForStoreApi(data)
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
				w.Write(resultVal.ToByte())

				return
			}

			currentData.DeviceBattery.DeviceId = WasteLibrary.StringIdToFloat64(resultVal.Retval.(string))

			resultVal = currentData.DeviceBattery.SaveToRedis()
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
				w.Write(resultVal.ToByte())

				return
			}
			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
				WasteLibrary.HTTP_DATA:   {currentData.DeviceBattery.ToString()},
			}
			resultVal = WasteLibrary.SaveReaderDbMainForStoreApi(data)
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
				w.Write(resultVal.ToByte())

				return
			}

			WasteLibrary.PublishRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_CHANNEL+currentHttpHeader.ToCustomerIdString(), WasteLibrary.DATATYPE_ULT_BATTERY_DEVICE, currentData.DeviceBattery.ToString())
		}

	} else {
		resultVal.Result = WasteLibrary.RESULT_OK
	}
	w.Write(resultVal.ToByte())

}
