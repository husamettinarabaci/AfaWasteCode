package main

import (
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/devafatek/WasteLibrary"
)

var serialNumber = "0"
var currentUser string
var opInterval time.Duration = 24 * 60 * 60
var wg sync.WaitGroup

const (
	AWS_S3_REGION = "eu-central-1"
	AWS_S3_BUCKET = "afatek-waste-files-s3"
)

type statusType struct {
	Name string `json:"Name"`
	Key  string `json:"Key"`
	Port string `json:"Port"`
}

var statusTypes []statusType = []statusType{
	{
		Name: WasteLibrary.RFID_APPNAME_UPDATER,
		Key:  WasteLibrary.CHECKTYPE_APP,
		Port: "10005",
	},
}

func initStart() {

	time.Sleep(5 * time.Second)
	WasteLibrary.LogStr("Successfully connected!")
	WasteLibrary.Version = "2"
	WasteLibrary.LogStr("Version : " + WasteLibrary.Version)
	currentUser = WasteLibrary.GetCurrentUser()
	serialNumber = getSerialNumber()
	WasteLibrary.LogStr(currentUser)
	WasteLibrary.LogStr(serialNumber)
}
func main() {

	initStart()

	go updateCheck()
	wg.Add(1)

	http.HandleFunc("/status", WasteLibrary.StatusHandler)
	http.ListenAndServe(":10006", nil)
	wg.Wait()

}

func updateCheck() {
	if currentUser == "pi" {
		for {
			time.Sleep(time.Second)
			startUpdate(WasteLibrary.RFID_APPNAME_UPDATER)

			time.Sleep(opInterval * time.Second)
		}
	}
	wg.Done()
}

func startUpdate(appType string) {
	var resultVal WasteLibrary.ResultType
	var currentHttpHeader WasteLibrary.HttpClientHeaderType
	currentHttpHeader.New()
	currentHttpHeader.DeviceNo = serialNumber
	currentHttpHeader.ReaderType = WasteLibrary.READERTYPE_UPDATE
	currentHttpHeader.DeviceType = WasteLibrary.DEVICETYPE_RFID
	var updaterType WasteLibrary.UpdaterType
	updaterType.New()
	updaterType.AppType = appType

	resultVal = checkApp(updaterType)
	if resultVal.Result == WasteLibrary.RESULT_OK {
		if resultVal.Retval != nil {
			updaterType.Version = resultVal.Retval.(string)
		}
	}

	data := url.Values{
		WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
		WasteLibrary.HTTP_DATA:   {updaterType.ToString()},
	}
	resultVal = WasteLibrary.HttpPostReq("http://listener.aws.afatek.com.tr/update", data)
	if resultVal.Result == WasteLibrary.RESULT_OK {
		updaterType.StringToType(resultVal.Retval.(string))
		if downloadApp(updaterType).Result == WasteLibrary.RESULT_OK {
			if stopApp(updaterType).Result == WasteLibrary.RESULT_OK {
				if rmApp(updaterType).Result == WasteLibrary.RESULT_OK {
					if mvApp(updaterType).Result == WasteLibrary.RESULT_OK {
						if chownApp(updaterType).Result == WasteLibrary.RESULT_OK {
							if chmodApp(updaterType).Result == WasteLibrary.RESULT_OK {
								if startApp(updaterType).Result == WasteLibrary.RESULT_OK {
									if checkApp(updaterType).Result == WasteLibrary.RESULT_OK {
										updateVersion(updaterType)
									} else {
									}
								} else {
								}
							} else {
							}
						} else {
						}
					} else {
					}
				} else {
				}
			} else {
			}
		} else {
		}
	}

}

func downloadApp(updaterType WasteLibrary.UpdaterType) WasteLibrary.ResultType {
	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL

	if WasteLibrary.IsFileExists("/home/pi/DOWNLOADED_APP/" + updaterType.AppType) {
		rmAppDownloaded(updaterType)
	}

	downFile := "FIRMWARE/RFID/" + updaterType.Version + "/" + updaterType.AppType
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String(AWS_S3_REGION)},
	)

	downloader := s3manager.NewDownloader(sess)

	file, err := os.Create("/home/pi/DOWNLOADED_APP/" + updaterType.AppType)
	if err != nil {
		WasteLibrary.LogErr(err)
	} else {
		_, err := downloader.Download(file,
			&s3.GetObjectInput{
				Bucket: aws.String(AWS_S3_BUCKET),
				Key:    aws.String(downFile),
			})

		if err != nil {
			WasteLibrary.LogErr(err)
		} else {
			if WasteLibrary.IsFileExists("/home/pi/DOWNLOADED_APP/" + updaterType.AppType) {
				resultVal.Result = WasteLibrary.RESULT_OK
				resultVal.Retval = WasteLibrary.RESULT_SUCCESS_OK
			}
		}
	}

	time.Sleep(time.Second * 2)
	return resultVal
}

func stopApp(updaterType WasteLibrary.UpdaterType) WasteLibrary.ResultType {
	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	cmd := exec.Command("sudo", "systemctl", "stop", updaterType.AppType+".service")
	err := cmd.Start()
	time.Sleep(time.Second * 5)
	if err != nil {
		WasteLibrary.LogErr(err)
		resultVal.Result = WasteLibrary.RESULT_FAIL
	} else {
		resultVal.Result = WasteLibrary.RESULT_OK
		resultVal.Retval = WasteLibrary.RESULT_SUCCESS_OK
	}
	time.Sleep(time.Second * 2)
	return resultVal
}

func rmAppDownloaded(updaterType WasteLibrary.UpdaterType) WasteLibrary.ResultType {
	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	cmd := exec.Command("sudo", "rm", "/home/pi/DOWNLOADED_APP/"+updaterType.AppType)
	err := cmd.Start()
	time.Sleep(time.Second * 5)
	if err != nil {
		WasteLibrary.LogErr(err)
		resultVal.Result = WasteLibrary.RESULT_FAIL
	} else {
		resultVal.Result = WasteLibrary.RESULT_OK
		resultVal.Retval = WasteLibrary.RESULT_SUCCESS_OK
	}
	time.Sleep(time.Second * 2)
	return resultVal
}

func rmApp(updaterType WasteLibrary.UpdaterType) WasteLibrary.ResultType {
	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	if !WasteLibrary.IsFileExists("/home/pi/" + updaterType.AppType) {
		resultVal.Result = WasteLibrary.RESULT_OK
		resultVal.Retval = WasteLibrary.RESULT_SUCCESS_OK
	} else {
		cmd := exec.Command("sudo", "rm", "/home/pi/"+updaterType.AppType)
		err := cmd.Start()
		time.Sleep(time.Second * 5)
		if err != nil {
			WasteLibrary.LogErr(err)
			resultVal.Result = WasteLibrary.RESULT_FAIL
		} else {
			if !WasteLibrary.IsFileExists("/home/pi/" + updaterType.AppType) {
				resultVal.Result = WasteLibrary.RESULT_OK
				resultVal.Retval = WasteLibrary.RESULT_SUCCESS_OK
			}
		}
	}
	time.Sleep(time.Second * 2)
	return resultVal
}

func mvApp(updaterType WasteLibrary.UpdaterType) WasteLibrary.ResultType {
	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	cmd := exec.Command("sudo", "cp", "/home/pi/DOWNLOADED_APP/"+updaterType.AppType, "/home/pi/"+updaterType.AppType)
	err := cmd.Start()
	time.Sleep(time.Second * 5)
	if err != nil {
		WasteLibrary.LogErr(err)
		resultVal.Result = WasteLibrary.RESULT_FAIL
	} else {
		if WasteLibrary.IsFileExists("/home/pi/" + updaterType.AppType) {
			resultVal.Result = WasteLibrary.RESULT_OK
			resultVal.Retval = WasteLibrary.RESULT_SUCCESS_OK
		}
	}
	time.Sleep(time.Second * 2)
	return resultVal
}

func chownApp(updaterType WasteLibrary.UpdaterType) WasteLibrary.ResultType {
	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	cmd := exec.Command("sudo", "chown", "pi:pi", "/home/pi/"+updaterType.AppType)
	err := cmd.Start()
	time.Sleep(time.Second * 5)
	if err != nil {
		WasteLibrary.LogErr(err)
		resultVal.Result = WasteLibrary.RESULT_FAIL
	} else {
		resultVal.Result = WasteLibrary.RESULT_OK
		resultVal.Retval = WasteLibrary.RESULT_SUCCESS_OK
	}
	time.Sleep(time.Second * 2)
	return resultVal
}

func chmodApp(updaterType WasteLibrary.UpdaterType) WasteLibrary.ResultType {
	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	cmd := exec.Command("sudo", "chmod", "+x", "/home/pi/"+updaterType.AppType)
	err := cmd.Start()
	time.Sleep(time.Second * 5)
	if err != nil {
		WasteLibrary.LogErr(err)
		resultVal.Result = WasteLibrary.RESULT_FAIL
	} else {
		resultVal.Result = WasteLibrary.RESULT_OK
		resultVal.Retval = WasteLibrary.RESULT_SUCCESS_OK
	}
	time.Sleep(time.Second * 2)
	return resultVal
}

func startApp(updaterType WasteLibrary.UpdaterType) WasteLibrary.ResultType {
	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	cmd := exec.Command("sudo", "systemctl", "start", updaterType.AppType+".service")
	err := cmd.Start()
	time.Sleep(time.Second * 5)
	if err != nil {
		WasteLibrary.LogErr(err)
		resultVal.Result = WasteLibrary.RESULT_FAIL
	} else {
		resultVal.Result = WasteLibrary.RESULT_OK
		resultVal.Retval = WasteLibrary.RESULT_SUCCESS_OK
	}
	time.Sleep(time.Second * 10)

	time.Sleep(time.Second * 2)
	return resultVal
}

func checkApp(updaterType WasteLibrary.UpdaterType) WasteLibrary.ResultType {
	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	for i := range statusTypes {
		if statusTypes[i].Name == updaterType.AppType {
			data := url.Values{
				WasteLibrary.HTTP_CHECKTYPE: {statusTypes[i].Key},
			}
			resultVal = WasteLibrary.HttpPostReq("http://127.0.0.1:"+statusTypes[i].Port+"/status", data)
			break
		}
	}

	time.Sleep(time.Second * 2)
	return resultVal
}

func updateVersion(updaterType WasteLibrary.UpdaterType) WasteLibrary.ResultType {
	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	var currentHttpHeader WasteLibrary.HttpClientHeaderType
	currentHttpHeader.New()
	currentHttpHeader.DeviceNo = serialNumber
	currentHttpHeader.ReaderType = WasteLibrary.READERTYPE_UPDATE
	currentHttpHeader.DeviceType = WasteLibrary.DEVICETYPE_RFID
	data := url.Values{
		WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
		WasteLibrary.HTTP_DATA:   {updaterType.ToString()},
	}
	resultVal = WasteLibrary.HttpPostReq("http://listener.aws.afatek.com.tr/update", data)
	return resultVal
}

func getSerialNumber() string {
	var tempNumber string = ""
	out, err := exec.Command("/home/pi/getSerialNumber.sh").Output()
	if err != nil {
		WasteLibrary.LogErr(err)
	}
	tempNumber = strings.TrimSuffix(string(out), "\n")

	return tempNumber
}
