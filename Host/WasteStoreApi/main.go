package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strconv"

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

	WasteLibrary.LogStr("Start")
	http.HandleFunc("/health", WasteLibrary.HealthHandler)
	http.HandleFunc("/readiness", WasteLibrary.ReadinessHandler)
	http.HandleFunc("/status", WasteLibrary.StatusHandler)
	http.HandleFunc("/getkey", getkey)
	http.HandleFunc("/setkey", setkey)
	http.HandleFunc("/saveBulkDbMain", saveBulkDbMain)
	http.HandleFunc("/saveStaticDbMain", saveStaticDbMain)
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
	appTypeVal := req.FormValue("APPTYPE")
	didVal := req.FormValue("DID")
	dataTypeVal := req.FormValue("DATATYPE")
	timeVal := req.FormValue("TIME")
	dataVal := req.FormValue("DATA")
	customerIdVal, _ := strconv.Atoi(req.FormValue("CUSTOMERID"))
	WasteLibrary.LogStr("Insert BulkDb : " + appTypeVal + " - " + didVal + " - " + dataTypeVal + " - " + timeVal + " - " + dataVal + "-" + string(customerIdVal))
	var insertSQL string = fmt.Sprintf(`INSERT INTO public.listenerdata(
		app_type,serial_number,data_type,data,customer_id,data_time)
	   VALUES ('%s','%s','%s','%s',%d,'%s');`, appTypeVal, didVal, dataTypeVal, dataVal, customerIdVal, timeVal)
	WasteLibrary.LogStr(insertSQL)
	_, errDb := bulkDb.Exec(insertSQL)
	WasteLibrary.LogErr(errDb)
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

	appTypeVal := req.FormValue("APPTYPE")
	opTypeVal := req.FormValue("OPTYPE")
	WasteLibrary.LogStr(appTypeVal + " - " + opTypeVal)
	if appTypeVal == "RFID" {

		didVal := req.FormValue("DID")
		dataTypeVal := req.FormValue("DATATYPE")
		dataTime := req.FormValue("TIME")
		repeat := req.FormValue("REPEAT")
		customerIdVal, _ := strconv.Atoi(req.FormValue("CUSTOMERID"))

		var execSQL string = ""
		if opTypeVal == "RF" {

			tagId := req.FormValue("TAGID")
			uId := req.FormValue("UID")
			latitude := req.FormValue("LATITUDE")
			longitude := req.FormValue("LONGITUDE")
			WasteLibrary.LogStr(repeat + " - " + dataTypeVal + " - " + latitude + " - " + longitude)

			var fLatitude float64 = 0
			var fLongitude float64 = 0
			if latitude != "" {
				if s, err := strconv.ParseFloat(latitude, 32); err == nil {
					fLatitude = s
				}
			}
			if longitude != "" {
				if s, err := strconv.ParseFloat(longitude, 32); err == nil {
					fLongitude = s
				}
			}

			var selectSQL string = fmt.Sprintf(`SELECT tag_id
			FROM public.tags WHERE epc='%s' AND customer_id=%d;`, tagId, customerIdVal)
			rows, errSel := staticDb.Query(selectSQL)
			WasteLibrary.LogErr(errSel)
			var tagsID int = 0
			for rows.Next() {
				rows.Scan(&tagsID)
			}
			if tagsID != 0 {
				if latitude == "" || longitude == "" {
					execSQL = fmt.Sprintf(`UPDATE public.tags
					SET uid='%s',read_time='%s',statu='1'
				   WHERE epc='%s' AND customer_id=%d;`, uId, dataTime, tagId, customerIdVal)
					WasteLibrary.LogStr(execSQL)
				} else {
					execSQL = fmt.Sprintf(`UPDATE public.tags
					SET uid='%s',read_time='%s',statu='1',latitude=%f,longitude=%f,gps_time='%s'
				   WHERE epc='%s' AND customer_id=%d;`, uId, dataTime, fLatitude, fLongitude, dataTime, tagId, customerIdVal)
					WasteLibrary.LogStr(execSQL)
				}
			} else {

				execSQL = fmt.Sprintf(`INSERT INTO public.tags(
			   app_type,serial_number,customer_id,epc,uid,container_no,latitude,longitude,statu)
			  VALUES ('%s','%s',%d,'%s','%s','%s',%f,%f,'1');`, appTypeVal, didVal, customerIdVal, tagId, uId, tagId, fLatitude, fLongitude)
				WasteLibrary.LogStr(execSQL)
			}

		} else if opTypeVal == "GPS" {

			latitude := req.FormValue("LATITUDE")
			longitude := req.FormValue("LONGITUDE")
			WasteLibrary.LogStr(repeat + " - " + dataTypeVal + " - " + latitude + " - " + longitude)

			var fLatitude float64 = 0
			var fLongitude float64 = 0
			if latitude != "" {
				if s, err := strconv.ParseFloat(latitude, 32); err == nil {
					fLatitude = s
				}
			}
			if longitude != "" {
				if s, err := strconv.ParseFloat(longitude, 32); err == nil {
					fLongitude = s
				}
			}

			execSQL = fmt.Sprintf(`UPDATE public.devices
			   SET gps_time='%s',latitude=%f,longitude=%f
			   WHERE serial_number='%s' AND customer_id=%d;`, dataTime, fLatitude, fLongitude, didVal, customerIdVal)
			WasteLibrary.LogStr(execSQL)

		} else if opTypeVal == "STATUS" {

			readerAppStatus := req.FormValue("READERAPPSTATUS")
			readerConnStatus := req.FormValue("READERCONNSTATUS")
			readerStatus := req.FormValue("READERSTATUS")
			camAppStatus := req.FormValue("CAMAPPSTATUS")
			camConnStatus := req.FormValue("CAMCONNSTATUS")
			camStatus := req.FormValue("CAMSTATUS")
			gpsAppStatus := req.FormValue("GPSAPPSTATUS")
			gpsConnStatus := req.FormValue("GPSCONNSTATUS")
			gpsStatus := req.FormValue("GPSSTATUS")
			thermAppStatus := req.FormValue("THERMAPSTATUS")
			transferAppStatus := req.FormValue("TRANSFERAPP")
			aliveStatus := req.FormValue("ALIVESTATUS")
			contactStatus := req.FormValue("CONTACTSTATUS")
			WasteLibrary.LogStr(repeat + " - " + dataTypeVal + " - " + readerAppStatus + " - " + readerConnStatus + " - " + readerStatus + " - " + camAppStatus + " - " + camConnStatus + " - " + camStatus + " - " + gpsAppStatus + " - " + gpsConnStatus + " - " + gpsStatus + " - " + thermAppStatus + " - " + transferAppStatus + " - " + aliveStatus + " - " + contactStatus)

			execSQL = fmt.Sprintf(`UPDATE public.devices
				SET status_time='%s',
				reader_app_status='%s',reader_conn_status='%s',reader_status='%s',cam_app_status='%s',cam_conn_status='%s',
				cam_status='%s',gps_app_status='%s',gps_conn_status='%s',gps_status='%s',therm_app_status='%s',
				transfer_app_status='%s',alive_status='%s',contact_status='%s'
			   WHERE serial_number='%s' AND customer_id=%d;
			   `, dataTime, readerAppStatus, readerConnStatus, readerStatus, camAppStatus, camConnStatus, camStatus, gpsAppStatus, gpsConnStatus, gpsStatus, thermAppStatus, transferAppStatus, aliveStatus, contactStatus, didVal, customerIdVal)
			WasteLibrary.LogStr(execSQL)

		} else if opTypeVal == "THERM" {

			therm := req.FormValue("THERM")
			WasteLibrary.LogStr(repeat + " - " + dataTypeVal + " - " + therm)

			execSQL = fmt.Sprintf(`UPDATE public.devices
			   SET therm='%s',therm_time='%s'
			   WHERE serial_number='%s' AND customer_id=%d;`, therm, dataTime, didVal, customerIdVal)
			WasteLibrary.LogStr(execSQL)

		} else if opTypeVal == "CAM_IMAGE" {

			uId := req.FormValue("UID")
			tagId := req.FormValue("TAGID")
			imageStatu := req.FormValue("IMAGE")
			WasteLibrary.LogStr(repeat + " - " + dataTypeVal + " - " + tagId + " - " + uId + " - " + imageStatu)

			execSQL = fmt.Sprintf(`UPDATE public.tags
			   SET image_statu='%s'
			   WHERE tag_id='%s' AND customer_id=%d;`, imageStatu, tagId, customerIdVal)
			WasteLibrary.LogStr(execSQL)

		} else {
			resultVal.Result = "FAIL"
		}

		if execSQL != "" {
			_, errDb := staticDb.Exec(execSQL)
			WasteLibrary.LogErr(errDb)
			if errDb != nil {
				WasteLibrary.LogErr(err)
			} else {
				resultVal.Result = "OK"
			}
		}
	} else if appTypeVal == "ULT" {
		resultVal.Result = "OK"
	} else if appTypeVal == "RECY" {
		resultVal.Result = "OK"
	} else if appTypeVal == "ADMIN" {
		var execSQL string = ""
		if opTypeVal == "CUSTOMER" {

			currentCustomer := WasteLibrary.StringToCustomerType(req.FormValue("DATA"))
			WasteLibrary.LogStr(currentCustomer.ToIdString() + " - " + currentCustomer.CustomerName + " - " + currentCustomer.Domain + " - " + currentCustomer.RfIdApp + " - " + currentCustomer.UltApp + " - " + currentCustomer.RecyApp)

			if currentCustomer.CustomerId != 0 {
				execSQL = fmt.Sprintf(`UPDATE public.customers
					SET customer_name='%s',domain='%s',rfid_app='%s',ult_app='%s',recy_app='%s'
				   WHERE customer_id=%f  RETURNING customer_id;`, currentCustomer.CustomerName, currentCustomer.Domain, currentCustomer.RfIdApp, currentCustomer.UltApp, currentCustomer.RecyApp, currentCustomer.CustomerId)
				WasteLibrary.LogStr(execSQL)
			} else {

				execSQL = fmt.Sprintf(`INSERT INTO public.customers(
					customer_name,domain,rfid_app,ult_app,recy_app)
			  VALUES ('%s','%s','%s','%s','%s')  RETURNING customer_id;`, currentCustomer.CustomerName, currentCustomer.Domain, currentCustomer.RfIdApp, currentCustomer.UltApp, currentCustomer.RecyApp)
				WasteLibrary.LogStr(execSQL)
			}
			var id int = 0
			errDb := staticDb.QueryRow(execSQL).Scan(&id)
			WasteLibrary.LogErr(errDb)
			if errDb != nil {
				WasteLibrary.LogErr(err)
			}

			var customerRetVal WasteLibrary.CustomerType
			var selectSQL string = fmt.Sprintf(`SELECT 
			customer_id,customer_name,domain,rfid_app,ult_app,recy_app,active,create_time
			 FROM public.customers WHERE customer_id=%d;`, id)
			rows, errSel := staticDb.Query(selectSQL)
			WasteLibrary.LogErr(errSel)
			for rows.Next() {
				rows.Scan(&customerRetVal.CustomerId, &customerRetVal.CustomerName, &customerRetVal.Domain, &customerRetVal.RfIdApp, &customerRetVal.UltApp,
					&customerRetVal.RecyApp, &customerRetVal.Active, &customerRetVal.CreateTime)
			}
			if customerRetVal.CustomerId != 0 {
				resultVal.Result = "OK"
				resultVal.Retval = customerRetVal.ToString()
			}

			WasteLibrary.LogStr("CustomerType : " + customerRetVal.ToString())

		} else if opTypeVal == "DEVICE" {

			currentDevice := WasteLibrary.StringToDeviceType(req.FormValue("DATA"))
			WasteLibrary.LogStr(currentDevice.ToIdString() + " - " + currentDevice.DeviceType + " - " + currentDevice.SerialNumber + " - " + currentDevice.DeviceName + " - " + currentDevice.ToCustomerIdString())

			if currentDevice.DeviceId != 0 {
				execSQL = fmt.Sprintf(`UPDATE public.devices 
				SET device_type='%s',serial_number='%s',device_name='%s',customer_id=%f 
	  			WHERE device_id=%f  
				RETURNING device_id;`, currentDevice.DeviceType, currentDevice.SerialNumber, currentDevice.DeviceName, currentDevice.CustomerId, currentDevice.DeviceId)
				WasteLibrary.LogStr(execSQL)
			} else {

				execSQL = fmt.Sprintf(`INSERT INTO public.devices 
				(device_type,serial_number,device_name,customer_id) 
  				VALUES ('%s','%s','%s',%f)   
  				RETURNING device_id;`, currentDevice.DeviceType, currentDevice.SerialNumber, currentDevice.DeviceName, currentDevice.CustomerId)
				WasteLibrary.LogStr(execSQL)
			}
			var id int = 0
			errDb := staticDb.QueryRow(execSQL).Scan(&id)
			WasteLibrary.LogErr(errDb)
			if errDb != nil {
				WasteLibrary.LogErr(err)
			}

			var deviceRetVal WasteLibrary.DeviceType
			var selectSQL string = fmt.Sprintf(`SELECT 
device_id,device_type,serial_number,device_name,customer_id,active,create_time
 FROM public.devices WHERE device_id=%d;`, id)
			rows, errSel := staticDb.Query(selectSQL)
			WasteLibrary.LogErr(errSel)
			for rows.Next() {
				rows.Scan(&deviceRetVal.DeviceId, &deviceRetVal.DeviceType, &deviceRetVal.SerialNumber, &deviceRetVal.DeviceName,
					&deviceRetVal.CustomerId, &deviceRetVal.Active, &deviceRetVal.CreateTime)
			}
			if deviceRetVal.DeviceId != 0 {
				resultVal.Result = "OK"
				resultVal.Retval = deviceRetVal.ToString()
			}

			WasteLibrary.LogStr("DeviceType : " + deviceRetVal.ToString())

		} else {
			resultVal.Result = "FAIL"
		}
	} else {
		resultVal.Result = "OK"
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
