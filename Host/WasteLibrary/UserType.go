package WasteLibrary

import (
	"encoding/base64"
	"fmt"
)

//UserType
type UserType struct {
	UserId       float64
	CustomerId   float64
	UserName     string
	UserType     string
	Email        string
	Token        string
	TokenEndTime string
	CreateTime   string
}

//ToId String
func (res UserType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.UserId)
}

//ToCustomerId String
func (res UserType) ToCustomerIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToByte
func (res UserType) ToByte() []byte {

	return InterfaceToGobBytes(res)
}

//ToString Get JSON
func (res UserType) ToString() string {
	return base64.StdEncoding.EncodeToString(res.ToByte())

}

//Byte To UserType
func ByteToUserType(retByte []byte) UserType {
	var retVal UserType
	GobBytestoInterface(retByte, &retVal)
	return retVal
}

//String To UserType
func StringToUserType(retStr string) UserType {
	bStr, _ := base64.StdEncoding.DecodeString(retStr)
	return ByteToUserType(bStr)
}
