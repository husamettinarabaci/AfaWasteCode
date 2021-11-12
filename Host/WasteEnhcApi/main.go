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
		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_SERIAL_RFID_DEVICE, currentHttpHeader.DeviceNo, currentDeviceMain.ToIdString())
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		data = url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentDeviceMain.ToString()},
		}
		resultVal = WasteLibrary.GetStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_GET
			w.Write(resultVal.ToByte())

			return
		}
		currentDeviceMain = WasteLibrary.StringToRfidDeviceMainType(resultVal.Retval.(string))

		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_RFID_MAIN_DEVICES, currentDeviceMain.ToIdString(), currentDeviceMain.ToString())
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
		data = url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentDeviceBase.ToString()},
		}
		resultVal = WasteLibrary.GetStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_GET
			w.Write(resultVal.ToByte())

			return
		}
		currentDeviceBase = WasteLibrary.StringToRfidDeviceBaseType(resultVal.Retval.(string))

		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_RFID_BASE_DEVICES, currentDeviceBase.ToIdString(), currentDeviceBase.ToString())
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
		data = url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentDeviceStatu.ToString()},
		}
		resultVal = WasteLibrary.GetStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_GET
			w.Write(resultVal.ToByte())

			return
		}
		currentDeviceStatu = WasteLibrary.StringToRfidDeviceStatuType(resultVal.Retval.(string))

		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_RFID_STATU_DEVICES, currentDeviceStatu.ToIdString(), currentDeviceStatu.ToString())
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
		data = url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentDeviceGps.ToString()},
		}
		resultVal = WasteLibrary.GetStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_GET
			w.Write(resultVal.ToByte())

			return
		}
		currentDeviceGps = WasteLibrary.StringToRfidDeviceGpsType(resultVal.Retval.(string))

		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_RFID_GPS_DEVICES, currentDeviceGps.ToIdString(), currentDeviceGps.ToString())
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
		data = url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentDeviceAlarm.ToString()},
		}
		resultVal = WasteLibrary.GetStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_GET
			w.Write(resultVal.ToByte())

			return
		}
		currentDeviceAlarm = WasteLibrary.StringToRfidDeviceAlarmType(resultVal.Retval.(string))

		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_RFID_ALARM_DEVICES, currentDeviceAlarm.ToIdString(), currentDeviceAlarm.ToString())
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
		data = url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentDeviceTherm.ToString()},
		}
		resultVal = WasteLibrary.GetStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_GET
			w.Write(resultVal.ToByte())

			return
		}
		currentDeviceTherm = WasteLibrary.StringToRfidDeviceThermType(resultVal.Retval.(string))

		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_RFID_THERM_DEVICES, currentDeviceTherm.ToIdString(), currentDeviceTherm.ToString())
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
		data = url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentDeviceVersion.ToString()},
		}
		resultVal = WasteLibrary.GetStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_GET
			w.Write(resultVal.ToByte())

			return
		}
		currentDeviceVersion = WasteLibrary.StringToRfidDeviceVersionType(resultVal.Retval.(string))

		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_RFID_VERSION_DEVICES, currentDeviceVersion.ToIdString(), currentDeviceVersion.ToString())
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
		data = url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentDeviceDetail.ToString()},
		}
		resultVal = WasteLibrary.GetStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_GET
			w.Write(resultVal.ToByte())

			return
		}
		currentDeviceDetail = WasteLibrary.StringToRfidDeviceDetailType(resultVal.Retval.(string))

		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_RFID_DETAIL_DEVICES, currentDeviceDetail.ToIdString(), currentDeviceDetail.ToString())
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
		data = url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentDeviceWorkHour.ToString()},
		}
		resultVal = WasteLibrary.GetStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_GET
			w.Write(resultVal.ToByte())

			return
		}
		currentDeviceWorkHour = WasteLibrary.StringToRfidDeviceWorkHourType(resultVal.Retval.(string))

		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_RFID_WORKHOUR_DEVICES, currentDeviceWorkHour.ToIdString(), currentDeviceWorkHour.ToString())
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_RFID_DEVICES, currentDeviceMain.ToCustomerIdString())
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_GET
			w.Write(resultVal.ToByte())

			return
		}

		var customerDevices WasteLibrary.CustomerRfidDevicesType = WasteLibrary.StringToCustomerRfidDevicesType(resultVal.Retval.(string))
		customerDevices.Devices[currentDeviceMain.ToIdString()] = currentDeviceMain.DeviceId
		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_RFID_DEVICES, customerDevices.ToIdString(), customerDevices.ToString())
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
		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_SERIAL_ULT_DEVICE, currentHttpHeader.DeviceNo, currentDeviceMain.ToIdString())
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		data = url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentDeviceMain.ToString()},
		}
		resultVal = WasteLibrary.GetStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_GET
			w.Write(resultVal.ToByte())

			return
		}
		currentDeviceMain = WasteLibrary.StringToUltDeviceMainType(resultVal.Retval.(string))

		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_ULT_MAIN_DEVICES, currentDeviceMain.ToIdString(), currentDeviceMain.ToString())
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
		data = url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentDeviceBase.ToString()},
		}
		resultVal = WasteLibrary.GetStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_GET
			w.Write(resultVal.ToByte())

			return
		}
		currentDeviceBase = WasteLibrary.StringToUltDeviceBaseType(resultVal.Retval.(string))

		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_ULT_BASE_DEVICES, currentDeviceBase.ToIdString(), currentDeviceBase.ToString())
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
		data = url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentDeviceStatu.ToString()},
		}
		resultVal = WasteLibrary.GetStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_GET
			w.Write(resultVal.ToByte())

			return
		}
		currentDeviceStatu = WasteLibrary.StringToUltDeviceStatuType(resultVal.Retval.(string))

		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_ULT_STATU_DEVICES, currentDeviceStatu.ToIdString(), currentDeviceStatu.ToString())
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
		data = url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentDeviceGps.ToString()},
		}
		resultVal = WasteLibrary.GetStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_GET
			w.Write(resultVal.ToByte())

			return
		}
		currentDeviceGps = WasteLibrary.StringToUltDeviceGpsType(resultVal.Retval.(string))

		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_ULT_GPS_DEVICES, currentDeviceGps.ToIdString(), currentDeviceGps.ToString())
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
		data = url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentDeviceAlarm.ToString()},
		}
		resultVal = WasteLibrary.GetStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_GET
			w.Write(resultVal.ToByte())

			return
		}
		currentDeviceAlarm = WasteLibrary.StringToUltDeviceAlarmType(resultVal.Retval.(string))

		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_ULT_ALARM_DEVICES, currentDeviceAlarm.ToIdString(), currentDeviceAlarm.ToString())
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
		data = url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentDeviceTherm.ToString()},
		}
		resultVal = WasteLibrary.GetStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_GET
			w.Write(resultVal.ToByte())

			return
		}
		currentDeviceTherm = WasteLibrary.StringToUltDeviceThermType(resultVal.Retval.(string))

		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_ULT_THERM_DEVICES, currentDeviceTherm.ToIdString(), currentDeviceTherm.ToString())
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
		data = url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentDeviceVersion.ToString()},
		}
		resultVal = WasteLibrary.GetStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_GET
			w.Write(resultVal.ToByte())

			return
		}
		currentDeviceVersion = WasteLibrary.StringToUltDeviceVersionType(resultVal.Retval.(string))

		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_ULT_VERSION_DEVICES, currentDeviceVersion.ToIdString(), currentDeviceVersion.ToString())
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
		data = url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentDeviceBattery.ToString()},
		}
		resultVal = WasteLibrary.GetStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_GET
			w.Write(resultVal.ToByte())

			return
		}
		currentDeviceBattery = WasteLibrary.StringToUltDeviceBatteryType(resultVal.Retval.(string))

		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_ULT_BATTERY_DEVICES, currentDeviceBattery.ToIdString(), currentDeviceBattery.ToString())
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
		data = url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentDeviceSens.ToString()},
		}
		resultVal = WasteLibrary.GetStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_GET
			w.Write(resultVal.ToByte())

			return
		}
		currentDeviceSens = WasteLibrary.StringToUltDeviceSensType(resultVal.Retval.(string))

		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_ULT_SENS_DEVICES, currentDeviceSens.ToIdString(), currentDeviceSens.ToString())
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_ULT_DEVICES, currentDeviceMain.ToCustomerIdString())
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_GET
			w.Write(resultVal.ToByte())

			return
		}

		var customerDevices WasteLibrary.CustomerUltDevicesType = WasteLibrary.StringToCustomerUltDevicesType(resultVal.Retval.(string))
		customerDevices.Devices[currentDeviceMain.ToIdString()] = currentDeviceMain.DeviceId
		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_ULT_DEVICES, customerDevices.ToIdString(), customerDevices.ToString())
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
		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_SERIAL_RECY_DEVICE, currentHttpHeader.DeviceNo, currentDeviceMain.ToIdString())
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		data = url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentDeviceMain.ToString()},
		}
		resultVal = WasteLibrary.GetStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_GET
			w.Write(resultVal.ToByte())

			return
		}
		currentDeviceMain = WasteLibrary.StringToRecyDeviceMainType(resultVal.Retval.(string))

		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_RECY_MAIN_DEVICES, currentDeviceMain.ToIdString(), currentDeviceMain.ToString())
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
		data = url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentDeviceBase.ToString()},
		}
		resultVal = WasteLibrary.GetStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_GET
			w.Write(resultVal.ToByte())

			return
		}
		currentDeviceBase = WasteLibrary.StringToRecyDeviceBaseType(resultVal.Retval.(string))

		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_RECY_BASE_DEVICES, currentDeviceBase.ToIdString(), currentDeviceBase.ToString())
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
		data = url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentDeviceStatu.ToString()},
		}
		resultVal = WasteLibrary.GetStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_GET
			w.Write(resultVal.ToByte())

			return
		}
		currentDeviceStatu = WasteLibrary.StringToRecyDeviceStatuType(resultVal.Retval.(string))

		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_RECY_STATU_DEVICES, currentDeviceStatu.ToIdString(), currentDeviceStatu.ToString())
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
		data = url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentDeviceGps.ToString()},
		}
		resultVal = WasteLibrary.GetStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_GET
			w.Write(resultVal.ToByte())

			return
		}
		currentDeviceGps = WasteLibrary.StringToRecyDeviceGpsType(resultVal.Retval.(string))

		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_RECY_GPS_DEVICES, currentDeviceGps.ToIdString(), currentDeviceGps.ToString())
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
		data = url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentDeviceAlarm.ToString()},
		}
		resultVal = WasteLibrary.GetStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_GET
			w.Write(resultVal.ToByte())

			return
		}
		currentDeviceAlarm = WasteLibrary.StringToRecyDeviceAlarmType(resultVal.Retval.(string))

		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_RECY_ALARM_DEVICES, currentDeviceAlarm.ToIdString(), currentDeviceAlarm.ToString())
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
		data = url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentDeviceTherm.ToString()},
		}
		resultVal = WasteLibrary.GetStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_GET
			w.Write(resultVal.ToByte())

			return
		}
		currentDeviceTherm = WasteLibrary.StringToRecyDeviceThermType(resultVal.Retval.(string))

		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_RECY_THERM_DEVICES, currentDeviceTherm.ToIdString(), currentDeviceTherm.ToString())
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
		data = url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentDeviceVersion.ToString()},
		}
		resultVal = WasteLibrary.GetStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_GET
			w.Write(resultVal.ToByte())

			return
		}
		currentDeviceVersion = WasteLibrary.StringToRecyDeviceVersionType(resultVal.Retval.(string))

		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_RECY_VERSION_DEVICES, currentDeviceVersion.ToIdString(), currentDeviceVersion.ToString())
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
		data = url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentDeviceDetail.ToString()},
		}
		resultVal = WasteLibrary.GetStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_GET
			w.Write(resultVal.ToByte())

			return
		}
		currentDeviceDetail = WasteLibrary.StringToRecyDeviceDetailType(resultVal.Retval.(string))

		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_RECY_DETAIL_DEVICES, currentDeviceDetail.ToIdString(), currentDeviceDetail.ToString())
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_RECY_DEVICES, currentDeviceMain.ToCustomerIdString())
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_GET
			w.Write(resultVal.ToByte())

			return
		}

		var customerDevices WasteLibrary.CustomerRecyDevicesType = WasteLibrary.StringToCustomerRecyDevicesType(resultVal.Retval.(string))
		customerDevices.Devices[currentDeviceMain.ToIdString()] = currentDeviceMain.DeviceId
		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_RECY_DEVICES, customerDevices.ToIdString(), customerDevices.ToString())
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
		data = url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentData.TagMain.ToString()},
		}
		resultVal = WasteLibrary.GetStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_GET
			w.Write(resultVal.ToByte())

			return
		}
		var currentDataMain WasteLibrary.TagMainType = WasteLibrary.StringToTagMainType(resultVal.Retval.(string))

		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_TAG_MAINS, currentDataMain.ToIdString(), currentDataMain.ToString())
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_TAG_EPC, currentDataMain.Epc, currentDataMain.ToIdString())
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
		data = url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentTagBase.ToString()},
		}
		resultVal = WasteLibrary.GetStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_GET
			w.Write(resultVal.ToByte())

			return
		}
		currentTagBase = WasteLibrary.StringToTagBaseType(resultVal.Retval.(string))

		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_TAG_BASES, currentTagBase.ToIdString(), currentTagBase.ToString())
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
		data = url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentTagStatu.ToString()},
		}
		resultVal = WasteLibrary.GetStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_GET
			w.Write(resultVal.ToByte())

			return
		}
		currentTagStatu = WasteLibrary.StringToTagStatuType(resultVal.Retval.(string))

		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_TAG_STATUS, currentTagStatu.ToIdString(), currentTagStatu.ToString())
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
		data = url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentTagGps.ToString()},
		}
		resultVal = WasteLibrary.GetStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_GET
			w.Write(resultVal.ToByte())

			return
		}
		currentTagGps = WasteLibrary.StringToTagGpsType(resultVal.Retval.(string))

		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_TAG_GPSES, currentTagGps.ToIdString(), currentTagGps.ToString())
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
		data = url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentTagReader.ToString()},
		}
		resultVal = WasteLibrary.GetStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_GET
			w.Write(resultVal.ToByte())

			return
		}
		currentTagReader = WasteLibrary.StringToTagReaderType(resultVal.Retval.(string))

		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_TAG_READERS, currentTagReader.ToIdString(), currentTagReader.ToString())
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_TAGS, currentDataMain.ToCustomerIdString())
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_GET
			w.Write(resultVal.ToByte())

			return
		}

		var customerTags WasteLibrary.CustomerTagsType = WasteLibrary.StringToCustomerTagsType(resultVal.Retval.(string))
		customerTags.Tags[currentDataMain.ToIdString()] = currentDataMain.TagId
		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_TAGS, customerTags.ToIdString(), customerTags.ToString())
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
