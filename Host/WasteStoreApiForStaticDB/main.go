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
var staticDb *sql.DB
var err error

func initStart() {

	WasteLibrary.LogStr("Successfully connected!")
	go WasteLibrary.InitLog()
}

func main() {

	initStart()

	var staticDbHost string = "waste-staticdb-cluster-ip"
	staticDbInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		staticDbHost, port, user, password, dbname)

	staticDb, err = sql.Open("postgres", staticDbInfo)
	WasteLibrary.LogErr(err)
	defer staticDb.Close()

	err = staticDb.Ping()
	WasteLibrary.LogErr(err)

	http.HandleFunc("/health", WasteLibrary.HealthHandler)
	http.HandleFunc("/readiness", WasteLibrary.ReadinessHandler)
	http.HandleFunc("/status", WasteLibrary.StatusHandler)
	http.HandleFunc("/saveStaticDbMain", saveStaticDbMain)
	http.HandleFunc("/getStaticDbMain", getStaticDbMain)
	http.ListenAndServe(":80", nil)
}

func saveStaticDbMain(w http.ResponseWriter, req *http.Request) {

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

	var currentHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.StringToHttpClientHeaderType(req.FormValue(WasteLibrary.HTTP_HEADER))

	var execSQL string = ""
	if currentHttpHeader.DataType == WasteLibrary.DATATYPE_RFID_MAIN {
		var currentData WasteLibrary.RfidDeviceMainType = WasteLibrary.StringToRfidDeviceMainType(req.FormValue(WasteLibrary.HTTP_DATA))
		if currentData.DeviceId != 0 {
			execSQL = currentData.UpdateSQL()
		} else {
			execSQL = currentData.InsertSQL()
		}
		var deviceId int = 0
		errDb := staticDb.QueryRow(execSQL).Scan(&deviceId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.DeviceId = float64(deviceId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_RFID_BASE {
		var currentData WasteLibrary.RfidDeviceBaseType = WasteLibrary.StringToRfidDeviceBaseType(req.FormValue(WasteLibrary.HTTP_DATA))
		if !currentData.NewData {
			execSQL = currentData.UpdateSQL()
		} else {
			execSQL = currentData.InsertSQL()
		}
		var deviceId int = 0
		errDb := staticDb.QueryRow(execSQL).Scan(&deviceId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.DeviceId = float64(deviceId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_RFID_STATU {
		var currentData WasteLibrary.RfidDeviceStatuType = WasteLibrary.StringToRfidDeviceStatuType(req.FormValue(WasteLibrary.HTTP_DATA))
		if !currentData.NewData {
			execSQL = currentData.UpdateSQL()
		} else {
			execSQL = currentData.InsertSQL()
		}
		var deviceId int = 0
		errDb := staticDb.QueryRow(execSQL).Scan(&deviceId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.DeviceId = float64(deviceId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_RFID_GPS {
		var currentData WasteLibrary.RfidDeviceGpsType = WasteLibrary.StringToRfidDeviceGpsType(req.FormValue(WasteLibrary.HTTP_DATA))
		if !currentData.NewData {
			execSQL = currentData.UpdateSQL()
		} else {
			execSQL = currentData.InsertSQL()
		}
		var deviceId int = 0
		errDb := staticDb.QueryRow(execSQL).Scan(&deviceId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.DeviceId = float64(deviceId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_RFID_EMBEDED_GPS {
		var currentData WasteLibrary.RfidDeviceEmbededGpsType = WasteLibrary.StringToRfidDeviceEmbededGpsType(req.FormValue(WasteLibrary.HTTP_DATA))
		if !currentData.NewData {
			execSQL = currentData.UpdateSQL()
		} else {
			execSQL = currentData.InsertSQL()
		}
		var deviceId int = 0
		errDb := staticDb.QueryRow(execSQL).Scan(&deviceId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.DeviceId = float64(deviceId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_RFID_ALARM {

		var currentData WasteLibrary.RfidDeviceAlarmType = WasteLibrary.StringToRfidDeviceAlarmType(req.FormValue(WasteLibrary.HTTP_DATA))
		if !currentData.NewData {
			execSQL = currentData.UpdateSQL()
		} else {
			execSQL = currentData.InsertSQL()
		}
		var deviceId int = 0
		errDb := staticDb.QueryRow(execSQL).Scan(&deviceId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.DeviceId = float64(deviceId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_RFID_THERM {

		var currentData WasteLibrary.RfidDeviceThermType = WasteLibrary.StringToRfidDeviceThermType(req.FormValue(WasteLibrary.HTTP_DATA))
		if !currentData.NewData {
			execSQL = currentData.UpdateSQL()
		} else {
			execSQL = currentData.InsertSQL()
		}
		var deviceId int = 0
		errDb := staticDb.QueryRow(execSQL).Scan(&deviceId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.DeviceId = float64(deviceId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_RFID_VERSION {

		var currentData WasteLibrary.RfidDeviceVersionType = WasteLibrary.StringToRfidDeviceVersionType(req.FormValue(WasteLibrary.HTTP_DATA))
		if !currentData.NewData {
			execSQL = currentData.UpdateSQL()
		} else {
			execSQL = currentData.InsertSQL()
		}
		var deviceId int = 0
		errDb := staticDb.QueryRow(execSQL).Scan(&deviceId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.DeviceId = float64(deviceId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_RFID_DETAIL {

		var currentData WasteLibrary.RfidDeviceDetailType = WasteLibrary.StringToRfidDeviceDetailType(req.FormValue(WasteLibrary.HTTP_DATA))
		if !currentData.NewData {
			execSQL = currentData.UpdateSQL()
		} else {
			execSQL = currentData.InsertSQL()
		}
		var deviceId int = 0
		errDb := staticDb.QueryRow(execSQL).Scan(&deviceId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.DeviceId = float64(deviceId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_RFID_WORKHOUR {

		var currentData WasteLibrary.RfidDeviceWorkHourType = WasteLibrary.StringToRfidDeviceWorkHourType(req.FormValue(WasteLibrary.HTTP_DATA))
		if !currentData.NewData {
			execSQL = currentData.UpdateSQL()
		} else {
			execSQL = currentData.InsertSQL()
		}
		var deviceId int = 0
		errDb := staticDb.QueryRow(execSQL).Scan(&deviceId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.DeviceId = float64(deviceId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_RFID_REPORT {

		var currentData WasteLibrary.RfidDeviceReportType = WasteLibrary.StringToRfidDeviceReportType(req.FormValue(WasteLibrary.HTTP_DATA))
		if !currentData.NewData {
			execSQL = currentData.UpdateSQL()
		} else {
			execSQL = currentData.InsertSQL()
		}
		var deviceId int = 0
		errDb := staticDb.QueryRow(execSQL).Scan(&deviceId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.DeviceId = float64(deviceId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_RFID_NOTE {

		var currentData WasteLibrary.RfidDeviceNoteType = WasteLibrary.StringToRfidDeviceNoteType(req.FormValue(WasteLibrary.HTTP_DATA))
		if !currentData.NewData {
			execSQL = currentData.UpdateSQL()
		} else {
			execSQL = currentData.InsertSQL()
		}
		var deviceId int = 0
		errDb := staticDb.QueryRow(execSQL).Scan(&deviceId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.DeviceId = float64(deviceId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_RECY_MAIN {
		var currentData WasteLibrary.RecyDeviceMainType = WasteLibrary.StringToRecyDeviceMainType(req.FormValue(WasteLibrary.HTTP_DATA))
		if currentData.DeviceId != 0 {
			execSQL = currentData.UpdateSQL()
		} else {
			execSQL = currentData.InsertSQL()
		}
		var deviceId int = 0
		errDb := staticDb.QueryRow(execSQL).Scan(&deviceId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.DeviceId = float64(deviceId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_RECY_BASE {
		var currentData WasteLibrary.RecyDeviceBaseType = WasteLibrary.StringToRecyDeviceBaseType(req.FormValue(WasteLibrary.HTTP_DATA))
		if !currentData.NewData {
			execSQL = currentData.UpdateSQL()
		} else {
			execSQL = currentData.InsertSQL()
		}
		var deviceId int = 0
		errDb := staticDb.QueryRow(execSQL).Scan(&deviceId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.DeviceId = float64(deviceId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_RECY_STATU {
		var currentData WasteLibrary.RecyDeviceStatuType = WasteLibrary.StringToRecyDeviceStatuType(req.FormValue(WasteLibrary.HTTP_DATA))
		if !currentData.NewData {
			execSQL = currentData.UpdateSQL()
		} else {
			execSQL = currentData.InsertSQL()
		}
		var deviceId int = 0
		errDb := staticDb.QueryRow(execSQL).Scan(&deviceId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.DeviceId = float64(deviceId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_RECY_GPS {
		var currentData WasteLibrary.RecyDeviceGpsType = WasteLibrary.StringToRecyDeviceGpsType(req.FormValue(WasteLibrary.HTTP_DATA))
		if !currentData.NewData {
			execSQL = currentData.UpdateSQL()
		} else {
			execSQL = currentData.InsertSQL()
		}
		var deviceId int = 0
		errDb := staticDb.QueryRow(execSQL).Scan(&deviceId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.DeviceId = float64(deviceId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_RECY_ALARM {

		var currentData WasteLibrary.RecyDeviceAlarmType = WasteLibrary.StringToRecyDeviceAlarmType(req.FormValue(WasteLibrary.HTTP_DATA))
		if !currentData.NewData {
			execSQL = currentData.UpdateSQL()
		} else {
			execSQL = currentData.InsertSQL()
		}
		var deviceId int = 0
		errDb := staticDb.QueryRow(execSQL).Scan(&deviceId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.DeviceId = float64(deviceId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_RECY_THERM {

		var currentData WasteLibrary.RecyDeviceThermType = WasteLibrary.StringToRecyDeviceThermType(req.FormValue(WasteLibrary.HTTP_DATA))
		if !currentData.NewData {
			execSQL = currentData.UpdateSQL()
		} else {
			execSQL = currentData.InsertSQL()
		}
		var deviceId int = 0
		errDb := staticDb.QueryRow(execSQL).Scan(&deviceId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.DeviceId = float64(deviceId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_RECY_VERSION {

		var currentData WasteLibrary.RecyDeviceVersionType = WasteLibrary.StringToRecyDeviceVersionType(req.FormValue(WasteLibrary.HTTP_DATA))
		if !currentData.NewData {
			execSQL = currentData.UpdateSQL()
		} else {
			execSQL = currentData.InsertSQL()
		}
		var deviceId int = 0
		errDb := staticDb.QueryRow(execSQL).Scan(&deviceId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.DeviceId = float64(deviceId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_RECY_DETAIL {

		var currentData WasteLibrary.RecyDeviceDetailType = WasteLibrary.StringToRecyDeviceDetailType(req.FormValue(WasteLibrary.HTTP_DATA))
		if !currentData.NewData {
			execSQL = currentData.UpdateSQL()
		} else {
			execSQL = currentData.InsertSQL()
		}
		var deviceId int = 0
		errDb := staticDb.QueryRow(execSQL).Scan(&deviceId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.DeviceId = float64(deviceId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_RECY_NOTE {
		var currentData WasteLibrary.RecyDeviceNoteType = WasteLibrary.StringToRecyDeviceNoteType(req.FormValue(WasteLibrary.HTTP_DATA))
		if !currentData.NewData {
			execSQL = currentData.UpdateSQL()
		} else {
			execSQL = currentData.InsertSQL()
		}
		var deviceId int = 0
		errDb := staticDb.QueryRow(execSQL).Scan(&deviceId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.DeviceId = float64(deviceId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_ULT_MAIN {
		var currentData WasteLibrary.UltDeviceMainType = WasteLibrary.StringToUltDeviceMainType(req.FormValue(WasteLibrary.HTTP_DATA))
		if currentData.DeviceId != 0 {
			execSQL = currentData.UpdateSQL()
		} else {
			execSQL = currentData.InsertSQL()
		}
		var deviceId int = 0
		errDb := staticDb.QueryRow(execSQL).Scan(&deviceId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.DeviceId = float64(deviceId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_ULT_BASE {
		var currentData WasteLibrary.UltDeviceBaseType = WasteLibrary.StringToUltDeviceBaseType(req.FormValue(WasteLibrary.HTTP_DATA))
		if !currentData.NewData {
			execSQL = currentData.UpdateSQL()
		} else {
			execSQL = currentData.InsertSQL()
		}
		var deviceId int = 0
		errDb := staticDb.QueryRow(execSQL).Scan(&deviceId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.DeviceId = float64(deviceId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_ULT_STATU {
		var currentData WasteLibrary.UltDeviceStatuType = WasteLibrary.StringToUltDeviceStatuType(req.FormValue(WasteLibrary.HTTP_DATA))
		if !currentData.NewData {
			execSQL = currentData.UpdateSQL()
		} else {
			execSQL = currentData.InsertSQL()
		}
		var deviceId int = 0
		errDb := staticDb.QueryRow(execSQL).Scan(&deviceId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.DeviceId = float64(deviceId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_ULT_GPS {
		var currentData WasteLibrary.UltDeviceGpsType = WasteLibrary.StringToUltDeviceGpsType(req.FormValue(WasteLibrary.HTTP_DATA))
		if !currentData.NewData {
			execSQL = currentData.UpdateSQL()
		} else {
			execSQL = currentData.InsertSQL()
		}
		var deviceId int = 0
		errDb := staticDb.QueryRow(execSQL).Scan(&deviceId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.DeviceId = float64(deviceId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_ULT_ALARM {

		var currentData WasteLibrary.UltDeviceAlarmType = WasteLibrary.StringToUltDeviceAlarmType(req.FormValue(WasteLibrary.HTTP_DATA))
		if !currentData.NewData {
			execSQL = currentData.UpdateSQL()
		} else {
			execSQL = currentData.InsertSQL()
		}
		var deviceId int = 0
		errDb := staticDb.QueryRow(execSQL).Scan(&deviceId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.DeviceId = float64(deviceId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_ULT_THERM {

		var currentData WasteLibrary.UltDeviceThermType = WasteLibrary.StringToUltDeviceThermType(req.FormValue(WasteLibrary.HTTP_DATA))
		if !currentData.NewData {
			execSQL = currentData.UpdateSQL()
		} else {
			execSQL = currentData.InsertSQL()
		}
		var deviceId int = 0
		errDb := staticDb.QueryRow(execSQL).Scan(&deviceId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.DeviceId = float64(deviceId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_ULT_VERSION {

		var currentData WasteLibrary.UltDeviceVersionType = WasteLibrary.StringToUltDeviceVersionType(req.FormValue(WasteLibrary.HTTP_DATA))
		if !currentData.NewData {
			execSQL = currentData.UpdateSQL()
		} else {
			execSQL = currentData.InsertSQL()
		}
		var deviceId int = 0
		errDb := staticDb.QueryRow(execSQL).Scan(&deviceId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.DeviceId = float64(deviceId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_ULT_BATTERY {

		var currentData WasteLibrary.UltDeviceBatteryType = WasteLibrary.StringToUltDeviceBatteryType(req.FormValue(WasteLibrary.HTTP_DATA))
		if !currentData.NewData {
			execSQL = currentData.UpdateSQL()
		} else {
			execSQL = currentData.InsertSQL()
		}
		var deviceId int = 0
		errDb := staticDb.QueryRow(execSQL).Scan(&deviceId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.DeviceId = float64(deviceId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_ULT_SENS {

		var currentData WasteLibrary.UltDeviceSensType = WasteLibrary.StringToUltDeviceSensType(req.FormValue(WasteLibrary.HTTP_DATA))
		if !currentData.NewData {
			execSQL = currentData.UpdateSQL()
		} else {
			execSQL = currentData.InsertSQL()
		}
		var deviceId int = 0
		errDb := staticDb.QueryRow(execSQL).Scan(&deviceId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.DeviceId = float64(deviceId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_ULT_NOTE {

		var currentData WasteLibrary.UltDeviceNoteType = WasteLibrary.StringToUltDeviceNoteType(req.FormValue(WasteLibrary.HTTP_DATA))
		if !currentData.NewData {
			execSQL = currentData.UpdateSQL()
		} else {
			execSQL = currentData.InsertSQL()
		}
		var deviceId int = 0
		errDb := staticDb.QueryRow(execSQL).Scan(&deviceId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.DeviceId = float64(deviceId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_TAG_MAIN {
		var currentData WasteLibrary.TagMainType = WasteLibrary.StringToTagMainType(req.FormValue(WasteLibrary.HTTP_DATA))
		if currentData.TagId != 0 {
			execSQL = currentData.UpdateSQL()
		} else {
			execSQL = currentData.InsertSQL()
		}
		var tagId int = 0
		errDb := staticDb.QueryRow(execSQL).Scan(&tagId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.TagId = float64(tagId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_TAG_BASE {
		var currentData WasteLibrary.TagBaseType = WasteLibrary.StringToTagBaseType(req.FormValue(WasteLibrary.HTTP_DATA))
		if !currentData.NewData {
			execSQL = currentData.UpdateSQL()
		} else {
			execSQL = currentData.InsertSQL()
		}
		var tagId int = 0
		errDb := staticDb.QueryRow(execSQL).Scan(&tagId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.TagId = float64(tagId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_TAG_STATU {
		var currentData WasteLibrary.TagStatuType = WasteLibrary.StringToTagStatuType(req.FormValue(WasteLibrary.HTTP_DATA))
		if !currentData.NewData {
			execSQL = currentData.UpdateSQL()
		} else {
			execSQL = currentData.InsertSQL()
		}
		var tagId int = 0
		errDb := staticDb.QueryRow(execSQL).Scan(&tagId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.TagId = float64(tagId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_TAG_GPS {
		var currentData WasteLibrary.TagGpsType = WasteLibrary.StringToTagGpsType(req.FormValue(WasteLibrary.HTTP_DATA))
		if !currentData.NewData {
			execSQL = currentData.UpdateSQL()
		} else {
			execSQL = currentData.InsertSQL()
		}
		var tagId int = 0
		errDb := staticDb.QueryRow(execSQL).Scan(&tagId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.TagId = float64(tagId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_TAG_READER {

		var currentData WasteLibrary.TagReaderType = WasteLibrary.StringToTagReaderType(req.FormValue(WasteLibrary.HTTP_DATA))
		if !currentData.NewData {
			execSQL = currentData.UpdateSQL()
		} else {
			execSQL = currentData.InsertSQL()
		}
		var tagId int = 0
		errDb := staticDb.QueryRow(execSQL).Scan(&tagId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.TagId = float64(tagId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_TAG_ALARM {

		var currentData WasteLibrary.TagAlarmType = WasteLibrary.StringToTagAlarmType(req.FormValue(WasteLibrary.HTTP_DATA))
		if !currentData.NewData {
			execSQL = currentData.UpdateSQL()
		} else {
			execSQL = currentData.InsertSQL()
		}
		var tagId int = 0
		errDb := staticDb.QueryRow(execSQL).Scan(&tagId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.TagId = float64(tagId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_TAG_NOTE {

		var currentData WasteLibrary.TagNoteType = WasteLibrary.StringToTagNoteType(req.FormValue(WasteLibrary.HTTP_DATA))
		if !currentData.NewData {
			execSQL = currentData.UpdateSQL()
		} else {
			execSQL = currentData.InsertSQL()
		}
		var tagId int = 0
		errDb := staticDb.QueryRow(execSQL).Scan(&tagId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.TagId = float64(tagId)
		resultVal.Retval = currentData.ToIdString()

	} else {
		resultVal.Result = WasteLibrary.RESULT_FAIL
	}

	w.Write(resultVal.ToByte())

}

func getStaticDbMain(w http.ResponseWriter, req *http.Request) {

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
	var currentHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.StringToHttpClientHeaderType(req.FormValue(WasteLibrary.HTTP_HEADER))
	if currentHttpHeader.DataType == WasteLibrary.DATATYPE_RFID_MAIN {
		var currentData WasteLibrary.RfidDeviceMainType = WasteLibrary.StringToRfidDeviceMainType(req.FormValue(WasteLibrary.HTTP_DATA))
		errDb := currentData.SelectWithDb(staticDb)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		resultVal.Retval = currentData.ToString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_RFID_BASE {
		var currentData WasteLibrary.RfidDeviceBaseType = WasteLibrary.StringToRfidDeviceBaseType(req.FormValue(WasteLibrary.HTTP_DATA))
		errDb := currentData.SelectWithDb(staticDb)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		resultVal.Retval = currentData.ToString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_RFID_STATU {
		var currentData WasteLibrary.RfidDeviceStatuType = WasteLibrary.StringToRfidDeviceStatuType(req.FormValue(WasteLibrary.HTTP_DATA))
		errDb := currentData.SelectWithDb(staticDb)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		resultVal.Retval = currentData.ToString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_RFID_GPS {
		var currentData WasteLibrary.RfidDeviceGpsType = WasteLibrary.StringToRfidDeviceGpsType(req.FormValue(WasteLibrary.HTTP_DATA))
		errDb := currentData.SelectWithDb(staticDb)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		resultVal.Retval = currentData.ToString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_RFID_EMBEDED_GPS {
		var currentData WasteLibrary.RfidDeviceEmbededGpsType = WasteLibrary.StringToRfidDeviceEmbededGpsType(req.FormValue(WasteLibrary.HTTP_DATA))
		errDb := currentData.SelectWithDb(staticDb)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		resultVal.Retval = currentData.ToString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_RFID_ALARM {
		var currentData WasteLibrary.RfidDeviceAlarmType = WasteLibrary.StringToRfidDeviceAlarmType(req.FormValue(WasteLibrary.HTTP_DATA))
		errDb := currentData.SelectWithDb(staticDb)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		resultVal.Retval = currentData.ToString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_RFID_THERM {
		var currentData WasteLibrary.RfidDeviceThermType = WasteLibrary.StringToRfidDeviceThermType(req.FormValue(WasteLibrary.HTTP_DATA))
		errDb := currentData.SelectWithDb(staticDb)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		resultVal.Retval = currentData.ToString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_RFID_VERSION {
		var currentData WasteLibrary.RfidDeviceVersionType = WasteLibrary.StringToRfidDeviceVersionType(req.FormValue(WasteLibrary.HTTP_DATA))
		errDb := currentData.SelectWithDb(staticDb)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		resultVal.Retval = currentData.ToString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_RFID_DETAIL {
		var currentData WasteLibrary.RfidDeviceDetailType = WasteLibrary.StringToRfidDeviceDetailType(req.FormValue(WasteLibrary.HTTP_DATA))
		errDb := currentData.SelectWithDb(staticDb)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		resultVal.Retval = currentData.ToString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_RFID_WORKHOUR {
		var currentData WasteLibrary.RfidDeviceWorkHourType = WasteLibrary.StringToRfidDeviceWorkHourType(req.FormValue(WasteLibrary.HTTP_DATA))
		errDb := currentData.SelectWithDb(staticDb)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		resultVal.Retval = currentData.ToString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_RFID_NOTE {
		var currentData WasteLibrary.RfidDeviceNoteType = WasteLibrary.StringToRfidDeviceNoteType(req.FormValue(WasteLibrary.HTTP_DATA))
		errDb := currentData.SelectWithDb(staticDb)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		resultVal.Retval = currentData.ToString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_RFID_REPORT {
		var currentData WasteLibrary.RfidDeviceReportType = WasteLibrary.StringToRfidDeviceReportType(req.FormValue(WasteLibrary.HTTP_DATA))
		errDb := currentData.SelectWithDb(staticDb)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		resultVal.Retval = currentData.ToString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_RECY_MAIN {
		var currentData WasteLibrary.RecyDeviceMainType = WasteLibrary.StringToRecyDeviceMainType(req.FormValue(WasteLibrary.HTTP_DATA))
		errDb := currentData.SelectWithDb(staticDb)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		resultVal.Retval = currentData.ToString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_RECY_BASE {
		var currentData WasteLibrary.RecyDeviceBaseType = WasteLibrary.StringToRecyDeviceBaseType(req.FormValue(WasteLibrary.HTTP_DATA))
		errDb := currentData.SelectWithDb(staticDb)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		resultVal.Retval = currentData.ToString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_RECY_STATU {
		var currentData WasteLibrary.RecyDeviceStatuType = WasteLibrary.StringToRecyDeviceStatuType(req.FormValue(WasteLibrary.HTTP_DATA))
		errDb := currentData.SelectWithDb(staticDb)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		resultVal.Retval = currentData.ToString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_RECY_GPS {
		var currentData WasteLibrary.RecyDeviceGpsType = WasteLibrary.StringToRecyDeviceGpsType(req.FormValue(WasteLibrary.HTTP_DATA))
		errDb := currentData.SelectWithDb(staticDb)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		resultVal.Retval = currentData.ToString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_RECY_ALARM {
		var currentData WasteLibrary.RecyDeviceAlarmType = WasteLibrary.StringToRecyDeviceAlarmType(req.FormValue(WasteLibrary.HTTP_DATA))
		errDb := currentData.SelectWithDb(staticDb)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		resultVal.Retval = currentData.ToString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_RECY_THERM {
		var currentData WasteLibrary.RecyDeviceThermType = WasteLibrary.StringToRecyDeviceThermType(req.FormValue(WasteLibrary.HTTP_DATA))
		errDb := currentData.SelectWithDb(staticDb)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		resultVal.Retval = currentData.ToString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_RECY_VERSION {
		var currentData WasteLibrary.RecyDeviceVersionType = WasteLibrary.StringToRecyDeviceVersionType(req.FormValue(WasteLibrary.HTTP_DATA))
		errDb := currentData.SelectWithDb(staticDb)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		resultVal.Retval = currentData.ToString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_RECY_DETAIL {
		var currentData WasteLibrary.RecyDeviceDetailType = WasteLibrary.StringToRecyDeviceDetailType(req.FormValue(WasteLibrary.HTTP_DATA))
		errDb := currentData.SelectWithDb(staticDb)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		resultVal.Retval = currentData.ToString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_RECY_NOTE {
		var currentData WasteLibrary.RecyDeviceNoteType = WasteLibrary.StringToRecyDeviceNoteType(req.FormValue(WasteLibrary.HTTP_DATA))
		errDb := currentData.SelectWithDb(staticDb)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		resultVal.Retval = currentData.ToString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_ULT_MAIN {
		var currentData WasteLibrary.UltDeviceMainType = WasteLibrary.StringToUltDeviceMainType(req.FormValue(WasteLibrary.HTTP_DATA))
		errDb := currentData.SelectWithDb(staticDb)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		resultVal.Retval = currentData.ToString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_ULT_BASE {
		var currentData WasteLibrary.UltDeviceBaseType = WasteLibrary.StringToUltDeviceBaseType(req.FormValue(WasteLibrary.HTTP_DATA))
		errDb := currentData.SelectWithDb(staticDb)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		resultVal.Retval = currentData.ToString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_ULT_STATU {
		var currentData WasteLibrary.UltDeviceStatuType = WasteLibrary.StringToUltDeviceStatuType(req.FormValue(WasteLibrary.HTTP_DATA))
		errDb := currentData.SelectWithDb(staticDb)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		resultVal.Retval = currentData.ToString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_ULT_GPS {
		var currentData WasteLibrary.UltDeviceGpsType = WasteLibrary.StringToUltDeviceGpsType(req.FormValue(WasteLibrary.HTTP_DATA))
		errDb := currentData.SelectWithDb(staticDb)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		resultVal.Retval = currentData.ToString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_ULT_ALARM {
		var currentData WasteLibrary.UltDeviceAlarmType = WasteLibrary.StringToUltDeviceAlarmType(req.FormValue(WasteLibrary.HTTP_DATA))
		errDb := currentData.SelectWithDb(staticDb)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		resultVal.Retval = currentData.ToString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_ULT_THERM {
		var currentData WasteLibrary.UltDeviceThermType = WasteLibrary.StringToUltDeviceThermType(req.FormValue(WasteLibrary.HTTP_DATA))
		errDb := currentData.SelectWithDb(staticDb)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		resultVal.Retval = currentData.ToString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_ULT_VERSION {
		var currentData WasteLibrary.UltDeviceVersionType = WasteLibrary.StringToUltDeviceVersionType(req.FormValue(WasteLibrary.HTTP_DATA))
		errDb := currentData.SelectWithDb(staticDb)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		resultVal.Retval = currentData.ToString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_ULT_BATTERY {
		var currentData WasteLibrary.UltDeviceBatteryType = WasteLibrary.StringToUltDeviceBatteryType(req.FormValue(WasteLibrary.HTTP_DATA))
		errDb := currentData.SelectWithDb(staticDb)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		resultVal.Retval = currentData.ToString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_ULT_SENS {
		var currentData WasteLibrary.UltDeviceSensType = WasteLibrary.StringToUltDeviceSensType(req.FormValue(WasteLibrary.HTTP_DATA))
		errDb := currentData.SelectWithDb(staticDb)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		resultVal.Retval = currentData.ToString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_ULT_NOTE {
		var currentData WasteLibrary.UltDeviceNoteType = WasteLibrary.StringToUltDeviceNoteType(req.FormValue(WasteLibrary.HTTP_DATA))
		errDb := currentData.SelectWithDb(staticDb)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		resultVal.Retval = currentData.ToString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_TAG_MAIN {
		var currentData WasteLibrary.TagMainType = WasteLibrary.StringToTagMainType(req.FormValue(WasteLibrary.HTTP_DATA))
		errDb := currentData.SelectWithDb(staticDb)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		resultVal.Retval = currentData.ToString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_TAG_BASE {
		var currentData WasteLibrary.TagBaseType = WasteLibrary.StringToTagBaseType(req.FormValue(WasteLibrary.HTTP_DATA))
		errDb := currentData.SelectWithDb(staticDb)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		resultVal.Retval = currentData.ToString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_TAG_STATU {
		var currentData WasteLibrary.TagStatuType = WasteLibrary.StringToTagStatuType(req.FormValue(WasteLibrary.HTTP_DATA))
		errDb := currentData.SelectWithDb(staticDb)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		resultVal.Retval = currentData.ToString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_TAG_GPS {
		var currentData WasteLibrary.TagGpsType = WasteLibrary.StringToTagGpsType(req.FormValue(WasteLibrary.HTTP_DATA))
		errDb := currentData.SelectWithDb(staticDb)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		resultVal.Retval = currentData.ToString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_TAG_READER {
		var currentData WasteLibrary.TagReaderType = WasteLibrary.StringToTagReaderType(req.FormValue(WasteLibrary.HTTP_DATA))
		errDb := currentData.SelectWithDb(staticDb)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		resultVal.Retval = currentData.ToString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_TAG_ALARM {
		var currentData WasteLibrary.TagAlarmType = WasteLibrary.StringToTagAlarmType(req.FormValue(WasteLibrary.HTTP_DATA))
		errDb := currentData.SelectWithDb(staticDb)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		resultVal.Retval = currentData.ToString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_TAG_NOTE {
		var currentData WasteLibrary.TagNoteType = WasteLibrary.StringToTagNoteType(req.FormValue(WasteLibrary.HTTP_DATA))
		errDb := currentData.SelectWithDb(staticDb)
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
