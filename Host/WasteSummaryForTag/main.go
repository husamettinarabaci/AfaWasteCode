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

		workStartTime := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), customerAdminConfig.WorkStartHour, customerAdminConfig.WorkStartMinute, 0, 0, time.Now().Location())
		workEndTime := workStartTime.Add(time.Duration(customerAdminConfig.WorkAddMinute) * time.Minute)
		workYestStartTime := workStartTime.Add(-24 * time.Hour)

		var opLimitTime time.Time
		if customerAdminConfig.DeviceBaseWork == WasteLibrary.STATU_PASSIVE {
			if time.Since(workEndTime).Seconds() > 0 {
				opLimitTime = workStartTime
			} else {
				opLimitTime = workYestStartTime
			}
		}

		var currentHttpHeader WasteLibrary.HttpClientHeaderType
		currentHttpHeader.New()
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

								if currentDeviceWorkHour.WorkCount == 3 {

									workStartTime = time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), currentDeviceWorkHour.Work3StartHour, currentDeviceWorkHour.Work3StartMinute, 0, 0, time.Now().Location())
									workEndTime = workStartTime.Add(time.Duration(currentDeviceWorkHour.Work3AddMinute) * time.Minute)
									workYestStartTime = workStartTime.Add(-24 * time.Hour)

									if time.Since(workEndTime).Seconds() > 0 {
										opLimitTime = workStartTime
									} else {
										opLimitTime = workYestStartTime
									}
								} else if currentDeviceWorkHour.WorkCount == 2 {

									workStartTime = time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), currentDeviceWorkHour.Work2StartHour, currentDeviceWorkHour.Work2StartMinute, 0, 0, time.Now().Location())
									workEndTime = workStartTime.Add(time.Duration(currentDeviceWorkHour.Work2AddMinute) * time.Minute)
									workYestStartTime = workStartTime.Add(-24 * time.Hour)

									if time.Since(workEndTime).Seconds() > 0 {
										opLimitTime = workStartTime
									} else {
										opLimitTime = workYestStartTime
									}
								} else if currentDeviceWorkHour.WorkCount == 1 {

									workStartTime = time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), currentDeviceWorkHour.Work1StartHour, currentDeviceWorkHour.Work1StartMinute, 0, 0, time.Now().Location())
									workEndTime = workStartTime.Add(time.Duration(currentDeviceWorkHour.Work1AddMinute) * time.Minute)
									workYestStartTime = workStartTime.Add(-24 * time.Hour)

									if time.Since(workEndTime).Seconds() > 0 {
										opLimitTime = workStartTime
									} else {
										opLimitTime = workYestStartTime
									}
								}

							}

						}

						var containerStatu = currentTag.TagStatu.ContainerStatu
						second := opLimitTime.Sub(WasteLibrary.StringToTime(currentTag.TagReader.ReadTime)).Seconds()
						if second < 0 {
							containerStatu = WasteLibrary.CONTAINER_FULLNESS_STATU_EMPTY
						} else {
							containerStatu = WasteLibrary.CONTAINER_FULLNESS_STATU_FULL
						}

						if containerStatu != currentTag.TagStatu.ContainerStatu {
							currentTag.TagStatu.ContainerStatu = containerStatu

							resultVal = currentTag.TagStatu.SaveToDb()
							if resultVal.Result != WasteLibrary.RESULT_OK {
								continue
							}

							resultVal = currentTag.TagStatu.SaveToRedis()
							if resultVal.Result != WasteLibrary.RESULT_OK {
								continue
							}
						}

					}
				}
			}

			data := url.Values{
				WasteLibrary.HTTP_DATA: {customerTags.ToString()},
			}

			resultVal = WasteLibrary.HttpPostReq("http://waste-summaryfortagview-cluster-ip/reader", data)
		}

		time.Sleep(opInterval * time.Second)

	}
}
