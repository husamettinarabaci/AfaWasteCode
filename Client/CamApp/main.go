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
	var resultVal WasteLibrary.ResultType
	if err := req.ParseForm(); err != nil {
		WasteLibrary.LogErr(err)
		return
	}

	opType := req.FormValue(WasteLibrary.OPTYPE)
	WasteLibrary.LogStr(opType)

	resultVal.Result = WasteLibrary.FAIL
	if opType == WasteLibrary.RF {
		var readerDataTypeVal WasteLibrary.TagType = WasteLibrary.StringToTagType(req.FormValue(WasteLibrary.DATA))
		if integratedPortInt == 4 {
			integratedPortInt = 1
		}
		go doRecord(readerDataTypeVal, strconv.Itoa(integratedPortInt), true)
		integratedPortInt++
		resultVal.Result = WasteLibrary.OK
	} else {
		resultVal.Result = WasteLibrary.FAIL
	}
	w.Write(resultVal.ToByte())
}

func doRecord(readerDataTypeVal WasteLibrary.TagType, integratedPort string, repeat bool) {
	WasteLibrary.CurrentCheckStatu.DeviceStatu = WasteLibrary.PASSIVE
	WasteLibrary.LogStr("Do Record : " + readerDataTypeVal.Epc + " - " + integratedPort + " - " + readerDataTypeVal.UID + " - " + strconv.FormatBool(repeat))
	cmd := exec.Command("timeout", "30", "ffmpeg", "-y", "-v", "0", "-loglevel", "0", "-hide_banner", "-f", "mpegts", "-i", "udp://localhost:1000"+integratedPort, "-t", "7", "-vb", "128k", "-threads", "7", "-map", "0:0", "-map", "-0:1", "-map", "-0:2", "-c:v", "libx264", "-pix_fmt", "yuvj420p", "-f", "mp4", "WAIT_CAM/"+readerDataTypeVal.UID+".mp4")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := cmd.Run()

	if err != nil && !strings.Contains(err.Error(), "124") {
		WasteLibrary.LogErr(err)
		if repeat {
			WasteLibrary.LogStr("Do Record repeat for err : " + readerDataTypeVal.Epc + " - " + integratedPort + " - " + readerDataTypeVal.UID + " - " + strconv.FormatBool(repeat))
			doRecord(readerDataTypeVal, integratedPort, false)
			return
		}
	} else {
		time.Sleep(5 * time.Second)

		if WasteLibrary.IsFileExists("WAIT_CAM/" + readerDataTypeVal.UID + ".mp4") {
			fi, err := os.Stat("WAIT_CAM/" + readerDataTypeVal.UID + ".mp4")
			if err != nil {
				if repeat {
					WasteLibrary.LogStr("Do Record repeat for not file : " + readerDataTypeVal.Epc + " - " + integratedPort + " - " + readerDataTypeVal.UID + " - " + strconv.FormatBool(repeat))
					doRecord(readerDataTypeVal, integratedPort, false)
					return
				}
			}
			size := fi.Size()
			if size < 10000 {
				if repeat {
					WasteLibrary.LogStr("Do Record repeat for file size : " + readerDataTypeVal.Epc + " - " + integratedPort + " - " + readerDataTypeVal.UID + " - " + strconv.FormatBool(repeat))
					doRecord(readerDataTypeVal, integratedPort, false)
					return
				}
			} else {

				cmdPic := exec.Command("ffmpeg", "-y", "-i", "WAIT_CAM/"+readerDataTypeVal.UID+".mp4", "-vf", "thumbnail,scale=150:100", "-frames:v", "1", "WAIT_CAM/"+readerDataTypeVal.UID+".png")
				cmdPic.Run()

				WasteLibrary.CurrentCheckStatu.DeviceStatu = WasteLibrary.ACTIVE
				sendCam(readerDataTypeVal)
			}
		} else {
			if repeat {
				WasteLibrary.LogStr("Do Record repeat for not file : " + readerDataTypeVal.Epc + " - " + integratedPort + " - " + readerDataTypeVal.UID + " - " + strconv.FormatBool(repeat))
				doRecord(readerDataTypeVal, integratedPort, false)
				return
			}
		}
	}
}

func sendCam(readerDataTypeVal WasteLibrary.TagType) WasteLibrary.ResultType {
	var resultVal WasteLibrary.ResultType

	data := url.Values{
		WasteLibrary.OPTYPE: {WasteLibrary.CAM},
		WasteLibrary.DATA:   {readerDataTypeVal.ToString()},
	}

	resultVal = WasteLibrary.HttpPostReq("http://127.0.0.1:10000/trans", data)
	return resultVal
}

func camCheck() {
	for {

		ifaces, err := net.Interfaces()
		if err != nil {
			WasteLibrary.LogErr(err)
			WasteLibrary.CurrentCheckStatu.ConnStatu = WasteLibrary.PASSIVE
		}

		WasteLibrary.CurrentCheckStatu.ConnStatu = WasteLibrary.PASSIVE
		for _, i := range ifaces {
			addrs, err := i.Addrs()
			if err != nil {
				WasteLibrary.LogErr(err)
				WasteLibrary.CurrentCheckStatu.ConnStatu = WasteLibrary.PASSIVE
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
					WasteLibrary.CurrentCheckStatu.ConnStatu = WasteLibrary.ACTIVE
				}
			}
		}

		if time.Since(lastCamRelayTime).Seconds() > 60*60 && WasteLibrary.CurrentCheckStatu.ConnStatu == WasteLibrary.PASSIVE {

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
