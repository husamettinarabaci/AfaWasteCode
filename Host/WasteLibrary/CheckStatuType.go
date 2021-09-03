package WasteLibrary

import (
	"encoding/json"
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
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res CheckStatuType) ToString() string {
	return string(res.ToByte())
}

//Byte To CheckStatuType
func ByteToCheckStatuType(retByte []byte) CheckStatuType {
	var retVal CheckStatuType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To CheckStatuType
func StringToCheckStatuType(retStr string) CheckStatuType {
	return ByteToCheckStatuType([]byte(retStr))
}
