package WasteLibrary

import (
	"encoding/json"
	"fmt"
)

//SystemDateType
type SystemDateType struct {
	LastDay  int
	DayDates [31]string
}

//New
func (res *SystemDateType) New() {
	res.LastDay = 0
	res.DayDates = [31]string{}
}

//GetByRedis
func (res *SystemDateType) GetByRedis() ResultType {
	var resultVal ResultType
	resultVal = GetRedisForStoreApi("0", REDIS_SYSTEM_DATE, REDIS_SYSTEM_DATE)
	if resultVal.Result == RESULT_OK {
		res.StringToType(resultVal.Retval.(string))
	} else {
		return resultVal
	}

	resultVal.Retval = res.ToString()
	return resultVal
}

//SaveToRedis
func (res *SystemDateType) SaveToRedis() ResultType {
	var resultVal ResultType
	resultVal = SaveRedisForStoreApi(REDIS_SYSTEM_DATE, REDIS_SYSTEM_DATE, res.ToString())
	return resultVal
}

//ToLastDay String
func (res *SystemDateType) ToLastDayString() string {
	return fmt.Sprintf("%d", res.LastDay)
}

//ToByte
func (res *SystemDateType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//ToString Get JSON
func (res *SystemDateType) ToString() string {
	return string(res.ToByte())

}

//Byte To SystemDateType
func ByteToSystemDateType(retByte []byte) SystemDateType {
	var retVal SystemDateType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To SystemDateType
func StringToSystemDateType(retStr string) SystemDateType {
	return ByteToSystemDateType([]byte(retStr))
}

//ByteToType
func (res *SystemDateType) ByteToType(retByte []byte) {
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *SystemDateType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}
