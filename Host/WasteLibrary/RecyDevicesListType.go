package WasteLibrary

import (
	"encoding/json"
)

//RecyDevicesListType
type RecyDevicesListType struct {
	Devices map[string]RecyDeviceType
}

//New
func (res *RecyDevicesListType) New() {
	res.Devices = make(map[string]RecyDeviceType)
}

//ToByte
func (res *RecyDevicesListType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//ToString Get JSON
func (res *RecyDevicesListType) ToString() string {
	return string(res.ToByte())

}

//ByteToType
func (res *RecyDevicesListType) ByteToType(retByte []byte) {
	res.New()
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *RecyDevicesListType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}
