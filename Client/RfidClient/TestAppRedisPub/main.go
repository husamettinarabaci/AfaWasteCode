package main

import (
	"context"

	"github.com/devafatek/WasteLibrary"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var redisDb *redis.Client

func initStart() {

	WasteLibrary.LogStr("Successfully connected!")
	redisDb = redis.NewClient(&redis.Options{
		Addr:     "redis.aws.afatek.com.tr:6379",
		Password: "Amca151200!Furkan",
		DB:       0,
	})

	pong, err := redisDb.Ping(ctx).Result()
	WasteLibrary.LogErr(err)
	WasteLibrary.LogStr(pong)

}

func main() {

	initStart()

	if err := redisDb.Publish(ctx, "customer-3-data", "Hello").Err(); err != nil {
		panic(err)
	}

}
