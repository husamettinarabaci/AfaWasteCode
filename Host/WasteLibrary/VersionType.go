package WasteLibrary

import (
	"encoding/json"
)

//VersionType
type VersionType struct {
	RfidGpsAppVersion      string
	RfidThermAppVersion    string
	RfidTransferAppVersion string
	RfidCheckerAppVersion  string
	RfidCamAppVersion      string
	RfidReaderAppVersion   string
	RfidSystemAppVersion   string
	UltFirmwareVersion     string
	RecyWebAppVersion      string
	RecyMotorAppVersion    string
	RecyThermAppVersion    string
	RecyTransferAppVersion string
	RecyCheckerAppVersion  string
	RecyCamAppVersion      string
	RecyReaderAppVersion   string
	RecySystemAppVersion   string
}

//New
func (res *VersionType) New() {
	res.RfidGpsAppVersion = "1"
	res.RfidThermAppVersion = "1"
	res.RfidTransferAppVersion = "1"
	res.RfidCheckerAppVersion = "1"
	res.RfidCamAppVersion = "1"
	res.RfidReaderAppVersion = "1"
	res.RfidSystemAppVersion = "1"
	res.UltFirmwareVersion = "1"
	res.RecyWebAppVersion = "1"
	res.RecyMotorAppVersion = "1"
	res.RecyThermAppVersion = "1"
	res.RecyTransferAppVersion = "1"
	res.RecyCheckerAppVersion = "1"
	res.RecyCamAppVersion = "1"
	res.RecyReaderAppVersion = "1"
	res.RecySystemAppVersion = "1"
}

//ToByte
func (res VersionType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//ToString Get JSON
func (res VersionType) ToString() string {
	return string(res.ToByte())

}

//Byte To VersionType
func ByteToVersionType(retByte []byte) VersionType {
	var retVal VersionType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To VersionType
func StringToVersionType(retStr string) VersionType {
	return ByteToVersionType([]byte(retStr))
}
