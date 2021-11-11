package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

//RfidDeviceWorkHourType
type RfidDeviceWorkHourType struct {
	DeviceId       float64
	WorkHour1Start string
	WorkHour1Add   float64
	WorkHour2Start string
	WorkHour2Add   float64
	WorkHour3Start string
	WorkHour3Add   float64
	NewData        bool
}

//New
func (res *RfidDeviceWorkHourType) New() {
	res.DeviceId = 0
	res.WorkHour1Start = GetTime()
	res.WorkHour1Add = 0
	res.WorkHour2Start = GetTime()
	res.WorkHour2Add = 0
	res.WorkHour3Start = GetTime()
	res.WorkHour3Add = 0
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
	 WorkHour1Start,WorkHour1Add,
	 WorkHour2Start,WorkHour2Add,
	 WorkHour3Start,WorkHour3Add
	 FROM public.rfid_workhour_devices
	 WHERE DeviceId=%f ;`, res.DeviceId)
}

//InsertSQL
func (res *RfidDeviceWorkHourType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.rfid_workhour_devices (DeviceId,
	 WorkHour1Start,WorkHour1Add,
	 WorkHour2Start,WorkHour2Add,
	 WorkHour3Start,WorkHour3Add) 
	  VALUES (%f
	,'%s',%f,'%s',%f,'%s',%f) 
	  RETURNING DeviceId;`, res.DeviceId,
		res.WorkHour1Start, res.WorkHour1Add,
		res.WorkHour2Start, res.WorkHour2Add,
		res.WorkHour3Start, res.WorkHour3Add)
}

//UpdateSQL
func (res *RfidDeviceWorkHourType) UpdateSQL() string {

	return fmt.Sprintf(`UPDATE public.rfid_workhour_devices 
	  SET WorkHour1Start='%s',WorkHour1Add=%f,
	  WorkHour2Start='%s',WorkHour2Add=%f,
	  WorkHour3Start='%s',WorkHour3Add=%f
	  WHERE DeviceId=%f  
	  RETURNING DeviceId;`,
		res.WorkHour1Start, res.WorkHour1Add,
		res.WorkHour2Start, res.WorkHour2Add,
		res.WorkHour3Start, res.WorkHour3Add,
		res.DeviceId)
}

//SelectWithDb
func (res *RfidDeviceWorkHourType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(
		&res.WorkHour1Start, &res.WorkHour1Add,
		&res.WorkHour2Start, &res.WorkHour2Add,
		&res.WorkHour3Start, &res.WorkHour3Add)
	return errDb
}
