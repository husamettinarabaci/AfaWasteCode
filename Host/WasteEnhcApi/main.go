package main

import (
	"net/http"
	"net/url"

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
		//DeviceMMainType    RfidDeviceMainType
		currentHttpHeader.DataType = WasteLibrary.DATATYPE_RFID_MAIN_DEVICE
		var currentDeviceMain WasteLibrary.RfidDeviceMainType
		currentDeviceMain.New()
		currentDeviceMain.SerialNumber = currentHttpHeader.DeviceNo
		data := url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentDeviceMain.ToString()},
		}
		resultVal = WasteLibrary.SaveStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		currentDeviceMain.DeviceId = WasteLibrary.StringIdToFloat64(resultVal.Retval.(string))
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
		currentHttpHeader.DataType = WasteLibrary.DATATYPE_RFID_BASE_DEVICE
		var currentDeviceBase WasteLibrary.RfidDeviceBaseType
		currentDeviceBase.New()
		currentDeviceBase.NewData = true
		currentDeviceBase.DeviceId = currentDeviceMain.DeviceId
		data = url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentDeviceBase.ToString()},
		}
		resultVal = WasteLibrary.SaveStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		currentDeviceBase.DeviceId = WasteLibrary.StringIdToFloat64(resultVal.Retval.(string))

		resultVal = currentDeviceBase.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		//DeviceStatu   RfidDeviceStatuType
		currentHttpHeader.DataType = WasteLibrary.DATATYPE_RFID_STATU_DEVICE
		var currentDeviceStatu WasteLibrary.RfidDeviceStatuType
		currentDeviceStatu.New()
		currentDeviceStatu.NewData = true
		currentDeviceStatu.DeviceId = currentDeviceMain.DeviceId
		data = url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentDeviceStatu.ToString()},
		}
		resultVal = WasteLibrary.SaveStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		currentDeviceStatu.DeviceId = WasteLibrary.StringIdToFloat64(resultVal.Retval.(string))

		resultVal = currentDeviceStatu.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		//DeviceGps     RfidDeviceGpsType
		currentHttpHeader.DataType = WasteLibrary.DATATYPE_RFID_GPS_DEVICE
		var currentDeviceGps WasteLibrary.RfidDeviceGpsType
		currentDeviceGps.New()
		currentDeviceGps.NewData = true
		currentDeviceGps.DeviceId = currentDeviceMain.DeviceId
		data = url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentDeviceGps.ToString()},
		}
		resultVal = WasteLibrary.SaveStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		currentDeviceGps.DeviceId = WasteLibrary.StringIdToFloat64(resultVal.Retval.(string))

		resultVal = currentDeviceGps.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		//DeviceAlarm   RfidDeviceAlarmType
		currentHttpHeader.DataType = WasteLibrary.DATATYPE_RFID_ALARM_DEVICE
		var currentDeviceAlarm WasteLibrary.RfidDeviceAlarmType
		currentDeviceAlarm.New()
		currentDeviceAlarm.NewData = true
		currentDeviceAlarm.DeviceId = currentDeviceMain.DeviceId
		data = url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentDeviceAlarm.ToString()},
		}
		resultVal = WasteLibrary.SaveStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		currentDeviceAlarm.DeviceId = WasteLibrary.StringIdToFloat64(resultVal.Retval.(string))

		resultVal = currentDeviceAlarm.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		//DeviceTherm   RfidDeviceThermType
		currentHttpHeader.DataType = WasteLibrary.DATATYPE_RFID_THERM_DEVICE
		var currentDeviceTherm WasteLibrary.RfidDeviceThermType
		currentDeviceTherm.New()
		currentDeviceTherm.NewData = true
		currentDeviceTherm.DeviceId = currentDeviceMain.DeviceId
		data = url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentDeviceTherm.ToString()},
		}
		resultVal = WasteLibrary.SaveStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		currentDeviceTherm.DeviceId = WasteLibrary.StringIdToFloat64(resultVal.Retval.(string))

		resultVal = currentDeviceTherm.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		//DeviceVersion RfidDeviceVersionType
		currentHttpHeader.DataType = WasteLibrary.DATATYPE_RFID_VERSION_DEVICE
		var currentDeviceVersion WasteLibrary.RfidDeviceVersionType
		currentDeviceVersion.New()
		currentDeviceVersion.NewData = true
		currentDeviceVersion.DeviceId = currentDeviceMain.DeviceId
		data = url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentDeviceVersion.ToString()},
		}
		resultVal = WasteLibrary.SaveStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		currentDeviceVersion.DeviceId = WasteLibrary.StringIdToFloat64(resultVal.Retval.(string))

		resultVal = currentDeviceVersion.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		//DeviceDetail  RfidDeviceDetailType
		currentHttpHeader.DataType = WasteLibrary.DATATYPE_RFID_DETAIL_DEVICE
		var currentDeviceDetail WasteLibrary.RfidDeviceDetailType
		currentDeviceDetail.New()
		currentDeviceDetail.NewData = true
		currentDeviceDetail.DeviceId = currentDeviceMain.DeviceId
		data = url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentDeviceDetail.ToString()},
		}
		resultVal = WasteLibrary.SaveStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		currentDeviceDetail.DeviceId = WasteLibrary.StringIdToFloat64(resultVal.Retval.(string))

		resultVal = currentDeviceDetail.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		//DeviceWorkHour RfidDeviceWorkHourType
		currentHttpHeader.DataType = WasteLibrary.DATATYPE_RFID_WORKHOUR_DEVICE
		var currentDeviceWorkHour WasteLibrary.RfidDeviceWorkHourType
		currentDeviceWorkHour.New()
		currentDeviceWorkHour.NewData = true
		currentDeviceWorkHour.DeviceId = currentDeviceMain.DeviceId
		data = url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentDeviceWorkHour.ToString()},
		}
		resultVal = WasteLibrary.SaveStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		currentDeviceWorkHour.DeviceId = WasteLibrary.StringIdToFloat64(resultVal.Retval.(string))

		resultVal = currentDeviceWorkHour.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		var customerDevices WasteLibrary.CustomerRfidDevicesType
		customerDevices.CustomerId = currentDeviceMain.CustomerId
		resultVal = customerDevices.GetByRedis()
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
		var currentUlt WasteLibrary.UltDeviceType = WasteLibrary.StringToUltDeviceType(req.FormValue(WasteLibrary.HTTP_DATA))
		//DeviceMainType    UltDeviceMainType
		currentHttpHeader.DataType = WasteLibrary.DATATYPE_ULT_MAIN_DEVICE
		var currentDeviceMain WasteLibrary.UltDeviceMainType
		currentDeviceMain.New()
		currentDeviceMain.SerialNumber = currentHttpHeader.DeviceNo
		data := url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentDeviceMain.ToString()},
		}
		resultVal = WasteLibrary.SaveStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		currentDeviceMain.DeviceId = WasteLibrary.StringIdToFloat64(resultVal.Retval.(string))
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
		currentHttpHeader.DataType = WasteLibrary.DATATYPE_ULT_BASE_DEVICE
		var currentDeviceBase WasteLibrary.UltDeviceBaseType
		currentDeviceBase.New()
		currentDeviceBase.NewData = true
		currentDeviceBase.DeviceId = currentDeviceMain.DeviceId
		currentDeviceBase.Imei = currentUlt.DeviceBase.Imei
		currentDeviceBase.Imsi = currentUlt.DeviceBase.Imsi
		data = url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentDeviceBase.ToString()},
		}
		resultVal = WasteLibrary.SaveStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		currentDeviceBase.DeviceId = WasteLibrary.StringIdToFloat64(resultVal.Retval.(string))

		resultVal = currentDeviceBase.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		//DeviceStatu   UltDeviceStatuType
		currentHttpHeader.DataType = WasteLibrary.DATATYPE_ULT_STATU_DEVICE
		var currentDeviceStatu WasteLibrary.UltDeviceStatuType
		currentDeviceStatu.New()
		currentDeviceStatu.NewData = true
		currentDeviceStatu.DeviceId = currentDeviceMain.DeviceId
		data = url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentDeviceStatu.ToString()},
		}
		resultVal = WasteLibrary.SaveStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		currentDeviceStatu.DeviceId = WasteLibrary.StringIdToFloat64(resultVal.Retval.(string))

		resultVal = currentDeviceStatu.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		//DeviceGps     UltDeviceGpsType
		currentHttpHeader.DataType = WasteLibrary.DATATYPE_ULT_GPS_DEVICE
		var currentDeviceGps WasteLibrary.UltDeviceGpsType
		currentDeviceGps.New()
		currentDeviceGps.NewData = true
		currentDeviceGps.DeviceId = currentDeviceMain.DeviceId
		data = url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentDeviceGps.ToString()},
		}
		resultVal = WasteLibrary.SaveStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		currentDeviceGps.DeviceId = WasteLibrary.StringIdToFloat64(resultVal.Retval.(string))

		resultVal = currentDeviceGps.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		//DeviceAlarm   UltDeviceAlarmType
		currentHttpHeader.DataType = WasteLibrary.DATATYPE_ULT_ALARM_DEVICE
		var currentDeviceAlarm WasteLibrary.UltDeviceAlarmType
		currentDeviceAlarm.New()
		currentDeviceAlarm.NewData = true
		currentDeviceAlarm.DeviceId = currentDeviceMain.DeviceId
		data = url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentDeviceAlarm.ToString()},
		}
		resultVal = WasteLibrary.SaveStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		currentDeviceAlarm.DeviceId = WasteLibrary.StringIdToFloat64(resultVal.Retval.(string))

		resultVal = currentDeviceAlarm.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		//DeviceTherm   UltDeviceThermType
		currentHttpHeader.DataType = WasteLibrary.DATATYPE_ULT_THERM_DEVICE
		var currentDeviceTherm WasteLibrary.UltDeviceThermType
		currentDeviceTherm.New()
		currentDeviceTherm.NewData = true
		currentDeviceTherm.DeviceId = currentDeviceMain.DeviceId
		data = url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentDeviceTherm.ToString()},
		}
		resultVal = WasteLibrary.SaveStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		currentDeviceTherm.DeviceId = WasteLibrary.StringIdToFloat64(resultVal.Retval.(string))

		resultVal = currentDeviceTherm.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		//DeviceVersion UltDeviceVersionType
		currentHttpHeader.DataType = WasteLibrary.DATATYPE_ULT_VERSION_DEVICE
		var currentDeviceVersion WasteLibrary.UltDeviceVersionType
		currentDeviceVersion.New()
		currentDeviceVersion.NewData = true
		currentDeviceVersion.DeviceId = currentDeviceMain.DeviceId
		data = url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentDeviceVersion.ToString()},
		}
		resultVal = WasteLibrary.SaveStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		currentDeviceVersion.DeviceId = WasteLibrary.StringIdToFloat64(resultVal.Retval.(string))

		resultVal = currentDeviceVersion.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		//DeviceBattery  UltDeviceBatteryType
		currentHttpHeader.DataType = WasteLibrary.DATATYPE_ULT_BATTERY_DEVICE
		var currentDeviceBattery WasteLibrary.UltDeviceBatteryType
		currentDeviceBattery.New()
		currentDeviceBattery.NewData = true
		currentDeviceBattery.DeviceId = currentDeviceMain.DeviceId
		data = url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentDeviceBattery.ToString()},
		}
		resultVal = WasteLibrary.SaveStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		currentDeviceBattery.DeviceId = WasteLibrary.StringIdToFloat64(resultVal.Retval.(string))

		resultVal = currentDeviceBattery.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		//DeviceSens  UltDeviceSensType
		currentHttpHeader.DataType = WasteLibrary.DATATYPE_ULT_SENS_DEVICE
		var currentDeviceSens WasteLibrary.UltDeviceSensType
		currentDeviceSens.New()
		currentDeviceSens.NewData = true
		currentDeviceSens.DeviceId = currentDeviceMain.DeviceId
		data = url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentDeviceSens.ToString()},
		}
		resultVal = WasteLibrary.SaveStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		currentDeviceSens.DeviceId = WasteLibrary.StringIdToFloat64(resultVal.Retval.(string))

		resultVal = currentDeviceSens.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		var customerDevices WasteLibrary.CustomerUltDevicesType
		customerDevices.CustomerId = currentDeviceMain.CustomerId
		resultVal = customerDevices.GetByRedis()
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
		currentHttpHeader.DataType = WasteLibrary.DATATYPE_RECY_MAIN_DEVICE
		var currentDeviceMain WasteLibrary.RecyDeviceMainType
		currentDeviceMain.New()
		currentDeviceMain.SerialNumber = currentHttpHeader.DeviceNo
		data := url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentDeviceMain.ToString()},
		}
		resultVal = WasteLibrary.SaveStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		currentDeviceMain.DeviceId = WasteLibrary.StringIdToFloat64(resultVal.Retval.(string))
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
		currentHttpHeader.DataType = WasteLibrary.DATATYPE_RECY_BASE_DEVICE
		var currentDeviceBase WasteLibrary.RecyDeviceBaseType
		currentDeviceBase.New()
		currentDeviceBase.NewData = true
		currentDeviceBase.DeviceId = currentDeviceMain.DeviceId
		data = url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentDeviceBase.ToString()},
		}
		resultVal = WasteLibrary.SaveStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		currentDeviceBase.DeviceId = WasteLibrary.StringIdToFloat64(resultVal.Retval.(string))

		resultVal = currentDeviceBase.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		//DeviceStatu   RecyDeviceStatuType
		currentHttpHeader.DataType = WasteLibrary.DATATYPE_RECY_STATU_DEVICE
		var currentDeviceStatu WasteLibrary.RecyDeviceStatuType
		currentDeviceStatu.New()
		currentDeviceStatu.NewData = true
		currentDeviceStatu.DeviceId = currentDeviceMain.DeviceId
		data = url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentDeviceStatu.ToString()},
		}
		resultVal = WasteLibrary.SaveStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		currentDeviceStatu.DeviceId = WasteLibrary.StringIdToFloat64(resultVal.Retval.(string))

		resultVal = currentDeviceStatu.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		//DeviceGps     RecyDeviceGpsType
		currentHttpHeader.DataType = WasteLibrary.DATATYPE_RECY_GPS_DEVICE
		var currentDeviceGps WasteLibrary.RecyDeviceGpsType
		currentDeviceGps.New()
		currentDeviceGps.NewData = true
		currentDeviceGps.DeviceId = currentDeviceMain.DeviceId
		data = url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentDeviceGps.ToString()},
		}
		resultVal = WasteLibrary.SaveStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		currentDeviceGps.DeviceId = WasteLibrary.StringIdToFloat64(resultVal.Retval.(string))

		resultVal = currentDeviceGps.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		//DeviceAlarm   RecyDeviceAlarmType
		currentHttpHeader.DataType = WasteLibrary.DATATYPE_RECY_ALARM_DEVICE
		var currentDeviceAlarm WasteLibrary.RecyDeviceAlarmType
		currentDeviceAlarm.New()
		currentDeviceAlarm.NewData = true
		currentDeviceAlarm.DeviceId = currentDeviceMain.DeviceId
		data = url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentDeviceAlarm.ToString()},
		}
		resultVal = WasteLibrary.SaveStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		currentDeviceAlarm.DeviceId = WasteLibrary.StringIdToFloat64(resultVal.Retval.(string))

		resultVal = currentDeviceAlarm.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		//DeviceTherm   RecyDeviceThermType
		currentHttpHeader.DataType = WasteLibrary.DATATYPE_RECY_THERM_DEVICE
		var currentDeviceTherm WasteLibrary.RecyDeviceThermType
		currentDeviceTherm.New()
		currentDeviceTherm.NewData = true
		currentDeviceTherm.DeviceId = currentDeviceMain.DeviceId
		data = url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentDeviceTherm.ToString()},
		}
		resultVal = WasteLibrary.SaveStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		currentDeviceTherm.DeviceId = WasteLibrary.StringIdToFloat64(resultVal.Retval.(string))

		resultVal = currentDeviceTherm.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		//DeviceVersion RecyDeviceVersionType
		currentHttpHeader.DataType = WasteLibrary.DATATYPE_RECY_VERSION_DEVICE
		var currentDeviceVersion WasteLibrary.RecyDeviceVersionType
		currentDeviceVersion.New()
		currentDeviceVersion.NewData = true
		currentDeviceVersion.DeviceId = currentDeviceMain.DeviceId
		data = url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentDeviceVersion.ToString()},
		}
		resultVal = WasteLibrary.SaveStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		currentDeviceVersion.DeviceId = WasteLibrary.StringIdToFloat64(resultVal.Retval.(string))

		resultVal = currentDeviceVersion.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		//DeviceDetail  RecyDeviceDetailType
		currentHttpHeader.DataType = WasteLibrary.DATATYPE_RECY_DETAIL_DEVICE
		var currentDeviceDetail WasteLibrary.RecyDeviceDetailType
		currentDeviceDetail.New()
		currentDeviceDetail.NewData = true
		currentDeviceDetail.DeviceId = currentDeviceMain.DeviceId
		data = url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentDeviceDetail.ToString()},
		}
		resultVal = WasteLibrary.SaveStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		currentDeviceDetail.DeviceId = WasteLibrary.StringIdToFloat64(resultVal.Retval.(string))

		resultVal = currentDeviceDetail.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		var customerDevices WasteLibrary.CustomerRecyDevicesType
		customerDevices.CustomerId = currentDeviceMain.CustomerId
		resultVal = customerDevices.GetByRedis()
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
	var currentData WasteLibrary.TagType = WasteLibrary.StringToTagType(req.FormValue(WasteLibrary.HTTP_DATA))
	if currentData.TagMain.CustomerId > 1 {
		//TagMainType    TagMainType
		currentHttpHeader.DataType = WasteLibrary.DATATYPE_TAG_MAIN
		data := url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentData.TagMain.ToString()},
		}
		resultVal = WasteLibrary.SaveStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		currentData.TagId = WasteLibrary.StringIdToFloat64(resultVal.Retval.(string))
		currentData.TagMain.TagId = currentData.TagId

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
		currentHttpHeader.DataType = WasteLibrary.DATATYPE_TAG_BASE
		var currentTagBase WasteLibrary.TagBaseType
		currentTagBase.New()
		currentTagBase.NewData = true
		currentTagBase.TagId = currentData.TagId
		data = url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentTagBase.ToString()},
		}
		resultVal = WasteLibrary.SaveStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		currentTagBase.TagId = WasteLibrary.StringIdToFloat64(resultVal.Retval.(string))

		resultVal = currentTagBase.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		//TagStatu   TagStatuType
		currentHttpHeader.DataType = WasteLibrary.DATATYPE_TAG_STATU
		var currentTagStatu WasteLibrary.TagStatuType
		currentTagStatu.New()
		currentTagStatu.NewData = true
		currentTagStatu.TagId = currentData.TagId
		data = url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentTagStatu.ToString()},
		}
		resultVal = WasteLibrary.SaveStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		currentTagStatu.TagId = WasteLibrary.StringIdToFloat64(resultVal.Retval.(string))

		resultVal = currentTagStatu.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		//TagGps     TagGpsType
		currentHttpHeader.DataType = WasteLibrary.DATATYPE_TAG_GPS
		var currentTagGps WasteLibrary.TagGpsType
		currentTagGps.New()
		currentTagGps.NewData = true
		currentTagGps.TagId = currentData.TagId
		data = url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentTagGps.ToString()},
		}
		resultVal = WasteLibrary.SaveStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		currentTagGps.TagId = WasteLibrary.StringIdToFloat64(resultVal.Retval.(string))

		resultVal = currentTagGps.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		//TagReader   TagReaderType
		currentHttpHeader.DataType = WasteLibrary.DATATYPE_TAG_READER
		var currentTagReader WasteLibrary.TagReaderType
		currentTagReader.New()
		currentTagReader.NewData = true
		currentTagReader.TagId = currentData.TagId
		data = url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentTagReader.ToString()},
		}
		resultVal = WasteLibrary.SaveStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		currentTagReader.TagId = WasteLibrary.StringIdToFloat64(resultVal.Retval.(string))

		resultVal = currentTagReader.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		var customerTags WasteLibrary.CustomerTagsType
		customerTags.CustomerId = currentData.TagMain.CustomerId
		resultVal = customerTags.GetByRedis()
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
