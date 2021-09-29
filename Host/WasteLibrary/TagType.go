package WasteLibrary

import (
	"encoding/base64"
	"fmt"
)

//TagType
type TagType struct {
	TagID       float64
	CustomerId  float64
	DeviceId    float64
	UID         string
	Epc         string
	ContainerNo string
	Latitude    float64
	Longitude   float64
	Statu       string
	ImageStatu  string
	Active      string
	ReadTime    string
	CheckTime   string
	CreateTime  string
}

//ToId String
func (res TagType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.TagID)
}

//ToCustomerId String
func (res TagType) ToCustomerIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToDeviceId String
func (res TagType) ToDeviceIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToByte
func (res TagType) ToByte() []byte {
	return InterfaceToGobBytes(res)

}

//ToString Get JSON
func (res TagType) ToString() string {
	return base64.StdEncoding.EncodeToString(res.ToByte())

}

//Byte To TagType
func ByteToTagType(retByte []byte) TagType {
	var retVal TagType
	GobBytestoInterface(retByte, &retVal)
	return retVal
}

//String To TagType
func StringToTagType(retStr string) TagType {
	bStr, _ := base64.StdEncoding.DecodeString(retStr)
	return ByteToTagType(bStr)
}
