package main

import (
	"fmt"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/devafatek/WasteLibrary"
)

func main() {

	http.HandleFunc("/health", WasteLibrary.HealthHandler)
	http.HandleFunc("/readiness", WasteLibrary.ReadinessHandler)
	http.HandleFunc("/status", WasteLibrary.StatusHandler)
	http.HandleFunc("/log", log)
	http.ListenAndServe(":80", nil)
}

func log(w http.ResponseWriter, req *http.Request) {

	if WasteLibrary.AllowCors {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,access-control-allow-origin, access-control-allow-headers")
	}
	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_OK
	resultVal.Retval = WasteLibrary.RESULT_SUCCESS_OK

	if err := req.ParseForm(); err != nil {
		logErr(err)
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_HTTP_PARSE
		w.Write(resultVal.ToByte())

		return
	}
	container := req.FormValue(WasteLibrary.LOGGER_CONTAINER)
	logType := req.FormValue(WasteLibrary.LOGGER_LOGTYPE)
	funcVal := req.FormValue(WasteLibrary.LOGGER_FUNC)
	logVal := req.FormValue(WasteLibrary.LOGGER_LOG)
	logStr("Time : " + WasteLibrary.GetTime() + " - Container : " + container + " - LogType : " + logType + " - Func : " + funcVal + " - Log : " + logVal + " - IP : " + req.RemoteAddr)
	w.Write(resultVal.ToByte())

}

func logErr(err error) {
	if err != nil {
		logStr(err.Error())
	}
}

func logStr(value string) {
	fmt.Println(value)
}
