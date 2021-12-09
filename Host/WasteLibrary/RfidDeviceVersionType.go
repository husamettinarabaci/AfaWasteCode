package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/url"
)

//RfidDeviceVersionType
type RfidDeviceVersionType struct {
	DeviceId           float64
	GpsAppVersion      string
	ThermAppVersion    string
	TransferAppVersion string
	CheckerAppVersion  string
	CamAppVersion      string
	ReaderAppVersion   string
	SystemAppVersion   string
	NewData            bool
}

//New
func (res *RfidDeviceVersionType) New() {
	res.DeviceId = 0
	res.GpsAppVersion = "1"
	res.ThermAppVersion = "1"
	res.TransferAppVersion = "1"
	res.CheckerAppVersion = "1"
	res.CamAppVersion = "1"
	res.ReaderAppVersion = "1"
	res.SystemAppVersion = "1"
	res.NewData = false
}

//GetByRedis
func (res *RfidDeviceVersionType) GetByRedis(dbIndex string) ResultType {

	resultVal := GetRedisForStoreApi(dbIndex, REDIS_RFID_VERSION, res.ToIdString())
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
func (res *RfidDeviceVersionType) SaveToRedis() ResultType {
	resultVal := SaveRedisForStoreApi(REDIS_RFID_VERSION, res.ToIdString(), res.ToString())
	return resultVal
}

//SaveToDb
func (res *RfidDeviceVersionType) SaveToDb() ResultType {
	var resultVal ResultType
	var currentHttpHeader HttpClientHeaderType
	currentHttpHeader.New()
	currentHttpHeader.DataType = DATATYPE_RFID_VERSION

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
func (res *RfidDeviceVersionType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToByte
func (res *RfidDeviceVersionType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res *RfidDeviceVersionType) ToString() string {
	return string(res.ToByte())

}

//ByteToType
func (res *RfidDeviceVersionType) ByteToType(retByte []byte) {
	res.New()
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *RfidDeviceVersionType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}

//SelectSQL
func (res *RfidDeviceVersionType) SelectSQL() string {
	return fmt.Sprintf(`SELECT GpsAppVersion,ThermAppVersion,TransferAppVersion,CheckerAppVersion,CamAppVersion,ReaderAppVersion,SystemAppVersion
	 FROM public.`+REDIS_RFID_VERSION+` 
	 WHERE DeviceId=%f ;`, res.DeviceId)
}

//InsertSQL
func (res *RfidDeviceVersionType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.`+REDIS_RFID_VERSION+`  (DeviceId,GpsAppVersion,ThermAppVersion,TransferAppVersion,CheckerAppVersion,CamAppVersion,ReaderAppVersion,SystemAppVersion) 
	  VALUES (%f,'%s','%s','%s','%s','%s','%s','%s') 
	  RETURNING DeviceId;`, res.DeviceId, res.GpsAppVersion, res.ThermAppVersion, res.TransferAppVersion, res.CheckerAppVersion, res.CamAppVersion, res.ReaderAppVersion, res.SystemAppVersion)
}

//UpdateSQL
func (res *RfidDeviceVersionType) UpdateSQL() string {
	return fmt.Sprintf(`UPDATE public.`+REDIS_RFID_VERSION+`  
	  SET GpsAppVersion='%s',ThermAppVersion='%s',TransferAppVersion='%s',CheckerAppVersion='%s',CamAppVersion='%s',ReaderAppVersion='%s',SystemAppVersion='%s' 
	  WHERE DeviceId=%f  
	  RETURNING DeviceId;`,
		res.GpsAppVersion,
		res.ThermAppVersion,
		res.TransferAppVersion,
		res.CheckerAppVersion,
		res.CamAppVersion,
		res.ReaderAppVersion,
		res.SystemAppVersion,
		res.DeviceId)
}

//SelectWithDb
func (res *RfidDeviceVersionType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(
		&res.GpsAppVersion,
		&res.ThermAppVersion,
		&res.TransferAppVersion,
		&res.CheckerAppVersion,
		&res.CamAppVersion,
		&res.ReaderAppVersion,
		&res.SystemAppVersion)
	return errDb
}

//CreateDb
func (res *RfidDeviceVersionType) CreateDb(currentDb *sql.DB) {
	createSQL := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS ` + DATATYPE_RFID_VERSION + `  (
	DataId serial PRIMARY KEY,
	DeviceId INT NOT NULL DEFAULT -1,
	GpsAppVersion varchar(50) NOT NULL DEFAULT '1',
	ThermAppVersion varchar(50) NOT NULL DEFAULT '1',
	TransferAppVersion varchar(50) NOT NULL DEFAULT '1',
	CheckerAppVersion varchar(50) NOT NULL DEFAULT '1',
	CamAppVersion varchar(50) NOT NULL DEFAULT '1',
	ReaderAppVersion varchar(50) NOT NULL DEFAULT '1',
	SystemAppVersion varchar(50) NOT NULL DEFAULT '1',
	CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);`)
	_, err := currentDb.Exec(createSQL)
	LogErr(err)
}
