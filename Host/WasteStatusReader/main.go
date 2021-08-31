package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"
)

var debug bool = os.Getenv("DEBUG") == "1"
var appStatus string = "1"

func initStart() {

	logStr("Successfully connected!")
}
func main() {

	initStart()

	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/readiness", readinessHandler)
	http.HandleFunc("/status", status)
	http.HandleFunc("/reader", reader)
	http.ListenAndServe(":80", nil)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func readinessHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func status(w http.ResponseWriter, req *http.Request) {

	if err := req.ParseForm(); err != nil {
		logErr(err)
		return
	}
	opType := req.FormValue("OPTYPE")
	logStr(opType)

	if opType == "TYPE" {
		w.Write([]byte("WasteStatusReader"))
	} else if opType == "APP" {
		if appStatus == "1" {
			w.Write([]byte("OK"))
		} else {
			w.Write([]byte("FAIL"))
		}
	} else {
		w.Write([]byte("FAIL"))
	}
}

func reader(w http.ResponseWriter, req *http.Request) {

	var retVal string = "FAIL"
	if err := req.ParseForm(); err != nil {
		logErr(err)
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
		logErr(err)
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
		logStr(repeat + " - " + dataTypeVal + " - " + readerAppStatus + " - " + readerConnStatus + " - " + readerStatus + " - " + camAppStatus + " - " + camConnStatus + " - " + camStatus + " - " + gpsAppStatus + " - " + gpsConnStatus + " - " + gpsStatus + " - " + thermAppStatus + " - " + transferAppStatus + " - " + aliveStatus + " - " + contactStatus)

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
		retVal = saveStaticDbMainForStoreApi(data)
		logStr("Save StaticDbMain : " + appTypeVal + " - " + dataTypeVal + " - " + retVal)

		if retVal == "OK" {
			retVal = saveStatusRedisForStoreApi(didVal, dataVal)
			logStr("Save Redis : " + appTypeVal + " - " + dataTypeVal + " - " + retVal + " - " + dataVal)
		}
	} else {
		retVal = "OK"
	}
	w.Write([]byte(retVal))
}

func saveStaticDbMainForStoreApi(data url.Values) string {
	var retVal string = "FAIL"
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.PostForm("http://waste-storeapi-cluster-ip/saveStaticDbMain", data)
	if err != nil {
		logErr(err)

	} else {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logErr(err)
		}
		bodyString := string(bodyBytes)
		if bodyString == "OK" {
			retVal = "OK"
		}
		logStr(bodyString)
	}
	return retVal
}

func saveStatusRedisForStoreApi(didVal string, kVal string) string {
	var customerId string = "-1"
	data := url.Values{
		"HASHKEY":  {"device-status"},
		"SUBKEY":   {didVal},
		"KEYVALUE": {kVal},
	}
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.PostForm("http://waste-storeapi-cluster-ip/setkey", data)
	if err != nil {
		logErr(err)

	} else {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logErr(err)
		}
		bodyString := string(bodyBytes)
		if bodyString != "NOT" {
			customerId = bodyString
		}
	}

	return customerId
}

func logErr(err error) {
	if err != nil {
		sendLogServer("ERR", err.Error())
	}
}

func logStr(value string) {
	if debug {
		sendLogServer("INFO", value)
	}
}

var container string = os.Getenv("CONTAINER_TYPE")

func sendLogServer(logType string, logVal string) string {
	var retVal string = "FAIL"
	data := url.Values{
		"CONTAINER": {container},
		"LOGTYPE":   {logType},
		"LOG":       {logVal},
	}
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.PostForm("http://waste-logserver-cluster-ip/log", data)
	if err != nil {
		logErr(err)

	} else {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logErr(err)
		}
		bodyString := string(bodyBytes)
		if bodyString != "NOT" {
			retVal = bodyString
		}
	}

	return retVal
}
