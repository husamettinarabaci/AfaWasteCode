package WasteLibrary

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
)

//CustomersType
type CustomersType struct {
	Customers map[float64]float64
}

//ToByte
func (res CustomersType) ToByte() []byte {

	bBuf := bytes.Buffer{}
	gobEn := gob.NewEncoder(&bBuf)
	gobEn.Encode(res)
	return bBuf.Bytes()
}

//ToString Get JSON
func (res CustomersType) ToString() string {
	return base64.StdEncoding.EncodeToString(res.ToByte())

}

//Byte To CustomersType
func ByteToCustomersType(retByte []byte) CustomersType {
	var retVal CustomersType
	bBuf := bytes.Buffer{}
	bBuf.Write(retByte)
	gobDe := gob.NewDecoder(&bBuf)
	gobDe.Decode(&retVal)
	return retVal
}

//String To CustomersType
func StringToCustomersType(retStr string) CustomersType {
	bStr, _ := base64.StdEncoding.DecodeString(retStr)
	return ByteToCustomersType(bStr)
}
