package WasteLibrary

import (
	"encoding/json"
	"fmt"
)

//CustomerDevicesListType
type CustomerDevicesListType struct {
	CustomerId float64
	Devices    map[string]DeviceType
}

//New
func (res CustomerDevicesListType) New() {
	res.CustomerId = 0
	res.Devices = make(map[string]DeviceType)
}

//ToId String
func (res CustomerDevicesListType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToByte
func (res CustomerDevicesListType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//ToString Get JSON
func (res CustomerDevicesListType) ToString() string {
	return string(res.ToByte())

}

//Byte To CustomerDevicesListType
func ByteToCustomerDevicesListType(retByte []byte) CustomerDevicesListType {
	var retVal CustomerDevicesListType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To CustomerDevicesListType
func StringToCustomerDevicesListType(retStr string) CustomerDevicesListType {
	return ByteToCustomerDevicesListType([]byte(retStr))
}
