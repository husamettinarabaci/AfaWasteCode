package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

//RecyDeviceMainType
type RecyDeviceMainType struct {
	DeviceId     float64
	CustomerId   float64
	SerialNumber string
	Active       string
	CreateTime   string
}

//New
func (res *RecyDeviceMainType) New() {
	res.DeviceId = 0
	res.CustomerId = 1
	res.SerialNumber = ""
	res.Active = STATU_ACTIVE
	res.CreateTime = GetTime()
}

//ToId String
func (res *RecyDeviceMainType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToCustomerId String
func (res *RecyDeviceMainType) ToCustomerIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToByte
func (res *RecyDeviceMainType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res *RecyDeviceMainType) ToString() string {
	return string(res.ToByte())

}

//Byte To RecyDeviceMainType
func ByteToRecyDeviceMainType(retByte []byte) RecyDeviceMainType {
	var retVal RecyDeviceMainType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To RecyDeviceMainType
func StringToRecyDeviceMainType(retStr string) RecyDeviceMainType {
	return ByteToRecyDeviceMainType([]byte(retStr))
}

//SelectSQL
func (res *RecyDeviceMainType) SelectSQL() string {
	return fmt.Sprintf(`SELECT CustomerId,SerialNumber,Active,CreateTime
	 FROM public.recy_main_devices
	 WHERE DeviceId=%f AND Active=`+STATU_ACTIVE+` ;`, res.DeviceId)
}

//InsertSQL
func (res *RecyDeviceMainType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.recy_main_devices (CustomerId,SerialNumber) 
	  VALUES (%f,'%s') 
	  RETURNING DeviceId;`, res.CustomerId, res.SerialNumber)
}

//InsertDataSQL
func (res *RecyDeviceMainType) InsertDataSQL() string {
	return fmt.Sprintf(`INSERT INTO public.recy_main_devices (DeviceId,CustomerId,SerialNumber) 
	  VALUES (%f,%f,'%s') 
	  RETURNING DeviceId;`, res.DeviceId, res.CustomerId, res.SerialNumber)
}

//UpdateSQL
func (res *RecyDeviceMainType) UpdateSQL() string {
	return fmt.Sprintf(`UPDATE public.recy_main_devices 
	  SET CustomerId=%f 
	  WHERE DeviceId=%f  
	  RETURNING DeviceId;`,
		res.CustomerId,
		res.DeviceId)
}

//SelectWithDb
func (res *RecyDeviceMainType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(
		&res.CustomerId,
		&res.SerialNumber,
		&res.Active,
		&res.CreateTime)
	return errDb
}
