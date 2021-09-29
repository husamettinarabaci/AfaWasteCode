package WasteLibrary

import (
	"encoding/base64"
	"fmt"
)

//AdminConfigType
type AdminConfigType struct {
	CustomerId float64
	Active     string
	CreateTime string
}

//ToId String
func (res AdminConfigType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToByte
func (res AdminConfigType) ToByte() []byte {
	return InterfaceToGobBytes(res)

}

//ToString Get JSON
func (res AdminConfigType) ToString() string {
	return base64.StdEncoding.EncodeToString(res.ToByte())

}

//Byte To AdminConfigType
func ByteToAdminConfigType(retByte []byte) AdminConfigType {
	var retVal AdminConfigType
	GobBytestoInterface(retByte, &retVal)
	return retVal
}

//String To AdminConfigType
func StringToAdminConfigType(retStr string) AdminConfigType {
	bStr, _ := base64.StdEncoding.DecodeString(retStr)
	return ByteToAdminConfigType(bStr)
}
