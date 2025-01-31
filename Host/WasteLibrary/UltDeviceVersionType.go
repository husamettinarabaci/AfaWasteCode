package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/url"
)

//UltDeviceVersionType
type UltDeviceVersionType struct {
	DeviceId        float64
	FirmwareVersion string
	NewData         bool
}

//New
func (res *UltDeviceVersionType) New() {
	res.DeviceId = 0
	res.FirmwareVersion = "1"
	res.NewData = false

}

//GetByRedis
func (res *UltDeviceVersionType) GetByRedis(dbIndex string) ResultType {
	resultVal := GetRedisForStoreApi(dbIndex, REDIS_ULT_VERSION, res.ToIdString())
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
func (res *UltDeviceVersionType) SaveToRedis() ResultType {
	resultVal := SaveRedisForStoreApi(REDIS_ULT_VERSION, res.ToIdString(), res.ToString())
	return resultVal
}

//SaveToDb
func (res *UltDeviceVersionType) SaveToDb() ResultType {
	var resultVal ResultType
	var currentHttpHeader HttpClientHeaderType
	currentHttpHeader.New()
	currentHttpHeader.DataType = DATATYPE_ULT_VERSION

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
func (res *UltDeviceVersionType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToByte
func (res *UltDeviceVersionType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res *UltDeviceVersionType) ToString() string {
	return string(res.ToByte())

}

//ByteToType
func (res *UltDeviceVersionType) ByteToType(retByte []byte) {
	res.New()
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *UltDeviceVersionType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}

//SelectSQL
func (res *UltDeviceVersionType) SelectSQL() string {
	return fmt.Sprintf(`SELECT FirmwareVersion
	 FROM public.`+DATATYPE_ULT_VERSION+` 
	 WHERE DeviceId=%f ;`, res.DeviceId)
}

//InsertSQL
func (res *UltDeviceVersionType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.`+DATATYPE_ULT_VERSION+`  (DeviceId,FirmwareVersion) 
	  VALUES (%f,'%s') 
	  RETURNING DeviceId;`, res.DeviceId, res.FirmwareVersion)
}

//UpdateSQL
func (res *UltDeviceVersionType) UpdateSQL() string {
	return fmt.Sprintf(`UPDATE public.`+DATATYPE_ULT_VERSION+`  
	  SET FirmwareVersion='%s'
	  WHERE DeviceId=%f  
	  RETURNING DeviceId;`,
		res.FirmwareVersion,
		res.DeviceId)
}

//SelectWithDb
func (res *UltDeviceVersionType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(
		&res.FirmwareVersion)
	return errDb
}

//CreateDb
func (res *UltDeviceVersionType) CreateDb(currentDb *sql.DB) {
	createSQL := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS ` + DATATYPE_ULT_VERSION + `  (
	DataId serial PRIMARY KEY,
	DeviceId INT NOT NULL DEFAULT -1,
	FirmwareVersion  varchar(50) NOT NULL DEFAULT '1',
	CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);`)
	_, err := currentDb.Exec(createSQL)
	LogErr(err)
}
