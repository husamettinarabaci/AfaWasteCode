package WasteLibrary

import (
	"encoding/base64"
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
	return InterfaceToGobBytes(res)

}

//ToString Get JSON
func (res CustomerDevicesType) ToString() string {
	return base64.StdEncoding.EncodeToString(res.ToByte())

}

//Byte To CustomerDevicesType
func ByteToCustomerDevicesType(retByte []byte) CustomerDevicesType {
	var retVal CustomerDevicesType
	GobBytestoInterface(retByte, &retVal)
	return retVal
}

//String To CustomerDevicesType
func StringToCustomerDevicesType(retStr string) CustomerDevicesType {
	bStr, _ := base64.StdEncoding.DecodeString(retStr)
	return ByteToCustomerDevicesType(bStr)
}
