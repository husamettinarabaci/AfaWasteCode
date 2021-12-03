package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/url"
)

//RecyDeviceThermType
type RecyDeviceThermType struct {
	DeviceId    float64
	Therm       string
	ThermTime   string
	ThermStatus string
	NewData     bool
}

//New
func (res *RecyDeviceThermType) New() {
	res.DeviceId = 0
	res.Therm = "00"
	res.ThermTime = GetTime()
	res.ThermStatus = THERMSTATU_NONE
	res.NewData = false
}

//GetByRedis
func (res *RecyDeviceThermType) GetByRedis(dbIndex string) ResultType {
	var resultVal ResultType
	resultVal = GetRedisForStoreApi(dbIndex, REDIS_RECY_THERM, res.ToIdString())
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
func (res *RecyDeviceThermType) SaveToRedis() ResultType {
	var resultVal ResultType
	resultVal = SaveRedisForStoreApi(REDIS_RECY_THERM, res.ToIdString(), res.ToString())
	return resultVal
}

//SaveToDb
func (res *RecyDeviceThermType) SaveToDb() ResultType {
	var resultVal ResultType
	var currentHttpHeader HttpClientHeaderType
	currentHttpHeader.New()
	currentHttpHeader.DataType = DATATYPE_RECY_THERM

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
func (res *RecyDeviceThermType) SaveToReaderDb() ResultType {
	var resultVal ResultType
	var currentHttpHeader HttpClientHeaderType
	currentHttpHeader.New()
	currentHttpHeader.DataType = DATATYPE_RECY_THERM

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
func (res *RecyDeviceThermType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToByte
func (res *RecyDeviceThermType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res *RecyDeviceThermType) ToString() string {
	return string(res.ToByte())

}

//Byte To RecyDeviceThermType
func ByteToRecyDeviceThermType(retByte []byte) RecyDeviceThermType {
	var retVal RecyDeviceThermType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To RecyDeviceThermType
func StringToRecyDeviceThermType(retStr string) RecyDeviceThermType {
	return ByteToRecyDeviceThermType([]byte(retStr))
}

//ByteToType
func (res *RecyDeviceThermType) ByteToType(retByte []byte) {
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *RecyDeviceThermType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}

//SelectSQL
func (res *RecyDeviceThermType) SelectSQL() string {
	return fmt.Sprintf(`SELECT Therm,ThermTime,ThermStatus
	 FROM public.`+DATATYPE_RECY_THERM+` 
	 WHERE DeviceId=%f ;`, res.DeviceId)
}

//InsertSQL
func (res *RecyDeviceThermType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.`+DATATYPE_RECY_THERM+`  (DeviceId,Therm,ThermTime,ThermStatus) 
	  VALUES (%f,'%s','%s','%s') 
	  RETURNING DeviceId;`, res.DeviceId, res.Therm, res.ThermTime, res.ThermStatus)
}

//UpdateSQL
func (res *RecyDeviceThermType) UpdateSQL() string {
	return fmt.Sprintf(`UPDATE public.`+DATATYPE_RECY_THERM+`  
	  SET Therm='%s',ThermTime='%s',ThermStatus='%s' 
	  WHERE DeviceId=%f  
	  RETURNING DeviceId;`,
		res.Therm,
		res.ThermTime,
		res.ThermStatus,
		res.DeviceId)
}

//SelectWithDb
func (res *RecyDeviceThermType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(
		&res.Therm,
		&res.ThermTime,
		&res.ThermStatus)
	return errDb
}
