package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/url"
)

//RfidDeviceThermType
type RfidDeviceThermType struct {
	DeviceId    float64
	Therm       string
	ThermTime   string
	ThermStatus string
	NewData     bool
}

//New
func (res *RfidDeviceThermType) New() {
	res.DeviceId = 0
	res.Therm = "00"
	res.ThermTime = GetTime()
	res.ThermStatus = THERMSTATU_NONE
	res.NewData = false
}

//GetByRedis
func (res *RfidDeviceThermType) GetByRedis(dbIndex string) ResultType {

	var resultVal ResultType
	resultVal = GetRedisForStoreApi(dbIndex, REDIS_RFID_THERM_DEVICES, res.ToIdString())
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
func (res *RfidDeviceThermType) SaveToRedis() ResultType {
	var resultVal ResultType
	resultVal = SaveRedisForStoreApi(REDIS_RFID_THERM_DEVICES, res.ToIdString(), res.ToString())
	return resultVal
}

//SaveToDb
func (res *RfidDeviceThermType) SaveToDb() ResultType {
	var resultVal ResultType
	var currentHttpHeader HttpClientHeaderType
	currentHttpHeader.New()
	currentHttpHeader.DataType = DATATYPE_RFID_THERM_DEVICE

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
func (res *RfidDeviceThermType) SaveToReaderDb() ResultType {
	var resultVal ResultType
	var currentHttpHeader HttpClientHeaderType
	currentHttpHeader.New()
	currentHttpHeader.DataType = DATATYPE_RFID_THERM_DEVICE

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
func (res *RfidDeviceThermType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToByte
func (res *RfidDeviceThermType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res *RfidDeviceThermType) ToString() string {
	return string(res.ToByte())

}

//Byte To RfidDeviceThermType
func ByteToRfidDeviceThermType(retByte []byte) RfidDeviceThermType {
	var retVal RfidDeviceThermType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To RfidDeviceThermType
func StringToRfidDeviceThermType(retStr string) RfidDeviceThermType {
	return ByteToRfidDeviceThermType([]byte(retStr))
}

//ByteToType
func (res *RfidDeviceThermType) ByteToType(retByte []byte) {
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *RfidDeviceThermType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}

//SelectSQL
func (res *RfidDeviceThermType) SelectSQL() string {
	return fmt.Sprintf(`SELECT Therm,ThermTime,ThermStatus
	 FROM public.rfid_therm_devices
	 WHERE DeviceId=%f ;`, res.DeviceId)
}

//InsertSQL
func (res *RfidDeviceThermType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.rfid_therm_devices (DeviceId,Therm,ThermTime,ThermStatus) 
	  VALUES (%f,'%s','%s','%s') 
	  RETURNING DeviceId;`, res.DeviceId, res.Therm, res.ThermTime, res.ThermStatus)
}

//UpdateSQL
func (res *RfidDeviceThermType) UpdateSQL() string {
	return fmt.Sprintf(`UPDATE public.rfid_therm_devices 
	  SET Therm='%s',ThermTime='%s',ThermStatus='%s' 
	  WHERE DeviceId=%f  
	  RETURNING DeviceId;`,
		res.Therm,
		res.ThermTime,
		res.ThermStatus,
		res.DeviceId)
}

//SelectWithDb
func (res *RfidDeviceThermType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(
		&res.Therm,
		&res.ThermTime,
		&res.ThermStatus)
	return errDb
}
