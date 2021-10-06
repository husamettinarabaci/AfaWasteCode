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

	WasteLibrary.LogStr("Start")
	http.HandleFunc("/health", WasteLibrary.HealthHandler)
	http.HandleFunc("/readiness", WasteLibrary.ReadinessHandler)
	http.HandleFunc("/status", WasteLibrary.StatusHandler)
	http.HandleFunc("/saveConfigDbMain", saveConfigDbMain)
	http.HandleFunc("/getConfigDbMain", getConfigDbMain)
	http.ListenAndServe(":80", nil)

}

func saveConfigDbMain(w http.ResponseWriter, req *http.Request) {
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
	WasteLibrary.LogStr("Header : " + currentHttpHeader.ToString())
	WasteLibrary.LogStr("Data : " + req.FormValue(WasteLibrary.HTTP_DATA))
	if currentHttpHeader.AppType == WasteLibrary.APPTYPE_ADMIN {
		var execSQL string = ""
		if currentHttpHeader.OpType == WasteLibrary.OPTYPE_USER {

			var currentData WasteLibrary.UserType = WasteLibrary.StringToUserType(req.FormValue(WasteLibrary.HTTP_DATA))
			WasteLibrary.LogStr("Data : " + currentData.ToString())
			if currentData.UserId != 0 {
				execSQL = fmt.Sprintf(`UPDATE public.users 
				SET user_role='%s',email='%s',user_name='%s',customer_id=%f,password='%s' 
	  			WHERE user_id=%f  
				RETURNING user_id;`,
					currentData.UserRole, currentData.Email, currentData.UserName,
					currentData.CustomerId, currentData.Password, currentData.UserId)
				WasteLibrary.LogStr(execSQL)
			} else {

				execSQL = fmt.Sprintf(`INSERT INTO public.users 
				(user_role,email,user_name,customer_id,password) 
  				VALUES ('%s','%s','%s',%f,'%s')   
  				RETURNING user_id;`,
					currentData.UserRole, currentData.Email, currentData.UserName,
					currentData.CustomerId, currentData.Password)
				WasteLibrary.LogStr(execSQL)
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

		} else {
			resultVal.Result = WasteLibrary.RESULT_FAIL
		}
	} else {
		resultVal.Result = WasteLibrary.RESULT_OK
	}
	w.Write(resultVal.ToByte())
}

func getConfigDbMain(w http.ResponseWriter, req *http.Request) {
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
	WasteLibrary.LogStr("Header : " + currentHttpHeader.ToString())
	WasteLibrary.LogStr("Data : " + req.FormValue(WasteLibrary.HTTP_DATA))
	var execSQL string = ""
	if currentHttpHeader.BaseDataType == WasteLibrary.BASETYPE_USER {

		var currentData WasteLibrary.UserType = WasteLibrary.StringToUserType(req.FormValue(WasteLibrary.HTTP_DATA))
		WasteLibrary.LogStr("Data : " + currentData.ToString())

		if currentData.UserId != 0 {
			execSQL = fmt.Sprintf(`SELECT 
			user_role,
			email,
			user_name,
			customer_id,
			password,
			create_time
			FROM public.users
				   WHERE user_id=%f ;`, currentData.UserId)
			WasteLibrary.LogStr(execSQL)
		}

		errDb := configDb.QueryRow(execSQL).Scan(&currentData.UserRole,
			&currentData.Email,
			&currentData.UserName,
			&currentData.CustomerId,
			&currentData.Password,
			&currentData.CreateTime)
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
