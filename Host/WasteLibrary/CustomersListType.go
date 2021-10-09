package WasteLibrary

import (
	"encoding/json"
)

//CustomersListType
type CustomersListType struct {
	Customers map[string]CustomerType
}

//ToByte
func (res CustomersListType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res CustomersListType) ToString() string {
	return string(res.ToByte())

}

//Byte To CustomersListType
func ByteToCustomersListType(retByte []byte) CustomersListType {
	var retVal CustomersListType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To CustomersListType
func StringToCustomersListType(retStr string) CustomersListType {
	return ByteToCustomersListType([]byte(retStr))
}
