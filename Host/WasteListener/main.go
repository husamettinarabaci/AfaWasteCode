package main

import (
	"math"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/AfatekDevelopers/result_lib_go/devafatekresult"
	"github.com/devafatek/WasteLibrary"
)

func initStart() {

	WasteLibrary.LogStr("Successfully connected!")
}

const (
	CONN_HOST = "0.0.0.0"
	CONN_PORT = "20000"
	CONN_TYPE = "tcp"
)

func main() {

	initStart()

	go tcpServer()
	http.HandleFunc("/health", WasteLibrary.HealthHandler)
	http.HandleFunc("/readiness", WasteLibrary.ReadinessHandler)
	http.HandleFunc("/status", WasteLibrary.StatusHandler)
	http.HandleFunc("/data", data)
	http.ListenAndServe(":80", nil)
}

func tcpServer() {
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		WasteLibrary.LogErr(err)
		os.Exit(1)
	}
	defer l.Close()
	WasteLibrary.LogStr("Listening on " + CONN_HOST + ":" + CONN_PORT)
	for {
		conn, err := l.Accept()
		if err != nil {
			WasteLibrary.LogErr(err)
			os.Exit(1)
		}
		go handleTcpRequest(conn)
	}
}

func handleTcpRequest(conn net.Conn) {
	buf := make([]byte, 1024)
	reqLen, err := conn.Read(buf)
	if err != nil {
		WasteLibrary.LogErr(err)
	}
	if reqLen == 111 {
		var strBuf = string(buf)
		WasteLibrary.LogStr(strBuf)
		split := strings.Split(strBuf, "#")
		if len(split) == 17 {
			appType := split[0]
			opType := split[1]
			serialNo := split[2]
			temp := split[3]
			battery := split[4]
			lat := split[5]
			long := split[6]
			r1 := WasteLibrary.StringToFloat64(split[7])
			r2 := WasteLibrary.StringToFloat64(split[8])
			r3 := WasteLibrary.StringToFloat64(split[9])
			r4 := WasteLibrary.StringToFloat64(split[10])
			r5 := WasteLibrary.StringToFloat64(split[11])
			r6 := WasteLibrary.StringToFloat64(split[12])
			r7 := WasteLibrary.StringToFloat64(split[13])
			r8 := WasteLibrary.StringToFloat64(split[14])
			r9 := WasteLibrary.StringToFloat64(split[15])
			r10 := WasteLibrary.StringToFloat64(split[16])
			var ultrange = math.Mod((r1+r2+r3+r4+r5+r6+r7+r8+r9+r10)/10, 10)

			if appType == "ULT" && (opType == "SENS" || opType == "ATMP" || opType == "AGPS") {
				conn.Write([]byte("OK"))
				conn.Close()

				var currentHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.HttpClientHeaderType{}
				currentHttpHeader.AppType = appType
				currentHttpHeader.OpType = opType
				currentHttpHeader.DeviceNo = serialNo
				currentHttpHeader.Time = time.Now().String()
				currentHttpHeader.Repeat = 0
				currentHttpHeader.DeviceId = 0
				currentHttpHeader.CustomerId = 0
				var currentDevice WasteLibrary.DeviceType = WasteLibrary.DeviceType{}
				currentDevice.DeviceType = appType
				currentDevice.SerialNumber = serialNo
				currentDevice.TransferAppStatus = "1"
				currentDevice.TransferAppLastOkTime = time.Now().String()
				currentDevice.AliveStatus = "1"
				currentDevice.AliveLastOkTime = time.Now().String()

				if lat != "00.00000" && long != "000.00000" {
					currentDevice.GpsAppStatus = "1"
					currentDevice.GpsAppLastOkTime = time.Now().String()
					currentDevice.GpsConnStatus = "1"
					currentDevice.GpsConnLastOkTime = time.Now().String()
					currentDevice.GpsStatus = "1"
					currentDevice.GpsLastOkTime = time.Now().String()
					currentDevice.Latitude = WasteLibrary.StringToFloat64(lat)
					currentDevice.Longitude = WasteLibrary.StringToFloat64(long)
					currentDevice.GpsTime = time.Now().String()
				}
				if temp != "00" {
					currentDevice.ThermAppStatus = "1"
					currentDevice.ThermAppLastOkTime = time.Now().String()
					currentDevice.Therm = temp
					currentDevice.ThermTime = time.Now().String()
					currentDevice.ThermStatus = "1"
				}
				if battery != "0000" {
					currentDevice.Battery = battery
					currentDevice.BatteryTime = time.Now().String()

					if WasteLibrary.StringToFloat64(battery) > 3200 {
						currentDevice.BatteryStatus = "1"
					} else {
						currentDevice.BatteryStatus = "2"
					}

				}

				currentDevice.DeviceStatus = "1"
				currentDevice.StatusTime = time.Now().String()
				if ultrange != 0 {
					currentDevice.UltRange = ultrange
					currentDevice.UltStatus = "0"
					currentDevice.UltTime = time.Now().String()
				}
				if opType == "ATMP" || opType == "AGPS" {
					currentDevice.AlarmStatus = "1"
					currentDevice.AlarmTime = time.Now().String()
					currentDevice.AlarmType = opType
					currentDevice.Alarm = opType
				}
				WasteLibrary.LogStr(currentHttpHeader.ToString())
				WasteLibrary.LogStr(currentDevice.ToString())
				/*data := url.Values{
					"HEADER": {currentHttpHeader.ToString()},
					"DATA":   {currentDevice.ToString()},
				}*/
				//resultVal := WasteLibrary.HttpPostReq("http://waste-enhc-cluster-ip/data", data)
				//WasteLibrary.LogStr(resultVal.ToString())
			} else {
				conn.Write([]byte("FAIL"))
				conn.Close()
			}
		} else {
			conn.Write([]byte("FAIL"))
			conn.Close()
		}
	} else {
		conn.Write([]byte("FAIL"))
		conn.Close()
	}
}

func data(w http.ResponseWriter, req *http.Request) {
	var resultVal devafatekresult.ResultType
	resultVal.Result = "FAIL"
	if err := req.ParseForm(); err != nil {
		WasteLibrary.LogErr(err)
		return
	}
	resultVal = WasteLibrary.HttpPostReq("http://waste-enhc-cluster-ip/data", req.Form)
	w.Write(resultVal.ToByte())
}
