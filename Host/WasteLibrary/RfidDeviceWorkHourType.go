package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

//RfidDeviceWorkHourType
type RfidDeviceWorkHourType struct {
	DeviceId        float64
	WorkStartHour   int
	WorkStartMinute int
	WorkEndHour     int
	WorkEndMinute   int
	NewData         bool
}

//New
func (res *RfidDeviceWorkHourType) New() {
	res.DeviceId = 0
	res.WorkStartHour = 06
	res.WorkStartMinute = 0
	res.WorkEndHour = 18
	res.WorkEndMinute = 30
	res.NewData = false
}

//ToId String
func (res *RfidDeviceWorkHourType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToByte
func (res *RfidDeviceWorkHourType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res *RfidDeviceWorkHourType) ToString() string {
	return string(res.ToByte())

}

//Byte To RfidDeviceWorkHourType
func ByteToRfidDeviceWorkHourType(retByte []byte) RfidDeviceWorkHourType {
	var retVal RfidDeviceWorkHourType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To RfidDeviceWorkHourType
func StringToRfidDeviceWorkHourType(retStr string) RfidDeviceWorkHourType {
	return ByteToRfidDeviceWorkHourType([]byte(retStr))
}

//SelectSQL
func (res *RfidDeviceWorkHourType) SelectSQL() string {
	return fmt.Sprintf(`SELECT 
	WorkStartHour,WorkStartMinute,
	WorkEndHour,WorkEndMinute
	 FROM public.rfid_workhour_devices
	 WHERE DeviceId=%f ;`, res.DeviceId)
}

//InsertSQL
func (res *RfidDeviceWorkHourType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.rfid_workhour_devices (DeviceId,
		WorkStartHour,WorkStartMinute,
		WorkEndHour,WorkEndMinute) 
	  VALUES (%f,%d,%d,%d,%d) 
	  RETURNING DeviceId;`, res.DeviceId,
		res.WorkStartHour, res.WorkStartMinute,
		res.WorkEndHour, res.WorkEndMinute)
}

//UpdateSQL
func (res *RfidDeviceWorkHourType) UpdateSQL() string {

	return fmt.Sprintf(`UPDATE public.rfid_workhour_devices 
	  SET WorkStartHour=%d,WorkStartMinute=%d,
	  WorkEndHour=%d,WorkEndMinute=%d
	  WHERE DeviceId=%f  
	  RETURNING DeviceId;`,
		res.WorkStartHour, res.WorkStartMinute,
		res.WorkEndHour, res.WorkEndMinute,
		res.DeviceId)
}

//SelectWithDb
func (res *RfidDeviceWorkHourType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(
		&res.WorkStartHour, &res.WorkStartMinute,
		&res.WorkEndHour, &res.WorkEndMinute)
	return errDb
}
