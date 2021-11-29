package WasteLibrary

import (
	"encoding/json"
	"fmt"
)

//RecyDeviceViewType
type RecyDeviceViewType struct {
	DeviceId  float64
	Latitude  float64
	Longitude float64
}

//New
func (res *RecyDeviceViewType) New() {
	res.DeviceId = 0
	res.Latitude = 0
	res.Longitude = 0
}

//ToId String
func (res *RecyDeviceViewType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToTagId String
func (res *RecyDeviceViewType) ToTagIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToByte
func (res *RecyDeviceViewType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//ToString Get JSON
func (res *RecyDeviceViewType) ToString() string {
	return string(res.ToByte())

}

//Byte To RecyDeviceViewType
func ByteToRecyDeviceViewType(retByte []byte) RecyDeviceViewType {
	var retVal RecyDeviceViewType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To RecyDeviceViewType
func StringToRecyDeviceViewType(retStr string) RecyDeviceViewType {
	return ByteToRecyDeviceViewType([]byte(retStr))
}
