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
		var customerId int = 10
		var currentDevice WasteLibrary.DeviceType = WasteLibrary.DeviceType{
			DeviceId:     float64(deviceId),
			CustomerId:   float64(customerId),
			DeviceName:   "06 AFA 01",
			DeviceType:   "RFID",
			SerialNumber: "00000000c1b1d188",
		}

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
	} else {
		var gopVal string = "/gGy/4MDAQEKRGV2aWNlVHlwZQH/hAABGgEIRGV2aWNlSWQBCAABCkN1c3RvbWVySWQBCAABCkRldmljZU5hbWUBDAABCkRldmljZVR5cGUBDAABDFNlcmlhbE51bWJlcgEMAAEPUmVhZGVyQXBwU3RhdHVzAQwAARBSZWFkZXJDb25uU3RhdHVzAQwAAQxSZWFkZXJTdGF0dXMBDAABDENhbUFwcFN0YXR1cwEMAAENQ2FtQ29ublN0YXR1cwEMAAEJQ2FtU3RhdHVzAQwAAQxHcHNBcHBTdGF0dXMBDAABDUdwc0Nvbm5TdGF0dXMBDAABCUdwc1N0YXR1cwEMAAEOVGhlcm1BcHBTdGF0dXMBDAABEVRyYW5zZmVyQXBwU3RhdHVzAQwAAQtBbGl2ZVN0YXR1cwEMAAENQ29udGFjdFN0YXR1cwEMAAEFVGhlcm0BDAABCExhdGl0dWRlAQwAAQlMb25naXR1ZGUBDAABBkFjdGl2ZQEMAAEJVGhlcm1UaW1lAQwAAQdHcHNUaW1lAQwAAQpTdGF0dXNUaW1lAQwAAQpDcmVhdGVUaW1lAQwAAAD/5/+EAf4gQAH+JEABCTA2IEFGQSAwMQEEUkZJRAEQMDAwMDAwMDBjMWIxZDE4OAEBMQEBMAEBMAEBMQEBMAEBMAEBMQEBMQEBMAEBMQEBMQEBMQEBMQEJdGVtcD0zOC42AQ0wLjAwMDAwMDAwMDAwAQ0wLjAwMDAwMDAwMDAwAQExARQyMDIxLTA5LTA2VDEyOjE3OjA3WgEbMjAyMS0wOS0wNFQyMTo1NTowNS4zOTU5MThaARQyMDIxLTA5LTA2VDEyOjE3OjA2WgEbMjAyMS0wOS0wNFQyMTo1NTowNS4zOTU5MThaAA=="

		var currentData WasteLibrary.DeviceType = WasteLibrary.StringToDeviceType(gopVal)
		fmt.Println(currentData)
	}

}
