package main

import (
	"bytes"
	"net/http"
	"net/url"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/devafatek/WasteLibrary"
)

var opInterval time.Duration = 5 * 60
var wg sync.WaitGroup
var currentUser string
var currentThermDataType WasteLibrary.RecyDeviceType

func initStart() {
	time.Sleep(5 * time.Second)
	WasteLibrary.LogStr("Successfully connected!")
	WasteLibrary.Version = "1"
	WasteLibrary.LogStr("Version : " + WasteLibrary.Version)
	currentUser = WasteLibrary.GetCurrentUser()
	WasteLibrary.LogStr(currentUser)
	currentThermDataType.New()
}
func main() {

	initStart()

	time.Sleep(5 * time.Second)
	go thermCheck()
	wg.Add(1)
	time.Sleep(5 * time.Second)
	go sendTherm()
	wg.Add(1)

	http.HandleFunc("/status", WasteLibrary.StatusHandler)
	http.ListenAndServe(":10004", nil)

	wg.Wait()
}

func thermCheck() {
	if currentUser == "pi" {
		for {
			time.Sleep(opInterval * time.Second)
			cmd := exec.Command("vcgencmd", "measure_temp")
			var outb, errb bytes.Buffer
			cmd.Stdout = &outb
			cmd.Stderr = &errb
			err := cmd.Run()
			currentThermDataType.DeviceTherm.Therm = strings.TrimSuffix(outb.String(), "'C\n")
			WasteLibrary.LogStr(currentThermDataType.DeviceTherm.Therm)
			if err != nil {
				WasteLibrary.LogErr(err)

			}
		}
	}
	wg.Done()
}

func sendTherm() {

	for {
		time.Sleep(opInterval * time.Second)
		data := url.Values{
			WasteLibrary.HTTP_READERTYPE: {WasteLibrary.READERTYPE_THERM},
			WasteLibrary.HTTP_DATA:       {currentThermDataType.ToString()},
		}
		WasteLibrary.HttpPostReq("http://127.0.0.1:10000/trans", data)
	}
	wg.Done()
}
