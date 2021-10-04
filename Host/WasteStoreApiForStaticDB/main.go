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
	WasteLibrary.LogStr("Finished")
}

func saveStaticDbMain(w http.ResponseWriter, req *http.Request) {
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
		if currentHttpHeader.OpType == WasteLibrary.OPTYPE_RF {
			var currentData WasteLibrary.TagType = WasteLibrary.StringToTagType(req.FormValue(WasteLibrary.HTTP_DATA))
			WasteLibrary.LogStr("Data : " + currentData.ToString())

			var selectSQL string = fmt.Sprintf(`SELECT tag_id
			FROM public.tags WHERE epc='%s' AND customer_id=%f;`, currentData.Epc, currentHttpHeader.CustomerId)
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
				if currentData.Latitude == 0 || currentData.Longitude == 0 {
					execSQL = fmt.Sprintf(`UPDATE public.tags
					SET uid='%s',read_time='%s',statu='%s',device_id=%f
				   WHERE epc='%s' AND customer_id=%f 
				   RETURNING tag_id;`, currentData.UID, currentData.ReadTime, WasteLibrary.STATU_ACTIVE,
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
				resultVal.Result = WasteLibrary.RESULT_FAIL
			} else {
				resultVal.Result = WasteLibrary.RESULT_OK
			}

			currentData.TagID = float64(tagID)
			resultVal.Retval = currentData.ToIdString()

		} else if currentHttpHeader.OpType == WasteLibrary.OPTYPE_GPS {

			var currentData WasteLibrary.DeviceType = WasteLibrary.StringToDeviceType(req.FormValue(WasteLibrary.HTTP_DATA))
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
				resultVal.Result = WasteLibrary.RESULT_FAIL
			} else {
				resultVal.Result = WasteLibrary.RESULT_OK
			}

			currentData.DeviceId = float64(deviceID)
			resultVal.Retval = currentData.ToIdString()

		} else if currentHttpHeader.OpType == WasteLibrary.OPTYPE_ARVENTO {

			var currentData WasteLibrary.DeviceType = WasteLibrary.StringToDeviceType(req.FormValue(WasteLibrary.HTTP_DATA))
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
				resultVal.Result = WasteLibrary.RESULT_FAIL
			} else {
				resultVal.Result = WasteLibrary.RESULT_OK
			}

			currentData.DeviceId = float64(deviceID)
			resultVal.Retval = currentData.ToIdString()

		} else if currentHttpHeader.OpType == WasteLibrary.OPTYPE_STATUS {

			var currentData WasteLibrary.DeviceType = WasteLibrary.StringToDeviceType(req.FormValue(WasteLibrary.HTTP_DATA))
			WasteLibrary.LogStr("Data : " + currentData.ToString())
			var execSqlExt = ""
			if currentData.ReaderAppStatus == WasteLibrary.STATU_ACTIVE {
				execSqlExt += ",reader_app_last_ok_time='" + currentData.ReaderAppLastOkTime + "'"
			}
			if currentData.ReaderConnStatus == WasteLibrary.STATU_ACTIVE {
				execSqlExt += ",reader_conn_last_ok_time='" + currentData.ReaderConnLastOkTime + "'"
			}
			if currentData.ReaderStatus == WasteLibrary.STATU_ACTIVE {
				execSqlExt += ",reader_last_ok_time='" + currentData.ReaderLastOkTime + "'"
			}

			if currentData.CamAppStatus == WasteLibrary.STATU_ACTIVE {
				execSqlExt += ",cam_app_last_ok_time='" + currentData.CamAppLastOkTime + "'"
			}
			if currentData.CamConnStatus == WasteLibrary.STATU_ACTIVE {
				execSqlExt += ",cam_conn_last_ok_time='" + currentData.CamConnLastOkTime + "'"
			}
			if currentData.CamStatus == WasteLibrary.STATU_ACTIVE {
				execSqlExt += ",cam_last_ok_time='" + currentData.CamLastOkTime + "'"
			}

			if currentData.GpsAppStatus == WasteLibrary.STATU_ACTIVE {
				execSqlExt += ",gps_app_last_ok_time='" + currentData.GpsAppLastOkTime + "'"
			}
			if currentData.GpsConnStatus == WasteLibrary.STATU_ACTIVE {
				execSqlExt += ",gps_conn_last_ok_time='" + currentData.GpsConnLastOkTime + "'"
			}
			if currentData.GpsStatus == WasteLibrary.STATU_ACTIVE {
				execSqlExt += ",gps_last_ok_time='" + currentData.GpsLastOkTime + "'"
			}

			if currentData.ThermAppStatus == WasteLibrary.STATU_ACTIVE {
				execSqlExt += ",therm_app_last_ok_time='" + currentData.ThermAppLastOkTime + "'"
			}
			if currentData.TransferAppStatus == WasteLibrary.STATU_ACTIVE {
				execSqlExt += ",transfer_app_last_ok_time='" + currentData.TransferAppLastOkTime + "'"
			}
			if currentData.AliveStatus == WasteLibrary.STATU_ACTIVE {
				execSqlExt += ",alive_last_ok_time='" + currentData.AliveLastOkTime + "'"
			}
			if currentData.ContactStatus == WasteLibrary.STATU_ACTIVE {
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
				resultVal.Result = WasteLibrary.RESULT_FAIL
			} else {
				resultVal.Result = WasteLibrary.RESULT_OK
			}

			currentData.DeviceId = float64(deviceID)
			resultVal.Retval = currentData.ToIdString()

		} else if currentHttpHeader.OpType == WasteLibrary.OPTYPE_THERM {

			var currentData WasteLibrary.DeviceType = WasteLibrary.StringToDeviceType(req.FormValue(WasteLibrary.HTTP_DATA))
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
	} else if currentHttpHeader.AppType == WasteLibrary.APPTYPE_ADMIN {
		var execSQL string = ""
		if currentHttpHeader.OpType == WasteLibrary.OPTYPE_CUSTOMER {

			var currentData WasteLibrary.CustomerType = WasteLibrary.StringToCustomerType(req.FormValue(WasteLibrary.HTTP_DATA))
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
				resultVal.Result = WasteLibrary.RESULT_FAIL
			} else {
				resultVal.Result = WasteLibrary.RESULT_OK
			}

			currentData.DeviceId = float64(deviceId)
			resultVal.Retval = currentData.ToIdString()

		} else if currentHttpHeader.OpType == WasteLibrary.OPTYPE_USER {

			var currentData WasteLibrary.UserType = WasteLibrary.StringToUserType(req.FormValue(WasteLibrary.HTTP_DATA))
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

func getStaticDbMain(w http.ResponseWriter, req *http.Request) {
	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	if err := req.ParseForm(); err != nil {
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
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		resultVal.Retval = currentData.ToString()

	} else if currentHttpHeader.BaseDataType == WasteLibrary.BASETYPE_DEVICE {

		var currentData WasteLibrary.DeviceType = WasteLibrary.StringToDeviceType(req.FormValue(WasteLibrary.HTTP_DATA))
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
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		resultVal.Retval = currentData.ToString()

	} else if currentHttpHeader.BaseDataType == WasteLibrary.BASETYPE_USER {

		var currentData WasteLibrary.UserType = WasteLibrary.StringToUserType(req.FormValue(WasteLibrary.HTTP_DATA))
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
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			resultVal.Result = WasteLibrary.RESULT_OK
		}

		resultVal.Retval = currentData.ToString()

	} else if currentHttpHeader.BaseDataType == WasteLibrary.BASETYPE_TAG {

		var currentData WasteLibrary.TagType = WasteLibrary.StringToTagType(req.FormValue(WasteLibrary.HTTP_DATA))
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
