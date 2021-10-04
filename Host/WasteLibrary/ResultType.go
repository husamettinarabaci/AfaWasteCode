package WasteLibrary

import (
	"encoding/json"
)

//ResultType
type ResultType struct {
	Result string
	Retval interface{}
}

//ToByte
func (res ResultType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res ResultType) ToString() string {
	return string(res.ToByte())

}

//Byte To ResultType
func ByteToResultType(retByte []byte) ResultType {
	var retVal ResultType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To ResultType
func StringToResultType(retStr string) ResultType {
	return ByteToResultType([]byte(retStr))
}
