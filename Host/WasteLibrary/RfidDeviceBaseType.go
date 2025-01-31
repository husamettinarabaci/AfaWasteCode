package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/url"
)

//RfidDeviceBaseType
type RfidDeviceBaseType struct {
	DeviceId      float64
	PlateNo       string
	DriverName    string
	DriverSurName string
	NewData       bool
}

//New
func (res *RfidDeviceBaseType) New() {
	res.DeviceId = 0
	res.PlateNo = ""
	res.DriverName = ""
	res.DriverSurName = ""
	res.NewData = false
}

//GetByRedis
func (res *RfidDeviceBaseType) GetByRedis(dbIndex string) ResultType {
	resultVal := GetRedisForStoreApi(dbIndex, REDIS_RFID_BASE, res.ToIdString())
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
func (res *RfidDeviceBaseType) SaveToRedis() ResultType {
	resultVal := SaveRedisForStoreApi(REDIS_RFID_BASE, res.ToIdString(), res.ToString())
	return resultVal
}

//SaveToDb
func (res *RfidDeviceBaseType) SaveToDb() ResultType {
	var resultVal ResultType
	var currentHttpHeader HttpClientHeaderType
	currentHttpHeader.New()
	currentHttpHeader.DataType = DATATYPE_RFID_BASE

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
func (res *RfidDeviceBaseType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToByte
func (res *RfidDeviceBaseType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res *RfidDeviceBaseType) ToString() string {
	return string(res.ToByte())

}

//ByteToType
func (res *RfidDeviceBaseType) ByteToType(retByte []byte) {
	res.New()
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *RfidDeviceBaseType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}

//SelectSQL
func (res *RfidDeviceBaseType) SelectSQL() string {
	return fmt.Sprintf(`SELECT PlateNo,DriverName,DriverSurName
	 FROM public.`+DATATYPE_RFID_BASE+` 
	 WHERE DeviceId=%f ;`, res.DeviceId)
}

//InsertSQL
func (res *RfidDeviceBaseType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.`+DATATYPE_RFID_BASE+`  (DeviceId,PlateNo,DriverName,DriverSurName) 
	  VALUES (%f,'%s','%s','%s') 
	  RETURNING DeviceId;`, res.DeviceId, res.PlateNo, res.DriverName, res.DriverSurName)
}

//UpdateSQL
func (res *RfidDeviceBaseType) UpdateSQL() string {
	return fmt.Sprintf(`UPDATE public.`+DATATYPE_RFID_BASE+`  
	  SET PlateNo='%s',DriverName='%s',DriverSurName='%s'
	  WHERE DeviceId=%f  
	  RETURNING DeviceId;`,
		res.PlateNo,
		res.DriverName,
		res.DriverSurName,
		res.DeviceId)
}

//SelectWithDb
func (res *RfidDeviceBaseType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(
		&res.PlateNo,
		&res.DriverName,
		&res.DriverSurName)
	return errDb
}

//CreateDb
func (res *RfidDeviceBaseType) CreateDb(currentDb *sql.DB) {
	createSQL := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS ` + DATATYPE_RFID_BASE + `  ( 
	DataId serial PRIMARY KEY,
	DeviceId INT NOT NULL DEFAULT -1,
	PlateNo varchar(50) NOT NULL DEFAULT '',
	DriverName varchar(50) NOT NULL DEFAULT '',
	DriverSurName varchar(50) NOT NULL DEFAULT '',
	CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);`)
	_, err := currentDb.Exec(createSQL)
	LogErr(err)
}
