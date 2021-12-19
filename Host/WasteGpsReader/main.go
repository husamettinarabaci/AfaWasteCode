package main

import (
	"math/rand"
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
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,access-control-allow-origin, access-control-allow-headers")
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
	var currentHttpHeader WasteLibrary.HttpClientHeaderType
	currentHttpHeader.StringToType(req.FormValue(WasteLibrary.HTTP_HEADER))
	if currentHttpHeader.Repeat == WasteLibrary.STATU_PASSIVE {
		if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_RFID {
			if currentHttpHeader.ReaderType == WasteLibrary.READERTYPE_GPS_API {
				WasteLibrary.LogStr("Data : " + req.FormValue(WasteLibrary.HTTP_DATA))
				var currentData WasteLibrary.RfidDeviceType
				currentData.StringToType(req.FormValue(WasteLibrary.HTTP_DATA))
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
					if int(currentData.DeviceGps.DeviceId)%(rand.Intn(10)+1) == 0 {
						WasteLibrary.PublishRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_CHANNEL+currentHttpHeader.ToCustomerIdString(), WasteLibrary.DATATYPE_RFID_GPS, currentData.DeviceGps.ToString())
					}

					var customerDevicesList WasteLibrary.CustomerRfidDevicesViewListType
					customerDevicesList.CustomerId = currentHttpHeader.CustomerId
					resultVal = customerDevicesList.GetByRedisByReel("0")
					if resultVal.Result == WasteLibrary.RESULT_OK {

						customerDeviceView := customerDevicesList.Devices[currentData.ToIdString()]
						customerDeviceView.Latitude = currentData.DeviceGps.Latitude
						customerDeviceView.Longitude = currentData.DeviceGps.Longitude
						customerDevicesList.Devices[currentData.ToIdString()] = customerDeviceView
						customerDevicesList.SaveToRedisWODb()
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

							//TO DO
							//gps stop data
							//data := url.Values{
							//	WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
							//	WasteLibrary.HTTP_DATA:   {currentData.ToString()},
							//}
							//resultVal = WasteLibrary.HttpPostReq("http://waste-gpsstopreader-cluster-ip/reader", data)

						}
					}
				} else {
					resultVal.Result = WasteLibrary.RESULT_OK
					resultVal.Retval = WasteLibrary.RESULT_SUCCESS_OK
				}
			} else {
				WasteLibrary.LogStr("Data : " + req.FormValue(WasteLibrary.HTTP_DATA))
				var currentData WasteLibrary.RfidDeviceType
				currentData.StringToType(req.FormValue(WasteLibrary.HTTP_DATA))
				currentData.DeviceId = currentHttpHeader.DeviceId
				var currentEmbGps WasteLibrary.RfidDeviceEmbededGpsType
				currentEmbGps.New()
				currentEmbGps.DeviceId = currentData.DeviceId

				currentEmbGps.GpsTime = currentHttpHeader.Time
				if currentData.DeviceGps.Longitude != 0 && currentData.DeviceGps.Latitude != 0 {

					currentEmbGps.Longitude = currentData.DeviceGps.Longitude
					currentEmbGps.Latitude = currentData.DeviceGps.Latitude
					resultVal = currentEmbGps.SaveToDb()
					if resultVal.Result != WasteLibrary.RESULT_OK {
						resultVal.Result = WasteLibrary.RESULT_FAIL
						resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
						w.Write(resultVal.ToByte())

						return
					}

					resultVal = currentEmbGps.SaveToRedis()
					if resultVal.Result != WasteLibrary.RESULT_OK {
						resultVal.Result = WasteLibrary.RESULT_FAIL
						resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
						w.Write(resultVal.ToByte())

						return
					}

					resultVal = currentEmbGps.SaveToReaderDb()
					if resultVal.Result != WasteLibrary.RESULT_OK {
						resultVal.Result = WasteLibrary.RESULT_FAIL
						resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
						w.Write(resultVal.ToByte())

						return
					}

					var oldDeviceGps WasteLibrary.RfidDeviceGpsType
					oldDeviceGps.New()
					oldDeviceGps.DeviceId = currentData.DeviceId
					oldDeviceGps.GetByRedis("0")
					if time.Since(WasteLibrary.StringToTime(oldDeviceGps.GpsTime)) > 15*60 {
						var customerDevicesList WasteLibrary.CustomerRfidDevicesViewListType
						customerDevicesList.CustomerId = currentHttpHeader.CustomerId
						resultVal = customerDevicesList.GetByRedisByReel("0")
						if resultVal.Result == WasteLibrary.RESULT_OK {

							customerDeviceView := customerDevicesList.Devices[currentData.ToIdString()]
							customerDeviceView.Latitude = currentEmbGps.Latitude
							customerDeviceView.Longitude = currentEmbGps.Longitude
							customerDevicesList.Devices[currentData.ToIdString()] = customerDeviceView
							customerDevicesList.SaveToRedisWODb()
						}
						if int(currentEmbGps.DeviceId)%(rand.Intn(10)+1) == 0 {
							WasteLibrary.PublishRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_CHANNEL+currentHttpHeader.ToCustomerIdString(), WasteLibrary.DATATYPE_RFID_GPS, currentEmbGps.ToString())
						}
					}

				} else {
					resultVal.Result = WasteLibrary.RESULT_OK
					resultVal.Retval = WasteLibrary.RESULT_SUCCESS_OK
				}
			}
		} else if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_ULT {
			var currentData WasteLibrary.UltDeviceType
			currentData.StringToType(req.FormValue(WasteLibrary.HTTP_DATA))
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

				WasteLibrary.PublishRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_CHANNEL+currentHttpHeader.ToCustomerIdString(), WasteLibrary.DATATYPE_ULT_GPS, currentData.DeviceGps.ToString())
			} else {
				resultVal.Result = WasteLibrary.RESULT_OK
				resultVal.Retval = WasteLibrary.RESULT_SUCCESS_OK
			}
		}
	} else {
		resultVal.Result = WasteLibrary.RESULT_OK
		resultVal.Retval = WasteLibrary.RESULT_SUCCESS_OK
	}
	w.Write(resultVal.ToByte())

}
