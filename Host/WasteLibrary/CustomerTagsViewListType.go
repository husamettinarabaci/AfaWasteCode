package WasteLibrary

import (
	"encoding/json"
	"fmt"
)

//CustomerTagsViewListType
type CustomerTagsViewListType struct {
	CustomerId float64
	Tags       map[string]TagViewType
}

//New
func (res *CustomerTagsViewListType) New() {
	res.CustomerId = 1
	res.Tags = make(map[string]TagViewType)
}

//GetByRedis
func (res *CustomerTagsViewListType) GetByRedis(dbIndex int) ResultType {
	var resultVal ResultType
	resultVal = GetRedisForStoreApi(REDIS_CUSTOMER_TAGVIEWS, res.ToIdString())
	if resultVal.Result == RESULT_OK {
		res.StringToType(resultVal.Retval.(string))
	} else {
		return resultVal
	}

	resultVal.Retval = res.ToString()
	return resultVal
}

//GetByRedisByReel
func (res *CustomerTagsViewListType) GetByRedisByReel() ResultType {
	var resultVal ResultType
	resultVal = GetRedisWODbForStoreApi(REDIS_CUSTOMER_TAGVIEWS_REEL, REDIS_CUSTOMER_TAGVIEWS, res.ToIdString())
	if resultVal.Result == RESULT_OK {
		res.StringToType(resultVal.Retval.(string))
	} else {
		return resultVal
	}

	resultVal.Retval = res.ToString()
	return resultVal
}

//SaveToRedis
func (res *CustomerTagsViewListType) SaveToRedis() ResultType {
	var resultVal ResultType
	resultVal = SaveRedisForStoreApi(REDIS_CUSTOMER_TAGVIEWS, res.ToIdString(), res.ToString())
	return resultVal
}

//SaveToRedisWODb
func (res *CustomerTagsViewListType) SaveToRedisWODb() ResultType {
	var resultVal ResultType
	resultVal = SaveRedisWODbForStoreApi(REDIS_CUSTOMER_TAGVIEWS_REEL, res.ToIdString(), res.ToString())
	return resultVal
}

//ToId String
func (res *CustomerTagsViewListType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToByte
func (res *CustomerTagsViewListType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//ToString Get JSON
func (res *CustomerTagsViewListType) ToString() string {
	return string(res.ToByte())

}

//Byte To CustomerTagsViewListType
func ByteToCustomerTagsViewListType(retByte []byte) CustomerTagsViewListType {
	var retVal CustomerTagsViewListType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To CustomerTagsViewListType
func StringToCustomerTagsViewListType(retStr string) CustomerTagsViewListType {
	return ByteToCustomerTagsViewListType([]byte(retStr))
}

//ByteToType
func (res *CustomerTagsViewListType) ByteToType(retByte []byte) {
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *CustomerTagsViewListType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}
