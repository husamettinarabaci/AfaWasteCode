package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

//RecyDeviceVersionType
type RecyDeviceVersionType struct {
	DeviceId           float64
	WebAppVersion      string
	MotorAppVersion    string
	ThermAppVersion    string
	TransferAppVersion string
	CheckerAppVersion  string
	CamAppVersion      string
	ReaderAppVersion   string
	SystemAppVersion   string
	NewData            bool
}

//New
func (res *RecyDeviceVersionType) New() {
	res.DeviceId = 0
	res.WebAppVersion = "1"
	res.MotorAppVersion = "1"
	res.ThermAppVersion = "1"
	res.TransferAppVersion = "1"
	res.CheckerAppVersion = "1"
	res.CamAppVersion = "1"
	res.ReaderAppVersion = "1"
	res.SystemAppVersion = "1"
	res.NewData = false
}

//GetByRedis
func (res *RecyDeviceVersionType) GetByRedis() ResultType {
	var resultVal ResultType
	resultVal = GetRedisForStoreApi(REDIS_RECY_VERSION_DEVICES, res.ToIdString())
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
func (res *RecyDeviceVersionType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToByte
func (res *RecyDeviceVersionType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res *RecyDeviceVersionType) ToString() string {
	return string(res.ToByte())

}

//Byte To RecyDeviceVersionType
func ByteToRecyDeviceVersionType(retByte []byte) RecyDeviceVersionType {
	var retVal RecyDeviceVersionType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To RecyDeviceVersionType
func StringToRecyDeviceVersionType(retStr string) RecyDeviceVersionType {
	return ByteToRecyDeviceVersionType([]byte(retStr))
}

//ByteToType
func (res *RecyDeviceVersionType) ByteToType(retByte []byte) {
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *RecyDeviceVersionType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}

//SelectSQL
func (res *RecyDeviceVersionType) SelectSQL() string {
	return fmt.Sprintf(`SELECT WebAppVersion,MotorAppVersion,ThermAppVersion,TransferAppVersion,CheckerAppVersion,CamAppVersion,ReaderAppVersion,SystemAppVersion
	 FROM public.recy_version_devices
	 WHERE DeviceId=%f ;`, res.DeviceId)
}

//InsertSQL
func (res *RecyDeviceVersionType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.recy_version_devices (DeviceId,WebAppVersion,MotorAppVersion,ThermAppVersion,TransferAppVersion,CheckerAppVersion,CamAppVersion,ReaderAppVersion,SystemAppVersion) 
	  VALUES (%f,'%s','%s','%s','%s','%s','%s','%s','%s') 
	  RETURNING DeviceId;`, res.DeviceId, res.WebAppVersion, res.MotorAppVersion, res.ThermAppVersion, res.TransferAppVersion, res.CheckerAppVersion, res.CamAppVersion, res.ReaderAppVersion, res.SystemAppVersion)
}

//UpdateSQL
func (res *RecyDeviceVersionType) UpdateSQL() string {
	return fmt.Sprintf(`UPDATE public.recy_version_devices 
	  SET WebAppVersion='%s',MotorAppVersion='%s',ThermAppVersion='%s',TransferAppVersion='%s',CheckerAppVersion='%s',CamAppVersion='%s',ReaderAppVersion='%s',SystemAppVersion='%s' 
	  WHERE DeviceId=%f  
	  RETURNING DeviceId;`,
		res.WebAppVersion,
		res.MotorAppVersion,
		res.ThermAppVersion,
		res.TransferAppVersion,
		res.CheckerAppVersion,
		res.CamAppVersion,
		res.ReaderAppVersion,
		res.SystemAppVersion,
		res.DeviceId)
}

//SelectWithDb
func (res *RecyDeviceVersionType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(
		&res.WebAppVersion,
		&res.MotorAppVersion,
		&res.ThermAppVersion,
		&res.TransferAppVersion,
		&res.CheckerAppVersion,
		&res.CamAppVersion,
		&res.ReaderAppVersion,
		&res.SystemAppVersion)
	return errDb
}
