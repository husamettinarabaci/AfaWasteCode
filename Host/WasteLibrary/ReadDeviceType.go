package WasteLibrary

import (
	"encoding/json"
)

//ReadDeviceType
type ReadDeviceType struct {
	DeviceId float64
	ReadTime string
}

//New
func (res *ReadDeviceType) New() {
	res.DeviceId = 0
	res.ReadTime = GetTime()
}

//ToByte
func (res *ReadDeviceType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//ToString Get JSON
func (res *ReadDeviceType) ToString() string {
	return string(res.ToByte())

}

//ByteToType
func (res *ReadDeviceType) ByteToType(retByte []byte) {
	res.New()
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *ReadDeviceType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}
