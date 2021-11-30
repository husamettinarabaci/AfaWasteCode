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
	var resultVal ResultType
	resultVal = GetRedisForStoreApi(dbIndex, REDIS_CUSTOMER_ULT_DEVICEVIEWS, res.ToIdString())
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
	var resultVal ResultType
	resultVal = GetRedisWODbForStoreApi(dbIndex, REDIS_CUSTOMER_ULT_DEVICEVIEWS_REEL, REDIS_CUSTOMER_ULT_DEVICEVIEWS, res.ToIdString())
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
	var resultVal ResultType
	resultVal = SaveRedisForStoreApi(REDIS_CUSTOMER_ULT_DEVICEVIEWS, res.ToIdString(), res.ToString())
	return resultVal
}

//SaveToRedisWODb
func (res *CustomerUltDevicesViewListType) SaveToRedisWODb() ResultType {
	var resultVal ResultType
	resultVal = SaveRedisWODbForStoreApi(REDIS_CUSTOMER_ULT_DEVICEVIEWS_REEL, res.ToIdString(), res.ToString())
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

//Byte To CustomerUltDevicesViewListType
func ByteToCustomerUltDevicesViewListType(retByte []byte) CustomerUltDevicesViewListType {
	var retVal CustomerUltDevicesViewListType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To CustomerUltDevicesViewListType
func StringToCustomerUltDevicesViewListType(retStr string) CustomerUltDevicesViewListType {
	return ByteToCustomerUltDevicesViewListType([]byte(retStr))
}

//ByteToType
func (res *CustomerUltDevicesViewListType) ByteToType(retByte []byte) {
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *CustomerUltDevicesViewListType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}
