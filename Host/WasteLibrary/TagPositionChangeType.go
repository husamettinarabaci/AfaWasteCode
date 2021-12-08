package WasteLibrary

import (
	"encoding/json"
	"fmt"
)

//TagPositionChangeType
type TagPositionChangeType struct {
	TagId           float64
	PositionChanges []PositionChangeType
}

//New
func (res *TagPositionChangeType) New() {
	res.TagId = 0
	res.PositionChanges = []PositionChangeType{}
}

//GetByRedis
func (res *TagPositionChangeType) GetByRedis(dbIndex string) ResultType {

	var resultVal ResultType
	resultVal = GetRedisForStoreApi(dbIndex, REDIS_TAG_POSITION_CHANGE, res.ToIdString())
	if resultVal.Result == RESULT_OK {
		res.StringToType(resultVal.Retval.(string))
	} else {
		return resultVal
	}

	resultVal.Retval = res.ToString()
	return resultVal
}

//SaveToRedis
func (res *TagPositionChangeType) SaveToRedis() ResultType {
	var resultVal ResultType
	resultVal = SaveRedisForStoreApi(REDIS_TAG_POSITION_CHANGE, res.ToIdString(), res.ToString())
	return resultVal
}

//ToId String
func (res *TagPositionChangeType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.TagId)
}

//ToByte
func (res *TagPositionChangeType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res *TagPositionChangeType) ToString() string {
	return string(res.ToByte())

}

//ByteToType
func (res *TagPositionChangeType) ByteToType(retByte []byte) {
	res.New()
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *TagPositionChangeType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}
