package WasteLibrary

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

//AppStatus
var AppStatus string = ACTIVE

//ConnStatus
var ConnStatus string = PASSIVE

//DeviceStatus
var DeviceStatus string = PASSIVE

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
	var resultVal ResultType
	if err := req.ParseForm(); err != nil {
		LogErr(err)
		resultVal.Result = FAIL
	} else {

		opType := req.FormValue(OPTYPE)
		LogStr(opType)
		resultVal.Result = FAIL
		if opType == APP {
			if CurrentCheckStatu.AppStatu == ACTIVE {
				resultVal.Result = OK
			} else {
				resultVal.Result = FAIL
			}
		} else if opType == CONN {
			if CurrentCheckStatu.ConnStatu == ACTIVE {
				resultVal.Result = OK
			} else {
				resultVal.Result = FAIL
			}
		} else if opType == CAM {
			if CurrentCheckStatu.DeviceStatu == ACTIVE {
				resultVal.Result = OK
			} else {
				resultVal.Result = FAIL
			}
		} else {
			resultVal.Result = FAIL
		}
	}
	w.Write(resultVal.ToByte())
}

//HttpPostReq
func HttpPostReq(url string, data url.Values) ResultType {
	var resultVal ResultType
	resultVal.Result = FAIL
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
