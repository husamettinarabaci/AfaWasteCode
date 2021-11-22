package WasteLibrary

import (
	"encoding/json"
	"fmt"
)

//CustomerRfidDevicesType
type CustomerRfidDevicesType struct {
	CustomerId float64
	Devices    map[string]float64
}

//New
func (res *CustomerRfidDevicesType) New() {
	res.CustomerId = 1
	res.Devices = make(map[string]float64)
}

//GetByRedis
func (res *CustomerRfidDevicesType) GetByRedis() ResultType {
	var resultVal ResultType
	resultVal = GetRedisForStoreApi(REDIS_CUSTOMER_RFID_DEVICES, res.ToIdString())
	if resultVal.Result == RESULT_OK {
		res.StringToType(resultVal.Retval.(string))
	} else {
		return resultVal
	}

	resultVal.Retval = res.ToString()
	return resultVal
}

//SaveToRedis
func (res *CustomerRfidDevicesType) SaveToRedis() ResultType {
	var resultVal ResultType
	resultVal = SaveRedisForStoreApi(REDIS_CUSTOMER_RFID_DEVICES, res.ToIdString(), res.ToString())
	return resultVal
}

//ToId String
func (res *CustomerRfidDevicesType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToByte
func (res *CustomerRfidDevicesType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//ToString Get JSON
func (res *CustomerRfidDevicesType) ToString() string {
	return string(res.ToByte())

}

//Byte To CustomerRfidDevicesType
func ByteToCustomerRfidDevicesType(retByte []byte) CustomerRfidDevicesType {
	var retVal CustomerRfidDevicesType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To CustomerRfidDevicesType
func StringToCustomerRfidDevicesType(retStr string) CustomerRfidDevicesType {
	return ByteToCustomerRfidDevicesType([]byte(retStr))
}

//ByteToType
func (res *CustomerRfidDevicesType) ByteToType(retByte []byte) {
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *CustomerRfidDevicesType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}
