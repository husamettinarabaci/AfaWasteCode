package WasteLibrary

import (
	"encoding/json"
	"fmt"
)

//DeviceType
type DeviceType struct {
	DeviceId              float64
	CustomerId            float64
	DeviceName            string
	DeviceType            string
	SerialNumber          string
	ReaderAppStatus       string
	ReaderAppLastOkTime   string
	ReaderConnStatus      string
	ReaderConnLastOkTime  string
	ReaderStatus          string
	ReaderLastOkTime      string
	CamAppStatus          string
	CamAppLastOkTime      string
	CamConnStatus         string
	CamConnLastOkTime     string
	CamStatus             string
	CamLastOkTime         string
	GpsAppStatus          string
	GpsAppLastOkTime      string
	GpsConnStatus         string
	GpsConnLastOkTime     string
	GpsStatus             string
	GpsLastOkTime         string
	ThermAppStatus        string
	ThermAppLastOkTime    string
	TransferAppStatus     string
	TransferAppLastOkTime string
	AliveStatus           string
	AliveLastOkTime       string
	ContactStatus         string
	ContactLastOkTime     string
	Therm                 string
	ThermStatus           string
	Battery               string
	BatteryStatus         string
	Latitude              float64
	Longitude             float64
	Speed                 float64
	UltRange              float64
	UltStatus             string
	DeviceStatus          string
	AlarmStatus           string
	TotalGlassCount       float64
	TotalPlasticCount     float64
	TotalMetalCount       float64
	UltTime               string
	AlarmTime             string
	AlarmType             string
	Alarm                 string
	RecyTime              string
	Active                string
	ThermTime             string
	BatteryTime           string
	GpsTime               string
	StatusTime            string
	CreateTime            string
}

//ToId String
func (res DeviceType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToCustomerId String
func (res DeviceType) ToCustomerIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToByte
func (res DeviceType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res DeviceType) ToString() string {
	return string(res.ToByte())

}

//Byte To DeviceType
func ByteToDeviceType(retByte []byte) DeviceType {
	var retVal DeviceType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To DeviceType
func StringToDeviceType(retStr string) DeviceType {
	return ByteToDeviceType([]byte(retStr))
}
