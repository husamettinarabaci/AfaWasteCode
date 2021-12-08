package WasteLibrary

import (
	"encoding/json"
)

//CustomersListType
type CustomersListType struct {
	Customers map[string]CustomerType
}

//New
func (res *CustomersListType) New() {
	res.Customers = make(map[string]CustomerType)
}

//ToByte
func (res *CustomersListType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res *CustomersListType) ToString() string {
	return string(res.ToByte())

}

//ByteToType
func (res *CustomersListType) ByteToType(retByte []byte) {
	res.New()
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *CustomersListType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}
