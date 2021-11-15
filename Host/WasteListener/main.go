package main

import (
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
			//ultRange1 := split[8]
			//ultRange2 := split[9]
			//ultRange3 := split[10]
			//ultRange4 := split[11]
			//ultRange5 := split[12]
			//ultRange6 := split[13]
			//ultRange7 := split[14]
			//ultRange8 := split[15]
			//ultRange9 := split[16]
			//ultRange10 := split[17]
			//ultRange11 := split[18]
			//ultRange12 := split[19]
			//ultRange13 := split[20]
			//ultRange14 := split[21]
			//ultRange15 := split[22]
			//ultRange16 := split[23]
			//ultRange17 := split[24]
			//ultRange18 := split[25]
			//ultRange19 := split[26]
			//ultRange20 := split[27]
			//ultRange21 := split[28]
			//ultRange22 := split[29]
			//ultRange23 := split[30]
			//ultRange24 := split[31]

			if deviceType == WasteLibrary.DEVICETYPE_ULT && imsi != "***************" {

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
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
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
	var currentHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.StringToHttpClientHeaderType(req.FormValue(WasteLibrary.HTTP_HEADER))
	WasteLibrary.LogStr("Header : " + currentHttpHeader.ToString())
	WasteLibrary.LogStr("Data : " + req.FormValue(WasteLibrary.HTTP_DATA))
	if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_RFID || currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_RECY {
		resultVal = WasteLibrary.HttpPostReq("http://waste-enhc-cluster-ip/data", req.Form)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_HTTP_POST
			w.Write(resultVal.ToByte())

			return
		}
	} else {
	}
	w.Write(resultVal.ToByte())

}

func update(w http.ResponseWriter, req *http.Request) {

	if WasteLibrary.AllowCors {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
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
	WasteLibrary.LogReport("Result : "+resultVal.ToString())
				WasteLibrary.LogErr(err)
				return
			}
			resultVal = WasteLibrary.HttpPostReq("http://waste-enhc-cluster-ip/update", req.Form)
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_HTTP_POST
				w.Write(resultVal.ToByte())
	WasteLibrary.LogReport("Result : "+resultVal.ToString())
				return
			}*/
	w.Write(resultVal.ToByte())

}
