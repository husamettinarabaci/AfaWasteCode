package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/AfatekDevelopers/result_lib_go/devafatekresult"
	"github.com/devafatek/WasteLibrary"
)

func main() {
	var resultVal WasteLibrary.ResultType
	var opType string = "NO"
	WasteLibrary.Debug = true
	if opType == "CUSTOMER" {
		var currentHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.HttpClientHeaderType{
			AppType:      "ADMIN",
			DeviceNo:     "",
			OpType:       "CUSTOMER",
			Time:         WasteLibrary.GetTime(),
			Repeat:       "0",
			DeviceId:     0,
			CustomerId:   0,
			BaseDataType: "CUSTOMER",
		}

		var customerId int = 0
		var currentCustomer WasteLibrary.CustomerType = WasteLibrary.CustomerType{
			CustomerId:   float64(customerId),
			CustomerName: "Afatek",
			CustomerLink: "afatek.aws.afatek.com.tr",
			RfIdApp:      "1",
			UltApp:       "0",
			RecyApp:      "1",
		}

		data := url.Values{
			"HEADER": {currentHttpHeader.ToString()},
			"DATA":   {currentCustomer.ToString()},
		}

		client := http.Client{
			Timeout: 10 * time.Second,
		}
		resp, err := client.PostForm("http://a579dddf21ea745a49c7237860760244-1808420299.eu-central-1.elb.amazonaws.com/setCustomer", data)

		if err != nil {
			WasteLibrary.LogErr(err)

		} else {
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				WasteLibrary.LogErr(err)
			}
			resultVal = devafatekresult.ByteToResultType(bodyBytes)
			WasteLibrary.LogStr(resultVal.ToString())
		}
	} else if opType == "DEVICE" {
		var currentHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.HttpClientHeaderType{
			AppType:      "ADMIN",
			DeviceNo:     "",
			OpType:       "DEVICE",
			Time:         WasteLibrary.GetTime(),
			Repeat:       "0",
			DeviceId:     0,
			CustomerId:   0,
			BaseDataType: "DEVICE",
		}

		var deviceId int = 0
		var customerId int = 1
		var currentDevice WasteLibrary.DeviceType = WasteLibrary.DeviceType{
			DeviceId:     float64(deviceId),
			CustomerId:   float64(customerId),
			DeviceName:   "07 MVS 33",
			DeviceType:   "RFID",
			SerialNumber: "00000000c1b1d188",
		}
		WasteLibrary.LogStr(currentDevice.ToString())
		data := url.Values{
			"HEADER": {currentHttpHeader.ToString()},
			"DATA":   {currentDevice.ToString()},
		}

		client := http.Client{
			Timeout: 10 * time.Second,
		}
		resp, err := client.PostForm("http://a579dddf21ea745a49c7237860760244-1808420299.eu-central-1.elb.amazonaws.com/setDevice", data)

		if err != nil {
			WasteLibrary.LogErr(err)

		} else {
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				WasteLibrary.LogErr(err)
			}
			resultVal = devafatekresult.ByteToResultType(bodyBytes)
			WasteLibrary.LogStr(resultVal.ToString())
		}
	} else if opType == "CUSTOMERCONFIG" {
		var currentHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.HttpClientHeaderType{
			AppType:      "ADMIN",
			DeviceNo:     "",
			OpType:       "CUSTOMER",
			Time:         WasteLibrary.GetTime(),
			Repeat:       "0",
			DeviceId:     0,
			CustomerId:   0,
			BaseDataType: "DEVICE",
		}

		var customerId int = 1
		var currentData WasteLibrary.CustomerConfigType = WasteLibrary.CustomerConfigType{
			CustomerId:      float64(customerId),
			ArventoApp:      "1",
			ArventoUserName: "afatekbilisim",
			ArventoPin1:     "Amca151200!Furkan",
			ArventoPin2:     "Amca151200!Furkan",
			SystemProblem:   "1",
			TruckStopTrace:  "1",
			Active:          "1",
			CreateTime:      WasteLibrary.GetTime(),
		}
		WasteLibrary.LogStr(currentData.ToString())
		data := url.Values{
			"HEADER": {currentHttpHeader.ToString()},
			"DATA":   {currentData.ToString()},
		}

		client := http.Client{
			Timeout: 10 * time.Second,
		}
		resp, err := client.PostForm("http://a579dddf21ea745a49c7237860760244-1808420299.eu-central-1.elb.amazonaws.com/setConfig", data)

		if err != nil {
			WasteLibrary.LogErr(err)

		} else {
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				WasteLibrary.LogErr(err)
			}
			resultVal = devafatekresult.ByteToResultType(bodyBytes)
			WasteLibrary.LogStr(resultVal.ToString())
		}
	} else if opType == "INGRESS" {

		data := url.Values{}

		client := http.Client{
			Timeout: 10 * time.Second,
		}
		resp, err := client.PostForm("http://a36dca1ff41a94299939de76fcf7f5e3-535996505.eu-central-1.elb.amazonaws.com/status1", data)

		if err != nil {
			WasteLibrary.LogErr(err)

		} else {
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				WasteLibrary.LogErr(err)
			}
			resultVal = devafatekresult.ByteToResultType(bodyBytes)
			WasteLibrary.LogStr(resultVal.ToString())
		}
	} else if opType == "TEST" {

		data := url.Values{}

		client := http.Client{
			Timeout: 10 * time.Second,
		}
		resp, err := client.PostForm("http://aa594028d9f43457ea027f6f0928e021-567581812.eu-central-1.elb.amazonaws.com/setConfig", data)

		if err != nil {
			WasteLibrary.LogErr(err)

		} else {
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				WasteLibrary.LogErr(err)
			}
			resultVal = devafatekresult.ByteToResultType(bodyBytes)
			WasteLibrary.LogStr(resultVal.ToString())
		}
	} else if opType == "TCP" {

		strEcho := "ULT#SENS#864450040594790#40#3241#E32.12345#N032.45678#10341#09872#11234#10341#09872#11234#10341#09872#11234#11234"
		servAddr := "listener.aws.afatek.com.tr:20000"
		tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
		if err != nil {
			println("ResolveTCPAddr failed:", err.Error())
			os.Exit(1)
		}

		conn, err := net.DialTCP("tcp", nil, tcpAddr)
		if err != nil {
			println("Dial failed:", err.Error())
			os.Exit(1)
		}

		_, err = conn.Write([]byte(strEcho))
		if err != nil {
			println("Write to server failed:", err.Error())
			os.Exit(1)
		}

		println("write to server = ", strEcho)

		reply := make([]byte, 1024)

		_, err = conn.Read(reply)
		if err != nil {
			println("Write to server failed:", err.Error())
			os.Exit(1)
		}

		println("reply from server=", string(reply))

		conn.Close()
	} else {
		var test1 TestType = TestType{
			TagID:  45,
			Status: WasteLibrary.OK,
			Tags:   make(map[string]WasteLibrary.TagType),
		}

		var tag0 WasteLibrary.TagType = WasteLibrary.TagType{
			TagID: 1,
		}
		var tag1 WasteLibrary.TagType = WasteLibrary.TagType{
			TagID: 2,
		}

		test1.Tags[tag0.ToIdString()] = tag0
		test1.Tags[tag1.ToIdString()] = tag1

		fmt.Println(test1)
		fmt.Println(test1.ToString())

		var test2 = StringToTestType(test1.ToString())

		fmt.Println(test2)
		fmt.Println(test2.ToString())

		var conval = WasteLibrary.CONNECTED
		if conval == "CONNECTED" {
			fmt.Println("1")
		} else {
			fmt.Println("2")
		}

	}

}

type TestType struct {
	TagID  float64
	Status WasteLibrary.StatusType
	Tags   map[string]WasteLibrary.TagType
}

func (res TestType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.TagID)
}

func (res TestType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

func (res TestType) ToString() string {
	return string(res.ToByte())

}

func ByteToTestType(retByte []byte) TestType {
	var retVal TestType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

func StringToTestType(retStr string) TestType {
	return ByteToTestType([]byte(retStr))
}
