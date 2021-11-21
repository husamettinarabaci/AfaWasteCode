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
	UltStatus       string
	ContainerStatu  string
	SensPercent     float64
	NewData         bool
}

//New
func (res *UltDeviceStatuType) New() {
	res.DeviceId = 0
	res.StatusTime = GetTime()
	res.ContainerStatu = CONTAINER_FULLNESS_STATU_NONE
	res.UltStatus = ULT_STATU_NONE
	res.AliveStatus = STATU_PASSIVE
	res.AliveLastOkTime = GetTime()
	res.SensPercent = 0
	res.NewData = false
}

//GetByRedis
func (res *UltDeviceStatuType) GetByRedis() ResultType {
	var resultVal ResultType
	resultVal = GetRedisForStoreApi(REDIS_ULT_STATU_DEVICES, res.ToIdString())
	if resultVal.Result == RESULT_OK {
		res.StringToType(resultVal.Retval.(string))
		res.NewData = false
	} else {
		return resultVal
	}

	resultVal.Retval = res.ToString()
	return resultVal
}

//ToId String
func (res *UltDeviceStatuType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToByte
func (res *UltDeviceStatuType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res *UltDeviceStatuType) ToString() string {
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

//ByteToType
func (res *UltDeviceStatuType) ByteToType(retByte []byte) {
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *UltDeviceStatuType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}

//SelectSQL
func (res *UltDeviceStatuType) SelectSQL() string {
	return fmt.Sprintf(`SELECT StatusTime,AliveStatus,AliveLastOkTime,ContainerStatu,UltStatus,SensPercent
	 FROM public.ult_statu_devices
	 WHERE DeviceId=%f ;`, res.DeviceId)
}

//InsertSQL
func (res *UltDeviceStatuType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.ult_statu_devices (DeviceId,StatusTime,AliveStatus,AliveLastOkTime,ContainerStatu,UltStatus,SensPercent) 
	  VALUES (%f,'%s','%s','%s','%s','%s',%f) 
	  RETURNING DeviceId;`, res.DeviceId,
		res.StatusTime, res.AliveStatus, res.AliveLastOkTime, res.ContainerStatu, res.UltStatus, res.SensPercent)
}

//UpdateSQL
func (res *UltDeviceStatuType) UpdateSQL() string {
	var execSqlExt = ""
	if res.AliveStatus == STATU_ACTIVE {
		execSqlExt += ",AliveLastOkTime='" + res.AliveLastOkTime + "'"
	}

	return fmt.Sprintf(`UPDATE public.ult_statu_devices 
	  SET StatusTime='%s',AliveStatus='%s',ContainerStatu='%s',UltStatus='%s',SensPercent=%f`+execSqlExt+`
	  WHERE DeviceId=%f  
	  RETURNING DeviceId;`,
		res.StatusTime, res.AliveStatus, res.ContainerStatu, res.UltStatus, res.SensPercent, res.DeviceId)
}

//SelectWithDb
func (res *UltDeviceStatuType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(
		&res.StatusTime,
		&res.AliveStatus,
		&res.AliveLastOkTime,
		&res.ContainerStatu,
		&res.UltStatus,
		&res.SensPercent)
	return errDb
}
