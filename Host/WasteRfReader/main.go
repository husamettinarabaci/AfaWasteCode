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
	http.HandleFunc("/reader", reader)
	http.ListenAndServe(":80", nil)
}

func reader(w http.ResponseWriter, req *http.Request) {

	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.FAIL
	if err := req.ParseForm(); err != nil {
		WasteLibrary.LogErr(err)
		return
	}
	var currentHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.StringToHttpClientHeaderType(req.FormValue(WasteLibrary.HEADER))
	if currentHttpHeader.Repeat == WasteLibrary.PASSIVE {
		var currentData WasteLibrary.TagType = WasteLibrary.StringToTagType(req.FormValue(WasteLibrary.DATA))
		WasteLibrary.LogStr("Header : " + currentHttpHeader.ToString())
		WasteLibrary.LogStr("Data : " + currentData.ToString())
		var currentDevice WasteLibrary.DeviceType = WasteLibrary.StringToDeviceType(WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_DEVICES, currentHttpHeader.ToDeviceIdString()).Retval.(string))
		currentData.Latitude = currentDevice.Latitude
		currentData.Longitude = currentDevice.Longitude
		currentData.ReadTime = currentHttpHeader.Time
		data := url.Values{
			WasteLibrary.HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.DATA:   {currentData.ToString()},
		}
		resultVal = WasteLibrary.SaveStaticDbMainForStoreApi(data)

		if resultVal.Result == WasteLibrary.OK {

			currentData.TagID = WasteLibrary.StringIdToFloat64(resultVal.Retval.(string))
			data := url.Values{
				WasteLibrary.HEADER: {currentHttpHeader.ToString()},
				WasteLibrary.DATA:   {currentData.ToString()},
			}
			var currentTag WasteLibrary.TagType = WasteLibrary.StringToTagType(WasteLibrary.GetStaticDbMainForStoreApi(data).Retval.(string))

			resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_TAGS, currentTag.ToCustomerIdString())
			if resultVal.Result == WasteLibrary.OK {
				var currentCustomerTags WasteLibrary.CustomerTagsType = WasteLibrary.StringToCustomerTagsType(resultVal.Retval.(string))
				currentCustomerTags.Tags[currentTag.ToIdString()] = currentTag
				resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_TAGS, currentCustomerTags.ToIdString(), currentCustomerTags.ToString())
			}

			var newCurrentHttpHeader WasteLibrary.HttpClientHeaderType
			newCurrentHttpHeader.AppType = WasteLibrary.RFID
			newCurrentHttpHeader.OpType = WasteLibrary.TAG
			data = url.Values{
				WasteLibrary.HEADER: {newCurrentHttpHeader.ToString()},
				WasteLibrary.DATA:   {currentTag.ToString()},
			}
			WasteLibrary.SaveReaderDbMainForStoreApi(data)
		}

	} else {
		resultVal.Result = WasteLibrary.OK
	}
	w.Write(resultVal.ToByte())
}
