package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/url"
)

//TagMainType
type TagMainType struct {
	TagId      float64
	CustomerId float64
	DeviceId   float64
	Epc        string
	Active     string
	CreateTime string
}

//New
func (res *TagMainType) New() {
	res.TagId = 0
	res.DeviceId = 0
	res.CustomerId = 1
	res.TagId = 0
	res.Epc = ""
	res.Active = STATU_ACTIVE
	res.CreateTime = GetTime()
}

//GetByRedis
func (res *TagMainType) GetByRedis(dbIndex string) ResultType {

	resultVal := GetRedisForStoreApi(dbIndex, REDIS_TAG_MAIN, res.ToIdString())
	if resultVal.Result == RESULT_OK {
		res.StringToType(resultVal.Retval.(string))
	} else {
		return resultVal
	}

	resultVal.Retval = res.ToString()
	return resultVal
}

//SaveToRedis
func (res *TagMainType) SaveToRedis() ResultType {
	resultVal := SaveRedisForStoreApi(REDIS_TAG_MAIN, res.ToIdString(), res.ToString())
	return resultVal
}

//SaveToRedisByEpc
func (res *TagMainType) SaveToRedisByEpc() ResultType {
	resultVal := SaveRedisForStoreApi(REDIS_TAG_EPC, res.Epc, res.ToIdString())
	return resultVal
}

//SaveToDb
func (res *TagMainType) SaveToDb() ResultType {
	var resultVal ResultType
	var currentHttpHeader HttpClientHeaderType
	currentHttpHeader.New()
	currentHttpHeader.DataType = DATATYPE_TAG_MAIN

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

//ToId String
func (res *TagMainType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.TagId)
}

//ToCustomerId String
func (res *TagMainType) ToCustomerIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToDeviceId String
func (res *TagMainType) ToDeviceIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToByte
func (res *TagMainType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//ToString Get JSON
func (res *TagMainType) ToString() string {
	return string(res.ToByte())

}

//ByteToType
func (res *TagMainType) ByteToType(retByte []byte) {
	res.New()
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *TagMainType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}

//SelectSQL
func (res *TagMainType) SelectSQL() string {
	return fmt.Sprintf(`SELECT CustomerId,DeviceId,Epc,Active,CreateTime
	 FROM public.`+DATATYPE_TAG_MAIN+` 
	 WHERE TagId=%f  ;`, res.TagId)
}

//InsertSQL
func (res *TagMainType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.`+DATATYPE_TAG_MAIN+`  (CustomerId,DeviceId,Epc) 
	  VALUES (%f,%f,'%s') 
	  RETURNING TagId;`, res.CustomerId, res.DeviceId, res.Epc)
}

//InsertDataSQL
func (res *TagMainType) InsertDataSQL() string {
	return fmt.Sprintf(`INSERT INTO public.`+DATATYPE_TAG_MAIN+`  (TagId,CustomerId,DeviceId,Epc) 
	  VALUES (%f,%f,%f,'%s') 
	  RETURNING TagId;`, res.TagId, res.CustomerId, res.DeviceId, res.Epc)
}

//UpdateSQL
func (res *TagMainType) UpdateSQL() string {
	return fmt.Sprintf(`UPDATE public.`+DATATYPE_TAG_MAIN+`  
	  SET CustomerId=%f,DeviceId=%f,Epc='%s' 
	  WHERE TagId=%f  
	  RETURNING TagId;`,
		res.CustomerId,
		res.DeviceId,
		res.Epc,
		res.TagId)
}

//SelectWithDb
func (res *TagMainType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(
		&res.CustomerId,
		&res.DeviceId,
		&res.Epc,
		&res.Active,
		&res.CreateTime)
	return errDb
}

//CreateDb
func (res *TagMainType) CreateDb(currentDb *sql.DB) {
	createSQL := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS ` + DATATYPE_TAG_MAIN + `  ( 
	TagID serial PRIMARY KEY,
	CustomerId INT NOT NULL DEFAULT -1,
	DeviceId INT NOT NULL DEFAULT -1,
	Epc varchar(50) NOT NULL DEFAULT '',
	Active varchar(50) NOT NULL DEFAULT '` + STATU_ACTIVE + `',
	CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);`)
	_, err := currentDb.Exec(createSQL)
	LogErr(err)
}

//CreateDb
func (res *TagMainType) CreateReaderDb(currentDb *sql.DB) {
	createSQL := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS ` + DATATYPE_TAG_MAIN + `  ( 
	DataId serial PRIMARY KEY,
	TagID INT NOT NULL DEFAULT -1,
	CustomerId INT NOT NULL DEFAULT -1,
	DeviceId INT NOT NULL DEFAULT -1,
	Epc varchar(50) NOT NULL DEFAULT '',
	Active varchar(50) NOT NULL DEFAULT '` + STATU_ACTIVE + `',
	CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);`)
	_, err := currentDb.Exec(createSQL)
	LogErr(err)
}
