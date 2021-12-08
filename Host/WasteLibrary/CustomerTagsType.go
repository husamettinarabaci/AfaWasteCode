package WasteLibrary

import (
	"encoding/json"
	"fmt"
)

//CustomerTagsType
type CustomerTagsType struct {
	CustomerId float64
	Tags       map[string]float64
}

//New
func (res *CustomerTagsType) New() {
	res.CustomerId = 1
	res.Tags = make(map[string]float64)
}

//GetByRedis
func (res *CustomerTagsType) GetByRedis(dbIndex string) ResultType {
	var resultVal ResultType
	resultVal = GetRedisForStoreApi(dbIndex, REDIS_CUSTOMER_TAGS, res.ToIdString())
	if resultVal.Result == RESULT_OK {
		res.StringToType(resultVal.Retval.(string))
	} else {
		return resultVal
	}

	resultVal.Retval = res.ToString()
	return resultVal
}

//SaveToRedis
func (res *CustomerTagsType) SaveToRedis() ResultType {
	var resultVal ResultType
	resultVal = SaveRedisForStoreApi(REDIS_CUSTOMER_TAGS, res.ToIdString(), res.ToString())
	return resultVal
}

//ToId String
func (res *CustomerTagsType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToByte
func (res *CustomerTagsType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//ToString Get JSON
func (res *CustomerTagsType) ToString() string {
	return string(res.ToByte())

}

//ByteToType
func (res *CustomerTagsType) ByteToType(retByte []byte) {
	res.New()
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *CustomerTagsType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}
