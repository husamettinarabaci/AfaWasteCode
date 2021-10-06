package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	/*var resultVal WasteLibrary.ResultType
	var opType string = "NO"
	WasteLibrary.Debug = true
	if opType == WasteLibrary.CUSTOMER {
		var currentHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.HttpClientHeaderType{
			AppType:      WasteLibrary.ADMIN,
			DeviceNo:     "",
			OpType:       WasteLibrary.CUSTOMER,
			Time:         WasteLibrary.GetTime(),
			Repeat:       WasteLibrary.STATU_PASSIVE,
			DeviceId:     0,
			CustomerId:   0,
			BaseDataType: WasteLibrary.CUSTOMER,
		}

		var customerId int = 0
		var currentCustomer WasteLibrary.CustomerType = WasteLibrary.CustomerType{
			CustomerId:   float64(customerId),
			CustomerName: "Afatek",
			CustomerLink: "afatek.aws.afatek.com.tr",
			RfIdApp:      WasteLibrary.STATU_ACTIVE,
			UltApp:       WasteLibrary.STATU_PASSIVE,
			RecyApp:      WasteLibrary.STATU_ACTIVE,
		}

		data := url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentCustomer.ToString()},
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
			resultVal = WasteLibrary.ByteToResultType(bodyBytes)
			WasteLibrary.LogStr(resultVal.ToString())
		}
	} else if opType == WasteLibrary.DEVICE {
		var currentHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.HttpClientHeaderType{
			AppType:      WasteLibrary.ADMIN,
			DeviceNo:     "",
			OpType:       WasteLibrary.DEVICE,
			Time:         WasteLibrary.GetTime(),
			Repeat:       WasteLibrary.STATU_PASSIVE,
			DeviceId:     0,
			CustomerId:   0,
			BaseDataType: WasteLibrary.DEVICE,
		}

		var deviceId int = 0
		var customerId int = 1
		var currentDevice WasteLibrary.DeviceType = WasteLibrary.DeviceType{
			DeviceId:     float64(deviceId),
			CustomerId:   float64(customerId),
			DeviceName:   "07 MVS 33",
			DeviceType:   WasteLibrary.RFID,
			SerialNumber: "00000000c1b1d188",
		}
		WasteLibrary.LogStr(currentDevice.ToString())
		data := url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentDevice.ToString()},
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
			resultVal = WasteLibrary.ByteToResultType(bodyBytes)
			WasteLibrary.LogStr(resultVal.ToString())
		}
	} else if opType == "CUSTOMERCONFIG" {
		var currentHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.HttpClientHeaderType{
			AppType:      WasteLibrary.ADMIN,
			DeviceNo:     "",
			OpType:       WasteLibrary.CUSTOMER,
			Time:         WasteLibrary.GetTime(),
			Repeat:       WasteLibrary.STATU_PASSIVE,
			DeviceId:     0,
			CustomerId:   0,
			BaseDataType: WasteLibrary.DEVICE,
		}

		var customerId int = 1
		var currentData WasteLibrary.CustomerConfigType = WasteLibrary.CustomerConfigType{
			CustomerId:      float64(customerId),
			ArventoApp:      WasteLibrary.STATU_ACTIVE,
			ArventoUserName: "afatekbilisim",
			ArventoPin1:     "Amca151200!Furkan",
			ArventoPin2:     "Amca151200!Furkan",
			SystemProblem:   WasteLibrary.STATU_ACTIVE,
			TruckStopTrace:  WasteLibrary.STATU_ACTIVE,
			Active:          WasteLibrary.STATU_ACTIVE,
			CreateTime:      WasteLibrary.GetTime(),
		}
		WasteLibrary.LogStr(currentData.ToString())
		data := url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentData.ToString()},
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
			resultVal = WasteLibrary.ByteToResultType(bodyBytes)
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
			resultVal = WasteLibrary.ByteToResultType(bodyBytes)
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
			resultVal = WasteLibrary.ByteToResultType(bodyBytes)
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
	} else {*/

	testVal := GetMD5Hash("123")
	fmt.Println(testVal)

	//}

}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

/*
func test123(testt string) {
	fmt.Println((testt))
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
*/
