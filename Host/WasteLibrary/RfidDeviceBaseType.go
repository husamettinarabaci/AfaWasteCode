package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

//RfidDeviceBaseType
type RfidDeviceBaseType struct {
	DeviceId   float64
	DeviceType string
	TruckType  string
	NewData    bool
}

//New
func (res *RfidDeviceBaseType) New() {
	res.DeviceId = 0
	res.DeviceType = RFID_DEVICE_TYPE_NONE
	res.TruckType = TRUCKTYPE_NONE
	res.NewData = false
}

//GetByRedis
func (res *RfidDeviceBaseType) GetByRedis() ResultType {

	var resultVal ResultType
	resultVal = GetRedisForStoreApi(REDIS_RFID_BASE_DEVICES, res.ToIdString())
	if resultVal.Result == RESULT_OK {
		res.StringToType(resultVal.Retval.(string))
		res.NewData = false
	} else {
		return resultVal
	}

	resultVal.Retval = res.ToString()
	return resultVal
}

//SaveToRedis
func (res *RfidDeviceBaseType) SaveToRedis() ResultType {
	var resultVal ResultType
	resultVal = SaveRedisForStoreApi(REDIS_RFID_BASE_DEVICES, res.ToIdString(), res.ToString())
	return resultVal
}

//ToId String
func (res *RfidDeviceBaseType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToByte
func (res *RfidDeviceBaseType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res *RfidDeviceBaseType) ToString() string {
	return string(res.ToByte())

}

//Byte To RfidDeviceBaseType
func ByteToRfidDeviceBaseType(retByte []byte) RfidDeviceBaseType {
	var retVal RfidDeviceBaseType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To RfidDeviceBaseType
func StringToRfidDeviceBaseType(retStr string) RfidDeviceBaseType {
	return ByteToRfidDeviceBaseType([]byte(retStr))
}

//ByteToType
func (res *RfidDeviceBaseType) ByteToType(retByte []byte) {
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *RfidDeviceBaseType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}

//SelectSQL
func (res *RfidDeviceBaseType) SelectSQL() string {
	return fmt.Sprintf(`SELECT DeviceType,TruckType
	 FROM public.rfid_base_devices
	 WHERE DeviceId=%f ;`, res.DeviceId)
}

//InsertSQL
func (res *RfidDeviceBaseType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.rfid_base_devices (DeviceId,DeviceType,TruckType) 
	  VALUES (%f,'%s','%s') 
	  RETURNING DeviceId;`, res.DeviceId, res.DeviceType, res.TruckType)
}

//UpdateSQL
func (res *RfidDeviceBaseType) UpdateSQL() string {
	return fmt.Sprintf(`UPDATE public.rfid_base_devices 
	  SET DeviceType='%s',TruckType='%s'
	  WHERE DeviceId=%f  
	  RETURNING DeviceId;`,
		res.DeviceType,
		res.TruckType,
		res.DeviceId)
}

//SelectWithDb
func (res *RfidDeviceBaseType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(
		&res.DeviceType,
		&res.TruckType)
	return errDb
}
