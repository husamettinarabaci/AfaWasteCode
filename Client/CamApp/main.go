package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"strings"
	"sync"
	"time"

	"gitee.com/wiseai/go-rpio"
)

var debug bool = os.Getenv("DEBUG") == "1"
var camPort string = os.Getenv("CAM_PORT")
var opInterval time.Duration = 5 * 60
var wg sync.WaitGroup
var appStatus string = "1"
var connStatus string = "0"
var camStatus string = "0"
var integratedPortInt = 1
var currentUser string
var lastCamRelayTime time.Time

type rfType struct {
	TagID string `json:"TagID"`
	UID   string `json:"UID"`
}

func initStart() {

	if !debug {
		time.Sleep(60 * time.Second)
	}
	lastCamRelayTime = time.Now()
	logStr("Successfully connected!")
	currentUser = getCurrentUser()
	logStr(currentUser)
}
func main() {

	initStart()

	time.Sleep(5 * time.Second)
	go camCheck()
	wg.Add(1)

	http.HandleFunc("/status", status)
	http.HandleFunc("/trigger", trigger)
	http.ListenAndServe(":10002", nil)
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
	} else if opType == "CAM" {
		if camStatus == "1" {
			w.Write([]byte("OK"))
		} else {
			w.Write([]byte("FAIL"))
		}
	} else {
		w.Write([]byte("FAIL"))
	}
}

func trigger(w http.ResponseWriter, req *http.Request) {

	if err := req.ParseForm(); err != nil {
		logErr(err)
		return
	}
	opType := req.FormValue("OPTYPE")
	logStr(opType)

	if opType == "RF" {
		rf := req.FormValue("RF")
		var currentRfType rfType
		json.Unmarshal([]byte(rf), &currentRfType)
		if integratedPortInt == 3 {
			integratedPortInt = 1
		}
		doRecord(currentRfType, strconv.Itoa(integratedPortInt), true)
	} else {
		w.Write([]byte("FAIL"))
	}
}

func doRecord(currentRfType rfType, integratedPort string, repeat bool) {
	camStatus = "0"
	logStr("Do Record : " + currentRfType.TagID + " - " + integratedPort + " - " + currentRfType.UID + " - " + strconv.FormatBool(repeat))
	cmd := exec.Command("timeout", "30", "ffmpeg", "-y", "-v", "0", "-loglevel", "0", "-hide_banner", "-f", "mpegts", "-i", "udp://localhost:1000"+integratedPort, "-t", "7", "-vb", "128k", "-threads", "7", "-map", "0:0", "-map", "-0:1", "-map", "-0:2", "-c:v", "libx264", "-pix_fmt", "yuvj420p", "-f", "mp4", "WAIT_CAM/"+currentRfType.UID+".mp4")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := cmd.Run()

	if err != nil && !strings.Contains(err.Error(), "124") {
		logErr(err)
		if repeat {
			logStr("Do Record repeat for err : " + currentRfType.TagID + " - " + integratedPort + " - " + currentRfType.UID + " - " + strconv.FormatBool(repeat))
			doRecord(currentRfType, integratedPort, false)
			return
		}
	} else {
		time.Sleep(5 * time.Second)

		if fileExists("WAIT_CAM/" + currentRfType.UID + ".mp4") {
			fi, err := os.Stat("WAIT_CAM/" + currentRfType.UID + ".mp4")
			if err != nil {
				if repeat {
					logStr("Do Record repeat for not file : " + currentRfType.TagID + " - " + integratedPort + " - " + currentRfType.UID + " - " + strconv.FormatBool(repeat))
					doRecord(currentRfType, integratedPort, false)
					return
				}
			}
			size := fi.Size()
			if size < 10000 {
				if repeat {
					logStr("Do Record repeat for file size : " + currentRfType.TagID + " - " + integratedPort + " - " + currentRfType.UID + " - " + strconv.FormatBool(repeat))
					doRecord(currentRfType, integratedPort, false)
					return
				}
			} else {
				camStatus = "1"
				sendCam(currentRfType)
			}
		} else {
			if repeat {
				logStr("Do Record repeat for not file : " + currentRfType.TagID + " - " + integratedPort + " - " + currentRfType.UID + " - " + strconv.FormatBool(repeat))
				doRecord(currentRfType, integratedPort, false)
				return
			}
		}
	}
}

func sendCam(currentRfType rfType) {
	data := url.Values{
		"OPTYPE": {"CAM"},
		"UID":    {currentRfType.UID},
		"TAGID":  {currentRfType.TagID},
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

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func camCheck() {
	for {

		ifaces, err := net.Interfaces()
		if err != nil {
			logErr(err)
			connStatus = "0"
		}

		connStatus = "0"
		for _, i := range ifaces {
			addrs, err := i.Addrs()
			if err != nil {
				logErr(err)
				connStatus = "0"
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
				logStr(ipStr)
				if ipStr == "10.0.0.1" {
					connStatus = "1"
				}
			}
		}

		if time.Since(lastCamRelayTime).Seconds() > 60*60 && connStatus == "0" {

			lastCamRelayTime = time.Now()
			logStr("Restart cam...")
			rpio.Open()
			logStr(camPort)
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
