package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/user"
	"strconv"
	"strings"
	"sync"
	"time"

	"gitee.com/wiseai/go-rpio"
	"github.com/AfatekDevelopers/serial_lib_go/devafatekserial"
	"github.com/google/uuid"
)

var debug bool = os.Getenv("DEBUG") == "1"
var readerPort string = os.Getenv("READER_PORT")
var appStatus string = "1"
var connStatus string = "0"
var readerStatus string = "0"
var opInterval time.Duration = 5 * 60
var wg sync.WaitGroup
var currentUser string

var lastReadTime time.Time
var lastSendTime time.Time
var lastRfTag string = ""
var readTags map[string]time.Time

var serialOptions devafatekserial.OpenOptions = devafatekserial.OpenOptions{
	PortName:        "/dev/ttyUSB0",
	BaudRate:        115200,
	DataBits:        8,
	StopBits:        1,
	MinimumReadSize: 4,
}

type rfType struct {
	TagID string `json:"TagID"`
	UID   string `json:"UID"`
}

var currentRfType rfType

func initStart() {

	currentRfType.TagID = ""
	currentRfType.UID = ""
	lastReadTime = time.Now()
	lastSendTime = time.Now()
	readTags = make(map[string]time.Time)
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
	go rfCheck()
	wg.Add(1)

	time.Sleep(5 * time.Second)
	go deviceCheck()
	wg.Add(1)

	http.HandleFunc("/status", status)
	http.ListenAndServe(":10001", nil)

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
	} else if opType == "READER" {
		if readerStatus == "1" {
			w.Write([]byte("OK"))
		} else {
			w.Write([]byte("FAIL"))
		}
	} else {
		w.Write([]byte("FAIL"))
	}
}

func rfCheck() {
	if currentUser == "pi" {

		serialPort, err := devafatekserial.Open(serialOptions)
		if err != nil {
			logErr(err)
			connStatus = "0"
		} else {
			connStatus = "1"
		}
		defer serialPort.Close()
		var data string = ""
		var tempData string = ""
		for {
			buf := make([]byte, 256)
			n, err := serialPort.Read(buf)
			if err != nil {
				if err != io.EOF {
					logErr(err)
				}
			} else {
				buf = buf[:n]
				data = hex.EncodeToString(buf)
				lastReadTime = time.Now()
				data = strings.ToUpper(data)
				logStr(data)

				if strings.Contains(data, "5379") || strings.Contains(data, "4354") {
					readerStatus = "1"
				} else {
					readerStatus = "0"
				}

				tempData += data

				if len(tempData) == 64 && tempData[:4] == "4354" && tempData[10:12] == "45" && tempData[36:50] == "AFA09012018AFA" {
					if time.Since(readTags[tempData[36:60]]).Seconds() > 15*60 {
						lastRfTag = tempData[36:60]
						nid, _ := uuid.NewUUID()
						lastSendTime = time.Now()
						readTags[tempData[36:60]] = lastSendTime
						currentRfType.TagID = lastRfTag
						currentRfType.UID = nid.String()
						sendRf()
						sendRfToCam()
					}
				}
				if len(tempData) >= 64 {
					tempData = ""
				}

			}
		}
	}
	wg.Done()
}

func sendRf() {

	data := url.Values{
		"OPTYPE": {"RF"},
		"TAGID":  {string(currentRfType.TagID)},
		"UID":    {string(currentRfType.UID)},
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

func sendRfToCam() {

	currentRfTypeJSON, err := json.Marshal(currentRfType)
	if err != nil {
		logErr(err)

	}
	data := url.Values{
		"OPTYPE": {"RF"},
		"RF":     {string(currentRfTypeJSON)},
	}
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.PostForm("http://127.0.0.1:10002/trigger", data)
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

func deviceCheck() {
	for {
		if time.Since(lastReadTime).Seconds() > 60*60 {

			logStr("Restart device...")
			rpio.Open()
			readerPort, _ := strconv.Atoi(readerPort)
			pin := rpio.Pin(readerPort)
			pin.Output()
			pin.High()
			time.Sleep(10 * time.Second)
			pin.Low()
			rpio.Close()
		}

		time.Sleep(opInterval * time.Second)
	}
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
