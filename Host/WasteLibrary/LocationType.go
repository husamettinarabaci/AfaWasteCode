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

//ToByte
func (res *LocationType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//ToString Get JSON
func (res *LocationType) ToString() string {
	return string(res.ToByte())

}

//Byte To LocationType
func ByteToLocationType(retByte []byte) LocationType {
	var retVal LocationType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To LocationType
func StringToLocationType(retStr string) LocationType {
	return ByteToLocationType([]byte(retStr))
}
