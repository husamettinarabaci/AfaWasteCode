package main

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/devafatek/WasteLibrary"
)

var currentCustomerList WasteLibrary.CustomersType

func initStart() {

	WasteLibrary.LogStr("Successfully connected!")
	go WasteLibrary.InitLog()
	currentCustomerList.New()
}

func main() {

	initStart()

	go setCustomerList()

	http.HandleFunc("/health", WasteLibrary.HealthHandler)
	http.HandleFunc("/readiness", WasteLibrary.ReadinessHandler)
	http.HandleFunc("/status", WasteLibrary.StatusHandler)
	http.ListenAndServe(":80", nil)
}

func setCustomerList() {
	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	for {

		resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMERS, WasteLibrary.REDIS_CUSTOMERS)

		if resultVal.Result == WasteLibrary.RESULT_OK {

			var currentCustomers WasteLibrary.CustomersType = WasteLibrary.StringToCustomersType(resultVal.Retval.(string))
			for _, customerId := range currentCustomers.Customers {
				if customerId != 0 {
					if _, ok := currentCustomerList.Customers[WasteLibrary.Float64IdToString(customerId)]; !ok {
						WasteLibrary.LogStr("Add Customer : " + WasteLibrary.Float64IdToString(customerId))
						currentCustomerList.Customers[WasteLibrary.Float64IdToString(customerId)] = customerId
						go customerProc(customerId)
						time.Sleep(60 * time.Second)
					}
				}
			}
		}
		time.Sleep(60 * 60 * time.Second)
	}
}

func customerProc(customerId float64) {
	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	var loopCount = 0
	var currentCustomerConfig WasteLibrary.CustomerConfigType
	var currentCustomerDevices WasteLibrary.CustomerRfidDevicesType
	var plateDevice map[string]string
	resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_CUSTOMERCONFIG, WasteLibrary.Float64IdToString(customerId))

	if resultVal.Result == WasteLibrary.RESULT_OK {
		currentCustomerConfig = WasteLibrary.StringToCustomerConfigType(resultVal.Retval.(string))
	}
	for {
		if currentCustomerConfig.ArventoApp == WasteLibrary.STATU_ACTIVE {
			if loopCount == 180 {
				loopCount = 0
			}
			if loopCount == 0 {
				loopCount = 0
				resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_CUSTOMERCONFIG, WasteLibrary.Float64IdToString(customerId))
				if resultVal.Result == WasteLibrary.RESULT_OK {
					currentCustomerConfig = WasteLibrary.StringToCustomerConfigType(resultVal.Retval.(string))
				}
				resultVal = WasteLibrary.GetRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_RFID_DEVICES, currentCustomerConfig.ToIdString())
				if resultVal.Result == WasteLibrary.RESULT_OK {
					WasteLibrary.LogStr("Add Devices : " + WasteLibrary.Float64IdToString(customerId))
					currentCustomerDevices = WasteLibrary.StringToCustomerRfidDevicesType(resultVal.Retval.(string))
				}
				resultVal = getDevice(currentCustomerConfig)
				if resultVal.Result == WasteLibrary.RESULT_OK {
					plateDevice = resultVal.Retval.(map[string]string)
				}
			}

			resultVal = getLocation(currentCustomerConfig)
			if resultVal.Result == WasteLibrary.RESULT_OK {
				var deviceLocations WasteLibrary.ArventoDeviceGpsListType = WasteLibrary.StringToArventoDeviceGpsListType(resultVal.Retval.(string))
				for _, vDevice := range currentCustomerDevices.Devices {
					if vDevice == 0 {
						continue
					}
					var currentDevice WasteLibrary.RfidDeviceType
					currentDevice.New()
					currentDevice.DeviceId = vDevice
					resultVal = currentDevice.GetAll()
					if resultVal.Result == WasteLibrary.RESULT_OK {
						var arventoId string = plateDevice[currentDevice.DeviceDetail.PlateNo]
						if currentDeviceLocation, ok := deviceLocations.ArventoDeviceGpsList[arventoId]; ok {
							if currentDeviceLocation.Latitude != 0 && currentDeviceLocation.Longitude != 0 {
								currentDevice.DeviceGps.Latitude = currentDeviceLocation.Latitude
								currentDevice.DeviceGps.Longitude = currentDeviceLocation.Longitude
								currentDevice.DeviceGps.GpsTime = WasteLibrary.TimeToString(WasteLibrary.AddTimeToBase(WasteLibrary.StringToTime(currentDeviceLocation.GpsTime), 3*time.Hour))
								currentDevice.DeviceGps.Speed = currentDeviceLocation.Speed
								WasteLibrary.LogStr("Devices Gps : " +
									WasteLibrary.Float64IdToString(customerId) + " - " +
									WasteLibrary.Float64IdToString(currentDevice.DeviceId) + " - " +
									WasteLibrary.Float64ToString(currentDevice.DeviceGps.Latitude) + " - " +
									WasteLibrary.Float64ToString(currentDevice.DeviceGps.Longitude) + " - " +
									WasteLibrary.Float64ToString(currentDevice.DeviceGps.Speed))
								var newCurrentHttpHeader WasteLibrary.HttpClientHeaderType
								newCurrentHttpHeader.New()
								newCurrentHttpHeader.DeviceType = WasteLibrary.DEVICETYPE_RFID
								newCurrentHttpHeader.DeviceId = currentDevice.DeviceId
								newCurrentHttpHeader.CustomerId = currentDevice.DeviceMain.CustomerId
								newCurrentHttpHeader.DataType = WasteLibrary.DATATYPE_RFID_GPS_DEVICE
								data := url.Values{
									WasteLibrary.HTTP_HEADER: {newCurrentHttpHeader.ToString()},
									WasteLibrary.HTTP_DATA:   {currentDevice.ToString()},
								}

								WasteLibrary.LogStr("Send Gps Reader" + currentDevice.ToString())
								resultVal = WasteLibrary.HttpPostReq("http://waste-gpsreader-cluster-ip/reader", data)
							}
						}

					}
				}
			}

			time.Sleep(20 * time.Second)
		} else {
			delete(currentCustomerList.Customers, WasteLibrary.Float64IdToString(customerId))
			break
		}
		loopCount++
	}
}

func getLocation(currentCustomerConfig WasteLibrary.CustomerConfigType) WasteLibrary.ResultType {
	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
	var deviceLocation WasteLibrary.ArventoDeviceGpsListType
	deviceLocation.New()

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
	var gpstime string = ""
	for {

		if strings.Contains(bodys, "<Device_x0020_No>") {
			deviceID = bodys[strings.Index(bodys, "<Device_x0020_No>")+17 : strings.Index(bodys, "</Device_x0020_No>")]
			gpstime = bodys[strings.Index(bodys, "<GMT_x0020_Date_x002F_Time>")+27 : strings.Index(bodys, "</GMT_x0020_Date_x002F_Time>")]
			latitude = bodys[strings.Index(bodys, "<Latitude>")+10 : strings.Index(bodys, "</Latitude>")]
			longitude = bodys[strings.Index(bodys, "<Longitude>")+11 : strings.Index(bodys, "</Longitude>")]
			speed = bodys[strings.Index(bodys, "<Speed>")+7 : strings.Index(bodys, "</Speed>")]
			bodys = bodys[strings.Index(bodys, "</Longitude>")+1:]
			var currrentArventoDeviceGpsType WasteLibrary.ArventoDeviceGpsType
			currrentArventoDeviceGpsType.New()
			if latitude != "" {
				currrentArventoDeviceGpsType.Latitude = WasteLibrary.StringToFloat64(latitude)
			} else {
				currrentArventoDeviceGpsType.Latitude = 0
			}
			if longitude != "" {
				currrentArventoDeviceGpsType.Longitude = WasteLibrary.StringToFloat64(longitude)
			} else {
				currrentArventoDeviceGpsType.Longitude = 0
			}
			if speed != "" {
				currrentArventoDeviceGpsType.Speed = WasteLibrary.StringToFloat64(speed)
			} else {
				currrentArventoDeviceGpsType.Speed = -1
			}
			currrentArventoDeviceGpsType.GpsTime = gpstime
			deviceLocation.ArventoDeviceGpsList[deviceID] = currrentArventoDeviceGpsType

		} else {
			break
		}
	}

	resultVal.Result = WasteLibrary.RESULT_OK
	resultVal.Retval = deviceLocation.ToString()
	return resultVal

}

func getDevice(currentCustomerConfig WasteLibrary.CustomerConfigType) WasteLibrary.ResultType {
	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL
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
		if strings.Contains(bodys, "<Device_x0020_No>") {
			deviceID = bodys[strings.Index(bodys, "<Device_x0020_No>")+17 : strings.Index(bodys, "</Device_x0020_No>")]
			plateNO = bodys[strings.Index(bodys, "<License_x0020_Plate>")+21 : strings.Index(bodys, "</License_x0020_Plate>")]
			bodys = bodys[strings.Index(bodys, "</License_x0020_Plate>")+1:]
			plateDevice[plateNO] = deviceID
		} else {
			break
		}
	}
	resultVal.Result = WasteLibrary.RESULT_OK
	resultVal.Retval = plateDevice
	return resultVal
}
