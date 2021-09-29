package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/AfatekDevelopers/result_lib_go/devafatekresult"
	"github.com/devafatek/WasteLibrary"
)

func main() {
	var resultVal devafatekresult.ResultType
	var opType string = "TEST"
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
			AdminLink:    "admin.aws.afatek.com.tr",
			WebLink:      "web.aws.afatek.com.tr",
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
	} else {

		var currentData WasteLibrary.DeviceType = WasteLibrary.StringToDeviceType("/gLX/4EDAQEKRGV2aWNlVHlwZQH/ggABKAEIRGV2aWNlSWQBCAABCkN1c3RvbWVySWQBCAABCkRldmljZU5hbWUBDAABCkRldmljZVR5cGUBDAABDFNlcmlhbE51bWJlcgEMAAEPUmVhZGVyQXBwU3RhdHVzAQwAARNSZWFkZXJBcHBMYXN0T2tUaW1lAQwAARBSZWFkZXJDb25uU3RhdHVzAQwAARRSZWFkZXJDb25uTGFzdE9rVGltZQEMAAEMUmVhZGVyU3RhdHVzAQwAARBSZWFkZXJMYXN0T2tUaW1lAQwAAQxDYW1BcHBTdGF0dXMBDAABEENhbUFwcExhc3RPa1RpbWUBDAABDUNhbUNvbm5TdGF0dXMBDAABEUNhbUNvbm5MYXN0T2tUaW1lAQwAAQlDYW1TdGF0dXMBDAABDUNhbUxhc3RPa1RpbWUBDAABDEdwc0FwcFN0YXR1cwEMAAEQR3BzQXBwTGFzdE9rVGltZQEMAAENR3BzQ29ublN0YXR1cwEMAAERR3BzQ29ubkxhc3RPa1RpbWUBDAABCUdwc1N0YXR1cwEMAAENR3BzTGFzdE9rVGltZQEMAAEOVGhlcm1BcHBTdGF0dXMBDAABElRoZXJtQXBwTGFzdE9rVGltZQEMAAERVHJhbnNmZXJBcHBTdGF0dXMBDAABFVRyYW5zZmVyQXBwTGFzdE9rVGltZQEMAAELQWxpdmVTdGF0dXMBDAABD0FsaXZlTGFzdE9rVGltZQEMAAENQ29udGFjdFN0YXR1cwEMAAERQ29udGFjdExhc3RPa1RpbWUBDAABBVRoZXJtAQwAAQhMYXRpdHVkZQEIAAEJTG9uZ2l0dWRlAQgAAQVTcGVlZAEIAAEGQWN0aXZlAQwAAQlUaGVybVRpbWUBDAABB0dwc1RpbWUBDAABClN0YXR1c1RpbWUBDAABCkNyZWF0ZVRpbWUBDAAAAAf/ggH+FEAA")
		fmt.Println(currentData)
	}

}
