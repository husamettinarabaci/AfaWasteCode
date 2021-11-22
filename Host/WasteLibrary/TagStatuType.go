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
func (res *TagStatuType) GetByRedis() ResultType {

	var resultVal ResultType
	resultVal = GetRedisForStoreApi(REDIS_TAG_STATUS, res.ToIdString())
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
	var resultVal ResultType
	resultVal = SaveRedisForStoreApi(REDIS_TAG_STATUS, res.ToIdString(), res.ToString())
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

//Byte To TagStatuType
func ByteToTagStatuType(retByte []byte) TagStatuType {
	var retVal TagStatuType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To TagStatuType
func StringToTagStatuType(retStr string) TagStatuType {
	return ByteToTagStatuType([]byte(retStr))
}

//ByteToType
func (res *TagStatuType) ByteToType(retByte []byte) {
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *TagStatuType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}

//SelectSQL
func (res *TagStatuType) SelectSQL() string {
	return fmt.Sprintf(`SELECT TagStatu,ImageStatu,CheckTime,ContainerStatu
	 FROM public.tag_status
	 WHERE TagId=%f ;`, res.TagId)
}

//InsertSQL
func (res *TagStatuType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.tag_status (TagId,TagStatu,ImageStatu,CheckTime,ContainerStatu) 
	  VALUES (%f,'%s','%s','%s','%s') 
	  RETURNING TagId;`, res.TagId, res.TagStatu, res.ImageStatu, res.CheckTime, res.ContainerStatu)
}

//UpdateSQL
func (res *TagStatuType) UpdateSQL() string {
	return fmt.Sprintf(`UPDATE public.tag_status 
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
