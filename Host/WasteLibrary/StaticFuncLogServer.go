package WasteLibrary

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

//LogErr
func LogErr(err error) {
	if err != nil {
		SendLogServer("ERR", err.Error())
	}
}

//LogStr
func LogStr(value string) {
	if Debug {
		SendLogServer("INFO", value)
	}
}

//SendLogServer
func SendLogServer(logType string, logVal string) {
	if Container == "" {
		if Debug {
			fmt.Println("Time : " + time.Now().String() + " - LogType : " + logType + " - Log : " + logVal)
		}
	} else {
		data := url.Values{
			"CONTAINER": {Container},
			"LOGTYPE":   {logType},
			"LOG":       {logVal},
		}
		client := http.Client{
			Timeout: 10 * time.Second,
		}
		client.PostForm("http://waste-logserver-cluster-ip/log", data)
	}
}
