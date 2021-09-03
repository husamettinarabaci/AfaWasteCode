package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/AfatekDevelopers/result_lib_go/devafatekresult"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/devafatek/WasteLibrary"
)

var applicationType = "RFID"
var serialNumber = "0"
var currentUser string
var opInterval time.Duration = 5 * 60
var wg sync.WaitGroup

const (
	AWS_S3_REGION = "eu-central-1"
	AWS_S3_BUCKET = "afatek-waste-videos-s3"
)

func initStart() {

	time.Sleep(5 * time.Second)
	WasteLibrary.LogStr("Successfully connected!")
	currentUser = WasteLibrary.GetCurrentUser()
	serialNumber = getSerialNumber()
	WasteLibrary.LogStr(currentUser)
	WasteLibrary.LogStr(serialNumber)
}
func main() {

	initStart()

	time.Sleep(time.Second)
	go fileCheck("RF")
	wg.Add(1)

	time.Sleep(time.Second)
	go fileCheck("CAM")
	wg.Add(1)

	time.Sleep(time.Second)
	go fileCheck("GPS")
	wg.Add(1)

	time.Sleep(time.Second)
	go fileCheck("THERM")
	wg.Add(1)

	time.Sleep(time.Second)
	go fileCheck("STATUS")
	wg.Add(1)

	http.HandleFunc("/status", status)
	http.HandleFunc("/trans", trans)
	http.ListenAndServe(":10000", nil)

	wg.Wait()

}

func status(w http.ResponseWriter, req *http.Request) {
	var resultVal devafatekresult.ResultType

	if err := req.ParseForm(); err != nil {
		WasteLibrary.LogErr(err)
		return
	}
	opType := req.FormValue("OPTYPE")
	WasteLibrary.LogStr(opType)

	resultVal.Result = "FAIL"
	if opType == "APP" {
		resultVal.Result = "OK"
	} else {
		resultVal.Result = "FAIL"
	}
	w.Write(resultVal.ToByte())
}

func trans(w http.ResponseWriter, req *http.Request) {
	var resultVal devafatekresult.ResultType

	if err := req.ParseForm(); err != nil {
		WasteLibrary.LogErr(err)
		resultVal.Result = "FAIL"

	} else {

		opType := req.FormValue("OPTYPE")
		dataVal := req.FormValue("DATA")
		resultVal = sendDataToServer(opType, dataVal, WasteLibrary.GetTime(), "0")
		WasteLibrary.LogStr("Send Data To Server : " + resultVal.ToString())
		if resultVal.Result != "OK" {
			if opType != "CAM" {
				storeData(opType, dataVal)
			}
		}
		if opType == "CAM" {

			sendFileToServer(req.FormValue("UID"))
		}
		resultVal.Result = "OK"
	}
	w.Write(resultVal.ToByte())
}

func sendFileToServer(fileName string) {
	session, err := session.NewSession(&aws.Config{Region: aws.String(AWS_S3_REGION)})
	if err != nil {
		WasteLibrary.LogErr(err)
	} else {
		err = uploadFile(session, "WAIT_CAM/"+fileName+".mp4")
		if err != nil {
			WasteLibrary.LogErr(err)
		} else {
			WasteLibrary.RemoveFile("WAIT_CAM/" + fileName + ".mp4")
		}
	}
}

func uploadFile(session *session.Session, uploadFileDir string) error {

	upFile, err := os.Open(uploadFileDir)
	if err != nil {
		return err
	}
	defer upFile.Close()

	upFileInfo, _ := upFile.Stat()
	var fileSize int64 = upFileInfo.Size()
	fileBuffer := make([]byte, fileSize)
	upFile.Read(fileBuffer)

	_, err = s3.New(session).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(AWS_S3_BUCKET),
		Key:                  aws.String(uploadFileDir),
		ACL:                  aws.String("private"),
		Body:                 bytes.NewReader(fileBuffer),
		ContentLength:        aws.Int64(fileSize),
		ContentType:          aws.String(http.DetectContentType(fileBuffer)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
	})
	return err
}

func sendDataToServer(datatype string, sendData string, dataTime string, repeat string) devafatekresult.ResultType {
	var resultVal devafatekresult.ResultType
	data := url.Values{
		"APPTYPE":  {applicationType},
		"DID":      {serialNumber},
		"DATATYPE": {datatype},
		"TIME":     {dataTime},
		"DATA":     {sendData},
		"REPEAT":   {repeat},
	}
	resultVal = WasteLibrary.HttpPostReq("http://aws.afatek.com.tr/data", data)
	return resultVal
}

func storeData(dataType string, sendData string) {
	err := ioutil.WriteFile("WAIT_"+dataType+"/"+WasteLibrary.GetTime(), []byte(sendData), 0644)
	if err != nil {
		WasteLibrary.LogErr(err)
	}
}

func resendData(opType string, fileName string) {
	var resultVal devafatekresult.ResultType
	if opType == "CAM" {
		sendFileToServer(fileName)
	} else {

		var dataJSON string = ""
		var dataTime string = fileName

		readByte, err := ioutil.ReadFile("WAIT_" + opType + "/" + fileName)
		if err != nil {
			WasteLibrary.LogErr(err)
		} else {

			dataJSON = string(readByte)

			WasteLibrary.LogStr("Read File : " + dataJSON)

			resultVal = sendDataToServer(opType, string(dataJSON), dataTime, "1")
			WasteLibrary.LogStr("Send Data To Server Again : " + resultVal.ToString())
			if resultVal.Result == "OK" {
				WasteLibrary.RemoveFile("WAIT_" + opType + "/" + fileName)
			}
		}
	}
}

func fileCheck(opType string) {
	WasteLibrary.LogStr("File Check :" + opType)
	for {
		time.Sleep(opInterval * time.Second)

		f, err := os.Open("WAIT_" + opType)
		if err != nil {
			WasteLibrary.LogErr(err)
			continue
		}
		fileInfo, err := f.Readdir(-1)
		f.Close()
		if err != nil {
			WasteLibrary.LogErr(err)
			continue
		}

		for _, file := range fileInfo {
			time.Sleep(time.Second)
			second := time.Since(file.ModTime()).Seconds()
			if second > 60*60 && second < 24*60*60 {
				var fileName string = file.Name()
				if opType == "CAM" {
					spData := strings.Split(strings.TrimSpace(file.Name()), ".")
					fileName = spData[0]
				}
				resendData(opType, fileName)
			}
			if second > 24*60*60 {
				WasteLibrary.RemoveFile("WAIT_" + opType + "/" + file.Name())
			}
		}
	}

	wg.Done()
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
