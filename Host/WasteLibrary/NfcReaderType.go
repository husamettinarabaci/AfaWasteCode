package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/url"
)

//NfcReaderType
type NfcReaderType struct {
	NfcId    float64
	UID      string
	ReadTime string
	NewData  bool
}

//New
func (res *NfcReaderType) New() {
	res.NfcId = 0
	res.UID = ""
	res.ReadTime = GetTime()
	res.NewData = false
}

//GetByRedis
func (res *NfcReaderType) GetByRedis() ResultType {
	var resultVal ResultType
	resultVal = GetRedisForStoreApi(REDIS_NFC_READERS, res.ToIdString())
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
func (res *NfcReaderType) SaveToRedis() ResultType {
	var resultVal ResultType
	resultVal = SaveRedisForStoreApi(REDIS_NFC_READERS, res.ToIdString(), res.ToString())
	return resultVal
}

//SaveToDb
func (res *NfcReaderType) SaveToDb() ResultType {
	var resultVal ResultType
	var currentHttpHeader HttpClientHeaderType
	currentHttpHeader.New()
	currentHttpHeader.DataType = DATATYPE_NFC_READER

	data := url.Values{
		HTTP_HEADER: {currentHttpHeader.ToString()},
		HTTP_DATA:   {res.ToString()},
	}
	resultVal = SaveStaticDbMainForStoreApi(data)
	if resultVal.Result == RESULT_OK {
		res.NfcId = StringIdToFloat64(resultVal.Retval.(string))
		resultVal.Retval = res.ToString()
	}

	return resultVal
}

//ToId String
func (res *NfcReaderType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.NfcId)
}

//ToByte
func (res *NfcReaderType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//ToString Get JSON
func (res *NfcReaderType) ToString() string {
	return string(res.ToByte())

}

//Byte To NfcReaderType
func ByteToNfcReaderType(retByte []byte) NfcReaderType {
	var retVal NfcReaderType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To NfcReaderType
func StringToNfcReaderType(retStr string) NfcReaderType {
	return ByteToNfcReaderType([]byte(retStr))
}

//ByteToType
func (res *NfcReaderType) ByteToType(retByte []byte) {
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *NfcReaderType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}

//SelectSQL
func (res *NfcReaderType) SelectSQL() string {
	return fmt.Sprintf(`SELECT UID,ReadTime
	 FROM public.nfc_readers
	 WHERE NfcId=%f ;`, res.NfcId)
}

//InsertSQL
func (res *NfcReaderType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.nfc_readers (NfcId,UID,ReadTime) 
	  VALUES (%f,'%s','%s') 
	  RETURNING NfcId;`, res.NfcId, res.UID, res.ReadTime)
}

//UpdateSQL
func (res *NfcReaderType) UpdateSQL() string {
	return fmt.Sprintf(`UPDATE public.nfc_readers 
	  SET UID='%s',ReadTime='%s'
	  WHERE NfcId=%f  
	  RETURNING NfcId;`,
		res.UID,
		res.ReadTime,
		res.NfcId)
}

//SelectWithDb
func (res *NfcReaderType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(
		&res.UID,
		&res.ReadTime)
	return errDb
}
