package WasteLibrary

import (
	"encoding/json"
)

//ArventoDeviceGpsListType
type ArventoDeviceGpsListType struct {
	ArventoDeviceGpsList map[string]ArventoDeviceGpsType
}

//New
func (res *ArventoDeviceGpsListType) New() {
	res.ArventoDeviceGpsList = make(map[string]ArventoDeviceGpsType)

}

//ToByte
func (res *ArventoDeviceGpsListType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//ToString Get JSON
func (res *ArventoDeviceGpsListType) ToString() string {
	return string(res.ToByte())

}

//ByteToType
func (res *ArventoDeviceGpsListType) ByteToType(retByte []byte) {
	res.New()
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *ArventoDeviceGpsListType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}
