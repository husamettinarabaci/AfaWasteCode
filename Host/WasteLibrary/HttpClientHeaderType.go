package WasteLibrary

import (
	"encoding/base64"
	"fmt"
)

//HttpClientHeaderType
type HttpClientHeaderType struct {
	AppType      string
	DeviceNo     string
	OpType       string
	Time         string
	Repeat       string
	DeviceId     float64
	CustomerId   float64
	BaseDataType string
}

//ToCustomerId String
func (res HttpClientHeaderType) ToCustomerIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToDeviceId String
func (res HttpClientHeaderType) ToDeviceIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToByte
func (res HttpClientHeaderType) ToByte() []byte {
	return InterfaceToGobBytes(res)

}

//ToString Get JSON
func (res HttpClientHeaderType) ToString() string {
	return base64.StdEncoding.EncodeToString(res.ToByte())

}

//Byte To HttpClientHeaderType
func ByteToHttpClientHeaderType(retByte []byte) HttpClientHeaderType {
	var retVal HttpClientHeaderType
	GobBytestoInterface(retByte, &retVal)
	return retVal
}

//String To HttpClientHeaderType
func StringToHttpClientHeaderType(retStr string) HttpClientHeaderType {
	bStr, _ := base64.StdEncoding.DecodeString(retStr)
	return ByteToHttpClientHeaderType(bStr)
}
