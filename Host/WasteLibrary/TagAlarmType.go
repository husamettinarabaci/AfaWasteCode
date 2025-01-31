package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/url"
)

//TagAlarmType
type TagAlarmType struct {
	TagId       float64
	AlarmStatus string
	AlarmTime   string
	AlarmType   string
	Alarm       string
	NewData     bool
}

//New
func (res *TagAlarmType) New() {
	res.TagId = 0
	res.AlarmStatus = ALARMSTATU_NONE
	res.AlarmTime = GetTime()
	res.AlarmType = ALARMTYPE_NONE
	res.Alarm = ""
	res.NewData = false
}

//GetByRedis
func (res *TagAlarmType) GetByRedis(dbIndex string) ResultType {

	resultVal := GetRedisForStoreApi(dbIndex, REDIS_TAG_ALARM, res.ToIdString())
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
func (res *TagAlarmType) SaveToRedis() ResultType {
	resultVal := SaveRedisForStoreApi(REDIS_TAG_ALARM, res.ToIdString(), res.ToString())
	return resultVal
}

//SaveToDb
func (res *TagAlarmType) SaveToDb() ResultType {
	var resultVal ResultType
	var currentHttpHeader HttpClientHeaderType
	currentHttpHeader.New()
	currentHttpHeader.DataType = DATATYPE_TAG_ALARM

	data := url.Values{
		HTTP_HEADER: {currentHttpHeader.ToString()},
		HTTP_DATA:   {res.ToString()},
	}
	resultVal = SaveStaticDbMainForStoreApi(data)
	if resultVal.Result == RESULT_OK {
		res.TagId = StringIdToFloat64(resultVal.Retval.(string))
		resultVal.Retval = res.ToString()
	}

	return resultVal
}

//ToId String
func (res *TagAlarmType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.TagId)
}

//ToByte
func (res *TagAlarmType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res *TagAlarmType) ToString() string {
	return string(res.ToByte())

}

//ByteToType
func (res *TagAlarmType) ByteToType(retByte []byte) {
	res.New()
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *TagAlarmType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}

//SelectSQL
func (res *TagAlarmType) SelectSQL() string {
	return fmt.Sprintf(`SELECT AlarmStatus,AlarmTime,AlarmType,Alarm
	 FROM public.`+DATATYPE_TAG_ALARM+` 
	 WHERE TagId=%f ;`, res.TagId)
}

//InsertSQL
func (res *TagAlarmType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.`+DATATYPE_TAG_ALARM+`  (TagId,AlarmStatus,AlarmTime,AlarmType,Alarm) 
	  VALUES (%f,'%s','%s','%s','%s') 
	  RETURNING TagId;`, res.TagId, res.AlarmStatus, res.AlarmTime, res.AlarmType, res.Alarm)
}

//UpdateSQL
func (res *TagAlarmType) UpdateSQL() string {
	return fmt.Sprintf(`UPDATE public.`+DATATYPE_TAG_ALARM+`  
	  SET AlarmStatus='%s',AlarmTime='%s',AlarmType='%s',Alarm='%s' 
	  WHERE TagId=%f  
	  RETURNING TagId;`,
		res.AlarmStatus,
		res.AlarmTime,
		res.AlarmType,
		res.Alarm,
		res.TagId)
}

//SelectWithDb
func (res *TagAlarmType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(
		&res.AlarmStatus,
		&res.AlarmTime,
		&res.AlarmType,
		&res.Alarm)
	return errDb
}

//CreateDb
func (res *TagAlarmType) CreateDb(currentDb *sql.DB) {
	createSQL := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS ` + DATATYPE_TAG_ALARM + `  (
	DataId serial PRIMARY KEY,
	TagId INT NOT NULL DEFAULT -1,
	AlarmStatus varchar(50) NOT NULL DEFAULT '` + ALARMSTATU_NONE + `',
	AlarmTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	AlarmType varchar(50) NOT NULL DEFAULT '` + ALARMTYPE_NONE + `',
	Alarm varchar(50) NOT NULL DEFAULT '',
	CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);`)
	_, err := currentDb.Exec(createSQL)
	LogErr(err)
}
