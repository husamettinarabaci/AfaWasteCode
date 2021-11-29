package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/url"
)

//TagBaseType
type TagBaseType struct {
	TagId         float64
	ContainerNo   string
	ContainerType string
	NewData       bool
}

//New
func (res *TagBaseType) New() {
	res.TagId = 0
	res.ContainerNo = ""
	res.ContainerType = CONTAINERTYPE_NONE
	res.NewData = false
}

//GetByRedis
func (res *TagBaseType) GetByRedis(dbIndex int) ResultType {

	var resultVal ResultType
	resultVal = GetRedisForStoreApi(REDIS_TAG_BASES, res.ToIdString())
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
func (res *TagBaseType) SaveToRedis() ResultType {
	var resultVal ResultType
	resultVal = SaveRedisForStoreApi(REDIS_TAG_BASES, res.ToIdString(), res.ToString())
	return resultVal
}

//SaveToDb
func (res *TagBaseType) SaveToDb() ResultType {
	var resultVal ResultType
	var currentHttpHeader HttpClientHeaderType
	currentHttpHeader.New()
	currentHttpHeader.DataType = DATATYPE_TAG_BASE

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
func (res *TagBaseType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.TagId)
}

//ToByte
func (res *TagBaseType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//ToString Get JSON
func (res *TagBaseType) ToString() string {
	return string(res.ToByte())

}

//Byte To TagBaseType
func ByteToTagBaseType(retByte []byte) TagBaseType {
	var retVal TagBaseType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To TagBaseType
func StringToTagBaseType(retStr string) TagBaseType {
	return ByteToTagBaseType([]byte(retStr))
}

//ByteToType
func (res *TagBaseType) ByteToType(retByte []byte) {
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *TagBaseType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}

//SelectSQL
func (res *TagBaseType) SelectSQL() string {
	return fmt.Sprintf(`SELECT ContainerNo,ContainerType
	 FROM public.tag_bases
	 WHERE TagId=%f ;`, res.TagId)
}

//InsertSQL
func (res *TagBaseType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.tag_bases (TagId,ContainerNo,ContainerType) 
	  VALUES (%f,'%s','%s') 
	  RETURNING TagId;`, res.TagId, res.ContainerNo, res.ContainerType)
}

//UpdateSQL
func (res *TagBaseType) UpdateSQL() string {
	return fmt.Sprintf(`UPDATE public.tag_bases 
	  SET ContainerNo='%s',ContainerType='%s'
	  WHERE TagId=%f  
	  RETURNING TagId;`,
		res.ContainerNo,
		res.ContainerType,
		res.TagId)
}

//SelectWithDb
func (res *TagBaseType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(
		&res.ContainerNo,
		&res.ContainerType)
	return errDb
}
