package WasteLibrary

import (
	"encoding/json"
	"fmt"
)

//AdminConfigType
type AdminConfigType struct {
	CustomerId float64
	Active     string
	CreateTime string
}

//New
func (res *AdminConfigType) New() {
	res.CustomerId = 0
	res.Active = STATU_ACTIVE
	res.CreateTime = GetTime()
}

//ToId String
func (res AdminConfigType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToByte
func (res AdminConfigType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//ToString Get JSON
func (res AdminConfigType) ToString() string {
	return string(res.ToByte())

}

//Byte To AdminConfigType
func ByteToAdminConfigType(retByte []byte) AdminConfigType {
	var retVal AdminConfigType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To AdminConfigType
func StringToAdminConfigType(retStr string) AdminConfigType {
	return ByteToAdminConfigType([]byte(retStr))
}
