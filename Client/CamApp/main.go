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

	"gitee.com/wiseai/go-rpio"
	"github.com/devafatek/WasteLibrary"
)

var camPort string = os.Getenv("CAM_PORT")
var opInterval time.Duration = 5 * 60
var wg sync.WaitGroup
var integratedPortInt = 1
var currentUser string
var lastCamRelayTime time.Time

func initStart() {
	time.Sleep(5 * time.Second)

	lastCamRelayTime = time.Now()
	WasteLibrary.LogStr("Successfully connected!")
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
	}
	var resultVal WasteLibrary.ResultType
	if err := req.ParseForm(); err != nil {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_HTTP_PARSE
		w.Write(resultVal.ToByte())
		WasteLibrary.LogErr(err)
		return
	}

	opType := req.FormValue(WasteLibrary.HTTP_OPTYPE)
	WasteLibrary.LogStr(opType)

	resultVal.Result = WasteLibrary.RESULT_FAIL
	if opType == WasteLibrary.OPTYPE_RF {
		var readerDataTypeVal WasteLibrary.TagType = WasteLibrary.StringToTagType(req.FormValue(WasteLibrary.HTTP_DATA))
		if integratedPortInt == 4 {
			integratedPortInt = 1
		}
		go doRecord(readerDataTypeVal, strconv.Itoa(integratedPortInt), WasteLibrary.STATU_ACTIVE)
		integratedPortInt++
		resultVal.Result = WasteLibrary.RESULT_OK
	} else {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_OPTYPE
		w.Write(resultVal.ToByte())
		return
	}
	w.Write(resultVal.ToByte())
}

func doRecord(readerDataTypeVal WasteLibrary.TagType, integratedPort string, repeat string) {
	WasteLibrary.CurrentCheckStatu.DeviceStatu = WasteLibrary.STATU_PASSIVE
	WasteLibrary.LogStr("Do Record : " + readerDataTypeVal.Epc + " - " + integratedPort + " - " + readerDataTypeVal.UID + " - " + repeat)
	cmd := exec.Command("timeout", "30", "ffmpeg", "-y", "-v", "0", "-loglevel", "0", "-hide_banner", "-f", "mpegts", "-i", "udp://localhost:1000"+integratedPort, "-t", "7", "-vb", "128k", "-threads", "7", "-map", "0:0", "-map", "-0:1", "-map", "-0:2", "-c:v", "libx264", "-pix_fmt", "yuvj420p", "-f", "mp4", "WAIT_CAM/"+readerDataTypeVal.UID+".mp4")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := cmd.Run()

	if err != nil && !strings.Contains(err.Error(), "124") {
		WasteLibrary.LogErr(err)
		if repeat == WasteLibrary.STATU_ACTIVE {
			WasteLibrary.LogStr("Do Record repeat for err : " + readerDataTypeVal.Epc + " - " + integratedPort + " - " + readerDataTypeVal.UID + " - " + repeat)
			doRecord(readerDataTypeVal, integratedPort, WasteLibrary.STATU_PASSIVE)
			return
		}
	} else {
		time.Sleep(5 * time.Second)

		if WasteLibrary.IsFileExists("WAIT_CAM/" + readerDataTypeVal.UID + ".mp4") {
			fi, err := os.Stat("WAIT_CAM/" + readerDataTypeVal.UID + ".mp4")
			if err != nil {
				if repeat == WasteLibrary.STATU_ACTIVE {
					WasteLibrary.LogStr("Do Record repeat for not file : " + readerDataTypeVal.Epc + " - " + integratedPort + " - " + readerDataTypeVal.UID + " - " + repeat)
					doRecord(readerDataTypeVal, integratedPort, WasteLibrary.STATU_PASSIVE)
					return
				}
			}
			size := fi.Size()
			if size < 10000 {
				if repeat == WasteLibrary.STATU_ACTIVE {
					WasteLibrary.LogStr("Do Record repeat for file size : " + readerDataTypeVal.Epc + " - " + integratedPort + " - " + readerDataTypeVal.UID + " - " + repeat)
					doRecord(readerDataTypeVal, integratedPort, WasteLibrary.STATU_PASSIVE)
					return
				}
			} else {

				cmdPic := exec.Command("ffmpeg", "-y", "-i", "WAIT_CAM/"+readerDataTypeVal.UID+".mp4", "-vf", "thumbnail,scale=150:100", "-frames:v", "1", "WAIT_CAM/"+readerDataTypeVal.UID+".png")
				cmdPic.Run()

				WasteLibrary.CurrentCheckStatu.DeviceStatu = WasteLibrary.STATU_ACTIVE
				sendCam(readerDataTypeVal)
			}
		} else {
			if repeat == WasteLibrary.STATU_ACTIVE {
				WasteLibrary.LogStr("Do Record repeat for not file : " + readerDataTypeVal.Epc + " - " + integratedPort + " - " + readerDataTypeVal.UID + " - " + repeat)
				doRecord(readerDataTypeVal, integratedPort, WasteLibrary.STATU_PASSIVE)
				return
			}
		}
	}
}

func sendCam(readerDataTypeVal WasteLibrary.TagType) {

	data := url.Values{
		WasteLibrary.HTTP_OPTYPE: {WasteLibrary.OPTYPE_CAM},
		WasteLibrary.HTTP_DATA:   {readerDataTypeVal.ToString()},
	}

	WasteLibrary.HttpPostReq("http://127.0.0.1:10000/trans", data)
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
				WasteLibrary.LogStr(ipStr)
				if ipStr == "10.0.0.1" {
					WasteLibrary.CurrentCheckStatu.ConnStatu = WasteLibrary.STATU_ACTIVE
				}
			}
		}

		if time.Since(lastCamRelayTime).Seconds() > 60*60 && WasteLibrary.CurrentCheckStatu.ConnStatu == WasteLibrary.STATU_PASSIVE {

			lastCamRelayTime = time.Now()
			WasteLibrary.LogStr("Restart cam...")
			rpio.Open()
			WasteLibrary.LogStr(camPort)
			camPort, _ := strconv.Atoi(camPort)
			pin := rpio.Pin(camPort)
			pin.Output()
			pin.High()
			time.Sleep(10 * time.Second)
			pin.Low()
			rpio.Close()
		}

		time.Sleep(opInterval * time.Second)
	}
}
