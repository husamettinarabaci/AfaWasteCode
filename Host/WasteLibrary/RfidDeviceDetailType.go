package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/url"
)

//RfidDeviceDetailType
type RfidDeviceDetailType struct {
	DeviceId      float64
	PlateNo       string
	DriverName    string
	DriverSurName string
	NewData       bool
}

//New
func (res *RfidDeviceDetailType) New() {
	res.DeviceId = 0
	res.PlateNo = ""
	res.DriverName = ""
	res.DriverSurName = ""
	res.NewData = false
}

//GetByRedis
func (res *RfidDeviceDetailType) GetByRedis(dbIndex string) ResultType {

	var resultVal ResultType
	resultVal = GetRedisForStoreApi(dbIndex, REDIS_RFID_DETAIL, res.ToIdString())
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
func (res *RfidDeviceDetailType) SaveToRedis() ResultType {
	var resultVal ResultType
	resultVal = SaveRedisForStoreApi(REDIS_RFID_DETAIL, res.ToIdString(), res.ToString())
	return resultVal
}

//SaveToDb
func (res *RfidDeviceDetailType) SaveToDb() ResultType {
	var resultVal ResultType
	var currentHttpHeader HttpClientHeaderType
	currentHttpHeader.New()
	currentHttpHeader.DataType = DATATYPE_RFID_DETAIL

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
func (res *RfidDeviceDetailType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToByte
func (res *RfidDeviceDetailType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res *RfidDeviceDetailType) ToString() string {
	return string(res.ToByte())

}

//Byte To RfidDeviceDetailType
func ByteToRfidDeviceDetailType(retByte []byte) RfidDeviceDetailType {
	var retVal RfidDeviceDetailType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To RfidDeviceDetailType
func StringToRfidDeviceDetailType(retStr string) RfidDeviceDetailType {
	return ByteToRfidDeviceDetailType([]byte(retStr))
}

//ByteToType
func (res *RfidDeviceDetailType) ByteToType(retByte []byte) {
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *RfidDeviceDetailType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}

//SelectSQL
func (res *RfidDeviceDetailType) SelectSQL() string {
	return fmt.Sprintf(`SELECT PlateNo,DriverName,DriverSurName
	 FROM public.`+DATATYPE_RFID_DETAIL+` 
	 WHERE DeviceId=%f ;`, res.DeviceId)
}

//InsertSQL
func (res *RfidDeviceDetailType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.`+DATATYPE_RFID_DETAIL+`  (DeviceId,PlateNo,DriverName,DriverSurName) 
	  VALUES (%f,'%s','%s','%s') 
	  RETURNING DeviceId;`, res.DeviceId, res.PlateNo, res.DriverName, res.DriverSurName)
}

//UpdateSQL
func (res *RfidDeviceDetailType) UpdateSQL() string {
	return fmt.Sprintf(`UPDATE public.`+DATATYPE_RFID_DETAIL+`  
	  SET PlateNo='%s',DriverName='%s',DriverSurName='%s'
	  WHERE DeviceId=%f  
	  RETURNING DeviceId;`,
		res.PlateNo,
		res.DriverName,
		res.DriverSurName,
		res.DeviceId)
}

//SelectWithDb
func (res *RfidDeviceDetailType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(
		&res.PlateNo,
		&res.DriverName,
		&res.DriverSurName)
	return errDb
}

//CreateDb
func (res *RfidDeviceDetailType) CreateDb(currentDb *sql.DB) {
	createSQL := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS ` + DATATYPE_RFID_DETAIL + `  ( 
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
