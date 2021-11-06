package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

//UltDeviceSensType
type UltDeviceSensType struct {
	DeviceId  float64
	UltTime   string
	UltRange  float64
	UltStatus string
	NewData   bool
}

//New
func (res *UltDeviceSensType) New() {
	res.DeviceId = 0
	res.UltTime = GetTime()
	res.UltRange = 0
	res.UltStatus = CONTINER_FULLNESS_STATU_NONE
	res.NewData = false
}

//ToId String
func (res *UltDeviceSensType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToByte
func (res *UltDeviceSensType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res *UltDeviceSensType) ToString() string {
	return string(res.ToByte())

}

//Byte To UltDeviceSensType
func ByteToUltDeviceSensType(retByte []byte) UltDeviceSensType {
	var retVal UltDeviceSensType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To UltDeviceSensType
func StringToUltDeviceSensType(retStr string) UltDeviceSensType {
	return ByteToUltDeviceSensType([]byte(retStr))
}

//SelectSQL
func (res *UltDeviceSensType) SelectSQL() string {
	return fmt.Sprintf(`SELECT UltTime,UltRange,UltStatus
	 FROM public.ult_sens_devices
	 WHERE DeviceId=%f ;`, res.DeviceId)
}

//InsertSQL
func (res *UltDeviceSensType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.ult_sens_devices (DeviceId,UltTime,UltRange,UltStatus) 
	  VALUES (%f,'%s',%f,'%s') 
	  RETURNING DeviceId;`, res.DeviceId,
		res.UltTime, res.UltRange, res.UltStatus)
}

//UpdateSQL
func (res *UltDeviceSensType) UpdateSQL() string {
	return fmt.Sprintf(`UPDATE public.ult_sens_devices 
	  SET UltTime='%s',UltRange=%f,UltStatus='%s'
	  WHERE DeviceId=%f  
	  RETURNING DeviceId;`,
		res.UltTime, res.UltRange, res.UltStatus, res.DeviceId)
}

//SelectWithDb
func (res *UltDeviceSensType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(
		&res.UltTime,
		&res.UltRange,
		&res.UltStatus)
	return errDb
}
