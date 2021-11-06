package WasteLibrary

import (
	"encoding/json"
	"fmt"
)

//CustomerUltDevicesListType
type CustomerUltDevicesListType struct {
	CustomerId float64
	Devices    map[string]UltDeviceType
}

//New
func (res *CustomerUltDevicesListType) New() {
	res.CustomerId = 1
	res.Devices = make(map[string]UltDeviceType)
}

//ToId String
func (res *CustomerUltDevicesListType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToByte
func (res *CustomerUltDevicesListType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//ToString Get JSON
func (res *CustomerUltDevicesListType) ToString() string {
	return string(res.ToByte())

}

//Byte To CustomerUltDevicesListType
func ByteToCustomerUltDevicesListType(retByte []byte) CustomerUltDevicesListType {
	var retVal CustomerUltDevicesListType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To CustomerUltDevicesListType
func StringToCustomerUltDevicesListType(retStr string) CustomerUltDevicesListType {
	return ByteToCustomerUltDevicesListType([]byte(retStr))
}
