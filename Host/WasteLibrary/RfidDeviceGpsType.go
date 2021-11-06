package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

//RfidDeviceGpsType
type RfidDeviceGpsType struct {
	DeviceId  float64
	Latitude  float64
	Longitude float64
	Speed     float64
	GpsTime   string
	NewData   bool
}

//New
func (res *RfidDeviceGpsType) New() {
	res.DeviceId = 0
	res.Latitude = 0
	res.Longitude = 0
	res.Speed = -1
	res.GpsTime = GetTime()
	res.NewData = false
}

//ToId String
func (res *RfidDeviceGpsType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToByte
func (res *RfidDeviceGpsType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res *RfidDeviceGpsType) ToString() string {
	return string(res.ToByte())

}

//Byte To RfidDeviceGpsType
func ByteToRfidDeviceGpsType(retByte []byte) RfidDeviceGpsType {
	var retVal RfidDeviceGpsType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To RfidDeviceGpsType
func StringToRfidDeviceGpsType(retStr string) RfidDeviceGpsType {
	return ByteToRfidDeviceGpsType([]byte(retStr))
}

//SelectSQL
func (res *RfidDeviceGpsType) SelectSQL() string {
	return fmt.Sprintf(`SELECT Latitude,Longitude,Speed,GpsTime
	 FROM public.rfid_gps_devices
	 WHERE DeviceId=%f ;`, res.DeviceId)
}

//InsertSQL
func (res *RfidDeviceGpsType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.rfid_gps_devices (DeviceId,Latitude,Longitude,Speed,GpsTime) 
	  VALUES (%f,%f,%f,%f,'%s') 
	  RETURNING DeviceId;`, res.DeviceId, res.Latitude, res.Longitude, res.Speed, res.GpsTime)
}

//UpdateSQL
func (res *RfidDeviceGpsType) UpdateSQL() string {
	return fmt.Sprintf(`UPDATE public.rfid_gps_devices 
	  SET Latitude=%f,Longitude=%f,Speed=%f,GpsTime='%s' 
	  WHERE DeviceId=%f  
	  RETURNING DeviceId;`,
		res.Latitude,
		res.Longitude,
		res.Speed,
		res.GpsTime,
		res.DeviceId)
}

//SelectWithDb
func (res *RfidDeviceGpsType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(
		&res.Latitude,
		&res.Longitude,
		&res.Speed,
		&res.GpsTime)
	return errDb
}
