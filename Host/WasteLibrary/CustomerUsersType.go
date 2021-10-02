package WasteLibrary

import (
	"encoding/json"
	"fmt"
)

//CustomerUsersType
type CustomerUsersType struct {
	CustomerId float64
	Users      map[string]float64
}

//ToId String
func (res CustomerUsersType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToByte
func (res CustomerUsersType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//ToString Get JSON
func (res CustomerUsersType) ToString() string {
	return string(res.ToByte())

}

//Byte To CustomerUsersType
func ByteToCustomerUsersType(retByte []byte) CustomerUsersType {
	var retVal CustomerUsersType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To CustomerUsersType
func StringToCustomerUsersType(retStr string) CustomerUsersType {
	return ByteToCustomerUsersType([]byte(retStr))
}
