package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/url"
)

//RecyDeviceNoteType
type RecyDeviceNoteType struct {
	DeviceId float64
	Note     string
	NoteType string
	NoteTime string
	NewData  bool
}

//New
func (res *RecyDeviceNoteType) New() {
	res.DeviceId = 0
	res.Note = ""
	res.NoteType = NOTETYPE_NONE
	res.NoteTime = GetTime()
	res.NewData = false
}

//GetByRedis
func (res *RecyDeviceNoteType) GetByRedis(dbIndex string) ResultType {

	resultVal := GetRedisForStoreApi(dbIndex, REDIS_RECY_NOTE, res.ToIdString())
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
func (res *RecyDeviceNoteType) SaveToRedis() ResultType {
	resultVal := SaveRedisForStoreApi(REDIS_RECY_NOTE, res.ToIdString(), res.ToString())
	return resultVal
}

//SaveToDb
func (res *RecyDeviceNoteType) SaveToDb() ResultType {
	var resultVal ResultType
	var currentHttpHeader HttpClientHeaderType
	currentHttpHeader.New()
	currentHttpHeader.DataType = DATATYPE_RECY_NOTE

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
func (res *RecyDeviceNoteType) SaveToReaderDb() ResultType {
	var resultVal ResultType
	var currentHttpHeader HttpClientHeaderType
	currentHttpHeader.New()
	currentHttpHeader.DataType = DATATYPE_RECY_NOTE

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
func (res *RecyDeviceNoteType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToByte
func (res *RecyDeviceNoteType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//ToString Get JSON
func (res *RecyDeviceNoteType) ToString() string {
	return string(res.ToByte())

}

//ByteToType
func (res *RecyDeviceNoteType) ByteToType(retByte []byte) {
	res.New()
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *RecyDeviceNoteType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}

//SelectSQL
func (res *RecyDeviceNoteType) SelectSQL() string {
	return fmt.Sprintf(`SELECT Note,NoteTime,NoteType
	 FROM public.`+DATATYPE_RECY_NOTE+`  
	 WHERE DeviceId=%f ;`, res.DeviceId)
}

//InsertSQL
func (res *RecyDeviceNoteType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.`+DATATYPE_RECY_NOTE+`  (DeviceId,Note,NoteTime,NoteType) 
	  VALUES (%f,'%s','%s','%s') 
	  RETURNING DeviceId;`, res.DeviceId, res.Note, res.NoteTime, res.NoteType)
}

//UpdateSQL
func (res *RecyDeviceNoteType) UpdateSQL() string {
	return fmt.Sprintf(`UPDATE public.`+DATATYPE_RECY_NOTE+`  
	  SET Note='%s',NoteTime='%s',NoteType='%s'
	  WHERE DeviceId=%f  
	  RETURNING DeviceId;`,
		res.Note,
		res.NoteTime,
		res.NoteType,
		res.DeviceId)
}

//SelectWithDb
func (res *RecyDeviceNoteType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(
		&res.Note,
		&res.NoteTime,
		&res.NoteType)
	return errDb
}

//CreateDb
func (res *RecyDeviceNoteType) CreateDb(currentDb *sql.DB) {
	createSQL := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS ` + DATATYPE_RECY_NOTE + `  ( 
	DataId serial PRIMARY KEY,
	DeviceId INT NOT NULL DEFAULT -1,
	Note varchar(500) NOT NULL DEFAULT '',
	NoteTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	NoteType varchar(50) NOT NULL DEFAULT '` + NOTETYPE_NONE + `',
	CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);`)
	_, err := currentDb.Exec(createSQL)
	LogErr(err)
}
