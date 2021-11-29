package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/devafatek/WasteLibrary"
)

var opInterval time.Duration = 60 * 60

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

	var currentHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.StringToHttpClientHeaderType(req.FormValue(WasteLibrary.HTTP_HEADER))
	if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_RFID {
		var customerDevices WasteLibrary.CustomerRfidDevicesType = WasteLibrary.StringToCustomerTagsType(req.FormValue(WasteLibrary.HTTP_DATA))
		go customerProcRfid(customerDevices)
	} else if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_RECY {
		var customerDevices WasteLibrary.CustomerRecyDevicesType = WasteLibrary.StringToCustomerTagsType(req.FormValue(WasteLibrary.HTTP_DATA))
		go customerProcRecy(customerDevices)
	} else if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_ULT {
		var customerDevices WasteLibrary.CustomerUltDevicesType = WasteLibrary.StringToCustomerTagsType(req.FormValue(WasteLibrary.HTTP_DATA))
		go customerProcUlt(customerDevices)
	} else {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_HTTP_PARSE
		w.Write(resultVal.ToByte())

		return
	}
	w.Write(resultVal.ToByte())
}

func customerProcRfid(customerDevices WasteLibrary.CustomerRfidDevicesType) {
	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL

	var customerDevicesList WasteLibrary.CustomerRfidDevicesViewListType
	customerDevicesList.New()
	customerDevicesList.CustomerId = customerDevices.CustomerId
	for _, deviceId := range customerDevices.Devices {

		if deviceId != 0 {

			var currentDevice WasteLibrary.RfidDeviceType
			currentDevice.New()
			currentDevice.DeviceId = deviceId
			resultVal = currentDevice.GetByRedis()
			if resultVal.Result == WasteLibrary.RESULT_OK && currentDevice.DeviceMain.Active == WasteLibrary.STATU_ACTIVE {
				var currentViewDevice WasteLibrary.RfidDeviceViewType
				currentViewDevice.New()
				currentViewDevice.DeviceId = currentDevice.DeviceId
				currentViewDevice.PlateNo = currentDevice.PlateNo
				currentViewDevice.Latitude = currentDevice.Latitude
				currentViewDevice.Longitude = currentDevice.Latitude
				customerDevicesList.Tags[currentViewDevice.ToIdString()] = currentViewDevice
			}
		}
	}
	resultVal = customerDevicesList.SaveToRedis()
	resultVal = customerDevicesList.SaveToRedisWODb()

}

func customerProcRecy(customerDevices WasteLibrary.CustomerRecyDevicesType) {
	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL

	var customerDevicesList WasteLibrary.CustomerRecyDevicesViewListType
	customerDevicesList.New()
	customerDevicesList.CustomerId = customerDevices.CustomerId
	for _, deviceId := range customerDevices.Devices {

		if deviceId != 0 {

			var currentDevice WasteLibrary.RecyDeviceType
			currentDevice.New()
			currentDevice.DeviceId = deviceId
			resultVal = currentDevice.GetByRedis()
			if resultVal.Result == WasteLibrary.RESULT_OK && currentDevice.DeviceMain.Active == WasteLibrary.STATU_ACTIVE {
				var currentViewDevice WasteLibrary.RecyDeviceViewType
				currentViewDevice.New()
				currentViewDevice.DeviceId = currentDevice.DeviceId
				currentViewDevice.PlateNo = currentDevice.PlateNo
				if time.Since(currentDevice.DeviceGps.GpsTime) < 15*60 {
					currentViewDevice.Latitude = currentDevice.DeviceGps.Latitude
					currentViewDevice.Longitude = currentDevice.DeviceGps.Latitude
				} else {
					if time.Since(currentDevice.DeviceEmbededGps.GpsTime) < 5*60 {
						currentViewDevice.Latitude = currentDevice.DeviceEmbededGps.Latitude
						currentViewDevice.Longitude = currentDevice.DeviceEmbededGps.Latitude
					} else {
						currentViewDevice.Latitude = currentDevice.DeviceGps.Latitude
						currentViewDevice.Longitude = currentDevice.DeviceGps.Latitude
					}
				}
				customerDevicesList.Tags[currentViewDevice.ToIdString()] = currentViewDevice
			}
		}
	}
	resultVal = customerDevicesList.SaveToRedis()
	resultVal = customerDevicesList.SaveToRedisWODb()

}

func customerProcUlt(customerDevices WasteLibrary.CustomerUltDevicesType) {
	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL

	var customerDevicesList WasteLibrary.CustomerUltDevicesViewListType
	customerDevicesList.New()
	customerDevicesList.CustomerId = customerDevices.CustomerId
	for _, deviceId := range customerDevices.Devices {

		if deviceId != 0 {

			var currentDevice WasteLibrary.UltDeviceType
			currentDevice.New()
			currentDevice.DeviceId = deviceId
			resultVal = currentDevice.GetByRedis()
			if resultVal.Result == WasteLibrary.RESULT_OK && currentDevice.DeviceMain.Active == WasteLibrary.STATU_ACTIVE {
				var currentViewDevice WasteLibrary.UltDeviceViewType
				currentViewDevice.New()
				currentViewDevice.DeviceId = currentDevice.DeviceId
				if currentDevice.DeviceBase.ContainerNo == "" {
					currentViewDevice.ContainerNo = fmt.Sprintf("%05d", int(currentDevice.DeviceId))
				} else {
					currentViewDevice.ContainerNo = currentDevice.DeviceBase.ContainerNo
				}
				currentViewDevice.ContainerStatu = currentDevice.DeviceStatu.ContainerStatu
				currentViewDevice.UltStatus = currentDevice.DeviceStatu.UltStatus
				currentViewDevice.Latitude = currentDevice.DeviceGps.Latitude
				currentViewDevice.Longitude = currentDevice.DeviceGps.Longitude
				currentViewDevice.SensPercent = currentDevice.DeviceStatu.SensPercent
				customerDevicesList.Tags[currentViewDevice.ToIdString()] = currentViewDevice
			}
		}
	}
	resultVal = customerDevicesList.SaveToRedis()
	resultVal = customerDevicesList.SaveToRedisWODb()

}
