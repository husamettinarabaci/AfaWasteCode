package WasteLibrary

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"fmt"
)

//CustomerDevicesType
type CustomerDevicesType struct {
	CustomerId float64
	Devices    map[float64]float64
}

//ToId String
func (res CustomerDevicesType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToByte
func (res CustomerDevicesType) ToByte() []byte {

	bBuf := bytes.Buffer{}
	gobEn := gob.NewEncoder(&bBuf)
	gobEn.Encode(res)
	return bBuf.Bytes()
}

//ToString Get JSON
func (res CustomerDevicesType) ToString() string {
	return base64.StdEncoding.EncodeToString(res.ToByte())

}

//Byte To CustomerDevicesType
func ByteToCustomerDevicesType(retByte []byte) CustomerDevicesType {
	var retVal CustomerDevicesType
	bBuf := bytes.Buffer{}
	bBuf.Write(retByte)
	gobDe := gob.NewDecoder(&bBuf)
	gobDe.Decode(&retVal)
	return retVal
}

//String To CustomerDevicesType
func StringToCustomerDevicesType(retStr string) CustomerDevicesType {
	bStr, _ := base64.StdEncoding.DecodeString(retStr)
	return ByteToCustomerDevicesType(bStr)
}
