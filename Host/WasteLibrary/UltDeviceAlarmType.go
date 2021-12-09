package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/url"
)

//UltDeviceAlarmType
type UltDeviceAlarmType struct {
	DeviceId    float64
	AlarmStatus string
	AlarmTime   string
	AlarmType   string
	Alarm       string
	NewData     bool
}

//New
func (res *UltDeviceAlarmType) New() {
	res.DeviceId = 0
	res.AlarmStatus = ALARMSTATU_NONE
	res.AlarmTime = GetTime()
	res.AlarmType = ALARMTYPE_NONE
	res.Alarm = ""
	res.NewData = false
}

//GetByRedis
func (res *UltDeviceAlarmType) GetByRedis(dbIndex string) ResultType {
	resultVal := GetRedisForStoreApi(dbIndex, REDIS_ULT_ALARM, res.ToIdString())
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
func (res *UltDeviceAlarmType) SaveToRedis() ResultType {
	resultVal := SaveRedisForStoreApi(REDIS_ULT_ALARM, res.ToIdString(), res.ToString())
	return resultVal
}

//SaveToDb
func (res *UltDeviceAlarmType) SaveToDb() ResultType {
	var resultVal ResultType
	var currentHttpHeader HttpClientHeaderType
	currentHttpHeader.New()
	currentHttpHeader.DataType = DATATYPE_ULT_ALARM

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
func (res *UltDeviceAlarmType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToByte
func (res *UltDeviceAlarmType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res *UltDeviceAlarmType) ToString() string {
	return string(res.ToByte())

}

//ByteToType
func (res *UltDeviceAlarmType) ByteToType(retByte []byte) {
	res.New()
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *UltDeviceAlarmType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}

//SelectSQL
func (res *UltDeviceAlarmType) SelectSQL() string {
	return fmt.Sprintf(`SELECT AlarmStatus,AlarmTime,AlarmType,Alarm
	 FROM public.`+DATATYPE_ULT_ALARM+` 
	 WHERE DeviceId=%f ;`, res.DeviceId)
}

//InsertSQL
func (res *UltDeviceAlarmType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.`+DATATYPE_ULT_ALARM+`  (DeviceId,AlarmStatus,AlarmTime,AlarmType,Alarm) 
	  VALUES (%f,'%s','%s','%s','%s') 
	  RETURNING DeviceId;`, res.DeviceId, res.AlarmStatus, res.AlarmTime, res.AlarmType, res.Alarm)
}

//UpdateSQL
func (res *UltDeviceAlarmType) UpdateSQL() string {
	return fmt.Sprintf(`UPDATE public.`+DATATYPE_ULT_ALARM+`  
	  SET AlarmStatus='%s',AlarmTime='%s',AlarmType='%s',Alarm='%s' 
	  WHERE DeviceId=%f  
	  RETURNING DeviceId;`,
		res.AlarmStatus,
		res.AlarmTime,
		res.AlarmType,
		res.Alarm,
		res.DeviceId)
}

//SelectWithDb
func (res *UltDeviceAlarmType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(
		&res.AlarmStatus,
		&res.AlarmTime,
		&res.AlarmType,
		&res.Alarm)
	return errDb
}

//CreateDb
func (res *UltDeviceAlarmType) CreateDb(currentDb *sql.DB) {
	createSQL := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS ` + DATATYPE_ULT_ALARM + `  (
	DataId serial PRIMARY KEY,
	DeviceId INT NOT NULL DEFAULT -1,
	AlarmStatus varchar(50) NOT NULL DEFAULT '` + ALARMSTATU_NONE + `',
	AlarmTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	AlarmType varchar(50) NOT NULL DEFAULT '` + ALARMTYPE_NONE + `',
	Alarm varchar(50) NOT NULL DEFAULT '',
	CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);`)
	_, err := currentDb.Exec(createSQL)
	LogErr(err)
}
