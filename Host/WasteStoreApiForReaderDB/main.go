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
var readerDb *sql.DB
var err error

func initStart() {

	WasteLibrary.LogStr("Successfully connected!")
	go WasteLibrary.InitLog()
}

func main() {

	initStart()

	var readerDbHost string = "waste-readerdb-cluster-ip"
	readerDbInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		readerDbHost, port, user, password, dbname)

	readerDb, err = sql.Open("postgres", readerDbInfo)
	WasteLibrary.LogErr(err)
	defer readerDb.Close()

	err = readerDb.Ping()
	WasteLibrary.LogErr(err)

	http.HandleFunc("/health", WasteLibrary.HealthHandler)
	http.HandleFunc("/readiness", WasteLibrary.ReadinessHandler)
	http.HandleFunc("/status", WasteLibrary.StatusHandler)
	http.HandleFunc("/saveReaderDbMain", saveReaderDbMain)
	http.ListenAndServe(":80", nil)
}

func saveReaderDbMain(w http.ResponseWriter, req *http.Request) {

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
	var execSQL string = ""
	if currentHttpHeader.DataType == WasteLibrary.DATATYPE_RFID_MAIN {
		var currentData WasteLibrary.RfidDeviceMainType
		currentData.StringToType(req.FormValue(WasteLibrary.HTTP_DATA))
		execSQL = currentData.InsertDataSQL()

		var deviceId int = 0
		errDb := readerDb.QueryRow(execSQL).Scan(&deviceId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.DeviceId = float64(deviceId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_RFID_BASE {
		var currentData WasteLibrary.RfidDeviceBaseType
		currentData.StringToType(req.FormValue(WasteLibrary.HTTP_DATA))
		execSQL = currentData.InsertSQL()
		var deviceId int = 0
		errDb := readerDb.QueryRow(execSQL).Scan(&deviceId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.DeviceId = float64(deviceId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_RFID_STATU {
		var currentData WasteLibrary.RfidDeviceStatuType
		currentData.StringToType(req.FormValue(WasteLibrary.HTTP_DATA))
		execSQL = currentData.InsertSQL()
		var deviceId int = 0
		errDb := readerDb.QueryRow(execSQL).Scan(&deviceId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.DeviceId = float64(deviceId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_RFID_GPS {
		var currentData WasteLibrary.RfidDeviceGpsType
		currentData.StringToType(req.FormValue(WasteLibrary.HTTP_DATA))
		execSQL = currentData.InsertSQL()
		var deviceId int = 0
		errDb := readerDb.QueryRow(execSQL).Scan(&deviceId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.DeviceId = float64(deviceId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_RFID_EMBEDED_GPS {
		var currentData WasteLibrary.RfidDeviceEmbededGpsType
		currentData.StringToType(req.FormValue(WasteLibrary.HTTP_DATA))
		execSQL = currentData.InsertSQL()
		var deviceId int = 0
		errDb := readerDb.QueryRow(execSQL).Scan(&deviceId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.DeviceId = float64(deviceId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_RFID_ALARM {

		var currentData WasteLibrary.RfidDeviceAlarmType
		currentData.StringToType(req.FormValue(WasteLibrary.HTTP_DATA))
		execSQL = currentData.InsertSQL()

		var deviceId int = 0
		errDb := readerDb.QueryRow(execSQL).Scan(&deviceId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.DeviceId = float64(deviceId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_RFID_THERM {

		var currentData WasteLibrary.RfidDeviceThermType
		currentData.StringToType(req.FormValue(WasteLibrary.HTTP_DATA))

		execSQL = currentData.InsertSQL()

		var deviceId int = 0
		errDb := readerDb.QueryRow(execSQL).Scan(&deviceId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.DeviceId = float64(deviceId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_RFID_VERSION {

		var currentData WasteLibrary.RfidDeviceVersionType
		currentData.StringToType(req.FormValue(WasteLibrary.HTTP_DATA))

		execSQL = currentData.InsertSQL()

		var deviceId int = 0
		errDb := readerDb.QueryRow(execSQL).Scan(&deviceId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.DeviceId = float64(deviceId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_RFID_WORKHOUR {

		var currentData WasteLibrary.RfidDeviceWorkHourType
		currentData.StringToType(req.FormValue(WasteLibrary.HTTP_DATA))

		execSQL = currentData.InsertSQL()

		var deviceId int = 0
		errDb := readerDb.QueryRow(execSQL).Scan(&deviceId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.DeviceId = float64(deviceId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_RFID_NOTE {

		var currentData WasteLibrary.RfidDeviceNoteType
		currentData.StringToType(req.FormValue(WasteLibrary.HTTP_DATA))

		execSQL = currentData.InsertSQL()

		var deviceId int = 0
		errDb := readerDb.QueryRow(execSQL).Scan(&deviceId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.DeviceId = float64(deviceId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_RFID_REPORT {

		var currentData WasteLibrary.RfidDeviceReportType
		currentData.StringToType(req.FormValue(WasteLibrary.HTTP_DATA))

		execSQL = currentData.InsertSQL()

		var deviceId int = 0
		errDb := readerDb.QueryRow(execSQL).Scan(&deviceId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.DeviceId = float64(deviceId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_RECY_MAIN {
		var currentData WasteLibrary.RecyDeviceMainType
		currentData.StringToType(req.FormValue(WasteLibrary.HTTP_DATA))

		execSQL = currentData.InsertDataSQL()

		var deviceId int = 0
		errDb := readerDb.QueryRow(execSQL).Scan(&deviceId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.DeviceId = float64(deviceId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_RECY_BASE {
		var currentData WasteLibrary.RecyDeviceBaseType
		currentData.StringToType(req.FormValue(WasteLibrary.HTTP_DATA))

		execSQL = currentData.InsertSQL()

		var deviceId int = 0
		errDb := readerDb.QueryRow(execSQL).Scan(&deviceId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.DeviceId = float64(deviceId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_RECY_STATU {
		var currentData WasteLibrary.RecyDeviceStatuType
		currentData.StringToType(req.FormValue(WasteLibrary.HTTP_DATA))

		execSQL = currentData.InsertSQL()

		var deviceId int = 0
		errDb := readerDb.QueryRow(execSQL).Scan(&deviceId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.DeviceId = float64(deviceId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_RECY_GPS {
		var currentData WasteLibrary.RecyDeviceGpsType
		currentData.StringToType(req.FormValue(WasteLibrary.HTTP_DATA))

		execSQL = currentData.InsertSQL()

		var deviceId int = 0
		errDb := readerDb.QueryRow(execSQL).Scan(&deviceId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.DeviceId = float64(deviceId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_RECY_ALARM {

		var currentData WasteLibrary.RecyDeviceAlarmType
		currentData.StringToType(req.FormValue(WasteLibrary.HTTP_DATA))

		execSQL = currentData.InsertSQL()

		var deviceId int = 0
		errDb := readerDb.QueryRow(execSQL).Scan(&deviceId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.DeviceId = float64(deviceId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_RECY_THERM {

		var currentData WasteLibrary.RecyDeviceThermType
		currentData.StringToType(req.FormValue(WasteLibrary.HTTP_DATA))

		execSQL = currentData.InsertSQL()

		var deviceId int = 0
		errDb := readerDb.QueryRow(execSQL).Scan(&deviceId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.DeviceId = float64(deviceId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_RECY_VERSION {

		var currentData WasteLibrary.RecyDeviceVersionType
		currentData.StringToType(req.FormValue(WasteLibrary.HTTP_DATA))
		execSQL = currentData.InsertSQL()

		var deviceId int = 0
		errDb := readerDb.QueryRow(execSQL).Scan(&deviceId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.DeviceId = float64(deviceId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_RECY_DETAIL {

		var currentData WasteLibrary.RecyDeviceDetailType
		currentData.StringToType(req.FormValue(WasteLibrary.HTTP_DATA))

		execSQL = currentData.InsertSQL()

		var deviceId int = 0
		errDb := readerDb.QueryRow(execSQL).Scan(&deviceId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.DeviceId = float64(deviceId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_RECY_NOTE {
		var currentData WasteLibrary.RecyDeviceNoteType
		currentData.StringToType(req.FormValue(WasteLibrary.HTTP_DATA))

		execSQL = currentData.InsertSQL()

		var deviceId int = 0
		errDb := readerDb.QueryRow(execSQL).Scan(&deviceId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.DeviceId = float64(deviceId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_ULT_MAIN {
		var currentData WasteLibrary.UltDeviceMainType
		currentData.StringToType(req.FormValue(WasteLibrary.HTTP_DATA))

		execSQL = currentData.InsertDataSQL()

		var deviceId int = 0
		errDb := readerDb.QueryRow(execSQL).Scan(&deviceId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.DeviceId = float64(deviceId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_ULT_BASE {
		var currentData WasteLibrary.UltDeviceBaseType
		currentData.StringToType(req.FormValue(WasteLibrary.HTTP_DATA))

		execSQL = currentData.InsertSQL()

		var deviceId int = 0
		errDb := readerDb.QueryRow(execSQL).Scan(&deviceId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.DeviceId = float64(deviceId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_ULT_STATU {
		var currentData WasteLibrary.UltDeviceStatuType
		currentData.StringToType(req.FormValue(WasteLibrary.HTTP_DATA))

		execSQL = currentData.InsertSQL()

		var deviceId int = 0
		errDb := readerDb.QueryRow(execSQL).Scan(&deviceId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.DeviceId = float64(deviceId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_ULT_GPS {
		var currentData WasteLibrary.UltDeviceGpsType
		currentData.StringToType(req.FormValue(WasteLibrary.HTTP_DATA))

		execSQL = currentData.InsertSQL()

		var deviceId int = 0
		errDb := readerDb.QueryRow(execSQL).Scan(&deviceId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.DeviceId = float64(deviceId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_ULT_ALARM {

		var currentData WasteLibrary.UltDeviceAlarmType
		currentData.StringToType(req.FormValue(WasteLibrary.HTTP_DATA))

		execSQL = currentData.InsertSQL()

		var deviceId int = 0
		errDb := readerDb.QueryRow(execSQL).Scan(&deviceId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.DeviceId = float64(deviceId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_ULT_THERM {

		var currentData WasteLibrary.UltDeviceThermType
		currentData.StringToType(req.FormValue(WasteLibrary.HTTP_DATA))

		execSQL = currentData.InsertSQL()

		var deviceId int = 0
		errDb := readerDb.QueryRow(execSQL).Scan(&deviceId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.DeviceId = float64(deviceId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_ULT_VERSION {

		var currentData WasteLibrary.UltDeviceVersionType
		currentData.StringToType(req.FormValue(WasteLibrary.HTTP_DATA))

		execSQL = currentData.InsertSQL()

		var deviceId int = 0
		errDb := readerDb.QueryRow(execSQL).Scan(&deviceId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.DeviceId = float64(deviceId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_ULT_BATTERY {

		var currentData WasteLibrary.UltDeviceBatteryType
		currentData.StringToType(req.FormValue(WasteLibrary.HTTP_DATA))

		execSQL = currentData.InsertSQL()

		var deviceId int = 0
		errDb := readerDb.QueryRow(execSQL).Scan(&deviceId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.DeviceId = float64(deviceId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_ULT_SENS {

		var currentData WasteLibrary.UltDeviceSensType
		currentData.StringToType(req.FormValue(WasteLibrary.HTTP_DATA))

		execSQL = currentData.InsertSQL()

		var deviceId int = 0
		errDb := readerDb.QueryRow(execSQL).Scan(&deviceId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.DeviceId = float64(deviceId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_ULT_NOTE {
		var currentData WasteLibrary.UltDeviceNoteType
		currentData.StringToType(req.FormValue(WasteLibrary.HTTP_DATA))

		execSQL = currentData.InsertSQL()

		var deviceId int = 0
		errDb := readerDb.QueryRow(execSQL).Scan(&deviceId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.DeviceId = float64(deviceId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_ULT_SIM {
		var currentData WasteLibrary.UltDeviceSimType
		currentData.StringToType(req.FormValue(WasteLibrary.HTTP_DATA))

		execSQL = currentData.InsertSQL()

		var deviceId int = 0
		errDb := readerDb.QueryRow(execSQL).Scan(&deviceId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.DeviceId = float64(deviceId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_TAG_MAIN {
		var currentData WasteLibrary.TagMainType
		currentData.StringToType(req.FormValue(WasteLibrary.HTTP_DATA))

		execSQL = currentData.InsertDataSQL()

		var tagId int = 0
		errDb := readerDb.QueryRow(execSQL).Scan(&tagId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.TagId = float64(tagId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_TAG_BASE {
		var currentData WasteLibrary.TagBaseType
		currentData.StringToType(req.FormValue(WasteLibrary.HTTP_DATA))

		execSQL = currentData.InsertSQL()

		var tagId int = 0
		errDb := readerDb.QueryRow(execSQL).Scan(&tagId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.TagId = float64(tagId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_TAG_STATU {
		var currentData WasteLibrary.TagStatuType
		currentData.StringToType(req.FormValue(WasteLibrary.HTTP_DATA))

		execSQL = currentData.InsertSQL()

		var tagId int = 0
		errDb := readerDb.QueryRow(execSQL).Scan(&tagId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.TagId = float64(tagId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_TAG_GPS {
		var currentData WasteLibrary.TagGpsType
		currentData.StringToType(req.FormValue(WasteLibrary.HTTP_DATA))

		execSQL = currentData.InsertSQL()

		var tagId int = 0
		errDb := readerDb.QueryRow(execSQL).Scan(&tagId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.TagId = float64(tagId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_TAG_READER {

		var currentData WasteLibrary.TagReaderType
		currentData.StringToType(req.FormValue(WasteLibrary.HTTP_DATA))

		execSQL = currentData.InsertSQL()

		var tagId int = 0
		errDb := readerDb.QueryRow(execSQL).Scan(&tagId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.TagId = float64(tagId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_TAG_ALARM {

		var currentData WasteLibrary.TagAlarmType
		currentData.StringToType(req.FormValue(WasteLibrary.HTTP_DATA))

		execSQL = currentData.InsertSQL()

		var tagId int = 0
		errDb := readerDb.QueryRow(execSQL).Scan(&tagId)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		currentData.TagId = float64(tagId)
		resultVal.Retval = currentData.ToIdString()

	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_TAG_NOTE {

		var currentData WasteLibrary.TagNoteType
		currentData.StringToType(req.FormValue(WasteLibrary.HTTP_DATA))

		execSQL = currentData.InsertSQL()

		var tagId int = 0
		errDb := readerDb.QueryRow(execSQL).Scan(&tagId)
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
