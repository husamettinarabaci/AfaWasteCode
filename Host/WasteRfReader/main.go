package main

import (
	"net/http"
	"net/url"
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
			var currentData WasteLibrary.TagType = WasteLibrary.StringToTagType(req.FormValue(WasteLibrary.HTTP_DATA))
			currentData.TagMain.DeviceId = currentHttpHeader.DeviceId
			currentData.TagMain.CustomerId = currentHttpHeader.CustomerId
			var redisTag WasteLibrary.TagType
			resultVal = redisTag.GetByRedisByEpc(currentData.TagMain.Epc)
			if resultVal.Result != WasteLibrary.RESULT_OK {
				var createHttpHeader WasteLibrary.HttpClientHeaderType
				createHttpHeader.New()
				data := url.Values{
					WasteLibrary.HTTP_HEADER: {createHttpHeader.ToString()},
					WasteLibrary.HTTP_DATA:   {currentData.ToString()},
				}

				resultVal = WasteLibrary.HttpPostReq("http://waste-enhcapi-cluster-ip/createTag", data)
				if resultVal.Result != WasteLibrary.RESULT_OK {
					resultVal.Result = WasteLibrary.RESULT_FAIL
					resultVal.Retval = WasteLibrary.RESULT_ERROR_HTTP_POST
					w.Write(resultVal.ToByte())

					return
				}
				resultVal = redisTag.GetByRedisByEpc(currentData.TagMain.Epc)
				if resultVal.Result != WasteLibrary.RESULT_OK {
					resultVal.Result = WasteLibrary.RESULT_FAIL
					resultVal.Retval = WasteLibrary.RESULT_ERROR_TAG_NOTFOUND
					w.Write(resultVal.ToByte())

					return
				}
			}

			currentData.TagId = redisTag.TagId

			if currentData.TagMain.DeviceId != redisTag.TagMain.DeviceId {
				//TagMain
				redisTag.TagMain.DeviceId = currentData.TagMain.DeviceId
				resultVal = redisTag.TagMain.SaveToDb()
				if resultVal.Result != WasteLibrary.RESULT_OK {
					resultVal.Result = WasteLibrary.RESULT_FAIL
					resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
					w.Write(resultVal.ToByte())

					return
				}

				resultVal = redisTag.TagMain.SaveToRedis()
				if resultVal.Result != WasteLibrary.RESULT_OK {
					resultVal.Result = WasteLibrary.RESULT_FAIL
					resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
					w.Write(resultVal.ToByte())

					return
				}

			}

			//TagReader
			currentData.TagReader.TagId = currentData.TagId
			currentData.TagReader.ReadTime = currentHttpHeader.Time

			resultVal = currentData.TagReader.SaveToDb()
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
				w.Write(resultVal.ToByte())

				return
			}

			resultVal = currentData.TagReader.SaveToRedis()
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
				w.Write(resultVal.ToByte())

				return
			}

			resultVal = currentData.TagReader.SaveToReaderDb()
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
				w.Write(resultVal.ToByte())

				return
			}

			WasteLibrary.PublishRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_CHANNEL+currentHttpHeader.ToCustomerIdString(), WasteLibrary.DATATYPE_TAG_READER, currentData.TagReader.ToString())

			//TagGps
			var currentDeviceGps WasteLibrary.RfidDeviceGpsType
			currentDeviceGps.DeviceId = currentHttpHeader.DeviceId
			resultVal = currentDeviceGps.GetByRedis("0")
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_DEVICE_NOTFOUND
				w.Write(resultVal.ToByte())

				return
			}
			var isDeviceGpsCorrect bool = false
			var deviceLatitude float64 = 0
			var deviceLongitude float64 = 0
			var deviceGpsTime string = ""
			if time.Since(WasteLibrary.StringToTime(currentDeviceGps.GpsTime)).Seconds() < 15*60 {
				isDeviceGpsCorrect = true
				deviceLatitude = currentDeviceGps.Latitude
				deviceLongitude = currentDeviceGps.Longitude
				deviceGpsTime = currentDeviceGps.GpsTime

			} else {
				var currentDeviceEmbededGps WasteLibrary.RfidDeviceEmbededGpsType
				currentDeviceEmbededGps.DeviceId = currentHttpHeader.DeviceId
				resultVal = currentDeviceEmbededGps.GetByRedis("0")
				if resultVal.Result != WasteLibrary.RESULT_OK {
					resultVal.Result = WasteLibrary.RESULT_FAIL
					resultVal.Retval = WasteLibrary.RESULT_ERROR_DEVICE_NOTFOUND
					w.Write(resultVal.ToByte())

					return
				}
				if time.Since(WasteLibrary.StringToTime(currentDeviceEmbededGps.GpsTime)).Seconds() < 5*60 {

					isDeviceGpsCorrect = true
					deviceLatitude = currentDeviceEmbededGps.Latitude
					deviceLongitude = currentDeviceEmbededGps.Longitude
					deviceGpsTime = currentDeviceEmbededGps.GpsTime

				}
			}

			if isDeviceGpsCorrect {
				var adminConfig WasteLibrary.AdminConfigType
				adminConfig.New()
				adminConfig.CustomerId = currentHttpHeader.CustomerId
				adminConfig.GetByRedis()
				if len(adminConfig.ParkLocations) > 0 {
					for _, parkLocation := range adminConfig.ParkLocations {
						if WasteLibrary.DistanceInKmBetweenEarthCoordinates(deviceLatitude, deviceLongitude, parkLocation.Latitude, parkLocation.Longitude) < parkLocation.ZoneRadius {
							isDeviceGpsCorrect = false
							break
						}
					}
				}
			}

			if isDeviceGpsCorrect {
				currentData.TagGps.TagId = currentData.TagId
				currentData.TagGps.Latitude = deviceLatitude
				currentData.TagGps.Longitude = deviceLongitude
				currentData.TagGps.GpsTime = deviceGpsTime
				resultVal = currentData.TagGps.SaveToDb()
				if resultVal.Result != WasteLibrary.RESULT_OK {
					resultVal.Result = WasteLibrary.RESULT_FAIL
					resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
					w.Write(resultVal.ToByte())

					return
				}

				resultVal = currentData.TagGps.SaveToRedis()
				if resultVal.Result != WasteLibrary.RESULT_OK {
					resultVal.Result = WasteLibrary.RESULT_FAIL
					resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
					w.Write(resultVal.ToByte())

					return
				}

				resultVal = currentData.TagGps.SaveToReaderDb()
				if resultVal.Result != WasteLibrary.RESULT_OK {
					resultVal.Result = WasteLibrary.RESULT_FAIL
					resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
					w.Write(resultVal.ToByte())

					return
				}

				WasteLibrary.PublishRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_CHANNEL+currentHttpHeader.ToCustomerIdString(), WasteLibrary.DATATYPE_TAG_GPS, currentData.TagGps.ToString())
			}

			//TagStatu
			redisTag.TagStatu.TagId = currentData.TagId
			redisTag.TagStatu.ContainerStatu = WasteLibrary.CONTAINER_FULLNESS_STATU_EMPTY
			redisTag.TagStatu.TagStatu = WasteLibrary.TAG_STATU_READ
			resultVal = redisTag.TagStatu.SaveToDb()
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
				w.Write(resultVal.ToByte())

				return
			}

			resultVal = redisTag.TagStatu.SaveToRedis()
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
				w.Write(resultVal.ToByte())

				return
			}

			resultVal = redisTag.TagStatu.SaveToReaderDb()
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
				w.Write(resultVal.ToByte())

				return
			}

			WasteLibrary.PublishRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_CHANNEL+currentHttpHeader.ToCustomerIdString(), WasteLibrary.DATATYPE_TAG_STATU, redisTag.TagStatu.ToString())
			var customerTagsList WasteLibrary.CustomerTagsViewListType
			customerTagsList.CustomerId = currentHttpHeader.CustomerId
			resultVal = customerTagsList.GetByRedisByReel("0")
			if resultVal.Result == WasteLibrary.RESULT_OK {

				customerTag := customerTagsList.Tags[redisTag.TagStatu.ToIdString()]
				customerTag.ContainerStatu = redisTag.TagStatu.ContainerStatu
				customerTag.TagStatu = redisTag.TagStatu.TagStatu
				customerTag.ReadTime = WasteLibrary.GetTime()
				customerTag.UID = redisTag.TagReader.UID
				customerTag.DeviceId = currentHttpHeader.DeviceId
				customerTag.Latitude = currentData.TagGps.Latitude
				customerTag.Longitude = currentData.TagGps.Longitude
				customerTagsList.Tags[redisTag.TagStatu.ToIdString()] = customerTag
				customerTagsList.SaveToRedisWODb()
			}

		}

	} else {
		resultVal.Result = WasteLibrary.RESULT_OK
	}
	w.Write(resultVal.ToByte())

}
