package WasteLibrary

import (
	"encoding/json"
	"fmt"
)

//CustomerDevicesType
type CustomerDevicesType struct {
	CustomerId float64
	Devices    map[string]float64
}

//ToId String
func (res CustomerDevicesType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToByte
func (res CustomerDevicesType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//ToString Get JSON
func (res CustomerDevicesType) ToString() string {
	return string(res.ToByte())

}

//Byte To CustomerDevicesType
func ByteToCustomerDevicesType(retByte []byte) CustomerDevicesType {
	var retVal CustomerDevicesType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To CustomerDevicesType
func StringToCustomerDevicesType(retStr string) CustomerDevicesType {
	return ByteToCustomerDevicesType([]byte(retStr))
}
