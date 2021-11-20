package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
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
