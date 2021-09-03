package WasteLibrary

import (
	"encoding/json"
)

//ReaderDataType
type ReaderDataType struct {
	UID   string
	TagID string
}

//ToByte
func (res ReaderDataType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res ReaderDataType) ToString() string {
	return string(res.ToByte())
}

//Byte To ReaderDataType
func ByteToReaderDataType(retByte []byte) ReaderDataType {
	var retVal ReaderDataType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To ReaderDataType
func StringToReaderDataType(retStr string) ReaderDataType {
	return ByteToReaderDataType([]byte(retStr))
}
