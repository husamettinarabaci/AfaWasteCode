package WasteLibrary

import (
	"encoding/base64"
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
	return InterfaceToGobBytes(res)

}

//ToString Get JSON
func (res CheckStatuType) ToString() string {
	return base64.StdEncoding.EncodeToString(res.ToByte())

}

//Byte To CheckStatuType
func ByteToCheckStatuType(retByte []byte) CheckStatuType {
	var retVal CheckStatuType
	GobBytestoInterface(retByte, &retVal)
	return retVal
}

//String To CheckStatuType
func StringToCheckStatuType(retStr string) CheckStatuType {
	bStr, _ := base64.StdEncoding.DecodeString(retStr)
	return ByteToCheckStatuType(bStr)
}
