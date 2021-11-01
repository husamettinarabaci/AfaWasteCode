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

var applicationType = WasteLibrary.APPTYPE_RFID
var serialNumber = "0"
var currentUser string
var opInterval time.Duration = 60 * 60
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
		Name: WasteLibrary.RFID_APPTYPE_READER,
		Key:  WasteLibrary.OPTYPE_APP,
		Port: "10001",
	},
	{
		Name: WasteLibrary.RFID_APPTYPE_CAM,
		Key:  WasteLibrary.OPTYPE_APP,
		Port: "10002",
	},
	{
		Name: WasteLibrary.RFID_APPTYPE_GPS,
		Key:  WasteLibrary.OPTYPE_APP,
		Port: "10003",
	},
	{
		Name: WasteLibrary.RFID_APPTYPE_THERM,
		Key:  WasteLibrary.OPTYPE_APP,
		Port: "10004",
	},
	{
		Name: WasteLibrary.RFID_APPTYPE_SYSTEM,
		Key:  WasteLibrary.OPTYPE_APP,
		Port: "10006",
	},
	{
		Name: WasteLibrary.RFID_APPTYPE_TRANSFER,
		Key:  WasteLibrary.OPTYPE_APP,
		Port: "10000",
	},
	{
		Name: WasteLibrary.RFID_APPTYPE_CHECKER,
		Key:  WasteLibrary.OPTYPE_APP,
		Port: "10007",
	},
}

func initStart() {

	time.Sleep(5 * time.Second)
	WasteLibrary.LogStr("Successfully connected!")
	WasteLibrary.Version = "1"
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
	http.ListenAndServe(":10005", nil)
	wg.Wait()

}

func updateCheck() {
	if currentUser == "pi" {
		for {
			time.Sleep(time.Second)
			startUpdate(WasteLibrary.RFID_APPTYPE_GPS)

			time.Sleep(time.Second)
			startUpdate(WasteLibrary.RFID_APPTYPE_CAM)

			time.Sleep(time.Second)
			startUpdate(WasteLibrary.RFID_APPTYPE_CHECKER)

			time.Sleep(time.Second)
			startUpdate(WasteLibrary.RFID_APPTYPE_READER)

			time.Sleep(time.Second)
			startUpdate(WasteLibrary.RFID_APPTYPE_THERM)

			time.Sleep(time.Second)
			startUpdate(WasteLibrary.RFID_APPTYPE_TRANSFER)

			time.Sleep(time.Second)
			startUpdate(WasteLibrary.RFID_APPTYPE_SYSTEM)

			time.Sleep(opInterval * time.Second)
		}
	}
	wg.Done()
}

func startUpdate(appType string) {
	var resultVal WasteLibrary.ResultType
	var currentHttpHeader WasteLibrary.HttpClientHeaderType
	currentHttpHeader.New()
	currentHttpHeader.AppType = applicationType
	currentHttpHeader.DeviceNo = serialNumber
	currentHttpHeader.OpType = WasteLibrary.OPTYPE_UPDATE
	currentHttpHeader.Time = WasteLibrary.GetTime()
	currentHttpHeader.Repeat = WasteLibrary.STATU_PASSIVE
	currentHttpHeader.DeviceId = 0
	currentHttpHeader.CustomerId = 0
	currentHttpHeader.DeviceType = WasteLibrary.DEVICE_TYPE_RFID
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
		updaterType = WasteLibrary.StringToUpdaterType(resultVal.Retval.(string))
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
										WasteLibrary.LogStr("Error : Check App - " + updaterType.AppType)
									}
								} else {
									WasteLibrary.LogStr("Error : Start App - " + updaterType.AppType)
								}
							} else {
								WasteLibrary.LogStr("Error : Chmod App - " + updaterType.AppType)
							}
						} else {
							WasteLibrary.LogStr("Error : Chown App - " + updaterType.AppType)
						}
					} else {
						WasteLibrary.LogStr("Error : Cp App - " + updaterType.AppType)
					}
				} else {
					WasteLibrary.LogStr("Error : Rm App - " + updaterType.AppType)
				}
			} else {
				WasteLibrary.LogStr("Error : Stop App - " + updaterType.AppType)
			}
		} else {
			WasteLibrary.LogStr("Error : Download App - " + updaterType.AppType)
		}
	}

}

func downloadApp(updaterType WasteLibrary.UpdaterType) WasteLibrary.ResultType {
	WasteLibrary.LogStr("Download App : " + updaterType.AppType + " - " + updaterType.Version)
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
			}
		}
	}

	time.Sleep(time.Second * 2)
	return resultVal
}

func stopApp(updaterType WasteLibrary.UpdaterType) WasteLibrary.ResultType {
	WasteLibrary.LogStr("Stop App : " + updaterType.AppType)
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
	}
	time.Sleep(time.Second * 2)
	return resultVal
}

func rmAppDownloaded(updaterType WasteLibrary.UpdaterType) WasteLibrary.ResultType {
	WasteLibrary.LogStr("Rm Downloaded App : " + updaterType.AppType)
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
	}
	time.Sleep(time.Second * 2)
	return resultVal
}

func rmApp(updaterType WasteLibrary.UpdaterType) WasteLibrary.ResultType {
	WasteLibrary.LogStr("Rm App : " + updaterType.AppType)
	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	if !WasteLibrary.IsFileExists("/home/pi/" + updaterType.AppType) {
		resultVal.Result = WasteLibrary.RESULT_OK
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
			}
		}
	}
	time.Sleep(time.Second * 2)
	return resultVal
}

func mvApp(updaterType WasteLibrary.UpdaterType) WasteLibrary.ResultType {
	WasteLibrary.LogStr("Cp App : " + updaterType.AppType)
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
		}
	}
	time.Sleep(time.Second * 2)
	return resultVal
}

func chownApp(updaterType WasteLibrary.UpdaterType) WasteLibrary.ResultType {
	WasteLibrary.LogStr("Chown App : " + updaterType.AppType)
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
	}
	time.Sleep(time.Second * 2)
	return resultVal
}

func chmodApp(updaterType WasteLibrary.UpdaterType) WasteLibrary.ResultType {
	WasteLibrary.LogStr("Chmod App : " + updaterType.AppType)
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
	}
	time.Sleep(time.Second * 2)
	return resultVal
}

func startApp(updaterType WasteLibrary.UpdaterType) WasteLibrary.ResultType {
	WasteLibrary.LogStr("Start App : " + updaterType.AppType)
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
	}
	time.Sleep(time.Second * 10)

	time.Sleep(time.Second * 2)
	return resultVal
}

func checkApp(updaterType WasteLibrary.UpdaterType) WasteLibrary.ResultType {
	WasteLibrary.LogStr("Check App : " + updaterType.AppType)
	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	for i := range statusTypes {
		if statusTypes[i].Name == updaterType.AppType {
			data := url.Values{
				WasteLibrary.HTTP_OPTYPE: {statusTypes[i].Key},
			}
			resultVal = WasteLibrary.HttpPostReq("http://127.0.0.1:"+statusTypes[i].Port+"/status", data)
			break
		}
	}

	time.Sleep(time.Second * 2)
	return resultVal
}

func updateVersion(updaterType WasteLibrary.UpdaterType) WasteLibrary.ResultType {
	WasteLibrary.LogStr("Update Version : " + updaterType.AppType + " - " + updaterType.Version)
	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	var currentHttpHeader WasteLibrary.HttpClientHeaderType
	currentHttpHeader.New()
	currentHttpHeader.AppType = applicationType
	currentHttpHeader.DeviceNo = serialNumber
	currentHttpHeader.OpType = WasteLibrary.OPTYPE_UPDATE
	currentHttpHeader.Time = WasteLibrary.GetTime()
	currentHttpHeader.Repeat = WasteLibrary.STATU_PASSIVE
	currentHttpHeader.DeviceId = 0
	currentHttpHeader.CustomerId = 0
	currentHttpHeader.DeviceType = WasteLibrary.DEVICE_TYPE_RFID
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
