package WasteLibrary

import (
	"encoding/json"
	"fmt"
)

//CustomerRecyDevicesListType
type CustomerRecyDevicesListType struct {
	CustomerId float64
	Devices    map[string]RecyDeviceType
}

//New
func (res CustomerRecyDevicesListType) New() {
	res.CustomerId = 0
	res.Devices = make(map[string]RecyDeviceType)
}

//ToId String
func (res CustomerRecyDevicesListType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToByte
func (res CustomerRecyDevicesListType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//ToString Get JSON
func (res CustomerRecyDevicesListType) ToString() string {
	return string(res.ToByte())

}

//Byte To CustomerRecyDevicesListType
func ByteToCustomerRecyDevicesListType(retByte []byte) CustomerRecyDevicesListType {
	var retVal CustomerRecyDevicesListType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To CustomerRecyDevicesListType
func StringToCustomerRecyDevicesListType(retStr string) CustomerRecyDevicesListType {
	return ByteToCustomerRecyDevicesListType([]byte(retStr))
}
