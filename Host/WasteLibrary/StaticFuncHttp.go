package WasteLibrary

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/AfatekDevelopers/result_lib_go/devafatekresult"
)

//AppStatus
var AppStatus string = "1"

//ConnStatus
var ConnStatus string = "0"

//DeviceStatus
var DeviceStatus string = "0"

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
	var resultVal WasteLibrary.ResultType
	if err := req.ParseForm(); err != nil {
		LogErr(err)
		resultVal.Result = "FAIL"
	} else {

		opType := req.FormValue("OPTYPE")
		LogStr(opType)
		resultVal.Result = "FAIL"
		if opType == "APP" {
			if CurrentCheckStatu.AppStatu == "1" {
				resultVal.Result = "OK"
			} else {
				resultVal.Result = "FAIL"
			}
		} else if opType == "CONN" {
			if CurrentCheckStatu.ConnStatu == "1" {
				resultVal.Result = "OK"
			} else {
				resultVal.Result = "FAIL"
			}
		} else if opType == "CAM" {
			if CurrentCheckStatu.DeviceStatu == "1" {
				resultVal.Result = "OK"
			} else {
				resultVal.Result = "FAIL"
			}
		} else {
			resultVal.Result = "FAIL"
		}
	}
	w.Write(resultVal.ToByte())
}

//HttpPostReq
func HttpPostReq(url string, data url.Values) WasteLibrary.ResultType {
	var resultVal WasteLibrary.ResultType
	resultVal.Result = "FAIL"
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
		resultVal = devafatekresult.ByteToResultType(bodyBytes)
		LogStr(resultVal.ToString())
	}

	return resultVal
}
