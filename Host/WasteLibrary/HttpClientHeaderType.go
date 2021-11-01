package WasteLibrary

import (
	"encoding/json"
	"fmt"
)

//HttpClientHeaderType
type HttpClientHeaderType struct {
	AppType      string
	DeviceNo     string
	OpType       string
	Time         string
	Repeat       string
	DeviceType   string
	DeviceId     float64
	CustomerId   float64
	BaseDataType string
	Token        string
}

//New
func (res *HttpClientHeaderType) New() {
	res.AppType = APPTYPE_NONE
	res.DeviceNo = ""
	res.OpType = OPTYPE_NONE
	res.Time = GetTime()
	res.Repeat = STATU_PASSIVE
	res.DeviceType = DEVICE_TYPE_NONE
	res.DeviceId = 0
	res.CustomerId = 0
	res.BaseDataType = BASETYPE_NONE
	res.Token = ""
}

//ToCustomerId String
func (res HttpClientHeaderType) ToCustomerIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToDeviceId String
func (res HttpClientHeaderType) ToDeviceIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToByte
func (res HttpClientHeaderType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res HttpClientHeaderType) ToString() string {
	return string(res.ToByte())

}

//Byte To HttpClientHeaderType
func ByteToHttpClientHeaderType(retByte []byte) HttpClientHeaderType {
	var retVal HttpClientHeaderType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To HttpClientHeaderType
func StringToHttpClientHeaderType(retStr string) HttpClientHeaderType {
	return ByteToHttpClientHeaderType([]byte(retStr))
}
