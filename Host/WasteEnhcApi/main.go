package main

import (
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
	http.HandleFunc("/createDevice", createDevice)
	http.HandleFunc("/createTag", createTag)
	http.ListenAndServe(":80", nil)
}

func createDevice(w http.ResponseWriter, req *http.Request) {

	if WasteLibrary.AllowCors {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,access-control-allow-origin, access-control-allow-headers")
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

	var currentHttpHeader WasteLibrary.HttpClientHeaderType
	currentHttpHeader.StringToType(req.FormValue(WasteLibrary.HTTP_HEADER))
	if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_RFID {
		//DeviceMMainType    RfidDeviceMainType
		var currentDeviceMain WasteLibrary.RfidDeviceMainType
		currentDeviceMain.New()
		//TO DO
		//remove after bodrum
		currentDeviceMain.CustomerId = 3
		currentDeviceMain.SerialNumber = currentHttpHeader.DeviceNo

		resultVal = currentDeviceMain.SaveToDb()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		resultVal = currentDeviceMain.SaveToRedisBySerial()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		resultVal = currentDeviceMain.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		//DeviceBase    RfidDeviceBaseType
		var currentDeviceBase WasteLibrary.RfidDeviceBaseType
		currentDeviceBase.New()
		currentDeviceBase.NewData = true

		//TO DO
		//remove after bodrum
		if currentHttpHeader.DeviceNo == "0000000000823756" {
			currentDeviceBase.PlateNo = "07 ADB 029"
		}
		if currentHttpHeader.DeviceNo == "00000000012e772a" {
			currentDeviceBase.PlateNo = "07 AAU 821"
		}
		if currentHttpHeader.DeviceNo == "0000000048b73c8b" {
			currentDeviceBase.PlateNo = "07 MVS 23"
		}
		if currentHttpHeader.DeviceNo == "000000005679401f" {
			currentDeviceBase.PlateNo = "07 MVS 29"
		}
		if currentHttpHeader.DeviceNo == "00000000ce54360b" {
			currentDeviceBase.PlateNo = "07 AAU 614"
		}
		if currentHttpHeader.DeviceNo == "00000000ff38daee" {
			currentDeviceBase.PlateNo = "07 ADA 288"
		}
		if currentHttpHeader.DeviceNo == "00000000b8b55550" {
			currentDeviceBase.PlateNo = "07 AAV 297"
		}
		if currentHttpHeader.DeviceNo == "0000000041e0f19b" {
			currentDeviceBase.PlateNo = "07 AAU 701"
		}
		if currentHttpHeader.DeviceNo == "00000000d7b45eaf" {
			currentDeviceBase.PlateNo = "07 AAV 086"
		}
		if currentHttpHeader.DeviceNo == "0000000038a0df52" {
			currentDeviceBase.PlateNo = "07 AAU 674"
		}
		if currentHttpHeader.DeviceNo == "00000000af4aaeab" {
			currentDeviceBase.PlateNo = "07 AJV 476"
		}
		if currentHttpHeader.DeviceNo == "000000009fb47031" {
			currentDeviceBase.PlateNo = "07 MVS 21"
		}
		if currentHttpHeader.DeviceNo == "0000000064a4a3fb" {
			currentDeviceBase.PlateNo = "07 AAV 271"
		}
		if currentHttpHeader.DeviceNo == "00000000dddd5738" {
			currentDeviceBase.PlateNo = "07 MVS 25"
		}
		if currentHttpHeader.DeviceNo == "00000000a931fb0d" {
			currentDeviceBase.PlateNo = "07 ADA 780"
		}
		if currentHttpHeader.DeviceNo == "00000000581973e0" {
			currentDeviceBase.PlateNo = "07 ADA 695"
		}
		if currentHttpHeader.DeviceNo == "000000007c6d676b" {
			currentDeviceBase.PlateNo = "07 AAV 316"
		}
		if currentHttpHeader.DeviceNo == "00000000c358d010" {
			currentDeviceBase.PlateNo = "48 UM 616"
		}
		if currentHttpHeader.DeviceNo == "000000009ec0646a" {
			currentDeviceBase.PlateNo = "07 MVS 27"
		}
		if currentHttpHeader.DeviceNo == "00000000580f3aa4" {
			currentDeviceBase.PlateNo = "07 ADA 303"
		}
		if currentHttpHeader.DeviceNo == "000000001bf125dc" {
			currentDeviceBase.PlateNo = "48 TR 218"
		}
		if currentHttpHeader.DeviceNo == "00000000ff2e9d8c" {
			currentDeviceBase.PlateNo = "07 AAV 113"
		}
		if currentHttpHeader.DeviceNo == "00000000cb22e4a1" {
			currentDeviceBase.PlateNo = "07 AAV 235"
		}
		if currentHttpHeader.DeviceNo == "000000004a65f340" {
			currentDeviceBase.PlateNo = "07 ABE 913"
		}
		if currentHttpHeader.DeviceNo == "00000000be0d41f4" {
			currentDeviceBase.PlateNo = "07 ADA 718"
		}
		if currentHttpHeader.DeviceNo == "00000000ad38f882" {
			currentDeviceBase.PlateNo = "48 UC 593"
		}
		if currentHttpHeader.DeviceNo == "000000008ce06a67" {
			currentDeviceBase.PlateNo = "07 ADA 714"
		}
		if currentHttpHeader.DeviceNo == "00000000196a3d01" {
			currentDeviceBase.PlateNo = "07 AAU 773"
		}
		if currentHttpHeader.DeviceNo == "00000000613ea9b3" {
			currentDeviceBase.PlateNo = "07 AAV 123"
		}
		if currentHttpHeader.DeviceNo == "00000000de434b82" {
			currentDeviceBase.PlateNo = "07 ADA 703"
		}
		if currentHttpHeader.DeviceNo == "0000000064231683" {
			currentDeviceBase.PlateNo = "07 AAU 461"
		}
		if currentHttpHeader.DeviceNo == "000000000f2d7706" {
			currentDeviceBase.PlateNo = "48 UE 548"
		}
		if currentHttpHeader.DeviceNo == "000000007f994bb3" {
			currentDeviceBase.PlateNo = "07 AAV 247"
		}
		if currentHttpHeader.DeviceNo == "000000004e64f0f2" {
			currentDeviceBase.PlateNo = "07 AJE 720"
		}
		if currentHttpHeader.DeviceNo == "0000000087abde8b" {
			currentDeviceBase.PlateNo = "07 ADA 770"
		}
		if currentHttpHeader.DeviceNo == "000000004154c1fe" {
			currentDeviceBase.PlateNo = "48 VS 510"
		}
		if currentHttpHeader.DeviceNo == "000000009a5b3613" {
			currentDeviceBase.PlateNo = "07 ADA 255"
		}
		if currentHttpHeader.DeviceNo == "00000000dbf8d383" {
			currentDeviceBase.PlateNo = "07 AHV 882"
		}
		if currentHttpHeader.DeviceNo == "000000007d4a013b" {
			currentDeviceBase.PlateNo = "07 AAU 876"
		}
		if currentHttpHeader.DeviceNo == "0000000000620c13" {
			currentDeviceBase.PlateNo = "07 AAV 200"
		}
		if currentHttpHeader.DeviceNo == "000000009b0c8bdd" {
			currentDeviceBase.PlateNo = "07 MVS 19"
		}
		if currentHttpHeader.DeviceNo == "000000005adcff3f" {
			currentDeviceBase.PlateNo = "07 MVS 31"
		}
		if currentHttpHeader.DeviceNo == "00000000b74fb0f0" {
			currentDeviceBase.PlateNo = "48 VJ 050"
		}
		if currentHttpHeader.DeviceNo == "00000000ca7c41bc" {
			currentDeviceBase.PlateNo = "07 AAV 165"
		}
		if currentHttpHeader.DeviceNo == "00000000156ce38c" {
			currentDeviceBase.PlateNo = "07 ABE 586"
		}
		if currentHttpHeader.DeviceNo == "00000000170c24a8" {
			currentDeviceBase.PlateNo = "07 ABE 301"
		}
		if currentHttpHeader.DeviceNo == "000000006c12e23d" {
			currentDeviceBase.PlateNo = "07 MVS 33"
		}
		if currentHttpHeader.DeviceNo == "000000005b8da4c2" {
			currentDeviceBase.PlateNo = "48 VJ 783"
		}
		if currentHttpHeader.DeviceNo == "0000000076ef40bd" {
			currentDeviceBase.PlateNo = "07 AAU 537"
		}
		if currentHttpHeader.DeviceNo == "00000000e549e8ed" {
			currentDeviceBase.PlateNo = "07 ADA 229"
		}

		currentDeviceBase.DeviceId = currentDeviceMain.DeviceId
		resultVal = currentDeviceBase.SaveToDb()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		resultVal = currentDeviceBase.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		//DeviceStatu   RfidDeviceStatuType
		var currentDeviceStatu WasteLibrary.RfidDeviceStatuType
		currentDeviceStatu.New()
		currentDeviceStatu.NewData = true
		currentDeviceStatu.DeviceId = currentDeviceMain.DeviceId
		resultVal = currentDeviceStatu.SaveToDb()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		resultVal = currentDeviceStatu.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		//DeviceGps     RfidDeviceGpsType
		var currentDeviceGps WasteLibrary.RfidDeviceGpsType
		currentDeviceGps.New()
		currentDeviceGps.NewData = true
		currentDeviceGps.DeviceId = currentDeviceMain.DeviceId
		resultVal = currentDeviceGps.SaveToDb()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		resultVal = currentDeviceGps.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		//DeviceEmbededGps     RfidDeviceEmbededGpsType
		var currentDeviceEmbededGps WasteLibrary.RfidDeviceEmbededGpsType
		currentDeviceEmbededGps.New()
		currentDeviceEmbededGps.NewData = true
		currentDeviceEmbededGps.DeviceId = currentDeviceMain.DeviceId
		resultVal = currentDeviceEmbededGps.SaveToDb()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		resultVal = currentDeviceEmbededGps.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		//DeviceAlarm   RfidDeviceAlarmType
		var currentDeviceAlarm WasteLibrary.RfidDeviceAlarmType
		currentDeviceAlarm.New()
		currentDeviceAlarm.NewData = true
		currentDeviceAlarm.DeviceId = currentDeviceMain.DeviceId
		resultVal = currentDeviceAlarm.SaveToDb()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		resultVal = currentDeviceAlarm.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		//DeviceTherm   RfidDeviceThermType
		var currentDeviceTherm WasteLibrary.RfidDeviceThermType
		currentDeviceTherm.New()
		currentDeviceTherm.NewData = true
		currentDeviceTherm.DeviceId = currentDeviceMain.DeviceId
		resultVal = currentDeviceTherm.SaveToDb()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		resultVal = currentDeviceTherm.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		//DeviceVersion RfidDeviceVersionType
		var currentDeviceVersion WasteLibrary.RfidDeviceVersionType
		currentDeviceVersion.New()
		currentDeviceVersion.NewData = true
		currentDeviceVersion.DeviceId = currentDeviceMain.DeviceId
		resultVal = currentDeviceVersion.SaveToDb()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		resultVal = currentDeviceVersion.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		//DeviceWorkHour RfidDeviceWorkHourType
		var currentDeviceWorkHour WasteLibrary.RfidDeviceWorkHourType
		currentDeviceWorkHour.New()
		currentDeviceWorkHour.NewData = true
		currentDeviceWorkHour.DeviceId = currentDeviceMain.DeviceId
		resultVal = currentDeviceWorkHour.SaveToDb()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		resultVal = currentDeviceWorkHour.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		//DeviceNote RfidDeviceNoteType
		var currentDeviceNote WasteLibrary.RfidDeviceNoteType
		currentDeviceNote.New()
		currentDeviceNote.NewData = true
		currentDeviceNote.DeviceId = currentDeviceMain.DeviceId
		resultVal = currentDeviceNote.SaveToDb()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		resultVal = currentDeviceNote.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		//DeviceReport RfidDeviceReportType
		var currentDeviceReport WasteLibrary.RfidDeviceReportType
		currentDeviceReport.New()
		currentDeviceReport.NewData = true
		currentDeviceReport.DeviceId = currentDeviceMain.DeviceId
		resultVal = currentDeviceReport.SaveToDb()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		resultVal = currentDeviceReport.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		//DeviceMotion RfidDeviceMotionType WODb
		var currentMotion WasteLibrary.RfidDeviceMotionType
		currentMotion.New()
		currentMotion.DeviceId = currentDeviceMain.DeviceId
		resultVal = currentMotion.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		//DeviceTag RfidDeviceTagType WODb
		var currentTag WasteLibrary.RfidDeviceTagType
		currentTag.New()
		currentTag.DeviceId = currentDeviceMain.DeviceId
		resultVal = currentTag.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		var customerDevices WasteLibrary.CustomerRfidDevicesType
		customerDevices.CustomerId = currentDeviceMain.CustomerId
		resultVal = customerDevices.GetByRedis("0")
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_GET
			w.Write(resultVal.ToByte())

			return
		}

		customerDevices.Devices[currentDeviceMain.ToIdString()] = currentDeviceMain.DeviceId
		resultVal = customerDevices.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}

	} else if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_ULT {
		var currentUlt WasteLibrary.UltDeviceType
		currentUlt.StringToType(req.FormValue(WasteLibrary.HTTP_DATA))
		//DeviceMainType    UltDeviceMainType
		var currentDeviceMain WasteLibrary.UltDeviceMainType
		currentDeviceMain.New()
		//TO DO
		//remove after bodrum
		currentDeviceMain.CustomerId = 3
		if currentHttpHeader.DeviceNo == "286016316690395" {
			currentDeviceMain.Latitude = 37.052582
			currentDeviceMain.Longitude = 27.382131
		} else if currentHttpHeader.DeviceNo == "286016316690635" {
			currentDeviceMain.Latitude = 37.052591
			currentDeviceMain.Longitude = 27.382524
		} else {
			currentDeviceMain.Latitude = 37.0351013
			currentDeviceMain.Longitude = 27.4305283

		}

		currentDeviceMain.SerialNumber = currentHttpHeader.DeviceNo
		resultVal = currentDeviceMain.SaveToDb()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		resultVal = currentDeviceMain.SaveToRedisBySerial()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		resultVal = currentDeviceMain.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		//DeviceBase    UltDeviceBaseType
		var currentDeviceBase WasteLibrary.UltDeviceBaseType
		currentDeviceBase.New()
		currentDeviceBase.NewData = true
		currentDeviceBase.DeviceId = currentDeviceMain.DeviceId
		resultVal = currentDeviceBase.SaveToDb()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		resultVal = currentDeviceBase.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		//DeviceStatu   UltDeviceStatuType
		var currentDeviceStatu WasteLibrary.UltDeviceStatuType
		currentDeviceStatu.New()
		currentDeviceStatu.NewData = true
		currentDeviceStatu.DeviceId = currentDeviceMain.DeviceId
		resultVal = currentDeviceStatu.SaveToDb()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		resultVal = currentDeviceStatu.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		//DeviceGps     UltDeviceGpsType
		var currentDeviceGps WasteLibrary.UltDeviceGpsType
		currentDeviceGps.New()
		currentDeviceGps.NewData = true
		currentDeviceGps.DeviceId = currentDeviceMain.DeviceId
		resultVal = currentDeviceGps.SaveToDb()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		resultVal = currentDeviceGps.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		//DeviceAlarm   UltDeviceAlarmType
		var currentDeviceAlarm WasteLibrary.UltDeviceAlarmType
		currentDeviceAlarm.New()
		currentDeviceAlarm.NewData = true
		currentDeviceAlarm.DeviceId = currentDeviceMain.DeviceId
		resultVal = currentDeviceAlarm.SaveToDb()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		resultVal = currentDeviceAlarm.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		//DeviceTherm   UltDeviceThermType
		var currentDeviceTherm WasteLibrary.UltDeviceThermType
		currentDeviceTherm.New()
		currentDeviceTherm.NewData = true
		currentDeviceTherm.DeviceId = currentDeviceMain.DeviceId
		resultVal = currentDeviceTherm.SaveToDb()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		resultVal = currentDeviceTherm.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		//DeviceVersion UltDeviceVersionType
		var currentDeviceVersion WasteLibrary.UltDeviceVersionType
		currentDeviceVersion.New()
		currentDeviceVersion.NewData = true
		currentDeviceVersion.DeviceId = currentDeviceMain.DeviceId
		resultVal = currentDeviceVersion.SaveToDb()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		resultVal = currentDeviceVersion.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		//DeviceBattery  UltDeviceBatteryType
		var currentDeviceBattery WasteLibrary.UltDeviceBatteryType
		currentDeviceBattery.New()
		currentDeviceBattery.NewData = true
		currentDeviceBattery.DeviceId = currentDeviceMain.DeviceId
		resultVal = currentDeviceBattery.SaveToDb()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		resultVal = currentDeviceBattery.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		//DeviceSens  UltDeviceSensType
		var currentDeviceSens WasteLibrary.UltDeviceSensType
		currentDeviceSens.New()
		currentDeviceSens.NewData = true
		currentDeviceSens.DeviceId = currentDeviceMain.DeviceId
		resultVal = currentDeviceSens.SaveToDb()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		resultVal = currentDeviceSens.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		//DeviceNote UltDeviceNoteType
		var currentDeviceNote WasteLibrary.UltDeviceNoteType
		currentDeviceNote.New()
		currentDeviceNote.NewData = true
		currentDeviceNote.DeviceId = currentDeviceMain.DeviceId
		resultVal = currentDeviceNote.SaveToDb()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		resultVal = currentDeviceNote.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		//DeviceSim UltDeviceSimType
		var currentDeviceSim WasteLibrary.UltDeviceSimType
		currentDeviceSim.New()
		currentDeviceSim.NewData = true
		currentDeviceSim.DeviceId = currentDeviceMain.DeviceId
		currentDeviceSim.Imei = currentUlt.DeviceSim.Imei
		currentDeviceSim.Imsi = currentUlt.DeviceSim.Imsi
		resultVal = currentDeviceSim.SaveToDb()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		resultVal = currentDeviceSim.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		var customerDevices WasteLibrary.CustomerUltDevicesType
		customerDevices.CustomerId = currentDeviceMain.CustomerId
		resultVal = customerDevices.GetByRedis("0")
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_GET
			w.Write(resultVal.ToByte())

			return
		}

		customerDevices.Devices[currentDeviceMain.ToIdString()] = currentDeviceMain.DeviceId
		resultVal = customerDevices.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}

	} else if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_RECY {
		//DeviceMainType    RecyDeviceMainType
		var currentDeviceMain WasteLibrary.RecyDeviceMainType
		currentDeviceMain.New()
		//TO DO
		//remove after bodrum
		currentDeviceMain.CustomerId = 3
		currentDeviceMain.Latitude = 37.0351013
		currentDeviceMain.Longitude = 27.4305283

		currentDeviceMain.SerialNumber = currentHttpHeader.DeviceNo
		resultVal = currentDeviceMain.SaveToDb()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		resultVal = currentDeviceMain.SaveToRedisBySerial()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		resultVal = currentDeviceMain.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		//DeviceBase    RecyDeviceBaseType
		var currentDeviceBase WasteLibrary.RecyDeviceBaseType
		currentDeviceBase.New()
		currentDeviceBase.NewData = true
		currentDeviceBase.DeviceId = currentDeviceMain.DeviceId
		resultVal = currentDeviceBase.SaveToDb()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		resultVal = currentDeviceBase.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		//DeviceStatu   RecyDeviceStatuType
		var currentDeviceStatu WasteLibrary.RecyDeviceStatuType
		currentDeviceStatu.New()
		currentDeviceStatu.NewData = true
		currentDeviceStatu.DeviceId = currentDeviceMain.DeviceId
		resultVal = currentDeviceStatu.SaveToDb()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		resultVal = currentDeviceStatu.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		//DeviceGps     RecyDeviceGpsType
		var currentDeviceGps WasteLibrary.RecyDeviceGpsType
		currentDeviceGps.New()
		currentDeviceGps.NewData = true
		currentDeviceGps.DeviceId = currentDeviceMain.DeviceId
		resultVal = currentDeviceGps.SaveToDb()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		resultVal = currentDeviceGps.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		//DeviceAlarm   RecyDeviceAlarmType
		var currentDeviceAlarm WasteLibrary.RecyDeviceAlarmType
		currentDeviceAlarm.New()
		currentDeviceAlarm.NewData = true
		currentDeviceAlarm.DeviceId = currentDeviceMain.DeviceId
		resultVal = currentDeviceAlarm.SaveToDb()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		resultVal = currentDeviceAlarm.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		//DeviceTherm   RecyDeviceThermType
		var currentDeviceTherm WasteLibrary.RecyDeviceThermType
		currentDeviceTherm.New()
		currentDeviceTherm.NewData = true
		currentDeviceTherm.DeviceId = currentDeviceMain.DeviceId
		resultVal = currentDeviceTherm.SaveToDb()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		resultVal = currentDeviceTherm.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		//DeviceVersion RecyDeviceVersionType
		var currentDeviceVersion WasteLibrary.RecyDeviceVersionType
		currentDeviceVersion.New()
		currentDeviceVersion.NewData = true
		currentDeviceVersion.DeviceId = currentDeviceMain.DeviceId
		resultVal = currentDeviceVersion.SaveToDb()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		resultVal = currentDeviceVersion.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		//DeviceDetail  RecyDeviceDetailType
		var currentDeviceDetail WasteLibrary.RecyDeviceDetailType
		currentDeviceDetail.New()
		currentDeviceDetail.NewData = true
		currentDeviceDetail.DeviceId = currentDeviceMain.DeviceId
		resultVal = currentDeviceDetail.SaveToDb()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		resultVal = currentDeviceDetail.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		//DeviceNote RecyDeviceNoteType
		var currentDeviceNote WasteLibrary.RecyDeviceNoteType
		currentDeviceNote.New()
		currentDeviceNote.NewData = true
		currentDeviceNote.DeviceId = currentDeviceMain.DeviceId
		resultVal = currentDeviceNote.SaveToDb()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		resultVal = currentDeviceNote.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		var customerDevices WasteLibrary.CustomerRecyDevicesType
		customerDevices.CustomerId = currentDeviceMain.CustomerId
		resultVal = customerDevices.GetByRedis("0")
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_GET
			w.Write(resultVal.ToByte())

			return
		}

		customerDevices.Devices[currentDeviceMain.ToIdString()] = currentDeviceMain.DeviceId
		resultVal = customerDevices.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}

	} else {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_DEVICE_NOTFOUND
	}

	w.Write(resultVal.ToByte())

}

func createTag(w http.ResponseWriter, req *http.Request) {

	if WasteLibrary.AllowCors {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,access-control-allow-origin, access-control-allow-headers")
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

	var currentData WasteLibrary.TagType
	currentData.StringToType(req.FormValue(WasteLibrary.HTTP_DATA))
	if currentData.TagMain.CustomerId > 1 {
		//TagMainType    TagMainType
		resultVal = currentData.TagMain.SaveToDb()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		currentData.TagId = currentData.TagMain.TagId

		resultVal = currentData.TagMain.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		resultVal = currentData.TagMain.SaveToRedisByEpc()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		//TagBase    TagBaseType
		var currentTagBase WasteLibrary.TagBaseType
		currentTagBase.New()
		currentTagBase.NewData = true
		currentTagBase.TagId = currentData.TagId
		resultVal = currentTagBase.SaveToDb()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		resultVal = currentTagBase.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		//TagStatu   TagStatuType
		var currentTagStatu WasteLibrary.TagStatuType
		currentTagStatu.New()
		currentTagStatu.NewData = true
		currentTagStatu.TagId = currentData.TagId
		resultVal = currentTagStatu.SaveToDb()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		resultVal = currentTagStatu.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		//TagGps     TagGpsType
		var currentTagGps WasteLibrary.TagGpsType
		currentTagGps.New()
		currentTagGps.NewData = true
		currentTagGps.TagId = currentData.TagId
		resultVal = currentTagGps.SaveToDb()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		resultVal = currentTagGps.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		//TagReader   TagReaderType
		var currentTagReader WasteLibrary.TagReaderType
		currentTagReader.New()
		currentTagReader.NewData = true
		currentTagReader.TagId = currentData.TagId
		resultVal = currentTagReader.SaveToDb()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		resultVal = currentTagReader.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		//TagNote   TagNoteType
		var currentTagNote WasteLibrary.TagNoteType
		currentTagNote.New()
		currentTagNote.NewData = true
		currentTagNote.TagId = currentData.TagId
		resultVal = currentTagNote.SaveToDb()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		resultVal = currentTagNote.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		//TagAlarm   TagAlarmType
		var currentTagAlarm WasteLibrary.TagAlarmType
		currentTagAlarm.New()
		currentTagAlarm.NewData = true
		currentTagAlarm.TagId = currentData.TagId
		resultVal = currentTagAlarm.SaveToDb()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		resultVal = currentTagAlarm.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		//TagReadDevice   TagReadDeviceType WODb
		var currentTagReadDevice WasteLibrary.TagReadDeviceType
		currentTagReadDevice.New()
		currentTagReadDevice.TagId = currentData.TagId
		resultVal = currentTagReadDevice.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		//TagPositionChange   TagPositionChangeType WODb
		var currentTagPositionChange WasteLibrary.TagPositionChangeType
		currentTagPositionChange.New()
		currentTagPositionChange.TagId = currentData.TagId
		resultVal = currentTagPositionChange.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		var customerTags WasteLibrary.CustomerTagsType
		customerTags.CustomerId = currentData.TagMain.CustomerId
		resultVal = customerTags.GetByRedis("0")
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_GET
			w.Write(resultVal.ToByte())

			return
		}

		customerTags.Tags[currentData.TagMain.ToIdString()] = currentData.TagMain.TagId
		resultVal = customerTags.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		w.Write(resultVal.ToByte())
	} else {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_TAG_CUSTOMER_NOTFOUND
		w.Write(resultVal.ToByte())
		return
	}

}
