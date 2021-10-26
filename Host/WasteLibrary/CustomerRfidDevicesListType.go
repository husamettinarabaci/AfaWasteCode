package WasteLibrary

import (
	"encoding/json"
	"fmt"
)

//CustomerRfidDevicesListType
type CustomerRfidDevicesListType struct {
	CustomerId float64
	Devices    map[string]RfidDeviceType
}

//New
func (res CustomerRfidDevicesListType) New() {
	res.CustomerId = 0
	res.Devices = make(map[string]RfidDeviceType)
}

//ToId String
func (res CustomerRfidDevicesListType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToByte
func (res CustomerRfidDevicesListType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//ToString Get JSON
func (res CustomerRfidDevicesListType) ToString() string {
	return string(res.ToByte())

}

//Byte To CustomerRfidDevicesListType
func ByteToCustomerRfidDevicesListType(retByte []byte) CustomerRfidDevicesListType {
	var retVal CustomerRfidDevicesListType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To CustomerRfidDevicesListType
func StringToCustomerRfidDevicesListType(retStr string) CustomerRfidDevicesListType {
	return ByteToCustomerRfidDevicesListType([]byte(retStr))
}
