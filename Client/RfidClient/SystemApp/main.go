package main

import (
	"net/http"
	"sync"
	"time"

	"github.com/devafatek/WasteLibrary"
)

var currentUser string
var opInterval time.Duration = 60 * 60
var wg sync.WaitGroup

func initStart() {

	time.Sleep(5 * time.Second)
	WasteLibrary.LogStr("Successfully connected!")
	WasteLibrary.Version = "1"
	WasteLibrary.LogStr("Version : " + WasteLibrary.Version)
	currentUser = WasteLibrary.GetCurrentUser()
	WasteLibrary.LogStr(currentUser)
}
func main() {

	initStart()

	time.Sleep(time.Second)
	go systemCheck(WasteLibrary.RFID_APPNAME_SYSTEM)
	wg.Add(1)

	http.HandleFunc("/status", WasteLibrary.StatusHandler)
	http.ListenAndServe(":10006", nil)

	wg.Wait()

}

func systemCheck(appType string) {
	if currentUser == "pi" {
		for {

			time.Sleep(opInterval * time.Second)
		}
	}
	wg.Done()
}
