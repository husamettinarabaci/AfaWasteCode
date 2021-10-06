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
var dbname string = os.Getenv("POSTGRES_DB")

var ctx = context.Background()
var bulkDb *sql.DB
var err error

func initStart() {

	WasteLibrary.LogStr("Successfully connected!")

}

func main() {

	initStart()

	var bulkDbHost string = "waste-bulkdb-cluster-ip"
	bulkDbInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		bulkDbHost, port, user, password, dbname)

	bulkDb, err = sql.Open("postgres", bulkDbInfo)
	WasteLibrary.LogErr(err)
	defer bulkDb.Close()

	err = bulkDb.Ping()
	WasteLibrary.LogErr(err)

	WasteLibrary.LogStr("Start")
	http.HandleFunc("/health", WasteLibrary.HealthHandler)
	http.HandleFunc("/readiness", WasteLibrary.ReadinessHandler)
	http.HandleFunc("/status", WasteLibrary.StatusHandler)
	http.HandleFunc("/saveBulkDbMain", saveBulkDbMain)
	http.ListenAndServe(":80", nil)
}

func saveBulkDbMain(w http.ResponseWriter, req *http.Request) {
	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	if err := req.ParseForm(); err != nil {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_HTTP_PARSE
		w.Write(resultVal.ToByte())
		WasteLibrary.LogErr(err)
		return
	}
	var currentHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.StringToHttpClientHeaderType(req.FormValue(WasteLibrary.HTTP_HEADER))
	dataVal := req.FormValue(WasteLibrary.HTTP_DATA)
	WasteLibrary.LogStr("Header : " + currentHttpHeader.ToString())
	WasteLibrary.LogStr("Data : " + dataVal)
	var insertSQL string = fmt.Sprintf(`INSERT INTO public.listenerdata(
		app_type,device_id,optype,data,customer_id,data_time)
	   VALUES ('%s',%f,'%s','%s',%f,'%s');`, currentHttpHeader.AppType, currentHttpHeader.DeviceId, currentHttpHeader.OpType, dataVal, currentHttpHeader.CustomerId, currentHttpHeader.Time)
	WasteLibrary.LogStr(insertSQL)
	_, errDb := bulkDb.Exec(insertSQL)
	if errDb != nil {
		WasteLibrary.LogErr(err)
		resultVal.Result = WasteLibrary.RESULT_FAIL
	} else {
		resultVal.Result = WasteLibrary.RESULT_OK
	}
	w.Write(resultVal.ToByte())
}
