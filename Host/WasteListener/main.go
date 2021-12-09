package main

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/devafatek/WasteLibrary"
)

func initStart() {

	WasteLibrary.LogStr("Successfully connected!")
	go WasteLibrary.InitLog()
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
			continue
		}
		go handleTcpRequest(conn)
	}
}

func handleTcpRequest(conn net.Conn) {
	buf := make([]byte, 1024)
	reqLen, err := conn.Read(buf)
	if err != nil {
		return
	} else {
		var strBuf = string(buf)
		if reqLen == 215 {
			split := strings.Split(strBuf, "#")
			if len(split) == 32 {
				deviceType := split[0]
				imei := split[1]
				therm := split[2]
				battery := split[3]
				//latitude := split[4]
				//longitude := split[5]
				imsi := split[6]
				ultCount := split[7]

				if deviceType == WasteLibrary.DEVICETYPE_ULT && imsi != "***************" {

					resultVal := WasteLibrary.GetRedisForStoreApi("0", WasteLibrary.REDIS_SERIAL_ALARM, imsi)
					if resultVal.Result == WasteLibrary.RESULT_OK {
						conn.Write([]byte(WasteLibrary.RESULT_OK + " - " + resultVal.Retval.(string)))
						conn.Close()
					} else {
						conn.Write([]byte(WasteLibrary.RESULT_OK))
						conn.Close()
					}

					var currentHttpHeader WasteLibrary.HttpClientHeaderType
					currentHttpHeader.New()
					currentHttpHeader.ReaderType = WasteLibrary.READERTYPE_ULT
					currentHttpHeader.DeviceNo = imsi
					currentHttpHeader.DeviceType = WasteLibrary.DEVICETYPE_ULT
					var currentDevice WasteLibrary.UltDeviceType
					currentDevice.New()
					currentDevice.DeviceMain.SerialNumber = imsi
					currentDevice.DeviceStatu.StatusTime = WasteLibrary.GetTime()
					currentDevice.DeviceStatu.AliveStatus = WasteLibrary.STATU_ACTIVE
					currentDevice.DeviceStatu.AliveLastOkTime = WasteLibrary.GetTime()
					currentDevice.DeviceTherm.Therm = therm
					currentDevice.DeviceTherm.ThermTime = WasteLibrary.GetTime()
					currentDevice.DeviceBattery.Battery = battery
					currentDevice.DeviceBattery.BatteryTime = WasteLibrary.GetTime()

					currentDevice.DeviceSens.UltCount = WasteLibrary.StringIdToFloat64(ultCount)
					tempCount, _ := strconv.Atoi(ultCount)
					currentDevice.DeviceSens.UltRange1 = WasteLibrary.StringIdToFloat64(split[tempCount+7])
					tempCount--
					if tempCount > 0 {
						currentDevice.DeviceSens.UltRange2 = WasteLibrary.StringIdToFloat64(split[tempCount+7])
						tempCount--
					}
					if tempCount > 0 {
						currentDevice.DeviceSens.UltRange3 = WasteLibrary.StringIdToFloat64(split[tempCount+7])
						tempCount--
					}

					if tempCount > 0 {
						currentDevice.DeviceSens.UltRange4 = WasteLibrary.StringIdToFloat64(split[tempCount+7])
						tempCount--
					}

					if tempCount > 0 {
						currentDevice.DeviceSens.UltRange5 = WasteLibrary.StringIdToFloat64(split[tempCount+7])
						tempCount--
					}

					if tempCount > 0 {
						currentDevice.DeviceSens.UltRange6 = WasteLibrary.StringIdToFloat64(split[tempCount+7])
						tempCount--
					}

					if tempCount > 0 {
						currentDevice.DeviceSens.UltRange7 = WasteLibrary.StringIdToFloat64(split[tempCount+7])
						tempCount--
					}

					if tempCount > 0 {
						currentDevice.DeviceSens.UltRange8 = WasteLibrary.StringIdToFloat64(split[tempCount+7])
						tempCount--
					}

					if tempCount > 0 {
						currentDevice.DeviceSens.UltRange9 = WasteLibrary.StringIdToFloat64(split[tempCount+7])
						tempCount--
					}

					if tempCount > 0 {
						currentDevice.DeviceSens.UltRange10 = WasteLibrary.StringIdToFloat64(split[tempCount+7])
						tempCount--
					}

					if tempCount > 0 {
						currentDevice.DeviceSens.UltRange11 = WasteLibrary.StringIdToFloat64(split[tempCount+7])
						tempCount--
					}

					if tempCount > 0 {
						currentDevice.DeviceSens.UltRange12 = WasteLibrary.StringIdToFloat64(split[tempCount+7])
						tempCount--
					}

					if tempCount > 0 {
						currentDevice.DeviceSens.UltRange13 = WasteLibrary.StringIdToFloat64(split[tempCount+7])
						tempCount--
					}

					if tempCount > 0 {
						currentDevice.DeviceSens.UltRange14 = WasteLibrary.StringIdToFloat64(split[tempCount+7])
						tempCount--
					}

					if tempCount > 0 {
						currentDevice.DeviceSens.UltRange15 = WasteLibrary.StringIdToFloat64(split[tempCount+7])
						tempCount--
					}

					if tempCount > 0 {
						currentDevice.DeviceSens.UltRange16 = WasteLibrary.StringIdToFloat64(split[tempCount+7])
						tempCount--
					}

					if tempCount > 0 {
						currentDevice.DeviceSens.UltRange17 = WasteLibrary.StringIdToFloat64(split[tempCount+7])
						tempCount--
					}

					if tempCount > 0 {
						currentDevice.DeviceSens.UltRange18 = WasteLibrary.StringIdToFloat64(split[tempCount+7])
						tempCount--
					}

					if tempCount > 0 {
						currentDevice.DeviceSens.UltRange19 = WasteLibrary.StringIdToFloat64(split[tempCount+7])
						tempCount--
					}
					if tempCount > 0 {
						currentDevice.DeviceSens.UltRange20 = WasteLibrary.StringIdToFloat64(split[tempCount+7])
						tempCount--
					}
					if tempCount > 0 {
						currentDevice.DeviceSens.UltRange21 = WasteLibrary.StringIdToFloat64(split[tempCount+7])
						tempCount--
					}
					if tempCount > 0 {
						currentDevice.DeviceSens.UltRange22 = WasteLibrary.StringIdToFloat64(split[tempCount+7])
						tempCount--
					}
					if tempCount > 0 {
						currentDevice.DeviceSens.UltRange23 = WasteLibrary.StringIdToFloat64(split[tempCount+7])
						tempCount--
					}
					if tempCount > 0 {
						currentDevice.DeviceSens.UltRange24 = WasteLibrary.StringIdToFloat64(split[tempCount+7])
						tempCount--
					}

					currentDevice.DeviceBase.Imei = imei
					currentDevice.DeviceBase.Imsi = imsi

					data := url.Values{
						WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
						WasteLibrary.HTTP_DATA:   {currentDevice.ToString()},
					}
					resultVal = WasteLibrary.HttpPostReq("http://waste-enhc-cluster-ip/data", data)

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
}

func data(w http.ResponseWriter, req *http.Request) {

	if WasteLibrary.AllowCors {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,access-control-allow-origin, access-control-allow-headers")
	}
	var resultVal WasteLibrary.ResultType
	if err := req.ParseForm(); err != nil {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_HTTP_PARSE
		return
	}
	var currentHttpHeader WasteLibrary.HttpClientHeaderType
	currentHttpHeader.StringToType(req.FormValue(WasteLibrary.HTTP_HEADER))
	if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_RFID || currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_RECY {
		resultVal = WasteLibrary.HttpPostReq("http://waste-enhc-cluster-ip/data", req.Form)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_HTTP_POST
		}
	} else {
	}
	w.Write(resultVal.ToByte())
}

func update(w http.ResponseWriter, req *http.Request) {

	if WasteLibrary.AllowCors {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,access-control-allow-origin, access-control-allow-headers")
	}
	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	//TO DO
	//udate proc

	if err := req.ParseForm(); err != nil {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_HTTP_PARSE
		w.Write(resultVal.ToByte())
		return
	}

	/*resultVal = WasteLibrary.HttpPostReq("http://waste-enhc-cluster-ip/update", req.Form)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_HTTP_POST
		w.Write(resultVal.ToByte())
		return
	}*/
	var currentHttpHeader WasteLibrary.HttpClientHeaderType
	currentHttpHeader.StringToType(req.FormValue(WasteLibrary.HTTP_HEADER))
	var updaterType WasteLibrary.UpdaterType
	updaterType.StringToType(req.FormValue(WasteLibrary.HTTP_DATA))

	if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_RFID {
		if currentHttpHeader.DeviceNo == "0000000064231683" ||
			currentHttpHeader.DeviceNo == "000000008ce06a67" {
			fmt.Println(req.FormValue(WasteLibrary.HTTP_HEADER))
			fmt.Println(req.FormValue(WasteLibrary.HTTP_DATA))
		}
		if updaterType.AppType == WasteLibrary.RFID_APPTYPE_TRANSFER {
			if updaterType.Version == "1" {
				updaterType.Version = "2"
				updaterType.AppType = WasteLibrary.RFID_APPNAME_TRANSFER
				resultVal.Result = WasteLibrary.RESULT_OK
				resultVal.Retval = updaterType.ToString()

			}
		}

		if updaterType.AppType == WasteLibrary.RFID_APPNAME_UPDATER {
			if updaterType.Version == "1" {
				updaterType.Version = "2"
				updaterType.AppType = WasteLibrary.RFID_APPNAME_UPDATER
				resultVal.Result = WasteLibrary.RESULT_OK
				resultVal.Retval = updaterType.ToString()

			}
		}

		if updaterType.AppType == WasteLibrary.RFID_APPTYPE_SYSTEM {
			if updaterType.Version == "1" {
				updaterType.Version = "2"
				updaterType.AppType = WasteLibrary.RFID_APPNAME_SYSTEM
				resultVal.Result = WasteLibrary.RESULT_OK
				resultVal.Retval = updaterType.ToString()

			}
		}
	}

	w.Write(resultVal.ToByte())

}
