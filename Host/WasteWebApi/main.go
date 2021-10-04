package main

import (
	"net/http"
	"net/url"
	"time"

	"github.com/devafatek/WasteLibrary"
)

var testapp bool = true

func initStart() {

	WasteLibrary.LogStr("Successfully connected!")
}
func main() {

	initStart()

	http.HandleFunc("/health", WasteLibrary.HealthHandler)
	http.HandleFunc("/readiness", WasteLibrary.ReadinessHandler)
	http.HandleFunc("/status", WasteLibrary.StatusHandler)
	http.HandleFunc("/getCustomer", getCustomer)
	http.HandleFunc("/getDevice", getDevice)
	http.HandleFunc("/getDevices", getDevices)
	http.HandleFunc("/getConfig", getConfig)
	http.HandleFunc("/getTags", getTags)
	http.HandleFunc("/getTag", getTag)
	http.HandleFunc("/getLink", getLink)
	http.ListenAndServe(":80", nil)
}

func getLink(w http.ResponseWriter, req *http.Request) {

	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL

	var linkVal string = req.Host
	WasteLibrary.LogStr("Get Link : " + linkVal)
	w.Write(resultVal.ToByte())
}

func getCustomer(w http.ResponseWriter, req *http.Request) {

	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	if err := req.ParseForm(); err != nil {
		WasteLibrary.LogErr(err)
		return
	}
	resultVal = checkAuth(req.Form)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		w.Write(resultVal.ToByte())
		return
	}
	if testapp {

		var retVal WasteLibrary.CustomerType = WasteLibrary.CustomerType{
			CustomerId:   1,
			CustomerName: "Afatek",
			CustomerLink: "afatek.aws.afatek.com.tr",
			RfIdApp:      WasteLibrary.STATU_ACTIVE,
			UltApp:       WasteLibrary.STATU_ACTIVE,
			RecyApp:      WasteLibrary.STATU_ACTIVE,
			Active:       WasteLibrary.STATU_ACTIVE,
			CreateTime:   time.Now().String(),
		}
		resultVal.Result = WasteLibrary.RESULT_OK
		resultVal.Retval = retVal.ToString()

		w.Write(resultVal.ToByte())
		return
	}

	resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_LINK, req.Host)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		w.Write(resultVal.ToByte())
		return
	}
	var customerId string = resultVal.Retval.(string)

	resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMERS, customerId)

	w.Write(resultVal.ToByte())
}

func getDevice(w http.ResponseWriter, req *http.Request) {

	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	if err := req.ParseForm(); err != nil {
		WasteLibrary.LogErr(err)
		return
	}
	resultVal = checkAuth(req.Form)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		w.Write(resultVal.ToByte())
		return
	}
	if testapp {
		var currentData WasteLibrary.DeviceType = WasteLibrary.StringToDeviceType(req.FormValue(WasteLibrary.HTTP_DATA))
		if currentData.DeviceId == 1 {
			var retVal WasteLibrary.DeviceType = WasteLibrary.DeviceType{
				DeviceId:              1,
				CustomerId:            1,
				DeviceName:            "06 AFA 001",
				DeviceType:            WasteLibrary.APPTYPE_RFID,
				SerialNumber:          "000001",
				ReaderAppStatus:       WasteLibrary.STATU_ACTIVE,
				ReaderAppLastOkTime:   WasteLibrary.STATU_ACTIVE,
				ReaderConnStatus:      WasteLibrary.STATU_ACTIVE,
				ReaderConnLastOkTime:  time.Now().String(),
				ReaderStatus:          WasteLibrary.STATU_ACTIVE,
				ReaderLastOkTime:      time.Now().String(),
				CamAppStatus:          WasteLibrary.STATU_ACTIVE,
				CamAppLastOkTime:      time.Now().String(),
				CamConnStatus:         WasteLibrary.STATU_ACTIVE,
				CamConnLastOkTime:     time.Now().String(),
				CamStatus:             WasteLibrary.STATU_ACTIVE,
				CamLastOkTime:         time.Now().String(),
				GpsAppStatus:          WasteLibrary.STATU_ACTIVE,
				GpsAppLastOkTime:      time.Now().String(),
				GpsConnStatus:         WasteLibrary.STATU_ACTIVE,
				GpsConnLastOkTime:     time.Now().String(),
				GpsStatus:             WasteLibrary.STATU_ACTIVE,
				GpsLastOkTime:         time.Now().String(),
				ThermAppStatus:        WasteLibrary.STATU_ACTIVE,
				ThermAppLastOkTime:    time.Now().String(),
				TransferAppStatus:     WasteLibrary.STATU_ACTIVE,
				TransferAppLastOkTime: time.Now().String(),
				AliveStatus:           WasteLibrary.STATU_ACTIVE,
				AliveLastOkTime:       time.Now().String(),
				ContactStatus:         WasteLibrary.STATU_ACTIVE,
				ContactLastOkTime:     time.Now().String(),
				Therm:                 "39",
				Latitude:              37.03633,
				Longitude:             27.41585,
				Speed:                 0,
				UltRange:              4785,
				UltStatus:             WasteLibrary.STATU_ACTIVE,
				DeviceStatus:          WasteLibrary.STATU_ACTIVE,
				AlarmStatus:           WasteLibrary.STATU_ACTIVE,
				TotalGlassCount:       345,
				TotalPlasticCount:     567,
				TotalMetalCount:       890,
				UltTime:               time.Now().String(),
				AlarmTime:             time.Now().String(),
				AlarmType:             WasteLibrary.ALARM_ATMP,
				Alarm:                 "Therm : 85",
				RecyTime:              time.Now().String(),
				Active:                WasteLibrary.STATU_ACTIVE,
				ThermTime:             time.Now().String(),
				GpsTime:               time.Now().String(),
				StatusTime:            time.Now().String(),
				CreateTime:            time.Now().String(),
			}
			resultVal.Result = WasteLibrary.RESULT_OK
			resultVal.Retval = retVal.ToString()
		} else if currentData.DeviceId == 2 {
			var retVal WasteLibrary.DeviceType = WasteLibrary.DeviceType{
				DeviceId:              2,
				CustomerId:            1,
				DeviceName:            "000002",
				DeviceType:            WasteLibrary.APPTYPE_ULT,
				SerialNumber:          "000002",
				ReaderAppStatus:       WasteLibrary.STATU_ACTIVE,
				ReaderAppLastOkTime:   WasteLibrary.STATU_ACTIVE,
				ReaderConnStatus:      WasteLibrary.STATU_ACTIVE,
				ReaderConnLastOkTime:  time.Now().String(),
				ReaderStatus:          WasteLibrary.STATU_ACTIVE,
				ReaderLastOkTime:      time.Now().String(),
				CamAppStatus:          WasteLibrary.STATU_ACTIVE,
				CamAppLastOkTime:      time.Now().String(),
				CamConnStatus:         WasteLibrary.STATU_ACTIVE,
				CamConnLastOkTime:     time.Now().String(),
				CamStatus:             WasteLibrary.STATU_ACTIVE,
				CamLastOkTime:         time.Now().String(),
				GpsAppStatus:          WasteLibrary.STATU_ACTIVE,
				GpsAppLastOkTime:      time.Now().String(),
				GpsConnStatus:         WasteLibrary.STATU_ACTIVE,
				GpsConnLastOkTime:     time.Now().String(),
				GpsStatus:             WasteLibrary.STATU_ACTIVE,
				GpsLastOkTime:         time.Now().String(),
				ThermAppStatus:        WasteLibrary.STATU_ACTIVE,
				ThermAppLastOkTime:    time.Now().String(),
				TransferAppStatus:     WasteLibrary.STATU_ACTIVE,
				TransferAppLastOkTime: time.Now().String(),
				AliveStatus:           WasteLibrary.STATU_ACTIVE,
				AliveLastOkTime:       time.Now().String(),
				ContactStatus:         WasteLibrary.STATU_ACTIVE,
				ContactLastOkTime:     time.Now().String(),
				Therm:                 "39",
				Latitude:              37.03652,
				Longitude:             27.42111,
				Speed:                 0,
				UltRange:              4785,
				UltStatus:             WasteLibrary.STATU_ACTIVE,
				DeviceStatus:          WasteLibrary.STATU_ACTIVE,
				AlarmStatus:           WasteLibrary.STATU_ACTIVE,
				TotalGlassCount:       345,
				TotalPlasticCount:     567,
				TotalMetalCount:       890,
				UltTime:               time.Now().String(),
				AlarmTime:             time.Now().String(),
				AlarmType:             WasteLibrary.ALARM_AGPS,
				Alarm:                 "Therm : 85",
				RecyTime:              time.Now().String(),
				Active:                WasteLibrary.STATU_ACTIVE,
				ThermTime:             time.Now().String(),
				GpsTime:               time.Now().String(),
				StatusTime:            time.Now().String(),
				CreateTime:            time.Now().String(),
			}
			resultVal.Result = WasteLibrary.RESULT_OK
			resultVal.Retval = retVal.ToString()
		} else if currentData.DeviceId == 3 {
			var retVal WasteLibrary.DeviceType = WasteLibrary.DeviceType{
				DeviceId:              3,
				CustomerId:            1,
				DeviceName:            "000003",
				DeviceType:            WasteLibrary.APPTYPE_RECY,
				SerialNumber:          "000003",
				ReaderAppStatus:       WasteLibrary.STATU_ACTIVE,
				ReaderAppLastOkTime:   WasteLibrary.STATU_ACTIVE,
				ReaderConnStatus:      WasteLibrary.STATU_ACTIVE,
				ReaderConnLastOkTime:  time.Now().String(),
				ReaderStatus:          WasteLibrary.STATU_ACTIVE,
				ReaderLastOkTime:      time.Now().String(),
				CamAppStatus:          WasteLibrary.STATU_ACTIVE,
				CamAppLastOkTime:      time.Now().String(),
				CamConnStatus:         WasteLibrary.STATU_ACTIVE,
				CamConnLastOkTime:     time.Now().String(),
				CamStatus:             WasteLibrary.STATU_ACTIVE,
				CamLastOkTime:         time.Now().String(),
				GpsAppStatus:          WasteLibrary.STATU_ACTIVE,
				GpsAppLastOkTime:      time.Now().String(),
				GpsConnStatus:         WasteLibrary.STATU_ACTIVE,
				GpsConnLastOkTime:     time.Now().String(),
				GpsStatus:             WasteLibrary.STATU_ACTIVE,
				GpsLastOkTime:         time.Now().String(),
				ThermAppStatus:        WasteLibrary.STATU_ACTIVE,
				ThermAppLastOkTime:    time.Now().String(),
				TransferAppStatus:     WasteLibrary.STATU_ACTIVE,
				TransferAppLastOkTime: time.Now().String(),
				AliveStatus:           WasteLibrary.STATU_ACTIVE,
				AliveLastOkTime:       time.Now().String(),
				ContactStatus:         WasteLibrary.STATU_ACTIVE,
				ContactLastOkTime:     time.Now().String(),
				Therm:                 "39",
				Latitude:              37.03732,
				Longitude:             27.41609,
				Speed:                 0,
				UltRange:              4785,
				UltStatus:             WasteLibrary.STATU_ACTIVE,
				DeviceStatus:          WasteLibrary.STATU_ACTIVE,
				AlarmStatus:           WasteLibrary.STATU_ACTIVE,
				TotalGlassCount:       345,
				TotalPlasticCount:     567,
				TotalMetalCount:       890,
				UltTime:               time.Now().String(),
				AlarmTime:             time.Now().String(),
				AlarmType:             WasteLibrary.ALARM_NONE,
				Alarm:                 "Therm : 85",
				RecyTime:              time.Now().String(),
				Active:                WasteLibrary.STATU_ACTIVE,
				ThermTime:             time.Now().String(),
				GpsTime:               time.Now().String(),
				StatusTime:            time.Now().String(),
				CreateTime:            time.Now().String(),
			}
			resultVal.Result = WasteLibrary.RESULT_OK
			resultVal.Retval = retVal.ToString()
		} else {
		}
		w.Write(resultVal.ToByte())

		return
	}

	resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_LINK, req.Host)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		w.Write(resultVal.ToByte())
		return
	}
	var customerId string = resultVal.Retval.(string)

	var currentData WasteLibrary.DeviceType = WasteLibrary.StringToDeviceType(req.FormValue(WasteLibrary.HTTP_DATA))
	resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_DEVICES, currentData.ToIdString())

	currentData = WasteLibrary.StringToDeviceType(resultVal.Retval.(string))
	if currentData.ToCustomerIdString() == customerId {
		w.Write(resultVal.ToByte())
	} else {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = ""
		w.Write(resultVal.ToByte())
	}
}

func getDevices(w http.ResponseWriter, req *http.Request) {

	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	if err := req.ParseForm(); err != nil {
		WasteLibrary.LogErr(err)
		return
	}
	resultVal = checkAuth(req.Form)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		w.Write(resultVal.ToByte())
		return
	}

	if testapp {

		var retVal WasteLibrary.CustomerDevicesListType = WasteLibrary.CustomerDevicesListType{
			CustomerId: 1,
			Devices:    make(map[string]WasteLibrary.DeviceType),
		}

		var retVal1 WasteLibrary.DeviceType = WasteLibrary.DeviceType{
			DeviceId:              1,
			CustomerId:            1,
			DeviceName:            "06 AFA 001",
			DeviceType:            WasteLibrary.APPTYPE_RFID,
			SerialNumber:          "000001",
			ReaderAppStatus:       WasteLibrary.STATU_ACTIVE,
			ReaderAppLastOkTime:   WasteLibrary.STATU_ACTIVE,
			ReaderConnStatus:      WasteLibrary.STATU_ACTIVE,
			ReaderConnLastOkTime:  time.Now().String(),
			ReaderStatus:          WasteLibrary.STATU_ACTIVE,
			ReaderLastOkTime:      time.Now().String(),
			CamAppStatus:          WasteLibrary.STATU_ACTIVE,
			CamAppLastOkTime:      time.Now().String(),
			CamConnStatus:         WasteLibrary.STATU_ACTIVE,
			CamConnLastOkTime:     time.Now().String(),
			CamStatus:             WasteLibrary.STATU_ACTIVE,
			CamLastOkTime:         time.Now().String(),
			GpsAppStatus:          WasteLibrary.STATU_ACTIVE,
			GpsAppLastOkTime:      time.Now().String(),
			GpsConnStatus:         WasteLibrary.STATU_ACTIVE,
			GpsConnLastOkTime:     time.Now().String(),
			GpsStatus:             WasteLibrary.STATU_ACTIVE,
			GpsLastOkTime:         time.Now().String(),
			ThermAppStatus:        WasteLibrary.STATU_ACTIVE,
			ThermAppLastOkTime:    time.Now().String(),
			TransferAppStatus:     WasteLibrary.STATU_ACTIVE,
			TransferAppLastOkTime: time.Now().String(),
			AliveStatus:           WasteLibrary.STATU_ACTIVE,
			AliveLastOkTime:       time.Now().String(),
			ContactStatus:         WasteLibrary.STATU_ACTIVE,
			ContactLastOkTime:     time.Now().String(),
			Therm:                 "39",
			Latitude:              37.03633,
			Longitude:             27.41585,
			Speed:                 0,
			UltRange:              4785,
			UltStatus:             WasteLibrary.STATU_ACTIVE,
			DeviceStatus:          WasteLibrary.STATU_ACTIVE,
			AlarmStatus:           WasteLibrary.STATU_ACTIVE,
			TotalGlassCount:       345,
			TotalPlasticCount:     567,
			TotalMetalCount:       890,
			UltTime:               time.Now().String(),
			AlarmTime:             time.Now().String(),
			AlarmType:             WasteLibrary.ALARM_ATMP,
			Alarm:                 "Therm : 85",
			RecyTime:              time.Now().String(),
			Active:                WasteLibrary.STATU_ACTIVE,
			ThermTime:             time.Now().String(),
			GpsTime:               time.Now().String(),
			StatusTime:            time.Now().String(),
			CreateTime:            time.Now().String(),
		}

		var retVal2 WasteLibrary.DeviceType = WasteLibrary.DeviceType{
			DeviceId:              2,
			CustomerId:            1,
			DeviceName:            "000002",
			DeviceType:            WasteLibrary.APPTYPE_ULT,
			SerialNumber:          "000002",
			ReaderAppStatus:       WasteLibrary.STATU_ACTIVE,
			ReaderAppLastOkTime:   WasteLibrary.STATU_ACTIVE,
			ReaderConnStatus:      WasteLibrary.STATU_ACTIVE,
			ReaderConnLastOkTime:  time.Now().String(),
			ReaderStatus:          WasteLibrary.STATU_ACTIVE,
			ReaderLastOkTime:      time.Now().String(),
			CamAppStatus:          WasteLibrary.STATU_ACTIVE,
			CamAppLastOkTime:      time.Now().String(),
			CamConnStatus:         WasteLibrary.STATU_ACTIVE,
			CamConnLastOkTime:     time.Now().String(),
			CamStatus:             WasteLibrary.STATU_ACTIVE,
			CamLastOkTime:         time.Now().String(),
			GpsAppStatus:          WasteLibrary.STATU_ACTIVE,
			GpsAppLastOkTime:      time.Now().String(),
			GpsConnStatus:         WasteLibrary.STATU_ACTIVE,
			GpsConnLastOkTime:     time.Now().String(),
			GpsStatus:             WasteLibrary.STATU_ACTIVE,
			GpsLastOkTime:         time.Now().String(),
			ThermAppStatus:        WasteLibrary.STATU_ACTIVE,
			ThermAppLastOkTime:    time.Now().String(),
			TransferAppStatus:     WasteLibrary.STATU_ACTIVE,
			TransferAppLastOkTime: time.Now().String(),
			AliveStatus:           WasteLibrary.STATU_ACTIVE,
			AliveLastOkTime:       time.Now().String(),
			ContactStatus:         WasteLibrary.STATU_ACTIVE,
			ContactLastOkTime:     time.Now().String(),
			Therm:                 "39",
			Latitude:              37.03652,
			Longitude:             27.42111,
			Speed:                 0,
			UltRange:              4785,
			UltStatus:             WasteLibrary.STATU_ACTIVE,
			DeviceStatus:          WasteLibrary.STATU_ACTIVE,
			AlarmStatus:           WasteLibrary.STATU_ACTIVE,
			TotalGlassCount:       345,
			TotalPlasticCount:     567,
			TotalMetalCount:       890,
			UltTime:               time.Now().String(),
			AlarmTime:             time.Now().String(),
			AlarmType:             WasteLibrary.ALARM_NONE,
			Alarm:                 "Therm : 85",
			RecyTime:              time.Now().String(),
			Active:                WasteLibrary.STATU_ACTIVE,
			ThermTime:             time.Now().String(),
			GpsTime:               time.Now().String(),
			StatusTime:            time.Now().String(),
			CreateTime:            time.Now().String(),
		}

		var retVal3 WasteLibrary.DeviceType = WasteLibrary.DeviceType{
			DeviceId:              3,
			CustomerId:            1,
			DeviceName:            "000003",
			DeviceType:            WasteLibrary.APPTYPE_RECY,
			SerialNumber:          "000003",
			ReaderAppStatus:       WasteLibrary.STATU_ACTIVE,
			ReaderAppLastOkTime:   WasteLibrary.STATU_ACTIVE,
			ReaderConnStatus:      WasteLibrary.STATU_ACTIVE,
			ReaderConnLastOkTime:  time.Now().String(),
			ReaderStatus:          WasteLibrary.STATU_ACTIVE,
			ReaderLastOkTime:      time.Now().String(),
			CamAppStatus:          WasteLibrary.STATU_ACTIVE,
			CamAppLastOkTime:      time.Now().String(),
			CamConnStatus:         WasteLibrary.STATU_ACTIVE,
			CamConnLastOkTime:     time.Now().String(),
			CamStatus:             WasteLibrary.STATU_ACTIVE,
			CamLastOkTime:         time.Now().String(),
			GpsAppStatus:          WasteLibrary.STATU_ACTIVE,
			GpsAppLastOkTime:      time.Now().String(),
			GpsConnStatus:         WasteLibrary.STATU_ACTIVE,
			GpsConnLastOkTime:     time.Now().String(),
			GpsStatus:             WasteLibrary.STATU_ACTIVE,
			GpsLastOkTime:         time.Now().String(),
			ThermAppStatus:        WasteLibrary.STATU_ACTIVE,
			ThermAppLastOkTime:    time.Now().String(),
			TransferAppStatus:     WasteLibrary.STATU_ACTIVE,
			TransferAppLastOkTime: time.Now().String(),
			AliveStatus:           WasteLibrary.STATU_ACTIVE,
			AliveLastOkTime:       time.Now().String(),
			ContactStatus:         WasteLibrary.STATU_ACTIVE,
			ContactLastOkTime:     time.Now().String(),
			Therm:                 "39",
			Latitude:              37.03732,
			Longitude:             27.41609,
			Speed:                 0,
			UltRange:              4785,
			UltStatus:             WasteLibrary.STATU_ACTIVE,
			DeviceStatus:          WasteLibrary.STATU_ACTIVE,
			AlarmStatus:           WasteLibrary.STATU_ACTIVE,
			TotalGlassCount:       345,
			TotalPlasticCount:     567,
			TotalMetalCount:       890,
			UltTime:               time.Now().String(),
			AlarmTime:             time.Now().String(),
			AlarmType:             WasteLibrary.ALARM_AGPS,
			Alarm:                 "Therm : 85",
			RecyTime:              time.Now().String(),
			Active:                WasteLibrary.STATU_ACTIVE,
			ThermTime:             time.Now().String(),
			GpsTime:               time.Now().String(),
			StatusTime:            time.Now().String(),
			CreateTime:            time.Now().String(),
		}

		retVal.Devices[retVal1.ToIdString()] = retVal1
		retVal.Devices[retVal2.ToIdString()] = retVal2
		retVal.Devices[retVal3.ToIdString()] = retVal3

		resultVal.Result = WasteLibrary.RESULT_OK
		resultVal.Retval = retVal.ToString()

		w.Write(resultVal.ToByte())

		return
	}

	resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_LINK, req.Host)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		w.Write(resultVal.ToByte())
		return
	}
	var customerId string = resultVal.Retval.(string)

	resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_DEVICES, customerId)

	if resultVal.Result == WasteLibrary.RESULT_OK {

		var customerDevices WasteLibrary.CustomerDevicesType = WasteLibrary.StringToCustomerDevicesType(resultVal.Retval.(string))
		var customerDevicesList WasteLibrary.CustomerDevicesListType = WasteLibrary.CustomerDevicesListType{
			CustomerId: WasteLibrary.StringIdToFloat64(customerId),
			Devices:    make(map[string]WasteLibrary.DeviceType),
		}
		for _, deviceId := range customerDevices.Devices {

			if deviceId != 0 {
				resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_DEVICES, WasteLibrary.Float64IdToString(deviceId))
				if resultVal.Result == WasteLibrary.RESULT_OK {
					var currentDevice WasteLibrary.DeviceType = WasteLibrary.StringToDeviceType(resultVal.Retval.(string))
					customerDevicesList.Devices[currentDevice.ToIdString()] = currentDevice
				}

			}
		}
		resultVal.Result = WasteLibrary.RESULT_OK
		resultVal.Retval = customerDevicesList.ToString()
	}

	w.Write(resultVal.ToByte())
}

func getConfig(w http.ResponseWriter, req *http.Request) {

	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	if err := req.ParseForm(); err != nil {
		WasteLibrary.LogErr(err)
		return
	}
	resultVal = checkAuth(req.Form)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		w.Write(resultVal.ToByte())
		return
	}
	if testapp {

		var currentHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.StringToHttpClientHeaderType(req.FormValue(WasteLibrary.HTTP_HEADER))
		if currentHttpHeader.OpType == WasteLibrary.OPTYPE_CUSTOMERCONFIG {
			var retVal WasteLibrary.CustomerConfigType = WasteLibrary.CustomerConfigType{
				CustomerId:      1,
				ArventoApp:      WasteLibrary.STATU_ACTIVE,
				ArventoUserName: "test1",
				ArventoPin1:     "pin1",
				ArventoPin2:     "pin2",
				SystemProblem:   WasteLibrary.STATU_ACTIVE,
				TruckStopTrace:  WasteLibrary.STATU_ACTIVE,
				Active:          WasteLibrary.STATU_ACTIVE,
				CreateTime:      time.Now().String(),
			}
			resultVal.Result = WasteLibrary.RESULT_OK
			resultVal.Retval = retVal.ToString()
		} else if currentHttpHeader.OpType == WasteLibrary.OPTYPE_ADMINCONFIG {
			var retVal WasteLibrary.AdminConfigType = WasteLibrary.AdminConfigType{
				CustomerId: 1,
				Active:     WasteLibrary.STATU_ACTIVE,
				CreateTime: time.Now().String(),
			}
			resultVal.Result = WasteLibrary.RESULT_OK
			resultVal.Retval = retVal.ToString()
		} else if currentHttpHeader.OpType == WasteLibrary.OPTYPE_LOCALCONFIG {
			var retVal WasteLibrary.LocalConfigType = WasteLibrary.LocalConfigType{
				CustomerId: 1,
				Active:     WasteLibrary.STATU_ACTIVE,
				CreateTime: time.Now().String(),
			}
			resultVal.Result = WasteLibrary.RESULT_OK
			resultVal.Retval = retVal.ToString()
		} else {
			resultVal.Result = WasteLibrary.RESULT_FAIL
		}

		w.Write(resultVal.ToByte())
		return
	}

	resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_LINK, req.Host)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		w.Write(resultVal.ToByte())
		return
	}
	var customerId string = resultVal.Retval.(string)

	var currentHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.StringToHttpClientHeaderType(req.FormValue(WasteLibrary.HTTP_HEADER))
	if currentHttpHeader.OpType == WasteLibrary.OPTYPE_CUSTOMERCONFIG {
		resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_CUSTOMERCONFIG, customerId)
	} else if currentHttpHeader.OpType == WasteLibrary.OPTYPE_ADMINCONFIG {
		resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_ADMINCONFIG, customerId)
	} else if currentHttpHeader.OpType == WasteLibrary.OPTYPE_LOCALCONFIG {
		resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_LOCALCONFIG, customerId)
	} else {
		resultVal.Result = WasteLibrary.RESULT_FAIL
	}

	w.Write(resultVal.ToByte())
}

func getTags(w http.ResponseWriter, req *http.Request) {

	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	if err := req.ParseForm(); err != nil {
		WasteLibrary.LogErr(err)
		return
	}
	resultVal = checkAuth(req.Form)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		w.Write(resultVal.ToByte())
		return
	}

	if testapp {

		var retVal WasteLibrary.CustomerTagsType = WasteLibrary.CustomerTagsType{
			CustomerId: 1,
			Tags:       make(map[string]WasteLibrary.TagType),
		}

		var retVal1 WasteLibrary.TagType = WasteLibrary.TagType{
			TagID:       1,
			CustomerId:  1,
			DeviceId:    1,
			UID:         "03f35539-0eea-11ec-976d-b827ebb1d188",
			Epc:         "00058",
			ContainerNo: "00058",
			Latitude:    37.03780,
			Longitude:   27.41151,
			Statu:       WasteLibrary.STATU_ACTIVE,
			ImageStatu:  WasteLibrary.STATU_ACTIVE,
			Active:      WasteLibrary.STATU_ACTIVE,
			ReadTime:    time.Now().String(),
			CheckTime:   time.Now().String(),
			CreateTime:  time.Now().String(),
		}

		var retVal2 WasteLibrary.TagType = WasteLibrary.TagType{
			TagID:       2,
			CustomerId:  1,
			DeviceId:    1,
			UID:         "992539a1-0ee1-11ec-8fff-b827ebb1d188",
			Epc:         "00059",
			ContainerNo: "00059",
			Latitude:    37.03899,
			Longitude:   27.42267,
			Statu:       WasteLibrary.STATU_ACTIVE,
			ImageStatu:  WasteLibrary.STATU_ACTIVE,
			Active:      WasteLibrary.STATU_ACTIVE,
			ReadTime:    time.Now().String(),
			CheckTime:   time.Now().String(),
			CreateTime:  time.Now().String(),
		}

		var retVal3 WasteLibrary.TagType = WasteLibrary.TagType{
			TagID:       3,
			CustomerId:  1,
			DeviceId:    1,
			UID:         "d3e415a7-0ee0-11ec-853d-b827ebb1d188",
			Epc:         "00060",
			ContainerNo: "00060",
			Latitude:    37.03528,
			Longitude:   27.41040,
			Statu:       "2",
			ImageStatu:  WasteLibrary.STATU_ACTIVE,
			Active:      WasteLibrary.STATU_ACTIVE,
			ReadTime:    time.Now().String(),
			CheckTime:   time.Now().String(),
			CreateTime:  time.Now().String(),
		}

		retVal.Tags[retVal1.ToIdString()] = retVal1
		retVal.Tags[retVal2.ToIdString()] = retVal2
		retVal.Tags[retVal3.ToIdString()] = retVal3

		resultVal.Result = WasteLibrary.RESULT_OK
		resultVal.Retval = retVal.ToString()

		w.Write(resultVal.ToByte())

		return
	}

	resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_LINK, req.Host)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		w.Write(resultVal.ToByte())
		return
	}
	var customerId string = resultVal.Retval.(string)

	resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_TAGS, customerId)

	w.Write(resultVal.ToByte())
}

func getTag(w http.ResponseWriter, req *http.Request) {

	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	if err := req.ParseForm(); err != nil {
		WasteLibrary.LogErr(err)
		return
	}
	resultVal = checkAuth(req.Form)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		w.Write(resultVal.ToByte())
		return
	}

	if testapp {
		var currentData WasteLibrary.TagType = WasteLibrary.StringToTagType(req.FormValue(WasteLibrary.HTTP_DATA))
		if currentData.TagID == 1 {

			var retVal WasteLibrary.TagType = WasteLibrary.TagType{
				TagID:       1,
				CustomerId:  1,
				DeviceId:    1,
				UID:         "03f35539-0eea-11ec-976d-b827ebb1d188",
				Epc:         "00058",
				ContainerNo: "00058",
				Latitude:    37.03780,
				Longitude:   27.41151,
				Statu:       WasteLibrary.STATU_ACTIVE,
				ImageStatu:  WasteLibrary.STATU_ACTIVE,
				Active:      WasteLibrary.STATU_ACTIVE,
				ReadTime:    time.Now().String(),
				CheckTime:   time.Now().String(),
				CreateTime:  time.Now().String(),
			}

			resultVal.Result = WasteLibrary.RESULT_OK
			resultVal.Retval = retVal.ToString()
		} else if currentData.TagID == 2 {

			var retVal WasteLibrary.TagType = WasteLibrary.TagType{
				TagID:       2,
				CustomerId:  1,
				DeviceId:    1,
				UID:         "992539a1-0ee1-11ec-8fff-b827ebb1d188",
				Epc:         "00059",
				ContainerNo: "00059",
				Latitude:    37.03899,
				Longitude:   27.42267,
				Statu:       WasteLibrary.STATU_ACTIVE,
				ImageStatu:  WasteLibrary.STATU_ACTIVE,
				Active:      WasteLibrary.STATU_ACTIVE,
				ReadTime:    time.Now().String(),
				CheckTime:   time.Now().String(),
				CreateTime:  time.Now().String(),
			}

			resultVal.Result = WasteLibrary.RESULT_OK
			resultVal.Retval = retVal.ToString()
		} else if currentData.TagID == 3 {

			var retVal WasteLibrary.TagType = WasteLibrary.TagType{
				TagID:       3,
				CustomerId:  1,
				DeviceId:    1,
				UID:         "d3e415a7-0ee0-11ec-853d-b827ebb1d188",
				Epc:         "00060",
				ContainerNo: "00060",
				Latitude:    37.03528,
				Longitude:   27.41040,
				Statu:       "2",
				ImageStatu:  WasteLibrary.STATU_ACTIVE,
				Active:      WasteLibrary.STATU_ACTIVE,
				ReadTime:    time.Now().String(),
				CheckTime:   time.Now().String(),
				CreateTime:  time.Now().String(),
			}
			resultVal.Result = WasteLibrary.RESULT_OK
			resultVal.Retval = retVal.ToString()
		} else {
		}
		w.Write(resultVal.ToByte())

		return
	}

	resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_LINK, req.Host)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		w.Write(resultVal.ToByte())
		return
	}
	var customerId string = resultVal.Retval.(string)

	var currentData WasteLibrary.TagType = WasteLibrary.StringToTagType(req.FormValue(WasteLibrary.HTTP_DATA))

	resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_TAGS, customerId)
	if resultVal.Result == WasteLibrary.RESULT_OK {
		var currentCustomerTags WasteLibrary.CustomerTagsType = WasteLibrary.StringToCustomerTagsType(resultVal.Retval.(string))
		currentData = currentCustomerTags.Tags[currentData.ToIdString()]
		resultVal.Retval = currentData.ToString()
	}
	w.Write(resultVal.ToByte())
}

func checkAuth(data url.Values) WasteLibrary.ResultType {
	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	resultVal.Result = WasteLibrary.RESULT_OK
	return resultVal
}
