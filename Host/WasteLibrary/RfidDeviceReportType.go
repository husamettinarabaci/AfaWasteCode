package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/url"
)

//RfidDeviceReportType
type RfidDeviceReportType struct {
	DeviceId      float64
	DailyCapacity float64
	DailyKm       float64
	NewData       bool
}

//New
func (res *RfidDeviceReportType) New() {
	res.DeviceId = 0
	res.DailyCapacity = 0
	res.DailyKm = 0
	res.NewData = false
}

//GetByRedis
func (res *RfidDeviceReportType) GetByRedis(dbIndex string) ResultType {

	var resultVal ResultType
	resultVal = GetRedisForStoreApi(dbIndex, REDIS_RFID_REPORT_DEVICES, res.ToIdString())
	if resultVal.Result == RESULT_OK {
		res.StringToType(resultVal.Retval.(string))
		res.NewData = false
	} else {
		return resultVal
	}

	resultVal.Retval = res.ToString()
	return resultVal
}

//SaveToRedis
func (res *RfidDeviceReportType) SaveToRedis() ResultType {
	var resultVal ResultType
	resultVal = SaveRedisForStoreApi(REDIS_RFID_REPORT_DEVICES, res.ToIdString(), res.ToString())
	return resultVal
}

//SaveToDb
func (res *RfidDeviceReportType) SaveToDb() ResultType {
	var resultVal ResultType
	var currentHttpHeader HttpClientHeaderType
	currentHttpHeader.New()
	currentHttpHeader.DataType = DATATYPE_RFID_REPORT_DEVICE

	data := url.Values{
		HTTP_HEADER: {currentHttpHeader.ToString()},
		HTTP_DATA:   {res.ToString()},
	}
	resultVal = SaveStaticDbMainForStoreApi(data)
	if resultVal.Result == RESULT_OK {
		res.DeviceId = StringIdToFloat64(resultVal.Retval.(string))
		resultVal.Retval = res.ToString()
	}

	return resultVal
}

//SaveToReaderDb
func (res *RfidDeviceReportType) SaveToReaderDb() ResultType {
	var resultVal ResultType
	var currentHttpHeader HttpClientHeaderType
	currentHttpHeader.New()
	currentHttpHeader.DataType = DATATYPE_RFID_REPORT_DEVICE

	data := url.Values{
		HTTP_HEADER: {currentHttpHeader.ToString()},
		HTTP_DATA:   {res.ToString()},
	}
	resultVal = SaveReaderDbMainForStoreApi(data)
	if resultVal.Result == RESULT_OK {
		res.DeviceId = StringIdToFloat64(resultVal.Retval.(string))
		resultVal.Retval = res.ToString()
	}

	return resultVal
}

//ToId String
func (res *RfidDeviceReportType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToByte
func (res *RfidDeviceReportType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res *RfidDeviceReportType) ToString() string {
	return string(res.ToByte())

}

//Byte To RfidDeviceReportType
func ByteToRfidDeviceReportType(retByte []byte) RfidDeviceReportType {
	var retVal RfidDeviceReportType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To RfidDeviceReportType
func StringToRfidDeviceReportType(retStr string) RfidDeviceReportType {
	return ByteToRfidDeviceReportType([]byte(retStr))
}

//ByteToType
func (res *RfidDeviceReportType) ByteToType(retByte []byte) {
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *RfidDeviceReportType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}

//SelectSQL
func (res *RfidDeviceReportType) SelectSQL() string {
	return fmt.Sprintf(`SELECT DailyCapacity,DailyKm
	 FROM public.`+DATATYPE_RFID_REPORT_DEVICE+` 
	 WHERE DeviceId=%f ;`, res.DeviceId)
}

//InsertSQL
func (res *RfidDeviceReportType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.`+DATATYPE_RFID_REPORT_DEVICE+`  (DeviceId,DailyCapacity,DailyKm) 
	  VALUES (%f,%f,%f) 
	  RETURNING DeviceId;`, res.DeviceId, res.DailyCapacity, res.DailyKm)
}

//UpdateSQL
func (res *RfidDeviceReportType) UpdateSQL() string {
	return fmt.Sprintf(`UPDATE public.`+DATATYPE_RFID_REPORT_DEVICE+`  
	  SET DailyCapacity=%f,DailyKm=%f 
	  WHERE DeviceId=%f  
	  RETURNING DeviceId;`,
		res.DailyCapacity,
		res.DailyKm,
		res.DeviceId)
}

//SelectWithDb
func (res *RfidDeviceReportType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(
		&res.DailyCapacity,
		&res.DailyKm)
	return errDb
}
