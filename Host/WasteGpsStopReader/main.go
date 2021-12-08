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
	var currentHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.StringToHttpClientHeaderType(req.FormValue(WasteLibrary.HTTP_HEADER))
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

	var customerTagsList WasteLibrary.CustomerTagsViewListType
	customerTagsList.CustomerId = currentHttpHeader.CustomerId
	resultVal = customerTagsList.GetByRedis("0")
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_GET
		return
	}

	for _, currentViewTag := range customerTagsList.Tags {
		var distance float64 = WasteLibrary.DistanceInKmBetweenEarthCoordinates(currentViewTag.Latitude, currentViewTag.Longitude, currentData.DeviceGps.Latitude, currentData.DeviceGps.Longitude)
		if distance < 50 {

			var currentTag WasteLibrary.TagType
			currentTag.New()
			currentTag.TagId = currentViewTag.TagId
			resultVal = currentTag.GetByRedis("0")
			if resultVal.Result == WasteLibrary.RESULT_OK && currentTag.TagMain.Active == WasteLibrary.STATU_ACTIVE {
				second := time.Since(WasteLibrary.StringToTime(currentTag.TagReader.ReadTime)).Seconds()
				if second < 1*60*60 {

					currentTag.TagStatu.TagId = currentTag.TagId
					currentTag.TagStatu.CheckTime = WasteLibrary.GetTime()
					currentTag.TagStatu.ContainerStatu = WasteLibrary.CONTAINER_FULLNESS_STATU_EMPTY
					currentTag.TagStatu.TagStatu = WasteLibrary.TAG_STATU_STOP
					resultVal = currentTag.TagStatu.SaveToDb()
					if resultVal.Result != WasteLibrary.RESULT_OK {
						resultVal.Result = WasteLibrary.RESULT_FAIL
						resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
						continue
					}

					resultVal = currentTag.TagStatu.SaveToRedis()
					if resultVal.Result != WasteLibrary.RESULT_OK {
						resultVal.Result = WasteLibrary.RESULT_FAIL
						resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
						continue
					}

					resultVal = currentTag.TagStatu.SaveToReaderDb()
					if resultVal.Result != WasteLibrary.RESULT_OK {
						resultVal.Result = WasteLibrary.RESULT_FAIL
						resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
						continue
					}

					WasteLibrary.PublishRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_CHANNEL+currentHttpHeader.ToCustomerIdString(), WasteLibrary.DATATYPE_TAG_STATU, currentTag.TagStatu.ToString())
					var customerTagsList WasteLibrary.CustomerTagsViewListType
					customerTagsList.CustomerId = currentHttpHeader.CustomerId
					resultVal = customerTagsList.GetByRedisByReel("0")
					if resultVal.Result == WasteLibrary.RESULT_OK {

						customerTag := customerTagsList.Tags[currentTag.TagStatu.ToIdString()]
						customerTag.ContainerStatu = currentTag.TagStatu.ContainerStatu
						customerTag.TagStatu = currentTag.TagStatu.TagStatu
						customerTagsList.Tags[currentTag.TagStatu.ToIdString()] = customerTag
						customerTagsList.SaveToRedisWODb()
					}
				}
			}
		}
	}
}
