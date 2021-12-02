package WasteLibrary

import (
	"encoding/json"
	"fmt"
)

//RedisDbDateType
type RedisDbDateType struct {
	LastDay  int
	DayDates [31]string
}

//New
func (res *RedisDbDateType) New() {
	res.LastDay = 0
	res.DayDates = [31]string{}
}

//GetByRedis
func (res *RedisDbDateType) GetByRedis() ResultType {
	var resultVal ResultType
	resultVal = GetRedisForStoreApi("0", REDIS_DB_DATE, REDIS_DB_DATE)
	if resultVal.Result == RESULT_OK {
		res.StringToType(resultVal.Retval.(string))
	} else {
		return resultVal
	}

	resultVal.Retval = res.ToString()
	return resultVal
}

//SaveToRedis
func (res *RedisDbDateType) SaveToRedis() ResultType {
	var resultVal ResultType
	resultVal = SaveRedisForStoreApi(REDIS_DB_DATE, REDIS_DB_DATE, res.ToString())
	return resultVal
}

//ToLastDay String
func (res *RedisDbDateType) ToLastDayString() string {
	return fmt.Sprintf("%d", res.LastDay)
}

//ToByte
func (res *RedisDbDateType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//ToString Get JSON
func (res *RedisDbDateType) ToString() string {
	return string(res.ToByte())

}

//Byte To RedisDbDateType
func ByteToRedisDbDateType(retByte []byte) RedisDbDateType {
	var retVal RedisDbDateType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To RedisDbDateType
func StringToRedisDbDateType(retStr string) RedisDbDateType {
	return ByteToRedisDbDateType([]byte(retStr))
}

//ByteToType
func (res *RedisDbDateType) ByteToType(retByte []byte) {
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *RedisDbDateType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}
