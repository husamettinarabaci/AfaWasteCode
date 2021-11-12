package main

import (
	"net/http"
	"time"

	"github.com/devafatek/WasteLibrary"
)

var currentCustomerList WasteLibrary.CustomersType
var opInterval time.Duration = 60 * 60

func initStart() {

	WasteLibrary.LogStr("Successfully connected!")
	go WasteLibrary.InitLog()
	currentCustomerList.New()
}

func main() {

	initStart()

	go setCustomerList()

	http.HandleFunc("/health", WasteLibrary.HealthHandler)
	http.HandleFunc("/readiness", WasteLibrary.ReadinessHandler)
	http.HandleFunc("/status", WasteLibrary.StatusHandler)
	http.ListenAndServe(":80", nil)
}

func setCustomerList() {
	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	for {

		resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMERS, WasteLibrary.REDIS_CUSTOMERS)

		if resultVal.Result == WasteLibrary.RESULT_OK {

			var currentCustomers WasteLibrary.CustomersType = WasteLibrary.StringToCustomersType(resultVal.Retval.(string))
			for _, customerId := range currentCustomers.Customers {
				if customerId != 0 {
					if _, ok := currentCustomerList.Customers[WasteLibrary.Float64IdToString(customerId)]; !ok {
						WasteLibrary.LogStr("Add Customer : " + WasteLibrary.Float64IdToString(customerId))
						currentCustomerList.Customers[WasteLibrary.Float64IdToString(customerId)] = customerId
						go customerProc(customerId)
					}
				}
			}
		}
		time.Sleep(opInterval * time.Second)
	}
}

func customerProc(customerId float64) {
	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	for {

		resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_TAGS, WasteLibrary.Float64IdToString(customerId))
		if resultVal.Result == WasteLibrary.RESULT_OK {

			var customerTags WasteLibrary.CustomerTagsType = WasteLibrary.StringToCustomerTagsType(resultVal.Retval.(string))
			var customerTagsList WasteLibrary.CustomerTagsListType
			customerTagsList.New()
			customerTagsList.CustomerId = customerId
			for _, tagId := range customerTags.Tags {

				if tagId != 0 {

					var currentTag WasteLibrary.TagType
					currentTag.New()
					currentTag.TagId = tagId
					resultVal = currentTag.GetAll()
					if resultVal.Result == WasteLibrary.RESULT_OK && currentTag.TagMain.Active == WasteLibrary.STATU_ACTIVE {
						var currentViewTag WasteLibrary.TagViewType
						currentViewTag.New()
						currentViewTag.TagId = currentTag.TagId
						currentViewTag.ContainerNo = currentTag.TagBase.ContainerNo
						currentViewTag.ContainerStatu = currentTag.TagStatu.ContainerStatu
						currentViewTag.TagStatu = currentTag.TagStatu.TagStatu
						currentViewTag.Latitude = currentTag.TagGps.Latitude
						currentViewTag.Longitude = currentTag.TagGps.Longitude
						customerTagsList.Tags = append(customerTagsList.Tags, currentViewTag)
					}
				}
			}
			resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_TAGVIEWS, customerTagsList.ToIdString(), customerTagsList.ToString())
		}

		time.Sleep(opInterval * time.Second)

	}
}
