package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

//TagType
type TagType struct {
	TagId      float64
	CustomerId float64
	DeviceId   float64
	Epc        string
	TagBase    TagBaseType
	TagStatu   TagStatuType
	TagGps     TagGpsType
	TagReader  TagReaderType
	Active     string
	CreateTime string
}

//New
func (res *TagType) New() {
	res.TagId = 0
	res.CustomerId = 1
	res.TagId = 0
	res.Epc = ""
	res.TagBase.New()
	res.TagStatu.New()
	res.TagGps.New()
	res.TagReader.New()
	res.Active = STATU_ACTIVE
	res.CreateTime = GetTime()
}

//ToId String
func (res TagType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.TagId)
}

//ToCustomerId String
func (res TagType) ToCustomerIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToTagId String
func (res TagType) ToTagIdString() string {
	return fmt.Sprintf("%.0f", res.TagId)
}

//ToByte
func (res TagType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//ToString Get JSON
func (res TagType) ToString() string {
	return string(res.ToByte())

}

//Byte To TagType
func ByteToTagType(retByte []byte) TagType {
	var retVal TagType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To TagType
func StringToTagType(retStr string) TagType {
	return ByteToTagType([]byte(retStr))
}

//SelectSQL
func (res TagType) SelectSQL() string {
	return fmt.Sprintf(`SELECT CustomerId,DeviceId,Epc,Active,CreateTime
	 FROM public.tags
	 WHERE TagId=%f ;`, res.TagId)
}

//InsertSQL
func (res TagType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.tags (CustomerId,DeviceId,Epc) 
	  VALUES (%f,%f,'%s') 
	  RETURNING TagId;`, res.CustomerId, res.DeviceId, res.Epc)
}

//InsertDataSQL
func (res TagType) InsertDataSQL() string {
	return fmt.Sprintf(`INSERT INTO public.tags (TagId,CustomerId,DeviceId,Epc) 
	  VALUES (%f,%f,%f,'%s') 
	  RETURNING TagId;`, res.TagId, res.CustomerId, res.DeviceId, res.Epc)
}

//UpdateSQL
func (res TagType) UpdateSQL() string {
	return fmt.Sprintf(`UPDATE public.tags 
	  SET CustomerId=%f,DeviceId=%f,Epc='%s' 
	  WHERE TagId=%f  
	  RETURNING TagId;`,
		res.CustomerId,
		res.DeviceId,
		res.Epc,
		res.TagId)
}

//SelectWithDb
func (res TagType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(
		&res.CustomerId,
		&res.DeviceId,
		&res.Epc,
		&res.Active,
		&res.CreateTime)
	return errDb
}
