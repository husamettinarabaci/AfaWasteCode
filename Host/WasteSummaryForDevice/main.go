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
						go customerProcRfid(customerId)
						go customerProcRecy(customerId)
						go customerProcUlt(customerId)
					}
				}
			}
		}
		time.Sleep(opInterval * time.Second)
	}
}

func customerProcRfid(customerId float64) {
	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	var currentHttpHeader WasteLibrary.HttpClientHeaderType
	currentHttpHeader.New()
	var customerDevices WasteLibrary.CustomerRfidDevicesType
	customerDevices.CustomerId = customerId
	resultVal = customerDevices.GetByRedis("0")
	if resultVal.Result == WasteLibrary.RESULT_OK {

		var currentHttpHeader WasteLibrary.HttpClientHeaderType
		currentHttpHeader.New()
		currentHttpHeader.DeviceType = WasteLibrary.DEVICETYPE_RFID
		data := url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {customerDevices.ToString()},
		}
		//TO DO
		//summary for device

		resultVal = WasteLibrary.HttpPostReq("http://waste-summaryfordeviceview-cluster-ip/reader", data)
	}

}

func customerProcRecy(customerId float64) {
	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL

	var currentHttpHeader WasteLibrary.HttpClientHeaderType
	currentHttpHeader.New()
	var customerDevices WasteLibrary.CustomerRecyDevicesType
	customerDevices.CustomerId = customerId
	resultVal = customerDevices.GetByRedis("0")
	if resultVal.Result == WasteLibrary.RESULT_OK {

		var currentHttpHeader WasteLibrary.HttpClientHeaderType
		currentHttpHeader.New()
		currentHttpHeader.DeviceType = WasteLibrary.DEVICETYPE_RECY
		data := url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {customerDevices.ToString()},
		}
		//TO DO
		//summary for device

		resultVal = WasteLibrary.HttpPostReq("http://waste-summaryfordeviceview-cluster-ip/reader", data)
	}

}

func customerProcUlt(customerId float64) {
	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL

	var currentHttpHeader WasteLibrary.HttpClientHeaderType
	currentHttpHeader.New()
	var customerDevices WasteLibrary.CustomerUltDevicesType
	customerDevices.CustomerId = customerId
	resultVal = customerDevices.GetByRedis("0")
	if resultVal.Result == WasteLibrary.RESULT_OK {

		var currentHttpHeader WasteLibrary.HttpClientHeaderType
		currentHttpHeader.New()
		currentHttpHeader.DeviceType = WasteLibrary.DEVICETYPE_ULT
		data := url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {customerDevices.ToString()},
		}
		//TO DO
		//summary for device

		resultVal = WasteLibrary.HttpPostReq("http://waste-summaryfordeviceview-cluster-ip/reader", data)
	}

}
