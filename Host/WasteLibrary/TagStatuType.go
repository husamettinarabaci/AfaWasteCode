package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/url"
)

//TagStatuType
type TagStatuType struct {
	TagId          float64
	ContainerStatu string
	TagStatu       string
	ImageStatu     string
	CheckTime      string
	NewData        bool
}

//New
func (res *TagStatuType) New() {
	res.TagId = 0
	res.ContainerStatu = CONTAINER_FULLNESS_STATU_NONE
	res.TagStatu = TAG_STATU_NONE
	res.ImageStatu = STATU_PASSIVE
	res.CheckTime = GetTime()
	res.NewData = false
}

//GetByRedis
func (res *TagStatuType) GetByRedis(dbIndex string) ResultType {

	resultVal := GetRedisForStoreApi(dbIndex, REDIS_TAG_STATU, res.ToIdString())
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
func (res *TagStatuType) SaveToRedis() ResultType {
	resultVal := SaveRedisForStoreApi(REDIS_TAG_STATU, res.ToIdString(), res.ToString())
	return resultVal
}

//SaveToDb
func (res *TagStatuType) SaveToDb() ResultType {
	var resultVal ResultType
	var currentHttpHeader HttpClientHeaderType
	currentHttpHeader.New()
	currentHttpHeader.DataType = DATATYPE_TAG_STATU

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
func (res *TagStatuType) SaveToReaderDb() ResultType {
	var resultVal ResultType
	var currentHttpHeader HttpClientHeaderType
	currentHttpHeader.New()
	currentHttpHeader.DataType = DATATYPE_TAG_STATU

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
func (res *TagStatuType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.TagId)
}

//ToByte
func (res *TagStatuType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//ToString Get JSON
func (res *TagStatuType) ToString() string {
	return string(res.ToByte())

}

//ByteToType
func (res *TagStatuType) ByteToType(retByte []byte) {
	res.New()
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *TagStatuType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}

//SelectSQL
func (res *TagStatuType) SelectSQL() string {
	return fmt.Sprintf(`SELECT TagStatu,ImageStatu,CheckTime,ContainerStatu
	 FROM public.`+DATATYPE_TAG_STATU+` 
	 WHERE TagId=%f ;`, res.TagId)
}

//InsertSQL
func (res *TagStatuType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.`+DATATYPE_TAG_STATU+`  (TagId,TagStatu,ImageStatu,CheckTime,ContainerStatu) 
	  VALUES (%f,'%s','%s','%s','%s') 
	  RETURNING TagId;`, res.TagId, res.TagStatu, res.ImageStatu, res.CheckTime, res.ContainerStatu)
}

//UpdateSQL
func (res *TagStatuType) UpdateSQL() string {
	return fmt.Sprintf(`UPDATE public.`+DATATYPE_TAG_STATU+`  
	  SET TagStatu='%s',ImageStatu='%s',CheckTime='%s',ContainerStatu='%s'
	  WHERE TagId=%f  
	  RETURNING TagId;`,
		res.TagStatu,
		res.ImageStatu,
		res.CheckTime,
		res.ContainerStatu,
		res.TagId)
}

//SelectWithDb
func (res *TagStatuType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(
		&res.TagStatu,
		&res.ImageStatu,
		&res.CheckTime,
		&res.ContainerStatu)
	return errDb
}

//CreateDb
func (res *TagStatuType) CreateDb(currentDb *sql.DB) {
	createSQL := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS ` + DATATYPE_TAG_STATU + `  ( 
	DataId serial PRIMARY KEY,
	TagID INT NOT NULL DEFAULT -1,
	ContainerStatu varchar(50) NOT NULL DEFAULT '` + CONTAINER_FULLNESS_STATU_NONE + `',
	TagStatu varchar(50) NOT NULL DEFAULT '` + TAG_STATU_NONE + `',
	ImageStatu varchar(50) NOT NULL DEFAULT '` + STATU_PASSIVE + `',
	CheckTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);`)
	_, err := currentDb.Exec(createSQL)
	LogErr(err)
}
