package main

import (
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/devafatek/WasteLibrary"
)

var wg sync.WaitGroup
var opInterval time.Duration = 5 * 60
var currentUser string
var currentDeviceType WasteLibrary.RecyDeviceType

type statusType struct {
	Name string `json:"Name"`
	Key  string `json:"Key"`
	Port string `json:"Port"`
}

var statusTypes []statusType = []statusType{
	{
		Name: "readerAppStatus",
		Key:  WasteLibrary.CHECKTYPE_APP,
		Port: "10001",
	},
	{
		Name: "readerConnStatus",
		Key:  WasteLibrary.CHECKTYPE_CONN,
		Port: "10001",
	},
	{
		Name: "readerStatus",
		Key:  WasteLibrary.CHECKTYPE_DEVICE,
		Port: "10001",
	},
	{
		Name: "camAppStatus",
		Key:  WasteLibrary.CHECKTYPE_APP,
		Port: "10002",
	},
	{
		Name: "camConnStatus",
		Key:  WasteLibrary.CHECKTYPE_CONN,
		Port: "10002",
	},
	{
		Name: "camStatus",
		Key:  WasteLibrary.CHECKTYPE_DEVICE,
		Port: "10002",
	},
	{
		Name: "webAppStatus",
		Key:  WasteLibrary.CHECKTYPE_APP,
		Port: "10003",
	},
	{
		Name: "thermAppStatus",
		Key:  WasteLibrary.CHECKTYPE_APP,
		Port: "10004",
	},
	{
		Name: "updaterAppStatus",
		Key:  WasteLibrary.CHECKTYPE_APP,
		Port: "10005",
	},
	{
		Name: "systemAppStatus",
		Key:  WasteLibrary.CHECKTYPE_APP,
		Port: "10006",
	},
	{
		Name: "motorAppStatus",
		Key:  WasteLibrary.CHECKTYPE_APP,
		Port: "10008",
	},
	{
		Name: "motorConnStatus",
		Key:  WasteLibrary.CHECKTYPE_CONN,
		Port: "10008",
	},
	{
		Name: "motorStatus",
		Key:  WasteLibrary.CHECKTYPE_DEVICE,
		Port: "10008",
	},
	{
		Name: "transferAppStatus",
		Key:  WasteLibrary.CHECKTYPE_APP,
		Port: "10000",
	},
	{
		Name: "aliveStatus",
		Key:  WasteLibrary.CHECKTYPE_NONE,
		Port: "",
	},
}

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

	for i := range statusTypes {
		if statusTypes[i].Key == WasteLibrary.CHECKTYPE_NONE {
			continue
		}
		go statusCheck(i)
		wg.Add(1)
	}

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
			WasteLibrary.HTTP_CHECKTYPE: {statusTypes[statusTypeIndex].Key},
		}

		resultVal = WasteLibrary.HttpPostReq("http://127.0.0.1:"+statusTypes[statusTypeIndex].Port+"/status", data)
		if resultVal.Result == WasteLibrary.RESULT_OK {
			lastStatus = WasteLibrary.STATU_ACTIVE
		}

		if statusTypes[statusTypeIndex].Name == "readerAppStatus" {
			currentDeviceType.DeviceStatu.ReaderAppStatus = lastStatus
		} else if statusTypes[statusTypeIndex].Name == "readerConnStatus" {
			currentDeviceType.DeviceStatu.ReaderConnStatus = lastStatus
		} else if statusTypes[statusTypeIndex].Name == "readerStatus" {
			currentDeviceType.DeviceStatu.ReaderStatus = lastStatus
		} else if statusTypes[statusTypeIndex].Name == "camAppStatus" {
			currentDeviceType.DeviceStatu.CamAppStatus = lastStatus
		} else if statusTypes[statusTypeIndex].Name == "camConnStatus" {
			currentDeviceType.DeviceStatu.CamConnStatus = lastStatus
		} else if statusTypes[statusTypeIndex].Name == "camStatus" {
			currentDeviceType.DeviceStatu.CamStatus = lastStatus
		} else if statusTypes[statusTypeIndex].Name == "webAppStatus" {
			currentDeviceType.DeviceStatu.WebAppStatus = lastStatus
		} else if statusTypes[statusTypeIndex].Name == "motorAppStatus" {
			currentDeviceType.DeviceStatu.MotorAppStatus = lastStatus
		} else if statusTypes[statusTypeIndex].Name == "motorConnStatus" {
			currentDeviceType.DeviceStatu.MotorConnStatus = lastStatus
		} else if statusTypes[statusTypeIndex].Name == "motorStatus" {
			currentDeviceType.DeviceStatu.MotorStatus = lastStatus
		} else if statusTypes[statusTypeIndex].Name == "thermAppStatus" {
			currentDeviceType.DeviceStatu.ThermAppStatus = lastStatus
		} else if statusTypes[statusTypeIndex].Name == "updaterAppStatus" {
			currentDeviceType.DeviceStatu.UpdaterAppStatus = lastStatus
		} else if statusTypes[statusTypeIndex].Name == "transferAppStatus" {
			currentDeviceType.DeviceStatu.TransferAppStatus = lastStatus
		} else if statusTypes[statusTypeIndex].Name == "systemAppStatus" {
			currentDeviceType.DeviceStatu.SystemAppStatus = lastStatus
		} else if statusTypes[statusTypeIndex].Name == "aliveStatus" {
			currentDeviceType.DeviceStatu.AliveStatus = lastStatus
		} else {

		}
	}
	wg.Done()
}

func sendStatus() {
	for {
		time.Sleep(opInterval * time.Second)

		data := url.Values{
			WasteLibrary.HTTP_READERTYPE: {WasteLibrary.READERTYPE_STATUS},
			WasteLibrary.HTTP_DATA:       {currentDeviceType.ToString()},
		}
		WasteLibrary.HttpPostReq("http://127.0.0.1:10000/trans", data)
	}
	wg.Done()
}
