package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

//UltDeviceThermType
type UltDeviceThermType struct {
	DeviceId    float64
	Therm       string
	ThermTime   string
	ThermStatus string
}

//New
func (res *UltDeviceThermType) New() {
	res.DeviceId = 0
	res.Therm = "0"
	res.ThermTime = ""
	res.ThermStatus = THERM_STATU_NONE
}

//ToId String
func (res UltDeviceThermType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToByte
func (res UltDeviceThermType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res UltDeviceThermType) ToString() string {
	return string(res.ToByte())

}

//Byte To UltDeviceThermType
func ByteToUltDeviceThermType(retByte []byte) UltDeviceThermType {
	var retVal UltDeviceThermType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To UltDeviceThermType
func StringToUltDeviceThermType(retStr string) UltDeviceThermType {
	return ByteToUltDeviceThermType([]byte(retStr))
}

//SelectSQL
func (res UltDeviceThermType) SelectSQL() string {
	return fmt.Sprintf(`SELECT Therm,ThermTime,ThermStatus
	 FROM public.ult_therm_devices
	 WHERE DeviceId=%f ;`, res.DeviceId)
}

//InsertSQL
func (res UltDeviceThermType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.ult_therm_devices (DeviceId,Therm,ThermTime,ThermStatus) 
	  VALUES (%f,'%s','%s','%s') 
	  RETURNING DeviceId;`, res.DeviceId, res.Therm, res.ThermTime, res.ThermStatus)
}

//UpdateSQL
func (res UltDeviceThermType) UpdateSQL() string {
	return fmt.Sprintf(`UPDATE public.ult_therm_devices 
	  SET Therm='%s',ThermTime='%s',ThermStatus='%s' 
	  WHERE DeviceId=%f  
	  RETURNING DeviceId;`,
		res.Therm,
		res.ThermTime,
		res.ThermStatus,
		res.DeviceId)
}

//SelectWithDb
func (res UltDeviceThermType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(
		&res.Therm,
		&res.ThermTime,
		&res.ThermStatus)
	return errDb
}
