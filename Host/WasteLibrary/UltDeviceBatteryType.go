package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/url"
)

//UltDeviceBatteryType
type UltDeviceBatteryType struct {
	DeviceId      float64
	Battery       string
	BatteryStatus string
	BatteryTime   string
	NewData       bool
}

//New
func (res *UltDeviceBatteryType) New() {
	res.DeviceId = 0
	res.Battery = "0000"
	res.BatteryStatus = BATTERYSTATU_NONE
	res.BatteryTime = GetTime()
	res.NewData = false
}

//GetByRedis
func (res *UltDeviceBatteryType) GetByRedis(dbIndex string) ResultType {
	resultVal := GetRedisForStoreApi(dbIndex, REDIS_ULT_BATTERY, res.ToIdString())
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
func (res *UltDeviceBatteryType) SaveToRedis() ResultType {
	resultVal := SaveRedisForStoreApi(REDIS_ULT_BATTERY, res.ToIdString(), res.ToString())
	return resultVal
}

//SaveToDb
func (res *UltDeviceBatteryType) SaveToDb() ResultType {
	var resultVal ResultType
	var currentHttpHeader HttpClientHeaderType
	currentHttpHeader.New()
	currentHttpHeader.DataType = DATATYPE_ULT_BATTERY

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
func (res *UltDeviceBatteryType) SaveToReaderDb() ResultType {
	var resultVal ResultType
	var currentHttpHeader HttpClientHeaderType
	currentHttpHeader.New()
	currentHttpHeader.DataType = DATATYPE_ULT_BATTERY

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
func (res *UltDeviceBatteryType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToByte
func (res *UltDeviceBatteryType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res *UltDeviceBatteryType) ToString() string {
	return string(res.ToByte())

}

//ByteToType
func (res *UltDeviceBatteryType) ByteToType(retByte []byte) {
	res.New()
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *UltDeviceBatteryType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}

//SelectSQL
func (res *UltDeviceBatteryType) SelectSQL() string {
	return fmt.Sprintf(`SELECT Battery,BatteryStatus,BatteryTime
	 FROM public.`+DATATYPE_ULT_BATTERY+` 
	 WHERE DeviceId=%f ;`, res.DeviceId)
}

//InsertSQL
func (res *UltDeviceBatteryType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.`+DATATYPE_ULT_BATTERY+`  (DeviceId,Battery,BatteryStatus,BatteryTime) 
	  VALUES (%f,'%s','%s','%s') 
	  RETURNING DeviceId;`, res.DeviceId, res.Battery, res.BatteryStatus, res.BatteryTime)
}

//UpdateSQL
func (res *UltDeviceBatteryType) UpdateSQL() string {
	return fmt.Sprintf(`UPDATE public.`+DATATYPE_ULT_BATTERY+`  
	  SET Battery='%s',BatteryStatus='%s',BatteryTime='%s'
	  WHERE DeviceId=%f  
	  RETURNING DeviceId;`,
		res.Battery,
		res.BatteryStatus,
		res.BatteryTime,
		res.DeviceId)
}

//SelectWithDb
func (res *UltDeviceBatteryType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(
		&res.Battery,
		&res.BatteryStatus,
		&res.BatteryTime)
	return errDb
}

//CreateDb
func (res *UltDeviceBatteryType) CreateDb(currentDb *sql.DB) {
	createSQL := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS ` + DATATYPE_ULT_BATTERY + `  (
	DataId serial PRIMARY KEY,
	DeviceId INT NOT NULL DEFAULT -1,
	Battery varchar(50) NOT NULL DEFAULT '0',
	BatteryStatus varchar(50) NOT NULL DEFAULT '` + BATTERYSTATU_NONE + `',
	BatteryTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);`)
	_, err := currentDb.Exec(createSQL)
	LogErr(err)
}
