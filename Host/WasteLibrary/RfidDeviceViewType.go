package WasteLibrary

import (
	"encoding/json"
	"fmt"
)

//RfidDeviceViewType
type RfidDeviceViewType struct {
	DeviceId  float64
	PlateNo   string
	Latitude  float64
	Longitude float64
}

//New
func (res *RfidDeviceViewType) New() {
	res.DeviceId = 0
	res.PlateNo = ""
	res.Latitude = 0
	res.Longitude = 0
}

//ToId String
func (res *RfidDeviceViewType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToTagId String
func (res *RfidDeviceViewType) ToTagIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToByte
func (res *RfidDeviceViewType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//ToString Get JSON
func (res *RfidDeviceViewType) ToString() string {
	return string(res.ToByte())

}

//Byte To RfidDeviceViewType
func ByteToRfidDeviceViewType(retByte []byte) RfidDeviceViewType {
	var retVal RfidDeviceViewType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To RfidDeviceViewType
func StringToRfidDeviceViewType(retStr string) RfidDeviceViewType {
	return ByteToRfidDeviceViewType([]byte(retStr))
}
