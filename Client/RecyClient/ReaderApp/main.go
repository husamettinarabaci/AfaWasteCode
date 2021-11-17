package main

import (
	"encoding/hex"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/devafatek/WasteLibrary"
)

var wg sync.WaitGroup
var currentUser string

var opInterval time.Duration = 5 * 60
var lastReadTime time.Time
var lastSendTime time.Time
var lastRfNfc string = ""
var readNfcs map[string]time.Time
var currentNfcDataType WasteLibrary.NfcType

func initStart() {

	lastReadTime = time.Now()
	lastSendTime = time.Now()
	readNfcs = make(map[string]time.Time)
	time.Sleep(5 * time.Second)
	WasteLibrary.LogStr("Successfully connected!")
	WasteLibrary.Version = "1"
	WasteLibrary.LogStr("Version : " + WasteLibrary.Version)
	currentUser = WasteLibrary.GetCurrentUser()
	WasteLibrary.LogStr(currentUser)
	currentNfcDataType.New()
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
	http.HandleFunc("/trigger", trigger)
	http.ListenAndServe(":10001", nil)

	wg.Wait()
}

func trigger(w http.ResponseWriter, req *http.Request) {

	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL

	if err := req.ParseForm(); err != nil {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_HTTP_PARSE
		w.Write(resultVal.ToByte())

		WasteLibrary.LogErr(err)
		return
	}

	cardId := req.URL.Query()["cardId"]
	WasteLibrary.LogStr(cardId[0])
	w.Write(resultVal.ToByte())

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

						if strings.Contains(data, WasteLibrary.RECY_READER_OKBIT) || strings.Contains(data, WasteLibrary.RECY_READER_STARTBIT) {
							WasteLibrary.CurrentCheckStatu.DeviceStatu = WasteLibrary.STATU_ACTIVE
						} else {
							WasteLibrary.CurrentCheckStatu.DeviceStatu = WasteLibrary.STATU_PASSIVE
						}

						tempData += data

						if len(tempData) == 64 && tempData[:4] == WasteLibrary.RECY_READER_STARTBIT && tempData[10:12] == WasteLibrary.RECY_READER_CHECKBIT && tempData[36:50] == WasteLibrary.RECY_NFC_PATTERN {
							if time.Since(readNfcs[tempData[36:60]]).Seconds() > 15*60 {
								lastRfNfc = tempData[36:60]
								nid, _ := uuid.NewUUID()
								lastSendTime = time.Now()
								readNfcs[tempData[36:60]] = lastSendTime
								currentNfcDataType.NfcMain.Epc = lastRfNfc
								currentNfcDataType.NfcReader.UID = nid.String()
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
		WasteLibrary.HTTP_READERTYPE: {WasteLibrary.READERTYPE_RF},
		WasteLibrary.HTTP_DATA:       {currentNfcDataType.ToString()},
	}
	WasteLibrary.HttpPostReq("http://127.0.0.1:10000/trans", data)
}

func sendRfToCam() {

	data := url.Values{
		WasteLibrary.HTTP_READERTYPE: {WasteLibrary.READERTYPE_CAMTRIGGER},
		WasteLibrary.HTTP_DATA:       {currentNfcDataType.ToString()},
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
