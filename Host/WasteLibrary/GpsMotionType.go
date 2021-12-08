package WasteLibrary

import (
	"encoding/json"
)

//GpsMotionType
type GpsMotionType struct {
	Angle      string
	Latitude   float64
	Longitude  float64
	MotionType string
	Speed      float64
	GpsTime    string
}

//New
func (res *GpsMotionType) New() {
	res.Angle = ANGLETYPE_NONE
	res.Latitude = 0
	res.Longitude = 0
	res.MotionType = MOTIONTYPE_NONE
	res.Speed = 0
	res.GpsTime = GetTime()
}

//ToByte
func (res *GpsMotionType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//ToString Get JSON
func (res *GpsMotionType) ToString() string {
	return string(res.ToByte())

}

//ByteToType
func (res *GpsMotionType) ByteToType(retByte []byte) {
	res.New()
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *GpsMotionType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}
