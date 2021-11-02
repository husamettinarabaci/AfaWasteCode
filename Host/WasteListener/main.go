package main

import (
	"math"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"

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
	http.HandleFunc("/update", update)
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
	var strBuf = string(buf)
	WasteLibrary.LogStr(strBuf)
	if reqLen == 128 {

		split := strings.Split(strBuf, "#")
		if len(split) == 17 {
			appType := split[0]
			imei := split[1]
			therm := split[2]
			battery := split[3]
			var r1 float64 = 0
			if split[6] != "*****" {
				r1 = WasteLibrary.StringToFloat64(split[6])
			}
			var r2 float64 = 0
			if split[7] != "*****" {
				r2 = WasteLibrary.StringToFloat64(split[7])
			}
			var r3 float64 = 0
			if split[8] != "*****" {
				r3 = WasteLibrary.StringToFloat64(split[8])
			}
			var r4 float64 = 0
			if split[9] != "*****" {
				r4 = WasteLibrary.StringToFloat64(split[9])
			}
			var r5 float64 = 0
			if split[10] != "*****" {
				r5 = WasteLibrary.StringToFloat64(split[10])
			}
			var r6 float64 = 0
			if split[11] != "*****" {
				r6 = WasteLibrary.StringToFloat64(split[11])
			}
			var r7 float64 = 0
			if split[12] != "*****" {
				r7 = WasteLibrary.StringToFloat64(split[12])
			}
			var r8 float64 = 0
			if split[13] != "*****" {
				r8 = WasteLibrary.StringToFloat64(split[13])
			}
			var r9 float64 = 0
			if split[14] != "*****" {
				r9 = WasteLibrary.StringToFloat64(split[14])
			}
			var r10 float64 = 0
			if split[15] != "*****" {
				r10 = WasteLibrary.StringToFloat64(split[15])
			}
			imsi := split[16]
			var ultrange = math.Mod((r1+r2+r3+r4+r5+r6+r7+r8+r9+r10)/10, 10)

			if appType == WasteLibrary.APPTYPE_ULT && imsi != "***************" {

				resultVal := WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_SERIAL_ALARM, imsi)
				if resultVal.Result == WasteLibrary.RESULT_OK {
					conn.Write([]byte(WasteLibrary.RESULT_OK + " - " + resultVal.Retval.(string)))
					conn.Close()
				} else {
					conn.Write([]byte(WasteLibrary.RESULT_OK))
					conn.Close()
				}

				var currentHttpHeader WasteLibrary.HttpClientHeaderType
				currentHttpHeader.New()
				currentHttpHeader.AppType = appType
				currentHttpHeader.ReaderType = WasteLibrary.READERTYPE_ULT
				currentHttpHeader.DeviceNo = imsi
				currentHttpHeader.DeviceType = WasteLibrary.DEVICETYPE_ULT
				var currentDevice WasteLibrary.UltDeviceType
				currentDevice.New()
				currentDevice.SerialNumber = imsi
				currentDevice.DeviceStatu.StatusTime = WasteLibrary.GetTime()
				currentDevice.DeviceStatu.AliveStatus = WasteLibrary.STATU_ACTIVE
				currentDevice.DeviceStatu.AliveLastOkTime = WasteLibrary.GetTime()
				if therm != "**" {
					currentDevice.DeviceTherm.Therm = therm
					currentDevice.DeviceTherm.ThermTime = WasteLibrary.GetTime()
				}
				if battery != "****" {
					currentDevice.DeviceBattery.Battery = battery
					currentDevice.DeviceBattery.BatteryTime = WasteLibrary.GetTime()
				}
				if ultrange != 0 {
					currentDevice.DeviceSens.UltTime = WasteLibrary.GetTime()
					currentDevice.DeviceSens.UltRange = ultrange
				}
				currentDevice.DeviceBase.Imei = imei
				currentDevice.DeviceBase.Imsi = imsi

				data := url.Values{
					WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
					WasteLibrary.HTTP_DATA:   {currentDevice.ToString()},
				}
				resultVal = WasteLibrary.HttpPostReq("http://waste-enhc-cluster-ip/data", data)
				WasteLibrary.LogStr(resultVal.ToString())

			} else {
				conn.Write([]byte(WasteLibrary.RESULT_FAIL))
				conn.Close()
			}
		} else {
			conn.Write([]byte(WasteLibrary.RESULT_FAIL))
			conn.Close()
		}
	} else {
		conn.Write([]byte(WasteLibrary.RESULT_FAIL))
		conn.Close()
	}
}

func data(w http.ResponseWriter, req *http.Request) {
	if WasteLibrary.AllowCors {

		w.Header().Set("Access-Control-Allow-Origin", "*")
	}
	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	if err := req.ParseForm(); err != nil {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_HTTP_PARSE
		w.Write(resultVal.ToByte())
		WasteLibrary.LogErr(err)
		return
	}
	resultVal = WasteLibrary.HttpPostReq("http://waste-enhc-cluster-ip/data", req.Form)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_HTTP_POST
		w.Write(resultVal.ToByte())
		return
	}
	w.Write(resultVal.ToByte())
}

func update(w http.ResponseWriter, req *http.Request) {
	if WasteLibrary.AllowCors {

		w.Header().Set("Access-Control-Allow-Origin", "*")
	}
	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	//TO DO
	//udate proc
	/*resultVal.Result = WasteLibrary.RESULT_FAIL
	if err := req.ParseForm(); err != nil {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_HTTP_PARSE
		w.Write(resultVal.ToByte())
		WasteLibrary.LogErr(err)
		return
	}
	resultVal = WasteLibrary.HttpPostReq("http://waste-enhc-cluster-ip/update", req.Form)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_HTTP_POST
		w.Write(resultVal.ToByte())
		return
	}*/
	w.Write(resultVal.ToByte())
}
