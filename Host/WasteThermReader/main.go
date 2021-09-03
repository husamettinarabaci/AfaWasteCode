package main

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/AfatekDevelopers/result_lib_go/devafatekresult"
	"github.com/devafatek/WasteLibrary"
)

func initStart() {

	WasteLibrary.LogStr("Successfully connected!")
}
func main() {

	initStart()

	http.HandleFunc("/health", WasteLibrary.HealthHandler)
	http.HandleFunc("/readiness", WasteLibrary.ReadinessHandler)
	http.HandleFunc("/status", WasteLibrary.StatusHandler)
	http.HandleFunc("/reader", reader)
	http.ListenAndServe(":80", nil)
}

func reader(w http.ResponseWriter, req *http.Request) {
	var resultVal devafatekresult.ResultType
	resultVal.Result = "FAIL"
	if err := req.ParseForm(); err != nil {
		WasteLibrary.LogErr(err)
		return
	}
	appTypeVal := req.FormValue("APPTYPE")
	didVal := req.FormValue("DID")
	dataTypeVal := req.FormValue("DATATYPE")
	dataVal := req.FormValue("DATA")
	dataTime := req.FormValue("TIME")
	repeat := req.FormValue("REPEAT")
	customerIdVal := req.FormValue("CUSTOMERID")

	var dataMap map[string][]string
	err := json.Unmarshal([]byte(dataVal), &dataMap)
	if err != nil {
		WasteLibrary.LogErr(err)
	}
	if resultVal.Result == "OK" {
		data := url.Values{
			"APPTYPE":    {appTypeVal},
			"DID":        {didVal},
			"DATATYPE":   {dataTypeVal},
			"TIME":       {dataTime},
			"CUSTOMERID": {customerIdVal},
		}

		therm := dataMap["THERM"][0]
		WasteLibrary.LogStr(repeat + " - " + dataTypeVal + " - " + therm)

		data.Add("OPTYPE", "thermdata")
		data.Add("THERM", therm)
		resultVal = WasteLibrary.SaveStaticDbMainForStoreApi(data)
		WasteLibrary.LogStr("Save StaticDbMain : " + appTypeVal + " - " + dataTypeVal + " - " + resultVal.ToString())

		resultVal = WasteLibrary.SaveRedisForStoreApi("device-therm", didVal, dataVal)
		WasteLibrary.LogStr("Save Redis : " + appTypeVal + " - " + dataTypeVal + " - " + resultVal.ToString() + " - " + dataVal)
	} else {
		resultVal.Result = "OK"
	}

	w.Write(resultVal.ToByte())
}
