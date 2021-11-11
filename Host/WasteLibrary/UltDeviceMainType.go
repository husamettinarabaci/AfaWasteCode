package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

//UltDeviceMainType
type UltDeviceMainType struct {
	DeviceId     float64
	CustomerId   float64
	SerialNumber string
	OldLatitude  float64
	OldLongitude float64
	Active       string
	CreateTime   string
}

//New
func (res *UltDeviceMainType) New() {
	res.DeviceId = 0
	res.CustomerId = 1
	res.SerialNumber = ""
	res.OldLatitude = 0
	res.OldLongitude = 0
	res.Active = STATU_ACTIVE
	res.CreateTime = GetTime()

}

//ToId String
func (res *UltDeviceMainType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToCustomerId String
func (res *UltDeviceMainType) ToCustomerIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToByte
func (res *UltDeviceMainType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res *UltDeviceMainType) ToString() string {
	return string(res.ToByte())

}

//Byte To UltDeviceMainType
func ByteToUltDeviceMainType(retByte []byte) UltDeviceMainType {
	var retVal UltDeviceMainType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To UltDeviceMainType
func StringToUltDeviceMainType(retStr string) UltDeviceMainType {
	return ByteToUltDeviceMainType([]byte(retStr))
}

//SelectSQL
func (res *UltDeviceMainType) SelectSQL() string {
	return fmt.Sprintf(`SELECT CustomerId,SerialNumber,Active,CreateTime,OldLatitude,OldLongitude
	 FROM public.ult_main_devices
	 WHERE DeviceId=%f  ;`, res.DeviceId)
}

//InsertSQL
func (res *UltDeviceMainType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.ult_main_devices (CustomerId,SerialNumber,OldLatitude,OldLongitude) 
	  VALUES (%f,'%s',%f,%f) 
	  RETURNING DeviceId;`, res.CustomerId, res.SerialNumber, res.OldLatitude, res.OldLongitude)
}

//InsertDataSQL
func (res *UltDeviceMainType) InsertDataSQL() string {
	return fmt.Sprintf(`INSERT INTO public.ult_main_devices (DeviceId,CustomerId,SerialNumber,OldLatitude,OldLongitude) 
	  VALUES (%f,%f,'%s',%f,%f) 
	  RETURNING DeviceId;`, res.DeviceId, res.CustomerId, res.SerialNumber, res.OldLatitude, res.OldLongitude)
}

//UpdateSQL
func (res *UltDeviceMainType) UpdateSQL() string {
	return fmt.Sprintf(`UPDATE public.ult_main_devices 
	  SET CustomerId=%f,OldLatitude=%f,OldLongitude=%f
	  WHERE DeviceId=%f  
	  RETURNING DeviceId;`,
		res.CustomerId,
		res.OldLatitude,
		res.OldLongitude,
		res.DeviceId)
}

//SelectWithDb
func (res *UltDeviceMainType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(
		&res.CustomerId,
		&res.SerialNumber,
		&res.Active,
		&res.CreateTime,
		&res.OldLatitude,
		&res.OldLongitude)
	return errDb
}
