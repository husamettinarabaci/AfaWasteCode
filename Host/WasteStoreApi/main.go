package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/AfatekDevelopers/result_lib_go/devafatekresult"
	"github.com/devafatek/WasteLibrary"
	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
)

var port int = 5432
var user string = os.Getenv("POSTGRES_USER")
var password string = os.Getenv("POSTGRES_PASSWORD")
var dbname string = os.Getenv("POSTGRES_DB")
var redisDb *redis.Client

var ctx = context.Background()
var sumDb *sql.DB
var bulkDb *sql.DB
var staticDb *sql.DB
var readerDb *sql.DB
var err error

func initStart() {

	WasteLibrary.LogStr("Successfully connected!")
	redisDb = redis.NewClient(&redis.Options{
		Addr:     "waste-redis-cluster-ip:6379",
		Password: "",
		DB:       0,
	})

	pong, err := redisDb.Ping(ctx).Result()
	WasteLibrary.LogErr(err)
	WasteLibrary.LogStr(pong)

}

func main() {

	initStart()

	var sumDbHost string = "waste-sumdb-cluster-ip"
	sumdDbInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		sumDbHost, port, user, password, dbname)

	sumDb, err = sql.Open("postgres", sumdDbInfo)
	WasteLibrary.LogErr(err)
	defer sumDb.Close()

	err = sumDb.Ping()
	WasteLibrary.LogErr(err)

	var bulkDbHost string = "waste-bulkdb-cluster-ip"
	bulkDbInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		bulkDbHost, port, user, password, dbname)

	bulkDb, err = sql.Open("postgres", bulkDbInfo)
	WasteLibrary.LogErr(err)
	defer bulkDb.Close()

	err = bulkDb.Ping()
	WasteLibrary.LogErr(err)

	var staticDbHost string = "waste-staticdb-cluster-ip"
	staticDbInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		staticDbHost, port, user, password, dbname)

	staticDb, err = sql.Open("postgres", staticDbInfo)
	WasteLibrary.LogErr(err)
	defer staticDb.Close()

	err = staticDb.Ping()
	WasteLibrary.LogErr(err)

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
	http.HandleFunc("/getkey", getkey)
	http.HandleFunc("/setkey", setkey)
	http.HandleFunc("/saveBulkDbMain", saveBulkDbMain)
	http.HandleFunc("/saveStaticDbMain", saveStaticDbMain)
	http.HandleFunc("/saveReaderDbMain", saveReaderDbMain)
	http.HandleFunc("/getStaticDbMain", getStaticDbMain)
	http.ListenAndServe(":80", nil)
	WasteLibrary.LogStr("Finished")
}

func saveBulkDbMain(w http.ResponseWriter, req *http.Request) {
	var resultVal devafatekresult.ResultType
	resultVal.Result = "FAIL"
	if err := req.ParseForm(); err != nil {
		WasteLibrary.LogErr(err)
		return
	}
	var currentHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.StringToHttpClientHeaderType(req.FormValue("HEADER"))
	dataVal := req.FormValue("DATA")
	WasteLibrary.LogStr("Header : " + currentHttpHeader.ToString())
	WasteLibrary.LogStr("Data : " + dataVal)
	var insertSQL string = fmt.Sprintf(`INSERT INTO public.listenerdata(
		app_type,device_id,optype,data,customer_id,data_time)
	   VALUES ('%s',%f,'%s','%s',%f,'%s');`, currentHttpHeader.AppType, currentHttpHeader.DeviceId, currentHttpHeader.OpType, dataVal, currentHttpHeader.CustomerId, currentHttpHeader.Time)
	WasteLibrary.LogStr(insertSQL)
	_, errDb := bulkDb.Exec(insertSQL)
	if errDb != nil {
		WasteLibrary.LogErr(err)
		resultVal.Result = "FAIL"
	} else {
		resultVal.Result = "OK"
	}
	w.Write(resultVal.ToByte())
}

func saveStaticDbMain(w http.ResponseWriter, req *http.Request) {
	var resultVal devafatekresult.ResultType
	resultVal.Result = "FAIL"
	if err := req.ParseForm(); err != nil {
		WasteLibrary.LogErr(err)
		return
	}

	var currentHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.StringToHttpClientHeaderType(req.FormValue("HEADER"))
	WasteLibrary.LogStr("Header : " + currentHttpHeader.ToString())
	WasteLibrary.LogStr("Data : " + req.FormValue("DATA"))
	if currentHttpHeader.AppType == "RFID" {
		var execSQL string = ""
		if currentHttpHeader.OpType == "RF" {
			var currentData WasteLibrary.TagType = WasteLibrary.StringToTagType(req.FormValue("DATA"))
			WasteLibrary.LogStr("Data : " + currentData.ToString())

			var selectSQL string = fmt.Sprintf(`SELECT tag_id
			FROM public.tags WHERE epc='%s' AND customer_id=%f;`, currentData.Epc, currentHttpHeader.CustomerId)
			rows, errSel := staticDb.Query(selectSQL)
			if errSel != nil {
				WasteLibrary.LogErr(errSel)
				resultVal.Result = "FAIL"
			} else {
				resultVal.Result = "OK"
			}
			var tagID int = 0
			for rows.Next() {
				rows.Scan(&tagID)
			}
			if tagID != 0 {
				if currentData.Latitude == 0 || currentData.Longitude == 0 {
					execSQL = fmt.Sprintf(`UPDATE public.tags
					SET uid='%s',read_time='%s',statu='1',device_id=%f
				   WHERE epc='%s' AND customer_id=%f 
				   RETURNING tag_id;`, currentData.UID, currentData.ReadTime,
						currentHttpHeader.DeviceId, currentData.Epc, currentHttpHeader.CustomerId)
					WasteLibrary.LogStr(execSQL)
				} else {
					execSQL = fmt.Sprintf(`UPDATE public.tags
					SET uid='%s',read_time='%s',latitude=%f,longitude=%f,device_id=%f
				   WHERE epc='%s' AND customer_id=%f 
				   RETURNING tag_id;`, currentData.UID, currentData.ReadTime,
						currentData.Latitude, currentData.Longitude,
						currentHttpHeader.DeviceId, currentData.Epc, currentHttpHeader.CustomerId)
					WasteLibrary.LogStr(execSQL)
				}
			} else {

				execSQL = fmt.Sprintf(`INSERT INTO public.tags(
			   epc,device_id,customer_id,uid,latitude,longitude)
			  VALUES ('%s',%f,%f,'%s',%f,%f)  
			  RETURNING tag_id;`, currentData.Epc, currentHttpHeader.DeviceId,
					currentHttpHeader.CustomerId, currentData.UID,
					currentData.Latitude, currentData.Longitude)
				WasteLibrary.LogStr(execSQL)
			}
			tagID = 0
			errDb := staticDb.QueryRow(execSQL).Scan(&tagID)
			if errDb != nil {
				WasteLibrary.LogErr(errDb)
				resultVal.Result = "FAIL"
			} else {
				resultVal.Result = "OK"
			}

			currentData.TagID = float64(tagID)
			resultVal.Retval = currentData.ToIdString()

		} else if currentHttpHeader.OpType == "GPS" {

			var currentData WasteLibrary.DeviceType = WasteLibrary.StringToDeviceType(req.FormValue("DATA"))
			WasteLibrary.LogStr("Data : " + currentData.ToString())

			execSQL = fmt.Sprintf(`UPDATE public.devices
			   SET gps_time='%s',latitude=%f,longitude=%f,speed=%f
			   WHERE device_id=%f AND customer_id=%f 
			   RETURNING device_id;`, currentData.GpsTime,
				currentData.Latitude, currentData.Longitude, currentData.Speed,
				currentHttpHeader.DeviceId, currentHttpHeader.CustomerId)
			WasteLibrary.LogStr(execSQL)
			var deviceID = 0
			errDb := staticDb.QueryRow(execSQL).Scan(&deviceID)
			if errDb != nil {
				WasteLibrary.LogErr(errDb)
				resultVal.Result = "FAIL"
			} else {
				resultVal.Result = "OK"
			}

			currentData.DeviceId = float64(deviceID)
			resultVal.Retval = currentData.ToIdString()

		} else if currentHttpHeader.OpType == "ARVENTO" {

			var currentData WasteLibrary.DeviceType = WasteLibrary.StringToDeviceType(req.FormValue("DATA"))
			WasteLibrary.LogStr("Data : " + currentData.ToString())

			execSQL = fmt.Sprintf(`UPDATE public.devices
			   SET latitude=%f,longitude=%f,speed=%f
			   WHERE device_id=%f AND customer_id=%f 
			   RETURNING device_id;`,
				currentData.Latitude, currentData.Longitude, currentData.Speed,
				currentData.DeviceId, currentData.CustomerId)
			WasteLibrary.LogStr(execSQL)
			var deviceID = 0
			errDb := staticDb.QueryRow(execSQL).Scan(&deviceID)
			if errDb != nil {
				WasteLibrary.LogErr(errDb)
				resultVal.Result = "FAIL"
			} else {
				resultVal.Result = "OK"
			}

			currentData.DeviceId = float64(deviceID)
			resultVal.Retval = currentData.ToIdString()

		} else if currentHttpHeader.OpType == "STATUS" {

			var currentData WasteLibrary.DeviceType = WasteLibrary.StringToDeviceType(req.FormValue("DATA"))
			WasteLibrary.LogStr("Data : " + currentData.ToString())
			var execSqlExt = ""
			if currentData.ReaderAppStatus == "1" {
				execSqlExt += ",reader_app_last_ok_time='" + currentData.ReaderAppLastOkTime + "'"
			}
			if currentData.ReaderConnStatus == "1" {
				execSqlExt += ",reader_conn_last_ok_time='" + currentData.ReaderConnLastOkTime + "'"
			}
			if currentData.ReaderStatus == "1" {
				execSqlExt += ",reader_last_ok_time='" + currentData.ReaderLastOkTime + "'"
			}

			if currentData.CamAppStatus == "1" {
				execSqlExt += ",cam_app_last_ok_time='" + currentData.CamAppLastOkTime + "'"
			}
			if currentData.CamConnStatus == "1" {
				execSqlExt += ",cam_conn_last_ok_time='" + currentData.CamConnLastOkTime + "'"
			}
			if currentData.CamStatus == "1" {
				execSqlExt += ",cam_last_ok_time='" + currentData.CamLastOkTime + "'"
			}

			if currentData.GpsAppStatus == "1" {
				execSqlExt += ",gps_app_last_ok_time='" + currentData.GpsAppLastOkTime + "'"
			}
			if currentData.GpsConnStatus == "1" {
				execSqlExt += ",gps_conn_last_ok_time='" + currentData.GpsConnLastOkTime + "'"
			}
			if currentData.GpsStatus == "1" {
				execSqlExt += ",gps_last_ok_time='" + currentData.GpsLastOkTime + "'"
			}

			if currentData.ThermAppStatus == "1" {
				execSqlExt += ",therm_app_last_ok_time='" + currentData.ThermAppLastOkTime + "'"
			}
			if currentData.TransferAppStatus == "1" {
				execSqlExt += ",transfer_app_last_ok_time='" + currentData.TransferAppLastOkTime + "'"
			}
			if currentData.AliveStatus == "1" {
				execSqlExt += ",alive_last_ok_time='" + currentData.AliveLastOkTime + "'"
			}
			if currentData.ContactStatus == "1" {
				execSqlExt += ",contact_last_ok_time='" + currentData.ContactLastOkTime + "'"
			}

			execSQL = fmt.Sprintf(`UPDATE public.devices
				SET status_time='%s',
				reader_app_status='%s',reader_conn_status='%s',reader_status='%s',cam_app_status='%s',cam_conn_status='%s',
				cam_status='%s',gps_app_status='%s',gps_conn_status='%s',gps_status='%s',therm_app_status='%s',
				transfer_app_status='%s',alive_status='%s',contact_status='%s'`+execSqlExt+`
			   WHERE device_id=%f AND customer_id=%f 
			   RETURNING device_id;`, currentData.StatusTime,
				currentData.ReaderAppStatus, currentData.ReaderConnStatus, currentData.ReaderStatus,
				currentData.CamAppStatus, currentData.CamConnStatus, currentData.CamStatus,
				currentData.GpsAppStatus, currentData.GpsConnStatus, currentData.GpsStatus,
				currentData.ThermAppStatus, currentData.TransferAppStatus, currentData.AliveStatus,
				currentData.ContactStatus, currentHttpHeader.DeviceId, currentHttpHeader.CustomerId)

			WasteLibrary.LogStr(execSQL)
			var deviceID = 0
			errDb := staticDb.QueryRow(execSQL).Scan(&deviceID)
			if errDb != nil {
				WasteLibrary.LogErr(errDb)
				resultVal.Result = "FAIL"
			} else {
				resultVal.Result = "OK"
			}

			currentData.DeviceId = float64(deviceID)
			resultVal.Retval = currentData.ToIdString()

		} else if currentHttpHeader.OpType == "THERM" {

			var currentData WasteLibrary.DeviceType = WasteLibrary.StringToDeviceType(req.FormValue("DATA"))
			WasteLibrary.LogStr("Data : " + currentData.ToString())

			execSQL = fmt.Sprintf(`UPDATE public.devices
			   SET therm='%s',therm_time='%s'
			   WHERE device_id=%f AND customer_id=%f 
			   RETURNING device_id;`,
				currentData.Therm, currentData.ThermTime, currentHttpHeader.DeviceId, currentHttpHeader.CustomerId)
			WasteLibrary.LogStr(execSQL)
			var deviceID = 0
			errDb := staticDb.QueryRow(execSQL).Scan(&deviceID)
			if errDb != nil {
				WasteLibrary.LogErr(errDb)
				resultVal.Result = "FAIL"
			} else {
				resultVal.Result = "OK"
			}

			currentData.DeviceId = float64(deviceID)
			resultVal.Retval = currentData.ToIdString()

		} else {
			resultVal.Result = "FAIL"
		}

	} else if currentHttpHeader.AppType == "ULT" {
		resultVal.Result = "OK"
	} else if currentHttpHeader.AppType == "RECY" {
		resultVal.Result = "OK"
	} else if currentHttpHeader.AppType == "ADMIN" {
		var execSQL string = ""
		if currentHttpHeader.OpType == "CUSTOMER" {

			var currentData WasteLibrary.CustomerType = WasteLibrary.StringToCustomerType(req.FormValue("DATA"))
			WasteLibrary.LogStr("Data : " + currentData.ToString())

			if currentData.CustomerId != 0 {
				execSQL = fmt.Sprintf(`UPDATE public.customers
					SET customer_name='%s',customer_link='%s',rfid_app='%s',ult_app='%s',recy_app='%s'
				   WHERE customer_id=%f  RETURNING customer_id;`,
					currentData.CustomerName, currentData.CustomerLink, currentData.RfIdApp,
					currentData.UltApp, currentData.RecyApp, currentData.CustomerId)
				WasteLibrary.LogStr(execSQL)
			} else {

				execSQL = fmt.Sprintf(`INSERT INTO public.customers(
					customer_name,customer_link,rfid_app,ult_app,recy_app)
			  VALUES ('%s','%s','%s','%s','%s')  RETURNING customer_id;`,
					currentData.CustomerName, currentData.CustomerLink, currentData.RfIdApp,
					currentData.UltApp, currentData.RecyApp)
				WasteLibrary.LogStr(execSQL)
			}
			var customerId int = 0
			errDb := staticDb.QueryRow(execSQL).Scan(&customerId)
			if errDb != nil {
				WasteLibrary.LogErr(errDb)
				resultVal.Result = "FAIL"
			} else {
				resultVal.Result = "OK"
			}

			currentData.CustomerId = float64(customerId)
			resultVal.Retval = currentData.ToIdString()

		} else if currentHttpHeader.OpType == "DEVICE" {

			var currentData WasteLibrary.DeviceType = WasteLibrary.StringToDeviceType(req.FormValue("DATA"))
			WasteLibrary.LogStr("Data : " + currentData.ToString())
			if currentData.DeviceId != 0 {
				execSQL = fmt.Sprintf(`UPDATE public.devices 
				SET device_type='%s',serial_number='%s',device_name='%s',customer_id=%f 
	  			WHERE device_id=%f  
				RETURNING device_id;`,
					currentData.DeviceType, currentData.SerialNumber, currentData.DeviceName,
					currentData.CustomerId, currentData.DeviceId)
				WasteLibrary.LogStr(execSQL)
			} else {

				execSQL = fmt.Sprintf(`INSERT INTO public.devices 
				(device_type,serial_number,device_name,customer_id) 
  				VALUES ('%s','%s','%s',%f)   
  				RETURNING device_id;`,
					currentData.DeviceType, currentData.SerialNumber, currentData.DeviceName,
					currentData.CustomerId)
				WasteLibrary.LogStr(execSQL)
			}
			var deviceId int = 0
			errDb := staticDb.QueryRow(execSQL).Scan(&deviceId)
			if errDb != nil {
				WasteLibrary.LogErr(errDb)
				resultVal.Result = "FAIL"
			} else {
				resultVal.Result = "OK"
			}

			currentData.DeviceId = float64(deviceId)
			resultVal.Retval = currentData.ToIdString()

		} else if currentHttpHeader.OpType == "USER" {

			var currentData WasteLibrary.UserType = WasteLibrary.StringToUserType(req.FormValue("DATA"))
			WasteLibrary.LogStr("Data : " + currentData.ToString())
			if currentData.UserId != 0 {
				execSQL = fmt.Sprintf(`UPDATE public.users 
				SET user_type='%s',email='%s',user_name='%s',customer_id=%f 
	  			WHERE user_id=%f  
				RETURNING user_id;`,
					currentData.UserType, currentData.Email, currentData.UserName,
					currentData.CustomerId, currentData.UserId)
				WasteLibrary.LogStr(execSQL)
			} else {

				execSQL = fmt.Sprintf(`INSERT INTO public.users 
				(user_type,email,user_name,customer_id) 
  				VALUES ('%s','%s','%s',%f)   
  				RETURNING user_id;`,
					currentData.UserType, currentData.Email, currentData.UserName,
					currentData.CustomerId)
				WasteLibrary.LogStr(execSQL)
			}
			var userId int = 0
			errDb := staticDb.QueryRow(execSQL).Scan(&userId)
			if errDb != nil {
				WasteLibrary.LogErr(errDb)
				resultVal.Result = "FAIL"
			} else {
				resultVal.Result = "OK"
			}

			currentData.UserId = float64(userId)
			resultVal.Retval = currentData.ToIdString()

		} else {
			resultVal.Result = "FAIL"
		}
	} else {
		resultVal.Result = "OK"
	}
	w.Write(resultVal.ToByte())
}

func saveReaderDbMain(w http.ResponseWriter, req *http.Request) {
	var resultVal devafatekresult.ResultType
	resultVal.Result = "FAIL"
	if err := req.ParseForm(); err != nil {
		WasteLibrary.LogErr(err)
		return
	}

	var currentHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.StringToHttpClientHeaderType(req.FormValue("HEADER"))
	WasteLibrary.LogStr("Header : " + currentHttpHeader.ToString())
	WasteLibrary.LogStr("Data : " + req.FormValue("DATA"))
	if currentHttpHeader.AppType == "RFID" {
		var execSQL string = ""
		if currentHttpHeader.OpType == "TAG" {

			var currentData WasteLibrary.TagType = WasteLibrary.StringToTagType(req.FormValue("DATA"))
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
				resultVal.Result = "FAIL"
			} else {
				resultVal.Result = "OK"
			}

			currentData.DeviceId = float64(deviceId)
			resultVal.Retval = currentData.ToIdString()

		} else if currentHttpHeader.OpType == "DEVICE" {

			var currentData WasteLibrary.DeviceType = WasteLibrary.StringToDeviceType(req.FormValue("DATA"))
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
				resultVal.Result = "FAIL"
			} else {
				resultVal.Result = "OK"
			}

			currentData.DeviceId = float64(deviceId)
			resultVal.Retval = currentData.ToIdString()

		} else {
			resultVal.Result = "FAIL"
		}
	} else {
		resultVal.Result = "OK"
	}
	w.Write(resultVal.ToByte())
}

func getStaticDbMain(w http.ResponseWriter, req *http.Request) {
	var resultVal devafatekresult.ResultType
	resultVal.Result = "FAIL"
	if err := req.ParseForm(); err != nil {
		WasteLibrary.LogErr(err)
		return
	}
	var currentHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.StringToHttpClientHeaderType(req.FormValue("HEADER"))
	WasteLibrary.LogStr("Header : " + currentHttpHeader.ToString())
	WasteLibrary.LogStr("Data : " + req.FormValue("DATA"))
	var execSQL string = ""
	if currentHttpHeader.BaseDataType == "CUSTOMER" {

		var currentData WasteLibrary.CustomerType = WasteLibrary.StringToCustomerType(req.FormValue("DATA"))
		WasteLibrary.LogStr("Data : " + currentData.ToString())

		if currentData.CustomerId != 0 {
			execSQL = fmt.Sprintf(`SELECT 
			customer_name,
			customer_link,
			rfid_app,
			ult_app,
			recy_app,
			active,
			create_time
			FROM public.customers
				   WHERE customer_id=%f ;`, currentData.CustomerId)
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
			resultVal.Result = "FAIL"
		} else {
			resultVal.Result = "OK"
		}

		resultVal.Retval = currentData.ToString()

	} else if currentHttpHeader.BaseDataType == "DEVICE" {

		var currentData WasteLibrary.DeviceType = WasteLibrary.StringToDeviceType(req.FormValue("DATA"))
		WasteLibrary.LogStr("Data : " + currentData.ToString())

		if currentData.DeviceId != 0 {
			execSQL = fmt.Sprintf(`SELECT 
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
					create_time
			 FROM public.devices
				   WHERE device_id=%f ;`, currentData.DeviceId)
			WasteLibrary.LogStr(execSQL)
		}

		errDb := staticDb.QueryRow(execSQL).Scan(
			&currentData.DeviceType,
			&currentData.SerialNumber,
			&currentData.DeviceName,
			&currentData.CustomerId,
			&currentData.ReaderAppStatus,
			&currentData.ReaderConnStatus,
			&currentData.ReaderStatus,
			&currentData.CamAppStatus,
			&currentData.CamConnStatus,
			&currentData.CamStatus,
			&currentData.GpsAppStatus,
			&currentData.GpsConnStatus,
			&currentData.GpsStatus,
			&currentData.ThermAppStatus,
			&currentData.TransferAppStatus,
			&currentData.AliveStatus,
			&currentData.ContactStatus,
			&currentData.Therm,
			&currentData.Latitude,
			&currentData.Longitude,
			&currentData.Speed,
			&currentData.UltRange,
			&currentData.UltStatus,
			&currentData.DeviceStatus,
			&currentData.TotalGlassCount,
			&currentData.TotalMetalCount,
			&currentData.TotalPlasticCount,
			&currentData.UltTime,
			&currentData.AlarmTime,
			&currentData.AlarmStatus,
			&currentData.ThermStatus,
			&currentData.BatteryStatus,
			&currentData.Battery,
			&currentData.AlarmType,
			&currentData.Alarm,
			&currentData.RecyTime,
			&currentData.ContactStatus,
			&currentData.Active,
			&currentData.ThermTime,
			&currentData.GpsTime,
			&currentData.StatusTime,
			&currentData.ReaderAppLastOkTime,
			&currentData.ReaderConnLastOkTime,
			&currentData.ReaderLastOkTime,
			&currentData.GpsAppLastOkTime,
			&currentData.GpsConnLastOkTime,
			&currentData.GpsLastOkTime,
			&currentData.CamAppLastOkTime,
			&currentData.BatteryTime,
			&currentData.CamConnLastOkTime,
			&currentData.CamLastOkTime,
			&currentData.ThermAppLastOkTime,
			&currentData.TransferAppLastOkTime,
			&currentData.AliveLastOkTime,
			&currentData.ContactLastOkTime,
			&currentData.CreateTime)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = "FAIL"
		} else {
			resultVal.Result = "OK"
		}

		resultVal.Retval = currentData.ToString()

	} else if currentHttpHeader.BaseDataType == "USER" {

		var currentData WasteLibrary.UserType = WasteLibrary.StringToUserType(req.FormValue("DATA"))
		WasteLibrary.LogStr("Data : " + currentData.ToString())

		if currentData.UserId != 0 {
			execSQL = fmt.Sprintf(`SELECT 
			user_type,
			email,
			user_name,
			customer_id,
			create_time
			FROM public.users
				   WHERE user_id=%f ;`, currentData.UserId)
			WasteLibrary.LogStr(execSQL)
		}

		errDb := staticDb.QueryRow(execSQL).Scan(&currentData.UserType,
			&currentData.Email,
			&currentData.UserName,
			&currentData.CustomerId,
			&currentData.Token,
			&currentData.TokenEndTime,
			&currentData.CreateTime)
		if errDb != nil {
			WasteLibrary.LogErr(errDb)
			resultVal.Result = "FAIL"
		} else {
			resultVal.Result = "OK"
		}

		resultVal.Retval = currentData.ToString()

	} else if currentHttpHeader.BaseDataType == "TAG" {

		var currentData WasteLibrary.TagType = WasteLibrary.StringToTagType(req.FormValue("DATA"))
		WasteLibrary.LogStr("Data : " + currentData.ToString())

		if currentData.TagID != 0 {
			execSQL = fmt.Sprintf(`SELECT 
			customer_id,
			device_id,
			epc,
			uid,
			container_no,
			latitude, 
			longitude, 
			statu,
			image_statu,
			active,
			read_time, 
			check_time, 
			create_time
			FROM public.tags
				   WHERE tag_id=%f ;`, currentData.TagID)
			WasteLibrary.LogStr(execSQL)
		}

		errDb := staticDb.QueryRow(execSQL).Scan(&currentData.CustomerId,
			&currentData.DeviceId,
			&currentData.Epc,
			&currentData.UID,
			&currentData.ContainerNo,
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
			resultVal.Result = "FAIL"
		} else {
			resultVal.Result = "OK"
		}

		resultVal.Retval = currentData.ToString()

	} else {
		resultVal.Result = "FAIL"
	}

	w.Write(resultVal.ToByte())
}

func getkey(w http.ResponseWriter, req *http.Request) {
	var resultVal devafatekresult.ResultType
	resultVal.Result = "FAIL"
	if err := req.ParseForm(); err != nil {
		WasteLibrary.LogErr(err)
		return
	}
	hKey := req.FormValue("HASHKEY")
	sKey := req.FormValue("SUBKEY")
	WasteLibrary.LogStr("GetKey : " + hKey + " - " + sKey)
	resultVal = getKeyRedis(hKey, sKey)
	WasteLibrary.LogStr("RetValByRedis : " + resultVal.ToString())
	if resultVal.Result == "FAIL" {
		resultVal = getKeyDb(hKey, sKey)
		if resultVal.Result != "FAIL" {
			setKeyRedis(hKey, sKey, resultVal.Retval.(string))
		}
	}
	w.Write(resultVal.ToByte())
}

func setkey(w http.ResponseWriter, req *http.Request) {
	var resultVal devafatekresult.ResultType
	resultVal.Result = "FAIL"
	if err := req.ParseForm(); err != nil {
		WasteLibrary.LogErr(err)
		return
	}
	hKey := req.FormValue("HASHKEY")
	sKey := req.FormValue("SUBKEY")
	kVal := req.FormValue("KEYVALUE")
	WasteLibrary.LogStr("GetKeyDb : " + resultVal.ToString())
	resultVal = getKeyDb(hKey, sKey)
	WasteLibrary.LogStr("GetKeyDb : " + resultVal.ToString())
	if resultVal.Result == "FAIL" {
		resultVal = insertKeyDb(hKey, sKey, kVal)
	} else {
		resultVal = updateKeyDb(hKey, sKey, kVal)
	}
	setKeyRedis(hKey, sKey, kVal)

	w.Write(resultVal.ToByte())
}

func getKeyRedis(hKey string, sKey string) devafatekresult.ResultType {
	var resultVal devafatekresult.ResultType
	resultVal.Result = "FAIL"
	var val string = ""
	var err error
	if hKey != "" {
		val, err = redisDb.HGet(ctx, hKey, sKey).Result()
	} else {
		val, err = redisDb.Get(ctx, sKey).Result()
	}
	switch {
	case err == redis.Nil:
		WasteLibrary.LogStr("Not Found")
		resultVal.Result = "FAIL"
	case err != nil:
		WasteLibrary.LogErr(err)
	case val == "":
		WasteLibrary.LogStr("Not Found")
		resultVal.Result = "FAIL"
	case val != "":
		resultVal.Result = "OK"
		resultVal.Retval = val
		WasteLibrary.LogStr(resultVal.ToString())
	}

	return resultVal
}

func setKeyRedis(hKey string, sKey string, kVal string) {
	var err error
	if hKey != "" {
		_, err = redisDb.HSet(ctx, hKey, sKey, kVal).Result()
	} else {
		_, err = redisDb.HSet(ctx, sKey, kVal).Result()

	}
	switch {
	case err == redis.Nil:
		WasteLibrary.LogStr("Not Found")
	case err != nil:
		WasteLibrary.LogErr(err)
	}
}

func getKeyDb(hKey string, sKey string) devafatekresult.ResultType {
	var resultVal devafatekresult.ResultType
	resultVal.Result = "FAIL"
	WasteLibrary.LogStr("Serach Db : " + hKey + " - " + sKey)

	var selectSQL string = fmt.Sprintf(`SELECT keyvalue 
	FROM public.redisdata WHERE hashkey='%s' AND subkey='%s';`, hKey, sKey)
	rows, errSel := sumDb.Query(selectSQL)
	WasteLibrary.LogErr(errSel)
	var kVal string = "NOT"
	for rows.Next() {
		rows.Scan(&kVal)
	}

	if kVal == "NOT" {
		resultVal.Result = "FAIL"
	} else {
		resultVal.Retval = kVal
		resultVal.Result = "OK"
	}
	WasteLibrary.LogStr("KeyValue : " + kVal)
	return resultVal
}

func insertKeyDb(hKey string, sKey string, kVal string) devafatekresult.ResultType {
	var resultVal devafatekresult.ResultType
	resultVal.Result = "FAIL"
	WasteLibrary.LogStr("Insert Db : " + hKey + " - " + sKey + " - " + kVal)
	var insertSQL string = fmt.Sprintf(`INSERT INTO public.redisdata(
		hashkey,subkey,keyvalue)
	   VALUES ('%s','%s','%s');`, hKey, sKey, kVal)
	_, errDb := sumDb.Exec(insertSQL)
	WasteLibrary.LogErr(errDb)
	resultVal.Result = "OK"
	return resultVal
}

func updateKeyDb(hKey string, sKey string, kVal string) devafatekresult.ResultType {
	var resultVal devafatekresult.ResultType
	resultVal.Result = "FAIL"
	WasteLibrary.LogStr("Update Db : " + hKey + " - " + sKey + " - " + kVal)
	var updateSQL string = fmt.Sprintf(`UPDATE public.redisdata SET keyvalue='%s' WHERE hashkey='%s' AND subkey='%s';`, kVal, hKey, sKey)
	_, errDb := sumDb.Exec(updateSQL)
	WasteLibrary.LogErr(errDb)
	resultVal.Result = "OK"
	return resultVal
}
