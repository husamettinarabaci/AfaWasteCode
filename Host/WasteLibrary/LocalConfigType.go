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

//New
func (res *LocalConfigType) New() {
	res.CustomerId = 1
	res.Active = STATU_ACTIVE
	res.CreateTime = GetTime()
}

//GetByRedis
func (res *LocalConfigType) GetByRedis() ResultType {
	var resultVal ResultType
	resultVal = GetRedisForStoreApi("0", REDIS_CUSTOMER_LOCALCONFIG, res.ToIdString())
	if resultVal.Result == RESULT_OK {
		res.StringToType(resultVal.Retval.(string))
	} else {
		return resultVal
	}

	resultVal.Retval = res.ToString()
	return resultVal
}

//SaveToRedis
func (res *LocalConfigType) SaveToRedis() ResultType {
	var resultVal ResultType
	resultVal = SaveRedisForStoreApi(REDIS_CUSTOMER_LOCALCONFIG, res.ToIdString(), res.ToString())
	return resultVal
}

//ToId String
func (res *LocalConfigType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToByte
func (res *LocalConfigType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res *LocalConfigType) ToString() string {
	return string(res.ToByte())

}

//ByteToType
func (res *LocalConfigType) ByteToType(retByte []byte) {
	res.New()
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *LocalConfigType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}
