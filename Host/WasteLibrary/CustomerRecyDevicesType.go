package WasteLibrary

import (
	"encoding/json"
	"fmt"
)

//CustomerRecyDevicesType
type CustomerRecyDevicesType struct {
	CustomerId float64
	Devices    map[string]float64
}

//New
func (res *CustomerRecyDevicesType) New() {
	res.CustomerId = 1
	res.Devices = make(map[string]float64)
}

//GetByRedis
func (res *CustomerRecyDevicesType) GetByRedis(dbIndex string) ResultType {
	resultVal := GetRedisForStoreApi(dbIndex, REDIS_CUSTOMER_RECY_DEVICES, res.ToIdString())
	if resultVal.Result == RESULT_OK {
		res.StringToType(resultVal.Retval.(string))
	} else {
		return resultVal
	}

	resultVal.Retval = res.ToString()
	return resultVal
}

//SaveToRedis
func (res *CustomerRecyDevicesType) SaveToRedis() ResultType {
	resultVal := SaveRedisForStoreApi(REDIS_CUSTOMER_RECY_DEVICES, res.ToIdString(), res.ToString())
	return resultVal
}

//ToId String
func (res *CustomerRecyDevicesType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToByte
func (res *CustomerRecyDevicesType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//ToString Get JSON
func (res *CustomerRecyDevicesType) ToString() string {
	return string(res.ToByte())

}

//ByteToType
func (res *CustomerRecyDevicesType) ByteToType(retByte []byte) {
	res.New()
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *CustomerRecyDevicesType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}
