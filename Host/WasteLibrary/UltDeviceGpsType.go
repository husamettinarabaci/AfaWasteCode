package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/url"
)

//UltDeviceGpsType
type UltDeviceGpsType struct {
	DeviceId  float64
	Latitude  float64
	Longitude float64
	GpsTime   string
	NewData   bool
}

//New
func (res *UltDeviceGpsType) New() {
	res.DeviceId = 0
	res.Latitude = 0
	res.Longitude = 0
	res.GpsTime = GetTime()
	res.NewData = false
}

//GetByRedis
func (res *UltDeviceGpsType) GetByRedis(dbIndex string) ResultType {
	var resultVal ResultType
	resultVal = GetRedisForStoreApi(dbIndex, REDIS_ULT_GPS, res.ToIdString())
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
func (res *UltDeviceGpsType) SaveToRedis() ResultType {
	var resultVal ResultType
	resultVal = SaveRedisForStoreApi(REDIS_ULT_GPS, res.ToIdString(), res.ToString())
	return resultVal
}

//SaveToDb
func (res *UltDeviceGpsType) SaveToDb() ResultType {
	var resultVal ResultType
	var currentHttpHeader HttpClientHeaderType
	currentHttpHeader.New()
	currentHttpHeader.DataType = DATATYPE_ULT_GPS

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
func (res *UltDeviceGpsType) SaveToReaderDb() ResultType {
	var resultVal ResultType
	var currentHttpHeader HttpClientHeaderType
	currentHttpHeader.New()
	currentHttpHeader.DataType = DATATYPE_ULT_GPS

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
func (res *UltDeviceGpsType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToByte
func (res *UltDeviceGpsType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res *UltDeviceGpsType) ToString() string {
	return string(res.ToByte())

}

//Byte To UltDeviceGpsType
func ByteToUltDeviceGpsType(retByte []byte) UltDeviceGpsType {
	var retVal UltDeviceGpsType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To UltDeviceGpsType
func StringToUltDeviceGpsType(retStr string) UltDeviceGpsType {
	return ByteToUltDeviceGpsType([]byte(retStr))
}

//ByteToType
func (res *UltDeviceGpsType) ByteToType(retByte []byte) {
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *UltDeviceGpsType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}

//SelectSQL
func (res *UltDeviceGpsType) SelectSQL() string {
	return fmt.Sprintf(`SELECT Latitude,Longitude,GpsTime
	 FROM public.`+DATATYPE_ULT_GPS+` 
	 WHERE DeviceId=%f ;`, res.DeviceId)
}

//InsertSQL
func (res *UltDeviceGpsType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.`+DATATYPE_ULT_GPS+`  (DeviceId,Latitude,Longitude,GpsTime) 
	  VALUES (%f,%f,%f,'%s') 
	  RETURNING DeviceId;`, res.DeviceId, res.Latitude, res.Longitude, res.GpsTime)
}

//UpdateSQL
func (res *UltDeviceGpsType) UpdateSQL() string {
	return fmt.Sprintf(`UPDATE public.`+DATATYPE_ULT_GPS+`  
	  SET Latitude=%f,Longitude=%f,GpsTime='%s' 
	  WHERE DeviceId=%f  
	  RETURNING DeviceId;`,
		res.Latitude,
		res.Longitude,
		res.GpsTime,
		res.DeviceId)
}

//SelectWithDb
func (res *UltDeviceGpsType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(
		&res.Latitude,
		&res.Longitude,
		&res.GpsTime)
	return errDb
}

//CreateDb
func (res *UltDeviceGpsType) CreateDb(currentDb *sql.DB) {
	createSQL := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS ` + DATATYPE_ULT_GPS + `  (
	DataId serial PRIMARY KEY,
	DeviceId INT NOT NULL DEFAULT -1,
	Latitude NUMERIC(14, 11)  NOT NULL DEFAULT 0,
	Longitude NUMERIC(14, 11)  NOT NULL DEFAULT 0,
	GpsTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);`)
	_, err := currentDb.Exec(createSQL)
	LogErr(err)
}
