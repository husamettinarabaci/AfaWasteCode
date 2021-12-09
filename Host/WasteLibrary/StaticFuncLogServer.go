package WasteLibrary

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"runtime"
	"time"

	"github.com/go-redis/redis/v8"
)

//LogErr
func LogErr(err error) {
	if err != nil {
		SendLogServer(LOGGER_ERROR, err.Error())
	}
}

//LogStr
func LogStr(value string) {
	if Debug {
		SendLogServer(LOGGER_INFO, value)
	}
}

//SendLogServer
func SendLogServer(logType string, logVal string) {
	if Container == "" {
		if Debug {
			fmt.Println("Time : " + GetTime() + " - LogType : " + logType + " - Func : " + GetFuncName(2).Function + " - Log : " + logVal)
		}
	} else {
		data := url.Values{
			LOGGER_CONTAINER: {Container},
			LOGGER_LOGTYPE:   {logType},
			LOGGER_FUNC:      {GetFuncName(2).Function},
			LOGGER_LOG:       {logVal},
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

//InitLog
func InitLog() {
	var ctxLog = context.Background()
	redisDbLog := redis.NewClient(&redis.Options{
		Addr:     "waste-redis-cluster-ip:6379",
		Password: "Amca151200!Furkan",
		DB:       0,
	})

	pong, err := redisDbLog.Ping(ctxLog).Result()
	LogErr(err)
	LogStr(pong)
	subscriber := redisDbLog.Subscribe(ctxLog, REDIS_APP_LOG_CHANNEL)
	var resultVal ResultType
	for {
		msg, err := subscriber.ReceiveMessage(ctxLog)
		if err != nil {
			continue
		}
		resultVal.StringToType(msg.Payload)
		if resultVal.Result == Container {
			LogStr("LogStatu : " + resultVal.Retval.(string) + " - Container : " + Container)
			Debug = resultVal.Retval.(string) == "open"
		}
	}
}
