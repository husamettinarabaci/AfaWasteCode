package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/url"
)

//RfidDeviceMainType
type RfidDeviceMainType struct {
	DeviceId     float64
	CustomerId   float64
	SerialNumber string
	Active       string
	CreateTime   string
}

//New
func (res *RfidDeviceMainType) New() {
	res.DeviceId = 0
	res.CustomerId = 1
	res.SerialNumber = ""
	res.Active = STATU_ACTIVE
	res.CreateTime = GetTime()
}

//GetByRedis
func (res *RfidDeviceMainType) GetByRedis(dbIndex int) ResultType {

	var resultVal ResultType
	resultVal = GetRedisForStoreApi(REDIS_RFID_MAIN_DEVICES, res.ToIdString())
	if resultVal.Result == RESULT_OK {
		res.StringToType(resultVal.Retval.(string))
	} else {
		return resultVal
	}

	resultVal.Retval = res.ToString()
	return resultVal
}

//SaveToRedis
func (res *RfidDeviceMainType) SaveToRedis() ResultType {
	var resultVal ResultType
	resultVal = SaveRedisForStoreApi(REDIS_RFID_MAIN_DEVICES, res.ToIdString(), res.ToString())
	return resultVal
}

//SaveToRedisBySerial
func (res *RfidDeviceMainType) SaveToRedisBySerial() ResultType {
	var resultVal ResultType
	resultVal = SaveRedisForStoreApi(REDIS_SERIAL_RFID_DEVICE, res.SerialNumber, res.ToIdString())
	return resultVal
}

//SaveToDb
func (res *RfidDeviceMainType) SaveToDb() ResultType {
	var resultVal ResultType
	var currentHttpHeader HttpClientHeaderType
	currentHttpHeader.New()
	currentHttpHeader.DataType = DATATYPE_RFID_MAIN_DEVICE

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

//ToId String
func (res *RfidDeviceMainType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToCustomerId String
func (res *RfidDeviceMainType) ToCustomerIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToByte
func (res *RfidDeviceMainType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res *RfidDeviceMainType) ToString() string {
	return string(res.ToByte())

}

//Byte To RfidDeviceMainType
func ByteToRfidDeviceMainType(retByte []byte) RfidDeviceMainType {
	var retVal RfidDeviceMainType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To RfidDeviceMainType
func StringToRfidDeviceMainType(retStr string) RfidDeviceMainType {
	return ByteToRfidDeviceMainType([]byte(retStr))
}

//ByteToType
func (res *RfidDeviceMainType) ByteToType(retByte []byte) {
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *RfidDeviceMainType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}

//SelectSQL
func (res *RfidDeviceMainType) SelectSQL() string {
	return fmt.Sprintf(`SELECT CustomerId,SerialNumber,Active,CreateTime
	 FROM public.rfid_main_devices
	 WHERE DeviceId=%f  ;`, res.DeviceId)
}

//InsertSQL
func (res *RfidDeviceMainType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.rfid_main_devices (CustomerId,SerialNumber) 
	  VALUES (%f,'%s') 
	  RETURNING DeviceId;`, res.CustomerId, res.SerialNumber)
}

//InsertDataSQL
func (res *RfidDeviceMainType) InsertDataSQL() string {
	return fmt.Sprintf(`INSERT INTO public.rfid_main_devices (DeviceId,CustomerId,SerialNumber) 
	  VALUES (%f,%f,'%s') 
	  RETURNING DeviceId;`, res.DeviceId, res.CustomerId, res.SerialNumber)
}

//UpdateSQL
func (res *RfidDeviceMainType) UpdateSQL() string {
	return fmt.Sprintf(`UPDATE public.rfid_main_devices 
	  SET CustomerId=%f
	  WHERE DeviceId=%f  
	  RETURNING DeviceId;`,
		res.CustomerId,
		res.DeviceId)
}

//SelectWithDb
func (res *RfidDeviceMainType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(
		&res.CustomerId,
		&res.SerialNumber,
		&res.Active,
		&res.CreateTime)
	return errDb
}
