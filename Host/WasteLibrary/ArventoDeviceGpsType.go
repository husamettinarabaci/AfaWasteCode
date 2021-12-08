package WasteLibrary

import (
	"encoding/json"
)

//ArventoDeviceGpsType
type ArventoDeviceGpsType struct {
	Latitude  float64
	Longitude float64
	Speed     float64
	GpsTime   string
}

//New
func (res *ArventoDeviceGpsType) New() {
	res.Latitude = 0
	res.Longitude = 0
	res.Speed = -1
	res.GpsTime = GetTime()

}

//ToByte
func (res *ArventoDeviceGpsType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//ToString Get JSON
func (res *ArventoDeviceGpsType) ToString() string {
	return string(res.ToByte())

}

//ByteToType
func (res *ArventoDeviceGpsType) ByteToType(retByte []byte) {
	res.New()
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *ArventoDeviceGpsType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}
