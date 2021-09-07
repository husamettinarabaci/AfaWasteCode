package WasteLibrary

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
)

//CheckStatuType
type CheckStatuType struct {
	AppStatu    string
	ConnStatu   string
	DeviceStatu string
}

var CurrentCheckStatu CheckStatuType = CheckStatuType{
	AppStatu:    "1",
	ConnStatu:   "0",
	DeviceStatu: "0",
}

//ToByte
func (res CheckStatuType) ToByte() []byte {

	bBuf := bytes.Buffer{}
	gobEn := gob.NewEncoder(&bBuf)
	gobEn.Encode(res)
	return bBuf.Bytes()
}

//ToString Get JSON
func (res CheckStatuType) ToString() string {
	return base64.StdEncoding.EncodeToString(res.ToByte())

}

//Byte To CheckStatuType
func ByteToCheckStatuType(retByte []byte) CheckStatuType {
	var retVal CheckStatuType
	bBuf := bytes.Buffer{}
	bBuf.Write(retByte)
	gobDe := gob.NewDecoder(&bBuf)
	gobDe.Decode(&retVal)
	return retVal
}

//String To CheckStatuType
func StringToCheckStatuType(retStr string) CheckStatuType {
	bStr, _ := base64.StdEncoding.DecodeString(retStr)
	return ByteToCheckStatuType(bStr)
}
