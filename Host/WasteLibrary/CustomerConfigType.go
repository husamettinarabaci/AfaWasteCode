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
