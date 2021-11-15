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
		if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_RFID {
			var currentData WasteLibrary.TagType = WasteLibrary.StringToTagType(req.FormValue(WasteLibrary.HTTP_DATA))
			currentData.TagMain.DeviceId = currentHttpHeader.DeviceId
			currentData.TagMain.CustomerId = currentHttpHeader.CustomerId
			WasteLibrary.LogStr("Header : " + currentHttpHeader.ToString())
			WasteLibrary.LogStr("Data : " + currentData.ToString())
			resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_TAG_EPC, currentData.TagMain.Epc)
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
				resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_TAG_EPC, currentData.TagMain.Epc)
				if resultVal.Result != WasteLibrary.RESULT_OK {
					resultVal.Result = WasteLibrary.RESULT_FAIL
					resultVal.Retval = WasteLibrary.RESULT_ERROR_TAG_NOTFOUND
					w.Write(resultVal.ToByte())

					return
				}
			}

			var redisTag WasteLibrary.TagType
			redisTag.New()
			redisTag.TagId = WasteLibrary.StringIdToFloat64(resultVal.Retval.(string))
			redisTag.GetAll()
			currentData.TagId = redisTag.TagId

			if currentData.TagMain.DeviceId != redisTag.TagMain.DeviceId {
				//TagMain
				redisTag.TagMain.DeviceId = currentData.TagMain.DeviceId
				currentHttpHeader.DataType = WasteLibrary.DATATYPE_TAG_MAIN
				data := url.Values{
					WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
					WasteLibrary.HTTP_DATA:   {redisTag.TagMain.ToString()},
				}
				resultVal = WasteLibrary.SaveStaticDbMainForStoreApi(data)
				if resultVal.Result != WasteLibrary.RESULT_OK {
					resultVal.Result = WasteLibrary.RESULT_FAIL
					resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
					w.Write(resultVal.ToByte())

					return
				}

				redisTag.TagMain.TagId = WasteLibrary.StringIdToFloat64(resultVal.Retval.(string))

				resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_TAG_MAINS, redisTag.TagMain.ToIdString(), redisTag.TagMain.ToString())
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
			currentHttpHeader.DataType = WasteLibrary.DATATYPE_TAG_READER
			WasteLibrary.LogStr("Data Reader : " + currentData.ToString())
			data := url.Values{
				WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
				WasteLibrary.HTTP_DATA:   {currentData.TagReader.ToString()},
			}
			resultVal = WasteLibrary.SaveStaticDbMainForStoreApi(data)
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
				w.Write(resultVal.ToByte())

				return
			}

			currentData.TagReader.TagId = WasteLibrary.StringIdToFloat64(resultVal.Retval.(string))

			resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_TAG_READERS, currentData.TagReader.ToIdString(), currentData.TagReader.ToString())
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
				w.Write(resultVal.ToByte())

				return
			}
			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
				WasteLibrary.HTTP_DATA:   {currentData.TagReader.ToString()},
			}
			resultVal = WasteLibrary.SaveReaderDbMainForStoreApi(data)
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
				w.Write(resultVal.ToByte())

				return
			}

			WasteLibrary.PublishRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_CHANNEL+currentHttpHeader.ToCustomerIdString(), WasteLibrary.DATATYPE_TAG_READER, currentData.TagReader.ToString())

			//TagGps
			resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_RFID_GPS_DEVICES, currentHttpHeader.ToDeviceIdString())
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_DEVICE_NOTFOUND
				w.Write(resultVal.ToByte())

				return
			}
			var currentDeviceGps WasteLibrary.RfidDeviceGpsType = WasteLibrary.StringToRfidDeviceGpsType(resultVal.Retval.(string))
			currentData.TagGps.TagId = currentData.TagId
			currentHttpHeader.DataType = WasteLibrary.DATATYPE_TAG_GPS
			currentData.TagGps.Latitude = currentDeviceGps.Latitude
			currentData.TagGps.Longitude = currentDeviceGps.Longitude
			currentData.TagGps.GpsTime = currentDeviceGps.GpsTime
			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
				WasteLibrary.HTTP_DATA:   {currentData.TagGps.ToString()},
			}
			resultVal = WasteLibrary.SaveStaticDbMainForStoreApi(data)
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
				w.Write(resultVal.ToByte())

				return
			}

			currentData.TagGps.TagId = WasteLibrary.StringIdToFloat64(resultVal.Retval.(string))

			resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_TAG_GPSES, currentData.TagGps.ToIdString(), currentData.TagGps.ToString())
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
				w.Write(resultVal.ToByte())

				return
			}
			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
				WasteLibrary.HTTP_DATA:   {currentData.TagGps.ToString()},
			}
			resultVal = WasteLibrary.SaveReaderDbMainForStoreApi(data)
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
				w.Write(resultVal.ToByte())

				return
			}

			WasteLibrary.PublishRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_CHANNEL+currentHttpHeader.ToCustomerIdString(), WasteLibrary.DATATYPE_TAG_GPS, currentData.TagGps.ToString())

			//TagStatu
			redisTag.TagStatu.TagId = currentData.TagId
			redisTag.TagStatu.ContainerStatu = WasteLibrary.CONTAINER_FULLNESS_STATU_EMPTY
			redisTag.TagStatu.TagStatu = WasteLibrary.TAG_STATU_READ
			currentHttpHeader.DataType = WasteLibrary.DATATYPE_TAG_STATU
			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
				WasteLibrary.HTTP_DATA:   {redisTag.TagStatu.ToString()},
			}
			resultVal = WasteLibrary.SaveStaticDbMainForStoreApi(data)
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
				w.Write(resultVal.ToByte())

				return
			}

			redisTag.TagStatu.TagId = WasteLibrary.StringIdToFloat64(resultVal.Retval.(string))

			resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_TAG_STATUS, redisTag.TagStatu.ToIdString(), redisTag.TagStatu.ToString())
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
				w.Write(resultVal.ToByte())

				return
			}
			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
				WasteLibrary.HTTP_DATA:   {redisTag.TagStatu.ToString()},
			}
			resultVal = WasteLibrary.SaveReaderDbMainForStoreApi(data)
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
				w.Write(resultVal.ToByte())

				return
			}

			WasteLibrary.PublishRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_CHANNEL+currentHttpHeader.ToCustomerIdString(), WasteLibrary.DATATYPE_TAG_STATU, redisTag.TagStatu.ToString())
		}

	} else {
		resultVal.Result = WasteLibrary.RESULT_OK
	}
	w.Write(resultVal.ToByte())

}
