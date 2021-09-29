package WasteLibrary

import (
	"encoding/base64"
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
	return InterfaceToGobBytes(res)

}

//ToString Get JSON
func (res CustomerType) ToString() string {
	return base64.StdEncoding.EncodeToString(res.ToByte())

}

//Byte To CustomerType
func ByteToCustomerType(retByte []byte) CustomerType {
	var retVal CustomerType
	GobBytestoInterface(retByte, &retVal)
	return retVal
}

//String To CustomerType
func StringToCustomerType(retStr string) CustomerType {
	bStr, _ := base64.StdEncoding.DecodeString(retStr)
	return ByteToCustomerType(bStr)
}
