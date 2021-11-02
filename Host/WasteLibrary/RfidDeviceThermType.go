package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

//RfidDeviceThermType
type RfidDeviceThermType struct {
	DeviceId    float64
	Therm       string
	ThermTime   string
	ThermStatus string
	NewData     bool
}

//New
func (res *RfidDeviceThermType) New() {
	res.DeviceId = 0
	res.Therm = "00"
	res.ThermTime = ""
	res.ThermStatus = THERMSTATU_NONE
	res.NewData = false
}

//ToId String
func (res RfidDeviceThermType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToByte
func (res RfidDeviceThermType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res RfidDeviceThermType) ToString() string {
	return string(res.ToByte())

}

//Byte To RfidDeviceThermType
func ByteToRfidDeviceThermType(retByte []byte) RfidDeviceThermType {
	var retVal RfidDeviceThermType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To RfidDeviceThermType
func StringToRfidDeviceThermType(retStr string) RfidDeviceThermType {
	return ByteToRfidDeviceThermType([]byte(retStr))
}

//SelectSQL
func (res RfidDeviceThermType) SelectSQL() string {
	return fmt.Sprintf(`SELECT Therm,ThermTime,ThermStatus
	 FROM public.rfid_therm_devices
	 WHERE DeviceId=%f ;`, res.DeviceId)
}

//InsertSQL
func (res RfidDeviceThermType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.rfid_therm_devices (DeviceId,Therm,ThermTime,ThermStatus) 
	  VALUES (%f,'%s','%s','%s') 
	  RETURNING DeviceId;`, res.DeviceId, res.Therm, res.ThermTime, res.ThermStatus)
}

//UpdateSQL
func (res RfidDeviceThermType) UpdateSQL() string {
	return fmt.Sprintf(`UPDATE public.rfid_therm_devices 
	  SET Therm='%s',ThermTime='%s',ThermStatus='%s' 
	  WHERE DeviceId=%f  
	  RETURNING DeviceId;`,
		res.Therm,
		res.ThermTime,
		res.ThermStatus,
		res.DeviceId)
}

//SelectWithDb
func (res RfidDeviceThermType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(
		&res.Therm,
		&res.ThermTime,
		&res.ThermStatus)
	return errDb
}
