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
	AppStatu:    STATU_ACTIVE,
	ConnStatu:   STATU_PASSIVE,
	DeviceStatu: STATU_PASSIVE,
}

//ToByte
func (res *CheckStatuType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//ToString Get JSON
func (res *CheckStatuType) ToString() string {
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
