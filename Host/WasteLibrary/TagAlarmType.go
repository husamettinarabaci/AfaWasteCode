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
func (res *TagAlarmType) GetByRedis() ResultType {

	var resultVal ResultType
	resultVal = GetRedisForStoreApi(REDIS_TAG_ALARMS, res.ToIdString())
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
	var resultVal ResultType
	resultVal = SaveRedisForStoreApi(REDIS_TAG_ALARMS, res.ToIdString(), res.ToString())
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

//Byte To TagAlarmType
func ByteToTagAlarmType(retByte []byte) TagAlarmType {
	var retVal TagAlarmType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To TagAlarmType
func StringToTagAlarmType(retStr string) TagAlarmType {
	return ByteToTagAlarmType([]byte(retStr))
}

//ByteToType
func (res *TagAlarmType) ByteToType(retByte []byte) {
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *TagAlarmType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}

//SelectSQL
func (res *TagAlarmType) SelectSQL() string {
	return fmt.Sprintf(`SELECT AlarmStatus,AlarmTime,AlarmType,Alarm
	 FROM public.tag_alarms
	 WHERE TagId=%f ;`, res.TagId)
}

//InsertSQL
func (res *TagAlarmType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.tag_alarms (TagId,AlarmStatus,AlarmTime,AlarmType,Alarm) 
	  VALUES (%f,'%s','%s','%s','%s') 
	  RETURNING TagId;`, res.TagId, res.AlarmStatus, res.AlarmTime, res.AlarmType, res.Alarm)
}

//UpdateSQL
func (res *TagAlarmType) UpdateSQL() string {
	return fmt.Sprintf(`UPDATE public.tag_alarms 
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
