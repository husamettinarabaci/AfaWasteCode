package main

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/devafatek/WasteLibrary"
)

func initStart() {

	WasteLibrary.LogStr("Successfully connected!")
	go WasteLibrary.InitLog()
}
func main() {

	initStart()

	http.HandleFunc("/health", WasteLibrary.HealthHandler)
	http.HandleFunc("/readiness", WasteLibrary.ReadinessHandler)
	http.HandleFunc("/status", WasteLibrary.StatusHandler)
	http.HandleFunc("/data", data)
	http.ListenAndServe(":80", nil)
}

func data(w http.ResponseWriter, req *http.Request) {

	if WasteLibrary.AllowCors {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
	}

	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL

	if err := req.ParseForm(); err != nil {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_HTTP_PARSE
		w.Write(resultVal.ToByte())

		WasteLibrary.LogErr(err)
		return
	}

	var currentHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.StringToHttpClientHeaderType(req.FormValue(WasteLibrary.HTTP_HEADER))
	if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_RFID {
		var currentDevice WasteLibrary.RfidDeviceType
		currentDevice.New()
		resultVal = currentDevice.GetByRedisBySerial(currentHttpHeader.DeviceNo)
		if resultVal.Result == WasteLibrary.RESULT_FAIL {
			if currentHttpHeader.ReaderType == WasteLibrary.READERTYPE_STATUS {
				resultVal = createDevice(currentHttpHeader, req.FormValue(WasteLibrary.HTTP_DATA))
				if resultVal.Result != WasteLibrary.RESULT_OK {
					resultVal.Result = WasteLibrary.RESULT_FAIL
					resultVal.Retval = WasteLibrary.RESULT_ERROR_HTTP_POST
					w.Write(resultVal.ToByte())

					return
				}
				resultVal = currentDevice.GetByRedisBySerial(currentHttpHeader.DeviceNo)
				if resultVal.Result != WasteLibrary.RESULT_OK {
					resultVal.Result = WasteLibrary.RESULT_FAIL
					resultVal.Retval = WasteLibrary.RESULT_ERROR_DEVICE_NOTFOUND
					w.Write(resultVal.ToByte())

					return
				}
			} else {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_IGNORE_FIRST_DATA
				w.Write(resultVal.ToByte())

				return
			}
		}
		currentHttpHeader.CustomerId = currentDevice.DeviceMain.CustomerId
		currentHttpHeader.DeviceId = currentDevice.DeviceId
	} else if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_ULT {
		var currentDevice WasteLibrary.UltDeviceType
		currentDevice.New()
		resultVal = currentDevice.GetByRedisBySerial(currentHttpHeader.DeviceNo)
		if resultVal.Result == WasteLibrary.RESULT_FAIL {
			resultVal = createDevice(currentHttpHeader, req.FormValue(WasteLibrary.HTTP_DATA))
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_HTTP_POST
				w.Write(resultVal.ToByte())

				return
			}
			resultVal = currentDevice.GetByRedisBySerial(currentHttpHeader.DeviceNo)
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_DEVICE_NOTFOUND
				w.Write(resultVal.ToByte())

				return
			}
		}
		currentHttpHeader.CustomerId = currentDevice.DeviceMain.CustomerId
		currentHttpHeader.DeviceId = currentDevice.DeviceId
	} else if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_RECY {
		var currentDevice WasteLibrary.RecyDeviceType
		currentDevice.New()
		resultVal = currentDevice.GetByRedisBySerial(currentHttpHeader.DeviceNo)
		if resultVal.Result == WasteLibrary.RESULT_FAIL {
			if currentHttpHeader.ReaderType == WasteLibrary.READERTYPE_STATUS {
				resultVal = createDevice(currentHttpHeader, req.FormValue(WasteLibrary.HTTP_DATA))
				if resultVal.Result != WasteLibrary.RESULT_OK {
					resultVal.Result = WasteLibrary.RESULT_FAIL
					resultVal.Retval = WasteLibrary.RESULT_ERROR_HTTP_POST
					w.Write(resultVal.ToByte())

					return
				}
				resultVal = currentDevice.GetByRedisBySerial(currentHttpHeader.DeviceNo)
				if resultVal.Result != WasteLibrary.RESULT_OK {
					resultVal.Result = WasteLibrary.RESULT_FAIL
					resultVal.Retval = WasteLibrary.RESULT_ERROR_DEVICE_NOTFOUND
					w.Write(resultVal.ToByte())

					return
				}
			} else {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_IGNORE_FIRST_DATA
				w.Write(resultVal.ToByte())

				return
			}
		}
		currentHttpHeader.CustomerId = currentDevice.DeviceMain.CustomerId
		currentHttpHeader.DeviceId = currentDevice.DeviceId
	} else {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_DEVICE_NOTFOUND
		w.Write(resultVal.ToByte())

		return
	}

	var serviceClusterIp string = ""
	if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_RFID {

		if currentHttpHeader.ReaderType == WasteLibrary.READERTYPE_RF {
			serviceClusterIp = "waste-rfreader-cluster-ip"
			resultVal = sendReader(serviceClusterIp, currentHttpHeader.ToString(), req.FormValue(WasteLibrary.HTTP_DATA))
		} else if currentHttpHeader.ReaderType == WasteLibrary.READERTYPE_GPS {
			serviceClusterIp = "waste-gpsreader-cluster-ip"
			resultVal = sendReader(serviceClusterIp, currentHttpHeader.ToString(), req.FormValue(WasteLibrary.HTTP_DATA))

		} else if currentHttpHeader.ReaderType == WasteLibrary.READERTYPE_STATUS {
			serviceClusterIp = "waste-statusreader-cluster-ip"
			resultVal = sendReader(serviceClusterIp, currentHttpHeader.ToString(), req.FormValue(WasteLibrary.HTTP_DATA))
		} else if currentHttpHeader.ReaderType == WasteLibrary.READERTYPE_THERM {
			serviceClusterIp = "waste-thermreader-cluster-ip"
			resultVal = sendReader(serviceClusterIp, currentHttpHeader.ToString(), req.FormValue(WasteLibrary.HTTP_DATA))
		} else if currentHttpHeader.ReaderType == WasteLibrary.READERTYPE_CAM {
			//serviceClusterIp = "waste-camreader-cluster-ip"
			//resultVal = sendReader(serviceClusterIp, currentHttpHeader.ToString(), req.FormValue(WasteLibrary.HTTP_DATA))
		} else {
			resultVal.Result = WasteLibrary.RESULT_FAIL
		}
	} else if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_ULT {
		if currentHttpHeader.ReaderType == WasteLibrary.READERTYPE_ULT {

			//serviceClusterIp = "waste-gpsreader-cluster-ip"
			//resultVal = sendReader(serviceClusterIp, currentHttpHeader.ToString(), req.FormValue(WasteLibrary.HTTP_DATA))

			serviceClusterIp = "waste-statusreader-cluster-ip"
			resultVal = sendReader(serviceClusterIp, currentHttpHeader.ToString(), req.FormValue(WasteLibrary.HTTP_DATA))

			serviceClusterIp = "waste-thermreader-cluster-ip"
			resultVal = sendReader(serviceClusterIp, currentHttpHeader.ToString(), req.FormValue(WasteLibrary.HTTP_DATA))

			serviceClusterIp = "waste-batteryreader-cluster-ip"
			resultVal = sendReader(serviceClusterIp, currentHttpHeader.ToString(), req.FormValue(WasteLibrary.HTTP_DATA))

			serviceClusterIp = "waste-sensreader-cluster-ip"
			resultVal = sendReader(serviceClusterIp, currentHttpHeader.ToString(), req.FormValue(WasteLibrary.HTTP_DATA))

		} else {
			resultVal.Result = WasteLibrary.RESULT_FAIL
		}
	} else if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_RECY {
		if currentHttpHeader.ReaderType == WasteLibrary.READERTYPE_RF {
			var currentNfc WasteLibrary.NfcType = WasteLibrary.StringToNfcType(req.FormValue(WasteLibrary.HTTP_DATA))
			currentNfc.NfcBase.Name = "İsim"
			currentNfc.NfcBase.SurName = "Soyisim"
			currentNfc.NfcBase.TotalAmount = 10
			currentNfc.NfcStatu.ItemStatu = WasteLibrary.RECY_ITEM_STATU_PLASTIC
			currentNfc.NfcBase.LastAmount = 0
			resultVal.Result = WasteLibrary.RESULT_OK
			resultVal.Retval = currentNfc.ToString()
		} else if currentHttpHeader.ReaderType == WasteLibrary.READERTYPE_STATUS {
			//serviceClusterIp = "waste-statusreader-cluster-ip"
			//resultVal = sendReader(serviceClusterIp, currentHttpHeader.ToString(), req.FormValue(WasteLibrary.HTTP_DATA))
		} else if currentHttpHeader.ReaderType == WasteLibrary.READERTYPE_THERM {
			//serviceClusterIp = "waste-thermreader-cluster-ip"
			//resultVal = sendReader(serviceClusterIp, currentHttpHeader.ToString(), req.FormValue(WasteLibrary.HTTP_DATA))
		} else if currentHttpHeader.ReaderType == WasteLibrary.READERTYPE_CAM {
			var currentNfc WasteLibrary.NfcType = WasteLibrary.StringToNfcType(req.FormValue(WasteLibrary.HTTP_DATA))

			client := http.Client{
				Timeout: 10 * time.Second,
			}
			resp, err := client.Get("http://161.35.160.129:3000/upload?uid=" + currentNfc.NfcReader.UID)
			if err != nil {
				WasteLibrary.LogErr(err)

			} else {
				_, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					WasteLibrary.LogErr(err)
				}
			}
		} else if currentHttpHeader.ReaderType == WasteLibrary.READERTYPE_GET_NFC {
			time.Sleep(20 * time.Second)
			var currentNfc WasteLibrary.NfcType = WasteLibrary.StringToNfcType(req.FormValue(WasteLibrary.HTTP_DATA))
			var retVal string
			retVal = "cam"
			client := http.Client{
				Timeout: 10 * time.Second,
			}

			resp, err := client.Get("http://161.35.160.129:3000/check?uid=" + currentNfc.NfcReader.UID)
			if err != nil {
				WasteLibrary.LogErr(err)

			} else {
				bodyBytes, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					WasteLibrary.LogErr(err)
				}
				retVal = string(bodyBytes)
			}
			WasteLibrary.LogStr(retVal)
			if strings.Contains(retVal, "cam") {
				retVal = "3"
			}
			if strings.Contains(retVal, "metal") {
				retVal = "2"
			}
			if strings.Contains(retVal, "plastik") {
				retVal = "1"
			}

			currentNfc.NfcBase.Name = "İsim"
			currentNfc.NfcBase.SurName = "Soyisim"
			if retVal == "3" {
				currentNfc.NfcStatu.ItemStatu = WasteLibrary.RECY_ITEM_STATU_GLASS
				currentNfc.NfcBase.LastAmount = 3
				currentNfc.NfcBase.TotalAmount = 13
			} else if retVal == "2" {
				currentNfc.NfcStatu.ItemStatu = WasteLibrary.RECY_ITEM_STATU_METAL
				currentNfc.NfcBase.LastAmount = 2
				currentNfc.NfcBase.TotalAmount = 12
			} else {
				currentNfc.NfcStatu.ItemStatu = WasteLibrary.RECY_ITEM_STATU_PLASTIC
				currentNfc.NfcBase.LastAmount = 1
				currentNfc.NfcBase.TotalAmount = 11
			}
			resultVal.Result = WasteLibrary.RESULT_OK
			resultVal.Retval = currentNfc.ToString()
		} else if currentHttpHeader.ReaderType == WasteLibrary.READERTYPE_GET_CUSTOMER {
			var currentCustomer WasteLibrary.CustomerType
			currentCustomer.CustomerId = currentHttpHeader.CustomerId
			currentCustomer.GetByRedis()
			currentCustomer.CustomerName = "BODRUM"
			resultVal.Result = WasteLibrary.RESULT_OK
			resultVal.Retval = currentCustomer.ToString()
		} else {
			resultVal.Result = WasteLibrary.RESULT_FAIL
		}
	} else {
		resultVal.Result = WasteLibrary.RESULT_FAIL
	}
	w.Write(resultVal.ToByte())

}

func sendReader(serviceClusterIp string, httpHeader string, httpData string) WasteLibrary.ResultType {
	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	data := url.Values{
		WasteLibrary.HTTP_HEADER: {httpHeader},
		WasteLibrary.HTTP_DATA:   {httpData},
	}

	resultVal = WasteLibrary.SaveBulkDbMainForStoreApi(data)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE

		return resultVal
	}
	if serviceClusterIp != "" {
		resultVal = WasteLibrary.HttpPostReq("http://"+serviceClusterIp+"/reader", data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_HTTP_POST

			return resultVal
		}
	}
	return resultVal
}

func createDevice(currentHttpHeader WasteLibrary.HttpClientHeaderType, currentData string) WasteLibrary.ResultType {
	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	var createHttpHeader WasteLibrary.HttpClientHeaderType
	createHttpHeader.New()
	createHttpHeader.DeviceType = currentHttpHeader.DeviceType
	createHttpHeader.DeviceNo = currentHttpHeader.DeviceNo
	data := url.Values{
		WasteLibrary.HTTP_HEADER: {createHttpHeader.ToString()},
		WasteLibrary.HTTP_DATA:   {currentData},
	}

	resultVal = WasteLibrary.HttpPostReq("http://waste-enhcapi-cluster-ip/createDevice", data)
	return resultVal
}
