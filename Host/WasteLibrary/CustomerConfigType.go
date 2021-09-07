package WasteLibrary

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"fmt"
)

//CustomerConfigType
type CustomerConfigType struct {
	CustomerId      float64
	ArventoApp      string
	ArventoUserName string
	ArventoPin1     string
	ArventoPin2     string
	SystemProblem   string
	TruckStopTrace  string
	Active          string
	CreateTime      string
}

//ToId String
func (res CustomerConfigType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToByte
func (res CustomerConfigType) ToByte() []byte {

	bBuf := bytes.Buffer{}
	gobEn := gob.NewEncoder(&bBuf)
	gobEn.Encode(res)
	return bBuf.Bytes()
}

//ToString Get JSON
func (res CustomerConfigType) ToString() string {
	return base64.StdEncoding.EncodeToString(res.ToByte())

}

//Byte To CustomerConfigType
func ByteToCustomerConfigType(retByte []byte) CustomerConfigType {
	var retVal CustomerConfigType
	bBuf := bytes.Buffer{}
	bBuf.Write(retByte)
	gobDe := gob.NewDecoder(&bBuf)
	gobDe.Decode(&retVal)
	return retVal
}

//String To CustomerConfigType
func StringToCustomerConfigType(retStr string) CustomerConfigType {
	bStr, _ := base64.StdEncoding.DecodeString(retStr)
	return ByteToCustomerConfigType(bStr)
}
