package WasteLibrary

import (
	"encoding/json"
)

//PositionChangeType
type PositionChangeType struct {
	DeviceId  float64
	Latitude  float64
	Longitude float64
	ReadTime  string
}

//New
func (res *PositionChangeType) New() {
	res.DeviceId = 0
	res.Latitude = 0
	res.Longitude = 0
	res.ReadTime = GetTime()
}

//ToByte
func (res *PositionChangeType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//ToString Get JSON
func (res *PositionChangeType) ToString() string {
	return string(res.ToByte())

}

//Byte To PositionChangeType
func ByteToPositionChangeType(retByte []byte) PositionChangeType {
	var retVal PositionChangeType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To PositionChangeType
func StringToPositionChangeType(retStr string) PositionChangeType {
	return ByteToPositionChangeType([]byte(retStr))
}
