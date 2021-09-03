package WasteLibrary

import (
	"encoding/json"
)

//ThermDataType
type ThermDataType struct {
	Therm string
}

//ToByte
func (res ThermDataType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res ThermDataType) ToString() string {
	return string(res.ToByte())
}

//Byte To ThermDataType
func ByteToThermDataType(retByte []byte) ThermDataType {
	var retVal ThermDataType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To ThermDataType
func StringToThermDataType(retStr string) ThermDataType {
	return ByteToThermDataType([]byte(retStr))
}
