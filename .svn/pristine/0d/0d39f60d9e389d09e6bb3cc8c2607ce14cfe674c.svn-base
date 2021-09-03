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
		w.Write([]byte("WasteCamReader"))
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

		uId := dataMap["UID"][0]
		tagId := dataMap["TAGID"][0]
		logStr(repeat + " - " + dataTypeVal + " - " + tagId + " - " + uId)

		data.Add("OPTYPE", "CAM_IMAGE")
		data.Add("TAGID", tagId)
		data.Add("UID", uId)
		data.Add("IMAGE", "1")
		retVal = saveStaticDbMainForStoreApi(data)
		logStr("Save StaticDbMain : " + appTypeVal + " - " + dataTypeVal + " - " + retVal)
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
