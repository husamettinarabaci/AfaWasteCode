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
	resultVal.Result = "FAIL"
	if err := req.ParseForm(); err != nil {
		WasteLibrary.LogErr(err)
		return
	}
	var currentHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.StringToHttpClientHeaderType(req.FormValue("HEADER"))
	if currentHttpHeader.Repeat == "0" {
		var currentData WasteLibrary.TagType = WasteLibrary.StringToTagType(req.FormValue("DATA"))
		WasteLibrary.LogStr("Header : " + currentHttpHeader.ToString())
		WasteLibrary.LogStr("Data : " + currentData.ToString())
		var currentDevice WasteLibrary.DeviceType = WasteLibrary.StringToDeviceType(WasteLibrary.GetRedisForStoreApi("devices", currentHttpHeader.ToDeviceIdString()).Retval.(string))
		currentData.Latitude = currentDevice.Latitude
		currentData.Longitude = currentDevice.Longitude
		currentData.ReadTime = currentHttpHeader.Time
		data := url.Values{
			"HEADER": {currentHttpHeader.ToString()},
			"DATA":   {currentData.ToString()},
		}
		resultVal = WasteLibrary.SaveStaticDbMainForStoreApi(data)

		if resultVal.Result == "OK" {

			currentData.TagID = WasteLibrary.StringIdToFloat64(resultVal.Retval.(string))
			data := url.Values{
				"HEADER": {currentHttpHeader.ToString()},
				"DATA":   {currentData.ToString()},
			}
			var currentTag WasteLibrary.TagType = WasteLibrary.StringToTagType(WasteLibrary.GetStaticDbMainForStoreApi(data).Retval.(string))

			resultVal = WasteLibrary.GetRedisForStoreApi("customer-tags", currentTag.ToCustomerIdString())
			if resultVal.Result == "OK" {
				var currentCustomerTags WasteLibrary.CustomerTagsType = WasteLibrary.StringToCustomerTagsType(resultVal.Retval.(string))
				currentCustomerTags.Tags[currentTag.ToIdString()] = currentTag
				resultVal = WasteLibrary.SaveRedisForStoreApi("customer-tags", currentCustomerTags.ToIdString(), currentCustomerTags.ToString())
			}

			var newCurrentHttpHeader WasteLibrary.HttpClientHeaderType
			newCurrentHttpHeader.AppType = "RFID"
			newCurrentHttpHeader.OpType = "TAG"
			data = url.Values{
				"HEADER": {newCurrentHttpHeader.ToString()},
				"DATA":   {currentTag.ToString()},
			}
			WasteLibrary.SaveReaderDbMainForStoreApi(data)
		}

	} else {
		resultVal.Result = "OK"
	}
	w.Write(resultVal.ToByte())
}
