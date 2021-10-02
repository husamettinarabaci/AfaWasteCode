package WasteLibrary

import (
	"encoding/json"
	"fmt"
)

//CustomerType
type CustomerType struct {
	CustomerId   float64
	CustomerName string
	AdminLink    string
	WebLink      string
	ReportLink   string
	RfIdApp      string
	UltApp       string
	RecyApp      string
	Active       string
	CreateTime   string
}

//ToId String
func (res CustomerType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToByte
func (res CustomerType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//ToString Get JSON
func (res CustomerType) ToString() string {
	return string(res.ToByte())

}

//Byte To CustomerType
func ByteToCustomerType(retByte []byte) CustomerType {
	var retVal CustomerType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To CustomerType
func StringToCustomerType(retStr string) CustomerType {
	return ByteToCustomerType([]byte(retStr))
}
