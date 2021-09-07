package WasteLibrary

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
)

//ArventoDeviceGpsType
type ArventoDeviceGpsType struct {
	Latitude  float64
	Longitude float64
	Speed     float64
}

//ToByte
func (res ArventoDeviceGpsType) ToByte() []byte {

	bBuf := bytes.Buffer{}
	gobEn := gob.NewEncoder(&bBuf)
	gobEn.Encode(res)
	return bBuf.Bytes()
}

//ToString Get JSON
func (res ArventoDeviceGpsType) ToString() string {
	return base64.StdEncoding.EncodeToString(res.ToByte())

}

//Byte To ArventoDeviceGpsType
func ByteToArventoDeviceGpsType(retByte []byte) ArventoDeviceGpsType {
	var retVal ArventoDeviceGpsType
	bBuf := bytes.Buffer{}
	bBuf.Write(retByte)
	gobDe := gob.NewDecoder(&bBuf)
	gobDe.Decode(&retVal)
	return retVal
}

//String To ArventoDeviceGpsType
func StringToArventoDeviceGpsType(retStr string) ArventoDeviceGpsType {
	bStr, _ := base64.StdEncoding.DecodeString(retStr)
	return ByteToArventoDeviceGpsType(bStr)
}
