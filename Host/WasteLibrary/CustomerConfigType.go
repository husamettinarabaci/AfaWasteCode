package WasteLibrary

import (
	"encoding/json"
	"fmt"
)

//CustomerConfigType
type CustomerConfigType struct {
	CustomerId      float64
	ArventoApp      string
	ArventoUserName string
	ArventoPin1     string
	ArventoPin2     string
	SystemProblem   string
	TruckStopTrace  string
	Active          string
	CreateTime      string
}

//New
func (res *CustomerConfigType) New() {
	res.CustomerId = 1
	res.ArventoApp = STATU_PASSIVE
	res.ArventoUserName = ""
	res.ArventoPin1 = ""
	res.ArventoPin2 = ""
	res.SystemProblem = STATU_PASSIVE
	res.TruckStopTrace = STATU_PASSIVE
	res.Active = STATU_ACTIVE
	res.CreateTime = GetTime()

}

//ToId String
func (res CustomerConfigType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToByte
func (res CustomerConfigType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//ToString Get JSON
func (res CustomerConfigType) ToString() string {
	return string(res.ToByte())

}

//Byte To CustomerConfigType
func ByteToCustomerConfigType(retByte []byte) CustomerConfigType {
	var retVal CustomerConfigType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To CustomerConfigType
func StringToCustomerConfigType(retStr string) CustomerConfigType {
	return ByteToCustomerConfigType([]byte(retStr))
}
