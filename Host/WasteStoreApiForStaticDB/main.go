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

	WasteLibrary.LogStr("Start")
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
	WasteLibrary.LogStr("Header : " + currentHttpHeader.ToString())
	WasteLibrary.LogStr("Data : " + req.FormValue(WasteLibrary.HTTP_DATA))

	WasteLibrary.LogStr("AfatekApi Receive Header : " + req.FormValue(WasteLibrary.HTTP_HEADER))
	WasteLibrary.LogStr("AfatekApi Receive Data : " + req.FormValue(WasteLibrary.HTTP_DATA))
	if currentHttpHeader.AppType == WasteLibrary.APPTYPE_RFID {
		var execSQL string = ""
		if currentHttpHeader.OpType == WasteLibrary.OPTYPE_RF {
			var currentData WasteLibrary.TagType = WasteLibrary.StringToTagType(req.FormValue(WasteLibrary.HTTP_DATA))
			WasteLibrary.LogStr("Data : " + currentData.ToString())

			var selectSQL string = fmt.Sprintf(`SELECT TagID
			FROM public.tags WHERE Epc='%s' AND CustomerId=%f;`, currentData.Epc, currentData.CustomerId)
			rows, errSel := staticDb.Query(selectSQL)
			if errSel != nil {
				WasteLibrary.LogErr(errSel)
				resultVal.Result = WasteLibrary.RESULT_FAIL
			} else {
				resultVal.Result = WasteLibrary.RESULT_OK
			}
			var tagID int = 0
			for rows.Next() {
				rows.Scan(&tagID)
			}
			if tagID != 0 {

				execSQL = currentData.UpdateSQL()
				WasteLibrary.LogStr(execSQL)
			} else {

				execSQL = currentData.InsertSQL()
				WasteLibrary.LogStr(execSQL)
			}
			tagID = 0
			errDb := staticDb.QueryRow(execSQL).Scan(&tagID)
			if errDb != nil {
				WasteLibrary.LogErr(errDb)
				resultVal.Result = WasteLibrary.RESULT_FAIL
			} else {
				resultVal.Result = WasteLibrary.RESULT_OK
			}

			currentData.TagID = float64(tagID)
			resultVal.Retval = currentData.ToIdString()

		} else if currentHttpHeader.OpType == WasteLibrary.OPTYPE_GPS {

			var currentData WasteLibrary.DeviceType = WasteLibrary.StringToDeviceType(req.FormValue(WasteLibrary.HTTP_DATA))
			WasteLibrary.LogStr("Data : " + currentData.ToString())

			execSQL = currentData.UpdateGpsSQL()
			WasteLibrary.LogStr(execSQL)
			var deviceID = 0
			errDb := staticDb.QueryRow(execSQL).Scan(&deviceID)
			if errDb != nil {
				WasteLibrary.LogErr(errDb)
				resultVal.Result = WasteLibrary.RESULT_FAIL
			} else {
				resultVal.Result = WasteLibrary.RESULT_OK
			}

			currentData.DeviceId = float64(deviceID)
			resultVal.Retval = currentData.ToIdString()

		} else if currentHttpHeader.OpType == WasteLibrary.OPTYPE_ARVENTO {

			var currentData WasteLibrary.DeviceType = WasteLibrary.StringToDeviceType(req.FormValue(WasteLibrary.HTTP_DATA))
			WasteLibrary.LogStr("Data : " + currentData.ToString())

			execSQL = currentData.UpdateGpsSQL()
			WasteLibrary.LogStr(execSQL)
			var deviceID = 0
			errDb := staticDb.QueryRow(execSQL).Scan(&deviceID)
			if errDb != nil {
				WasteLibrary.LogErr(errDb)
				resultVal.Result = WasteLibrary.RESULT_FAIL
			} else {
				resultVal.Result = WasteLibrary.RESULT_OK
			}

			currentData.DeviceId = float64(deviceID)
			resultVal.Retval = currentData.ToIdString()

		} else if currentHttpHeader.OpType == WasteLibrary.OPTYPE_STATUS {

			var currentData WasteLibrary.DeviceType = WasteLibrary.StringToDeviceType(req.FormValue(WasteLibrary.HTTP_DATA))
			WasteLibrary.LogStr("Data : " + currentData.ToString())

			execSQL = currentData.UpdateStatuSQL()

			WasteLibrary.LogStr(execSQL)
			var deviceID = 0
			errDb := staticDb.QueryRow(execSQL).Scan(&deviceID)
			if errDb != nil {
				WasteLibrary.LogErr(errDb)
				resultVal.Result = WasteLibrary.RESULT_FAIL
			} else {
				resultVal.Result = WasteLibrary.RESULT_OK
			}

			currentData.DeviceId = float64(deviceID)
			resultVal.Retval = currentData.ToIdString()

		} else if currentHttpHeader.OpType == WasteLibrary.OPTYPE_THERM {

			var currentData WasteLibrary.DeviceType = WasteLibrary.StringToDeviceType(req.FormValue(WasteLibrary.HTTP_DATA))
			WasteLibrary.LogStr("Data : " + currentData.ToString())

			execSQL = currentData.UpdateThermSQL()
			WasteLibrary.LogStr(execSQL)
			var deviceID = 0
			errDb := staticDb.QueryRow(execSQL).Scan(&deviceID)
			if errDb != nil {
				WasteLibrary.LogErr(errDb)
				resultVal.Result = WasteLibrary.RESULT_FAIL
			} else {
				resultVal.Result = WasteLibrary.RESULT_OK
			}

			currentData.DeviceId = float64(deviceID)
			resultVal.Retval = currentData.ToIdString()

		} else {
			resultVal.Result = WasteLibrary.RESULT_FAIL
		}

	} else if currentHttpHeader.AppType == WasteLibrary.APPTYPE_ULT {
		resultVal.Result = WasteLibrary.RESULT_OK
	} else if currentHttpHeader.AppType == WasteLibrary.APPTYPE_RECY {
		resultVal.Result = WasteLibrary.RESULT_OK
	} else if currentHttpHeader.AppType == WasteLibrary.APPTYPE_AFATEK {
		var execSQL string = ""
		if currentHttpHeader.OpType == WasteLibrary.OPTYPE_CUSTOMER {

			var currentData WasteLibrary.CustomerType = WasteLibrary.StringToCustomerType(req.FormValue(WasteLibrary.HTTP_DATA))
			WasteLibrary.LogStr("Data : " + currentData.ToString())

			if currentData.CustomerId != 0 {
				execSQL = currentData.UpdateSQL()
				WasteLibrary.LogStr(execSQL)
			} else {

				execSQL = currentData.InsertSQL()
				WasteLibrary.LogStr(execSQL)
			}
			var customerId int = 0
			errDb := staticDb.QueryRow(execSQL).Scan(&customerId)
			if errDb != nil {
				WasteLibrary.LogErr(errDb)
				resultVal.Result = WasteLibrary.RESULT_FAIL
			} else {
				resultVal.Result = WasteLibrary.RESULT_OK
			}

			currentData.CustomerId = float64(customerId)
			resultVal.Retval = currentData.ToIdString()

		} else if currentHttpHeader.OpType == WasteLibrary.OPTYPE_DEVICE {

			var currentData WasteLibrary.DeviceType = WasteLibrary.StringToDeviceType(req.FormValue(WasteLibrary.HTTP_DATA))
			WasteLibrary.LogStr("Data : " + currentData.ToString())
			if currentData.DeviceId != 0 {
				execSQL = currentData.UpdateSQL()
				WasteLibrary.LogStr(execSQL)
			} else {

				execSQL = currentData.InsertSQL()
				WasteLibrary.LogStr(execSQL)
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

		} else {
			resultVal.Result = WasteLibrary.RESULT_FAIL
		}
	} else {
		resultVal.Result = WasteLibrary.RESULT_OK
	}
	w.Write(resultVal.ToByte())
}

func getStaticDbMain(w http.ResponseWriter, req *http.Request) {
	if WasteLibrary.AllowCors {

		w.Header().Set("Access-Control-Allow-Origin", "*")
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
	WasteLibrary.LogStr("Header : " + currentHttpHeader.ToString())
	WasteLibrary.LogStr("Data : " + req.FormValue(WasteLibrary.HTTP_DATA))
	var execSQL string = ""
	if currentHttpHeader.BaseDataType == WasteLibrary.BASETYPE_CUSTOMER {

		var currentData WasteLibrary.CustomerType = WasteLibrary.StringToCustomerType(req.FormValue(WasteLibrary.HTTP_DATA))
		WasteLibrary.LogStr("Data : " + currentData.ToString())

		if currentData.CustomerId != 0 {
			execSQL = currentData.SelectSQL()
			WasteLibrary.LogStr(execSQL)
		}

		errDb := staticDb.QueryRow(execSQL).Scan(&currentData.CustomerName,
			&currentData.CustomerLink,
			&currentData.RfIdApp,
			&currentData.UltApp,
			&currentData.RecyApp,
			&currentData.Active,
			&currentData.CreateTime)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		resultVal.Retval = currentData.ToString()

	} else if currentHttpHeader.BaseDataType == WasteLibrary.BASETYPE_DEVICE {

		var currentData WasteLibrary.DeviceType = WasteLibrary.StringToDeviceType(req.FormValue(WasteLibrary.HTTP_DATA))
		WasteLibrary.LogStr("Data : " + currentData.ToString())

		if currentData.DeviceId != 0 {

			execSQL = currentData.SelectSQL()
			WasteLibrary.LogStr(execSQL)
		}

		errDb := staticDb.QueryRow(execSQL).Scan(
			&currentData.CustomerId,
			&currentData.DeviceName,
			&currentData.ContainerNo,
			&currentData.ContainerType,
			&currentData.DeviceType,
			&currentData.SerialNumber,
			&currentData.DeviceStatus,
			&currentData.StatusTime,
			&currentData.AliveStatus,
			&currentData.AliveLastOkTime,
			&currentData.Latitude,
			&currentData.Longitude,
			&currentData.GpsTime,
			&currentData.AlarmStatus,
			&currentData.AlarmTime,
			&currentData.AlarmType,
			&currentData.Alarm,
			&currentData.Therm,
			&currentData.ThermTime,
			&currentData.ThermStatus,
			&currentData.Active,
			&currentData.CreateTime,
			&currentData.ReaderAppStatus,
			&currentData.ReaderAppLastOkTime,
			&currentData.ReaderConnStatus,
			&currentData.ReaderConnLastOkTime,
			&currentData.ReaderStatus,
			&currentData.ReaderLastOkTime,
			&currentData.CamAppStatus,
			&currentData.CamAppLastOkTime,
			&currentData.CamConnStatus,
			&currentData.CamConnLastOkTime,
			&currentData.CamStatus,
			&currentData.CamLastOkTime,
			&currentData.GpsAppStatus,
			&currentData.GpsAppLastOkTime,
			&currentData.GpsConnStatus,
			&currentData.GpsConnLastOkTime,
			&currentData.GpsStatus,
			&currentData.GpsLastOkTime,
			&currentData.ThermAppStatus,
			&currentData.ThermAppLastOkTime,
			&currentData.TransferAppStatus,
			&currentData.TransferAppLastOkTime,
			&currentData.ContactStatus,
			&currentData.ContactLastOkTime,
			&currentData.Speed,
			&currentData.Battery,
			&currentData.BatteryStatus,
			&currentData.BatteryTime,
			&currentData.UltTime,
			&currentData.UltRange,
			&currentData.UltStatus,
			&currentData.TotalGlassCount,
			&currentData.TotalPlasticCount,
			&currentData.TotalMetalCount,
			&currentData.DailyGlassCount,
			&currentData.DailyPlasticCount,
			&currentData.DailyMetalCount,
			&currentData.RecyTime,
			&currentData.MotorAppStatus,
			&currentData.MotorAppLastOkTime,
			&currentData.MotorConnStatus,
			&currentData.MotorConnLastOkTime,
			&currentData.MotorStatus)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		resultVal.Retval = currentData.ToString()

	} else if currentHttpHeader.BaseDataType == WasteLibrary.BASETYPE_TAG {

		var currentData WasteLibrary.TagType = WasteLibrary.StringToTagType(req.FormValue(WasteLibrary.HTTP_DATA))
		WasteLibrary.LogStr("Data : " + currentData.ToString())

		if currentData.TagID != 0 {

			execSQL = currentData.SelectSQL()
			WasteLibrary.LogStr(execSQL)
		}

		errDb := staticDb.QueryRow(execSQL).Scan(&currentData.CustomerId,
			&currentData.CustomerId,
			&currentData.DeviceId,
			&currentData.UID,
			&currentData.Epc,
			&currentData.ContainerNo,
			&currentData.ContainerType,
			&currentData.Latitude,
			&currentData.Longitude,
			&currentData.Statu,
			&currentData.ImageStatu,
			&currentData.Active,
			&currentData.ReadTime,
			&currentData.CheckTime,
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
