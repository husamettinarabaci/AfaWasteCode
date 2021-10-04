package WasteLibrary

import (
	"net/url"
)

//SaveBulkDbMainForStoreApi
func SaveBulkDbMainForStoreApi(data url.Values) ResultType {
	var resultVal ResultType
	resultVal.Result = FAIL
	resultVal = HttpPostReq("http://waste-storeapi-cluster-ip/saveBulkDbMain", data)
	return resultVal
}

//SaveStaticDbMainForStoreApi
func SaveStaticDbMainForStoreApi(data url.Values) ResultType {
	var resultVal ResultType
	resultVal.Result = FAIL
	resultVal = HttpPostReq("http://waste-storeapi-cluster-ip/saveStaticDbMain", data)
	return resultVal
}

//GetStaticDbMainForStoreApi
func GetStaticDbMainForStoreApi(data url.Values) ResultType {
	var resultVal ResultType
	resultVal.Result = FAIL
	resultVal = HttpPostReq("http://waste-storeapi-cluster-ip/getStaticDbMain", data)
	return resultVal
}

//SaveReaderDbMainForStoreApi
func SaveReaderDbMainForStoreApi(data url.Values) ResultType {
	var resultVal ResultType
	resultVal.Result = FAIL
	resultVal = HttpPostReq("http://waste-storeapi-cluster-ip/saveReaderDbMain", data)
	return resultVal
}

//SaveRedisForStoreApi
func SaveRedisForStoreApi(hKey string, sKey string, kVal string) ResultType {
	var resultVal ResultType
	resultVal.Result = FAIL
	data := url.Values{
		HASHKEY:  {hKey},
		SUBKEY:   {sKey},
		KEYVALUE: {kVal},
	}
	resultVal = HttpPostReq("http://waste-storeapi-cluster-ip/setkey", data)
	return resultVal
}

//GetRedisForStoreApi
func GetRedisForStoreApi(hKey string, sKey string) ResultType {
	var resultVal ResultType
	resultVal.Result = FAIL
	data := url.Values{
		HASHKEY: {hKey},
		SUBKEY:  {sKey},
	}
	resultVal = HttpPostReq("http://waste-storeapi-cluster-ip/getkey", data)
	return resultVal
}
