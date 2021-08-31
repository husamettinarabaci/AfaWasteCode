package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/user"
	"strconv"
	"sync"
	"time"

	"gitee.com/wiseai/go-rpio"
)

var debug bool = os.Getenv("DEBUG") == "1"
var wg sync.WaitGroup
var opInterval time.Duration = 5 * 60
var contactPort string = os.Getenv("CONTACT_PORT")
var currentUser string

var readerAppStatus string = "0"
var readerConnStatus string = "0"
var readerStatus string = "0"
var camAppStatus string = "0"
var camConnStatus string = "0"
var camStatus string = "0"
var gpsAppStatus string = "0"
var gpsConnStatus string = "0"
var gpsStatus string = "0"
var thermAppStatus string = "0"
var transferAppStatus string = "0"
var aliveStatus string = "1"
var contactStatus string = "0"

type statusType struct {
	Name string `json:"Name"`
	Key  string `json:"Key"`
	Port string `json:"Port"`
}

var statusTypes []statusType = []statusType{
	{
		Name: "readerAppStatus",
		Key:  "APP",
		Port: "10001",
	},
	{
		Name: "readerConnStatus",
		Key:  "CONN",
		Port: "10001",
	},
	{
		Name: "readerStatus",
		Key:  "READER",
		Port: "10001",
	},
	{
		Name: "camAppStatus",
		Key:  "APP",
		Port: "10002",
	},
	{
		Name: "camConnStatus",
		Key:  "CONN",
		Port: "10002",
	},
	{
		Name: "camStatus",
		Key:  "CAM",
		Port: "10002",
	},
	{
		Name: "gpsAppStatus",
		Key:  "APP",
		Port: "10003",
	},
	{
		Name: "gpsConnStatus",
		Key:  "CONN",
		Port: "10003",
	},
	{
		Name: "gpsStatus",
		Key:  "GPS",
		Port: "10003",
	},
	{
		Name: "thermAppStatus",
		Key:  "APP",
		Port: "10004",
	},
	{
		Name: "transferAppStatus",
		Key:  "APP",
		Port: "10000",
	},
	{
		Name: "aliveStatus",
		Key:  "NONE",
		Port: "",
	},
	{
		Name: "contactStatus",
		Key:  "NONE",
		Port: "",
	},
}

func initStart() {

	if !debug {
		time.Sleep(60 * time.Second)
	}
	logStr("Successfully connected!")
	currentUser = getCurrentUser()
	logStr(currentUser)
}
func main() {

	initStart()

	for i := range statusTypes {
		if statusTypes[i].Key == "NONE" {
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

	wg.Wait()
}

func statusCheck(statusTypeIndex int) {

	for {
		var lastStatus = "0"
		time.Sleep(opInterval * time.Second)
		data := url.Values{
			"OPTYPE": {statusTypes[statusTypeIndex].Key},
		}
		client := http.Client{
			Timeout: 10 * time.Second,
		}
		resp, err := client.PostForm("http://127.0.0.1:"+statusTypes[statusTypeIndex].Port+"/status", data)
		if err != nil {
			logErr(err)

		} else {
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				logErr(err)
			}
			bodyString := string(bodyBytes)
			if bodyString == "OK" {
				lastStatus = "1"
			}
		}

		if statusTypes[statusTypeIndex].Name == "readerAppStatus" {
			readerAppStatus = lastStatus
		} else if statusTypes[statusTypeIndex].Name == "readerConnStatus" {
			readerConnStatus = lastStatus
		} else if statusTypes[statusTypeIndex].Name == "readerStatus" {
			readerStatus = lastStatus
		} else if statusTypes[statusTypeIndex].Name == "camAppStatus" {
			camAppStatus = lastStatus
		} else if statusTypes[statusTypeIndex].Name == "camConnStatus" {
			camConnStatus = lastStatus
		} else if statusTypes[statusTypeIndex].Name == "camStatus" {
			camStatus = lastStatus
		} else if statusTypes[statusTypeIndex].Name == "gpsAppStatus" {
			gpsAppStatus = lastStatus
		} else if statusTypes[statusTypeIndex].Name == "gpsConnStatus" {
			gpsConnStatus = lastStatus
		} else if statusTypes[statusTypeIndex].Name == "gpsStatus" {
			gpsStatus = lastStatus
		} else if statusTypes[statusTypeIndex].Name == "transferAppStatus" {
			transferAppStatus = lastStatus
		} else if statusTypes[statusTypeIndex].Name == "aliveStatus" {
			aliveStatus = lastStatus
		} else if statusTypes[statusTypeIndex].Name == "contactStatus" {
			contactStatus = lastStatus
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
				contactStatus = "1"
			} else {
				contactStatus = "0"
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
			"OPTYPE":           {"STATUS"},
			"READERAPPSTATUS":  {readerAppStatus},
			"READERCONNSTATUS": {readerConnStatus},
			"READERSTATUS":     {readerStatus},
			"CAMAPPSTATUS":     {camAppStatus},
			"CAMCONNSTATUS":    {camConnStatus},
			"CAMSTATUS":        {camStatus},
			"GPSAPPSTATUS":     {gpsAppStatus},
			"GPSCONNSTATUS":    {gpsConnStatus},
			"GPSSTATUS":        {gpsStatus},
			"THERMAPSTATUS":    {thermAppStatus},
			"TRANSFERAPP":      {transferAppStatus},
			"ALIVESTATUS":      {aliveStatus},
			"CONTACTSTATUS":    {contactStatus},
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
