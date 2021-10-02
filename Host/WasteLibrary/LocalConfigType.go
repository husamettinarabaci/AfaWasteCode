package WasteLibrary

import (
	"encoding/json"
	"fmt"
)

//LocalConfigType
type LocalConfigType struct {
	CustomerId float64
	Active     string
	CreateTime string
}

//ToId String
func (res LocalConfigType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToByte
func (res LocalConfigType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res LocalConfigType) ToString() string {
	return string(res.ToByte())

}

//Byte To LocalConfigType
func ByteToLocalConfigType(retByte []byte) LocalConfigType {
	var retVal LocalConfigType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To LocalConfigType
func StringToLocalConfigType(retStr string) LocalConfigType {
	return ByteToLocalConfigType([]byte(retStr))
}
