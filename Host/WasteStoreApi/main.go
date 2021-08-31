package main

import (
	"context"
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
)

var debug bool = os.Getenv("DEBUG") == "1"
var port int = 5432
var user string = os.Getenv("POSTGRES_USER")
var password string = os.Getenv("POSTGRES_PASSWORD")
var dbname string = os.Getenv("POSTGRES_DB")
var appStatus string = "1"
var redisDb *redis.Client

var ctx = context.Background()
var sumDb *sql.DB
var bulkDb *sql.DB
var staticDb *sql.DB
var err error

func initStart() {

	logStr("Successfully connected!")
	redisDb = redis.NewClient(&redis.Options{
		Addr:     "waste-redis-cluster-ip:6379",
		Password: "",
		DB:       0,
	})

	pong, err := redisDb.Ping(ctx).Result()
	logErr(err)
	logStr(pong)

}

func main() {

	initStart()

	var sumDbHost string = "waste-sumdb-cluster-ip"
	sumdDbInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		sumDbHost, port, user, password, dbname)

	sumDb, err = sql.Open("postgres", sumdDbInfo)
	logErr(err)
	defer sumDb.Close()

	err = sumDb.Ping()
	logErr(err)

	var bulkDbHost string = "waste-bulkdb-cluster-ip"
	bulkDbInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		bulkDbHost, port, user, password, dbname)

	bulkDb, err = sql.Open("postgres", bulkDbInfo)
	logErr(err)
	defer bulkDb.Close()

	err = bulkDb.Ping()
	logErr(err)

	var staticDbHost string = "waste-staticdb-cluster-ip"
	staticDbInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		staticDbHost, port, user, password, dbname)

	staticDb, err = sql.Open("postgres", staticDbInfo)
	logErr(err)
	defer staticDb.Close()

	err = staticDb.Ping()
	logErr(err)

	logStr("Start")
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/readiness", readinessHandler)
	http.HandleFunc("/status", status)
	http.HandleFunc("/getkey", getkey)
	http.HandleFunc("/setkey", setkey)
	http.HandleFunc("/saveBulkDbMain", saveBulkDbMain)
	http.HandleFunc("/saveStaticDbMain", saveStaticDbMain)
	http.ListenAndServe(":80", nil)
	logStr("Finished")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func readinessHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func status(w http.ResponseWriter, req *http.Request) {

	if err := req.ParseForm(); err != nil {
		logErr(err)
		return
	}
	opType := req.FormValue("OPTYPE")
	logStr(opType)

	if opType == "TYPE" {
		w.Write([]byte("WasteStoreApi"))
	} else if opType == "APP" {
		if appStatus == "1" {
			w.Write([]byte("OK"))
		} else {
			w.Write([]byte("FAIL"))
		}
	} else {
		w.Write([]byte("FAIL"))
	}
}

func saveBulkDbMain(w http.ResponseWriter, req *http.Request) {

	if err := req.ParseForm(); err != nil {
		logErr(err)
		return
	}
	appTypeVal := req.FormValue("APPTYPE")
	didVal := req.FormValue("DID")
	dataTypeVal := req.FormValue("DATATYPE")
	timeVal := req.FormValue("TIME")
	dataVal := req.FormValue("DATA")
	customerIdVal, _ := strconv.Atoi(req.FormValue("CUSTOMERID"))
	logStr("Insert BulkDb : " + appTypeVal + " - " + didVal + " - " + dataTypeVal + " - " + timeVal + " - " + dataVal + "-" + string(customerIdVal))
	var insertSQL string = fmt.Sprintf(`INSERT INTO public.listenerdata(
		app_type,serial_number,data_type,data,customer_id,data_time)
	   VALUES ('%s','%s','%s','%s',%d,'%s');`, appTypeVal, didVal, dataTypeVal, dataVal, customerIdVal, timeVal)
	logStr(insertSQL)
	_, errDb := bulkDb.Exec(insertSQL)
	logErr(errDb)
	if errDb != nil {
		logErr(err)
		w.Write([]byte("FAIL"))
	} else {
		w.Write([]byte("OK"))
	}

}

func saveStaticDbMain(w http.ResponseWriter, req *http.Request) {

	var retVal string = "FAIL"
	if err := req.ParseForm(); err != nil {
		logErr(err)
		return
	}

	appTypeVal := req.FormValue("APPTYPE")
	opTypeVal := req.FormValue("OPTYPE")
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
			logStr(repeat + " - " + dataTypeVal + " - " + latitude + " - " + longitude)

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
			logErr(errSel)
			var tagsID int = 0
			for rows.Next() {
				rows.Scan(&tagsID)
			}
			if tagsID != 0 {
				if latitude == "" || longitude == "" {
					execSQL = fmt.Sprintf(`UPDATE public.tags
					SET uid='%s',read_time='%s',statu='1'
				   WHERE epc='%s' AND customer_id=%d;`, uId, dataTime, tagId, customerIdVal)
					logStr(execSQL)
				} else {
					execSQL = fmt.Sprintf(`UPDATE public.tags
					SET uid='%s',read_time='%s',statu='1',latitude=%f,longitude=%f,gps_time='%s'
				   WHERE epc='%s' AND customer_id=%d;`, uId, dataTime, fLatitude, fLongitude, dataTime, tagId, customerIdVal)
					logStr(execSQL)
				}
			} else {

				execSQL = fmt.Sprintf(`INSERT INTO public.tags(
			   app_type,serial_number,customer_id,epc,uid,container_no,latitude,longitude,statu)
			  VALUES ('%s','%s',%d,'%s','%s','%s',%f,%f,'1');`, appTypeVal, didVal, customerIdVal, tagId, uId, tagId, fLatitude, fLongitude)
				logStr(execSQL)
			}

		} else if opTypeVal == "GPS" {

			latitude := req.FormValue("LATITUDE")
			longitude := req.FormValue("LONGITUDE")
			logStr(repeat + " - " + dataTypeVal + " - " + latitude + " - " + longitude)

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
			logStr(execSQL)

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
			logStr(repeat + " - " + dataTypeVal + " - " + readerAppStatus + " - " + readerConnStatus + " - " + readerStatus + " - " + camAppStatus + " - " + camConnStatus + " - " + camStatus + " - " + gpsAppStatus + " - " + gpsConnStatus + " - " + gpsStatus + " - " + thermAppStatus + " - " + transferAppStatus + " - " + aliveStatus + " - " + contactStatus)

			execSQL = fmt.Sprintf(`UPDATE public.devices
				SET status_time='%s',
				reader_app_status='%s',reader_conn_status='%s',reader_status='%s',cam_app_status='%s',cam_conn_status='%s',
				cam_status='%s',gps_app_status='%s',gps_conn_status='%s',gps_status='%s',therm_app_status='%s',
				transfer_app_status='%s',alive_status='%s',contact_status='%s'
			   WHERE serial_number='%s' AND customer_id=%d;
			   `, dataTime, readerAppStatus, readerConnStatus, readerStatus, camAppStatus, camConnStatus, camStatus, gpsAppStatus, gpsConnStatus, gpsStatus, thermAppStatus, transferAppStatus, aliveStatus, contactStatus, didVal, customerIdVal)
			logStr(execSQL)

		} else if opTypeVal == "THERM" {

			therm := req.FormValue("THERM")
			logStr(repeat + " - " + dataTypeVal + " - " + therm)

			execSQL = fmt.Sprintf(`UPDATE public.devices
			   SET therm='%s',therm_time='%s'
			   WHERE serial_number='%s' AND customer_id=%d;`, therm, dataTime, didVal, customerIdVal)
			logStr(execSQL)

		} else if opTypeVal == "CAM_IMAGE" {

			uId := req.FormValue("UID")
			tagId := req.FormValue("TAGID")
			imageStatu := req.FormValue("IMAGE")
			logStr(repeat + " - " + dataTypeVal + " - " + tagId + " - " + uId + " - " + imageStatu)

			execSQL = fmt.Sprintf(`UPDATE public.tags
			   SET image_statu='%s'
			   WHERE tag_id='%s' AND customer_id=%d;`, imageStatu, tagId, customerIdVal)
			logStr(execSQL)

		} else {
			retVal = "FAIL"
		}

		if execSQL != "" {
			_, errDb := staticDb.Exec(execSQL)
			logErr(errDb)
			if errDb != nil {
				logErr(err)
			} else {
				retVal = "OK"
			}
		}
		w.Write([]byte(retVal))
	} else if appTypeVal == "ULT" {
		w.Write([]byte(retVal))
	} else if appTypeVal == "RECY" {
		w.Write([]byte(retVal))
	} else if appTypeVal == "ADMIN" {
		var execSQL string = ""
		if opTypeVal == "CUSTOMER" {

			customerId, _ := strconv.Atoi(req.FormValue("CUSTOMERID"))
			customerName := req.FormValue("NAME")
			domain := req.FormValue("DOMAIN")
			rfidApp := req.FormValue("RFIDAPP")
			ultApp := req.FormValue("ULTAPP")
			recyApp := req.FormValue("RECYAPP")
			logStr(string(customerId) + " - " + customerName + " - " + domain + " - " + rfidApp + " - " + ultApp + " - " + recyApp)

			if customerId != 0 {
				execSQL = fmt.Sprintf(`UPDATE public.customers
					SET customer_name='%s',domain='%s',rfid_app='%s',ult_app='%s',recy_app='%s'
				   WHERE customer_id=%d;`, customerName, domain, rfidApp, ultApp, recyApp, customerId)
				logStr(execSQL)
			} else {

				execSQL = fmt.Sprintf(`INSERT INTO public.customers(
					customer_name,domain,rfid_app,ult_app,recy_app)
			  VALUES ('%s','%s','%s','%s','%s');`, customerName, domain, rfidApp, ultApp, recyApp)
				logStr(execSQL)
			}

		} else if opTypeVal == "DEVICE" {

			deviceId, _ := strconv.Atoi(req.FormValue("DEVICEID"))
			customerId, _ := strconv.Atoi(req.FormValue("CUSTOMERID"))
			deviceName := req.FormValue("NAME")
			serialNumber := req.FormValue("SERIALNO")
			deviceType := req.FormValue("DEVICETYPE")
			logStr(string(deviceId) + " - " + string(customerId) + " - " + deviceName + " - " + deviceType)

			if deviceId != 0 {
				execSQL = fmt.Sprintf(`UPDATE public.devices
					SET device_type='%s',serial_number='%s',device_name='%s',customer_id=%d
				   WHERE device_id=%d;`, deviceType, serialNumber, deviceName, customerId, deviceId)
				logStr(execSQL)
			} else {

				execSQL = fmt.Sprintf(`INSERT INTO public.devices(
					device_type,serial_number,device_name,customer_id)
			  VALUES ('%s','%s','%s',%d);`, deviceType, serialNumber, deviceName, customerId)
				logStr(execSQL)
			}

		} else {
			retVal = "FAIL"
		}
		if execSQL != "" {
			_, errDb := staticDb.Exec(execSQL)
			logErr(errDb)
			if errDb != nil {
				logErr(err)
			} else {
				retVal = "OK"
			}
		}
		w.Write([]byte(retVal))
	} else {
		w.Write([]byte(retVal))
	}

}

func getkey(w http.ResponseWriter, req *http.Request) {

	var retVal string = "FAIL"
	if err := req.ParseForm(); err != nil {
		logErr(err)
		return
	}
	hKey := req.FormValue("HASHKEY")
	sKey := req.FormValue("SUBKEY")
	logStr("GetKEy : " + hKey + " - " + sKey)
	retVal = getKeyRedis(hKey, sKey)
	logStr("RetValByRedis : " + retVal)
	if retVal == "NOT" {
		retVal = getKeyDb(hKey, sKey)
		if retVal != "NOT" && retVal != "FAIL" {
			setKeyRedis(hKey, sKey, retVal)
		}
	}
	w.Write([]byte(retVal))
}

func setkey(w http.ResponseWriter, req *http.Request) {

	var retVal string = "OK"
	if err := req.ParseForm(); err != nil {
		logErr(err)
		return
	}
	hKey := req.FormValue("HASHKEY")
	sKey := req.FormValue("SUBKEY")
	kVal := req.FormValue("KEYVALUE")

	retVal = getKeyDb(hKey, sKey)

	if retVal == "NOT" {
		insertKeyDb(hKey, sKey, kVal)
	} else {
		updateKeyDb(hKey, sKey, kVal)
	}
	setKeyRedis(hKey, sKey, kVal)

	w.Write([]byte(retVal))
}

func getKeyRedis(hKey string, sKey string) string {
	var retVal string = "FAIL"
	val, err := redisDb.HGet(ctx, hKey, sKey).Result()
	switch {
	case err == redis.Nil:
		logStr("Not Found")
		retVal = "NOT"
	case err != nil:
		logErr(err)
	case val == "":
		logStr("Not Found")
		retVal = "NOT"
	case val != "":
		retVal = val
		logStr(retVal)
	}

	return retVal
}

func setKeyRedis(hKey string, sKey string, kVal string) {
	redisDb.HSet(ctx, hKey, sKey, kVal).Result()
}

func getKeyDb(hKey string, sKey string) string {
	logStr("Serach Db : " + hKey + " - " + sKey)
	var retVal string = "FAIL"

	var selectSQL string = fmt.Sprintf(`SELECT keyvalue 
	FROM public.redisdata WHERE hashkey='%s' AND subkey='%s';`, hKey, sKey)
	rows, errSel := sumDb.Query(selectSQL)
	logErr(errSel)
	var kVal string = "NOT"
	for rows.Next() {
		rows.Scan(&kVal)
	}

	logStr("KeyValue : " + kVal)

	retVal = kVal
	return retVal
}

func insertKeyDb(hKey string, sKey string, kVal string) string {
	logStr("Insert Db : " + hKey + " - " + sKey + " - " + kVal)
	var retVal string = "OK"
	var insertSQL string = fmt.Sprintf(`INSERT INTO public.redisdata(
		hashkey,subkey,keyvalue)
	   VALUES ('%s','%s','%s');`, hKey, sKey, kVal)
	_, errDb := sumDb.Exec(insertSQL)
	logErr(errDb)
	return retVal
}

func updateKeyDb(hKey string, sKey string, kVal string) string {
	logStr("Update Db : " + hKey + " - " + sKey + " - " + kVal)
	var retVal string = "FAIL"
	var updateSQL string = fmt.Sprintf(`UPDATE public.redisdata SET keyvalue='%s' WHERE hashkey='%s' AND subkey='%s';`, kVal, hKey, sKey)
	_, errDb := sumDb.Exec(updateSQL)
	logErr(errDb)
	return retVal
}

func logErr(err error) {
	if err != nil {
		sendLogServer("ERR", err.Error())
	}
}

func logStr(value string) {
	if debug {
		sendLogServer("INFO", value)
	}
}

var container string = os.Getenv("CONTAINER_TYPE")

func sendLogServer(logType string, logVal string) string {
	var retVal string = "FAIL"
	data := url.Values{
		"CONTAINER": {container},
		"LOGTYPE":   {logType},
		"LOG":       {logVal},
	}
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.PostForm("http://waste-logserver-cluster-ip/log", data)
	if err != nil {
		logErr(err)

	} else {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logErr(err)
		}
		bodyString := string(bodyBytes)
		if bodyString != "NOT" {
			retVal = bodyString
		}
	}

	return retVal
}
