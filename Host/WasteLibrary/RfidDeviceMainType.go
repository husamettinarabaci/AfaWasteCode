package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

//RfidDeviceMainType
type RfidDeviceMainType struct {
	DeviceId     float64
	CustomerId   float64
	SerialNumber string
	Active       string
	CreateTime   string
}

//New
func (res *RfidDeviceMainType) New() {
	res.DeviceId = 0
	res.CustomerId = 1
	res.SerialNumber = ""
	res.Active = STATU_ACTIVE
	res.CreateTime = GetTime()
}

//ToId String
func (res *RfidDeviceMainType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToCustomerId String
func (res *RfidDeviceMainType) ToCustomerIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToByte
func (res *RfidDeviceMainType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res *RfidDeviceMainType) ToString() string {
	return string(res.ToByte())

}

//Byte To RfidDeviceMainType
func ByteToRfidDeviceMainType(retByte []byte) RfidDeviceMainType {
	var retVal RfidDeviceMainType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To RfidDeviceMainType
func StringToRfidDeviceMainType(retStr string) RfidDeviceMainType {
	return ByteToRfidDeviceMainType([]byte(retStr))
}

//SelectSQL
func (res *RfidDeviceMainType) SelectSQL() string {
	return fmt.Sprintf(`SELECT CustomerId,SerialNumber,Active,CreateTime
	 FROM public.rfid_main_devices
	 WHERE DeviceId=%f ;`, res.DeviceId)
}

//InsertSQL
func (res *RfidDeviceMainType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.rfid_main_devices (CustomerId,SerialNumber) 
	  VALUES (%f,'%s') 
	  RETURNING DeviceId;`, res.CustomerId, res.SerialNumber)
}

//InsertDataSQL
func (res *RfidDeviceMainType) InsertDataSQL() string {
	return fmt.Sprintf(`INSERT INTO public.rfid_main_devices (DeviceId,CustomerId,SerialNumber) 
	  VALUES (%f,%f,'%s') 
	  RETURNING DeviceId;`, res.DeviceId, res.CustomerId, res.SerialNumber)
}

//UpdateSQL
func (res *RfidDeviceMainType) UpdateSQL() string {
	return fmt.Sprintf(`UPDATE public.rfid_main_devices 
	  SET CustomerId=%f
	  WHERE DeviceId=%f  
	  RETURNING DeviceId;`,
		res.CustomerId,
		res.DeviceId)
}

//SelectWithDb
func (res *RfidDeviceMainType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(
		&res.CustomerId,
		&res.SerialNumber,
		&res.Active,
		&res.CreateTime)
	return errDb
}
