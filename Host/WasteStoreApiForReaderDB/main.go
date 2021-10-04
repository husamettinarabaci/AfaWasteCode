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

	WasteLibrary.LogStr("Start")
	http.HandleFunc("/health", WasteLibrary.HealthHandler)
	http.HandleFunc("/readiness", WasteLibrary.ReadinessHandler)
	http.HandleFunc("/status", WasteLibrary.StatusHandler)
	http.HandleFunc("/saveReaderDbMain", saveReaderDbMain)
	http.ListenAndServe(":80", nil)
}

func saveReaderDbMain(w http.ResponseWriter, req *http.Request) {
	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	if err := req.ParseForm(); err != nil {
		WasteLibrary.LogErr(err)
		return
	}

	var currentHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.StringToHttpClientHeaderType(req.FormValue(WasteLibrary.HTTP_HEADER))
	WasteLibrary.LogStr("Header : " + currentHttpHeader.ToString())
	WasteLibrary.LogStr("Data : " + req.FormValue(WasteLibrary.HTTP_DATA))
	if currentHttpHeader.AppType == WasteLibrary.APPTYPE_RFID {
		var execSQL string = ""
		if currentHttpHeader.OpType == WasteLibrary.OPTYPE_TAG {

			var currentData WasteLibrary.TagType = WasteLibrary.StringToTagType(req.FormValue(WasteLibrary.HTTP_DATA))
			WasteLibrary.LogStr("Data : " + currentData.ToString())

			execSQL = fmt.Sprintf(`INSERT INTO public.tagdata 
				(tag_id,customer_id,device_id,epc,
				uid,container_no,latitude, longitude, 
				statu,image_statu,active,
				read_time,check_time,create_time) 
  				VALUES (%f,%f,%f,'%s''%s''%s',%f,%f,'%s','%s','%s','%s','%s','%s')   
  				RETURNING data_id;`,
				currentData.TagID, currentData.CustomerId, currentData.DeviceId,
				currentData.Epc, currentData.UID, currentData.ContainerNo,
				currentData.Latitude, currentData.Longitude,
				currentData.Statu, currentData.ImageStatu, currentData.Active,
				currentData.ReadTime, currentData.CheckTime, currentData.CreateTime)
			WasteLibrary.LogStr(execSQL)

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

		} else if currentHttpHeader.OpType == WasteLibrary.OPTYPE_DEVICE {

			var currentData WasteLibrary.DeviceType = WasteLibrary.StringToDeviceType(req.FormValue(WasteLibrary.HTTP_DATA))
			WasteLibrary.LogStr("Data : " + currentData.ToString())

			execSQL = fmt.Sprintf(`INSERT INTO public.devicedata 
				(device_id,
					device_type,
					serial_number,
					device_name,
					customer_id,
					reader_app_status,
					reader_conn_status,
					reader_status,
					cam_app_status,
					cam_conn_status,
					cam_status,
					gps_app_status,
					gps_conn_status,
					gps_status,
					therm_app_status,
					transfer_app_status,
					alive_status,
					contact_status,
					therm,
					latitude, 
					longitude, 
					speed,
					ult_range,
					ult_status,
					device_status,
					total_glass_count,
					total_metal_count,
					total_plastic_count,
					ult_time, 
					alarm_time, 
					alarm_status,
					therm_status,
					battery_status,
					battery,
					alarm_type,
					alarm,
					recy_time, 
					contact_status,
					active, 
					therm_time, 
					gps_time, 
					status_time,
					reader_app_last_ok_time,
					reader_conn_last_ok_time,
					reader_last_ok_time,
					gps_app_last_ok_time,
					gps_conn_last_ok_time,
					gps_last_ok_time,
					cam_app_last_ok_time,
					battery_time,
					cam_conn_last_ok_time,
					cam_last_ok_time,
					therm_app_last_ok_time,
					transfer_app_last_ok_time,
					alive_last_ok_time,
					contact_last_ok_time,
					create_time) 
  				VALUES (%f,'%s',
					'%s','%s',%f,'%s','%s',
					'%s','%s','%s','%s','%s',
					'%s','%s','%s','%s','%s',
					'%s','%s',%f,%f,%f,
					%f,'%s','%s',%f,%f,
					%f,'%s','%s','%s','%s',
					'%s','%s','%s','%s','%s', 
					'%s','%s','%s','%s','%s',
					'%s','%s','%s','%s','%s',
					'%s','%s','%s','%s','%s',
					'%s','%s','%s','%s','%s') 
  				RETURNING device_id;`,
				currentData.DeviceId,
				currentData.DeviceType,
				currentData.SerialNumber,
				currentData.DeviceName,
				currentData.CustomerId,
				currentData.ReaderAppStatus,
				currentData.ReaderConnStatus,
				currentData.ReaderStatus,
				currentData.CamAppStatus,
				currentData.CamConnStatus,
				currentData.CamStatus,
				currentData.GpsAppStatus,
				currentData.GpsConnStatus,
				currentData.GpsStatus,
				currentData.ThermAppStatus,
				currentData.TransferAppStatus,
				currentData.AliveStatus,
				currentData.ContactStatus,
				currentData.Therm,
				currentData.Latitude,
				currentData.Longitude,
				currentData.Speed,
				currentData.UltRange,
				currentData.UltStatus,
				currentData.DeviceStatus,
				currentData.TotalGlassCount,
				currentData.TotalMetalCount,
				currentData.TotalPlasticCount,
				currentData.UltTime,
				currentData.AlarmTime,
				currentData.AlarmStatus,
				currentData.ThermStatus,
				currentData.BatteryStatus,
				currentData.Battery,
				currentData.AlarmType,
				currentData.Alarm,
				currentData.RecyTime,
				currentData.ContactStatus,
				currentData.Active,
				currentData.ThermTime,
				currentData.GpsTime,
				currentData.StatusTime,
				currentData.ReaderAppLastOkTime,
				currentData.ReaderConnLastOkTime,
				currentData.ReaderLastOkTime,
				currentData.GpsAppLastOkTime,
				currentData.GpsConnLastOkTime,
				currentData.GpsLastOkTime,
				currentData.CamAppLastOkTime,
				currentData.BatteryTime,
				currentData.CamConnLastOkTime,
				currentData.CamLastOkTime,
				currentData.ThermAppLastOkTime,
				currentData.TransferAppLastOkTime,
				currentData.AliveLastOkTime,
				currentData.ContactLastOkTime,
				currentData.CreateTime)

			WasteLibrary.LogStr(execSQL)

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

		} else {
			resultVal.Result = WasteLibrary.RESULT_FAIL
		}
	} else {
		resultVal.Result = WasteLibrary.RESULT_OK
	}
	w.Write(resultVal.ToByte())
}
