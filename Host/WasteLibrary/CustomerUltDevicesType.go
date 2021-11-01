package WasteLibrary

import (
	"encoding/json"
	"fmt"
)

//CustomerUltDevicesType
type CustomerUltDevicesType struct {
	CustomerId float64
	Devices    map[string]float64
}

//New
func (res *CustomerUltDevicesType) New() {
	res.CustomerId = 0
	res.Devices = make(map[string]float64)
}

//ToId String
func (res CustomerUltDevicesType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToByte
func (res CustomerUltDevicesType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//ToString Get JSON
func (res CustomerUltDevicesType) ToString() string {
	return string(res.ToByte())

}

//Byte To CustomerUltDevicesType
func ByteToCustomerUltDevicesType(retByte []byte) CustomerUltDevicesType {
	var retVal CustomerUltDevicesType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To CustomerUltDevicesType
func StringToCustomerUltDevicesType(retStr string) CustomerUltDevicesType {
	return ByteToCustomerUltDevicesType([]byte(retStr))
}
