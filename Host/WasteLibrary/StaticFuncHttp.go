package WasteLibrary

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

//AppStatus
var AppStatus string = STATU_ACTIVE

//ConnStatus
var ConnStatus string = STATU_PASSIVE

//DeviceStatus
var DeviceStatus string = STATU_PASSIVE

//HealthHandler
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

//ReadinessHandler
func ReadinessHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

//StaStatusHandlertus
func StatusHandler(w http.ResponseWriter, req *http.Request) {

	if AllowCors {

		w.Header().Set("Access-Control-Allow-Origin", "*")
	}
	var resultVal ResultType

	if err := req.ParseForm(); err != nil {
		LogErr(err)
		resultVal.Result = RESULT_FAIL
		resultVal.Retval = RESULT_ERROR_HTTP_PARSE
	} else {

		readerType := req.FormValue(HTTP_CHECKTYPE)
		LogStr(readerType)
		resultVal.Result = RESULT_FAIL
		if readerType == CHECKTYPE_APP {
			if CurrentCheckStatu.AppStatu == STATU_ACTIVE {
				resultVal.Result = RESULT_OK
				resultVal.Retval = Version
			} else {
				resultVal.Result = RESULT_FAIL
			}
		} else if readerType == CHECKTYPE_CONN {
			if CurrentCheckStatu.ConnStatu == STATU_ACTIVE {
				resultVal.Result = RESULT_OK
			} else {
				resultVal.Result = RESULT_FAIL
			}
		} else if readerType == CHECKTYPE_DEVICE {
			if CurrentCheckStatu.DeviceStatu == STATU_ACTIVE {
				resultVal.Result = RESULT_OK
			} else {
				resultVal.Result = RESULT_FAIL
			}
		} else {
			resultVal.Result = RESULT_FAIL
		}
	}
	w.Write(resultVal.ToByte())

}

//HttpPostReq
func HttpPostReq(url string, data url.Values) ResultType {
	var resultVal ResultType
	resultVal.Result = RESULT_FAIL
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.PostForm(url, data)
	if err != nil {
		LogErr(err)

	} else {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			LogErr(err)
		}
		resultVal = ByteToResultType(bodyBytes)
		LogStr(resultVal.ToString())
	}

	return resultVal
}
