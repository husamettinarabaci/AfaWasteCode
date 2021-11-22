package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

//NfcStatuType
type NfcStatuType struct {
	NfcId      float64
	ItemStatu  string
	NfcStatu   string
	ImageStatu string
	CheckTime  string
	NewData    bool
}

//New
func (res *NfcStatuType) New() {
	res.NfcId = 0
	res.ItemStatu = RECY_ITEM_STATU_NONE
	res.NfcStatu = NFC_STATU_NONE
	res.ImageStatu = STATU_PASSIVE
	res.CheckTime = GetTime()
	res.NewData = false
}

//GetByRedis
func (res *NfcStatuType) GetByRedis() ResultType {
	var resultVal ResultType
	resultVal = GetRedisForStoreApi(REDIS_NFC_STATUS, res.ToIdString())
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
func (res *NfcStatuType) SaveToRedis() ResultType {
	var resultVal ResultType
	resultVal = SaveRedisForStoreApi(REDIS_NFC_STATUS, res.ToIdString(), res.ToString())
	return resultVal
}

//ToId String
func (res *NfcStatuType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.NfcId)
}

//ToByte
func (res *NfcStatuType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//ToString Get JSON
func (res *NfcStatuType) ToString() string {
	return string(res.ToByte())

}

//Byte To NfcStatuType
func ByteToNfcStatuType(retByte []byte) NfcStatuType {
	var retVal NfcStatuType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To NfcStatuType
func StringToNfcStatuType(retStr string) NfcStatuType {
	return ByteToNfcStatuType([]byte(retStr))
}

//ByteToType
func (res *NfcStatuType) ByteToType(retByte []byte) {
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *NfcStatuType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}

//SelectSQL
func (res *NfcStatuType) SelectSQL() string {
	return fmt.Sprintf(`SELECT NfcStatu,ImageStatu,CheckTime,ItemStatu
	 FROM public.nfc_status
	 WHERE NfcId=%f ;`, res.NfcId)
}

//InsertSQL
func (res *NfcStatuType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.nfc_status (NfcId,NfcStatu,ImageStatu,CheckTime,ItemStatu) 
	  VALUES (%f,'%s','%s','%s','%s') 
	  RETURNING NfcId;`, res.NfcId, res.NfcStatu, res.ImageStatu, res.CheckTime, res.ItemStatu)
}

//UpdateSQL
func (res *NfcStatuType) UpdateSQL() string {
	return fmt.Sprintf(`UPDATE public.nfc_status 
	  SET NfcStatu='%s',ImageStatu='%s',CheckTime='%s',ItemStatu='%s'
	  WHERE NfcId=%f  
	  RETURNING NfcId;`,
		res.NfcStatu,
		res.ImageStatu,
		res.CheckTime,
		res.ItemStatu,
		res.NfcId)
}

//SelectWithDb
func (res *NfcStatuType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(
		&res.NfcStatu,
		&res.ImageStatu,
		&res.CheckTime,
		&res.ItemStatu)
	return errDb
}
