package WasteLibrary

import (
	"encoding/json"
	"fmt"
)

//UpdaterType
type UpdaterType struct {
	DeviceId   float64
	AppType    string
	Version    string
	Active     string
	CreateTime string
}

//New
func (res *UpdaterType) New() {
	res.DeviceId = 0
	res.AppType = RFID_APPNAME_NONE
	res.Version = "1"
	res.Active = STATU_ACTIVE
	res.CreateTime = GetTime()
}

//ToId String
func (res *UpdaterType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToByte
func (res *UpdaterType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//ToString Get JSON
func (res *UpdaterType) ToString() string {
	return string(res.ToByte())

}

//Byte To UpdaterType
func ByteToUpdaterType(retByte []byte) UpdaterType {
	var retVal UpdaterType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To UpdaterType
func StringToUpdaterType(retStr string) UpdaterType {
	return ByteToUpdaterType([]byte(retStr))
}
