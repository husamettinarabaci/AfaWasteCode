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
func (res *CustomerRfidDevicesListType) New() {
	res.CustomerId = 1
	res.Devices = make(map[string]RfidDeviceType)
}

//ToId String
func (res *CustomerRfidDevicesListType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToByte
func (res *CustomerRfidDevicesListType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//ToString Get JSON
func (res *CustomerRfidDevicesListType) ToString() string {
	return string(res.ToByte())

}

//ByteToType
func (res *CustomerRfidDevicesListType) ByteToType(retByte []byte) {
	res.New()
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *CustomerRfidDevicesListType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}
