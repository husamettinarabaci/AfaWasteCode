package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

//RfidDeviceDetailType
type RfidDeviceDetailType struct {
	DeviceId      float64
	PlateNo       string
	DriverName    string
	DriverSurName string
	NewData       bool
}

//New
func (res *RfidDeviceDetailType) New() {
	res.DeviceId = 0
	res.PlateNo = ""
	res.DriverName = ""
	res.DriverSurName = ""
	res.NewData = false
}

//ToId String
func (res RfidDeviceDetailType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToByte
func (res RfidDeviceDetailType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res RfidDeviceDetailType) ToString() string {
	return string(res.ToByte())

}

//Byte To RfidDeviceDetailType
func ByteToRfidDeviceDetailType(retByte []byte) RfidDeviceDetailType {
	var retVal RfidDeviceDetailType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To RfidDeviceDetailType
func StringToRfidDeviceDetailType(retStr string) RfidDeviceDetailType {
	return ByteToRfidDeviceDetailType([]byte(retStr))
}

//SelectSQL
func (res RfidDeviceDetailType) SelectSQL() string {
	return fmt.Sprintf(`SELECT PlateNo,DriverName,DriverSurName
	 FROM public.rfid_detail_devices
	 WHERE DeviceId=%f ;`, res.DeviceId)
}

//InsertSQL
func (res RfidDeviceDetailType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.rfid_detail_devices (DeviceId,PlateNo,DriverName,DriverSurName) 
	  VALUES (%f,'%s','%s','%s') 
	  RETURNING DeviceId;`, res.DeviceId, res.PlateNo, res.DriverName, res.DriverSurName)
}

//UpdateSQL
func (res RfidDeviceDetailType) UpdateSQL() string {
	return fmt.Sprintf(`UPDATE public.rfid_detail_devices 
	  SET PlateNo='%s',DriverName='%s',DriverSurName='%s'
	  WHERE DeviceId=%f  
	  RETURNING DeviceId;`,
		res.PlateNo,
		res.DriverName,
		res.DriverSurName,
		res.DeviceId)
}

//SelectWithDb
func (res RfidDeviceDetailType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(
		&res.PlateNo,
		&res.DriverName,
		&res.DriverSurName)
	return errDb
}
