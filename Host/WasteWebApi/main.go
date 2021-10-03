package main

import (
	"net/http"
	"net/url"
	"time"

	"github.com/AfatekDevelopers/result_lib_go/devafatekresult"
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

	var resultVal devafatekresult.ResultType
	resultVal.Result = "FAIL"

	var linkVal string = req.Host
	WasteLibrary.LogStr("Get Link : " + linkVal)
	w.Write(resultVal.ToByte())
}

func getCustomer(w http.ResponseWriter, req *http.Request) {

	var resultVal devafatekresult.ResultType
	resultVal.Result = "FAIL"
	if err := req.ParseForm(); err != nil {
		WasteLibrary.LogErr(err)
		return
	}
	resultVal = checkAuth(req.Form)
	if resultVal.Result != "OK" {
		w.Write(resultVal.ToByte())
		return
	}
	if testapp {

		var retVal WasteLibrary.CustomerType = WasteLibrary.CustomerType{
			CustomerId:   1,
			CustomerName: "Afatek",
			CustomerLink: "afatek.aws.afatek.com.tr",
			RfIdApp:      "1",
			UltApp:       "1",
			RecyApp:      "1",
			Active:       "1",
			CreateTime:   time.Now().String(),
		}
		resultVal.Result = "OK"
		resultVal.Retval = retVal.ToString()

		w.Write(resultVal.ToByte())
		return
	}

	resultVal = WasteLibrary.GetRedisForStoreApi("customer-link", req.Host)
	if resultVal.Result != "OK" {
		w.Write(resultVal.ToByte())
		return
	}
	var customerId string = resultVal.Retval.(string)

	resultVal = WasteLibrary.GetRedisForStoreApi("customers", customerId)

	w.Write(resultVal.ToByte())
}

func getDevice(w http.ResponseWriter, req *http.Request) {

	var resultVal devafatekresult.ResultType
	resultVal.Result = "FAIL"
	if err := req.ParseForm(); err != nil {
		WasteLibrary.LogErr(err)
		return
	}
	resultVal = checkAuth(req.Form)
	if resultVal.Result != "OK" {
		w.Write(resultVal.ToByte())
		return
	}
	if testapp {
		var currentData WasteLibrary.DeviceType = WasteLibrary.StringToDeviceType(req.FormValue("DATA"))
		if currentData.DeviceId == 1 {
			var retVal WasteLibrary.DeviceType = WasteLibrary.DeviceType{
				DeviceId:              1,
				CustomerId:            1,
				DeviceName:            "06 AFA 001",
				DeviceType:            "RFID",
				SerialNumber:          "000001",
				ReaderAppStatus:       "1",
				ReaderAppLastOkTime:   "1",
				ReaderConnStatus:      "1",
				ReaderConnLastOkTime:  time.Now().String(),
				ReaderStatus:          "1",
				ReaderLastOkTime:      time.Now().String(),
				CamAppStatus:          "1",
				CamAppLastOkTime:      time.Now().String(),
				CamConnStatus:         "1",
				CamConnLastOkTime:     time.Now().String(),
				CamStatus:             "1",
				CamLastOkTime:         time.Now().String(),
				GpsAppStatus:          "1",
				GpsAppLastOkTime:      time.Now().String(),
				GpsConnStatus:         "1",
				GpsConnLastOkTime:     time.Now().String(),
				GpsStatus:             "1",
				GpsLastOkTime:         time.Now().String(),
				ThermAppStatus:        "1",
				ThermAppLastOkTime:    time.Now().String(),
				TransferAppStatus:     "1",
				TransferAppLastOkTime: time.Now().String(),
				AliveStatus:           "1",
				AliveLastOkTime:       time.Now().String(),
				ContactStatus:         "1",
				ContactLastOkTime:     time.Now().String(),
				Therm:                 "39",
				Latitude:              37.03633,
				Longitude:             27.41585,
				Speed:                 0,
				UltRange:              4785,
				UltStatus:             "1",
				DeviceStatus:          "1",
				AlarmStatus:           "1",
				TotalGlassCount:       345,
				TotalPlasticCount:     567,
				TotalMetalCount:       890,
				UltTime:               time.Now().String(),
				AlarmTime:             time.Now().String(),
				AlarmType:             "THERM",
				Alarm:                 "Therm : 85",
				RecyTime:              time.Now().String(),
				Active:                "1",
				ThermTime:             time.Now().String(),
				GpsTime:               time.Now().String(),
				StatusTime:            time.Now().String(),
				CreateTime:            time.Now().String(),
			}
			resultVal.Result = "OK"
			resultVal.Retval = retVal.ToString()
		} else if currentData.DeviceId == 2 {
			var retVal WasteLibrary.DeviceType = WasteLibrary.DeviceType{
				DeviceId:              2,
				CustomerId:            1,
				DeviceName:            "000002",
				DeviceType:            "ULT",
				SerialNumber:          "000002",
				ReaderAppStatus:       "1",
				ReaderAppLastOkTime:   "1",
				ReaderConnStatus:      "1",
				ReaderConnLastOkTime:  time.Now().String(),
				ReaderStatus:          "1",
				ReaderLastOkTime:      time.Now().String(),
				CamAppStatus:          "1",
				CamAppLastOkTime:      time.Now().String(),
				CamConnStatus:         "1",
				CamConnLastOkTime:     time.Now().String(),
				CamStatus:             "1",
				CamLastOkTime:         time.Now().String(),
				GpsAppStatus:          "1",
				GpsAppLastOkTime:      time.Now().String(),
				GpsConnStatus:         "1",
				GpsConnLastOkTime:     time.Now().String(),
				GpsStatus:             "1",
				GpsLastOkTime:         time.Now().String(),
				ThermAppStatus:        "1",
				ThermAppLastOkTime:    time.Now().String(),
				TransferAppStatus:     "1",
				TransferAppLastOkTime: time.Now().String(),
				AliveStatus:           "1",
				AliveLastOkTime:       time.Now().String(),
				ContactStatus:         "1",
				ContactLastOkTime:     time.Now().String(),
				Therm:                 "39",
				Latitude:              37.03652,
				Longitude:             27.42111,
				Speed:                 0,
				UltRange:              4785,
				UltStatus:             "1",
				DeviceStatus:          "1",
				AlarmStatus:           "1",
				TotalGlassCount:       345,
				TotalPlasticCount:     567,
				TotalMetalCount:       890,
				UltTime:               time.Now().String(),
				AlarmTime:             time.Now().String(),
				AlarmType:             "THERM",
				Alarm:                 "Therm : 85",
				RecyTime:              time.Now().String(),
				Active:                "1",
				ThermTime:             time.Now().String(),
				GpsTime:               time.Now().String(),
				StatusTime:            time.Now().String(),
				CreateTime:            time.Now().String(),
			}
			resultVal.Result = "OK"
			resultVal.Retval = retVal.ToString()
		} else if currentData.DeviceId == 3 {
			var retVal WasteLibrary.DeviceType = WasteLibrary.DeviceType{
				DeviceId:              3,
				CustomerId:            1,
				DeviceName:            "000003",
				DeviceType:            "RECY",
				SerialNumber:          "000003",
				ReaderAppStatus:       "1",
				ReaderAppLastOkTime:   "1",
				ReaderConnStatus:      "1",
				ReaderConnLastOkTime:  time.Now().String(),
				ReaderStatus:          "1",
				ReaderLastOkTime:      time.Now().String(),
				CamAppStatus:          "1",
				CamAppLastOkTime:      time.Now().String(),
				CamConnStatus:         "1",
				CamConnLastOkTime:     time.Now().String(),
				CamStatus:             "1",
				CamLastOkTime:         time.Now().String(),
				GpsAppStatus:          "1",
				GpsAppLastOkTime:      time.Now().String(),
				GpsConnStatus:         "1",
				GpsConnLastOkTime:     time.Now().String(),
				GpsStatus:             "1",
				GpsLastOkTime:         time.Now().String(),
				ThermAppStatus:        "1",
				ThermAppLastOkTime:    time.Now().String(),
				TransferAppStatus:     "1",
				TransferAppLastOkTime: time.Now().String(),
				AliveStatus:           "1",
				AliveLastOkTime:       time.Now().String(),
				ContactStatus:         "1",
				ContactLastOkTime:     time.Now().String(),
				Therm:                 "39",
				Latitude:              37.03732,
				Longitude:             27.41609,
				Speed:                 0,
				UltRange:              4785,
				UltStatus:             "1",
				DeviceStatus:          "1",
				AlarmStatus:           "1",
				TotalGlassCount:       345,
				TotalPlasticCount:     567,
				TotalMetalCount:       890,
				UltTime:               time.Now().String(),
				AlarmTime:             time.Now().String(),
				AlarmType:             "THERM",
				Alarm:                 "Therm : 85",
				RecyTime:              time.Now().String(),
				Active:                "1",
				ThermTime:             time.Now().String(),
				GpsTime:               time.Now().String(),
				StatusTime:            time.Now().String(),
				CreateTime:            time.Now().String(),
			}
			resultVal.Result = "OK"
			resultVal.Retval = retVal.ToString()
		} else {
		}
		w.Write(resultVal.ToByte())

		return
	}

	resultVal = WasteLibrary.GetRedisForStoreApi("customer-link", req.Host)
	if resultVal.Result != "OK" {
		w.Write(resultVal.ToByte())
		return
	}
	var customerId string = resultVal.Retval.(string)

	var currentData WasteLibrary.DeviceType = WasteLibrary.StringToDeviceType(req.FormValue("DATA"))
	resultVal = WasteLibrary.GetRedisForStoreApi("devices", currentData.ToIdString())

	currentData = WasteLibrary.StringToDeviceType(resultVal.Retval.(string))
	if currentData.ToCustomerIdString() == customerId {
		w.Write(resultVal.ToByte())
	} else {
		resultVal.Result = "FAIL"
		resultVal.Retval = ""
		w.Write(resultVal.ToByte())
	}
}

func getDevices(w http.ResponseWriter, req *http.Request) {

	var resultVal devafatekresult.ResultType
	resultVal.Result = "FAIL"
	if err := req.ParseForm(); err != nil {
		WasteLibrary.LogErr(err)
		return
	}
	resultVal = checkAuth(req.Form)
	if resultVal.Result != "OK" {
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
			DeviceType:            "RFID",
			SerialNumber:          "000001",
			ReaderAppStatus:       "1",
			ReaderAppLastOkTime:   "1",
			ReaderConnStatus:      "1",
			ReaderConnLastOkTime:  time.Now().String(),
			ReaderStatus:          "1",
			ReaderLastOkTime:      time.Now().String(),
			CamAppStatus:          "1",
			CamAppLastOkTime:      time.Now().String(),
			CamConnStatus:         "1",
			CamConnLastOkTime:     time.Now().String(),
			CamStatus:             "1",
			CamLastOkTime:         time.Now().String(),
			GpsAppStatus:          "1",
			GpsAppLastOkTime:      time.Now().String(),
			GpsConnStatus:         "1",
			GpsConnLastOkTime:     time.Now().String(),
			GpsStatus:             "1",
			GpsLastOkTime:         time.Now().String(),
			ThermAppStatus:        "1",
			ThermAppLastOkTime:    time.Now().String(),
			TransferAppStatus:     "1",
			TransferAppLastOkTime: time.Now().String(),
			AliveStatus:           "1",
			AliveLastOkTime:       time.Now().String(),
			ContactStatus:         "1",
			ContactLastOkTime:     time.Now().String(),
			Therm:                 "39",
			Latitude:              37.03633,
			Longitude:             27.41585,
			Speed:                 0,
			UltRange:              4785,
			UltStatus:             "1",
			DeviceStatus:          "1",
			AlarmStatus:           "1",
			TotalGlassCount:       345,
			TotalPlasticCount:     567,
			TotalMetalCount:       890,
			UltTime:               time.Now().String(),
			AlarmTime:             time.Now().String(),
			AlarmType:             "THERM",
			Alarm:                 "Therm : 85",
			RecyTime:              time.Now().String(),
			Active:                "1",
			ThermTime:             time.Now().String(),
			GpsTime:               time.Now().String(),
			StatusTime:            time.Now().String(),
			CreateTime:            time.Now().String(),
		}

		var retVal2 WasteLibrary.DeviceType = WasteLibrary.DeviceType{
			DeviceId:              2,
			CustomerId:            1,
			DeviceName:            "000002",
			DeviceType:            "ULT",
			SerialNumber:          "000002",
			ReaderAppStatus:       "1",
			ReaderAppLastOkTime:   "1",
			ReaderConnStatus:      "1",
			ReaderConnLastOkTime:  time.Now().String(),
			ReaderStatus:          "1",
			ReaderLastOkTime:      time.Now().String(),
			CamAppStatus:          "1",
			CamAppLastOkTime:      time.Now().String(),
			CamConnStatus:         "1",
			CamConnLastOkTime:     time.Now().String(),
			CamStatus:             "1",
			CamLastOkTime:         time.Now().String(),
			GpsAppStatus:          "1",
			GpsAppLastOkTime:      time.Now().String(),
			GpsConnStatus:         "1",
			GpsConnLastOkTime:     time.Now().String(),
			GpsStatus:             "1",
			GpsLastOkTime:         time.Now().String(),
			ThermAppStatus:        "1",
			ThermAppLastOkTime:    time.Now().String(),
			TransferAppStatus:     "1",
			TransferAppLastOkTime: time.Now().String(),
			AliveStatus:           "1",
			AliveLastOkTime:       time.Now().String(),
			ContactStatus:         "1",
			ContactLastOkTime:     time.Now().String(),
			Therm:                 "39",
			Latitude:              37.03652,
			Longitude:             27.42111,
			Speed:                 0,
			UltRange:              4785,
			UltStatus:             "1",
			DeviceStatus:          "1",
			AlarmStatus:           "1",
			TotalGlassCount:       345,
			TotalPlasticCount:     567,
			TotalMetalCount:       890,
			UltTime:               time.Now().String(),
			AlarmTime:             time.Now().String(),
			AlarmType:             "THERM",
			Alarm:                 "Therm : 85",
			RecyTime:              time.Now().String(),
			Active:                "1",
			ThermTime:             time.Now().String(),
			GpsTime:               time.Now().String(),
			StatusTime:            time.Now().String(),
			CreateTime:            time.Now().String(),
		}

		var retVal3 WasteLibrary.DeviceType = WasteLibrary.DeviceType{
			DeviceId:              3,
			CustomerId:            1,
			DeviceName:            "000003",
			DeviceType:            "RECY",
			SerialNumber:          "000003",
			ReaderAppStatus:       "1",
			ReaderAppLastOkTime:   "1",
			ReaderConnStatus:      "1",
			ReaderConnLastOkTime:  time.Now().String(),
			ReaderStatus:          "1",
			ReaderLastOkTime:      time.Now().String(),
			CamAppStatus:          "1",
			CamAppLastOkTime:      time.Now().String(),
			CamConnStatus:         "1",
			CamConnLastOkTime:     time.Now().String(),
			CamStatus:             "1",
			CamLastOkTime:         time.Now().String(),
			GpsAppStatus:          "1",
			GpsAppLastOkTime:      time.Now().String(),
			GpsConnStatus:         "1",
			GpsConnLastOkTime:     time.Now().String(),
			GpsStatus:             "1",
			GpsLastOkTime:         time.Now().String(),
			ThermAppStatus:        "1",
			ThermAppLastOkTime:    time.Now().String(),
			TransferAppStatus:     "1",
			TransferAppLastOkTime: time.Now().String(),
			AliveStatus:           "1",
			AliveLastOkTime:       time.Now().String(),
			ContactStatus:         "1",
			ContactLastOkTime:     time.Now().String(),
			Therm:                 "39",
			Latitude:              37.03732,
			Longitude:             27.41609,
			Speed:                 0,
			UltRange:              4785,
			UltStatus:             "1",
			DeviceStatus:          "1",
			AlarmStatus:           "1",
			TotalGlassCount:       345,
			TotalPlasticCount:     567,
			TotalMetalCount:       890,
			UltTime:               time.Now().String(),
			AlarmTime:             time.Now().String(),
			AlarmType:             "THERM",
			Alarm:                 "Therm : 85",
			RecyTime:              time.Now().String(),
			Active:                "1",
			ThermTime:             time.Now().String(),
			GpsTime:               time.Now().String(),
			StatusTime:            time.Now().String(),
			CreateTime:            time.Now().String(),
		}

		retVal.Devices[retVal1.ToIdString()] = retVal1
		retVal.Devices[retVal2.ToIdString()] = retVal2
		retVal.Devices[retVal3.ToIdString()] = retVal3

		resultVal.Result = "OK"
		resultVal.Retval = retVal.ToString()

		w.Write(resultVal.ToByte())

		return
	}

	resultVal = WasteLibrary.GetRedisForStoreApi("customer-link", req.Host)
	if resultVal.Result != "OK" {
		w.Write(resultVal.ToByte())
		return
	}
	var customerId string = resultVal.Retval.(string)

	resultVal = WasteLibrary.GetRedisForStoreApi("customer-devices", customerId)

	if resultVal.Result == "OK" {

		var customerDevices WasteLibrary.CustomerDevicesType = WasteLibrary.StringToCustomerDevicesType(resultVal.Retval.(string))
		var customerDevicesList WasteLibrary.CustomerDevicesListType = WasteLibrary.CustomerDevicesListType{
			CustomerId: WasteLibrary.StringIdToFloat64(customerId),
			Devices:    make(map[string]WasteLibrary.DeviceType),
		}
		for _, deviceId := range customerDevices.Devices {

			if deviceId != 0 {
				resultVal = WasteLibrary.GetRedisForStoreApi("devices", WasteLibrary.Float64IdToString(deviceId))
				if resultVal.Result == "OK" {
					var currentDevice WasteLibrary.DeviceType = WasteLibrary.StringToDeviceType(resultVal.Retval.(string))
					customerDevicesList.Devices[currentDevice.ToIdString()] = currentDevice
				}

			}
		}
		resultVal.Result = "OK"
		resultVal.Retval = customerDevicesList.ToString()
	}

	w.Write(resultVal.ToByte())
}

func getConfig(w http.ResponseWriter, req *http.Request) {

	var resultVal devafatekresult.ResultType
	resultVal.Result = "FAIL"
	if err := req.ParseForm(); err != nil {
		WasteLibrary.LogErr(err)
		return
	}
	resultVal = checkAuth(req.Form)
	if resultVal.Result != "OK" {
		w.Write(resultVal.ToByte())
		return
	}
	if testapp {

		var currentHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.StringToHttpClientHeaderType(req.FormValue("HEADER"))
		if currentHttpHeader.OpType == "CUSTOMER" {
			var retVal WasteLibrary.CustomerConfigType = WasteLibrary.CustomerConfigType{
				CustomerId:      1,
				ArventoApp:      "1",
				ArventoUserName: "test1",
				ArventoPin1:     "pin1",
				ArventoPin2:     "pin2",
				SystemProblem:   "1",
				TruckStopTrace:  "1",
				Active:          "1",
				CreateTime:      time.Now().String(),
			}
			resultVal.Result = "OK"
			resultVal.Retval = retVal.ToString()
		} else if currentHttpHeader.OpType == "ADMIN" {
			var retVal WasteLibrary.AdminConfigType = WasteLibrary.AdminConfigType{
				CustomerId: 1,
				Active:     "1",
				CreateTime: time.Now().String(),
			}
			resultVal.Result = "OK"
			resultVal.Retval = retVal.ToString()
		} else if currentHttpHeader.OpType == "LOCAL" {
			var retVal WasteLibrary.LocalConfigType = WasteLibrary.LocalConfigType{
				CustomerId: 1,
				Active:     "1",
				CreateTime: time.Now().String(),
			}
			resultVal.Result = "OK"
			resultVal.Retval = retVal.ToString()
		} else {
			resultVal.Result = "FAIL"
		}

		w.Write(resultVal.ToByte())
		return
	}

	resultVal = WasteLibrary.GetRedisForStoreApi("customer-link", req.Host)
	if resultVal.Result != "OK" {
		w.Write(resultVal.ToByte())
		return
	}
	var customerId string = resultVal.Retval.(string)

	var currentHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.StringToHttpClientHeaderType(req.FormValue("HEADER"))
	if currentHttpHeader.OpType == "CUSTOMER" {
		resultVal = WasteLibrary.GetRedisForStoreApi("customer-customerconfig", customerId)
	} else if currentHttpHeader.OpType == "ADMIN" {
		resultVal = WasteLibrary.GetRedisForStoreApi("customer-adminconfig", customerId)
	} else if currentHttpHeader.OpType == "LOCAL" {
		resultVal = WasteLibrary.GetRedisForStoreApi("customer-localconfig", customerId)
	} else {
		resultVal.Result = "FAIL"
	}

	w.Write(resultVal.ToByte())
}

func getTags(w http.ResponseWriter, req *http.Request) {

	var resultVal devafatekresult.ResultType
	resultVal.Result = "FAIL"
	if err := req.ParseForm(); err != nil {
		WasteLibrary.LogErr(err)
		return
	}
	resultVal = checkAuth(req.Form)
	if resultVal.Result != "OK" {
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
			Statu:       "0",
			ImageStatu:  "1",
			Active:      "1",
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
			Statu:       "1",
			ImageStatu:  "1",
			Active:      "1",
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
			ImageStatu:  "1",
			Active:      "1",
			ReadTime:    time.Now().String(),
			CheckTime:   time.Now().String(),
			CreateTime:  time.Now().String(),
		}

		retVal.Tags[retVal1.ToIdString()] = retVal1
		retVal.Tags[retVal2.ToIdString()] = retVal2
		retVal.Tags[retVal3.ToIdString()] = retVal3

		resultVal.Result = "OK"
		resultVal.Retval = retVal.ToString()

		w.Write(resultVal.ToByte())

		return
	}

	resultVal = WasteLibrary.GetRedisForStoreApi("customer-link", req.Host)
	if resultVal.Result != "OK" {
		w.Write(resultVal.ToByte())
		return
	}
	var customerId string = resultVal.Retval.(string)

	resultVal = WasteLibrary.GetRedisForStoreApi("customer-tags", customerId)

	w.Write(resultVal.ToByte())
}

func getTag(w http.ResponseWriter, req *http.Request) {

	var resultVal devafatekresult.ResultType
	resultVal.Result = "FAIL"
	if err := req.ParseForm(); err != nil {
		WasteLibrary.LogErr(err)
		return
	}
	resultVal = checkAuth(req.Form)
	if resultVal.Result != "OK" {
		w.Write(resultVal.ToByte())
		return
	}

	if testapp {
		var currentData WasteLibrary.TagType = WasteLibrary.StringToTagType(req.FormValue("DATA"))
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
				Statu:       "0",
				ImageStatu:  "1",
				Active:      "1",
				ReadTime:    time.Now().String(),
				CheckTime:   time.Now().String(),
				CreateTime:  time.Now().String(),
			}

			resultVal.Result = "OK"
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
				Statu:       "1",
				ImageStatu:  "1",
				Active:      "1",
				ReadTime:    time.Now().String(),
				CheckTime:   time.Now().String(),
				CreateTime:  time.Now().String(),
			}

			resultVal.Result = "OK"
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
				ImageStatu:  "1",
				Active:      "1",
				ReadTime:    time.Now().String(),
				CheckTime:   time.Now().String(),
				CreateTime:  time.Now().String(),
			}
			resultVal.Result = "OK"
			resultVal.Retval = retVal.ToString()
		} else {
		}
		w.Write(resultVal.ToByte())

		return
	}

	resultVal = WasteLibrary.GetRedisForStoreApi("customer-link", req.Host)
	if resultVal.Result != "OK" {
		w.Write(resultVal.ToByte())
		return
	}
	var customerId string = resultVal.Retval.(string)

	var currentData WasteLibrary.TagType = WasteLibrary.StringToTagType(req.FormValue("DATA"))

	resultVal = WasteLibrary.GetRedisForStoreApi("customer-tags", customerId)
	if resultVal.Result == "OK" {
		var currentCustomerTags WasteLibrary.CustomerTagsType = WasteLibrary.StringToCustomerTagsType(resultVal.Retval.(string))
		currentData = currentCustomerTags.Tags[currentData.ToIdString()]
		resultVal.Retval = currentData.ToString()
	}
	w.Write(resultVal.ToByte())
}

func checkAuth(data url.Values) devafatekresult.ResultType {
	var resultVal devafatekresult.ResultType
	resultVal.Result = "FAIL"
	resultVal.Result = "OK"
	return resultVal
}
