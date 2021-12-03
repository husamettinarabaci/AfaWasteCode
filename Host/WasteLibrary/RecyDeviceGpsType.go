package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/url"
)

//RecyDeviceGpsType
type RecyDeviceGpsType struct {
	DeviceId  float64
	Latitude  float64
	Longitude float64
	GpsTime   string
	NewData   bool
}

//New
func (res *RecyDeviceGpsType) New() {
	res.DeviceId = 0
	res.Latitude = 0
	res.Longitude = 0
	res.GpsTime = GetTime()
	res.NewData = false
}

//GetByRedis
func (res *RecyDeviceGpsType) GetByRedis(dbIndex string) ResultType {
	var resultVal ResultType
	resultVal = GetRedisForStoreApi(dbIndex, REDIS_RECY_GPS, res.ToIdString())
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
func (res *RecyDeviceGpsType) SaveToRedis() ResultType {
	var resultVal ResultType
	resultVal = SaveRedisForStoreApi(REDIS_RECY_GPS, res.ToIdString(), res.ToString())
	return resultVal
}

//SaveToDb
func (res *RecyDeviceGpsType) SaveToDb() ResultType {
	var resultVal ResultType
	var currentHttpHeader HttpClientHeaderType
	currentHttpHeader.New()
	currentHttpHeader.DataType = DATATYPE_RECY_GPS

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
func (res *RecyDeviceGpsType) SaveToReaderDb() ResultType {
	var resultVal ResultType
	var currentHttpHeader HttpClientHeaderType
	currentHttpHeader.New()
	currentHttpHeader.DataType = DATATYPE_RECY_GPS

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
func (res *RecyDeviceGpsType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToByte
func (res *RecyDeviceGpsType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res *RecyDeviceGpsType) ToString() string {
	return string(res.ToByte())

}

//Byte To RecyDeviceGpsType
func ByteToRecyDeviceGpsType(retByte []byte) RecyDeviceGpsType {
	var retVal RecyDeviceGpsType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To RecyDeviceGpsType
func StringToRecyDeviceGpsType(retStr string) RecyDeviceGpsType {
	return ByteToRecyDeviceGpsType([]byte(retStr))
}

//ByteToType
func (res *RecyDeviceGpsType) ByteToType(retByte []byte) {
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *RecyDeviceGpsType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}

//SelectSQL
func (res *RecyDeviceGpsType) SelectSQL() string {
	return fmt.Sprintf(`SELECT Latitude,Longitude,GpsTime
	 FROM public.`+DATATYPE_RECY_GPS+` 
	 WHERE DeviceId=%f ;`, res.DeviceId)
}

//InsertSQL
func (res *RecyDeviceGpsType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.`+DATATYPE_RECY_GPS+`  (DeviceId,Latitude,Longitude,GpsTime) 
	  VALUES (%f,%f,%f,'%s') 
	  RETURNING DeviceId;`, res.DeviceId, res.Latitude, res.Longitude, res.GpsTime)
}

//UpdateSQL
func (res *RecyDeviceGpsType) UpdateSQL() string {
	return fmt.Sprintf(`UPDATE public.`+DATATYPE_RECY_GPS+`  
	  SET Latitude=%f,Longitude=%f,GpsTime='%s' 
	  WHERE DeviceId=%f  
	  RETURNING DeviceId;`,
		res.Latitude,
		res.Longitude,
		res.GpsTime,
		res.DeviceId)
}

//SelectWithDb
func (res *RecyDeviceGpsType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(
		&res.Latitude,
		&res.Longitude,
		&res.GpsTime)
	return errDb
}
