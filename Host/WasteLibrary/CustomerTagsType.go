package WasteLibrary

import (
	"encoding/base64"
	"fmt"
)

//CustomerTagsType
type CustomerTagsType struct {
	CustomerId float64
	Tags       map[float64]TagType
}

//ToId String
func (res CustomerTagsType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToByte
func (res CustomerTagsType) ToByte() []byte {
	return InterfaceToGobBytes(res)

}

//ToString Get JSON
func (res CustomerTagsType) ToString() string {
	return base64.StdEncoding.EncodeToString(res.ToByte())

}

//Byte To CustomerTagsType
func ByteToCustomerTagsType(retByte []byte) CustomerTagsType {
	var retVal CustomerTagsType
	GobBytestoInterface(retByte, &retVal)
	return retVal
}

//String To CustomerTagsType
func StringToCustomerTagsType(retStr string) CustomerTagsType {
	bStr, _ := base64.StdEncoding.DecodeString(retStr)
	return ByteToCustomerTagsType(bStr)
}
