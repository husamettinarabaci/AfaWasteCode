package WasteLibrary

import (
	"encoding/json"
)

//UltDevicesListType
type UltDevicesListType struct {
	Devices map[string]UltDeviceType
}

//New
func (res *UltDevicesListType) New() {
	res.Devices = make(map[string]UltDeviceType)
}

//ToByte
func (res *UltDevicesListType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//ToString Get JSON
func (res *UltDevicesListType) ToString() string {
	return string(res.ToByte())

}

//ByteToType
func (res *UltDevicesListType) ByteToType(retByte []byte) {
	res.New()
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *UltDevicesListType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}
