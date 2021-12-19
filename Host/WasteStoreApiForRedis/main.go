package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/devafatek/WasteLibrary"
	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
)

var port int = 5432
var user string = os.Getenv("POSTGRES_USER")
var password string = os.Getenv("POSTGRES_PASSWORD")
var dbname string = os.Getenv("POSTGRES_DB")
var redisRClts [31]*redis.Client
var redisWClts [31]*redis.Client

var ctx = context.Background()
var sumDb *sql.DB
var err error

func initStart() {

	WasteLibrary.LogStr("Successfully connected!")
	go WasteLibrary.InitLog()

}

func setRedisClts() {
	for i := 0; i < 31; i++ {
		var redisDb *redis.Client
		redisDb = redis.NewClient(&redis.Options{
			Addr:     "waste-redis-cluster-ip:6379",
			Password: "Amca151200!Furkan",
			DB:       i,
		})

		pong, err := redisDb.Ping(ctx).Result()
		WasteLibrary.LogErr(err)
		WasteLibrary.LogStr(pong)
		redisRClts[i] = redisDb
	}

	for i := 0; i < 31; i++ {
		var redisDb *redis.Client
		redisDb = redis.NewClient(&redis.Options{
			Addr:     "waste-redis-master-cluster-ip:6379",
			Password: "Amca151200!Furkan",
			DB:       i,
		})

		pong, err := redisDb.Ping(ctx).Result()
		WasteLibrary.LogErr(err)
		WasteLibrary.LogStr(pong)
		redisWClts[i] = redisDb
	}
}

func getRedisClts(index int, isWrite bool) *redis.Client {
	if isWrite {
		return redisWClts[index]
	} else {
		return redisRClts[index]
	}
}

func main() {

	initStart()

	setRedisClts()

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
	http.HandleFunc("/getkeylist", getkeylist)
	http.HandleFunc("/clonekey", clonekey)
	http.HandleFunc("/clonekeyWODb", clonekeyWODb)
	http.HandleFunc("/setkey", setkey)
	http.HandleFunc("/setkeyWODb", setkeyWODb)
	http.HandleFunc("/deletekey", deletekey)
	http.ListenAndServe(":80", nil)
}

func getkey(w http.ResponseWriter, req *http.Request) {

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
	dbIndex, _ := strconv.Atoi(req.FormValue(WasteLibrary.REDIS_DBINDEX))
	hKey := req.FormValue(WasteLibrary.REDIS_HASHKEY)
	sKey := req.FormValue(WasteLibrary.REDIS_SUBKEY)
	resultVal = getKeyRedis(dbIndex, hKey, sKey)
	if resultVal.Result == WasteLibrary.RESULT_FAIL {
		resultVal = getKeyDb(dbIndex, hKey, sKey)
		if resultVal.Result == WasteLibrary.RESULT_OK {
			setKeyRedis(dbIndex, hKey, sKey, resultVal.Retval.(string))
		}
	}
	w.Write(resultVal.ToByte())

}

func getkeyWODb(w http.ResponseWriter, req *http.Request) {

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
	dbIndex, _ := strconv.Atoi(req.FormValue(WasteLibrary.REDIS_DBINDEX))
	hKey := req.FormValue(WasteLibrary.REDIS_HASHKEY)
	sKey := req.FormValue(WasteLibrary.REDIS_SUBKEY)
	hBKey := req.FormValue(WasteLibrary.REDIS_HASHBASEKEY)
	resultVal = getKeyRedis(dbIndex, hKey, sKey)
	if resultVal.Result == WasteLibrary.RESULT_FAIL {
		if hBKey != "" {
			resultVal = getKeyRedis(dbIndex, hBKey, sKey)
		}
		if resultVal.Result == WasteLibrary.RESULT_FAIL {
			resultVal = getKeyDb(dbIndex, hBKey, sKey)
			if resultVal.Result == WasteLibrary.RESULT_OK {
				setKeyRedis(dbIndex, hKey, sKey, resultVal.Retval.(string))
			}
		}
	}
	w.Write(resultVal.ToByte())

}

func getkeylist(w http.ResponseWriter, req *http.Request) {

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

	pattern := req.FormValue(WasteLibrary.REDIS_PATTERN)

	resultVal = getKeyListRedis(pattern)

	w.Write(resultVal.ToByte())

}

func clonekey(w http.ResponseWriter, req *http.Request) {

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
	dbIndex, _ := strconv.Atoi(req.FormValue(WasteLibrary.REDIS_DBINDEX))
	dbIndexClone, _ := strconv.Atoi(req.FormValue(WasteLibrary.REDIS_DBINDEXCLONE))
	hKey := req.FormValue(WasteLibrary.REDIS_HASHKEY)
	resultVal = getKeyAllRedis(dbIndex, hKey)
	if resultVal.Result == WasteLibrary.RESULT_OK {
		for sKey, kVal := range WasteLibrary.StringToMapStringString(resultVal.Retval.(string)) {
			resultVal = getKeyDb(dbIndexClone, hKey, sKey)
			if resultVal.Result == WasteLibrary.RESULT_OK {
				resultVal = updateKeyDb(dbIndexClone, hKey, sKey, kVal)
			} else {
				resultVal = insertKeyDb(dbIndexClone, hKey, sKey, kVal)
			}
			setKeyRedis(dbIndexClone, hKey, sKey, kVal)
		}
	}
	w.Write(resultVal.ToByte())

}

func clonekeyWODb(w http.ResponseWriter, req *http.Request) {

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
	dbIndex, _ := strconv.Atoi(req.FormValue(WasteLibrary.REDIS_DBINDEX))
	dbIndexClone, _ := strconv.Atoi(req.FormValue(WasteLibrary.REDIS_DBINDEXCLONE))
	hKey := req.FormValue(WasteLibrary.REDIS_HASHKEY)
	hBKey := req.FormValue(WasteLibrary.REDIS_HASHBASEKEY)
	resultVal = getKeyAllRedis(dbIndex, hKey)
	if resultVal.Result == WasteLibrary.RESULT_FAIL {
		if hBKey != "" {
			resultVal = getKeyAllRedis(dbIndex, hBKey)
		}

	}
	if resultVal.Result == WasteLibrary.RESULT_OK {
		for sKey, kVal := range WasteLibrary.StringToMapStringString(resultVal.Retval.(string)) {
			setKeyRedis(dbIndexClone, hKey, sKey, kVal)
		}
	}
	w.Write(resultVal.ToByte())

}

func publishkey(w http.ResponseWriter, req *http.Request) {

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
	channelKey := req.FormValue(WasteLibrary.REDIS_CHANNELKEY)
	kVal := req.FormValue(WasteLibrary.REDIS_KEYVALUE)
	resultVal = publishKeyRedis(channelKey, kVal)
	w.Write(resultVal.ToByte())

}

func setkey(w http.ResponseWriter, req *http.Request) {

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
	hKey := req.FormValue(WasteLibrary.REDIS_HASHKEY)
	sKey := req.FormValue(WasteLibrary.REDIS_SUBKEY)
	kVal := req.FormValue(WasteLibrary.REDIS_KEYVALUE)
	resultVal = getKeyDb(0, hKey, sKey)
	if resultVal.Result == WasteLibrary.RESULT_OK {
		resultVal = updateKeyDb(0, hKey, sKey, kVal)
	} else {
		resultVal = insertKeyDb(0, hKey, sKey, kVal)
	}
	setKeyRedis(0, hKey, sKey, kVal)

	w.Write(resultVal.ToByte())

}

func setkeyWODb(w http.ResponseWriter, req *http.Request) {

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
	hKey := req.FormValue(WasteLibrary.REDIS_HASHKEY)
	sKey := req.FormValue(WasteLibrary.REDIS_SUBKEY)
	kVal := req.FormValue(WasteLibrary.REDIS_KEYVALUE)
	setKeyRedis(0, hKey, sKey, kVal)

	w.Write(resultVal.ToByte())

}

func deletekey(w http.ResponseWriter, req *http.Request) {

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
	hKey := req.FormValue(WasteLibrary.REDIS_HASHKEY)
	sKey := req.FormValue(WasteLibrary.REDIS_SUBKEY)
	resultVal = deleteKeyDb(0, hKey, sKey)
	deleteKeyRedis(0, hKey, sKey)

	w.Write(resultVal.ToByte())

}

func getKeyListRedis(pattern string) WasteLibrary.ResultType {
	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	var val []string
	var err error
	val, err = getRedisClts(0, false).Keys(ctx, pattern).Result()
	switch {
	case err == redis.Nil:
		resultVal.Result = WasteLibrary.RESULT_FAIL
	case err != nil:
		WasteLibrary.LogErr(err)
	case len(val) == 0:
		resultVal.Result = WasteLibrary.RESULT_FAIL
	case len(val) != 0:
		resultVal.Result = WasteLibrary.RESULT_OK
		resultVal.Retval = WasteLibrary.StringArrayToString(val)
	}

	return resultVal
}

func getKeyRedis(dbIndex int, hKey string, sKey string) WasteLibrary.ResultType {
	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	var val string = ""
	var err error
	if hKey != "" {
		val, err = getRedisClts(dbIndex, false).HGet(ctx, hKey, sKey).Result()
	} else {
		val, err = getRedisClts(dbIndex, false).Get(ctx, sKey).Result()
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

func getKeyAllRedis(dbIndex int, hKey string) WasteLibrary.ResultType {
	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	var val map[string]string
	var err error
	val, err = getRedisClts(dbIndex, false).HGetAll(ctx, hKey).Result()

	switch {
	case err == redis.Nil:
		resultVal.Result = WasteLibrary.RESULT_FAIL
	case err != nil:
		WasteLibrary.LogErr(err)
	case len(val) == 0:
		resultVal.Result = WasteLibrary.RESULT_FAIL
	case len(val) != 0:
		resultVal.Result = WasteLibrary.RESULT_OK
		resultVal.Retval = WasteLibrary.MapStringStringToString(val)
	}

	return resultVal
}

func setKeyRedis(dbIndex int, hKey string, sKey string, kVal string) {
	var err error
	if hKey != "" {
		_, err = getRedisClts(dbIndex, true).HSet(ctx, hKey, sKey, kVal).Result()
	} else {
		_, err = getRedisClts(dbIndex, true).HSet(ctx, sKey, kVal).Result()

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

	_, err = getRedisClts(0, true).Publish(ctx, channelKey, kVal).Result()
	switch {
	case err == redis.Nil:
	case err != nil:
		WasteLibrary.LogErr(err)
	}

	if err != nil {

	} else {
		resultVal.Result = WasteLibrary.RESULT_OK
		resultVal.Retval = WasteLibrary.RESULT_SUCCESS_OK
	}
	return resultVal
}

func deleteKeyRedis(dbIndex int, hKey string, sKey string) {
	var err error
	if hKey != "" {
		_, err = getRedisClts(dbIndex, true).HDel(ctx, hKey, sKey).Result()
	} else {
		_, err = getRedisClts(dbIndex, true).HDel(ctx, sKey).Result()

	}
	switch {
	case err == redis.Nil:
	case err != nil:
		WasteLibrary.LogErr(err)
	}
}

func getKeyDb(dbIndex int, hKey string, sKey string) WasteLibrary.ResultType {
	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL

	var selectSQL string = fmt.Sprintf(`SELECT KeyValue 
	FROM public.`+WasteLibrary.DATATYPE_REDIS_DATA+`_%d WHERE HashKey='%s' AND SubKey='%s';`, dbIndex, hKey, sKey)
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

func insertKeyDb(dbIndex int, hKey string, sKey string, kVal string) WasteLibrary.ResultType {
	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	var insertSQL string = fmt.Sprintf(`INSERT INTO public.`+WasteLibrary.DATATYPE_REDIS_DATA+`_%d (
		HashKey,SubKey,KeyValue)
	   VALUES ('%s','%s','%s');`, dbIndex, hKey, sKey, kVal)
	_, errDb := sumDb.Exec(insertSQL)
	WasteLibrary.LogErr(errDb)
	resultVal.Result = WasteLibrary.RESULT_OK
	resultVal.Retval = WasteLibrary.RESULT_SUCCESS_OK
	return resultVal
}

func deleteKeyDb(dbIndex int, hKey string, sKey string) WasteLibrary.ResultType {
	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	var deleteSQL string = fmt.Sprintf(`DELETE FROM public.`+WasteLibrary.DATATYPE_REDIS_DATA+`_%d 
	    WHERE HashKey='%s' AND SubKey='%s';`, dbIndex, hKey, sKey)
	_, errDb := sumDb.Exec(deleteSQL)
	WasteLibrary.LogErr(errDb)
	resultVal.Result = WasteLibrary.RESULT_OK
	resultVal.Retval = WasteLibrary.RESULT_SUCCESS_OK
	return resultVal
}

func updateKeyDb(dbIndex int, hKey string, sKey string, kVal string) WasteLibrary.ResultType {
	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	var updateSQL string = fmt.Sprintf(`UPDATE public.`+WasteLibrary.DATATYPE_REDIS_DATA+`_%d SET KeyValue='%s' WHERE HashKey='%s' AND SubKey='%s';`, dbIndex, kVal, hKey, sKey)
	_, errDb := sumDb.Exec(updateSQL)
	WasteLibrary.LogErr(errDb)
	resultVal.Result = WasteLibrary.RESULT_OK
	resultVal.Retval = WasteLibrary.RESULT_SUCCESS_OK
	return resultVal
}
