package WasteLibrary

import (
	"fmt"
	"net/http"
	"net/url"
	"runtime"
	"time"
)

//LogErr
func LogErr(err error) {
	if err != nil {
		SendLogServer(ERROR, err.Error())
	}
}

//LogStr
func LogStr(value string) {
	if Debug {
		SendLogServer(INFO, value)
	}
}

//SendLogServer
func SendLogServer(logType string, logVal string) {
	if Container == "" {
		if Debug {
			fmt.Println("Time : " + time.Now().String() + " - LogType : " + logType + " - Func : " + GetFuncName(2).Function + " - Log : " + logVal)
		}
	} else {
		data := url.Values{
			CONTAINER: {Container},
			LOGTYPE:   {logType},
			FUNC:      {GetFuncName(2).Function},
			LOG:       {logVal},
		}
		client := http.Client{
			Timeout: 10 * time.Second,
		}
		client.PostForm("http://waste-logserver-cluster-ip/log", data)
	}
}

//GetFuncName
func GetFuncName(skipFrames int) runtime.Frame {
	targetFrameIndex := skipFrames + 2
	programCounters := make([]uintptr, targetFrameIndex+2)
	n := runtime.Callers(0, programCounters)

	frame := runtime.Frame{Function: "unknown"}
	if n > 0 {
		frames := runtime.CallersFrames(programCounters[:n])
		for more, frameIndex := true, 0; more && frameIndex <= targetFrameIndex; frameIndex++ {
			var frameCandidate runtime.Frame
			frameCandidate, more = frames.Next()
			if frameIndex == targetFrameIndex {
				frame = frameCandidate
			}
		}
	}

	return frame
}
