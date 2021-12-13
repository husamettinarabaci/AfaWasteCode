package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/url"
)

//UltDeviceMainType
type UltDeviceMainType struct {
	DeviceId     float64
	CustomerId   float64
	SerialNumber string
	Latitude     float64
	Longitude    float64
	DeviceType   string
	Active       string
	CreateTime   string
}

//New
func (res *UltDeviceMainType) New() {
	res.DeviceId = 0
	res.CustomerId = 1
	res.SerialNumber = ""
	res.Latitude = 0
	res.Longitude = 0
	res.DeviceType = ULT_DEVICE_TYPE_NONE
	res.Active = STATU_ACTIVE
	res.CreateTime = GetTime()

}

//GetByRedis
func (res *UltDeviceMainType) GetByRedis(dbIndex string) ResultType {
	resultVal := GetRedisForStoreApi(dbIndex, REDIS_ULT_MAIN, res.ToIdString())
	if resultVal.Result == RESULT_OK {
		res.StringToType(resultVal.Retval.(string))
	} else {
		return resultVal
	}

	resultVal.Retval = res.ToString()
	return resultVal
}

//SaveToRedis
func (res *UltDeviceMainType) SaveToRedis() ResultType {
	resultVal := SaveRedisForStoreApi(REDIS_ULT_MAIN, res.ToIdString(), res.ToString())
	return resultVal
}

//SaveToRedisBySerial
func (res *UltDeviceMainType) SaveToRedisBySerial() ResultType {
	resultVal := SaveRedisForStoreApi(REDIS_SERIAL_ULT_DEVICE, res.SerialNumber, res.ToIdString())
	return resultVal
}

//SaveToDb
func (res *UltDeviceMainType) SaveToDb() ResultType {
	var resultVal ResultType
	var currentHttpHeader HttpClientHeaderType
	currentHttpHeader.New()
	currentHttpHeader.DataType = DATATYPE_ULT_MAIN

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
func (res *UltDeviceMainType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToCustomerId String
func (res *UltDeviceMainType) ToCustomerIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToByte
func (res *UltDeviceMainType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res *UltDeviceMainType) ToString() string {
	return string(res.ToByte())

}

//ByteToType
func (res *UltDeviceMainType) ByteToType(retByte []byte) {
	res.New()
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *UltDeviceMainType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}

//SelectSQL
func (res *UltDeviceMainType) SelectSQL() string {
	return fmt.Sprintf(`SELECT CustomerId,SerialNumber,Active,CreateTime,Latitude,Longitude,DeviceType
	 FROM public.`+DATATYPE_ULT_MAIN+` 
	 WHERE DeviceId=%f  ;`, res.DeviceId)
}

//InsertSQL
func (res *UltDeviceMainType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.`+DATATYPE_ULT_MAIN+`  (CustomerId,SerialNumber,Latitude,Longitude,DeviceType) 
	  VALUES (%f,'%s',%f,%f,'%s') 
	  RETURNING DeviceId;`, res.CustomerId, res.SerialNumber, res.Latitude, res.Longitude, res.DeviceType)
}

//InsertDataSQL
func (res *UltDeviceMainType) InsertDataSQL() string {
	return fmt.Sprintf(`INSERT INTO public.`+DATATYPE_ULT_MAIN+`  (DeviceId,CustomerId,SerialNumber,Latitude,Longitude,DeviceType) 
	  VALUES (%f,%f,'%s',%f,%f,'%s') 
	  RETURNING DeviceId;`, res.DeviceId, res.CustomerId, res.SerialNumber, res.Latitude, res.Longitude, res.DeviceType)
}

//UpdateSQL
func (res *UltDeviceMainType) UpdateSQL() string {
	return fmt.Sprintf(`UPDATE public.`+DATATYPE_ULT_MAIN+`  
	  SET CustomerId=%f,Latitude=%f,Longitude=%f,DeviceType='%s'
	  WHERE DeviceId=%f  
	  RETURNING DeviceId;`,
		res.CustomerId,
		res.Latitude,
		res.Longitude,
		res.DeviceType,
		res.DeviceId)
}

//SelectWithDb
func (res *UltDeviceMainType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(
		&res.CustomerId,
		&res.SerialNumber,
		&res.Active,
		&res.CreateTime,
		&res.Latitude,
		&res.Longitude,
		&res.DeviceType)
	return errDb
}

//CreateDb
func (res *UltDeviceMainType) CreateDb(currentDb *sql.DB) {
	createSQL := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS ` + DATATYPE_ULT_MAIN + `  (
	DeviceId  serial PRIMARY KEY,
	CustomerId INT NOT NULL DEFAULT -1,
	SerialNumber  varchar(50) NOT NULL DEFAULT '',
	Latitude NUMERIC(14, 11)  NOT NULL DEFAULT 0,
	Longitude NUMERIC(14, 11)  NOT NULL DEFAULT 0,
	DeviceType  varchar(50) NOT NULL DEFAULT '` + ULT_DEVICE_TYPE_NONE + `',
	Active varchar(50) NOT NULL DEFAULT '` + STATU_ACTIVE + `',
	CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);`)
	_, err := currentDb.Exec(createSQL)
	LogErr(err)
}

//CreateDb
func (res *UltDeviceMainType) CreateReaderDb(currentDb *sql.DB) {
	createSQL := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS ` + DATATYPE_ULT_MAIN + `  (
	DataId serial PRIMARY KEY,
	DeviceId INT NOT NULL DEFAULT -1,
	CustomerId INT NOT NULL DEFAULT -1,
	SerialNumber  varchar(50) NOT NULL DEFAULT '',
	Latitude NUMERIC(14, 11)  NOT NULL DEFAULT 0,
	Longitude NUMERIC(14, 11)  NOT NULL DEFAULT 0,
	DeviceType  varchar(50) NOT NULL DEFAULT '` + ULT_DEVICE_TYPE_NONE + `',
	Active varchar(50) NOT NULL DEFAULT '` + STATU_ACTIVE + `',
	CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);`)
	_, err := currentDb.Exec(createSQL)
	LogErr(err)
}
