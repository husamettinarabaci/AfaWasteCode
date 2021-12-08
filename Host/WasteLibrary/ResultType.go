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
func (res *ResultType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res *ResultType) ToString() string {
	return string(res.ToByte())

}

//ByteToType
func (res *ResultType) ByteToType(retByte []byte) {
	res.New()
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *ResultType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}
