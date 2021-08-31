package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/user"
	"sync"
	"time"

	"github.com/AfatekDevelopers/gps_lib_go/devafatekgps"
	"github.com/AfatekDevelopers/serial_lib_go/devafatekserial"
)

var debug bool = os.Getenv("DEBUG") == "1"
var appStatus string = "1"
var connStatus string = "0"
var gpsStatus string = "0"
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

type gpsType struct {
	Latitude  string `json:"Latitude"`
	Longitude string `json:"Longitude"`
}

var currentGpsType gpsType

func initStart() {

	currentGpsType.Latitude = ""
	currentGpsType.Longitude = ""

	if !debug {
		time.Sleep(60 * time.Second)
	}
	logStr("Successfully connected!")
	currentUser = getCurrentUser()
	logStr(currentUser)
}
func main() {

	initStart()

	time.Sleep(5 * time.Second)
	go gpsCheck()
	wg.Add(1)
	time.Sleep(5 * time.Second)
	go sendGps()
	wg.Add(1)

	http.HandleFunc("/status", status)
	http.ListenAndServe(":10003", nil)

	wg.Wait()
}

func status(w http.ResponseWriter, req *http.Request) {

	if err := req.ParseForm(); err != nil {
		logErr(err)
		return
	}
	opType := req.FormValue("OPTYPE")
	logStr(opType)

	if opType == "APP" {
		if appStatus == "1" {
			w.Write([]byte("OK"))
		} else {
			w.Write([]byte("FAIL"))
		}
	} else if opType == "CONN" {
		if connStatus == "1" {
			w.Write([]byte("OK"))
		} else {
			w.Write([]byte("FAIL"))
		}
	} else if opType == "GPS" {
		if gpsStatus == "1" {
			w.Write([]byte("OK"))
		} else {
			w.Write([]byte("FAIL"))
		}
	} else {
		w.Write([]byte("FAIL"))
	}
}

func gpsCheck() {
	if currentUser == "pi" {

		serialPort, err := devafatekserial.Open(serialOptions)
		if err != nil {
			logErr(err)
			connStatus = "0"
		} else {
			connStatus = "1"
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

					currentGpsType.Latitude = latitude
					currentGpsType.Longitude = longitude
					gpsStatus = "1"
				} else {
					currentGpsType.Latitude = ""
					currentGpsType.Longitude = ""
					gpsStatus = "0"
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
			"OPTYPE":    {"GPS"},
			"LATITUDE":  {string(currentGpsType.Latitude)},
			"LONGITUDE": {string(currentGpsType.Longitude)},
		}
		client := http.Client{
			Timeout: 10 * time.Second,
		}
		resp, err := client.PostForm("http://127.0.0.1:10000/trans", data)
		if err != nil {
			logErr(err)

		} else {
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				logErr(err)
			}
			bodyString := string(bodyBytes)
			logStr(bodyString)
		}
	}
	wg.Done()
}

func getCurrentUser() string {
	user, err := user.Current()
	if err != nil {
		logErr(err)
	}

	username := user.Username
	return username
}

func logErr(err error) {
	if err != nil {
		logStr(err.Error())
	}
}

func logStr(value string) {
	if debug {
		fmt.Println(value)
	}
}
