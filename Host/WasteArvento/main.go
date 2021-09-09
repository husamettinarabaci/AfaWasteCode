package main

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/AfatekDevelopers/result_lib_go/devafatekresult"
	"github.com/devafatek/WasteLibrary"
)

var currentCustomerList WasteLibrary.CustomersType = WasteLibrary.CustomersType{
	Customers: make(map[float64]float64),
}

func initStart() {

	WasteLibrary.LogStr("Successfully connected!")
}

func main() {

	initStart()

	//go mainProc()
	go setCustomerList()

	http.HandleFunc("/health", WasteLibrary.HealthHandler)
	http.HandleFunc("/readiness", WasteLibrary.ReadinessHandler)
	http.HandleFunc("/status", WasteLibrary.StatusHandler)
	http.ListenAndServe(":80", nil)
}

func setCustomerList() {
	var resultVal devafatekresult.ResultType
	resultVal.Result = "FAIL"
	for {
		time.Sleep(60 * 60 * time.Second)
		resultVal = WasteLibrary.GetRedisForStoreApi("customers", "customers")
		var currentCustomers WasteLibrary.CustomersType = WasteLibrary.StringToCustomersType(resultVal.Retval.(string))
		for _, customerId := range currentCustomers.Customers {
			if _, ok := currentCustomerList.Customers[customerId]; !ok {
				currentCustomerList.Customers[customerId] = customerId
				go customerProc(customerId)

			}
		}

	}
}

func customerProc(customerId float64) {
	var resultVal devafatekresult.ResultType
	resultVal.Result = "FAIL"
	var loopCount = 0
	var currentCustomerConfig WasteLibrary.CustomerConfigType
	var currentCustomerDevices WasteLibrary.CustomerDevicesType
	var plateDevice map[string]string
	resultVal = WasteLibrary.GetRedisForStoreApi("customer-customerconfig", WasteLibrary.Float64IdToString(customerId))

	if resultVal.Result == "OK" {
		currentCustomerConfig = WasteLibrary.StringToCustomerConfigType(resultVal.Retval.(string))
	}
	for {
		if currentCustomerConfig.ArventoApp == "1" {
			if loopCount == 180 {
				loopCount = 0
			}
			if loopCount == 0 {
				loopCount = 0
				resultVal = WasteLibrary.GetRedisForStoreApi("customer-customerconfig", WasteLibrary.Float64IdToString(customerId))
				if resultVal.Result == "OK" {
					currentCustomerConfig = WasteLibrary.StringToCustomerConfigType(resultVal.Retval.(string))
				}
				resultVal = WasteLibrary.GetRedisForStoreApi("customer-devices", currentCustomerConfig.ToIdString())
				if resultVal.Result == "OK" {
					currentCustomerDevices = WasteLibrary.StringToCustomerDevicesType(resultVal.Retval.(string))
				}
				resultVal = getDevice(currentCustomerConfig)
				if resultVal.Result == "OK" {
					plateDevice = resultVal.Retval.(map[string]string)
				}
			}

			resultVal = getLocation(currentCustomerConfig)
			if resultVal.Result == "OK" {
				var deviceLocations WasteLibrary.ArventoDeviceGpsListType = WasteLibrary.StringToArventoDeviceGpsListType(resultVal.Retval.(string))
				for _, vDevice := range currentCustomerDevices.Devices {
					resultVal = WasteLibrary.GetRedisForStoreApi("devices", WasteLibrary.Float64IdToString(vDevice))
					if resultVal.Result == "OK" {
						var currentDevice WasteLibrary.DeviceType = WasteLibrary.StringToDeviceType(resultVal.Retval.(string))
						if currentDevice.DeviceType == "RFID" {
							var arventoId string = plateDevice[currentDevice.DeviceName]
							if currentDeviceLocation, ok := deviceLocations.ArventoDeviceGpsList[arventoId]; ok {
								currentDevice.Latitude = currentDeviceLocation.Latitude
								currentDevice.Longitude = currentDeviceLocation.Longitude
								currentDevice.Speed = currentDeviceLocation.Speed

								var newCurrentHttpHeader WasteLibrary.HttpClientHeaderType
								newCurrentHttpHeader.AppType = "RFID"
								newCurrentHttpHeader.OpType = "ARVENTO"
								data := url.Values{
									"HEADER": {newCurrentHttpHeader.ToString()},
									"DATA":   {currentDevice.ToString()},
								}
								resultVal = WasteLibrary.SaveStaticDbMainForStoreApi(data)

								if resultVal.Result == "OK" {
									resultVal = WasteLibrary.SaveRedisForStoreApi("devices", currentDevice.ToIdString(), currentDevice.ToString())
								}
								if currentDevice.Speed == 0 {
									//TO DO
									//Speed Op
								}

							}
						}
					}
				}
			}

			time.Sleep(20 * time.Second)
		} else {
			delete(currentCustomerList.Customers, customerId)
			break
		}
		loopCount++
	}
}

func getLocation(currentCustomerConfig WasteLibrary.CustomerConfigType) devafatekresult.ResultType {
	var resultVal devafatekresult.ResultType
	resultVal.Result = "FAIL"
	var deviceLocation WasteLibrary.ArventoDeviceGpsListType
	resp, err := http.Get("http://ws.arvento.com/v1/report.asmx/GetVehicleStatus?Username=" + currentCustomerConfig.ArventoUserName + "&PIN1=" + currentCustomerConfig.ArventoPin1 + "&PIN2=" + currentCustomerConfig.ArventoPin2 + "&Language=tr")
	if err != nil {
		WasteLibrary.LogErr(err)
		return resultVal
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		WasteLibrary.LogErr(err)
		return resultVal
	}

	bodys := string(body)
	var deviceID string = ""
	var latitude string = ""
	var longitude string = ""
	var speed string = ""
	for {
		if strings.Index(bodys, "<Device_x0020_No>") > -1 {
			deviceID = bodys[strings.Index(bodys, "<Device_x0020_No>")+17 : strings.Index(bodys, "</Device_x0020_No>")]
			latitude = bodys[strings.Index(bodys, "<Latitude>")+10 : strings.Index(bodys, "</Latitude>")]
			longitude = bodys[strings.Index(bodys, "<Longitude>")+11 : strings.Index(bodys, "</Longitude>")]
			speed = bodys[strings.Index(bodys, "<Speed>")+7 : strings.Index(bodys, "</Speed>")]
			bodys = bodys[strings.Index(bodys, "</Longitude>")+1:]

			var currrentArventoDeviceGpsType WasteLibrary.ArventoDeviceGpsType
			if latitude != "" {
				currrentArventoDeviceGpsType.Latitude = WasteLibrary.StringIdToFloat64(latitude)
			} else {
				currrentArventoDeviceGpsType.Latitude = 0
			}
			if longitude != "" {
				currrentArventoDeviceGpsType.Longitude = WasteLibrary.StringIdToFloat64(longitude)
			} else {
				currrentArventoDeviceGpsType.Longitude = 0
			}
			if speed != "" {
				currrentArventoDeviceGpsType.Speed = WasteLibrary.StringIdToFloat64(speed)
			} else {
				currrentArventoDeviceGpsType.Speed = 0
			}
			deviceLocation.ArventoDeviceGpsList[deviceID] = currrentArventoDeviceGpsType

		} else {
			break
		}
	}

	resultVal.Result = "OK"
	resultVal.Retval = deviceLocation
	return resultVal

}

func getDevice(currentCustomerConfig WasteLibrary.CustomerConfigType) devafatekresult.ResultType {
	var resultVal devafatekresult.ResultType
	resultVal.Result = "FAIL"
	var plateDevice map[string]string = make(map[string]string)
	resp, err := http.Get("http://ws.arvento.com/v1/report.asmx/GetLicensePlateNodeMappings?Username=" + currentCustomerConfig.ArventoUserName + "&PIN1=" + currentCustomerConfig.ArventoPin1 + "&PIN2=" + currentCustomerConfig.ArventoPin2 + "&Language=tr")
	if err != nil {
		WasteLibrary.LogErr(err)
		return resultVal
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		WasteLibrary.LogErr(err)
		return resultVal
	}

	bodys := string(body)
	var deviceID string = ""
	var plateNO string = ""
	for {
		if strings.Index(bodys, "<Device_x0020_No>") > -1 {
			deviceID = bodys[strings.Index(bodys, "<Device_x0020_No>")+17 : strings.Index(bodys, "</Device_x0020_No>")]
			plateNO = bodys[strings.Index(bodys, "<License_x0020_Plate>")+21 : strings.Index(bodys, "</License_x0020_Plate>")]
			bodys = bodys[strings.Index(bodys, "</License_x0020_Plate>")+1:]
			plateDevice[plateNO] = deviceID
		} else {
			break
		}
	}
	resultVal.Result = "OK"
	resultVal.Retval = plateDevice
	return resultVal
}
