package main

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/devafatek/WasteLibrary"
)

var opInterval time.Duration = 5 * 60
var wg sync.WaitGroup
var integratedPortInt = 1
var currentUser string

func initStart() {
	time.Sleep(5 * time.Second)

	WasteLibrary.LogStr("Successfully connected!")
	WasteLibrary.Version = "1"
	WasteLibrary.LogStr("Version : " + WasteLibrary.Version)
	currentUser = WasteLibrary.GetCurrentUser()
	WasteLibrary.LogStr(currentUser)
}

func main() {

	initStart()

	time.Sleep(5 * time.Second)
	go camCheck()
	wg.Add(1)

	http.HandleFunc("/status", WasteLibrary.StatusHandler)
	http.HandleFunc("/trigger", trigger)
	http.ListenAndServe(":10002", nil)
}

func trigger(w http.ResponseWriter, req *http.Request) {

	if WasteLibrary.AllowCors {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,access-control-allow-origin, access-control-allow-headers")
	}
	var resultVal WasteLibrary.ResultType

	if err := req.ParseForm(); err != nil {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_HTTP_PARSE
		w.Write(resultVal.ToByte())

		WasteLibrary.LogErr(err)
		return
	}

	readerType := req.FormValue(WasteLibrary.HTTP_READERTYPE)

	resultVal.Result = WasteLibrary.RESULT_FAIL
	if readerType == WasteLibrary.READERTYPE_CAMTRIGGER {
		var readerDataTypeVal WasteLibrary.NfcType
		readerDataTypeVal.StringToType(req.FormValue(WasteLibrary.HTTP_DATA))
		if integratedPortInt == 4 {
			integratedPortInt = 1
		}
		doRecord(readerDataTypeVal, strconv.Itoa(integratedPortInt), WasteLibrary.STATU_ACTIVE)
		integratedPortInt++
		resultVal.Result = WasteLibrary.RESULT_OK
		resultVal.Retval = WasteLibrary.RESULT_SUCCESS_OK
	} else {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_READERTYPE
		w.Write(resultVal.ToByte())

		return
	}
	w.Write(resultVal.ToByte())

}

func doRecord(readerDataTypeVal WasteLibrary.NfcType, integratedPort string, repeat string) {
	WasteLibrary.CurrentCheckStatu.DeviceStatu = WasteLibrary.STATU_PASSIVE

	cmd := exec.Command("timeout", "30", "ffmpeg", "-y", "-v", "0", "-loglevel", "0", "-hide_banner", "-f", "mpegts", "-i", "udp://localhost:1000"+integratedPort, "-frames:v", "1", "WAIT_CAM/"+readerDataTypeVal.NfcReader.UID+".png")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := cmd.Run()

	if err != nil && !strings.Contains(err.Error(), "124") {
		WasteLibrary.LogErr(err)
		if repeat == WasteLibrary.STATU_ACTIVE {
			doRecord(readerDataTypeVal, integratedPort, WasteLibrary.STATU_PASSIVE)
			return
		}
	} else {
		time.Sleep(2 * time.Second)

		if WasteLibrary.IsFileExists("WAIT_CAM/" + readerDataTypeVal.NfcReader.UID + ".png") {
			_, err := os.Stat("WAIT_CAM/" + readerDataTypeVal.NfcReader.UID + ".png")
			if err != nil {
				if repeat == WasteLibrary.STATU_ACTIVE {
					doRecord(readerDataTypeVal, integratedPort, WasteLibrary.STATU_PASSIVE)
					return
				}
			}

			WasteLibrary.CurrentCheckStatu.DeviceStatu = WasteLibrary.STATU_ACTIVE
			sendCam(readerDataTypeVal)
			sendCamToWeb(readerDataTypeVal)

		} else {
			if repeat == WasteLibrary.STATU_ACTIVE {
				doRecord(readerDataTypeVal, integratedPort, WasteLibrary.STATU_PASSIVE)
				return
			}
		}
	}
}

func sendCam(readerDataTypeVal WasteLibrary.NfcType) {

	data := url.Values{
		WasteLibrary.HTTP_READERTYPE: {WasteLibrary.READERTYPE_CAM},
		WasteLibrary.HTTP_DATA:       {readerDataTypeVal.ToString()},
	}

	WasteLibrary.HttpPostReq("http://127.0.0.1:10000/trans", data)
}

func sendCamToWeb(readerDataTypeVal WasteLibrary.NfcType) {

	data := url.Values{
		WasteLibrary.HTTP_READERTYPE: {WasteLibrary.READERTYPE_CAM},
		WasteLibrary.HTTP_DATA:       {readerDataTypeVal.ToString()},
	}

	WasteLibrary.HttpPostReq("http://127.0.0.1:10003/trigger", data)
}

func camCheck() {
	for {
		ifaces, err := net.Interfaces()
		if err != nil {
			WasteLibrary.LogErr(err)
			WasteLibrary.CurrentCheckStatu.ConnStatu = WasteLibrary.STATU_PASSIVE
		}

		WasteLibrary.CurrentCheckStatu.ConnStatu = WasteLibrary.STATU_PASSIVE
		for _, i := range ifaces {
			addrs, err := i.Addrs()
			if err != nil {
				WasteLibrary.LogErr(err)
				WasteLibrary.CurrentCheckStatu.ConnStatu = WasteLibrary.STATU_PASSIVE
			}
			for _, addr := range addrs {
				var ip net.IP
				switch v := addr.(type) {
				case *net.IPNet:
					ip = v.IP
				case *net.IPAddr:
					ip = v.IP
				}
				ipStr := fmt.Sprintf("%s", ip)
				if ipStr == "10.0.0.1" {
					WasteLibrary.CurrentCheckStatu.ConnStatu = WasteLibrary.STATU_ACTIVE
				}
			}
		}
		time.Sleep(opInterval * time.Second)
	}
}
