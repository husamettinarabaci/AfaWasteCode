package main

import (
	"fmt"
	"net/http"

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

	w.Write(resultVal.ToByte())

	var customerTags WasteLibrary.CustomerTagsType = WasteLibrary.StringToCustomerTagsType(req.FormValue(WasteLibrary.HTTP_DATA))
	go customerProc(customerTags)

}

func customerProc(customerTags WasteLibrary.CustomerTagsType) {
	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL

	var customerTagsList WasteLibrary.CustomerTagsViewListType
	customerTagsList.New()
	customerTagsList.CustomerId = customerTags.CustomerId
	for _, tagId := range customerTags.Tags {

		if tagId != 0 {

			var currentTag WasteLibrary.TagType
			currentTag.New()
			currentTag.TagId = tagId
			resultVal = currentTag.GetByRedis("0")
			if resultVal.Result == WasteLibrary.RESULT_OK && currentTag.TagMain.Active == WasteLibrary.STATU_ACTIVE {
				var currentViewTag WasteLibrary.TagViewType
				currentViewTag.New()
				currentViewTag.TagId = currentTag.TagId
				currentViewTag.DeviceId = currentTag.TagMain.DeviceId
				if currentTag.TagBase.ContainerNo == "" {
					currentViewTag.ContainerNo = fmt.Sprintf("%05d", int(currentTag.TagId))
				} else {
					currentViewTag.ContainerNo = currentTag.TagBase.ContainerNo
				}
				currentViewTag.ContainerStatu = currentTag.TagStatu.ContainerStatu
				currentViewTag.TagStatu = currentTag.TagStatu.TagStatu
				currentViewTag.Latitude = currentTag.TagGps.Latitude
				currentViewTag.Longitude = currentTag.TagGps.Longitude
				currentViewTag.ReadTime = currentTag.TagReader.ReadTime
				currentViewTag.UID = currentTag.TagReader.UID
				customerTagsList.Tags[currentViewTag.ToIdString()] = currentViewTag
			}
		}
	}
	resultVal = customerTagsList.SaveToRedis()
	resultVal = customerTagsList.SaveToRedisWODb()

}
