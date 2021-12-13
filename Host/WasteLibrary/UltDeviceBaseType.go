package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/url"
)

//UltDeviceBaseType
type UltDeviceBaseType struct {
	DeviceId      float64
	ContainerNo   string
	ContainerType string
	NewData       bool
}

//New
func (res *UltDeviceBaseType) New() {
	res.DeviceId = 0
	res.ContainerNo = ""
	res.ContainerType = CONTAINERTYPE_NONE
	res.NewData = false
}

//GetByRedis
func (res *UltDeviceBaseType) GetByRedis(dbIndex string) ResultType {
	resultVal := GetRedisForStoreApi(dbIndex, REDIS_ULT_BASE, res.ToIdString())
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
func (res *UltDeviceBaseType) SaveToRedis() ResultType {
	resultVal := SaveRedisForStoreApi(REDIS_ULT_BASE, res.ToIdString(), res.ToString())
	return resultVal
}

//SaveToDb
func (res *UltDeviceBaseType) SaveToDb() ResultType {
	var resultVal ResultType
	var currentHttpHeader HttpClientHeaderType
	currentHttpHeader.New()
	currentHttpHeader.DataType = DATATYPE_ULT_BASE

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
func (res *UltDeviceBaseType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToByte
func (res *UltDeviceBaseType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res *UltDeviceBaseType) ToString() string {
	return string(res.ToByte())

}

//ByteToType
func (res *UltDeviceBaseType) ByteToType(retByte []byte) {
	res.New()
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *UltDeviceBaseType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}

//SelectSQL
func (res *UltDeviceBaseType) SelectSQL() string {
	return fmt.Sprintf(`SELECT ContainerNo,ContainerType
	 FROM public.`+DATATYPE_ULT_BASE+` 
	 WHERE DeviceId=%f ;`, res.DeviceId)
}

//InsertSQL
func (res *UltDeviceBaseType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.`+DATATYPE_ULT_BASE+`  (DeviceId,ContainerNo,ContainerType) 
	  VALUES (%f,'%s','%s') 
	  RETURNING DeviceId;`, res.DeviceId, res.ContainerNo, res.ContainerType)
}

//UpdateSQL
func (res *UltDeviceBaseType) UpdateSQL() string {
	return fmt.Sprintf(`UPDATE public.`+DATATYPE_ULT_BASE+`  
	  SET ContainerNo='%s',ContainerType='%s'
	  WHERE DeviceId=%f  
	  RETURNING DeviceId;`,
		res.ContainerNo,
		res.ContainerType,
		res.DeviceId)
}

//SelectWithDb
func (res *UltDeviceBaseType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(
		&res.ContainerNo,
		&res.ContainerType)
	return errDb
}

//CreateDb
func (res *UltDeviceBaseType) CreateDb(currentDb *sql.DB) {
	createSQL := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS ` + DATATYPE_ULT_BASE + `  (
	DataId serial PRIMARY KEY,
	DeviceId INT NOT NULL DEFAULT -1,
	ContainerNo  varchar(50) NOT NULL DEFAULT '',
	ContainerType varchar(50) NOT NULL DEFAULT '` + CONTAINERTYPE_NONE + `',
	CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);`)
	_, err := currentDb.Exec(createSQL)
	LogErr(err)
}
