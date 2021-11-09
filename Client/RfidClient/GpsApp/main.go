package main

import (
	"bufio"
	"io"
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
var serialPort io.ReadWriteCloser

var serialOptions0 devafatekserial.OpenOptions = devafatekserial.OpenOptions{
	PortName:        "/dev/ttyAMA0",
	BaudRate:        9600,
	DataBits:        8,
	StopBits:        1,
	MinimumReadSize: 4,
}

var serialOptions1 devafatekserial.OpenOptions = devafatekserial.OpenOptions{
	PortName:        "/dev/ttyAMA1",
	BaudRate:        9600,
	DataBits:        8,
	StopBits:        1,
	MinimumReadSize: 4,
}

var currentDeviceType WasteLibrary.RfidDeviceType

func initStart() {

	time.Sleep(5 * time.Second)
	WasteLibrary.LogStr("Successfully connected!")
	WasteLibrary.Version = "1"
	WasteLibrary.LogStr("Version : " + WasteLibrary.Version)
	currentUser = WasteLibrary.GetCurrentUser()
	WasteLibrary.LogStr(currentUser)
	currentDeviceType.New()
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
	http.HandleFunc("/openLog", WasteLibrary.OpenLogHandler)
	http.HandleFunc("/closeLog", WasteLibrary.CloseLogHandler)
	http.ListenAndServe(":10003", nil)

	wg.Wait()
}

func gpsCheck() {
	if currentUser == "pi" {
		var err error
		for {
			time.Sleep(time.Second)
			WasteLibrary.LogStr("Device Check")
			serialPort, err = devafatekserial.Open(serialOptions0)
			if err != nil {
				WasteLibrary.LogErr(err)
				WasteLibrary.CurrentCheckStatu.ConnStatu = WasteLibrary.STATU_PASSIVE
			} else {
				WasteLibrary.CurrentCheckStatu.ConnStatu = WasteLibrary.STATU_ACTIVE
			}
			if WasteLibrary.CurrentCheckStatu.ConnStatu == WasteLibrary.STATU_PASSIVE {
				serialPort, err = devafatekserial.Open(serialOptions1)
				if err != nil {
					WasteLibrary.LogErr(err)
					WasteLibrary.CurrentCheckStatu.ConnStatu = WasteLibrary.STATU_PASSIVE
				} else {
					WasteLibrary.CurrentCheckStatu.ConnStatu = WasteLibrary.STATU_ACTIVE
				}
			}

			if WasteLibrary.CurrentCheckStatu.ConnStatu == WasteLibrary.STATU_ACTIVE {
				reader := bufio.NewReader(serialPort)
				scanner := bufio.NewScanner(reader)
				WasteLibrary.LogStr("Device OK")
				for scanner.Scan() {
					gps, err := devafatekgps.ParseGpsLine(scanner.Text())
					if err == nil {
						if gps.GetFixQuality() == "1" || gps.GetFixQuality() == "2" {
							latitude, _ := gps.GetLatitude()
							longitude, _ := gps.GetLongitude()

							if latitude != "" {
								currentDeviceType.DeviceGps.Latitude = WasteLibrary.StringToFloat64(latitude)
							} else {
								currentDeviceType.DeviceGps.Latitude = 0
							}
							if longitude != "" {
								currentDeviceType.DeviceGps.Longitude = WasteLibrary.StringToFloat64(longitude)
							} else {
								currentDeviceType.DeviceGps.Longitude = 0
							}
							WasteLibrary.CurrentCheckStatu.DeviceStatu = WasteLibrary.STATU_ACTIVE
						} else {
							currentDeviceType.DeviceGps.Latitude = 0
							currentDeviceType.DeviceGps.Longitude = 0
							WasteLibrary.CurrentCheckStatu.DeviceStatu = WasteLibrary.STATU_PASSIVE
						}
					}
					time.Sleep(opInterval * time.Second)
				}
			}
			WasteLibrary.LogStr("Device NONE")
		}
	}
	wg.Done()
}

func sendGps() {
	for {
		time.Sleep(opInterval * time.Second)
		data := url.Values{
			WasteLibrary.HTTP_READERTYPE: {WasteLibrary.READERTYPE_GPS},
			WasteLibrary.HTTP_DATA:       {currentDeviceType.ToString()},
		}
		WasteLibrary.HttpPostReq("http://127.0.0.1:10000/trans", data)
	}
	wg.Done()
}
