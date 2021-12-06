package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/url"
)

//RfidDeviceWorkHourType
type RfidDeviceWorkHourType struct {
	DeviceId         float64
	WorkCount        int
	Work1StartHour   int
	Work1StartMinute int
	Work1AddMinute   int
	Work2StartHour   int
	Work2StartMinute int
	Work2AddMinute   int
	Work3StartHour   int
	Work3StartMinute int
	Work3AddMinute   int
	NewData          bool
}

//New
func (res *RfidDeviceWorkHourType) New() {
	res.DeviceId = 0
	res.WorkCount = 1
	res.Work1StartHour = 06
	res.Work1StartMinute = 0
	res.Work1AddMinute = 510
	res.Work2StartHour = 0
	res.Work2StartMinute = 0
	res.Work2AddMinute = 0
	res.Work3StartHour = 0
	res.Work3StartMinute = 0
	res.Work3AddMinute = 0
	res.NewData = false
}

//GetByRedis
func (res *RfidDeviceWorkHourType) GetByRedis(dbIndex string) ResultType {

	var resultVal ResultType
	resultVal = GetRedisForStoreApi(dbIndex, REDIS_RFID_WORKHOUR, res.ToIdString())
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
func (res *RfidDeviceWorkHourType) SaveToRedis() ResultType {
	var resultVal ResultType
	resultVal = SaveRedisForStoreApi(REDIS_RFID_WORKHOUR, res.ToIdString(), res.ToString())
	return resultVal
}

//SaveToDb
func (res *RfidDeviceWorkHourType) SaveToDb() ResultType {
	var resultVal ResultType
	var currentHttpHeader HttpClientHeaderType
	currentHttpHeader.New()
	currentHttpHeader.DataType = DATATYPE_RFID_WORKHOUR

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

//ToId String
func (res *RfidDeviceWorkHourType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToByte
func (res *RfidDeviceWorkHourType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res *RfidDeviceWorkHourType) ToString() string {
	return string(res.ToByte())

}

//Byte To RfidDeviceWorkHourType
func ByteToRfidDeviceWorkHourType(retByte []byte) RfidDeviceWorkHourType {
	var retVal RfidDeviceWorkHourType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To RfidDeviceWorkHourType
func StringToRfidDeviceWorkHourType(retStr string) RfidDeviceWorkHourType {
	return ByteToRfidDeviceWorkHourType([]byte(retStr))
}

//ByteToType
func (res *RfidDeviceWorkHourType) ByteToType(retByte []byte) {
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *RfidDeviceWorkHourType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}

//SelectSQL
func (res *RfidDeviceWorkHourType) SelectSQL() string {
	return fmt.Sprintf(`SELECT WorkCount,
	 Work1StartHour,Work1StartMinute,Work1AddMinute,
	 Work2StartHour,Work2StartMinute,Work2AddMinute,
	 Work3StartHour,Work3StartMinute,Work3AddMinute
	 FROM public.`+DATATYPE_RFID_WORKHOUR+` 
	 WHERE DeviceId=%f ;`, res.DeviceId)
}

//InsertSQL
func (res *RfidDeviceWorkHourType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.`+DATATYPE_RFID_WORKHOUR+`  (DeviceId,WorkCount,
	 Work1StartHour,Work1StartMinute,Work1AddMinute,
	 Work2StartHour,Work2StartMinute,Work2AddMinute,
	 Work3StartHour,Work3StartMinute,Work3AddMinute) 
	 VALUES (%f,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d) 
	 RETURNING DeviceId;`, res.DeviceId, res.WorkCount,
		res.Work1StartHour, res.Work1StartMinute, res.Work1AddMinute,
		res.Work2StartHour, res.Work2StartMinute, res.Work2AddMinute,
		res.Work3StartHour, res.Work3StartMinute, res.Work3AddMinute)
}

//UpdateSQL
func (res *RfidDeviceWorkHourType) UpdateSQL() string {

	return fmt.Sprintf(`UPDATE public.`+DATATYPE_RFID_WORKHOUR+`  
	  SET WorkCount=%d,
	  Work1StartHour=%d,Work1StartMinute=%d,Work1AddMinute=%d,
	  Work2StartHour=%d,Work2StartMinute=%d,Work2AddMinute=%d,
	  Work3StartHour=%d,Work3StartMinute=%d,Work3AddMinute=%d
	  WHERE DeviceId=%f  
	  RETURNING DeviceId;`, res.WorkCount,
		res.Work1StartHour, res.Work1StartMinute, res.Work1AddMinute,
		res.Work2StartHour, res.Work2StartMinute, res.Work2AddMinute,
		res.Work3StartHour, res.Work3StartMinute, res.Work3AddMinute,
		res.DeviceId)
}

//SelectWithDb
func (res *RfidDeviceWorkHourType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(
		&res.WorkCount,
		&res.Work1StartHour, &res.Work1StartMinute, &res.Work1AddMinute,
		&res.Work2StartHour, &res.Work2StartMinute, &res.Work2AddMinute,
		&res.Work3StartHour, &res.Work3StartMinute, &res.Work3AddMinute)
	return errDb
}

//CreateDb
func (res *RfidDeviceWorkHourType) CreateDb(currentDb *sql.DB) {
	createSQL := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS ` + DATATYPE_RFID_WORKHOUR + `  ( 
	DataId serial PRIMARY KEY,
	DeviceId INT NOT NULL DEFAULT -1,
	WorkCount INT NOT NULL DEFAULT 1,
	Work1StartHour  INT NOT NULL DEFAULT 6,
	Work1StartMinute  INT NOT NULL DEFAULT 0,
	Work1AddMinute  INT NOT NULL DEFAULT 510,
	Work2StartHour  INT NOT NULL DEFAULT 0,
	Work2StartMinute  INT NOT NULL DEFAULT 0,
	Work2AddMinute  INT NOT NULL DEFAULT 0,
	Work3StartHour  INT NOT NULL DEFAULT 0,
	Work3StartMinute  INT NOT NULL DEFAULT 0,
	Work3AddMinute  INT NOT NULL DEFAULT 0,
	CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);`)
	_, err := currentDb.Exec(createSQL)
	LogErr(err)
}
