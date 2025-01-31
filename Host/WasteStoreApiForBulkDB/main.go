package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/devafatek/WasteLibrary"
	_ "github.com/lib/pq"
)

var port int = 5432
var user string = os.Getenv("POSTGRES_USER")
var password string = os.Getenv("POSTGRES_PASSWORD")
var dbname string = "bulkdb"

var ctx = context.Background()
var bulkDb *sql.DB
var err error

func initStart() {

	WasteLibrary.LogStr("Successfully connected!")
	go WasteLibrary.InitLog()
}

func main() {

	initStart()

	var bulkDbHost string = "waste-psqldb-cluster-ip"
	bulkDbInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		bulkDbHost, port, user, password, dbname)

	bulkDb, err = sql.Open("postgres", bulkDbInfo)
	WasteLibrary.LogErr(err)
	defer bulkDb.Close()

	err = bulkDb.Ping()
	WasteLibrary.LogErr(err)

	http.HandleFunc("/health", WasteLibrary.HealthHandler)
	http.HandleFunc("/readiness", WasteLibrary.ReadinessHandler)
	http.HandleFunc("/status", WasteLibrary.StatusHandler)
	http.HandleFunc("/saveBulkDbMain", saveBulkDbMain)
	http.ListenAndServe(":80", nil)
}

func saveBulkDbMain(w http.ResponseWriter, req *http.Request) {

	if WasteLibrary.AllowCors {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,access-control-allow-origin, access-control-allow-headers")
	}
	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL

	if err := req.ParseForm(); err != nil {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_HTTP_PARSE
		w.Write(resultVal.ToByte())

		WasteLibrary.LogErr(err)
		return
	}
	var currentHttpHeader WasteLibrary.HttpClientHeaderType
	currentHttpHeader.StringToType(req.FormValue(WasteLibrary.HTTP_HEADER))
	dataVal := req.FormValue(WasteLibrary.HTTP_DATA)

	var insertSQL string = fmt.Sprintf(`INSERT INTO public.`+WasteLibrary.DATATYPE_LISTENER_DATA+`  
	  (AppType,DeviceNo,DeviceId,CustomerId,Time,Repeat,DeviceType,ReaderType,DataType,Token,Data) 
	  VALUES ('%s','%s',%f,%f,'%s','%s','%s','%s','%s','%s','%s');`,
		currentHttpHeader.AppType,
		currentHttpHeader.DeviceNo,
		currentHttpHeader.DeviceId,
		currentHttpHeader.CustomerId,
		currentHttpHeader.Time,
		currentHttpHeader.Repeat,
		currentHttpHeader.DeviceType,
		currentHttpHeader.ReaderType,
		currentHttpHeader.DataType,
		"",
		dataVal)
	_, errDb := bulkDb.Exec(insertSQL)
	if errDb != nil {
		WasteLibrary.LogErr(err)
		resultVal.Result = WasteLibrary.RESULT_FAIL
	} else {
		resultVal.Result = WasteLibrary.RESULT_OK
		resultVal.Retval = WasteLibrary.RESULT_SUCCESS_OK
	}
	w.Write(resultVal.ToByte())

}
