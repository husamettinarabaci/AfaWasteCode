package main

import (
	"net/http"
	"time"

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
		if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_RFID {
			var currentData WasteLibrary.RfidDeviceType = WasteLibrary.StringToRfidDeviceType(req.FormValue(WasteLibrary.HTTP_DATA))
			currentData.DeviceId = currentHttpHeader.DeviceId
			currentData.DeviceGps.DeviceId = currentData.DeviceId

			currentData.DeviceGps.GpsTime = currentHttpHeader.Time
			if currentData.DeviceGps.Longitude != 0 && currentData.DeviceGps.Latitude != 0 {
				resultVal = currentData.DeviceGps.SaveToDb()
				if resultVal.Result != WasteLibrary.RESULT_OK {
					resultVal.Result = WasteLibrary.RESULT_FAIL
					resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
					w.Write(resultVal.ToByte())

					return
				}

				resultVal = currentData.DeviceGps.SaveToRedis()
				if resultVal.Result != WasteLibrary.RESULT_OK {
					resultVal.Result = WasteLibrary.RESULT_FAIL
					resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
					w.Write(resultVal.ToByte())

					return
				}

				resultVal = currentData.DeviceGps.SaveToReaderDb()
				if resultVal.Result != WasteLibrary.RESULT_OK {
					resultVal.Result = WasteLibrary.RESULT_FAIL
					resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
					w.Write(resultVal.ToByte())

					return
				}

				if int(currentData.DeviceGps.DeviceId)%(time.Now().Second()+1) == 0 {
					WasteLibrary.PublishRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_CHANNEL+currentHttpHeader.ToCustomerIdString(), WasteLibrary.DATATYPE_RFID_GPS_DEVICE, currentData.DeviceGps.ToString())
				}
				if currentData.DeviceGps.Speed == 0 {
					var customerConfig WasteLibrary.CustomerConfigType
					customerConfig.CustomerId = currentHttpHeader.CustomerId
					resultVal = customerConfig.GetByRedis()
					if resultVal.Result != WasteLibrary.RESULT_OK {
						resultVal.Result = WasteLibrary.RESULT_FAIL
						resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_CUSTOMERCONFIG_NOTFOUND
						w.Write(resultVal.ToByte())

						return
					}
					if customerConfig.TruckStopTrace == WasteLibrary.STATU_ACTIVE {

						// data = url.Values{
						// WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
						// WasteLibrary.HTTP_DATA:   {currentData.ToString()},
						// }
						//TO DO
						//send gps stop reader
						//resultVal = WasteLibrary.HttpPostReq("http://waste-gpsstopreader-cluster-ip/reader", data)

					}
				}
			} else {
				resultVal.Result = WasteLibrary.RESULT_OK
			}
		} else if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_ULT {
			var currentData WasteLibrary.UltDeviceType = WasteLibrary.StringToUltDeviceType(req.FormValue(WasteLibrary.HTTP_DATA))
			currentData.DeviceId = currentHttpHeader.DeviceId
			currentData.DeviceGps.DeviceId = currentData.DeviceId
			currentData.DeviceGps.GpsTime = currentHttpHeader.Time
			if currentData.DeviceGps.Longitude != 0 && currentData.DeviceGps.Latitude != 0 {
				resultVal = currentData.DeviceGps.SaveToDb()
				if resultVal.Result != WasteLibrary.RESULT_OK {
					resultVal.Result = WasteLibrary.RESULT_FAIL
					resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
					w.Write(resultVal.ToByte())

					return
				}

				resultVal = currentData.DeviceGps.SaveToRedis()
				if resultVal.Result != WasteLibrary.RESULT_OK {
					resultVal.Result = WasteLibrary.RESULT_FAIL
					resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
					w.Write(resultVal.ToByte())

					return
				}

				resultVal = currentData.DeviceGps.SaveToReaderDb()
				if resultVal.Result != WasteLibrary.RESULT_OK {
					resultVal.Result = WasteLibrary.RESULT_FAIL
					resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
					w.Write(resultVal.ToByte())

					return
				}

				WasteLibrary.PublishRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_CHANNEL+currentHttpHeader.ToCustomerIdString(), WasteLibrary.DATATYPE_ULT_GPS_DEVICE, currentData.DeviceGps.ToString())
			} else {
				resultVal.Result = WasteLibrary.RESULT_OK
			}
		}
	} else {
		resultVal.Result = WasteLibrary.RESULT_OK
	}
	w.Write(resultVal.ToByte())

}
