package main

import (
	"net/http"
	"net/url"
	"os"
	"strconv"
	"sync"
	"time"

	"gitee.com/wiseai/go-rpio"
	"github.com/devafatek/WasteLibrary"
)

var wg sync.WaitGroup
var opInterval time.Duration = 5 * 60
var contactPort string = os.Getenv("CONTACT_PORT")
var currentUser string
var currentDeviceType WasteLibrary.DeviceType
var version = "1"

type statusType struct {
	Name string `json:"Name"`
	Key  string `json:"Key"`
	Port string `json:"Port"`
}

var statusTypes []statusType = []statusType{
	{
		Name: "readerAppStatus",
		Key:  WasteLibrary.OPTYPE_APP,
		Port: "10001",
	},
	{
		Name: "readerConnStatus",
		Key:  WasteLibrary.OPTYPE_CONN,
		Port: "10001",
	},
	{
		Name: "readerStatus",
		Key:  WasteLibrary.OPTYPE_READER,
		Port: "10001",
	},
	{
		Name: "camAppStatus",
		Key:  WasteLibrary.OPTYPE_APP,
		Port: "10002",
	},
	{
		Name: "camConnStatus",
		Key:  WasteLibrary.OPTYPE_CONN,
		Port: "10002",
	},
	{
		Name: "camStatus",
		Key:  WasteLibrary.OPTYPE_CAM,
		Port: "10002",
	},
	{
		Name: "gpsAppStatus",
		Key:  WasteLibrary.OPTYPE_APP,
		Port: "10003",
	},
	{
		Name: "gpsConnStatus",
		Key:  WasteLibrary.OPTYPE_CONN,
		Port: "10003",
	},
	{
		Name: "gpsStatus",
		Key:  WasteLibrary.OPTYPE_GPS,
		Port: "10003",
	},
	{
		Name: "thermAppStatus",
		Key:  WasteLibrary.OPTYPE_APP,
		Port: "10004",
	},
	{
		Name: "updaterAppStatus",
		Key:  WasteLibrary.OPTYPE_APP,
		Port: "10005",
	},
	{
		Name: "systemAppStatus",
		Key:  WasteLibrary.OPTYPE_APP,
		Port: "10006",
	},
	{
		Name: "transferAppStatus",
		Key:  WasteLibrary.OPTYPE_APP,
		Port: "10000",
	},
	{
		Name: "aliveStatus",
		Key:  WasteLibrary.OPTYPE_NONE,
		Port: "",
	},
	{
		Name: "contactStatus",
		Key:  WasteLibrary.OPTYPE_NONE,
		Port: "",
	},
}

func initStart() {

	time.Sleep(5 * time.Second)
	WasteLibrary.LogStr("Successfully connected!")
	WasteLibrary.LogStr("Version : " + version)
	currentUser = WasteLibrary.GetCurrentUser()
	WasteLibrary.LogStr(currentUser)
	currentDeviceType.New()
}
func main() {

	initStart()

	for i := range statusTypes {
		if statusTypes[i].Key == WasteLibrary.OPTYPE_NONE {
			continue
		}
		go statusCheck(i)
		wg.Add(1)
	}

	time.Sleep(5 * time.Second)
	go contactCheck()
	wg.Add(1)
	time.Sleep(5 * time.Second)
	go sendStatus()
	wg.Add(1)

	http.HandleFunc("/status", WasteLibrary.StatusHandler)
	http.ListenAndServe(":10007", nil)

	wg.Wait()
}

func statusCheck(statusTypeIndex int) {
	var resultVal WasteLibrary.ResultType
	for {
		var lastStatus = WasteLibrary.STATU_PASSIVE
		time.Sleep(opInterval * time.Second)
		data := url.Values{
			WasteLibrary.HTTP_OPTYPE: {statusTypes[statusTypeIndex].Key},
		}

		resultVal = WasteLibrary.HttpPostReq("http://127.0.0.1:"+statusTypes[statusTypeIndex].Port+"/status", data)
		if resultVal.Result == WasteLibrary.RESULT_OK {
			lastStatus = WasteLibrary.STATU_ACTIVE
		}

		if statusTypes[statusTypeIndex].Name == "readerAppStatus" {
			currentDeviceType.ReaderAppStatus = lastStatus
		} else if statusTypes[statusTypeIndex].Name == "readerConnStatus" {
			currentDeviceType.ReaderConnStatus = lastStatus
		} else if statusTypes[statusTypeIndex].Name == "readerStatus" {
			currentDeviceType.ReaderStatus = lastStatus
		} else if statusTypes[statusTypeIndex].Name == "camAppStatus" {
			currentDeviceType.CamAppStatus = lastStatus
		} else if statusTypes[statusTypeIndex].Name == "camConnStatus" {
			currentDeviceType.CamConnStatus = lastStatus
		} else if statusTypes[statusTypeIndex].Name == "camStatus" {
			currentDeviceType.CamStatus = lastStatus
		} else if statusTypes[statusTypeIndex].Name == "gpsAppStatus" {
			currentDeviceType.GpsAppStatus = lastStatus
		} else if statusTypes[statusTypeIndex].Name == "gpsConnStatus" {
			currentDeviceType.GpsConnStatus = lastStatus
		} else if statusTypes[statusTypeIndex].Name == "gpsStatus" {
			currentDeviceType.GpsStatus = lastStatus
		} else if statusTypes[statusTypeIndex].Name == "thermAppStatus" {
			currentDeviceType.ThermAppStatus = lastStatus
		} else if statusTypes[statusTypeIndex].Name == "updaterAppStatus" {
			currentDeviceType.UpdaterAppStatus = lastStatus
		} else if statusTypes[statusTypeIndex].Name == "transferAppStatus" {
			currentDeviceType.TransferAppStatus = lastStatus
		} else if statusTypes[statusTypeIndex].Name == "systemAppStatus" {
			currentDeviceType.SystemAppStatus = lastStatus
		} else if statusTypes[statusTypeIndex].Name == "aliveStatus" {
			currentDeviceType.AliveStatus = lastStatus
		} else if statusTypes[statusTypeIndex].Name == "contactStatus" {
			currentDeviceType.ContactStatus = lastStatus
		} else {

		}
	}
	wg.Done()
}

func contactCheck() {

	if currentUser == "pi" {
		for {
			time.Sleep(opInterval * time.Second)
			rpio.Open()
			conPort, _ := strconv.Atoi(contactPort)
			pin := rpio.Pin(conPort)
			var tempData = rpio.ReadPin(pin) == 1
			if tempData {
				currentDeviceType.ContactStatus = WasteLibrary.STATU_ACTIVE
			} else {
				currentDeviceType.ContactStatus = WasteLibrary.STATU_PASSIVE
			}
			rpio.Close()
		}
	}
	wg.Done()
}

func sendStatus() {
	for {
		time.Sleep(opInterval * time.Second)

		data := url.Values{
			WasteLibrary.HTTP_OPTYPE: {WasteLibrary.OPTYPE_STATUS},
			WasteLibrary.HTTP_DATA:   {currentDeviceType.ToString()},
		}
		WasteLibrary.HttpPostReq("http://127.0.0.1:10000/trans", data)
	}
	wg.Done()
}
