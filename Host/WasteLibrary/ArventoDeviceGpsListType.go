package WasteLibrary

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
)

//ArventoDeviceGpsListType
type ArventoDeviceGpsListType struct {
	ArventoDeviceGpsList map[string]ArventoDeviceGpsType
}

//ToByte
func (res ArventoDeviceGpsListType) ToByte() []byte {

	bBuf := bytes.Buffer{}
	gobEn := gob.NewEncoder(&bBuf)
	gobEn.Encode(res)
	return bBuf.Bytes()
}

//ToString Get JSON
func (res ArventoDeviceGpsListType) ToString() string {
	return base64.StdEncoding.EncodeToString(res.ToByte())

}

//Byte To ArventoDeviceGpsListType
func ByteToArventoDeviceGpsListType(retByte []byte) ArventoDeviceGpsListType {
	var retVal ArventoDeviceGpsListType
	bBuf := bytes.Buffer{}
	bBuf.Write(retByte)
	gobDe := gob.NewDecoder(&bBuf)
	gobDe.Decode(&retVal)
	return retVal
}

//String To ArventoDeviceGpsListType
func StringToArventoDeviceGpsListType(retStr string) ArventoDeviceGpsListType {
	bStr, _ := base64.StdEncoding.DecodeString(retStr)
	return ByteToArventoDeviceGpsListType(bStr)
}
