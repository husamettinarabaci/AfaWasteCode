package main

import (
	"bufio"
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/AfatekDevelopers/gps_lib_go/devafatekgps"
	"github.com/AfatekDevelopers/serial_lib_go/devafatekserial"
	"github.com/devafatek/WasteLibrary"
)

var opInterval time.Duration = 1 * 60
var wg sync.WaitGroup
var currentUser string

var serialOptions devafatekserial.OpenOptions = devafatekserial.OpenOptions{
	PortName:        "/dev/ttyAMA0",
	BaudRate:        9600,
	DataBits:        8,
	StopBits:        1,
	MinimumReadSize: 4,
}

var currentGpsDataType WasteLibrary.GpsDataType = WasteLibrary.GpsDataType{
	Latitude:   "",
	Longitude:  "",
	LatitudeF:  0,
	LongitudeF: 0,
}

func initStart() {

	time.Sleep(5 * time.Second)
	WasteLibrary.LogStr("Successfully connected!")
	currentUser = WasteLibrary.GetCurrentUser()
	WasteLibrary.LogStr(currentUser)
}
func main() {

	initStart()

	time.Sleep(5 * time.Second)
	go gpsCheck()
	wg.Add(1)
	time.Sleep(5 * time.Second)
	go sendGps()
	wg.Add(1)

	http.HandleFunc("/status", WasteLibrary.StatusHandler)
	http.ListenAndServe(":10003", nil)

	wg.Wait()
}

func gpsCheck() {
	if currentUser == "pi" {

		serialPort, err := devafatekserial.Open(serialOptions)
		if err != nil {
			WasteLibrary.LogErr(err)
			WasteLibrary.CurrentCheckStatu.ConnStatu = "0"
		} else {
			WasteLibrary.CurrentCheckStatu.ConnStatu = "1"
		}
		defer serialPort.Close()
		reader := bufio.NewReader(serialPort)
		scanner := bufio.NewScanner(reader)

		for scanner.Scan() {
			gps, err := devafatekgps.ParseGpsLine(scanner.Text())
			fmt.Println(gps)
			if err == nil {
				if gps.GetFixQuality() == "1" || gps.GetFixQuality() == "2" {
					latitude, _ := gps.GetLatitude()
					longitude, _ := gps.GetLongitude()

					currentGpsDataType.Latitude = latitude
					currentGpsDataType.Longitude = longitude
					WasteLibrary.CurrentCheckStatu.DeviceStatu = "1"
				} else {
					currentGpsDataType.Latitude = ""
					currentGpsDataType.Longitude = ""
					WasteLibrary.CurrentCheckStatu.DeviceStatu = "0"
				}
				time.Sleep(opInterval * time.Second)
			}
		}
	}
	wg.Done()
}

func sendGps() {
	for {
		time.Sleep(opInterval * time.Second)
		data := url.Values{
			"OPTYPE": {"GPS"},
			"DATA":   {currentGpsDataType.ToString()},
		}
		WasteLibrary.HttpPostReq("http://127.0.0.1:10000/trans", data)
	}
	wg.Done()
}
