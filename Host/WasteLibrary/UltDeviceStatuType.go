package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/url"
)

//UltDeviceStatuType
type UltDeviceStatuType struct {
	DeviceId        float64
	StatusTime      string
	AliveStatus     string
	AliveLastOkTime string
	UltStatus       string
	ContainerStatu  string
	SensPercent     float64
	NewData         bool
}

//New
func (res *UltDeviceStatuType) New() {
	res.DeviceId = 0
	res.StatusTime = GetTime()
	res.ContainerStatu = CONTAINER_FULLNESS_STATU_NONE
	res.UltStatus = ULT_STATU_NONE
	res.AliveStatus = STATU_PASSIVE
	res.AliveLastOkTime = GetTime()
	res.SensPercent = 0
	res.NewData = false
}

//GetByRedis
func (res *UltDeviceStatuType) GetByRedis(dbIndex string) ResultType {
	resultVal := GetRedisForStoreApi(dbIndex, REDIS_ULT_STATU, res.ToIdString())
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
func (res *UltDeviceStatuType) SaveToRedis() ResultType {
	resultVal := SaveRedisForStoreApi(REDIS_ULT_STATU, res.ToIdString(), res.ToString())
	return resultVal
}

//SaveToDb
func (res *UltDeviceStatuType) SaveToDb() ResultType {
	var resultVal ResultType
	var currentHttpHeader HttpClientHeaderType
	currentHttpHeader.New()
	currentHttpHeader.DataType = DATATYPE_ULT_STATU

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
func (res *UltDeviceStatuType) SaveToReaderDb() ResultType {
	var resultVal ResultType
	var currentHttpHeader HttpClientHeaderType
	currentHttpHeader.New()
	currentHttpHeader.DataType = DATATYPE_ULT_STATU

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
func (res *UltDeviceStatuType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToByte
func (res *UltDeviceStatuType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res *UltDeviceStatuType) ToString() string {
	return string(res.ToByte())

}

//ByteToType
func (res *UltDeviceStatuType) ByteToType(retByte []byte) {
	res.New()
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *UltDeviceStatuType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}

//SelectSQL
func (res *UltDeviceStatuType) SelectSQL() string {
	return fmt.Sprintf(`SELECT StatusTime,AliveStatus,AliveLastOkTime,ContainerStatu,UltStatus,SensPercent
	 FROM public.`+DATATYPE_ULT_STATU+` 
	 WHERE DeviceId=%f ;`, res.DeviceId)
}

//InsertSQL
func (res *UltDeviceStatuType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.`+DATATYPE_ULT_STATU+`  (DeviceId,StatusTime,AliveStatus,AliveLastOkTime,ContainerStatu,UltStatus,SensPercent) 
	  VALUES (%f,'%s','%s','%s','%s','%s',%f) 
	  RETURNING DeviceId;`, res.DeviceId,
		res.StatusTime, res.AliveStatus, res.AliveLastOkTime, res.ContainerStatu, res.UltStatus, res.SensPercent)
}

//UpdateSQL
func (res *UltDeviceStatuType) UpdateSQL() string {
	var execSqlExt = ""
	if res.AliveStatus == STATU_ACTIVE {
		execSqlExt += ",AliveLastOkTime='" + res.AliveLastOkTime + "'"
	}

	return fmt.Sprintf(`UPDATE public.`+DATATYPE_ULT_STATU+`  
	  SET StatusTime='%s',AliveStatus='%s',ContainerStatu='%s',UltStatus='%s',SensPercent=%f`+execSqlExt+`
	  WHERE DeviceId=%f  
	  RETURNING DeviceId;`,
		res.StatusTime, res.AliveStatus, res.ContainerStatu, res.UltStatus, res.SensPercent, res.DeviceId)
}

//SelectWithDb
func (res *UltDeviceStatuType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(
		&res.StatusTime,
		&res.AliveStatus,
		&res.AliveLastOkTime,
		&res.ContainerStatu,
		&res.UltStatus,
		&res.SensPercent)
	return errDb
}

//CreateDb
func (res *UltDeviceStatuType) CreateDb(currentDb *sql.DB) {
	createSQL := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS ` + DATATYPE_ULT_STATU + `  (
	DataId serial PRIMARY KEY,
	DeviceId INT NOT NULL DEFAULT -1,
	StatusTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	AliveStatus varchar(50) NOT NULL DEFAULT '` + STATU_PASSIVE + `',
	AliveLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	UltStatus varchar(50) NOT NULL DEFAULT '` + ULT_STATU_NONE + `',
	SensPercent NUMERIC(14, 11)  NOT NULL DEFAULT 0,
	ContainerStatu varchar(50) NOT NULL DEFAULT '` + CONTAINER_FULLNESS_STATU_NONE + `',
	CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);`)
	_, err := currentDb.Exec(createSQL)
	LogErr(err)
}
