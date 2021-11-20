package main

import (
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/devafatek/WasteLibrary"
	"github.com/google/uuid"
)

var wg sync.WaitGroup
var currentUser string

var opInterval time.Duration = 5 * 60
var lastSendTime time.Time
var readNfcs map[string]time.Time
var currentNfcDataType WasteLibrary.NfcType

func initStart() {

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

	w.Write(resultVal.ToByte())

	cardId := req.URL.Query()["cardId"]
	if time.Since(readNfcs[cardId[0]]).Seconds() > 30 {
		lastSendTime = time.Now()
		readNfcs[cardId[0]] = lastSendTime
		nid, _ := uuid.NewUUID()
		currentNfcDataType.NfcMain.Epc = cardId[0]
		currentNfcDataType.NfcReader.UID = nid.String()
		WasteLibrary.LogStr(currentNfcDataType.ToString())
		go sendRf()
		go sendRfToWeb()
		go sendRfToCam()
	}

}

func sendRf() {
	data := url.Values{
		WasteLibrary.HTTP_READERTYPE: {WasteLibrary.READERTYPE_RF},
		WasteLibrary.HTTP_DATA:       {currentNfcDataType.ToString()},
	}
	WasteLibrary.HttpPostReq("http://127.0.0.1:10000/trans", data)
}

func sendRfToWeb() {

	data := url.Values{
		WasteLibrary.HTTP_READERTYPE: {WasteLibrary.READERTYPE_RF},
		WasteLibrary.HTTP_DATA:       {currentNfcDataType.ToString()},
	}
	WasteLibrary.HttpPostReq("http://127.0.0.1:10003/trigger", data)
}

func sendRfToCam() {

	data := url.Values{
		WasteLibrary.HTTP_READERTYPE: {WasteLibrary.READERTYPE_CAMTRIGGER},
		WasteLibrary.HTTP_DATA:       {currentNfcDataType.ToString()},
	}
	WasteLibrary.HttpPostReq("http://127.0.0.1:10002/trigger", data)
}

func deviceCheck() {
	//TO DO
	//check /dev/hidraw
}
