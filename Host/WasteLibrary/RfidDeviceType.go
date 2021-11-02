package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

//RfidDeviceType
type RfidDeviceType struct {
	DeviceId      float64
	CustomerId    float64
	SerialNumber  string
	DeviceBase    RfidDeviceBaseType
	DeviceStatu   RfidDeviceStatuType
	DeviceGps     RfidDeviceGpsType
	DeviceAlarm   RfidDeviceAlarmType
	DeviceTherm   RfidDeviceThermType
	DeviceVersion RfidDeviceVersionType
	DeviceDetail  RfidDeviceDetailType
	Active        string
	CreateTime    string
}

//New
func (res *RfidDeviceType) New() {
	res.DeviceId = 0
	res.CustomerId = 1
	res.Active = STATU_ACTIVE
	res.CreateTime = ""
}

//ToId String
func (res RfidDeviceType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToCustomerId String
func (res RfidDeviceType) ToCustomerIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToByte
func (res RfidDeviceType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res RfidDeviceType) ToString() string {
	return string(res.ToByte())

}

//Byte To RfidDeviceType
func ByteToRfidDeviceType(retByte []byte) RfidDeviceType {
	var retVal RfidDeviceType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To RfidDeviceType
func StringToRfidDeviceType(retStr string) RfidDeviceType {
	return ByteToRfidDeviceType([]byte(retStr))
}

//SelectSQL
func (res RfidDeviceType) SelectSQL() string {
	return fmt.Sprintf(`SELECT CustomerId,SerialNumber,Active,CreateTime
	 FROM public.rfid_devices
	 WHERE DeviceId=%f ;`, res.DeviceId)
}

//InsertSQL
func (res RfidDeviceType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.rfid_devices (CustomerId,SerialNumber) 
	  VALUES (%f,'%s') 
	  RETURNING DeviceId;`, res.CustomerId, res.SerialNumber)
}

//InsertDataSQL
func (res RfidDeviceType) InsertDataSQL() string {
	return fmt.Sprintf(`INSERT INTO public.rfid_devices (DeviceId,CustomerId,SerialNumber) 
	  VALUES (%f,%f,'%s') 
	  RETURNING DeviceId;`, res.DeviceId, res.CustomerId, res.SerialNumber)
}

//UpdateSQL
func (res RfidDeviceType) UpdateSQL() string {
	return fmt.Sprintf(`UPDATE public.rfid_devices 
	  SET CustomerId=%f,SerialNumber='%s' 
	  WHERE DeviceId=%f  
	  RETURNING DeviceId;`,
		res.CustomerId,
		res.SerialNumber,
		res.DeviceId)
}

//SelectWithDb
func (res RfidDeviceType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(
		&res.CustomerId,
		&res.SerialNumber,
		&res.Active,
		&res.CreateTime)
	return errDb
}
