package WasteLibrary

import (
	"encoding/json"
	"fmt"
)

//CustomerUsersType
type CustomerUsersType struct {
	CustomerId float64
	Users      map[string]float64
}

//New
func (res *CustomerUsersType) New() {
	res.CustomerId = 1
	res.Users = make(map[string]float64)
}

//GetByRedis
func (res *CustomerUsersType) GetByRedis() ResultType {
	resultVal := GetRedisForStoreApi("0", REDIS_CUSTOMER_USERS, res.ToIdString())
	if resultVal.Result == RESULT_OK {
		res.StringToType(resultVal.Retval.(string))
	} else {
		return resultVal
	}

	resultVal.Retval = res.ToString()
	return resultVal
}

//SaveToRedis
func (res *CustomerUsersType) SaveToRedis() ResultType {
	resultVal := SaveRedisForStoreApi(REDIS_CUSTOMER_USERS, res.ToIdString(), res.ToString())
	return resultVal
}

//ToId String
func (res *CustomerUsersType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToByte
func (res *CustomerUsersType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//ToString Get JSON
func (res *CustomerUsersType) ToString() string {
	return string(res.ToByte())

}

//ByteToType
func (res *CustomerUsersType) ByteToType(retByte []byte) {
	res.New()
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *CustomerUsersType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}
