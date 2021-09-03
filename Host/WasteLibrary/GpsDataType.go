package WasteLibrary

import (
	"encoding/json"
)

//GpsDataType
type GpsDataType struct {
	Latitude   string
	Longitude  string
	LatitudeF  float64
	LongitudeF float64
}

//ToByte
func (res GpsDataType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res GpsDataType) ToString() string {
	return string(res.ToByte())
}

//Byte To GpsDataType
func ByteToGpsDataType(retByte []byte) GpsDataType {
	var retVal GpsDataType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To GpsDataType
func StringToGpsDataType(retStr string) GpsDataType {
	return ByteToGpsDataType([]byte(retStr))
}
