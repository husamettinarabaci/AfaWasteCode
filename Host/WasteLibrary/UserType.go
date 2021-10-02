package WasteLibrary

import (
	"encoding/json"
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
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res UserType) ToString() string {
	return string(res.ToByte())

}

//Byte To UserType
func ByteToUserType(retByte []byte) UserType {
	var retVal UserType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To UserType
func StringToUserType(retStr string) UserType {
	return ByteToUserType([]byte(retStr))
}
