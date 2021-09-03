package main

import (
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
	http.HandleFunc("/data", data)
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
		w.Write([]byte("WasteEnhc"))
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

func data(w http.ResponseWriter, req *http.Request) {

	var retVal string = "FAIL"
	if err := req.ParseForm(); err != nil {
		logErr(err)
		return
	}
	appTypeVal := req.FormValue("APPTYPE")
	didVal := req.FormValue("DID")
	dataTypeVal := req.FormValue("DATATYPE")
	customerIdVal := getCustomerIdByDeviceId(didVal)
	req.Form.Add("CUSTOMERID", customerIdVal)

	client := http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.PostForm("http://waste-storeapi-cluster-ip/saveBulkDbMain", req.Form)
	if err != nil {
		logErr(err)

	} else {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logErr(err)
		}
		bodyString := string(bodyBytes)
		retVal = bodyString
	}

	var serviceClusterIp string = ""
	if appTypeVal == "RFID" {

		if dataTypeVal == "RF" {

			serviceClusterIp = "waste-rfreader-cluster-ip"

		} else if dataTypeVal == "GPS" {

			serviceClusterIp = "waste-gpsreader-cluster-ip"

		} else if dataTypeVal == "STATUS" {

			serviceClusterIp = "waste-statusreader-cluster-ip"

		} else if dataTypeVal == "THERM" {

			serviceClusterIp = "waste-thermreader-cluster-ip"

		} else if dataTypeVal == "CAM" {

			serviceClusterIp = "waste-camreader-cluster-ip"

		} else {
			retVal = "FAIL"
		}
	} else if appTypeVal == "ULT" {
		retVal = "OK"
	} else if appTypeVal == "RECY" {
		retVal = "OK"
	} else {
		retVal = "FAIL"
	}
	if serviceClusterIp != "" {
		retVal = sendDataToReader(req.Form, serviceClusterIp)
	}
	w.Write([]byte(retVal))
}

func sendDataToReader(data url.Values, readerClusterIp string) string {
	var retVal string = "FAIL"
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.PostForm("http://"+readerClusterIp+"/reader", data)
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

func getCustomerIdByDeviceId(didVal string) string {
	var customerId string = "-1"
	data := url.Values{
		"HASHKEY": {"serial-customer"},
		"SUBKEY":  {didVal},
	}
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.PostForm("http://waste-storeapi-cluster-ip/getkey", data)
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
