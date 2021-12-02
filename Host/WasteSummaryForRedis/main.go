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

	go summaryRedis()

	http.HandleFunc("/health", WasteLibrary.HealthHandler)
	http.HandleFunc("/readiness", WasteLibrary.ReadinessHandler)
	http.HandleFunc("/status", WasteLibrary.StatusHandler)
	http.ListenAndServe(":80", nil)
}

func summaryRedis() {
	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	for {

		if time.Now().Hour() == 0 {
			var redisDbDate WasteLibrary.RedisDbDateType
			redisDbDate.New()
			redisDbDate.GetByRedis()
			redisDbDate.LastDay++
			if redisDbDate.LastDay == 31 {
				redisDbDate.LastDay = 1
			}

			redisDbDate.DayDates[0] = WasteLibrary.TimeToString(time.Now())
			redisDbDate.DayDates[redisDbDate.LastDay] = WasteLibrary.TimeToString(time.Now().Add(-24 * time.Hour))
			redisDbDate.SaveToRedis()

			resultVal = WasteLibrary.GetKeyListRedisForStoreApi("hsm-*")
			if resultVal.Result == WasteLibrary.RESULT_OK {
				for _, hKey := range resultVal.Retval.([]string) {

					if strings.Contains(hKey, "-reel") {
						hBaseKey := strings.Replace(hKey, "-reel", "", -1)
						WasteLibrary.CloneRedisWODbForStoreApi("0", redisDbDate.ToLastDayString(), hKey, hBaseKey)
					} else {
						WasteLibrary.CloneRedisForStoreApi("0", redisDbDate.ToLastDayString(), hKey)
					}
				}

			}

		}

		time.Sleep(opInterval * time.Second)

	}
}
