package WasteLibrary

import (
	"net/url"

	"github.com/AfatekDevelopers/result_lib_go/devafatekresult"
)

//SaveStaticDbMainForStoreApi
func SaveStaticDbMainForStoreApi(data url.Values) devafatekresult.ResultType {
	var resultVal devafatekresult.ResultType
	resultVal.Result = "FAIL"
	resultVal = HttpPostReq("http://waste-storeapi-cluster-ip/saveStaticDbMain", data)
	return resultVal
}

//SaveRedisForStoreApi
func SaveRedisForStoreApi(hKey string, sKey string, kVal string) devafatekresult.ResultType {
	var resultVal devafatekresult.ResultType
	resultVal.Result = "FAIL"
	data := url.Values{
		"HASHKEY":  {hKey},
		"SUBKEY":   {sKey},
		"KEYVALUE": {kVal},
	}
	resultVal = HttpPostReq("http://waste-storeapi-cluster-ip/setkey", data)
	return resultVal
}

//GetRedisForStoreApi
func GetRedisForStoreApi(hKey string, sKey string) devafatekresult.ResultType {
	var resultVal devafatekresult.ResultType
	resultVal.Result = "FAIL"
	data := url.Values{
		"HASHKEY": {hKey},
		"SUBKEY":  {sKey},
	}
	resultVal = HttpPostReq("http://waste-storeapi-cluster-ip/getkey", data)
	return resultVal
}
