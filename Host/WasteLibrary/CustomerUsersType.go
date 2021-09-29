package WasteLibrary

import (
	"encoding/base64"
	"fmt"
)

//CustomerUsersType
type CustomerUsersType struct {
	CustomerId float64
	Users      map[float64]float64
}

//ToId String
func (res CustomerUsersType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToByte
func (res CustomerUsersType) ToByte() []byte {
	return InterfaceToGobBytes(res)

}

//ToString Get JSON
func (res CustomerUsersType) ToString() string {
	return base64.StdEncoding.EncodeToString(res.ToByte())

}

//Byte To CustomerUsersType
func ByteToCustomerUsersType(retByte []byte) CustomerUsersType {
	var retVal CustomerUsersType
	GobBytestoInterface(retByte, &retVal)
	return retVal
}

//String To CustomerUsersType
func StringToCustomerUsersType(retStr string) CustomerUsersType {
	bStr, _ := base64.StdEncoding.DecodeString(retStr)
	return ByteToCustomerUsersType(bStr)
}
