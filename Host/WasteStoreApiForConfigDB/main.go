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
var configDb *sql.DB
var err error

func initStart() {

	WasteLibrary.LogStr("Successfully connected!")
	go WasteLibrary.InitLog()
}

func main() {

	initStart()

	var configDbHost string = "waste-configdb-cluster-ip"
	configDbInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		configDbHost, port, user, password, dbname)

	configDb, err = sql.Open("postgres", configDbInfo)
	WasteLibrary.LogErr(err)
	defer configDb.Close()

	err = configDb.Ping()
	WasteLibrary.LogErr(err)

	http.HandleFunc("/health", WasteLibrary.HealthHandler)
	http.HandleFunc("/readiness", WasteLibrary.ReadinessHandler)
	http.HandleFunc("/status", WasteLibrary.StatusHandler)
	http.HandleFunc("/saveConfigDbMain", saveConfigDbMain)
	http.HandleFunc("/getConfigDbMain", getConfigDbMain)
	http.ListenAndServe(":80", nil)

}

func saveConfigDbMain(w http.ResponseWriter, req *http.Request) {

	if WasteLibrary.AllowCors {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
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

	var currentHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.StringToHttpClientHeaderType(req.FormValue(WasteLibrary.HTTP_HEADER))
	var execSQL string = ""
	if currentHttpHeader.DataType == WasteLibrary.DATATYPE_USER {

		var currentData WasteLibrary.UserType = WasteLibrary.StringToUserType(req.FormValue(WasteLibrary.HTTP_DATA))
		if currentData.UserId != 0 {
			execSQL = currentData.UpdateSQL()
		} else {

			execSQL = currentData.InsertSQL()
		}
		var userId int = 0
		errDb := configDb.QueryRow(execSQL).Scan(&userId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.UserId = float64(userId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_CUSTOMER {

		var currentData WasteLibrary.CustomerType = WasteLibrary.StringToCustomerType(req.FormValue(WasteLibrary.HTTP_DATA))

		if currentData.CustomerId != 0 {
			execSQL = currentData.UpdateSQL()
		} else {

			execSQL = currentData.InsertSQL()
		}
		var customerId int = 0
		errDb := configDb.QueryRow(execSQL).Scan(&customerId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.CustomerId = float64(customerId)
		resultVal.Retval = currentData.ToIdString()

	} else {
		resultVal.Result = WasteLibrary.RESULT_FAIL
	}
	w.Write(resultVal.ToByte())

}

func getConfigDbMain(w http.ResponseWriter, req *http.Request) {

	if WasteLibrary.AllowCors {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
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
	var currentHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.StringToHttpClientHeaderType(req.FormValue(WasteLibrary.HTTP_HEADER))
	if currentHttpHeader.DataType == WasteLibrary.DATATYPE_USER {

		var currentData WasteLibrary.UserType = WasteLibrary.StringToUserType(req.FormValue(WasteLibrary.HTTP_DATA))

		errDb := currentData.SelectWithDb(configDb)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		resultVal.Retval = currentData.ToString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_CUSTOMER {
		var currentData WasteLibrary.CustomerType = WasteLibrary.StringToCustomerType(req.FormValue(WasteLibrary.HTTP_DATA))
		errDb := currentData.SelectWithDb(configDb)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		resultVal.Retval = currentData.ToString()

	} else {
		resultVal.Result = WasteLibrary.RESULT_FAIL
	}

	w.Write(resultVal.ToByte())

}
