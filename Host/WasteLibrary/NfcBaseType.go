package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/url"
)

//NfcBaseType
type NfcBaseType struct {
	NfcId       float64
	Name        string
	SurName     string
	TotalAmount float64
	LastAmount  float64
	NewData     bool
}

//New
func (res *NfcBaseType) New() {
	res.NfcId = 0
	res.Name = ""
	res.SurName = ""
	res.TotalAmount = 0
	res.LastAmount = 0
	res.NewData = false
}

//GetByRedis
func (res *NfcBaseType) GetByRedis(dbIndex string) ResultType {
	var resultVal ResultType
	resultVal = GetRedisForStoreApi(dbIndex, REDIS_NFC_BASES, res.ToIdString())
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
func (res *NfcBaseType) SaveToRedis() ResultType {
	var resultVal ResultType
	resultVal = SaveRedisForStoreApi(REDIS_NFC_BASES, res.ToIdString(), res.ToString())
	return resultVal
}

//SaveToDb
func (res *NfcBaseType) SaveToDb() ResultType {
	var resultVal ResultType
	var currentHttpHeader HttpClientHeaderType
	currentHttpHeader.New()
	currentHttpHeader.DataType = DATATYPE_NFC_BASE

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
func (res *NfcBaseType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.NfcId)
}

//ToByte
func (res *NfcBaseType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//ToString Get JSON
func (res *NfcBaseType) ToString() string {
	return string(res.ToByte())

}

//Byte To NfcBaseType
func ByteToNfcBaseType(retByte []byte) NfcBaseType {
	var retVal NfcBaseType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To NfcBaseType
func StringToNfcBaseType(retStr string) NfcBaseType {
	return ByteToNfcBaseType([]byte(retStr))
}

//ByteToType
func (res *NfcBaseType) ByteToType(retByte []byte) {
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *NfcBaseType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}

//SelectSQL
func (res *NfcBaseType) SelectSQL() string {
	return fmt.Sprintf(`SELECT Name,SurName,TotalAmount,LastAmount
	 FROM public.`+DATATYPE_NFC_BASE+` 
	 WHERE NfcId=%f ;`, res.NfcId)
}

//InsertSQL
func (res *NfcBaseType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.`+DATATYPE_NFC_BASE+`  (NfcId,Name,SurName,TotalAmount,LastAmount) 
	  VALUES (%f,'%s','%s',%f,%f) 
	  RETURNING NfcId;`, res.NfcId, res.Name, res.SurName, res.TotalAmount, res.LastAmount)
}

//UpdateSQL
func (res *NfcBaseType) UpdateSQL() string {
	return fmt.Sprintf(`UPDATE public.`+DATATYPE_NFC_BASE+`  
	  SET Name='%s',SurName='%s',TotalAmount=%f,LastAmount=%f
	  WHERE NfcId=%f  
	  RETURNING NfcId;`, res.Name, res.SurName, res.TotalAmount, res.LastAmount, res.NfcId)
}

//SelectWithDb
func (res *NfcBaseType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(
		&res.Name,
		&res.SurName,
		&res.TotalAmount,
		&res.LastAmount)
	return errDb
}
