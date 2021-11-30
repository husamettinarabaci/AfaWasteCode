package main

import (
	"context"
	"fmt"

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
	var val []string

	val, err = redisDb1.Keys(ctx, "*").Result()

	fmt.Println(val)

}
