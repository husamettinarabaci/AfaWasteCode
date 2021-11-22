package WasteLibrary

import (
	"net/url"
)

//SaveBulkDbMainForStoreApi
func SaveBulkDbMainForStoreApi(data url.Values) ResultType {
	var resultVal ResultType
	resultVal.Result = RESULT_FAIL
	resultVal = HttpPostReq("http://waste-storeapi-cluster-ip/saveBulkDbMain", data)
	return resultVal
}

//GetBulkDbMainForStoreApi
func GetBulkDbMainForStoreApi(data url.Values) ResultType {
	var resultVal ResultType
	resultVal.Result = RESULT_FAIL
	resultVal = HttpPostReq("http://waste-storeapi-cluster-ip/getBulkDbMain", data)
	return resultVal
}

//SaveStaticDbMainForStoreApi
func SaveStaticDbMainForStoreApi(data url.Values) ResultType {
	var resultVal ResultType
	resultVal.Result = RESULT_FAIL
	resultVal = HttpPostReq("http://waste-storeapi-cluster-ip/saveStaticDbMain", data)
	return resultVal
}

//GetStaticDbMainForStoreApi
func GetStaticDbMainForStoreApi(data url.Values) ResultType {
	var resultVal ResultType
	resultVal.Result = RESULT_FAIL
	resultVal = HttpPostReq("http://waste-storeapi-cluster-ip/getStaticDbMain", data)
	return resultVal
}

//SaveReaderDbMainForStoreApi
func SaveReaderDbMainForStoreApi(data url.Values) ResultType {
	var resultVal ResultType
	resultVal.Result = RESULT_FAIL
	resultVal = HttpPostReq("http://waste-storeapi-cluster-ip/saveReaderDbMain", data)
	return resultVal
}

//GetReaderDbMainForStoreApi
func GetReaderDbMainForStoreApi(data url.Values) ResultType {
	var resultVal ResultType
	resultVal.Result = RESULT_FAIL
	resultVal = HttpPostReq("http://waste-storeapi-cluster-ip/getReaderDbMain", data)
	return resultVal
}

//SaveConfigDbMainForStoreApi
func SaveConfigDbMainForStoreApi(data url.Values) ResultType {
	var resultVal ResultType
	resultVal.Result = RESULT_FAIL
	resultVal = HttpPostReq("http://waste-storeapi-cluster-ip/saveConfigDbMain", data)
	return resultVal
}

//GetConfigDbMainForStoreApi
func GetConfigDbMainForStoreApi(data url.Values) ResultType {
	var resultVal ResultType
	resultVal.Result = RESULT_FAIL
	resultVal = HttpPostReq("http://waste-storeapi-cluster-ip/getConfigDbMain", data)
	return resultVal
}

//SaveRedisForStoreApi
func SaveRedisForStoreApi(hKey string, sKey string, kVal string) ResultType {
	var resultVal ResultType
	resultVal.Result = RESULT_FAIL
	data := url.Values{
		REDIS_HASHKEY:  {hKey},
		REDIS_SUBKEY:   {sKey},
		REDIS_KEYVALUE: {kVal},
	}
	resultVal = HttpPostReq("http://waste-storeapi-cluster-ip/setkey", data)
	return resultVal
}

//SaveRedisWODbForStoreApi
func SaveRedisWODbForStoreApi(hKey string, sKey string, kVal string) ResultType {
	var resultVal ResultType
	resultVal.Result = RESULT_FAIL
	data := url.Values{
		REDIS_HASHKEY:  {hKey},
		REDIS_SUBKEY:   {sKey},
		REDIS_KEYVALUE: {kVal},
	}
	resultVal = HttpPostReq("http://waste-storeapi-cluster-ip/setkey", data)
	return resultVal
}

//PublishRedisForStoreApi
func PublishRedisForStoreApi(channelKey string, dataType string, dataVal string) ResultType {
	var resultVal ResultType
	resultVal.Result = dataType
	resultVal.Retval = dataVal
	data := url.Values{
		REDIS_CHANNELKEY: {channelKey},
		REDIS_KEYVALUE:   {resultVal.ToString()},
	}
	resultVal = HttpPostReq("http://waste-storeapi-cluster-ip/publishkey", data)
	return resultVal
}

//DeleteRedisForStoreApi
func DeleteRedisForStoreApi(hKey string, sKey string) ResultType {
	var resultVal ResultType
	resultVal.Result = RESULT_FAIL
	data := url.Values{
		REDIS_HASHKEY: {hKey},
		REDIS_SUBKEY:  {sKey},
	}
	resultVal = HttpPostReq("http://waste-storeapi-cluster-ip/deletekey", data)
	return resultVal
}

//GetRedisForStoreApi
func GetRedisForStoreApi(hKey string, sKey string) ResultType {
	var resultVal ResultType
	resultVal.Result = RESULT_FAIL
	data := url.Values{
		REDIS_HASHKEY: {hKey},
		REDIS_SUBKEY:  {sKey},
	}
	resultVal = HttpPostReq("http://waste-storeapi-cluster-ip/getkey", data)
	return resultVal
}

//GetRedisWODbForStoreApi
func GetRedisWODbForStoreApi(hKey string, hBaseKey string, sKey string) ResultType {
	var resultVal ResultType
	resultVal.Result = RESULT_FAIL
	data := url.Values{
		REDIS_HASHKEY:     {hKey},
		REDIS_HASHBASEKEY: {hBaseKey},
		REDIS_SUBKEY:      {sKey},
	}
	resultVal = HttpPostReq("http://waste-storeapi-cluster-ip/getkey", data)
	return resultVal
}
