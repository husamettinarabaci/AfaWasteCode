package WasteLibrary

import (
	"encoding/json"
	"fmt"
)

//CustomerRecyDevicesType
type CustomerRecyDevicesType struct {
	CustomerId float64
	Devices    map[string]float64
}

//New
func (res *CustomerRecyDevicesType) New() {
	res.CustomerId = 0
	res.Devices = make(map[string]float64)
}

//ToId String
func (res CustomerRecyDevicesType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToByte
func (res CustomerRecyDevicesType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//ToString Get JSON
func (res CustomerRecyDevicesType) ToString() string {
	return string(res.ToByte())

}

//Byte To CustomerRecyDevicesType
func ByteToCustomerRecyDevicesType(retByte []byte) CustomerRecyDevicesType {
	var retVal CustomerRecyDevicesType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To CustomerRecyDevicesType
func StringToCustomerRecyDevicesType(retStr string) CustomerRecyDevicesType {
	return ByteToCustomerRecyDevicesType([]byte(retStr))
}
