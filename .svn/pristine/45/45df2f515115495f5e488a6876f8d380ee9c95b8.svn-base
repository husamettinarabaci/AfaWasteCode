package main

import (
	"fmt"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

func main() {

	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/readiness", readinessHandler)
	http.HandleFunc("/status", status)
	http.HandleFunc("/log", log)
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
		w.Write([]byte("WasteLogServer"))
	} else if opType == "APP" {
		w.Write([]byte("OK"))
	} else {
		w.Write([]byte("FAIL"))
	}
}

func log(w http.ResponseWriter, req *http.Request) {

	var retVal string = "FAIL"
	if err := req.ParseForm(); err != nil {
		logErr(err)
		return
	}
	container := req.FormValue("CONTAINER")
	logType := req.FormValue("LOGTYPE")
	logVal := req.FormValue("LOG")
	logStr("Time : " + time.Now().String() + " - Container : " + container + " - LogType : " + logType + " - Log : " + logVal)
	w.Write([]byte(retVal))
}

func logErr(err error) {
	if err != nil {
		logStr(err.Error())
	}
}

func logStr(value string) {
	fmt.Println(value)
}
