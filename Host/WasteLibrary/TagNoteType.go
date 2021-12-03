package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/url"
)

//TagNoteType
type TagNoteType struct {
	TagId    float64
	Note     string
	NoteType string
	NoteTime string
	NewData  bool
}

//New
func (res *TagNoteType) New() {
	res.TagId = 0
	res.Note = ""
	res.NoteType = NOTETYPE_NONE
	res.NoteTime = GetTime()
	res.NewData = false
}

//GetByRedis
func (res *TagNoteType) GetByRedis(dbIndex string) ResultType {

	var resultVal ResultType
	resultVal = GetRedisForStoreApi(dbIndex, REDIS_TAG_NOTES, res.ToIdString())
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
func (res *TagNoteType) SaveToRedis() ResultType {
	var resultVal ResultType
	resultVal = SaveRedisForStoreApi(REDIS_TAG_NOTES, res.ToIdString(), res.ToString())
	return resultVal
}

//SaveToDb
func (res *TagNoteType) SaveToDb() ResultType {
	var resultVal ResultType
	var currentHttpHeader HttpClientHeaderType
	currentHttpHeader.New()
	currentHttpHeader.DataType = DATATYPE_TAG_NOTE

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
func (res *TagNoteType) SaveToReaderDb() ResultType {
	var resultVal ResultType
	var currentHttpHeader HttpClientHeaderType
	currentHttpHeader.New()
	currentHttpHeader.DataType = DATATYPE_TAG_NOTE

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
func (res *TagNoteType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.TagId)
}

//ToByte
func (res *TagNoteType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//ToString Get JSON
func (res *TagNoteType) ToString() string {
	return string(res.ToByte())

}

//Byte To TagNoteType
func ByteToTagNoteType(retByte []byte) TagNoteType {
	var retVal TagNoteType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To TagNoteType
func StringToTagNoteType(retStr string) TagNoteType {
	return ByteToTagNoteType([]byte(retStr))
}

//ByteToType
func (res *TagNoteType) ByteToType(retByte []byte) {
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *TagNoteType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}

//SelectSQL
func (res *TagNoteType) SelectSQL() string {
	return fmt.Sprintf(`SELECT Note,NoteTime,NoteType
	 FROM public.`+DATATYPE_TAG_NOTE+` 
	 WHERE TagId=%f ;`, res.TagId)
}

//InsertSQL
func (res *TagNoteType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.`+DATATYPE_TAG_NOTE+`  (TagId,Note,NoteTime,NoteType) 
	  VALUES (%f,'%s','%s','%s') 
	  RETURNING TagId;`, res.TagId, res.Note, res.NoteTime, res.NoteType)
}

//UpdateSQL
func (res *TagNoteType) UpdateSQL() string {
	return fmt.Sprintf(`UPDATE public.`+DATATYPE_TAG_NOTE+`  
	  SET Note='%s',NoteTime='%s',NoteType='%s'
	  WHERE TagId=%f  
	  RETURNING TagId;`,
		res.Note,
		res.NoteTime,
		res.NoteType,
		res.TagId)
}

//SelectWithDb
func (res *TagNoteType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(
		&res.Note,
		&res.NoteTime,
		&res.NoteType)
	return errDb
}
