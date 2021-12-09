package WasteLibrary

import (
	"encoding/json"
	"fmt"
)

//CustomerUltDevicesType
type CustomerUltDevicesType struct {
	CustomerId float64
	Devices    map[string]float64
}

//New
func (res *CustomerUltDevicesType) New() {
	res.CustomerId = 1
	res.Devices = make(map[string]float64)
}

//GetByRedis
func (res *CustomerUltDevicesType) GetByRedis(dbIndex string) ResultType {
	resultVal := GetRedisForStoreApi(dbIndex, REDIS_CUSTOMER_ULT_DEVICES, res.ToIdString())
	if resultVal.Result == RESULT_OK {
		res.StringToType(resultVal.Retval.(string))
	} else {
		return resultVal
	}

	resultVal.Retval = res.ToString()
	return resultVal
}

//SaveToRedis
func (res *CustomerUltDevicesType) SaveToRedis() ResultType {
	resultVal := SaveRedisForStoreApi(REDIS_CUSTOMER_ULT_DEVICES, res.ToIdString(), res.ToString())
	return resultVal
}

//ToId String
func (res *CustomerUltDevicesType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToByte
func (res *CustomerUltDevicesType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//ToString Get JSON
func (res *CustomerUltDevicesType) ToString() string {
	return string(res.ToByte())

}

//ByteToType
func (res *CustomerUltDevicesType) ByteToType(retByte []byte) {
	res.New()
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *CustomerUltDevicesType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}
