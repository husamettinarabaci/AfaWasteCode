package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

//TagGpsType
type TagGpsType struct {
	TagId     float64
	Latitude  float64
	Longitude float64
	GpsTime   string
	NewData   bool
}

//New
func (res *TagGpsType) New() {
	res.TagId = 0
	res.Latitude = 0
	res.Longitude = 0
	res.GpsTime = ""
	res.NewData = false
}

//ToId String
func (res TagGpsType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.TagId)
}

//ToByte
func (res TagGpsType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//ToString Get JSON
func (res TagGpsType) ToString() string {
	return string(res.ToByte())

}

//Byte To TagGpsType
func ByteToTagGpsType(retByte []byte) TagGpsType {
	var retVal TagGpsType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To TagGpsType
func StringToTagGpsType(retStr string) TagGpsType {
	return ByteToTagGpsType([]byte(retStr))
}

//SelectSQL
func (res TagGpsType) SelectSQL() string {
	return fmt.Sprintf(`SELECT Latitude,Longitude,GpsTime
	 FROM public.tag_gpses
	 WHERE TagId=%f ;`, res.TagId)
}

//InsertSQL
func (res TagGpsType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.tag_gpses (TagId,Latitude,Longitude,GpsTime) 
	  VALUES (%f,%f,%f,'%s') 
	  RETURNING TagId;`, res.TagId, res.Latitude, res.Longitude, res.GpsTime)
}

//UpdateSQL
func (res TagGpsType) UpdateSQL() string {
	return fmt.Sprintf(`UPDATE public.tag_gpses 
	  SET Latitude=%f,Longitude=%f,GpsTime='%s' 
	  WHERE TagId=%f  
	  RETURNING TagId;`,
		res.Latitude,
		res.Longitude,
		res.GpsTime,
		res.TagId)
}

//SelectWithDb
func (res TagGpsType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(
		&res.Latitude,
		&res.Longitude,
		&res.GpsTime)
	return errDb
}
