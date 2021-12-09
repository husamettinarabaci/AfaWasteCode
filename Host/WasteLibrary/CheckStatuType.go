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

//New
func (res *CheckStatuType) New() {
	res.AppStatu = STATU_ACTIVE
	res.ConnStatu = STATU_PASSIVE
	res.DeviceStatu = STATU_PASSIVE
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

//ByteToType
func (res *CheckStatuType) ByteToType(retByte []byte) {
	res.New()
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *CheckStatuType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}
