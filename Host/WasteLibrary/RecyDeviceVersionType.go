package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/url"
)

//RecyDeviceVersionType
type RecyDeviceVersionType struct {
	DeviceId           float64
	WebAppVersion      string
	MotorAppVersion    string
	ThermAppVersion    string
	TransferAppVersion string
	CheckerAppVersion  string
	CamAppVersion      string
	ReaderAppVersion   string
	SystemAppVersion   string
	NewData            bool
}

//New
func (res *RecyDeviceVersionType) New() {
	res.DeviceId = 0
	res.WebAppVersion = "1"
	res.MotorAppVersion = "1"
	res.ThermAppVersion = "1"
	res.TransferAppVersion = "1"
	res.CheckerAppVersion = "1"
	res.CamAppVersion = "1"
	res.ReaderAppVersion = "1"
	res.SystemAppVersion = "1"
	res.NewData = false
}

//GetByRedis
func (res *RecyDeviceVersionType) GetByRedis(dbIndex string) ResultType {
	var resultVal ResultType
	resultVal = GetRedisForStoreApi(dbIndex, REDIS_RECY_VERSION, res.ToIdString())
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
func (res *RecyDeviceVersionType) SaveToRedis() ResultType {
	var resultVal ResultType
	resultVal = SaveRedisForStoreApi(REDIS_RECY_VERSION, res.ToIdString(), res.ToString())
	return resultVal
}

//SaveToDb
func (res *RecyDeviceVersionType) SaveToDb() ResultType {
	var resultVal ResultType
	var currentHttpHeader HttpClientHeaderType
	currentHttpHeader.New()
	currentHttpHeader.DataType = DATATYPE_RECY_VERSION

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
func (res *RecyDeviceVersionType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToByte
func (res *RecyDeviceVersionType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res *RecyDeviceVersionType) ToString() string {
	return string(res.ToByte())

}

//ByteToType
func (res *RecyDeviceVersionType) ByteToType(retByte []byte) {
	res.New()
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *RecyDeviceVersionType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}

//SelectSQL
func (res *RecyDeviceVersionType) SelectSQL() string {
	return fmt.Sprintf(`SELECT WebAppVersion,MotorAppVersion,ThermAppVersion,TransferAppVersion,CheckerAppVersion,CamAppVersion,ReaderAppVersion,SystemAppVersion
	 FROM public.`+DATATYPE_RECY_VERSION+` 
	 WHERE DeviceId=%f ;`, res.DeviceId)
}

//InsertSQL
func (res *RecyDeviceVersionType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.`+DATATYPE_RECY_VERSION+`  (DeviceId,WebAppVersion,MotorAppVersion,ThermAppVersion,TransferAppVersion,CheckerAppVersion,CamAppVersion,ReaderAppVersion,SystemAppVersion) 
	  VALUES (%f,'%s','%s','%s','%s','%s','%s','%s','%s') 
	  RETURNING DeviceId;`, res.DeviceId, res.WebAppVersion, res.MotorAppVersion, res.ThermAppVersion, res.TransferAppVersion, res.CheckerAppVersion, res.CamAppVersion, res.ReaderAppVersion, res.SystemAppVersion)
}

//UpdateSQL
func (res *RecyDeviceVersionType) UpdateSQL() string {
	return fmt.Sprintf(`UPDATE public.`+DATATYPE_RECY_VERSION+`  
	  SET WebAppVersion='%s',MotorAppVersion='%s',ThermAppVersion='%s',TransferAppVersion='%s',CheckerAppVersion='%s',CamAppVersion='%s',ReaderAppVersion='%s',SystemAppVersion='%s' 
	  WHERE DeviceId=%f  
	  RETURNING DeviceId;`,
		res.WebAppVersion,
		res.MotorAppVersion,
		res.ThermAppVersion,
		res.TransferAppVersion,
		res.CheckerAppVersion,
		res.CamAppVersion,
		res.ReaderAppVersion,
		res.SystemAppVersion,
		res.DeviceId)
}

//SelectWithDb
func (res *RecyDeviceVersionType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(
		&res.WebAppVersion,
		&res.MotorAppVersion,
		&res.ThermAppVersion,
		&res.TransferAppVersion,
		&res.CheckerAppVersion,
		&res.CamAppVersion,
		&res.ReaderAppVersion,
		&res.SystemAppVersion)
	return errDb
}

//CreateDb
func (res *RecyDeviceVersionType) CreateDb(currentDb *sql.DB) {
	createSQL := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS ` + DATATYPE_RECY_VERSION + `  (
	DataId serial PRIMARY KEY,
	DeviceId INT NOT NULL DEFAULT -1,
	WebAppVersion varchar(50) NOT NULL DEFAULT '1',
	MotorAppVersion varchar(50) NOT NULL DEFAULT '1',
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
