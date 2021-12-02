package WasteLibrary

import (
	"encoding/json"
)

//ReadDeviceType
type ReadDeviceType struct {
	DeviceId float64
	ReadTime string
}

//New
func (res *ReadDeviceType) New() {
	res.DeviceId = 0
	res.ReadTime = GetTime()
}

//ToByte
func (res *ReadDeviceType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//ToString Get JSON
func (res *ReadDeviceType) ToString() string {
	return string(res.ToByte())

}

//Byte To ReadDeviceType
func ByteToReadDeviceType(retByte []byte) ReadDeviceType {
	var retVal ReadDeviceType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To ReadDeviceType
func StringToReadDeviceType(retStr string) ReadDeviceType {
	return ByteToReadDeviceType([]byte(retStr))
}
