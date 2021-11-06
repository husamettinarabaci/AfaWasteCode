package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

//UltDeviceVersionType
type UltDeviceVersionType struct {
	DeviceId        float64
	FirmwareVersion string
	NewData         bool
}

//New
func (res *UltDeviceVersionType) New() {
	res.DeviceId = 0
	res.FirmwareVersion = "1"
	res.NewData = false

}

//ToId String
func (res *UltDeviceVersionType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToByte
func (res *UltDeviceVersionType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res *UltDeviceVersionType) ToString() string {
	return string(res.ToByte())

}

//Byte To UltDeviceVersionType
func ByteToUltDeviceVersionType(retByte []byte) UltDeviceVersionType {
	var retVal UltDeviceVersionType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To UltDeviceVersionType
func StringToUltDeviceVersionType(retStr string) UltDeviceVersionType {
	return ByteToUltDeviceVersionType([]byte(retStr))
}

//SelectSQL
func (res *UltDeviceVersionType) SelectSQL() string {
	return fmt.Sprintf(`SELECT FirmwareVersion
	 FROM public.ult_version_devices
	 WHERE DeviceId=%f ;`, res.DeviceId)
}

//InsertSQL
func (res *UltDeviceVersionType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.ult_version_devices (DeviceId,FirmwareVersion) 
	  VALUES (%f,'%s') 
	  RETURNING DeviceId;`, res.DeviceId, res.FirmwareVersion)
}

//UpdateSQL
func (res *UltDeviceVersionType) UpdateSQL() string {
	return fmt.Sprintf(`UPDATE public.ult_version_devices 
	  SET FirmwareVersion='%s'
	  WHERE DeviceId=%f  
	  RETURNING DeviceId;`,
		res.FirmwareVersion,
		res.DeviceId)
}

//SelectWithDb
func (res *UltDeviceVersionType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(
		&res.FirmwareVersion)
	return errDb
}
