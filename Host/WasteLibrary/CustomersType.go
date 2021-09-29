package WasteLibrary

import (
	"encoding/base64"
)

//CustomersType
type CustomersType struct {
	Customers map[float64]float64
}

//ToByte
func (res CustomersType) ToByte() []byte {
	return InterfaceToGobBytes(res)

}

//ToString Get JSON
func (res CustomersType) ToString() string {
	return base64.StdEncoding.EncodeToString(res.ToByte())

}

//Byte To CustomersType
func ByteToCustomersType(retByte []byte) CustomersType {
	var retVal CustomersType
	GobBytestoInterface(retByte, &retVal)
	return retVal
}

//String To CustomersType
func StringToCustomersType(retStr string) CustomersType {
	bStr, _ := base64.StdEncoding.DecodeString(retStr)
	return ByteToCustomersType(bStr)
}
