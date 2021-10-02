package WasteLibrary

import (
	"encoding/json"
	"fmt"
)

//CustomerTagsType
type CustomerTagsType struct {
	CustomerId float64
	Tags       map[string]TagType
}

//ToId String
func (res CustomerTagsType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToByte
func (res CustomerTagsType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res CustomerTagsType) ToString() string {
	return string(res.ToByte())

}

//Byte To CustomerTagsType
func ByteToCustomerTagsType(retByte []byte) CustomerTagsType {
	var retVal CustomerTagsType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To CustomerTagsType
func StringToCustomerTagsType(retStr string) CustomerTagsType {
	return ByteToCustomerTagsType([]byte(retStr))
}
