package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

//NfcMainType
type NfcMainType struct {
	NfcId      float64
	CustomerId float64
	DeviceId   float64
	Epc        string
	Active     string
	CreateTime string
}

//GetByRedis
func (res *NfcMainType) GetByRedis() ResultType {
	var resultVal ResultType
	resultVal = GetRedisForStoreApi(REDIS_NFC_MAINS, res.ToIdString())
	if resultVal.Result == RESULT_OK {
		res.StringToType(resultVal.Retval.(string))
	} else {
		return resultVal
	}

	resultVal.Retval = res.ToString()
	return resultVal
}

//New
func (res *NfcMainType) New() {
	res.NfcId = 0
	res.CustomerId = 1
	res.NfcId = 0
	res.Epc = ""
	res.Active = STATU_ACTIVE
	res.CreateTime = GetTime()
}

//ToId String
func (res *NfcMainType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.NfcId)
}

//ToCustomerId String
func (res *NfcMainType) ToCustomerIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToNfcId String
func (res *NfcMainType) ToNfcIdString() string {
	return fmt.Sprintf("%.0f", res.NfcId)
}

//ToByte
func (res *NfcMainType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//ToString Get JSON
func (res *NfcMainType) ToString() string {
	return string(res.ToByte())

}

//Byte To NfcMainType
func ByteToNfcMainType(retByte []byte) NfcMainType {
	var retVal NfcMainType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To NfcMainType
func StringToNfcMainType(retStr string) NfcMainType {
	return ByteToNfcMainType([]byte(retStr))
}

//ByteToType
func (res *NfcMainType) ByteToType(retByte []byte) {
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *NfcMainType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}

//SelectSQL
func (res *NfcMainType) SelectSQL() string {
	return fmt.Sprintf(`SELECT CustomerId,DeviceId,Epc,Active,CreateTime
	 FROM public.nfc_mains
	 WHERE NfcId=%f  ;`, res.NfcId)
}

//InsertSQL
func (res *NfcMainType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.nfc_mains (CustomerId,DeviceId,Epc) 
	  VALUES (%f,%f,'%s') 
	  RETURNING NfcId;`, res.CustomerId, res.DeviceId, res.Epc)
}

//InsertDataSQL
func (res *NfcMainType) InsertDataSQL() string {
	return fmt.Sprintf(`INSERT INTO public.nfc_mains (NfcId,CustomerId,DeviceId,Epc) 
	  VALUES (%f,%f,%f,'%s') 
	  RETURNING NfcId;`, res.NfcId, res.CustomerId, res.DeviceId, res.Epc)
}

//UpdateSQL
func (res *NfcMainType) UpdateSQL() string {
	return fmt.Sprintf(`UPDATE public.nfc_mains 
	  SET CustomerId=%f,DeviceId=%f,Epc='%s' 
	  WHERE NfcId=%f  
	  RETURNING NfcId;`,
		res.CustomerId,
		res.DeviceId,
		res.Epc,
		res.NfcId)
}

//SelectWithDb
func (res *NfcMainType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(
		&res.CustomerId,
		&res.DeviceId,
		&res.Epc,
		&res.Active,
		&res.CreateTime)
	return errDb
}
