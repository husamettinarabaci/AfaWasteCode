package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
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
func (res *NfcBaseType) GetByRedis() ResultType {
	var resultVal ResultType
	resultVal = GetRedisForStoreApi(REDIS_NFC_BASES, res.ToIdString())
	if resultVal.Result == RESULT_OK {
		res.StringToType(resultVal.Retval.(string))
		res.NewData = false
	} else {
		return resultVal
	}

	resultVal.Retval = res.ToString()
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
	 FROM public.nfc_bases
	 WHERE NfcId=%f ;`, res.NfcId)
}

//InsertSQL
func (res *NfcBaseType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.nfc_bases (NfcId,Name,SurName,TotalAmount,LastAmount) 
	  VALUES (%f,'%s','%s',%f,%f) 
	  RETURNING NfcId;`, res.NfcId, res.Name, res.SurName, res.TotalAmount, res.LastAmount)
}

//UpdateSQL
func (res *NfcBaseType) UpdateSQL() string {
	return fmt.Sprintf(`UPDATE public.nfc_bases 
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
