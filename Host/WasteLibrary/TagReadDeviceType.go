package WasteLibrary

import (
	"encoding/json"
	"fmt"
)

//TagReadDeviceType
type TagReadDeviceType struct {
	TagId       float64
	ReadDevices []ReadDeviceType
}

//New
func (res *TagReadDeviceType) New() {
	res.TagId = 0
	res.ReadDevices = []ReadDeviceType{}
}

//GetByRedis
func (res *TagReadDeviceType) GetByRedis(dbIndex string) ResultType {

	var resultVal ResultType
	resultVal = GetRedisForStoreApi(dbIndex, REDIS_TAG_READDEVICES, res.ToIdString())
	if resultVal.Result == RESULT_OK {
		res.StringToType(resultVal.Retval.(string))
	} else {
		return resultVal
	}

	resultVal.Retval = res.ToString()
	return resultVal
}

//SaveToRedis
func (res *TagReadDeviceType) SaveToRedis() ResultType {
	var resultVal ResultType
	resultVal = SaveRedisForStoreApi(REDIS_TAG_READDEVICES, res.ToIdString(), res.ToString())
	return resultVal
}

//ToId String
func (res *TagReadDeviceType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.TagId)
}

//ToByte
func (res *TagReadDeviceType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//ToString Get JSON
func (res *TagReadDeviceType) ToString() string {
	return string(res.ToByte())

}

//Byte To TagReadDeviceType
func ByteToTagReadDeviceType(retByte []byte) TagReadDeviceType {
	var retVal TagReadDeviceType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To TagReadDeviceType
func StringToTagReadDeviceType(retStr string) TagReadDeviceType {
	return ByteToTagReadDeviceType([]byte(retStr))
}

//ByteToType
func (res *TagReadDeviceType) ByteToType(retByte []byte) {
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *TagReadDeviceType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}
