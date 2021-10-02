package WasteLibrary

import (
	"encoding/json"
)

//ArventoDeviceGpsListType
type ArventoDeviceGpsListType struct {
	ArventoDeviceGpsList map[string]ArventoDeviceGpsType
}

//ToByte
func (res ArventoDeviceGpsListType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//ToString Get JSON
func (res ArventoDeviceGpsListType) ToString() string {
	return string(res.ToByte())

}

//Byte To ArventoDeviceGpsListType
func ByteToArventoDeviceGpsListType(retByte []byte) ArventoDeviceGpsListType {
	var retVal ArventoDeviceGpsListType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To ArventoDeviceGpsListType
func StringToArventoDeviceGpsListType(retStr string) ArventoDeviceGpsListType {
	return ByteToArventoDeviceGpsListType([]byte(retStr))
}
