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

	resultVal := GetRedisForStoreApi(dbIndex, REDIS_TAG_READDEVICE, res.ToIdString())
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
	resultVal := SaveRedisForStoreApi(REDIS_TAG_READDEVICE, res.ToIdString(), res.ToString())
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

//ByteToType
func (res *TagReadDeviceType) ByteToType(retByte []byte) {
	res.New()
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *TagReadDeviceType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}
