package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/url"
)

//RfidDeviceEmbededGpsType
type RfidDeviceEmbededGpsType struct {
	DeviceId  float64
	Latitude  float64
	Longitude float64
	Speed     float64
	GpsTime   string
	NewData   bool
}

//New
func (res *RfidDeviceEmbededGpsType) New() {
	res.DeviceId = 0
	res.Latitude = 0
	res.Longitude = 0
	res.Speed = -1
	res.GpsTime = GetTime()
	res.NewData = false
}

//GetByRedis
func (res *RfidDeviceEmbededGpsType) GetByRedis(dbIndex string) ResultType {

	var resultVal ResultType
	resultVal = GetRedisForStoreApi(dbIndex, REDIS_RFID_EMBEDED_GPS, res.ToIdString())
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
func (res *RfidDeviceEmbededGpsType) SaveToRedis() ResultType {
	var resultVal ResultType
	resultVal = SaveRedisForStoreApi(REDIS_RFID_EMBEDED_GPS, res.ToIdString(), res.ToString())
	return resultVal
}

//SaveToDb
func (res *RfidDeviceEmbededGpsType) SaveToDb() ResultType {
	var resultVal ResultType
	var currentHttpHeader HttpClientHeaderType
	currentHttpHeader.New()
	currentHttpHeader.DataType = DATATYPE_RFID_EMBEDED_GPS

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
func (res *RfidDeviceEmbededGpsType) SaveToReaderDb() ResultType {
	var resultVal ResultType
	var currentHttpHeader HttpClientHeaderType
	currentHttpHeader.New()
	currentHttpHeader.DataType = DATATYPE_RFID_EMBEDED_GPS

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
func (res *RfidDeviceEmbededGpsType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToByte
func (res *RfidDeviceEmbededGpsType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res *RfidDeviceEmbededGpsType) ToString() string {
	return string(res.ToByte())

}

//Byte To RfidDeviceEmbededGpsType
func ByteToRfidDeviceEmbededGpsType(retByte []byte) RfidDeviceEmbededGpsType {
	var retVal RfidDeviceEmbededGpsType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To RfidDeviceEmbededGpsType
func StringToRfidDeviceEmbededGpsType(retStr string) RfidDeviceEmbededGpsType {
	return ByteToRfidDeviceEmbededGpsType([]byte(retStr))
}

//ByteToType
func (res *RfidDeviceEmbededGpsType) ByteToType(retByte []byte) {
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *RfidDeviceEmbededGpsType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}

//SelectSQL
func (res *RfidDeviceEmbededGpsType) SelectSQL() string {
	return fmt.Sprintf(`SELECT Latitude,Longitude,Speed,GpsTime
	 FROM public.`+DATATYPE_RFID_EMBEDED_GPS+` 
	 WHERE DeviceId=%f ;`, res.DeviceId)
}

//InsertSQL
func (res *RfidDeviceEmbededGpsType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.`+DATATYPE_RFID_EMBEDED_GPS+`  (DeviceId,Latitude,Longitude,Speed,GpsTime) 
	  VALUES (%f,%f,%f,%f,'%s') 
	  RETURNING DeviceId;`, res.DeviceId, res.Latitude, res.Longitude, res.Speed, res.GpsTime)
}

//UpdateSQL
func (res *RfidDeviceEmbededGpsType) UpdateSQL() string {
	return fmt.Sprintf(`UPDATE public.`+DATATYPE_RFID_EMBEDED_GPS+`  
	  SET Latitude=%f,Longitude=%f,Speed=%f,GpsTime='%s' 
	  WHERE DeviceId=%f  
	  RETURNING DeviceId;`,
		res.Latitude,
		res.Longitude,
		res.Speed,
		res.GpsTime,
		res.DeviceId)
}

//SelectWithDb
func (res *RfidDeviceEmbededGpsType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(
		&res.Latitude,
		&res.Longitude,
		&res.Speed,
		&res.GpsTime)
	return errDb
}

//CreateDb
func (res *RfidDeviceEmbededGpsType) CreateDb(currentDb *sql.DB) {
	createSQL := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS ` + DATATYPE_RFID_EMBEDED_GPS + `  (
	DataId serial PRIMARY KEY,
	DeviceId INT NOT NULL DEFAULT -1,
	Latitude NUMERIC(14, 11)  NOT NULL DEFAULT 0,
	Longitude NUMERIC(14, 11)  NOT NULL DEFAULT 0,
	Speed NUMERIC(14, 11)  NOT NULL DEFAULT 0,
	GpsTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);`)
	_, err := currentDb.Exec(createSQL)
	LogErr(err)
}
