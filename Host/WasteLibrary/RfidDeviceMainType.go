package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/url"
)

//RfidDeviceMainType
type RfidDeviceMainType struct {
	DeviceId     float64
	CustomerId   float64
	SerialNumber string
	DeviceType   string
	Active       string
	CreateTime   string
}

//New
func (res *RfidDeviceMainType) New() {
	res.DeviceId = 0
	res.CustomerId = 1
	res.SerialNumber = ""
	res.DeviceType = RFID_DEVICE_TYPE_NONE
	res.Active = STATU_ACTIVE
	res.CreateTime = GetTime()
}

//GetByRedis
func (res *RfidDeviceMainType) GetByRedis(dbIndex string) ResultType {

	resultVal := GetRedisForStoreApi(dbIndex, REDIS_RFID_MAIN, res.ToIdString())
	if resultVal.Result == RESULT_OK {
		res.StringToType(resultVal.Retval.(string))
	} else {
		return resultVal
	}

	resultVal.Retval = res.ToString()
	return resultVal
}

//SaveToRedis
func (res *RfidDeviceMainType) SaveToRedis() ResultType {
	resultVal := SaveRedisForStoreApi(REDIS_RFID_MAIN, res.ToIdString(), res.ToString())
	return resultVal
}

//SaveToRedisBySerial
func (res *RfidDeviceMainType) SaveToRedisBySerial() ResultType {
	resultVal := SaveRedisForStoreApi(REDIS_SERIAL_RFID_DEVICE, res.SerialNumber, res.ToIdString())
	return resultVal
}

//SaveToDb
func (res *RfidDeviceMainType) SaveToDb() ResultType {
	var resultVal ResultType
	var currentHttpHeader HttpClientHeaderType
	currentHttpHeader.New()
	currentHttpHeader.DataType = DATATYPE_RFID_MAIN

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
func (res *RfidDeviceMainType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToCustomerId String
func (res *RfidDeviceMainType) ToCustomerIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToByte
func (res *RfidDeviceMainType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res *RfidDeviceMainType) ToString() string {
	return string(res.ToByte())

}

//ByteToType
func (res *RfidDeviceMainType) ByteToType(retByte []byte) {
	res.New()
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *RfidDeviceMainType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}

//SelectSQL
func (res *RfidDeviceMainType) SelectSQL() string {
	return fmt.Sprintf(`SELECT CustomerId,SerialNumber,Active,CreateTime,DeviceType
	 FROM public.`+DATATYPE_RFID_MAIN+` 
	 WHERE DeviceId=%f  ;`, res.DeviceId)
}

//InsertSQL
func (res *RfidDeviceMainType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.`+DATATYPE_RFID_MAIN+`  (CustomerId,SerialNumber,DeviceType) 
	  VALUES (%f,'%s','%s') 
	  RETURNING DeviceId;`, res.CustomerId, res.SerialNumber, res.DeviceType)
}

//InsertDataSQL
func (res *RfidDeviceMainType) InsertDataSQL() string {
	return fmt.Sprintf(`INSERT INTO public.`+DATATYPE_RFID_MAIN+`  (DeviceId,CustomerId,SerialNumber,DeviceType) 
	  VALUES (%f,%f,'%s','%s') 
	  RETURNING DeviceId;`, res.DeviceId, res.CustomerId, res.SerialNumber, res.DeviceType)
}

//UpdateSQL
func (res *RfidDeviceMainType) UpdateSQL() string {
	return fmt.Sprintf(`UPDATE public.`+DATATYPE_RFID_MAIN+`  
	  SET CustomerId=%f,DeviceType='%s'
	  WHERE DeviceId=%f  
	  RETURNING DeviceId;`,
		res.CustomerId,
		res.DeviceType,
		res.DeviceId)
}

//SelectWithDb
func (res *RfidDeviceMainType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(
		&res.CustomerId,
		&res.SerialNumber,
		&res.Active,
		&res.CreateTime,
		&res.DeviceType)
	return errDb
}

//CreateDb
func (res *RfidDeviceMainType) CreateDb(currentDb *sql.DB) {
	createSQL := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS ` + DATATYPE_RFID_MAIN + `  (
	DeviceId  serial PRIMARY KEY,
	CustomerId INT NOT NULL DEFAULT -1,
	SerialNumber  varchar(50) NOT NULL DEFAULT '',
	DeviceType  varchar(50) NOT NULL DEFAULT '` + RFID_DEVICE_TYPE_NONE + `',
	Active varchar(50) NOT NULL DEFAULT '` + STATU_ACTIVE + `',
	CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);`)
	_, err := currentDb.Exec(createSQL)
	LogErr(err)
}

//CreateDb
func (res *RfidDeviceMainType) CreateReaderDb(currentDb *sql.DB) {
	createSQL := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS ` + DATATYPE_RFID_MAIN + `  (
	DataId serial PRIMARY KEY,
	DeviceId INT NOT NULL DEFAULT -1,
	CustomerId INT NOT NULL DEFAULT -1,
	SerialNumber  varchar(50) NOT NULL DEFAULT '',
	DeviceType  varchar(50) NOT NULL DEFAULT '` + RFID_DEVICE_TYPE_NONE + `',
	Active varchar(50) NOT NULL DEFAULT '` + STATU_ACTIVE + `',
	CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);`)
	_, err := currentDb.Exec(createSQL)
	LogErr(err)
}
