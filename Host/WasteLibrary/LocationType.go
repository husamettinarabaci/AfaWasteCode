package WasteLibrary

import (
	"encoding/json"
)

//LocationType
type LocationType struct {
	LocationName string
	Latitude     float64
	Longitude    float64
	ZoneRadius   float64
}

//New
func (res *LocationType) New() {
	res.LocationName = ""
	res.Latitude = 0
	res.Longitude = 0
	res.ZoneRadius = 0
}

//ToByte
func (res *LocationType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//ToString Get JSON
func (res *LocationType) ToString() string {
	return string(res.ToByte())

}

//ByteToType
func (res *LocationType) ByteToType(retByte []byte) {
	res.New()
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *LocationType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}
