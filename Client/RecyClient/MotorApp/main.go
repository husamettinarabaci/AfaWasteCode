package main

import (
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"gitee.com/wiseai/go-rpio"
	"github.com/devafatek/WasteLibrary"
)

var opInterval time.Duration = 1 * 60
var wg sync.WaitGroup
var currentUser string
var turnLeftEnb string = os.Getenv("MOTOR_L_ENB")
var turnLeftPwm string = os.Getenv("MOTOR_L_PWM")
var turnRigthEnb string = os.Getenv("MOTOR_R_ENB")
var turnRigthPwm string = os.Getenv("MOTOR_R_PWM")

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

	http.HandleFunc("/status", WasteLibrary.StatusHandler)
	http.HandleFunc("/trigger", trigger)
	http.ListenAndServe(":10008", nil)
}

func trigger(w http.ResponseWriter, req *http.Request) {

	if WasteLibrary.AllowCors {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
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
	if readerType == WasteLibrary.READERTYPE_MOTORTRIGGER {
		motorProc()
		resultVal.Result = WasteLibrary.RESULT_OK
	} else {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_READERTYPE
		w.Write(resultVal.ToByte())

		return
	}
	w.Write(resultVal.ToByte())

}

func motorProc() {

	rpio.Open()
	leftEnb, _ := strconv.Atoi(turnLeftEnb)
	rightEnb, _ := strconv.Atoi(turnRigthEnb)
	leftPwm, _ := strconv.Atoi(turnLeftPwm)
	rightPwm, _ := strconv.Atoi(turnRigthPwm)
	pinLeftEnb := rpio.Pin(leftEnb)
	pinRightEnb := rpio.Pin(rightEnb)
	pinLeftPwm := rpio.Pin(leftPwm)
	pinRightPwm := rpio.Pin(rightPwm)
	pinLeftEnb.Output()
	pinRightEnb.Output()
	pinLeftPwm.Output()
	pinRightPwm.Output()

	pinLeftEnb.High()
	pinRightEnb.High()
	pinLeftPwm.Low()
	pinRightPwm.Low()
	time.Sleep(1 * time.Second)
	pinLeftPwm.High()
	pinRightPwm.Low()
	time.Sleep(3 * time.Second)
	pinLeftPwm.Low()
	pinRightPwm.High()
	time.Sleep(3 * time.Second)
	pinLeftEnb.Low()
	pinRightEnb.Low()
	pinLeftPwm.Low()
	pinRightPwm.Low()
	rpio.Close()

}
