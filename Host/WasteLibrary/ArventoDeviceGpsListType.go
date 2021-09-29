package WasteLibrary

import (
	"encoding/base64"
)

//ArventoDeviceGpsListType
type ArventoDeviceGpsListType struct {
	ArventoDeviceGpsList map[string]ArventoDeviceGpsType
}

//ToByte
func (res ArventoDeviceGpsListType) ToByte() []byte {
	return InterfaceToGobBytes(res)

}

//ToString Get JSON
func (res ArventoDeviceGpsListType) ToString() string {
	return base64.StdEncoding.EncodeToString(res.ToByte())

}

//Byte To ArventoDeviceGpsListType
func ByteToArventoDeviceGpsListType(retByte []byte) ArventoDeviceGpsListType {
	var retVal ArventoDeviceGpsListType
	GobBytestoInterface(retByte, &retVal)
	return retVal
}

//String To ArventoDeviceGpsListType
func StringToArventoDeviceGpsListType(retStr string) ArventoDeviceGpsListType {
	bStr, _ := base64.StdEncoding.DecodeString(retStr)
	return ByteToArventoDeviceGpsListType(bStr)
}
