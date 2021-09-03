package WasteLibrary

import (
	"encoding/json"
	"fmt"
)

//DeviceType
type DeviceType struct {
	DeviceId     float64
	CustomerId   float64
	DeviceName   string
	DeviceType   string
	SerialNumber string
	Active       string
	CreateTime   string
}

//ToId String
func (res DeviceType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToCustomerId String
func (res DeviceType) ToCustomerIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToByte
func (res DeviceType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res DeviceType) ToString() string {
	return string(res.ToByte())
}

//Byte To DeviceType
func ByteToDeviceType(retByte []byte) DeviceType {
	var retVal DeviceType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To DeviceType
func StringToDeviceType(retStr string) DeviceType {
	return ByteToDeviceType([]byte(retStr))
}
