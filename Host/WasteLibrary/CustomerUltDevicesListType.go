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

//ByteToType
func (res *CustomerUltDevicesListType) ByteToType(retByte []byte) {
	res.New()
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *CustomerUltDevicesListType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}
