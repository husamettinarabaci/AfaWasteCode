package WasteLibrary

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"fmt"
)

//CustomerTagsType
type CustomerTagsType struct {
	CustomerId float64
	Tags       map[float64]TagType
}

//ToId String
func (res CustomerTagsType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToByte
func (res CustomerTagsType) ToByte() []byte {

	bBuf := bytes.Buffer{}
	gobEn := gob.NewEncoder(&bBuf)
	gobEn.Encode(res)
	return bBuf.Bytes()
}

//ToString Get JSON
func (res CustomerTagsType) ToString() string {
	return base64.StdEncoding.EncodeToString(res.ToByte())

}

//Byte To CustomerTagsType
func ByteToCustomerTagsType(retByte []byte) CustomerTagsType {
	var retVal CustomerTagsType
	bBuf := bytes.Buffer{}
	bBuf.Write(retByte)
	gobDe := gob.NewDecoder(&bBuf)
	gobDe.Decode(&retVal)
	return retVal
}

//String To CustomerTagsType
func StringToCustomerTagsType(retStr string) CustomerTagsType {
	bStr, _ := base64.StdEncoding.DecodeString(retStr)
	return ByteToCustomerTagsType(bStr)
}
