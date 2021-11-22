package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/url"
)

//UltDeviceBatteryType
type UltDeviceBatteryType struct {
	DeviceId      float64
	Battery       string
	BatteryStatus string
	BatteryTime   string
	NewData       bool
}

//New
func (res *UltDeviceBatteryType) New() {
	res.DeviceId = 0
	res.Battery = "0000"
	res.BatteryStatus = BATTERYSTATU_NONE
	res.BatteryTime = GetTime()
	res.NewData = false
}

//GetByRedis
func (res *UltDeviceBatteryType) GetByRedis() ResultType {
	var resultVal ResultType
	resultVal = GetRedisForStoreApi(REDIS_ULT_BATTERY_DEVICES, res.ToIdString())
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
func (res *UltDeviceBatteryType) SaveToRedis() ResultType {
	var resultVal ResultType
	resultVal = SaveRedisForStoreApi(REDIS_ULT_BATTERY_DEVICES, res.ToIdString(), res.ToString())
	return resultVal
}

//SaveToDb
func (res *UltDeviceBatteryType) SaveToDb() ResultType {
	var resultVal ResultType
	var currentHttpHeader HttpClientHeaderType
	currentHttpHeader.New()
	currentHttpHeader.DataType = DATATYPE_ULT_BATTERY_DEVICE

	data := url.Values{
		HTTP_HEADER: {currentHttpHeader.ToString()},
		HTTP_DATA:   {res.ToString()},
	}
	resultVal = SaveStaticDbMainForStoreApi(data)
	if resultVal.Result == RESULT_OK {
		res.DeviceId = StringIdToFloat64(resultVal.Retval.(string))
		resultVal.Retval = res.ToString()
	}

	return resultVal
}

//SaveToReaderDb
func (res *UltDeviceBatteryType) SaveToReaderDb() ResultType {
	var resultVal ResultType
	var currentHttpHeader HttpClientHeaderType
	currentHttpHeader.New()
	currentHttpHeader.DataType = DATATYPE_ULT_BATTERY_DEVICE

	data := url.Values{
		HTTP_HEADER: {currentHttpHeader.ToString()},
		HTTP_DATA:   {res.ToString()},
	}
	resultVal = SaveReaderDbMainForStoreApi(data)
	if resultVal.Result == RESULT_OK {
		res.DeviceId = StringIdToFloat64(resultVal.Retval.(string))
		resultVal.Retval = res.ToString()
	}

	return resultVal
}

//ToId String
func (res *UltDeviceBatteryType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToByte
func (res *UltDeviceBatteryType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res *UltDeviceBatteryType) ToString() string {
	return string(res.ToByte())

}

//Byte To UltDeviceBatteryType
func ByteToUltDeviceBatteryType(retByte []byte) UltDeviceBatteryType {
	var retVal UltDeviceBatteryType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To UltDeviceBatteryType
func StringToUltDeviceBatteryType(retStr string) UltDeviceBatteryType {
	return ByteToUltDeviceBatteryType([]byte(retStr))
}

//ByteToType
func (res *UltDeviceBatteryType) ByteToType(retByte []byte) {
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *UltDeviceBatteryType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}

//SelectSQL
func (res *UltDeviceBatteryType) SelectSQL() string {
	return fmt.Sprintf(`SELECT Battery,BatteryStatus,BatteryTime
	 FROM public.ult_battery_devices
	 WHERE DeviceId=%f ;`, res.DeviceId)
}

//InsertSQL
func (res *UltDeviceBatteryType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.ult_battery_devices (DeviceId,Battery,BatteryStatus,BatteryTime) 
	  VALUES (%f,'%s','%s','%s') 
	  RETURNING DeviceId;`, res.DeviceId, res.Battery, res.BatteryStatus, res.BatteryTime)
}

//UpdateSQL
func (res *UltDeviceBatteryType) UpdateSQL() string {
	return fmt.Sprintf(`UPDATE public.ult_battery_devices 
	  SET Battery='%s',BatteryStatus='%s',BatteryTime='%s'
	  WHERE DeviceId=%f  
	  RETURNING DeviceId;`,
		res.Battery,
		res.BatteryStatus,
		res.BatteryTime,
		res.DeviceId)
}

//SelectWithDb
func (res *UltDeviceBatteryType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(
		&res.Battery,
		&res.BatteryStatus,
		&res.BatteryTime)
	return errDb
}
