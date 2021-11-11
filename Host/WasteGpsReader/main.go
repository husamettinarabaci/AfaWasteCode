package main

import (
	"net/http"
	"net/url"
	"time"

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
	WasteLibrary.LogStr("Header : " + currentHttpHeader.ToString())
	WasteLibrary.LogStr("Data : " + req.FormValue(WasteLibrary.HTTP_DATA))
	if currentHttpHeader.Repeat == WasteLibrary.STATU_PASSIVE {
		if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_RFID {
			var currentData WasteLibrary.RfidDeviceType = WasteLibrary.StringToRfidDeviceType(req.FormValue(WasteLibrary.HTTP_DATA))
			currentData.DeviceId = currentHttpHeader.DeviceId
			currentData.DeviceGps.DeviceId = currentData.DeviceId

			currentHttpHeader.DataType = WasteLibrary.DATATYPE_RFID_GPS_DEVICE
			currentData.DeviceGps.GpsTime = currentHttpHeader.Time
			if currentData.DeviceGps.Longitude != 0 && currentData.DeviceGps.Latitude != 0 {
				data := url.Values{
					WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
					WasteLibrary.HTTP_DATA:   {currentData.DeviceGps.ToString()},
				}
				resultVal = WasteLibrary.SaveStaticDbMainForStoreApi(data)
				if resultVal.Result != WasteLibrary.RESULT_OK {
					resultVal.Result = WasteLibrary.RESULT_FAIL
					resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
					w.Write(resultVal.ToByte())

					return
				}

				currentData.DeviceGps.DeviceId = WasteLibrary.StringIdToFloat64(resultVal.Retval.(string))
				data = url.Values{
					WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
					WasteLibrary.HTTP_DATA:   {currentData.DeviceGps.ToString()},
				}
				resultVal = WasteLibrary.GetStaticDbMainForStoreApi(data)
				if resultVal.Result != WasteLibrary.RESULT_OK {
					resultVal.Result = WasteLibrary.RESULT_FAIL
					resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_GET
					w.Write(resultVal.ToByte())

					return
				}
				currentData.DeviceGps = WasteLibrary.StringToRfidDeviceGpsType(resultVal.Retval.(string))
				resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_RFID_GPS_DEVICES, currentData.DeviceGps.ToIdString(), currentData.DeviceGps.ToString())
				if resultVal.Result != WasteLibrary.RESULT_OK {
					resultVal.Result = WasteLibrary.RESULT_FAIL
					resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
					w.Write(resultVal.ToByte())

					return
				}
				data = url.Values{
					WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
					WasteLibrary.HTTP_DATA:   {currentData.DeviceGps.ToString()},
				}
				resultVal = WasteLibrary.SaveReaderDbMainForStoreApi(data)
				if resultVal.Result != WasteLibrary.RESULT_OK {
					resultVal.Result = WasteLibrary.RESULT_FAIL
					resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
					w.Write(resultVal.ToByte())

					return
				}

				WasteLibrary.PublishRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_CHANNEL+currentHttpHeader.ToCustomerIdString(), WasteLibrary.DATATYPE_RFID_GPS_DEVICE, currentData.DeviceGps.ToString())
				if currentData.DeviceGps.Speed == 0 {
					resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_CUSTOMERCONFIG, currentHttpHeader.ToCustomerIdString())
					if resultVal.Result != WasteLibrary.RESULT_OK {
						resultVal.Result = WasteLibrary.RESULT_FAIL
						resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_CUSTOMERCONFIG_NOTFOUND
						w.Write(resultVal.ToByte())

						return
					}
					var customerConfig WasteLibrary.CustomerConfigType = WasteLibrary.StringToCustomerConfigType(resultVal.Retval.(string))
					if customerConfig.TruckStopTrace == WasteLibrary.STATU_ACTIVE {

						resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_TAGS, currentHttpHeader.ToCustomerIdString())
						if resultVal.Result == WasteLibrary.RESULT_OK {

							var customerTags WasteLibrary.CustomerTagsType = WasteLibrary.StringToCustomerTagsType(resultVal.Retval.(string))
							for _, tagId := range customerTags.Tags {

								if tagId != 0 {

									var currentTag WasteLibrary.TagType
									currentTag.New()
									currentTag.TagId = tagId
									resultVal = currentTag.GetAll()
									if resultVal.Result == WasteLibrary.RESULT_OK && currentTag.TagMain.Active == WasteLibrary.STATU_ACTIVE {
										second := time.Since(WasteLibrary.StringToTime(currentTag.TagReader.ReadTime)).Seconds()
										if second < 1*60*60 {
											var distance float64 = WasteLibrary.DistanceInKmBetweenEarthCoordinates(currentTag.TagGps.Latitude, currentTag.TagGps.Longitude, currentData.DeviceGps.Latitude, currentData.DeviceGps.Longitude)
											if distance < 50 {

												currentTag.TagStatu.TagId = currentTag.TagId
												currentTag.TagStatu.CheckTime = WasteLibrary.GetTime()
												currentTag.TagStatu.ContainerStatu = WasteLibrary.CONTAINER_FULLNESS_STATU_EMPTY
												currentTag.TagStatu.TagStatu = WasteLibrary.TAG_STATU_STOP
												currentHttpHeader.DataType = WasteLibrary.DATATYPE_TAG_STATU
												data = url.Values{
													WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
													WasteLibrary.HTTP_DATA:   {currentTag.TagStatu.ToString()},
												}
												resultVal = WasteLibrary.SaveStaticDbMainForStoreApi(data)
												if resultVal.Result != WasteLibrary.RESULT_OK {
													resultVal.Result = WasteLibrary.RESULT_FAIL
													resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
													w.Write(resultVal.ToByte())

													return
												}

												currentTag.TagStatu.TagId = WasteLibrary.StringIdToFloat64(resultVal.Retval.(string))
												data = url.Values{
													WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
													WasteLibrary.HTTP_DATA:   {currentTag.TagStatu.ToString()},
												}
												resultVal = WasteLibrary.GetStaticDbMainForStoreApi(data)
												if resultVal.Result != WasteLibrary.RESULT_OK {
													resultVal.Result = WasteLibrary.RESULT_FAIL
													resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_GET
													w.Write(resultVal.ToByte())

													return
												}
												currentTag.TagStatu = WasteLibrary.StringToTagStatuType(resultVal.Retval.(string))
												resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_TAG_STATUS, currentTag.TagStatu.ToIdString(), currentTag.TagStatu.ToString())
												if resultVal.Result != WasteLibrary.RESULT_OK {
													resultVal.Result = WasteLibrary.RESULT_FAIL
													resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
													w.Write(resultVal.ToByte())

													return
												}
												data = url.Values{
													WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
													WasteLibrary.HTTP_DATA:   {currentTag.TagStatu.ToString()},
												}
												resultVal = WasteLibrary.SaveReaderDbMainForStoreApi(data)
												if resultVal.Result != WasteLibrary.RESULT_OK {
													resultVal.Result = WasteLibrary.RESULT_FAIL
													resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
													w.Write(resultVal.ToByte())

													return
												}

												WasteLibrary.PublishRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_CHANNEL+currentHttpHeader.ToCustomerIdString(), WasteLibrary.DATATYPE_TAG_STATU, currentTag.TagStatu.ToString())

											}
										}
									}
								}
							}

						}

					}
				}
			} else {
				resultVal.Result = WasteLibrary.RESULT_OK
			}
		} else if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_ULT {
			var currentData WasteLibrary.UltDeviceType = WasteLibrary.StringToUltDeviceType(req.FormValue(WasteLibrary.HTTP_DATA))
			currentData.DeviceId = currentHttpHeader.DeviceId
			currentData.DeviceGps.DeviceId = currentData.DeviceId
			currentHttpHeader.DataType = WasteLibrary.DATATYPE_ULT_GPS_DEVICE
			currentData.DeviceGps.GpsTime = currentHttpHeader.Time
			if currentData.DeviceGps.Longitude != 0 && currentData.DeviceGps.Latitude != 0 {
				data := url.Values{
					WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
					WasteLibrary.HTTP_DATA:   {currentData.DeviceGps.ToString()},
				}
				resultVal = WasteLibrary.SaveStaticDbMainForStoreApi(data)
				if resultVal.Result != WasteLibrary.RESULT_OK {
					resultVal.Result = WasteLibrary.RESULT_FAIL
					resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
					w.Write(resultVal.ToByte())

					return
				}

				currentData.DeviceGps.DeviceId = WasteLibrary.StringIdToFloat64(resultVal.Retval.(string))
				data = url.Values{
					WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
					WasteLibrary.HTTP_DATA:   {currentData.DeviceGps.ToString()},
				}
				resultVal = WasteLibrary.GetStaticDbMainForStoreApi(data)
				if resultVal.Result != WasteLibrary.RESULT_OK {
					resultVal.Result = WasteLibrary.RESULT_FAIL
					resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_GET
					w.Write(resultVal.ToByte())

					return
				}
				currentData.DeviceGps = WasteLibrary.StringToUltDeviceGpsType(resultVal.Retval.(string))
				resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_ULT_GPS_DEVICES, currentData.DeviceGps.ToIdString(), currentData.DeviceGps.ToString())
				if resultVal.Result != WasteLibrary.RESULT_OK {
					resultVal.Result = WasteLibrary.RESULT_FAIL
					resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
					w.Write(resultVal.ToByte())

					return
				}
				data = url.Values{
					WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
					WasteLibrary.HTTP_DATA:   {currentData.DeviceGps.ToString()},
				}
				resultVal = WasteLibrary.SaveReaderDbMainForStoreApi(data)
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
