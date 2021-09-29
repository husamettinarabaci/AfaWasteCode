package WasteLibrary

import (
	"encoding/base64"
)

//ArventoDeviceGpsType
type ArventoDeviceGpsType struct {
	Latitude  float64
	Longitude float64
	Speed     float64
}

//ToByte
func (res ArventoDeviceGpsType) ToByte() []byte {
	return InterfaceToGobBytes(res)

}

//ToString Get JSON
func (res ArventoDeviceGpsType) ToString() string {
	return base64.StdEncoding.EncodeToString(res.ToByte())

}

//Byte To ArventoDeviceGpsType
func ByteToArventoDeviceGpsType(retByte []byte) ArventoDeviceGpsType {
	var retVal ArventoDeviceGpsType
	GobBytestoInterface(retByte, &retVal)
	return retVal
}

//String To ArventoDeviceGpsType
func StringToArventoDeviceGpsType(retStr string) ArventoDeviceGpsType {
	bStr, _ := base64.StdEncoding.DecodeString(retStr)
	return ByteToArventoDeviceGpsType(bStr)
}
