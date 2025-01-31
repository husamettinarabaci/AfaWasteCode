package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/url"
)

//UltDeviceThermType
type UltDeviceThermType struct {
	DeviceId    float64
	Therm       string
	ThermTime   string
	ThermStatus string
	NewData     bool
}

//New
func (res *UltDeviceThermType) New() {
	res.DeviceId = 0
	res.Therm = "00"
	res.ThermTime = GetTime()
	res.ThermStatus = THERMSTATU_NONE
	res.NewData = false
}

//GetByRedis
func (res *UltDeviceThermType) GetByRedis(dbIndex string) ResultType {
	resultVal := GetRedisForStoreApi(dbIndex, REDIS_ULT_THERM, res.ToIdString())
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
func (res *UltDeviceThermType) SaveToRedis() ResultType {
	resultVal := SaveRedisForStoreApi(REDIS_ULT_THERM, res.ToIdString(), res.ToString())
	return resultVal
}

//SaveToDb
func (res *UltDeviceThermType) SaveToDb() ResultType {
	var resultVal ResultType
	var currentHttpHeader HttpClientHeaderType
	currentHttpHeader.New()
	currentHttpHeader.DataType = DATATYPE_ULT_THERM

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
func (res *UltDeviceThermType) SaveToReaderDb() ResultType {
	var resultVal ResultType
	var currentHttpHeader HttpClientHeaderType
	currentHttpHeader.New()
	currentHttpHeader.DataType = DATATYPE_ULT_THERM

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
func (res *UltDeviceThermType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToByte
func (res *UltDeviceThermType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res *UltDeviceThermType) ToString() string {
	return string(res.ToByte())

}

//ByteToType
func (res *UltDeviceThermType) ByteToType(retByte []byte) {
	res.New()
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *UltDeviceThermType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}

//SelectSQL
func (res *UltDeviceThermType) SelectSQL() string {
	return fmt.Sprintf(`SELECT Therm,ThermTime,ThermStatus
	 FROM public.`+DATATYPE_ULT_THERM+` 
	 WHERE DeviceId=%f ;`, res.DeviceId)
}

//InsertSQL
func (res *UltDeviceThermType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.`+DATATYPE_ULT_THERM+`  (DeviceId,Therm,ThermTime,ThermStatus) 
	  VALUES (%f,'%s','%s','%s') 
	  RETURNING DeviceId;`, res.DeviceId, res.Therm, res.ThermTime, res.ThermStatus)
}

//UpdateSQL
func (res *UltDeviceThermType) UpdateSQL() string {
	return fmt.Sprintf(`UPDATE public.`+DATATYPE_ULT_THERM+`  
	  SET Therm='%s',ThermTime='%s',ThermStatus='%s' 
	  WHERE DeviceId=%f  
	  RETURNING DeviceId;`,
		res.Therm,
		res.ThermTime,
		res.ThermStatus,
		res.DeviceId)
}

//SelectWithDb
func (res *UltDeviceThermType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(
		&res.Therm,
		&res.ThermTime,
		&res.ThermStatus)
	return errDb
}

//CreateDb
func (res *UltDeviceThermType) CreateDb(currentDb *sql.DB) {
	createSQL := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS ` + DATATYPE_ULT_THERM + `  (
	DataId serial PRIMARY KEY,
	DeviceId INT NOT NULL DEFAULT -1,
	Therm varchar(50) NOT NULL DEFAULT '0',
	ThermTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	ThermStatus varchar(50) NOT NULL DEFAULT '` + THERMSTATU_NONE + `',
	CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);`)
	_, err := currentDb.Exec(createSQL)
	LogErr(err)
}
