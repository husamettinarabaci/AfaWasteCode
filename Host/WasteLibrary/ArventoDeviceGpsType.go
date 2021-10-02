package WasteLibrary

import (
	"encoding/json"
)

//ArventoDeviceGpsType
type ArventoDeviceGpsType struct {
	Latitude  float64
	Longitude float64
	Speed     float64
}

//ToByte
func (res ArventoDeviceGpsType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//ToString Get JSON
func (res ArventoDeviceGpsType) ToString() string {
	return string(res.ToByte())

}

//Byte To ArventoDeviceGpsType
func ByteToArventoDeviceGpsType(retByte []byte) ArventoDeviceGpsType {
	var retVal ArventoDeviceGpsType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To ArventoDeviceGpsType
func StringToArventoDeviceGpsType(retStr string) ArventoDeviceGpsType {
	return ByteToArventoDeviceGpsType([]byte(retStr))
}
