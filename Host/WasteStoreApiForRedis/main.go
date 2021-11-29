package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"

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
var err error

func initStart() {

	WasteLibrary.LogStr("Successfully connected!")
	go WasteLibrary.InitLog()
	redisDb = redis.NewClient(&redis.Options{
		Addr:     "waste-redis-cluster-ip:6379",
		Password: "Amca151200!Furkan",
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

	http.HandleFunc("/health", WasteLibrary.HealthHandler)
	http.HandleFunc("/readiness", WasteLibrary.ReadinessHandler)
	http.HandleFunc("/status", WasteLibrary.StatusHandler)
	http.HandleFunc("/publishkey", publishkey)
	http.HandleFunc("/getkey", getkey)
	http.HandleFunc("/getkeyWODb", getkeyWODb)
	http.HandleFunc("/setkey", setkey)
	http.HandleFunc("/setkeyWODb", setkeyWODb)
	http.HandleFunc("/deletekey", deletekey)
	http.ListenAndServe(":80", nil)
}

func getkey(w http.ResponseWriter, req *http.Request) {

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
	hKey := req.FormValue(WasteLibrary.REDIS_HASHKEY)
	sKey := req.FormValue(WasteLibrary.REDIS_SUBKEY)
	resultVal = getKeyRedis(hKey, sKey)
	if resultVal.Result == WasteLibrary.RESULT_FAIL {
		resultVal = getKeyDb(hKey, sKey)
		if resultVal.Result == WasteLibrary.RESULT_OK {
			setKeyRedis(hKey, sKey, resultVal.Retval.(string))
		}
	}
	w.Write(resultVal.ToByte())

}

func getkeyWODb(w http.ResponseWriter, req *http.Request) {

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
	hKey := req.FormValue(WasteLibrary.REDIS_HASHKEY)
	sKey := req.FormValue(WasteLibrary.REDIS_SUBKEY)
	hBKey := req.FormValue(WasteLibrary.REDIS_HASHBASEKEY)
	resultVal = getKeyRedis(hKey, sKey)
	if resultVal.Result == WasteLibrary.RESULT_FAIL {
		if hBKey != "" {
			resultVal = getKeyRedis(hBKey, sKey)
		}
		if resultVal.Result == WasteLibrary.RESULT_FAIL {
			resultVal = getKeyDb(hBKey, sKey)
			if resultVal.Result == WasteLibrary.RESULT_OK {
				setKeyRedis(hKey, sKey, resultVal.Retval.(string))
			}
		}
	}
	w.Write(resultVal.ToByte())

}

func publishkey(w http.ResponseWriter, req *http.Request) {

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
	channelKey := req.FormValue(WasteLibrary.REDIS_CHANNELKEY)
	kVal := req.FormValue(WasteLibrary.REDIS_KEYVALUE)
	resultVal = publishKeyRedis(channelKey, kVal)
	w.Write(resultVal.ToByte())

}

func setkey(w http.ResponseWriter, req *http.Request) {

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
	hKey := req.FormValue(WasteLibrary.REDIS_HASHKEY)
	sKey := req.FormValue(WasteLibrary.REDIS_SUBKEY)
	kVal := req.FormValue(WasteLibrary.REDIS_KEYVALUE)
	resultVal = getKeyDb(hKey, sKey)
	if resultVal.Result == WasteLibrary.RESULT_OK {
		resultVal = updateKeyDb(hKey, sKey, kVal)
	} else {
		resultVal = insertKeyDb(hKey, sKey, kVal)
	}
	setKeyRedis(hKey, sKey, kVal)

	w.Write(resultVal.ToByte())

}

func setkeyWODb(w http.ResponseWriter, req *http.Request) {

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
	hKey := req.FormValue(WasteLibrary.REDIS_HASHKEY)
	sKey := req.FormValue(WasteLibrary.REDIS_SUBKEY)
	kVal := req.FormValue(WasteLibrary.REDIS_KEYVALUE)
	setKeyRedis(hKey, sKey, kVal)

	w.Write(resultVal.ToByte())

}

func deletekey(w http.ResponseWriter, req *http.Request) {

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
	hKey := req.FormValue(WasteLibrary.REDIS_HASHKEY)
	sKey := req.FormValue(WasteLibrary.REDIS_SUBKEY)
	resultVal = deleteKeyDb(hKey, sKey)
	deleteKeyRedis(hKey, sKey)

	w.Write(resultVal.ToByte())

}

func getKeyRedis(hKey string, sKey string) WasteLibrary.ResultType {
	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	var val string = ""
	var err error
	if hKey != "" {
		val, err = redisDb.HGet(ctx, hKey, sKey).Result()
	} else {
		val, err = redisDb.Get(ctx, sKey).Result()
	}
	switch {
	case err == redis.Nil:
		resultVal.Result = WasteLibrary.RESULT_FAIL
	case err != nil:
		WasteLibrary.LogErr(err)
	case val == "":
		resultVal.Result = WasteLibrary.RESULT_FAIL
	case val != "":
		resultVal.Result = WasteLibrary.RESULT_OK
		resultVal.Retval = val
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
	case err != nil:
		WasteLibrary.LogErr(err)
	}
}

func publishKeyRedis(channelKey string, kVal string) WasteLibrary.ResultType {
	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	var err error

	_, err = redisDb.Publish(ctx, channelKey, kVal).Result()
	switch {
	case err == redis.Nil:
	case err != nil:
		WasteLibrary.LogErr(err)
	}

	if err != nil {

	} else {
		resultVal.Result = WasteLibrary.RESULT_OK
	}
	return resultVal
}

func deleteKeyRedis(hKey string, sKey string) {
	var err error
	if hKey != "" {
		_, err = redisDb.HDel(ctx, hKey, sKey).Result()
	} else {
		_, err = redisDb.HDel(ctx, sKey).Result()

	}
	switch {
	case err == redis.Nil:
	case err != nil:
		WasteLibrary.LogErr(err)
	}
}

func getKeyDb(hKey string, sKey string) WasteLibrary.ResultType {
	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL

	var selectSQL string = fmt.Sprintf(`SELECT KeyValue 
	FROM public.redisdata WHERE HashKey='%s' AND SubKey='%s';`, hKey, sKey)
	rows, errSel := sumDb.Query(selectSQL)
	WasteLibrary.LogErr(errSel)
	var kVal string = WasteLibrary.RESULT_FAIL
	for rows.Next() {
		rows.Scan(&kVal)
	}

	if kVal == WasteLibrary.RESULT_FAIL {
		resultVal.Result = WasteLibrary.RESULT_FAIL
	} else {
		resultVal.Retval = kVal
		resultVal.Result = WasteLibrary.RESULT_OK
	}
	return resultVal
}

func insertKeyDb(hKey string, sKey string, kVal string) WasteLibrary.ResultType {
	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	var insertSQL string = fmt.Sprintf(`INSERT INTO public.redisdata(
		HashKey,SubKey,KeyValue)
	   VALUES ('%s','%s','%s');`, hKey, sKey, kVal)
	_, errDb := sumDb.Exec(insertSQL)
	WasteLibrary.LogErr(errDb)
	resultVal.Result = WasteLibrary.RESULT_OK
	return resultVal
}

func deleteKeyDb(hKey string, sKey string) WasteLibrary.ResultType {
	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	var deleteSQL string = fmt.Sprintf(`DELETE FROM public.redisdata
	    WHERE HashKey='%s' AND SubKey='%s';`, hKey, sKey)
	_, errDb := sumDb.Exec(deleteSQL)
	WasteLibrary.LogErr(errDb)
	resultVal.Result = WasteLibrary.RESULT_OK
	return resultVal
}

func updateKeyDb(hKey string, sKey string, kVal string) WasteLibrary.ResultType {
	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	var updateSQL string = fmt.Sprintf(`UPDATE public.redisdata SET KeyValue='%s' WHERE HashKey='%s' AND SubKey='%s';`, kVal, hKey, sKey)
	_, errDb := sumDb.Exec(updateSQL)
	WasteLibrary.LogErr(errDb)
	resultVal.Result = WasteLibrary.RESULT_OK
	return resultVal
}
