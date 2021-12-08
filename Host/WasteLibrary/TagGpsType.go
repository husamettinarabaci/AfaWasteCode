package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/url"
)

//TagGpsType
type TagGpsType struct {
	TagId     float64
	Latitude  float64
	Longitude float64
	GpsTime   string
	NewData   bool
}

//New
func (res *TagGpsType) New() {
	res.TagId = 0
	res.Latitude = 0
	res.Longitude = 0
	res.GpsTime = GetTime()
	res.NewData = false
}

//GetByRedis
func (res *TagGpsType) GetByRedis(dbIndex string) ResultType {

	var resultVal ResultType
	resultVal = GetRedisForStoreApi(dbIndex, REDIS_TAG_GPS, res.ToIdString())
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
func (res *TagGpsType) SaveToRedis() ResultType {
	var resultVal ResultType
	resultVal = SaveRedisForStoreApi(REDIS_TAG_GPS, res.ToIdString(), res.ToString())
	return resultVal
}

//SaveToDb
func (res *TagGpsType) SaveToDb() ResultType {
	var resultVal ResultType
	var currentHttpHeader HttpClientHeaderType
	currentHttpHeader.New()
	currentHttpHeader.DataType = DATATYPE_TAG_GPS

	data := url.Values{
		HTTP_HEADER: {currentHttpHeader.ToString()},
		HTTP_DATA:   {res.ToString()},
	}
	resultVal = SaveStaticDbMainForStoreApi(data)
	if resultVal.Result == RESULT_OK {
		res.TagId = StringIdToFloat64(resultVal.Retval.(string))
		resultVal.Retval = res.ToString()
	}

	return resultVal
}

//SaveToReaderDb
func (res *TagGpsType) SaveToReaderDb() ResultType {
	var resultVal ResultType
	var currentHttpHeader HttpClientHeaderType
	currentHttpHeader.New()
	currentHttpHeader.DataType = DATATYPE_TAG_GPS

	data := url.Values{
		HTTP_HEADER: {currentHttpHeader.ToString()},
		HTTP_DATA:   {res.ToString()},
	}
	resultVal = SaveReaderDbMainForStoreApi(data)
	if resultVal.Result == RESULT_OK {
		res.TagId = StringIdToFloat64(resultVal.Retval.(string))
		resultVal.Retval = res.ToString()
	}

	return resultVal
}

//ToId String
func (res *TagGpsType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.TagId)
}

//ToByte
func (res *TagGpsType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//ToString Get JSON
func (res *TagGpsType) ToString() string {
	return string(res.ToByte())

}

//Byte To TagGpsType
func ByteToTagGpsType(retByte []byte) TagGpsType {
	var retVal TagGpsType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To TagGpsType
func StringToTagGpsType(retStr string) TagGpsType {
	return ByteToTagGpsType([]byte(retStr))
}

//ByteToType
func (res *TagGpsType) ByteToType(retByte []byte) {
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *TagGpsType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}

//SelectSQL
func (res *TagGpsType) SelectSQL() string {
	return fmt.Sprintf(`SELECT Latitude,Longitude,GpsTime
	 FROM public.`+DATATYPE_TAG_GPS+` 
	 WHERE TagId=%f ;`, res.TagId)
}

//InsertSQL
func (res *TagGpsType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.`+DATATYPE_TAG_GPS+`  (TagId,Latitude,Longitude,GpsTime) 
	  VALUES (%f,%f,%f,'%s') 
	  RETURNING TagId;`, res.TagId, res.Latitude, res.Longitude, res.GpsTime)
}

//UpdateSQL
func (res *TagGpsType) UpdateSQL() string {
	return fmt.Sprintf(`UPDATE public.`+DATATYPE_TAG_GPS+`  
	  SET Latitude=%f,Longitude=%f,GpsTime='%s' 
	  WHERE TagId=%f  
	  RETURNING TagId;`,
		res.Latitude,
		res.Longitude,
		res.GpsTime,
		res.TagId)
}

//SelectWithDb
func (res *TagGpsType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(
		&res.Latitude,
		&res.Longitude,
		&res.GpsTime)
	return errDb
}

//CreateDb
func (res *TagGpsType) CreateDb(currentDb *sql.DB) {
	createSQL := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS ` + DATATYPE_TAG_GPS + `  ( 
	DataId serial PRIMARY KEY,
	TagID INT NOT NULL DEFAULT -1,
	Latitude NUMERIC(14, 11)  NOT NULL DEFAULT 0,
	Longitude NUMERIC(14, 11)  NOT NULL DEFAULT 0,
	GpsTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);`)
	_, err := currentDb.Exec(createSQL)
	LogErr(err)
}
