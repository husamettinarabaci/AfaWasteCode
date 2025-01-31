package WasteLibrary

import (
	"encoding/json"
	"fmt"
)

//UltDeviceViewType
type RecyDeviceViewType struct {
	DeviceId    float64
	ContainerNo string
	Latitude    float64
	Longitude   float64
}

//New
func (res *RecyDeviceViewType) New() {
	res.DeviceId = 0
	res.ContainerNo = ""
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

//ByteToType
func (res *RecyDeviceViewType) ByteToType(retByte []byte) {
	res.New()
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *RecyDeviceViewType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}
