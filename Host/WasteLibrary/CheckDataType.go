package WasteLibrary

import (
	"encoding/json"
)

//CheckDataType
type CheckDataType struct {
	ReaderAppStatus   string
	ReaderConnStatus  string
	ReaderStatus      string
	CamAppStatus      string
	CamConnStatus     string
	CamStatus         string
	GpsAppStatus      string
	GpsConnStatus     string
	GpsStatus         string
	ThermAppStatus    string
	TransferAppStatus string
	AliveStatus       string
	ContactStatus     string
}

//ToByte
func (res CheckDataType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res CheckDataType) ToString() string {
	return string(res.ToByte())
}

//Byte To CheckDataType
func ByteToCheckDataType(retByte []byte) CheckDataType {
	var retVal CheckDataType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To CheckDataType
func StringToCheckDataType(retStr string) CheckDataType {
	return ByteToCheckDataType([]byte(retStr))
}
