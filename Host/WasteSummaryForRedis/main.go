package main

import (
	"net/http"
	"strings"
	"time"

	"github.com/devafatek/WasteLibrary"
)

var opInterval time.Duration = 60 * 60

func initStart() {

	WasteLibrary.LogStr("Successfully connected!")
	go WasteLibrary.InitLog()
}

func main() {

	initStart()

	WasteLibrary.Debug = true
	go summaryRedis()

	http.HandleFunc("/health", WasteLibrary.HealthHandler)
	http.HandleFunc("/readiness", WasteLibrary.ReadinessHandler)
	http.HandleFunc("/status", WasteLibrary.StatusHandler)
	http.ListenAndServe(":80", nil)
}

func summaryRedis() {
	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	WasteLibrary.LogStr("Started")
	for {

		if time.Now().Hour() == 13 || time.Now().Hour() == 14 || time.Now().Hour() == 15 || time.Now().Hour() == 16 || time.Now().Hour() == 17 {
			var redisDbDate WasteLibrary.RedisDbDateType
			redisDbDate.New()
			redisDbDate.GetByRedis()
			redisDbDate.LastDay++
			WasteLibrary.LogStr(redisDbDate.ToString())
			if redisDbDate.LastDay == 31 {
				redisDbDate.LastDay = 1
			}

			redisDbDate.DayDates[0] = WasteLibrary.TimeToString(time.Now())
			redisDbDate.DayDates[redisDbDate.LastDay] = WasteLibrary.TimeToString(time.Now().Add(-24 * time.Hour))
			redisDbDate.SaveToRedis()
			WasteLibrary.LogStr(redisDbDate.ToString())
			resultVal = WasteLibrary.GetKeyListRedisForStoreApi("hsm-*")
			WasteLibrary.LogStr(resultVal.ToString())

			if resultVal.Result == WasteLibrary.RESULT_OK {
				for _, hKey := range resultVal.Retval.([]string) {
					var inResultVal WasteLibrary.ResultType
					WasteLibrary.LogStr(hKey)
					if strings.Contains(hKey, "-reel") {
						hBaseKey := strings.Replace(hKey, "-reel", "", -1)
						inResultVal = WasteLibrary.CloneRedisWODbForStoreApi("0", redisDbDate.ToLastDayString(), hKey, hBaseKey)
					} else {
						inResultVal = WasteLibrary.CloneRedisForStoreApi("0", redisDbDate.ToLastDayString(), hKey)
					}
					WasteLibrary.LogStr(inResultVal.ToString())
				}

			}

		}

		time.Sleep(opInterval * time.Second)

	}
}
