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
			var systemDate WasteLibrary.SystemDateType
			systemDate.New()
			systemDate.GetByRedis()
			systemDate.LastDay++
			if systemDate.LastDay == 31 {
				systemDate.LastDay = 1
			}

			systemDate.DayDates[0] = WasteLibrary.TimeToString(time.Now())
			systemDate.DayDates[systemDate.LastDay] = WasteLibrary.TimeToString(time.Now().Add(-24 * time.Hour))
			systemDate.SaveToRedis()

			resultVal = WasteLibrary.GetKeyListRedisForStoreApi("*")
			if resultVal.Result == WasteLibrary.RESULT_OK {
				for _, hKey := range resultVal.Retval.([]string) {

					if strings.Contains(hKey, "-reel") {
						hBaseKey := strings.Replace(hKey, "-reel", "", -1)
						WasteLibrary.CloneRedisWODbForStoreApi("0", systemDate.ToLastDayString(), hKey, hBaseKey)
					} else {
						WasteLibrary.CloneRedisForStoreApi("0", systemDate.ToLastDayString(), hKey)
					}
				}

			}

		}

		time.Sleep(opInterval * time.Second)

	}
}
