package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/url"
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
func (res *NfcStatuType) GetByRedis(dbIndex string) ResultType {
	var resultVal ResultType
	resultVal = GetRedisForStoreApi(dbIndex, REDIS_NFC_STATUS, res.ToIdString())
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

//SaveToDb
func (res *NfcStatuType) SaveToDb() ResultType {
	var resultVal ResultType
	var currentHttpHeader HttpClientHeaderType
	currentHttpHeader.New()
	currentHttpHeader.DataType = DATATYPE_NFC_STATU

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

//ByteToType
func (res *NfcStatuType) ByteToType(retByte []byte) {
	res.New()
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *NfcStatuType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}

//SelectSQL
func (res *NfcStatuType) SelectSQL() string {
	return fmt.Sprintf(`SELECT NfcStatu,ImageStatu,CheckTime,ItemStatu
	 FROM public.`+DATATYPE_NFC_STATU+` 
	 WHERE NfcId=%f ;`, res.NfcId)
}

//InsertSQL
func (res *NfcStatuType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.`+DATATYPE_NFC_STATU+`  (NfcId,NfcStatu,ImageStatu,CheckTime,ItemStatu) 
	  VALUES (%f,'%s','%s','%s','%s') 
	  RETURNING NfcId;`, res.NfcId, res.NfcStatu, res.ImageStatu, res.CheckTime, res.ItemStatu)
}

//UpdateSQL
func (res *NfcStatuType) UpdateSQL() string {
	return fmt.Sprintf(`UPDATE public.`+DATATYPE_NFC_STATU+`  
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

//CreateDb
func (res *NfcStatuType) CreateDb(currentDb *sql.DB) {
	createSQL := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS ` + DATATYPE_NFC_STATU + `  ( 
	DataId serial PRIMARY KEY,
	NfcID INT NOT NULL DEFAULT -1,
	ItemStatu varchar(50) NOT NULL DEFAULT '` + RECY_ITEM_STATU_NONE + `',
	NfcStatu varchar(50) NOT NULL DEFAULT '` + NFC_STATU_NONE + `',
	ImageStatu varchar(50) NOT NULL DEFAULT '` + STATU_PASSIVE + `',
	CheckTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);`)
	_, err := currentDb.Exec(createSQL)
	LogErr(err)
}
