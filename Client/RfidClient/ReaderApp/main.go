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
var serialPort io.ReadWriteCloser
var version = "1"

var serialOptions0 devafatekserial.OpenOptions = devafatekserial.OpenOptions{
	PortName:        "/dev/ttyUSB0",
	BaudRate:        115200,
	DataBits:        8,
	StopBits:        1,
	MinimumReadSize: 4,
}

var serialOptions1 devafatekserial.OpenOptions = devafatekserial.OpenOptions{
	PortName:        "/dev/ttyUSB1",
	BaudRate:        115200,
	DataBits:        8,
	StopBits:        1,
	MinimumReadSize: 4,
}

var currentTagDataType WasteLibrary.TagType

func initStart() {

	lastReadTime = time.Now()
	lastSendTime = time.Now()
	readTags = make(map[string]time.Time)
	time.Sleep(5 * time.Second)
	WasteLibrary.LogStr("Successfully connected!")
	WasteLibrary.LogStr("Version : " + version)
	currentUser = WasteLibrary.GetCurrentUser()
	WasteLibrary.LogStr(currentUser)
	currentTagDataType.New()
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

			var data string = ""
			var tempData string = ""

			if WasteLibrary.CurrentCheckStatu.ConnStatu == WasteLibrary.STATU_ACTIVE {
				for {
					WasteLibrary.LogStr("Device OK")
					buf := make([]byte, 256)
					n, err := serialPort.Read(buf)
					if err != nil {
						WasteLibrary.LogErr(err)
						if err != io.EOF {
							WasteLibrary.LogErr(err)
						}
						break
					} else {
						buf = buf[:n]
						data = hex.EncodeToString(buf)
						lastReadTime = time.Now()
						data = strings.ToUpper(data)
						WasteLibrary.LogStr(data)

						if strings.Contains(data, WasteLibrary.RFID_READER_OKBIT) || strings.Contains(data, WasteLibrary.RFID_READER_STARTBIT) {
							WasteLibrary.CurrentCheckStatu.DeviceStatu = WasteLibrary.STATU_ACTIVE
						} else {
							WasteLibrary.CurrentCheckStatu.DeviceStatu = WasteLibrary.STATU_PASSIVE
						}

						tempData += data

						if len(tempData) == 64 && tempData[:4] == WasteLibrary.RFID_READER_STARTBIT && tempData[10:12] == WasteLibrary.RFID_READER_CHECKBIT && tempData[36:50] == WasteLibrary.RFID_TAG_PATTERN {
							if time.Since(readTags[tempData[36:60]]).Seconds() > 15*60 {
								lastRfTag = tempData[36:60]
								nid, _ := uuid.NewUUID()
								lastSendTime = time.Now()
								readTags[tempData[36:60]] = lastSendTime
								currentTagDataType.Epc = lastRfTag
								currentTagDataType.UID = nid.String()
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
			WasteLibrary.LogStr("Device NONE")
		}
	}
	wg.Done()
}

func sendRf() {
	data := url.Values{
		WasteLibrary.HTTP_OPTYPE: {WasteLibrary.OPTYPE_RF},
		WasteLibrary.HTTP_DATA:   {currentTagDataType.ToString()},
	}
	WasteLibrary.HttpPostReq("http://127.0.0.1:10000/trans", data)
}

func sendRfToCam() {

	data := url.Values{
		WasteLibrary.HTTP_OPTYPE: {WasteLibrary.OPTYPE_RF},
		WasteLibrary.HTTP_DATA:   {currentTagDataType.ToString()},
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
