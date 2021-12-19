package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"time"

	"github.com/devafatek/WasteLibrary"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}
var socketCh chan string
var currentUser string
var currentStatu string
var lastTime time.Time
var opInterval time.Duration = 60 * 60

func initStart() {

	WasteLibrary.LogStr("Successfully connected!")
	WasteLibrary.Version = "1"
	WasteLibrary.LogStr("Version : " + WasteLibrary.Version)
	currentUser = WasteLibrary.GetCurrentUser()
	WasteLibrary.LogStr(currentUser)
	go WasteLibrary.InitLog()
	socketCh = make(chan string)
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	currentStatu = WasteLibrary.RECY_SOCKET_INDEX
	lastTime = time.Now()
}

func main() {

	initStart()

	go getCustomer()
	go checkStatu()
	http.HandleFunc("/health", WasteLibrary.HealthHandler)
	http.HandleFunc("/readiness", WasteLibrary.ReadinessHandler)
	http.HandleFunc("/status", WasteLibrary.StatusHandler)
	http.HandleFunc("/trigger", trigger)
	http.HandleFunc("/socket", socket)
	http.ListenAndServe(":10003", nil)
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
	resultVal.Result = WasteLibrary.RESULT_OK
	resultVal.Retval = WasteLibrary.RESULT_SUCCESS_OK
	w.Write(resultVal.ToByte())

	readerType := req.FormValue(WasteLibrary.HTTP_READERTYPE)
	var nfcTypeVal WasteLibrary.NfcType
	nfcTypeVal.StringToType(req.FormValue(WasteLibrary.HTTP_DATA))

	if readerType == WasteLibrary.READERTYPE_RF {

		resultVal = sendRf(nfcTypeVal)
		if resultVal.Result == WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RECY_SOCKET_ANALYZE
			nfcTypeVal.StringToType(resultVal.Retval.(string))
			go sendRfToCam(nfcTypeVal)
			resultVal.Retval = nfcTypeVal.ToString()
			socketCh <- resultVal.ToString()
			currentStatu = resultVal.Result
			lastTime = time.Now()
		} else {
			resultVal.Result = WasteLibrary.RECY_SOCKET_ERROR
			socketCh <- resultVal.ToString()
			currentStatu = resultVal.Result
			lastTime = time.Now()
		}
	} else if readerType == WasteLibrary.READERTYPE_CAM {
		go sendMotor()
		resultVal = getResult(nfcTypeVal)
		if resultVal.Result == WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RECY_SOCKET_FINISH
			nfcTypeVal.StringToType(resultVal.Retval.(string))
			resultVal.Retval = nfcTypeVal.ToString()

			socketCh <- resultVal.ToString()
			currentStatu = resultVal.Result
			lastTime = time.Now()
		} else {
			resultVal.Result = WasteLibrary.RECY_SOCKET_ERROR
			socketCh <- resultVal.ToString()
			currentStatu = resultVal.Result
			lastTime = time.Now()
		}

		time.Sleep(5 * time.Second)
		resultVal.Result = WasteLibrary.RECY_SOCKET_INDEX
		socketCh <- resultVal.ToString()
		currentStatu = resultVal.Result
		lastTime = time.Now()
	}

}

func getCustomer() {
	for {
		var resultVal WasteLibrary.ResultType
		resultVal.Result = WasteLibrary.RESULT_FAIL
		data := url.Values{
			WasteLibrary.HTTP_READERTYPE: {WasteLibrary.READERTYPE_GET_CUSTOMER},
		}
		resultVal = WasteLibrary.HttpPostReq("http://127.0.0.1:10000/trans", data)
		if resultVal.Result == WasteLibrary.RESULT_OK {
			var customer WasteLibrary.CustomerType
			customer.StringToType(resultVal.Retval.(string))
			checkCustomer(customer)
			time.Sleep(opInterval * time.Second)
		} else {
			time.Sleep(5 * 60 * time.Second)
		}

	}
}

func sendRf(nfcData WasteLibrary.NfcType) WasteLibrary.ResultType {
	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	data := url.Values{
		WasteLibrary.HTTP_READERTYPE: {WasteLibrary.READERTYPE_RF},
		WasteLibrary.HTTP_DATA:       {nfcData.ToString()},
	}
	resultVal = WasteLibrary.HttpPostReq("http://127.0.0.1:10000/trans", data)
	return resultVal
}

func sendRfToCam(nfcData WasteLibrary.NfcType) {

	data := url.Values{
		WasteLibrary.HTTP_READERTYPE: {WasteLibrary.READERTYPE_CAMTRIGGER},
		WasteLibrary.HTTP_DATA:       {nfcData.ToString()},
	}
	WasteLibrary.HttpPostReq("http://127.0.0.1:10002/trigger", data)
}

func getResult(nfcData WasteLibrary.NfcType) WasteLibrary.ResultType {
	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	data := url.Values{
		WasteLibrary.HTTP_READERTYPE: {WasteLibrary.READERTYPE_GET_NFC},
		WasteLibrary.HTTP_DATA:       {nfcData.ToString()},
	}
	resultVal = WasteLibrary.HttpPostReq("http://127.0.0.1:10000/trans", data)
	WasteLibrary.LogStr(resultVal.ToString())
	return resultVal
}

func sendMotor() {

	data := url.Values{
		WasteLibrary.HTTP_READERTYPE: {WasteLibrary.READERTYPE_MOTORTRIGGER},
	}

	WasteLibrary.HttpPostReq("http://127.0.0.1:10008/trigger", data)
}

func socket(w http.ResponseWriter, req *http.Request) {

	if WasteLibrary.AllowCors {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,access-control-allow-origin, access-control-allow-headers")
	}

	c, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		WasteLibrary.LogErr(err)
		return
	}
	defer c.Close()

	for {
		msg := <-socketCh
		err = c.WriteMessage(1, []byte(msg))
		if err != nil {
			WasteLibrary.LogErr(err)
			break
		}
	}
}

func checkStatu() {
	var resultVal WasteLibrary.ResultType
	for {
		second := time.Since(lastTime).Seconds()
		if second > 5*60 && currentStatu != WasteLibrary.RECY_SOCKET_INDEX {
			resultVal.Result = WasteLibrary.RECY_SOCKET_INDEX
			socketCh <- resultVal.ToString()
			currentStatu = resultVal.Result
			lastTime = time.Now()
		}
		time.Sleep(10 * time.Second)
	}

}

func checkCustomer(customer WasteLibrary.CustomerType) {
	var customerAutoStartVal string = fmt.Sprintf(`@xset s off
@xset -dpms
@xset s noblank
@chromium-browser --incognito --kiosk https://afatek-waste-recy-web-s3.s3.eu-central-1.amazonaws.com/%s/index.html
	`, customer.CustomerName)
	resultVal := readAutostart()
	if resultVal.Retval.(string) == customerAutoStartVal {
	} else {
		writeAutostart(customerAutoStartVal)
		reboot()
	}
}

func readAutostart() WasteLibrary.ResultType {
	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	dat, err := os.ReadFile("/home/pi/.config/lxsession/LXDE-pi/autostart")
	if err != nil {

	} else {
		resultVal.Result = WasteLibrary.RESULT_OK
		resultVal.Retval = string(dat)
	}
	return resultVal
}

func writeAutostart(customerAutoStartVal string) WasteLibrary.ResultType {
	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL

	err := ioutil.WriteFile("/home/pi/.config/lxsession/LXDE-pi/autostart", []byte(customerAutoStartVal), 0644)
	if err != nil {

	} else {
		resultVal.Result = WasteLibrary.RESULT_OK
		resultVal.Retval = WasteLibrary.RESULT_SUCCESS_OK
	}
	return resultVal
}

func reboot() WasteLibrary.ResultType {
	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	cmd := exec.Command("sudo", "./reboot")
	err := cmd.Start()
	if err != nil {
		WasteLibrary.LogErr(err)
		resultVal.Result = WasteLibrary.RESULT_FAIL
	} else {
		resultVal.Result = WasteLibrary.RESULT_OK
		resultVal.Retval = WasteLibrary.RESULT_SUCCESS_OK
	}
	return resultVal
}
