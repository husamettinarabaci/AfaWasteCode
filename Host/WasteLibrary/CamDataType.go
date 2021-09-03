package WasteLibrary

import (
	"encoding/json"
)

//CamDataType
type CamDataType struct {
	UID   string
	TagID string
}

//ToByte
func (res CamDataType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res CamDataType) ToString() string {
	return string(res.ToByte())
}

//Byte To CamDataType
func ByteToCamDataType(retByte []byte) CamDataType {
	var retVal CamDataType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To CamDataType
func StringToCamDataType(retStr string) CamDataType {
	return ByteToCamDataType([]byte(retStr))
}
