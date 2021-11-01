package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

//RecyDeviceGpsType
type RecyDeviceGpsType struct {
	DeviceId  float64
	Latitude  float64
	Longitude float64
	GpsTime   string
}

//New
func (res *RecyDeviceGpsType) New() {
	res.DeviceId = 0
	res.Latitude = 0
	res.Longitude = 0
	res.GpsTime = ""
}

//ToId String
func (res RecyDeviceGpsType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToByte
func (res RecyDeviceGpsType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res RecyDeviceGpsType) ToString() string {
	return string(res.ToByte())

}

//Byte To RecyDeviceGpsType
func ByteToRecyDeviceGpsType(retByte []byte) RecyDeviceGpsType {
	var retVal RecyDeviceGpsType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To RecyDeviceGpsType
func StringToRecyDeviceGpsType(retStr string) RecyDeviceGpsType {
	return ByteToRecyDeviceGpsType([]byte(retStr))
}

//SelectSQL
func (res RecyDeviceGpsType) SelectSQL() string {
	return fmt.Sprintf(`SELECT Latitude,Longitude,GpsTime
	 FROM public.recy_gps_devices
	 WHERE DeviceId=%f ;`, res.DeviceId)
}

//InsertSQL
func (res RecyDeviceGpsType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.recy_gps_devices (DeviceId,Latitude,Longitude,GpsTime) 
	  VALUES (%f,%f,%f,'%s') 
	  RETURNING DeviceId;`, res.DeviceId, res.Latitude, res.Longitude, res.GpsTime)
}

//UpdateSQL
func (res RecyDeviceGpsType) UpdateSQL() string {
	return fmt.Sprintf(`UPDATE public.recy_gps_devices 
	  SET Latitude=%f,Longitude=%f,GpsTime='%s' 
	  WHERE DeviceId=%f  
	  RETURNING DeviceId;`,
		res.Latitude,
		res.Longitude,
		res.GpsTime,
		res.DeviceId)
}

//SelectWithDb
func (res RecyDeviceGpsType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(
		&res.Latitude,
		&res.Longitude,
		&res.GpsTime)
	return errDb
}
