package main

import (
	"net/http"
	"net/url"

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

	//TO DO
	// Get Customer LocalConfig

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
	WasteLibrary.LogStr(readerType)

	if readerType == WasteLibrary.READERTYPE_RF {
		socketCh <- WasteLibrary.RECY_SOCKET_ANALYZE
	} else if readerType == WasteLibrary.READERTYPE_WEBTRIGGER {
		socketCh <- WasteLibrary.RECY_SOCKET_FINISH
		sendMotor()
	} else {

	}
	socketCh <- WasteLibrary.RECY_SOCKET_INDEX

}

func sendMotor() {

	data := url.Values{
		WasteLibrary.HTTP_READERTYPE: {WasteLibrary.READERTYPE_MOTORRIGGER},
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
			break
		}
	}
}
