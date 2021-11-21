package main

import (
	"net/http"
	"net/url"
	"time"

	"github.com/devafatek/WasteLibrary"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}
var socketCh chan string

func initStart() {

	WasteLibrary.LogStr("Successfully connected!")
	go WasteLibrary.InitLog()
	socketCh = make(chan string)
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
}

func main() {

	initStart()

	go getCustomer()

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
	resultVal.Result = WasteLibrary.RESULT_OK
	w.Write(resultVal.ToByte())

	readerType := req.FormValue(WasteLibrary.HTTP_READERTYPE)
	var nfcTypeVal WasteLibrary.NfcType = WasteLibrary.StringToNfcType(req.FormValue(WasteLibrary.HTTP_DATA))
	WasteLibrary.LogStr(readerType)

	if readerType == WasteLibrary.READERTYPE_RF {
		resultVal.Result = WasteLibrary.RECY_SOCKET_ANALYZE
		resultVal.Retval = nfcTypeVal.ToString()
		socketCh <- resultVal.ToString()
	} else if readerType == WasteLibrary.READERTYPE_WEBTRIGGER {
		go sendMotor()
		resultVal = getResult(nfcTypeVal)
		if resultVal.Result == WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RECY_SOCKET_FINISH
			socketCh <- resultVal.ToString()
		} else {
			resultVal.Result = WasteLibrary.RECY_SOCKET_ERROR
			socketCh <- resultVal.ToString()
		}

		time.Sleep(5 * time.Second)
		resultVal.Result = WasteLibrary.RECY_SOCKET_INDEX
		socketCh <- resultVal.ToString()
	} else {

	}

}

func getCustomer() WasteLibrary.ResultType {
	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	data := url.Values{
		WasteLibrary.HTTP_READERTYPE: {WasteLibrary.READERTYPE_GET_CUSTOMER},
	}
	resultVal = WasteLibrary.HttpPostReq("http://127.0.0.1:10000/trans", data)
	if resultVal.Result == WasteLibrary.RESULT_OK {
		//var customer WasteLibrary.CustomerType = WasteLibrary.StringToCustomerType(resultVal.Retval.(string))
		//var autoStartVal string
		//var customerAutoStartVal string
		//TO DO
		//create customerAutoStartVal by CustomerName
		//get file value
		//if autoStartVal != customerAutoStartVal {
		//TO DO
		// write customerAutoStartVal and reboot
		//}
	}
	return resultVal
}

func getResult(nfcData WasteLibrary.NfcType) WasteLibrary.ResultType {
	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	data := url.Values{
		WasteLibrary.HTTP_READERTYPE: {WasteLibrary.READERTYPE_GET_NFC},
		WasteLibrary.HTTP_DATA:       {nfcData.ToString()},
	}
	resultVal = WasteLibrary.HttpPostReq("http://127.0.0.1:10000/trans", data)
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
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
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
