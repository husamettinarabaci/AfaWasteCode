package WasteLibrary

import (
	"encoding/json"
	"fmt"
)

//CustomerUsersListType
type CustomerUsersListType struct {
	CustomerId float64
	Users      map[string]UserType
}

//ToId String
func (res CustomerUsersListType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToByte
func (res CustomerUsersListType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//ToString Get JSON
func (res CustomerUsersListType) ToString() string {
	return string(res.ToByte())

}

//Byte To CustomerUsersListType
func ByteToCustomerUsersListType(retByte []byte) CustomerUsersListType {
	var retVal CustomerUsersListType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To CustomerUsersListType
func StringToCustomerUsersListType(retStr string) CustomerUsersListType {
	return ByteToCustomerUsersListType([]byte(retStr))
}
