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

		var currentCustomers WasteLibrary.CustomersType
		resultVal = currentCustomers.GetByRedis()
		if resultVal.Result == WasteLibrary.RESULT_OK {

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

		var customerAdminConfig WasteLibrary.AdminConfigType
		customerAdminConfig.New()
		customerAdminConfig.CustomerId = customerId
		resultVal = customerAdminConfig.GetByRedis()

		var workEndHour int = customerAdminConfig.WorkEndHour
		var workEndMinute int = customerAdminConfig.WorkEndMinute

		var inWork bool = false
		var workStartTime time.Time
		if customerAdminConfig.DeviceBaseWork == WasteLibrary.STATU_PASSIVE {
			if time.Now().Hour() < workEndHour {
				inWork = true
			} else if time.Now().Hour() == workEndHour {
				if time.Now().Minute() < workEndMinute {
					inWork = true
				} else {
					inWork = false
				}
			} else {
				inWork = false
			}

			if !inWork {
				workStartTime = time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), customerAdminConfig.WorkStartHour, customerAdminConfig.WorkStartMinute, 0, 0, time.Now().Location())
			}
		}

		if inWork {

		} else {
			var currentHttpHeader WasteLibrary.HttpClientHeaderType
			currentHttpHeader.New()
			currentHttpHeader.DataType = WasteLibrary.DATATYPE_TAG_STATU
			var customerTags WasteLibrary.CustomerTagsType
			customerTags.CustomerId = customerId
			resultVal = customerTags.GetByRedis()
			if resultVal.Result == WasteLibrary.RESULT_OK {

				for _, tagId := range customerTags.Tags {

					if tagId != 0 {

						var currentTag WasteLibrary.TagType
						currentTag.New()
						currentTag.TagId = tagId
						resultVal = currentTag.GetByRedis()
						if resultVal.Result == WasteLibrary.RESULT_OK && currentTag.TagMain.Active == WasteLibrary.STATU_ACTIVE {

							if customerAdminConfig.DeviceBaseWork == WasteLibrary.STATU_ACTIVE {

								var currentDeviceWorkHour WasteLibrary.RfidDeviceWorkHourType
								currentDeviceWorkHour.DeviceId = currentTag.TagMain.DeviceId
								resultVal = currentDeviceWorkHour.GetByRedis()
								if resultVal.Result == WasteLibrary.RESULT_OK {

									workEndHour = currentDeviceWorkHour.WorkEndHour
									workEndMinute = currentDeviceWorkHour.WorkEndMinute
									if time.Now().Hour() < workEndHour {
										inWork = true
									} else if time.Now().Hour() == workEndHour {
										if time.Now().Minute() < workEndMinute {
											inWork = true
										} else {
											inWork = false
										}
									} else {
										inWork = false
									}

									if !inWork {
										workStartTime = time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), currentDeviceWorkHour.WorkStartHour, currentDeviceWorkHour.WorkStartMinute, 0, 0, time.Now().Location())
									}
								}

							}

							if inWork {

							} else {

								var containerStatu = currentTag.TagStatu.ContainerStatu
								second := workStartTime.Sub(WasteLibrary.StringToTime(currentTag.TagReader.ReadTime)).Seconds()
								if second < 0 {
									containerStatu = WasteLibrary.CONTAINER_FULLNESS_STATU_EMPTY
								} else {
									containerStatu = WasteLibrary.CONTAINER_FULLNESS_STATU_FULL
								}

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

									resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_TAG_STATUS, currentTag.TagStatu.ToIdString(), currentTag.TagStatu.ToString())
									if resultVal.Result != WasteLibrary.RESULT_OK {
										WasteLibrary.LogStr(resultVal.ToString())
										continue
									}
								}
							}

						}
					}
				}

				//TO DO
				//take tag statu spanshot
			}
		}
		time.Sleep(opInterval * time.Second)

	}
}
