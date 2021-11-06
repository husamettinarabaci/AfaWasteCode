package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

//TagStatuType
type TagStatuType struct {
	TagId          float64
	ContainerStatu string
	TagStatu       string
	ImageStatu     string
	CheckTime      string
	NewData        bool
}

//New
func (res *TagStatuType) New() {
	res.TagId = 0
	res.TagStatu = TAG_STATU_NONE
	res.ImageStatu = STATU_PASSIVE
	res.CheckTime = GetTime()
	res.NewData = false
}

//ToId String
func (res *TagStatuType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.TagId)
}

//ToByte
func (res *TagStatuType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//ToString Get JSON
func (res *TagStatuType) ToString() string {
	return string(res.ToByte())

}

//Byte To TagStatuType
func ByteToTagStatuType(retByte []byte) TagStatuType {
	var retVal TagStatuType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To TagStatuType
func StringToTagStatuType(retStr string) TagStatuType {
	return ByteToTagStatuType([]byte(retStr))
}

//SelectSQL
func (res *TagStatuType) SelectSQL() string {
	return fmt.Sprintf(`SELECT TagStatu,ImageStatu,CheckTime
	 FROM public.tag_status
	 WHERE TagId=%f ;`, res.TagId)
}

//InsertSQL
func (res *TagStatuType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.tag_status (TagId,TagStatu,ImageStatu,CheckTime) 
	  VALUES (%f,'%s','%s','%s') 
	  RETURNING TagId;`, res.TagId, res.TagStatu, res.ImageStatu, res.CheckTime)
}

//UpdateSQL
func (res *TagStatuType) UpdateSQL() string {
	return fmt.Sprintf(`UPDATE public.tag_status 
	  SET TagStatu='%s',ImageStatu='%s',CheckTime='%s'
	  WHERE TagId=%f  
	  RETURNING TagId;`,
		res.TagStatu,
		res.ImageStatu,
		res.CheckTime,
		res.TagId)
}

//SelectWithDb
func (res *TagStatuType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(
		&res.TagStatu,
		&res.ImageStatu,
		&res.CheckTime)
	return errDb
}
