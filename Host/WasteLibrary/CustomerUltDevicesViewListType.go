package WasteLibrary

import (
	"encoding/json"
	"fmt"
)

//CustomerUltDevicesViewListType
type CustomerUltDevicesViewListType struct {
	CustomerId float64
	Devices    map[string]UltDeviceViewType
}

//New
func (res *CustomerUltDevicesViewListType) New() {
	res.CustomerId = 1
	res.Devices = make(map[string]UltDeviceViewType)
}

//ToId String
func (res *CustomerUltDevicesViewListType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToByte
func (res *CustomerUltDevicesViewListType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//ToString Get JSON
func (res *CustomerUltDevicesViewListType) ToString() string {
	return string(res.ToByte())

}

//Byte To CustomerUltDevicesViewListType
func ByteToCustomerUltDevicesViewListType(retByte []byte) CustomerUltDevicesViewListType {
	var retVal CustomerUltDevicesViewListType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To CustomerUltDevicesViewListType
func StringToCustomerUltDevicesViewListType(retStr string) CustomerUltDevicesViewListType {
	return ByteToCustomerUltDevicesViewListType([]byte(retStr))
}
