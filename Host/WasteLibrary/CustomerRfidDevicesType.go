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
func (res CustomerRfidDevicesType) New() {
	res.CustomerId = 0
	res.Devices = make(map[string]float64)
}

//ToId String
func (res CustomerRfidDevicesType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToByte
func (res CustomerRfidDevicesType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//ToString Get JSON
func (res CustomerRfidDevicesType) ToString() string {
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
