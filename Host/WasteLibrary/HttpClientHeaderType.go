package WasteLibrary

import (
	"encoding/json"
	"fmt"
)

//HttpClientHeaderType
type HttpClientHeaderType struct {
	AppType    string
	DeviceNo   string
	DeviceId   float64
	CustomerId float64
	Time       string
	Repeat     string
	DeviceType string
	ReaderType string
	DataType   string
}

//New
func (res *HttpClientHeaderType) New() {
	res.AppType = APPTYPE_NONE
	res.DeviceNo = ""
	res.DeviceId = 0
	res.CustomerId = 1
	res.Time = GetTime()
	res.Repeat = STATU_PASSIVE
	res.DeviceType = DEVICETYPE_NONE
	res.ReaderType = READERTYPE_NONE
	res.DataType = DATATYPE_NONE
}

//ToCustomerId String
func (res *HttpClientHeaderType) ToCustomerIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToDeviceId String
func (res *HttpClientHeaderType) ToDeviceIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToByte
func (res *HttpClientHeaderType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res *HttpClientHeaderType) ToString() string {
	return string(res.ToByte())

}

//ByteToType
func (res *HttpClientHeaderType) ByteToType(retByte []byte) {
	res.New()
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *HttpClientHeaderType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}
