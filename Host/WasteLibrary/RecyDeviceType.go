package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

//RecyDeviceType
type RecyDeviceType struct {
	DeviceId      float64
	CustomerId    float64
	SerialNumber  string
	DeviceBase    RecyDeviceBaseType
	DeviceGps     RecyDeviceGpsType
	DeviceTherm   RecyDeviceThermType
	DeviceVersion RecyDeviceVersionType
	DeviceAlarm   RecyDeviceAlarmType
	DeviceStatu   RecyDeviceStatuType
	DeviceDetail  RecyDeviceDetailType
	Active        string
	CreateTime    string
}

//New
func (res *RecyDeviceType) New() {
	res.DeviceId = 0
	res.CustomerId = 0
	res.DeviceBase.New()
	res.DeviceGps.New()
	res.DeviceTherm.New()
	res.DeviceVersion.New()
	res.DeviceAlarm.New()
	res.DeviceStatu.New()
	res.DeviceDetail.New()
}

//ToId String
func (res RecyDeviceType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToCustomerId String
func (res RecyDeviceType) ToCustomerIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToByte
func (res RecyDeviceType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res RecyDeviceType) ToString() string {
	return string(res.ToByte())

}

//Byte To RecyDeviceType
func ByteToRecyDeviceType(retByte []byte) RecyDeviceType {
	var retVal RecyDeviceType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To RecyDeviceType
func StringToRecyDeviceType(retStr string) RecyDeviceType {
	return ByteToRecyDeviceType([]byte(retStr))
}

//SelectSQL
func (res RecyDeviceType) SelectSQL() string {
	return fmt.Sprintf(`SELECT CustomerId,SerialNumber,Active,CreateTime
	 FROM public.recy_devices
	 WHERE DeviceId=%f ;`, res.DeviceId)
}

//InsertSQL
func (res RecyDeviceType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.rfid_devices (CustomerId,SerialNumber) 
	  VALUES (%f,'%s') 
	  RETURNING DeviceId;`, res.CustomerId, res.SerialNumber)
}

//InsertDataSQL
func (res RecyDeviceType) InsertDataSQL() string {
	return fmt.Sprintf(`INSERT INTO public.rfid_devices (DeviceId,CustomerId,SerialNumber) 
	  VALUES (%f,%f,'%s') 
	  RETURNING DeviceId;`, res.DeviceId, res.CustomerId, res.SerialNumber)
}

//UpdateSQL
func (res RecyDeviceType) UpdateSQL() string {
	return fmt.Sprintf(`UPDATE public.recy_devices 
	  SET CustomerId=%f,SerialNumber='%s' 
	  WHERE DeviceId=%f  
	  RETURNING DeviceId;`,
		res.CustomerId,
		res.SerialNumber,
		res.DeviceId)
}

//SelectWithDb
func (res RecyDeviceType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(
		&res.CustomerId,
		&res.SerialNumber,
		&res.Active,
		&res.CreateTime)
	return errDb
}
