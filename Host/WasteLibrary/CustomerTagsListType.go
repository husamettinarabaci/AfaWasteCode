package WasteLibrary

import (
	"encoding/json"
	"fmt"
)

//CustomerTagsListType
type CustomerTagsListType struct {
	CustomerId float64
	Tags       map[string]TagViewType
}

//New
func (res *CustomerTagsListType) New() {
	res.CustomerId = 1
	res.Tags = make(map[string]TagViewType)
}

//ToId String
func (res *CustomerTagsListType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToByte
func (res *CustomerTagsListType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//ToString Get JSON
func (res *CustomerTagsListType) ToString() string {
	return string(res.ToByte())

}

//Byte To CustomerTagsListType
func ByteToCustomerTagsListType(retByte []byte) CustomerTagsListType {
	var retVal CustomerTagsListType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To CustomerTagsListType
func StringToCustomerTagsListType(retStr string) CustomerTagsListType {
	return ByteToCustomerTagsListType([]byte(retStr))
}
