package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

//TagReaderType
type TagReaderType struct {
	TagId    float64
	UID      string
	ReadTime string
	NewData  bool
}

//New
func (res *TagReaderType) New() {
	res.TagId = 0
	res.UID = ""
	res.ReadTime = GetTime()
	res.NewData = false
}

//GetByRedis
func (res *TagReaderType) GetByRedis() ResultType {

	var resultVal ResultType
	resultVal = GetRedisForStoreApi(REDIS_TAG_READERS, res.ToIdString())
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
func (res *TagReaderType) SaveToRedis() ResultType {
	var resultVal ResultType
	resultVal = SaveRedisForStoreApi(REDIS_TAG_READERS, res.ToIdString(), res.ToString())
	return resultVal
}

//ToId String
func (res *TagReaderType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.TagId)
}

//ToByte
func (res *TagReaderType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//ToString Get JSON
func (res *TagReaderType) ToString() string {
	return string(res.ToByte())

}

//Byte To TagReaderType
func ByteToTagReaderType(retByte []byte) TagReaderType {
	var retVal TagReaderType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To TagReaderType
func StringToTagReaderType(retStr string) TagReaderType {
	return ByteToTagReaderType([]byte(retStr))
}

//ByteToType
func (res *TagReaderType) ByteToType(retByte []byte) {
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *TagReaderType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}

//SelectSQL
func (res *TagReaderType) SelectSQL() string {
	return fmt.Sprintf(`SELECT UID,ReadTime
	 FROM public.tag_readers
	 WHERE TagId=%f ;`, res.TagId)
}

//InsertSQL
func (res *TagReaderType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.tag_readers (TagId,UID,ReadTime) 
	  VALUES (%f,'%s','%s') 
	  RETURNING TagId;`, res.TagId, res.UID, res.ReadTime)
}

//UpdateSQL
func (res *TagReaderType) UpdateSQL() string {
	return fmt.Sprintf(`UPDATE public.tag_readers 
	  SET UID='%s',ReadTime='%s'
	  WHERE TagId=%f  
	  RETURNING TagId;`,
		res.UID,
		res.ReadTime,
		res.TagId)
}

//SelectWithDb
func (res *TagReaderType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(
		&res.UID,
		&res.ReadTime)
	return errDb
}
