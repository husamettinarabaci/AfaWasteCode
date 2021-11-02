package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

//RecyDeviceThermType
type RecyDeviceThermType struct {
	DeviceId    float64
	Therm       string
	ThermTime   string
	ThermStatus string
	NewData     bool
}

//New
func (res *RecyDeviceThermType) New() {
	res.DeviceId = 0
	res.Therm = "00"
	res.ThermTime = ""
	res.ThermStatus = THERMSTATU_NONE
	res.NewData = false
}

//ToId String
func (res RecyDeviceThermType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToByte
func (res RecyDeviceThermType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res RecyDeviceThermType) ToString() string {
	return string(res.ToByte())

}

//Byte To RecyDeviceThermType
func ByteToRecyDeviceThermType(retByte []byte) RecyDeviceThermType {
	var retVal RecyDeviceThermType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To RecyDeviceThermType
func StringToRecyDeviceThermType(retStr string) RecyDeviceThermType {
	return ByteToRecyDeviceThermType([]byte(retStr))
}

//SelectSQL
func (res RecyDeviceThermType) SelectSQL() string {
	return fmt.Sprintf(`SELECT Therm,ThermTime,ThermStatus
	 FROM public.recy_therm_devices
	 WHERE DeviceId=%f ;`, res.DeviceId)
}

//InsertSQL
func (res RecyDeviceThermType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.recy_therm_devices (DeviceId,Therm,ThermTime,ThermStatus) 
	  VALUES (%f,'%s','%s','%s') 
	  RETURNING DeviceId;`, res.DeviceId, res.Therm, res.ThermTime, res.ThermStatus)
}

//UpdateSQL
func (res RecyDeviceThermType) UpdateSQL() string {
	return fmt.Sprintf(`UPDATE public.recy_therm_devices 
	  SET Therm='%s',ThermTime='%s',ThermStatus='%s' 
	  WHERE DeviceId=%f  
	  RETURNING DeviceId;`,
		res.Therm,
		res.ThermTime,
		res.ThermStatus,
		res.DeviceId)
}

//SelectWithDb
func (res RecyDeviceThermType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(
		&res.Therm,
		&res.ThermTime,
		&res.ThermStatus)
	return errDb
}
