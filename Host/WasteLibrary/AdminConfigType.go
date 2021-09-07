package WasteLibrary

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"fmt"
)

//AdminConfigType
type AdminConfigType struct {
	CustomerId float64
	Active     string
	CreateTime string
}

//ToId String
func (res AdminConfigType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToByte
func (res AdminConfigType) ToByte() []byte {

	bBuf := bytes.Buffer{}
	gobEn := gob.NewEncoder(&bBuf)
	gobEn.Encode(res)
	return bBuf.Bytes()
}

//ToString Get JSON
func (res AdminConfigType) ToString() string {
	return base64.StdEncoding.EncodeToString(res.ToByte())

}

//Byte To AdminConfigType
func ByteToAdminConfigType(retByte []byte) AdminConfigType {
	var retVal AdminConfigType
	bBuf := bytes.Buffer{}
	bBuf.Write(retByte)
	gobDe := gob.NewDecoder(&bBuf)
	gobDe.Decode(&retVal)
	return retVal
}

//String To AdminConfigType
func StringToAdminConfigType(retStr string) AdminConfigType {
	bStr, _ := base64.StdEncoding.DecodeString(retStr)
	return ByteToAdminConfigType(bStr)
}
