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
	if WasteLibrary.AllowCors {

		w.Header().Set("Access-Control-Allow-Origin", "*")
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
		var currentData WasteLibrary.TagType = WasteLibrary.StringToTagType(req.FormValue(WasteLibrary.HTTP_DATA))
		currentData.DeviceId = currentHttpHeader.DeviceId
		currentData.CustomerId = currentHttpHeader.CustomerId
		WasteLibrary.LogStr("Header : " + currentHttpHeader.ToString())
		WasteLibrary.LogStr("Data : " + currentData.ToString())
		resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_RFID_DEVICES, currentHttpHeader.ToDeviceIdString())
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DEVICE_NOTFOUND
			w.Write(resultVal.ToByte())
			return
		}
		var currentDevice WasteLibrary.RfidDeviceType = WasteLibrary.StringToRfidDeviceType(resultVal.Retval.(string))
		currentData.Latitude = currentDevice.Latitude
		currentData.Longitude = currentDevice.Longitude
		currentData.ReadTime = currentHttpHeader.Time
		data := url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentData.ToString()},
		}
		resultVal = WasteLibrary.SaveStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())
			return
		}

		currentData.TagID = WasteLibrary.StringIdToFloat64(resultVal.Retval.(string))
		data = url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentData.ToString()},
		}
		resultVal = WasteLibrary.GetStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_GET
			w.Write(resultVal.ToByte())
			return
		}
		var currentTag WasteLibrary.TagType = WasteLibrary.StringToTagType(resultVal.Retval.(string))
		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_TAGS, currentTag.ToIdString(), currentTag.ToString())
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())
			return
		}
		resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_TAGS, currentTag.ToCustomerIdString())
		if resultVal.Result == WasteLibrary.RESULT_OK {
			var currentCustomerTags WasteLibrary.CustomerTagsType = WasteLibrary.StringToCustomerTagsType(resultVal.Retval.(string))
			currentCustomerTags.Tags[currentTag.ToIdString()] = currentTag.TagID
			resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_TAGS, currentCustomerTags.ToIdString(), currentCustomerTags.ToString())
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
				w.Write(resultVal.ToByte())
				return
			}
		}

		var newCurrentHttpHeader WasteLibrary.HttpClientHeaderType
		newCurrentHttpHeader.AppType = WasteLibrary.APPTYPE_RFID
		newCurrentHttpHeader.OpType = WasteLibrary.OPTYPE_TAG
		data = url.Values{
			WasteLibrary.HTTP_HEADER: {newCurrentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentTag.ToString()},
		}
		resultVal = WasteLibrary.SaveReaderDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())
			return
		}

	} else {
		resultVal.Result = WasteLibrary.RESULT_OK
	}
	w.Write(resultVal.ToByte())
}
