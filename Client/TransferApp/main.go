package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"os/user"
	"strings"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var debug bool = os.Getenv("DEBUG") == "1"
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

	if !debug {
		time.Sleep(60 * time.Second)
	}
	logStr("Successfully connected!")
	currentUser = getCurrentUser()
	serialNumber = getSerialNumber()
	logStr(currentUser)
	logStr(serialNumber)
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

	if err := req.ParseForm(); err != nil {
		logErr(err)
		return
	}
	opType := req.FormValue("OPTYPE")
	logStr(opType)

	if opType == "APP" {
		w.Write([]byte("OK"))
	} else {
		w.Write([]byte("FAIL"))
	}
}

func trans(w http.ResponseWriter, req *http.Request) {

	if err := req.ParseForm(); err != nil {
		logErr(err)
		return
	}
	opType := req.FormValue("OPTYPE")
	logStr(opType)

	w.Write([]byte("OK"))
	dataMap := req.Form
	sendDataByte, err := json.Marshal(dataMap)
	if err != nil {
		logErr(err)
	}
	sendDataJson := string(sendDataByte)
	retVal := sendDataToServer(opType, sendDataJson, getTime(), "0")
	logStr("Send Data To Server : " + retVal)
	if retVal != "OK" {
		if opType != "CAM" {
			storeData(opType, sendDataJson)
		}
	}
	if opType == "CAM" {
		sendFileToServer(req.FormValue("UID"))
	}
}

func sendFileToServer(fileName string) {
	session, err := session.NewSession(&aws.Config{Region: aws.String(AWS_S3_REGION)})
	if err != nil {
		logErr(err)
	} else {
		err = uploadFile(session, "WAIT_CAM/"+fileName+".mp4")
		if err != nil {
			logErr(err)
		} else {
			removeFile("WAIT_CAM/" + fileName + ".mp4")
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

func sendDataToServer(datatype string, sendData string, dataTime string, repeat string) string {
	var retVal string = "FAIL"
	data := url.Values{
		"APPTYPE":  {applicationType},
		"DID":      {serialNumber},
		"DATATYPE": {datatype},
		"TIME":     {dataTime},
		"DATA":     {sendData},
		"REPEAT":   {repeat},
	}
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.PostForm("http://aws.afatek.com.tr/data", data)
	if err != nil {
		logErr(err)

	} else {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logErr(err)
		}
		bodyString := string(bodyBytes)
		if bodyString == "OK" {
			retVal = "OK"
		}
		logStr(bodyString)
	}
	return retVal
}

func storeData(dataType string, sendData string) {
	err := ioutil.WriteFile("WAIT_"+dataType+"/"+getTime(), []byte(sendData), 0644)
	if err != nil {
		logErr(err)
	}
}

func resendData(opType string, fileName string) {
	if opType == "CAM" {
		sendFileToServer(fileName)
	} else {

		var dataJSON string = ""
		var dataTime string = fileName

		readByte, err := ioutil.ReadFile("WAIT_" + opType + "/" + fileName)
		if err != nil {
			logErr(err)
		} else {

			dataJSON = string(readByte)

			logStr("Read File : " + dataJSON)

			retVal := sendDataToServer(opType, string(dataJSON), dataTime, "1")
			logStr("Send Data To Server Again : " + retVal)
			if retVal == "OK" {
				removeFile("WAIT_" + opType + "/" + fileName)
			}
		}
	}
}

func fileCheck(opType string) {
	logStr("File Check :" + opType)
	for {
		time.Sleep(opInterval * time.Second)

		f, err := os.Open("WAIT_" + opType)
		if err != nil {
			logErr(err)
			continue
		}
		fileInfo, err := f.Readdir(-1)
		f.Close()
		if err != nil {
			logErr(err)
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
				removeFile("WAIT_" + opType + "/" + file.Name())
			}
		}
	}

	wg.Done()
}

func removeFile(filePath string) {
	logStr("Remove File : " + filePath)
	cmdRm := exec.Command("rm", filePath)
	errRm := cmdRm.Start()
	if errRm != nil {
		logErr(errRm)
	}
}

func getCurrentUser() string {
	user, err := user.Current()
	if err != nil {
		logErr(err)
	}

	username := user.Username
	return username
}

func getSerialNumber() string {
	var tempNumber string = ""
	out, err := exec.Command("/home/pi/getSerialNumber.sh").Output()
	if err != nil {
		logErr(err)
	}
	tempNumber = strings.TrimSuffix(string(out), "\n")

	return tempNumber
}

func logErr(err error) {
	if err != nil {
		logStr(err.Error())
	}
}

func logStr(value string) {
	if debug {
		fmt.Println(value)
	}
}

func getTime() string {
	currentTime := time.Now()
	return currentTime.Format("2006.01.02 15:04:05")
}
