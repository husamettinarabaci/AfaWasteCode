package WasteLibrary

import (
	"encoding/base64"
	"fmt"
)

//LocalConfigType
type LocalConfigType struct {
	CustomerId float64
	Active     string
	CreateTime string
}

//ToId String
func (res LocalConfigType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToByte
func (res LocalConfigType) ToByte() []byte {
	return InterfaceToGobBytes(res)

}

//ToString Get JSON
func (res LocalConfigType) ToString() string {
	return base64.StdEncoding.EncodeToString(res.ToByte())

}

//Byte To LocalConfigType
func ByteToLocalConfigType(retByte []byte) LocalConfigType {
	var retVal LocalConfigType
	GobBytestoInterface(retByte, &retVal)
	return retVal
}

//String To LocalConfigType
func StringToLocalConfigType(retStr string) LocalConfigType {
	bStr, _ := base64.StdEncoding.DecodeString(retStr)
	return ByteToLocalConfigType(bStr)
}
