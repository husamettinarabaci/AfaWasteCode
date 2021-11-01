package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

//UltDeviceStatuType
type UltDeviceStatuType struct {
	DeviceId        float64
	StatusTime      string
	AliveStatus     string
	AliveLastOkTime string
}

//New
func (res *UltDeviceStatuType) New() {
	res.DeviceId = 0
	res.StatusTime = ""
	res.AliveStatus = STATU_PASSIVE
	res.AliveLastOkTime = ""
}

//ToId String
func (res UltDeviceStatuType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToByte
func (res UltDeviceStatuType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res UltDeviceStatuType) ToString() string {
	return string(res.ToByte())

}

//Byte To UltDeviceStatuType
func ByteToUltDeviceStatuType(retByte []byte) UltDeviceStatuType {
	var retVal UltDeviceStatuType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To UltDeviceStatuType
func StringToUltDeviceStatuType(retStr string) UltDeviceStatuType {
	return ByteToUltDeviceStatuType([]byte(retStr))
}

//SelectSQL
func (res UltDeviceStatuType) SelectSQL() string {
	return fmt.Sprintf(`SELECT StatusTime,AliveStatus,AliveLastOkTime
	 FROM public.ult_statu_devices
	 WHERE DeviceId=%f ;`, res.DeviceId)
}

//InsertSQL
func (res UltDeviceStatuType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.ult_statu_devices (DeviceId,StatusTime,AliveStatus,AliveLastOkTime) 
	  VALUES (%f,'%s','%s','%s') 
	  RETURNING DeviceId;`, res.DeviceId,
		res.StatusTime, res.AliveStatus, res.AliveLastOkTime)
}

//UpdateSQL
func (res UltDeviceStatuType) UpdateSQL() string {
	var execSqlExt = ""
	if res.AliveStatus == STATU_ACTIVE {
		execSqlExt += ",AliveLastOkTime='" + res.AliveLastOkTime + "'"
	}

	return fmt.Sprintf(`UPDATE public.ult_statu_devices 
	  SET StatusTime='%s',AliveStatus='%s'`+execSqlExt+`
	  WHERE DeviceId=%f  
	  RETURNING DeviceId;`,
		res.StatusTime, res.AliveStatus, res.DeviceId)
}

//SelectWithDb
func (res UltDeviceStatuType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(
		&res.StatusTime,
		&res.AliveStatus,
		&res.AliveLastOkTime)
	return errDb
}
