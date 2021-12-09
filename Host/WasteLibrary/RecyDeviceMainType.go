package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/url"
)

//RecyDeviceMainType
type RecyDeviceMainType struct {
	DeviceId     float64
	CustomerId   float64
	SerialNumber string
	Active       string
	CreateTime   string
}

//New
func (res *RecyDeviceMainType) New() {
	res.DeviceId = 0
	res.CustomerId = 1
	res.SerialNumber = ""
	res.Active = STATU_ACTIVE
	res.CreateTime = GetTime()
}

//GetByRedis
func (res *RecyDeviceMainType) GetByRedis(dbIndex string) ResultType {
	resultVal := GetRedisForStoreApi(dbIndex, REDIS_RECY_MAIN, res.ToIdString())
	if resultVal.Result == RESULT_OK {
		res.StringToType(resultVal.Retval.(string))
	} else {
		return resultVal
	}

	resultVal.Retval = res.ToString()
	return resultVal
}

//SaveToRedis
func (res *RecyDeviceMainType) SaveToRedis() ResultType {
	resultVal := SaveRedisForStoreApi(REDIS_RECY_MAIN, res.ToIdString(), res.ToString())
	return resultVal
}

//SaveToRedisBySerial
func (res *RecyDeviceMainType) SaveToRedisBySerial() ResultType {
	resultVal := SaveRedisForStoreApi(REDIS_SERIAL_RECY_DEVICE, res.SerialNumber, res.ToIdString())
	return resultVal
}

//SaveToDb
func (res *RecyDeviceMainType) SaveToDb() ResultType {
	var resultVal ResultType
	var currentHttpHeader HttpClientHeaderType
	currentHttpHeader.New()
	currentHttpHeader.DataType = DATATYPE_RECY_MAIN

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
func (res *RecyDeviceMainType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToCustomerId String
func (res *RecyDeviceMainType) ToCustomerIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToByte
func (res *RecyDeviceMainType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res *RecyDeviceMainType) ToString() string {
	return string(res.ToByte())

}

//ByteToType
func (res *RecyDeviceMainType) ByteToType(retByte []byte) {
	res.New()
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *RecyDeviceMainType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}

//SelectSQL
func (res *RecyDeviceMainType) SelectSQL() string {
	return fmt.Sprintf(`SELECT CustomerId,SerialNumber,Active,CreateTime
	 FROM public.`+DATATYPE_RECY_MAIN+` 
	 WHERE DeviceId=%f  ;`, res.DeviceId)
}

//InsertSQL
func (res *RecyDeviceMainType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.`+DATATYPE_RECY_MAIN+`  (CustomerId,SerialNumber) 
	  VALUES (%f,'%s') 
	  RETURNING DeviceId;`, res.CustomerId, res.SerialNumber)
}

//InsertDataSQL
func (res *RecyDeviceMainType) InsertDataSQL() string {
	return fmt.Sprintf(`INSERT INTO public.`+DATATYPE_RECY_MAIN+`  (DeviceId,CustomerId,SerialNumber) 
	  VALUES (%f,%f,'%s') 
	  RETURNING DeviceId;`, res.DeviceId, res.CustomerId, res.SerialNumber)
}

//UpdateSQL
func (res *RecyDeviceMainType) UpdateSQL() string {
	return fmt.Sprintf(`UPDATE public.`+DATATYPE_RECY_MAIN+`  
	  SET CustomerId=%f 
	  WHERE DeviceId=%f  
	  RETURNING DeviceId;`,
		res.CustomerId,
		res.DeviceId)
}

//SelectWithDb
func (res *RecyDeviceMainType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(
		&res.CustomerId,
		&res.SerialNumber,
		&res.Active,
		&res.CreateTime)
	return errDb
}

//CreateDb
func (res *RecyDeviceMainType) CreateDb(currentDb *sql.DB) {
	createSQL := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS ` + DATATYPE_RECY_MAIN + `  (
	DeviceId  serial PRIMARY KEY,
	CustomerId INT NOT NULL DEFAULT -1,
	SerialNumber  varchar(50) NOT NULL DEFAULT '',
	Active varchar(50) NOT NULL DEFAULT '` + STATU_ACTIVE + `',
	CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);`)
	_, err := currentDb.Exec(createSQL)
	LogErr(err)
}

//CreateDb
func (res *RecyDeviceMainType) CreateReaderDb(currentDb *sql.DB) {
	createSQL := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS ` + DATATYPE_RECY_MAIN + `  (
	DataId serial PRIMARY KEY,
	DeviceId INT NOT NULL DEFAULT -1,
	CustomerId INT NOT NULL DEFAULT -1,
	SerialNumber  varchar(50) NOT NULL DEFAULT '',
	Active varchar(50) NOT NULL DEFAULT '` + STATU_ACTIVE + `',
	CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);`)
	_, err := currentDb.Exec(createSQL)
	LogErr(err)
}
