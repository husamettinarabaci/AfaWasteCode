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

	if repeat == "0" {
		data := url.Values{
			"APPTYPE":    {appTypeVal},
			"DID":        {didVal},
			"DATATYPE":   {dataTypeVal},
			"TIME":       {dataTime},
			"CUSTOMERID": {customerIdVal},
		}

		readerAppStatus := dataMap["READERAPPSTATUS"][0]
		readerConnStatus := dataMap["READERCONNSTATUS"][0]
		readerStatus := dataMap["READERSTATUS"][0]
		camAppStatus := dataMap["CAMAPPSTATUS"][0]
		camConnStatus := dataMap["CAMCONNSTATUS"][0]
		camStatus := dataMap["CAMSTATUS"][0]
		gpsAppStatus := dataMap["GPSAPPSTATUS"][0]
		gpsConnStatus := dataMap["GPSCONNSTATUS"][0]
		gpsStatus := dataMap["GPSSTATUS"][0]
		thermAppStatus := dataMap["THERMAPSTATUS"][0]
		transferAppStatus := dataMap["TRANSFERAPP"][0]
		aliveStatus := dataMap["ALIVESTATUS"][0]
		contactStatus := dataMap["CONTACTSTATUS"][0]
		WasteLibrary.LogStr(repeat + " - " + dataTypeVal + " - " + readerAppStatus + " - " + readerConnStatus + " - " + readerStatus + " - " + camAppStatus + " - " + camConnStatus + " - " + camStatus + " - " + gpsAppStatus + " - " + gpsConnStatus + " - " + gpsStatus + " - " + thermAppStatus + " - " + transferAppStatus + " - " + aliveStatus + " - " + contactStatus)

		data.Add("OPTYPE", "STATUS")
		data.Add("READERAPPSTATUS", readerAppStatus)
		data.Add("READERCONNSTATUS", readerConnStatus)
		data.Add("READERSTATUS", readerStatus)
		data.Add("CAMAPPSTATUS", camAppStatus)
		data.Add("CAMCONNSTATUS", camConnStatus)
		data.Add("CAMSTATUS", camStatus)
		data.Add("GPSAPPSTATUS", gpsAppStatus)
		data.Add("GPSCONNSTATUS", gpsConnStatus)
		data.Add("GPSSTATUS", gpsStatus)
		data.Add("THERMAPSTATUS", thermAppStatus)
		data.Add("TRANSFERAPP", transferAppStatus)
		data.Add("ALIVESTATUS", aliveStatus)
		data.Add("CONTACTSTATUS", contactStatus)
		resultVal = WasteLibrary.SaveStaticDbMainForStoreApi(data)
		WasteLibrary.LogStr("Save StaticDbMain : " + appTypeVal + " - " + dataTypeVal + " - " + resultVal.ToString())

		if resultVal.Result == "OK" {
			resultVal = WasteLibrary.SaveRedisForStoreApi("device-status", didVal, dataVal)
			WasteLibrary.LogStr("Save Redis : " + appTypeVal + " - " + dataTypeVal + " - " + resultVal.ToString() + " - " + dataVal)
		}
	} else {
		resultVal.Result = "OK"
	}
	w.Write(resultVal.ToByte())
}
