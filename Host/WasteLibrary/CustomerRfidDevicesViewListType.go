package WasteLibrary

import (
	"encoding/json"
	"fmt"
)

//CustomerRfidDevicesViewListType
type CustomerRfidDevicesViewListType struct {
	CustomerId float64
	Devices    map[string]RfidDeviceViewType
}

//New
func (res *CustomerRfidDevicesViewListType) New() {
	res.CustomerId = 1
	res.Devices = make(map[string]RfidDeviceViewType)
}

//GetByRedis
func (res *CustomerRfidDevicesViewListType) GetByRedis(dbIndex string) ResultType {
	var resultVal ResultType
	resultVal = GetRedisForStoreApi(dbIndex, REDIS_CUSTOMER_RFID_DEVICEVIEWS, res.ToIdString())
	if resultVal.Result == RESULT_OK {
		res.StringToType(resultVal.Retval.(string))
	} else {
		return resultVal
	}

	resultVal.Retval = res.ToString()
	return resultVal
}

//GetByRedisByReel
func (res *CustomerRfidDevicesViewListType) GetByRedisByReel(dbIndex string) ResultType {
	var resultVal ResultType
	resultVal = GetRedisWODbForStoreApi(dbIndex, REDIS_CUSTOMER_RFID_DEVICEVIEWS_REEL, REDIS_CUSTOMER_RFID_DEVICEVIEWS, res.ToIdString())
	if resultVal.Result == RESULT_OK {
		res.StringToType(resultVal.Retval.(string))
	} else {
		return resultVal
	}

	resultVal.Retval = res.ToString()
	return resultVal
}

//SaveToRedis
func (res *CustomerRfidDevicesViewListType) SaveToRedis() ResultType {
	var resultVal ResultType
	resultVal = SaveRedisForStoreApi(REDIS_CUSTOMER_RFID_DEVICEVIEWS, res.ToIdString(), res.ToString())
	return resultVal
}

//SaveToRedisWODb
func (res *CustomerRfidDevicesViewListType) SaveToRedisWODb() ResultType {
	var resultVal ResultType
	resultVal = SaveRedisWODbForStoreApi(REDIS_CUSTOMER_RFID_DEVICEVIEWS_REEL, res.ToIdString(), res.ToString())
	return resultVal
}

//ToId String
func (res *CustomerRfidDevicesViewListType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToByte
func (res *CustomerRfidDevicesViewListType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//ToString Get JSON
func (res *CustomerRfidDevicesViewListType) ToString() string {
	return string(res.ToByte())

}

//ByteToType
func (res *CustomerRfidDevicesViewListType) ByteToType(retByte []byte) {
	res.New()
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *CustomerRfidDevicesViewListType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}
