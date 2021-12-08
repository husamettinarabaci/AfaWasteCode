package WasteLibrary

import (
	"encoding/json"
	"fmt"
)

//CustomerConfigType
type CustomerConfigType struct {
	CustomerId      float64
	ArventoApp      string
	ArventoUserName string
	ArventoPin1     string
	ArventoPin2     string
	SystemProblem   string
	TruckStopTrace  string
	Active          string
	CreateTime      string
}

//New
func (res *CustomerConfigType) New() {
	res.CustomerId = 1
	res.ArventoApp = STATU_PASSIVE
	res.ArventoUserName = ""
	res.ArventoPin1 = ""
	res.ArventoPin2 = ""
	res.SystemProblem = STATU_PASSIVE
	res.TruckStopTrace = STATU_PASSIVE
	res.Active = STATU_ACTIVE
	res.CreateTime = GetTime()

}

//GetByRedis
func (res *CustomerConfigType) GetByRedis() ResultType {
	var resultVal ResultType
	resultVal = GetRedisForStoreApi("0", REDIS_CUSTOMER_CUSTOMERCONFIG, res.ToIdString())
	if resultVal.Result == RESULT_OK {
		res.StringToType(resultVal.Retval.(string))
	} else {
		return resultVal
	}

	resultVal.Retval = res.ToString()
	return resultVal
}

//SaveToRedis
func (res *CustomerConfigType) SaveToRedis() ResultType {
	var resultVal ResultType
	resultVal = SaveRedisForStoreApi(REDIS_CUSTOMER_CUSTOMERCONFIG, res.ToIdString(), res.ToString())
	return resultVal
}

//ToId String
func (res *CustomerConfigType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToByte
func (res *CustomerConfigType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//ToString Get JSON
func (res *CustomerConfigType) ToString() string {
	return string(res.ToByte())

}

//ByteToType
func (res *CustomerConfigType) ByteToType(retByte []byte) {
	retVal.New()
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *CustomerConfigType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}
