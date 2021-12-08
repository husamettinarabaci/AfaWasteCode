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
func (res *CustomerRecyDevicesListType) New() {
	res.CustomerId = 1
	res.Devices = make(map[string]RecyDeviceType)
}

//ToId String
func (res *CustomerRecyDevicesListType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToByte
func (res *CustomerRecyDevicesListType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//ToString Get JSON
func (res *CustomerRecyDevicesListType) ToString() string {
	return string(res.ToByte())

}

//ByteToType
func (res *CustomerRecyDevicesListType) ByteToType(retByte []byte) {
	res.New()
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *CustomerRecyDevicesListType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}
