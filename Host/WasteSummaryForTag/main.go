package main

import (
	"net/http"
	"net/url"
	"time"

	"github.com/devafatek/WasteLibrary"
)

var currentCustomerList WasteLibrary.CustomersType
var opInterval time.Duration = 60 * 60

func initStart() {

	WasteLibrary.LogStr("Successfully connected!")
	currentCustomerList.New()
}

func main() {

	initStart()

	go setCustomerList()

	http.HandleFunc("/health", WasteLibrary.HealthHandler)
	http.HandleFunc("/readiness", WasteLibrary.ReadinessHandler)
	http.HandleFunc("/status", WasteLibrary.StatusHandler)
	http.HandleFunc("/openLog", WasteLibrary.OpenLogHandler)
	http.HandleFunc("/closeLog", WasteLibrary.CloseLogHandler)
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

		//TO DO
		//vardiya config
		// var customerDevices WasteLibrary.CustomerRfidDevicesType = WasteLibrary.StringToCustomerRfidDevicesType(resultVal.Retval.(string))
		// var customerDevicesList WasteLibrary.CustomerRfidDevicesListType
		// customerDevicesList.New()
		// customerDevicesList.CustomerId = WasteLibrary.StringIdToFloat64(WasteLibrary.Float64IdToString(customerId))
		// for _, deviceId := range customerDevices.Devices {
		//
		// if deviceId != 0 {
		//
		// var currentDevice WasteLibrary.RfidDeviceType
		// currentDevice.New()
		// currentDevice.DeviceId = deviceId
		// resultVal = currentDevice.GetAll()
		// if resultVal.Result == WasteLibrary.RESULT_OK {
		// customerDevicesList.Devices = append(customerDevicesList.Devices, currentDevice)
		// }
		//
		// }
		// }

		var currentHttpHeader WasteLibrary.HttpClientHeaderType
		currentHttpHeader.New()
		currentHttpHeader.DataType = WasteLibrary.DATATYPE_TAG_STATU
		resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_TAGS, WasteLibrary.Float64IdToString(customerId))
		if resultVal.Result == WasteLibrary.RESULT_OK {

			var customerTags WasteLibrary.CustomerTagsType = WasteLibrary.StringToCustomerTagsType(resultVal.Retval.(string))
			for _, tagId := range customerTags.Tags {

				if tagId != 0 {

					var currentTag WasteLibrary.TagType
					currentTag.New()
					currentTag.TagId = tagId
					resultVal = currentTag.GetAll()
					if resultVal.Result == WasteLibrary.RESULT_OK && currentTag.TagMain.Active == WasteLibrary.STATU_ACTIVE {
						var containerStatu string = WasteLibrary.CONTAINER_FULLNESS_STATU_FULL
						second := time.Since(WasteLibrary.StringToTime(currentTag.TagReader.ReadTime)).Seconds()
						if second < 25*60*60 {
							containerStatu = WasteLibrary.CONTAINER_FULLNESS_STATU_EMPTY
						}
						//TO DO
						//vardiya config
						//for _, device := range customerDevicesList.Devices {
						//	if device.DeviceId == currentTag.TagMain.DeviceId {
						//
						// var readTime time.Time      // tag read time
						// var work1Start time.Time    //device work1 start
						// var work1End time.Time      //device work1 start add 1 hour
						// var work1Exist bool = false //if device work1 exist
						// var work2Start time.Time    //device work1 start
						// var work2End time.Time      //device work1 start
						// var work2Exist bool = false //if device work1 exist
						// var work3Start time.Time    //device work1 start
						// var work3End time.Time      //device work1 start
						// var work3Exist bool = false //if device work1 exist
						//
						// if device.DeviceWorkHour.WorkHour1Add > 0 {
						// work1Exist = true
						// work1Start = WasteLibrary.StringToTime(device.DeviceWorkHour.WorkHour1Start)
						// work1Start.
						//
						// }
						//
						// var containerStatu string = WasteLibrary.CONTAINER_FULLNESS_STATU_FULL
						// if work1Exist && readTime > work1Start && readTime < work1End {
						// containerStatu = WasteLibrary.CONTAINER_FULLNESS_STATU_EMPTY
						// }
						// if work2Exist && readTime > work2Start && readTime < work2End {
						// containerStatu = WasteLibrary.CONTAINER_FULLNESS_STATU_EMPTY
						// }
						// if work3Exist && readTime > work3Start && readTime < work3End {
						// containerStatu = WasteLibrary.CONTAINER_FULLNESS_STATU_EMPTY
						// }

						//break
						//}
						//}
						if containerStatu != currentTag.TagStatu.ContainerStatu {
							currentTag.TagStatu.ContainerStatu = containerStatu
							data := url.Values{
								WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
								WasteLibrary.HTTP_DATA:   {currentTag.TagStatu.ToString()},
							}
							resultVal = WasteLibrary.SaveStaticDbMainForStoreApi(data)
							if resultVal.Result != WasteLibrary.RESULT_OK {
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
								WasteLibrary.LogStr(resultVal.ToString())
								continue
							}
							currentTag.TagStatu = WasteLibrary.StringToTagStatuType(resultVal.Retval.(string))
							resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_TAG_STATUS, currentTag.TagStatu.ToIdString(), currentTag.TagStatu.ToString())
							if resultVal.Result != WasteLibrary.RESULT_OK {
								WasteLibrary.LogStr(resultVal.ToString())
								continue
							}
							WasteLibrary.PublishRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_CHANNEL+currentHttpHeader.ToCustomerIdString(), WasteLibrary.DATATYPE_TAG_STATU, currentTag.TagStatu.ToString())
						}

					}
				}
			}

			//TO DO
			//take tag statu spanshot
		}

		time.Sleep(opInterval * time.Second)

	}
}
