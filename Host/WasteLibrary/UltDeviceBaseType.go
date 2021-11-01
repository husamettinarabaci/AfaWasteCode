package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

//UltDeviceBaseType
type UltDeviceBaseType struct {
	DeviceId      float64
	ContainerNo   string
	ContainerType string
	DeviceType    string
	Imei          string
	Imsi          string
}

//New
func (res *UltDeviceBaseType) New() {
	res.DeviceId = 0
	res.ContainerNo = ""
	res.ContainerType = CONTAINER_TYPE_NONE
	res.DeviceType = ULT_DEVICE_TYPE_NONE
	res.Imei = ""
	res.Imsi = ""
}

//ToId String
func (res UltDeviceBaseType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToByte
func (res UltDeviceBaseType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res UltDeviceBaseType) ToString() string {
	return string(res.ToByte())

}

//Byte To UltDeviceBaseType
func ByteToUltDeviceBaseType(retByte []byte) UltDeviceBaseType {
	var retVal UltDeviceBaseType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To UltDeviceBaseType
func StringToUltDeviceBaseType(retStr string) UltDeviceBaseType {
	return ByteToUltDeviceBaseType([]byte(retStr))
}

//SelectSQL
func (res UltDeviceBaseType) SelectSQL() string {
	return fmt.Sprintf(`SELECT ContainerNo,ContainerType,DeviceType,Imei,Imsi
	 FROM public.ult_base_devices
	 WHERE DeviceId=%f ;`, res.DeviceId)
}

//InsertSQL
func (res UltDeviceBaseType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.ult_base_devices (DeviceId,ContainerNo,ContainerType,DeviceType,Imei,Imsi) 
	  VALUES (%f,'%s','%s','%s','%s','%s') 
	  RETURNING DeviceId;`, res.DeviceId, res.ContainerNo, res.ContainerType, res.DeviceType, res.Imei, res.Imsi)
}

//UpdateSQL
func (res UltDeviceBaseType) UpdateSQL() string {
	return fmt.Sprintf(`UPDATE public.ult_base_devices 
	  SET ContainerNo='%s',ContainerType='%s',DeviceType='%s',Imei='%s',Imsi='%s'
	  WHERE DeviceId=%f  
	  RETURNING DeviceId;`,
		res.ContainerNo,
		res.ContainerType,
		res.DeviceType,
		res.Imei,
		res.Imsi,
		res.DeviceId)
}

//SelectWithDb
func (res UltDeviceBaseType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(
		&res.ContainerNo,
		&res.ContainerType,
		&res.DeviceType,
		&res.Imei,
		&res.Imsi)
	return errDb
}
