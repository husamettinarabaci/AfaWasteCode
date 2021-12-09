package WasteLibrary

import (
	"encoding/json"
	"fmt"
)

//CustomerUltDevicesViewListType
type CustomerUltDevicesViewListType struct {
	CustomerId float64
	Devices    map[string]UltDeviceViewType
}

//New
func (res *CustomerUltDevicesViewListType) New() {
	res.CustomerId = 1
	res.Devices = make(map[string]UltDeviceViewType)
}

//GetByRedis
func (res *CustomerUltDevicesViewListType) GetByRedis(dbIndex string) ResultType {
	resultVal := GetRedisForStoreApi(dbIndex, REDIS_CUSTOMER_ULT_DEVICEVIEWS, res.ToIdString())
	if resultVal.Result == RESULT_OK {
		res.StringToType(resultVal.Retval.(string))
	} else {
		return resultVal
	}

	resultVal.Retval = res.ToString()
	return resultVal
}

//GetByRedisByReel
func (res *CustomerUltDevicesViewListType) GetByRedisByReel(dbIndex string) ResultType {
	resultVal := GetRedisWODbForStoreApi(dbIndex, REDIS_CUSTOMER_ULT_DEVICEVIEWS_REEL, REDIS_CUSTOMER_ULT_DEVICEVIEWS, res.ToIdString())
	if resultVal.Result == RESULT_OK {
		res.StringToType(resultVal.Retval.(string))
	} else {
		return resultVal
	}

	resultVal.Retval = res.ToString()
	return resultVal
}

//SaveToRedis
func (res *CustomerUltDevicesViewListType) SaveToRedis() ResultType {
	resultVal := SaveRedisForStoreApi(REDIS_CUSTOMER_ULT_DEVICEVIEWS, res.ToIdString(), res.ToString())
	return resultVal
}

//SaveToRedisWODb
func (res *CustomerUltDevicesViewListType) SaveToRedisWODb() ResultType {
	resultVal := SaveRedisWODbForStoreApi(REDIS_CUSTOMER_ULT_DEVICEVIEWS_REEL, res.ToIdString(), res.ToString())
	return resultVal
}

//ToId String
func (res *CustomerUltDevicesViewListType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToByte
func (res *CustomerUltDevicesViewListType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//ToString Get JSON
func (res *CustomerUltDevicesViewListType) ToString() string {
	return string(res.ToByte())

}

//ByteToType
func (res *CustomerUltDevicesViewListType) ByteToType(retByte []byte) {
	res.New()
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *CustomerUltDevicesViewListType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}
