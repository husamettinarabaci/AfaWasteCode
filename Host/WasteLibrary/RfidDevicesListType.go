package WasteLibrary

import (
	"encoding/json"
)

//RfidDevicesListType
type RfidDevicesListType struct {
	Devices map[string]RfidDeviceType
}

//New
func (res *RfidDevicesListType) New() {
	res.Devices = make(map[string]RfidDeviceType)
}

//ToByte
func (res *RfidDevicesListType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//ToString Get JSON
func (res *RfidDevicesListType) ToString() string {
	return string(res.ToByte())

}

//ByteToType
func (res *RfidDevicesListType) ByteToType(retByte []byte) {
	res.New()
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *RfidDevicesListType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}
