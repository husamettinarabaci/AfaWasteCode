package main

import (
	"context"

	"github.com/devafatek/WasteLibrary"
	"github.com/go-redis/redis/v8"
)

var redisDb *redis.Client
var redisDb1 *redis.Client

var ctx = context.Background()

func main() {

	WasteLibrary.Debug = true
	redisDb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "Amca151200!Furkan",
		DB:       0,
	})

	pong, err := redisDb.Ping(ctx).Result()
	WasteLibrary.LogErr(err)
	WasteLibrary.LogStr(pong)

	redisDb1 = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "Amca151200!Furkan",
		DB:       1,
	})

	pong, err = redisDb1.Ping(ctx).Result()
	WasteLibrary.LogErr(err)
	WasteLibrary.LogStr(pong)

	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	var val map[string]string

	val, err = redisDb.HGetAll(ctx, "customers").Result()

	switch {
	case err == redis.Nil:
		resultVal.Result = WasteLibrary.RESULT_FAIL
	case err != nil:
		WasteLibrary.LogErr(err)
	case len(val) == 0:
		resultVal.Result = WasteLibrary.RESULT_FAIL
	case len(val) != 0:
		resultVal.Result = WasteLibrary.RESULT_OK
		resultVal.Retval = val
	}

	_, err = redisDb1.HMSet(ctx, "customers", resultVal.Retval.(map[string]string)).Result()
	switch {
	case err == redis.Nil:
	case err != nil:
		WasteLibrary.LogErr(err)
	}

	WasteLibrary.LogStr(resultVal.ToString())

}
