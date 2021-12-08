package main

import (
	"context"

	"github.com/devafatek/WasteLibrary"
	"github.com/go-redis/redis/v8"
)

var redisRClts [31]*redis.Client
var redisWClts [31]*redis.Client

var ctx = context.Background()

func main() {

	setRedisClts()

}

func setRedisClts() {
	for i := 0; i < 31; i++ {
		var redisDb *redis.Client
		redisDb = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
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
			Addr:     "localhost:6379",
			Password: "Amca151200!Furkan",
			DB:       i,
		})

		pong, err := redisDb.Ping(ctx).Result()
		WasteLibrary.LogErr(err)
		WasteLibrary.LogStr(pong)
		redisWClts[i] = redisDb
	}
}
