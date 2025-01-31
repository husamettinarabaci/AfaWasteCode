package WasteLibrary

import (
	"encoding/json"
	"fmt"
)

//CustomerRecyDevicesViewListType
type CustomerRecyDevicesViewListType struct {
	CustomerId float64
	Devices    map[string]RecyDeviceViewType
}

//New
func (res *CustomerRecyDevicesViewListType) New() {
	res.CustomerId = 1
	res.Devices = make(map[string]RecyDeviceViewType)
}

//GetByRedis
func (res *CustomerRecyDevicesViewListType) GetByRedis(dbIndex string) ResultType {
	resultVal := GetRedisForStoreApi(dbIndex, REDIS_CUSTOMER_RECY_DEVICEVIEWS, res.ToIdString())
	if resultVal.Result == RESULT_OK {
		res.StringToType(resultVal.Retval.(string))
	} else {
		return resultVal
	}

	resultVal.Retval = res.ToString()
	return resultVal
}

//GetByRedisByReel
func (res *CustomerRecyDevicesViewListType) GetByRedisByReel(dbIndex string) ResultType {
	resultVal := GetRedisWODbForStoreApi(dbIndex, REDIS_CUSTOMER_RECY_DEVICEVIEWS_REEL, REDIS_CUSTOMER_RECY_DEVICEVIEWS, res.ToIdString())
	if resultVal.Result == RESULT_OK {
		res.StringToType(resultVal.Retval.(string))
	} else {
		return resultVal
	}

	resultVal.Retval = res.ToString()
	return resultVal
}

//SaveToRedis
func (res *CustomerRecyDevicesViewListType) SaveToRedis() ResultType {
	resultVal := SaveRedisForStoreApi(REDIS_CUSTOMER_RECY_DEVICEVIEWS, res.ToIdString(), res.ToString())
	return resultVal
}

//SaveToRedisWODb
func (res *CustomerRecyDevicesViewListType) SaveToRedisWODb() ResultType {
	resultVal := SaveRedisWODbForStoreApi(REDIS_CUSTOMER_RECY_DEVICEVIEWS_REEL, res.ToIdString(), res.ToString())
	return resultVal
}

//ToId String
func (res *CustomerRecyDevicesViewListType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToByte
func (res *CustomerRecyDevicesViewListType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//ToString Get JSON
func (res *CustomerRecyDevicesViewListType) ToString() string {
	return string(res.ToByte())

}

//ByteToType
func (res *CustomerRecyDevicesViewListType) ByteToType(retByte []byte) {
	res.New()
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *CustomerRecyDevicesViewListType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}
