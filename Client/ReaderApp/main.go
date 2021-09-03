package main

import (
	"encoding/hex"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"gitee.com/wiseai/go-rpio"
	"github.com/AfatekDevelopers/serial_lib_go/devafatekserial"
	"github.com/devafatek/WasteLibrary"
	"github.com/google/uuid"
)

var readerPort string = os.Getenv("READER_PORT")
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

var currentReaderDataType WasteLibrary.ReaderDataType = WasteLibrary.ReaderDataType{
	UID:   "",
	TagID: "",
}

func initStart() {

	lastReadTime = time.Now()
	lastSendTime = time.Now()
	readTags = make(map[string]time.Time)
	time.Sleep(5 * time.Second)
	WasteLibrary.LogStr("Successfully connected!")
	currentUser = WasteLibrary.GetCurrentUser()
	WasteLibrary.LogStr(currentUser)
}
func main() {

	initStart()

	time.Sleep(5 * time.Second)
	go rfCheck()
	wg.Add(1)

	time.Sleep(5 * time.Second)
	go deviceCheck()
	wg.Add(1)

	http.HandleFunc("/status", WasteLibrary.StatusHandler)
	http.ListenAndServe(":10001", nil)

	wg.Wait()
}

func rfCheck() {
	if currentUser == "pi" {

		serialPort, err := devafatekserial.Open(serialOptions)
		if err != nil {
			WasteLibrary.LogErr(err)
			WasteLibrary.CurrentCheckStatu.ConnStatu = "0"
		} else {
			WasteLibrary.CurrentCheckStatu.ConnStatu = "1"
		}
		defer serialPort.Close()
		var data string = ""
		var tempData string = ""
		for {
			buf := make([]byte, 256)
			n, err := serialPort.Read(buf)
			if err != nil {
				if err != io.EOF {
					WasteLibrary.LogErr(err)
				}
			} else {
				buf = buf[:n]
				data = hex.EncodeToString(buf)
				lastReadTime = time.Now()
				data = strings.ToUpper(data)
				WasteLibrary.LogStr(data)

				if strings.Contains(data, "5379") || strings.Contains(data, "4354") {
					WasteLibrary.CurrentCheckStatu.DeviceStatu = "1"
				} else {
					WasteLibrary.CurrentCheckStatu.DeviceStatu = "0"
				}

				tempData += data

				if len(tempData) == 64 && tempData[:4] == "4354" && tempData[10:12] == "45" && tempData[36:50] == "AFA09012018AFA" {
					if time.Since(readTags[tempData[36:60]]).Seconds() > 15*60 {
						lastRfTag = tempData[36:60]
						nid, _ := uuid.NewUUID()
						lastSendTime = time.Now()
						readTags[tempData[36:60]] = lastSendTime
						currentReaderDataType.TagID = lastRfTag
						currentReaderDataType.UID = nid.String()
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
		"DATA":   {currentReaderDataType.ToString()},
	}
	WasteLibrary.HttpPostReq("http://127.0.0.1:10000/trans", data)
}

func sendRfToCam() {

	data := url.Values{
		"OPTYPE": {"RF"},
		"DATA":   {currentReaderDataType.ToString()},
	}
	WasteLibrary.HttpPostReq("http://127.0.0.1:10002/trigger", data)
}

func deviceCheck() {
	for {
		if time.Since(lastReadTime).Seconds() > 60*60 {

			WasteLibrary.LogStr("Restart device...")
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
