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

//New
func (res *CustomerUsersListType) New() {
	res.CustomerId = 1
	res.Users = make(map[string]UserType)
}

//ToId String
func (res *CustomerUsersListType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToByte
func (res *CustomerUsersListType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//ToString Get JSON
func (res *CustomerUsersListType) ToString() string {
	return string(res.ToByte())

}

//ByteToType
func (res *CustomerUsersListType) ByteToType(retByte []byte) {
	res.New()
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *CustomerUsersListType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}
