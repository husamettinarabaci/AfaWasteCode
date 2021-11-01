package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

//UltDeviceType
type UltDeviceType struct {
	DeviceId     float64
	CustomerId   float64
	SerialNumber string
	Active       string
	CreateTime   string
	OldLatitude  float64
	OldLongitude float64
}

//New
func (res *UltDeviceType) New() {
	res.DeviceId = 0
	res.CustomerId = 0
	res.SerialNumber = ""
	res.Active = STATU_ACTIVE
	res.CreateTime = ""
	res.OldLatitude = 0
	res.OldLongitude = 0
}

//ToId String
func (res UltDeviceType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToCustomerId String
func (res UltDeviceType) ToCustomerIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToByte
func (res UltDeviceType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res UltDeviceType) ToString() string {
	return string(res.ToByte())

}

//Byte To UltDeviceType
func ByteToUltDeviceType(retByte []byte) UltDeviceType {
	var retVal UltDeviceType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To UltDeviceType
func StringToUltDeviceType(retStr string) UltDeviceType {
	return ByteToUltDeviceType([]byte(retStr))
}

//SelectSQL
func (res UltDeviceType) SelectSQL() string {
	return fmt.Sprintf(`SELECT CustomerId,SerialNumber,Active,CreateTime,OldLatitude,OldLongitude
	 FROM public.ult_devices
	 WHERE DeviceId=%f ;`, res.DeviceId)
}

//InsertSQL
func (res UltDeviceType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.ult_devices (DeviceId,CustomerId,SerialNumber,OldLatitude,OldLongitude) 
	  VALUES (%f,%f,'%s',%f,%f) 
	  RETURNING DeviceId;`, res.DeviceId, res.CustomerId, res.SerialNumber, res.OldLatitude, res.OldLongitude)
}

//UpdateSQL
func (res UltDeviceType) UpdateSQL() string {
	return fmt.Sprintf(`UPDATE public.ult_devices 
	  SET CustomerId=%f,SerialNumber='%s',OldLatitude=%f,OldLongitude=%f
	  WHERE DeviceId=%f  
	  RETURNING DeviceId;`,
		res.CustomerId,
		res.SerialNumber,
		res.OldLatitude,
		res.OldLongitude,
		res.DeviceId)
}

//SelectWithDb
func (res UltDeviceType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(
		&res.CustomerId,
		&res.SerialNumber,
		&res.Active,
		&res.CreateTime,
		&res.OldLatitude,
		&res.OldLongitude)
	return errDb
}
