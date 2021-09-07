package WasteLibrary

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"fmt"
)

//LocalConfigType
type LocalConfigType struct {
	CustomerId float64
	Active     string
	CreateTime string
}

//ToId String
func (res LocalConfigType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToByte
func (res LocalConfigType) ToByte() []byte {

	bBuf := bytes.Buffer{}
	gobEn := gob.NewEncoder(&bBuf)
	gobEn.Encode(res)
	return bBuf.Bytes()
}

//ToString Get JSON
func (res LocalConfigType) ToString() string {
	return base64.StdEncoding.EncodeToString(res.ToByte())

}

//Byte To LocalConfigType
func ByteToLocalConfigType(retByte []byte) LocalConfigType {
	var retVal LocalConfigType
	bBuf := bytes.Buffer{}
	bBuf.Write(retByte)
	gobDe := gob.NewDecoder(&bBuf)
	gobDe.Decode(&retVal)
	return retVal
}

//String To LocalConfigType
func StringToLocalConfigType(retStr string) LocalConfigType {
	bStr, _ := base64.StdEncoding.DecodeString(retStr)
	return ByteToLocalConfigType(bStr)
}
