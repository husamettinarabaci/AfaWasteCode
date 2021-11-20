package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

//NfcBaseType
type NfcBaseType struct {
	NfcId   float64
	Name    string
	SurName string
	NewData bool
}

//New
func (res *NfcBaseType) New() {
	res.NfcId = 0
	res.Name = ""
	res.SurName = ""
	res.NewData = false
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

//SelectSQL
func (res *NfcBaseType) SelectSQL() string {
	return fmt.Sprintf(`SELECT Name,SurName
	 FROM public.nfc_bases
	 WHERE NfcId=%f ;`, res.NfcId)
}

//InsertSQL
func (res *NfcBaseType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.nfc_bases (NfcId,Name,SurName) 
	  VALUES (%f,'%s','%s') 
	  RETURNING NfcId;`, res.NfcId, res.Name, res.SurName)
}

//UpdateSQL
func (res *NfcBaseType) UpdateSQL() string {
	return fmt.Sprintf(`UPDATE public.nfc_bases 
	  SET Name='%s',SurName='%s'
	  WHERE NfcId=%f  
	  RETURNING NfcId;`, res.Name, res.SurName, res.NfcId)
}

//SelectWithDb
func (res *NfcBaseType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(
		&res.Name,
		&res.SurName)
	return errDb
}
