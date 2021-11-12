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
	WasteLibrary.LogStr("Header : " + currentHttpHeader.ToString())
	WasteLibrary.LogStr("Data : " + req.FormValue(WasteLibrary.HTTP_DATA))
	if currentHttpHeader.Repeat == WasteLibrary.STATU_PASSIVE {
		if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_RFID {
			var currentData WasteLibrary.RfidDeviceType = WasteLibrary.StringToRfidDeviceType(req.FormValue(WasteLibrary.HTTP_DATA))
			go procGpsStopDevice(currentData, currentHttpHeader)
		}
	} else {
		resultVal.Result = WasteLibrary.RESULT_OK
	}
	w.Write(resultVal.ToByte())

}

func procGpsStopDevice(currentData WasteLibrary.RfidDeviceType, currentHttpHeader WasteLibrary.HttpClientHeaderType) {
	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL

	resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_TAGVIEWS, currentHttpHeader.ToCustomerIdString())
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_GET
		WasteLibrary.LogStr(resultVal.ToString())
		return
	}

	var customerTagsList WasteLibrary.CustomerTagsListType = WasteLibrary.StringToCustomerTagsListType(resultVal.Retval.(string))
	for _, currentViewTag := range customerTagsList.Tags {
		var distance float64 = WasteLibrary.DistanceInKmBetweenEarthCoordinates(currentViewTag.Latitude, currentViewTag.Longitude, currentData.DeviceGps.Latitude, currentData.DeviceGps.Longitude)
		if distance < 50 {

			var currentTag WasteLibrary.TagType
			currentTag.New()
			currentTag.TagId = currentViewTag.TagId
			resultVal = currentTag.GetAll()
			WasteLibrary.LogStr("Stop Operation Device : " + currentData.ToString())
			WasteLibrary.LogStr("Stop Operation Distance : " + WasteLibrary.Float64IdToString(distance))
			WasteLibrary.LogStr("Stop Operation Tag : " + currentTag.ToString())
			if resultVal.Result == WasteLibrary.RESULT_OK && currentTag.TagMain.Active == WasteLibrary.STATU_ACTIVE {
				second := time.Since(WasteLibrary.StringToTime(currentTag.TagReader.ReadTime)).Seconds()
				if second < 1*60*60 {

					currentTag.TagStatu.TagId = currentTag.TagId
					currentTag.TagStatu.CheckTime = WasteLibrary.GetTime()
					currentTag.TagStatu.ContainerStatu = WasteLibrary.CONTAINER_FULLNESS_STATU_EMPTY
					currentTag.TagStatu.TagStatu = WasteLibrary.TAG_STATU_STOP
					currentHttpHeader.DataType = WasteLibrary.DATATYPE_TAG_STATU
					data := url.Values{
						WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
						WasteLibrary.HTTP_DATA:   {currentTag.TagStatu.ToString()},
					}
					resultVal = WasteLibrary.SaveStaticDbMainForStoreApi(data)
					if resultVal.Result != WasteLibrary.RESULT_OK {
						resultVal.Result = WasteLibrary.RESULT_FAIL
						resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
						WasteLibrary.LogStr(resultVal.ToString())
						continue
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
						WasteLibrary.LogStr(resultVal.ToString())
						continue
					}
					currentTag.TagStatu = WasteLibrary.StringToTagStatuType(resultVal.Retval.(string))
					resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_TAG_STATUS, currentTag.TagStatu.ToIdString(), currentTag.TagStatu.ToString())
					if resultVal.Result != WasteLibrary.RESULT_OK {
						resultVal.Result = WasteLibrary.RESULT_FAIL
						resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
						WasteLibrary.LogStr(resultVal.ToString())
						continue
					}
					data = url.Values{
						WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
						WasteLibrary.HTTP_DATA:   {currentTag.TagStatu.ToString()},
					}
					resultVal = WasteLibrary.SaveReaderDbMainForStoreApi(data)
					if resultVal.Result != WasteLibrary.RESULT_OK {
						resultVal.Result = WasteLibrary.RESULT_FAIL
						resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
						WasteLibrary.LogStr(resultVal.ToString())
						continue
					}

					WasteLibrary.PublishRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_CHANNEL+currentHttpHeader.ToCustomerIdString(), WasteLibrary.DATATYPE_TAG_STATU, currentTag.TagStatu.ToString())

				}
			}
		}

	}
}
