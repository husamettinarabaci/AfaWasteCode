package main

import (
	"net/http"

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
	http.HandleFunc("/logApp", logApp)
	http.HandleFunc("/setCustomer", setCustomer)
	http.HandleFunc("/getCustomer", getCustomer)
	http.HandleFunc("/getCustomers", getCustomers)
	http.HandleFunc("/setDevice", setDevice)
	http.HandleFunc("/getDevice", getDevice)
	http.HandleFunc("/getConfig", getConfig)
	http.HandleFunc("/setConfig", setConfig)
	http.HandleFunc("/getDevices", getDevices)
	http.HandleFunc("/startSystem", startSystem)
	http.ListenAndServe(":80", nil)
}

func startSystem(w http.ResponseWriter, req *http.Request) {

	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL

	if err := req.ParseForm(); err != nil {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_HTTP_PARSE
		w.Write(resultVal.ToByte())

		WasteLibrary.LogErr(err)
		return
	}

	var customerBase WasteLibrary.CustomerType
	customerBase.CustomerId = 1
	resultVal = customerBase.GetByRedis()
	if resultVal.Result == WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_APP_STARTED
	} else {
		WasteLibrary.LogStr("AfatekApi Start System Add Customer AFATEK")
		resultVal = createCustomer("AFATEK", "afatek.aws.afatek.com.tr")
		resultVal = createCustomer("AFATEKWEB", "atik.afatek.com.tr")
		resultVal = createCustomer("BODRUM", "atik.bodrum.bel.tr")

	}
	w.Write(resultVal.ToByte())

}

func createCustomer(customerName string, customerLink string) WasteLibrary.ResultType {
	var resultVal WasteLibrary.ResultType
	var currentData WasteLibrary.CustomerType
	currentData.New()
	currentData.CustomerName = customerName
	currentData.CustomerLink = customerLink
	currentData.RfIdApp = WasteLibrary.STATU_ACTIVE
	currentData.UltApp = WasteLibrary.STATU_ACTIVE
	currentData.RecyApp = WasteLibrary.STATU_ACTIVE
	currentData.Active = WasteLibrary.STATU_ACTIVE
	currentData.CreateTime = WasteLibrary.GetTime()

	resultVal = currentData.SaveToDb()
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE

		return resultVal
	}

	resultVal = currentData.SaveToRedis()
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE

		return resultVal
	}

	resultVal = currentData.SaveToRedisLink()
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE

		return resultVal
	}

	resultVal = createCustomerPart(currentData)

	return resultVal
}

func createCustomerPart(currentData WasteLibrary.CustomerType) WasteLibrary.ResultType {
	var resultVal WasteLibrary.ResultType

	var currentCustomerRfidDevices WasteLibrary.CustomerRfidDevicesType
	currentCustomerRfidDevices.New()
	currentCustomerRfidDevices.CustomerId = currentData.CustomerId
	currentCustomerRfidDevices.Devices["0"] = 0

	var currentCustomerRfidDeviceViews WasteLibrary.CustomerRfidDevicesViewListType
	currentCustomerRfidDeviceViews.New()
	currentCustomerRfidDeviceViews.CustomerId = currentData.CustomerId

	var currentCustomerRecyDevices WasteLibrary.CustomerRecyDevicesType
	currentCustomerRecyDevices.New()
	currentCustomerRecyDevices.CustomerId = currentData.CustomerId
	currentCustomerRecyDevices.Devices["0"] = 0

	var currentCustomerRecyDeviceViews WasteLibrary.CustomerRecyDevicesViewListType
	currentCustomerRecyDeviceViews.New()
	currentCustomerRecyDeviceViews.CustomerId = currentData.CustomerId

	var currentCustomerUltDevices WasteLibrary.CustomerUltDevicesType
	currentCustomerUltDevices.New()
	currentCustomerUltDevices.CustomerId = currentData.CustomerId
	currentCustomerUltDevices.Devices["0"] = 0

	var currentCustomerUltDeviceViews WasteLibrary.CustomerUltDevicesViewListType
	currentCustomerUltDeviceViews.New()
	currentCustomerUltDeviceViews.CustomerId = currentData.CustomerId

	var currentCustomerUsers WasteLibrary.CustomerUsersType
	currentCustomerUsers.New()
	currentCustomerUsers.CustomerId = currentData.CustomerId
	currentCustomerUsers.Users["0"] = 0

	var currentCustomerTags WasteLibrary.CustomerTagsType
	currentCustomerTags.New()
	currentCustomerTags.CustomerId = currentData.CustomerId
	currentCustomerTags.Tags["0"] = 0

	var currentCustomerTagViews WasteLibrary.CustomerTagsViewListType
	currentCustomerTagViews.New()
	currentCustomerTagViews.CustomerId = currentData.CustomerId

	var currentCustomerConfig WasteLibrary.CustomerConfigType
	currentCustomerConfig.New()
	currentCustomerConfig.CustomerId = currentData.CustomerId
	currentCustomerConfig.ArventoApp = WasteLibrary.STATU_ACTIVE
	currentCustomerConfig.ArventoUserName = "afatekbilisim"
	currentCustomerConfig.ArventoPin1 = "Amca151200!Furkan"
	currentCustomerConfig.ArventoPin2 = "Amca151200!Furkan"

	var currentAdminConfig WasteLibrary.AdminConfigType
	currentAdminConfig.New()
	currentAdminConfig.CustomerId = currentData.CustomerId
	currentAdminConfig.CenterLatitude = 37.036604
	currentAdminConfig.CenterLongitude = 27.424406

	var parkLoc1 WasteLibrary.LocationType
	parkLoc1.LocationName = "TURGUT REİS"
	parkLoc1.ZoneRadius = 50
	parkLoc1.Latitude = 37.012583
	parkLoc1.Longitude = 27.283889

	var parkLoc2 WasteLibrary.LocationType
	parkLoc2.LocationName = "GÜNDOĞAN"
	parkLoc2.ZoneRadius = 50
	parkLoc2.Latitude = 37.114556
	parkLoc2.Longitude = 27.347444

	var parkLoc3 WasteLibrary.LocationType
	parkLoc3.LocationName = "KONACIK"
	parkLoc3.ZoneRadius = 50
	parkLoc3.Latitude = 37.058611
	parkLoc3.Longitude = 27.383417

	var parkLoc4 WasteLibrary.LocationType
	parkLoc4.LocationName = "MUMCULAR 1"
	parkLoc4.ZoneRadius = 50
	parkLoc4.Latitude = 37.106806
	parkLoc4.Longitude = 27.661806

	var parkLoc5 WasteLibrary.LocationType
	parkLoc5.LocationName = "MUMCULAR 2"
	parkLoc5.ZoneRadius = 50
	parkLoc5.Latitude = 36.997361
	parkLoc5.Longitude = 27.516944

	currentAdminConfig.ParkLocations = append(currentAdminConfig.ParkLocations, parkLoc1, parkLoc2, parkLoc3, parkLoc4, parkLoc5)

	var currentLocalConfig WasteLibrary.LocalConfigType
	currentLocalConfig.New()
	currentLocalConfig.CustomerId = currentData.CustomerId

	var currentCustomers WasteLibrary.CustomersType
	resultVal = currentCustomers.GetByRedis()
	if resultVal.Result == WasteLibrary.RESULT_OK {

	} else {
		currentCustomers = WasteLibrary.CustomersType{
			Customers: make(map[string]float64),
		}
	}
	currentCustomers.Customers[currentData.ToIdString()] = currentData.CustomerId
	resultVal = currentCustomers.SaveToRedis()
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE

		return resultVal
	}

	resultVal = currentCustomerRfidDevices.SaveToRedis()
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE

		return resultVal
	}
	resultVal = currentCustomerRfidDeviceViews.SaveToRedis()
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE

		return resultVal
	}

	resultVal = currentCustomerRecyDevices.SaveToRedis()
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE

		return resultVal
	}

	resultVal = currentCustomerRecyDeviceViews.SaveToRedis()
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE

		return resultVal
	}

	resultVal = currentCustomerUltDevices.SaveToRedis()
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE

		return resultVal
	}

	resultVal = currentCustomerUltDeviceViews.SaveToRedis()
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE

		return resultVal
	}

	resultVal = currentCustomerUsers.SaveToRedis()
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE

		return resultVal
	}
	resultVal = currentCustomerTags.SaveToRedis()
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE

		return resultVal
	}

	resultVal = currentCustomerTagViews.SaveToRedis()
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE

		return resultVal
	}

	resultVal = currentCustomerConfig.SaveToRedis()
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE

		return resultVal
	}
	resultVal = currentAdminConfig.SaveToRedis()
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE

		return resultVal
	}
	resultVal = currentLocalConfig.SaveToRedis()
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE

		return resultVal
	}

	return resultVal

}

func logApp(w http.ResponseWriter, req *http.Request) {

	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL

	if err := req.ParseForm(); err != nil {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_HTTP_PARSE
		w.Write(resultVal.ToByte())

		WasteLibrary.LogErr(err)
		return
	}

	appTypes := req.URL.Query()["appType"]
	opTypes := req.URL.Query()["opType"]
	resultVal = WasteLibrary.PublishRedisForStoreApi(WasteLibrary.REDIS_APP_LOG_CHANNEL, appTypes[0], opTypes[0])

	w.Write(resultVal.ToByte())

}

func setCustomer(w http.ResponseWriter, req *http.Request) {

	if WasteLibrary.AllowCors {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,access-control-allow-origin, access-control-allow-headers")
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

	var linkCustomer WasteLibrary.CustomerType
	resultVal = linkCustomer.GetByRedisByLink(req.Host)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_NOTFOUND
		w.Write(resultVal.ToByte())

		return
	}

	resultVal = checkAuth(req.Header, linkCustomer.ToIdString())

	if resultVal.Result != WasteLibrary.RESULT_OK {
		w.Write(resultVal.ToByte())

		return
	}

	var currentData WasteLibrary.CustomerType
	currentData.StringToType(req.FormValue(WasteLibrary.HTTP_DATA))
	var oldCustomer WasteLibrary.CustomerType
	var isCustomerExist = false
	if currentData.CustomerId != 0 {
		isCustomerExist = true
	}

	if currentData.CustomerId == 1 {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_NOTFOUND
		w.Write(resultVal.ToByte())
		return
	}

	if currentData.CustomerId != 0 {
		oldCustomer.CustomerId = currentData.CustomerId
		resultVal = oldCustomer.GetByRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_NOTFOUND
			w.Write(resultVal.ToByte())
			return
		}
	}

	resultVal = currentData.SaveToDb()
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
		w.Write(resultVal.ToByte())

		return
	}

	resultVal = currentData.SaveToRedis()
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
		w.Write(resultVal.ToByte())

		return
	}
	resultVal = currentData.SaveToRedisLink()
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
		w.Write(resultVal.ToByte())

		return
	}
	if isCustomerExist {
		if oldCustomer.CustomerLink != currentData.CustomerLink {
			resultVal = WasteLibrary.DeleteRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_LINK, oldCustomer.CustomerLink)
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
				w.Write(resultVal.ToByte())

				return
			}
		}
	} else {
		resultVal = createCustomerPart(currentData)
	}

	resultVal = WasteLibrary.PublishRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_CHANNEL+currentData.ToIdString(), WasteLibrary.DATATYPE_CUSTOMER, currentData.ToString())

	w.Write(resultVal.ToByte())

}

func getCustomers(w http.ResponseWriter, req *http.Request) {

	if WasteLibrary.AllowCors {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,access-control-allow-origin, access-control-allow-headers")
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

	var linkCustomer WasteLibrary.CustomerType
	resultVal = linkCustomer.GetByRedisByLink(req.Host)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_NOTFOUND
		w.Write(resultVal.ToByte())

		return
	}

	resultVal = checkAuth(req.Header, linkCustomer.ToIdString())

	if resultVal.Result != WasteLibrary.RESULT_OK {
		w.Write(resultVal.ToByte())

		return
	}

	var customers WasteLibrary.CustomersType
	resultVal = customers.GetByRedis()
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMERS_NOTFOUND
		w.Write(resultVal.ToByte())

		return
	}

	var customersList WasteLibrary.CustomersListType
	customersList.New()
	for _, customerId := range customers.Customers {

		if customerId != 0 {
			var currentCustomer WasteLibrary.CustomerType
			currentCustomer.CustomerId = customerId
			resultVal = currentCustomer.GetByRedis()
			if resultVal.Result == WasteLibrary.RESULT_OK {
				customersList.Customers[currentCustomer.ToIdString()] = currentCustomer
			}
		}
	}

	resultVal.Result = WasteLibrary.RESULT_OK
	resultVal.Retval = customersList.ToString()
	w.Write(resultVal.ToByte())

}

func getCustomer(w http.ResponseWriter, req *http.Request) {

	if WasteLibrary.AllowCors {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,access-control-allow-origin, access-control-allow-headers")
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

	var linkCustomer WasteLibrary.CustomerType
	resultVal = linkCustomer.GetByRedisByLink(req.Host)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_NOTFOUND
		w.Write(resultVal.ToByte())

		return
	}

	resultVal = checkAuth(req.Header, linkCustomer.ToIdString())

	if resultVal.Result != WasteLibrary.RESULT_OK {
		w.Write(resultVal.ToByte())

		return
	}

	var currentData WasteLibrary.CustomerType
	currentData.StringToType(req.FormValue(WasteLibrary.HTTP_DATA))
	resultVal = currentData.GetByRedis()
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_NOTFOUND
		w.Write(resultVal.ToByte())

		return
	}

	w.Write(resultVal.ToByte())

}

func setDevice(w http.ResponseWriter, req *http.Request) {

	if WasteLibrary.AllowCors {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,access-control-allow-origin, access-control-allow-headers")
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

	var linkCustomer WasteLibrary.CustomerType
	resultVal = linkCustomer.GetByRedisByLink(req.Host)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_NOTFOUND
		w.Write(resultVal.ToByte())

		return
	}

	resultVal = checkAuth(req.Header, linkCustomer.ToIdString())

	if resultVal.Result != WasteLibrary.RESULT_OK {
		w.Write(resultVal.ToByte())

		return
	}

	var currentHttpHeader WasteLibrary.HttpClientHeaderType
	currentHttpHeader.StringToType(req.FormValue(WasteLibrary.HTTP_HEADER))
	if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_RFID {
		var currentDeviceMain WasteLibrary.RfidDeviceMainType
		currentDeviceMain.StringToType(req.FormValue(WasteLibrary.HTTP_DATA))
		var currentOldData WasteLibrary.RfidDeviceType
		currentOldData.DeviceId = currentDeviceMain.DeviceId
		currentOldData.GetByRedis("0")
		resultVal = currentDeviceMain.SaveToDb()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		resultVal = currentDeviceMain.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		if currentDeviceMain.CustomerId != currentOldData.DeviceMain.CustomerId {
			var customerDevices WasteLibrary.CustomerRfidDevicesType
			customerDevices.CustomerId = currentDeviceMain.CustomerId
			resultVal = customerDevices.GetByRedis("0")
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_GET
				w.Write(resultVal.ToByte())

				return
			}

			customerDevices.Devices[currentDeviceMain.ToIdString()] = currentDeviceMain.DeviceId
			resultVal = customerDevices.SaveToRedis()
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
				w.Write(resultVal.ToByte())

				return
			}

			var oldCustomerDevices WasteLibrary.CustomerRfidDevicesType
			oldCustomerDevices.CustomerId = currentOldData.DeviceMain.CustomerId
			resultVal = oldCustomerDevices.GetByRedis("0")
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_GET
				w.Write(resultVal.ToByte())

				return
			}

			delete(oldCustomerDevices.Devices, currentDeviceMain.ToIdString())
			resultVal = oldCustomerDevices.SaveToRedis()
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
				w.Write(resultVal.ToByte())

				return
			}
		}

		resultVal = WasteLibrary.PublishRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_CHANNEL+currentDeviceMain.ToCustomerIdString(), WasteLibrary.DATATYPE_RFID_MAIN, currentDeviceMain.ToString())

	} else if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_ULT {
		var currentDeviceMain WasteLibrary.UltDeviceMainType
		currentDeviceMain.StringToType(req.FormValue(WasteLibrary.HTTP_DATA))
		var currentOldData WasteLibrary.UltDeviceType
		currentOldData.DeviceId = currentDeviceMain.DeviceId
		currentOldData.GetByRedis("0")
		resultVal = currentDeviceMain.SaveToDb()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		resultVal = currentDeviceMain.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		if currentDeviceMain.CustomerId != currentOldData.DeviceMain.CustomerId {
			var customerDevices WasteLibrary.CustomerUltDevicesType
			customerDevices.CustomerId = currentDeviceMain.CustomerId
			resultVal = customerDevices.GetByRedis("0")
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_GET
				w.Write(resultVal.ToByte())

				return
			}

			customerDevices.Devices[currentDeviceMain.ToIdString()] = currentDeviceMain.DeviceId
			resultVal = customerDevices.SaveToRedis()
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
				w.Write(resultVal.ToByte())

				return
			}
			var oldCustomerDevices WasteLibrary.CustomerUltDevicesType
			oldCustomerDevices.CustomerId = currentOldData.DeviceMain.CustomerId
			resultVal = oldCustomerDevices.GetByRedis("0")
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_GET
				w.Write(resultVal.ToByte())

				return
			}

			delete(oldCustomerDevices.Devices, currentDeviceMain.ToIdString())
			resultVal = oldCustomerDevices.SaveToRedis()
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
				w.Write(resultVal.ToByte())

				return
			}
		}
		resultVal = WasteLibrary.PublishRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_CHANNEL+currentDeviceMain.ToCustomerIdString(), WasteLibrary.DATATYPE_ULT_MAIN, currentDeviceMain.ToString())

	} else if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_RECY {
		var currentDeviceMain WasteLibrary.RecyDeviceMainType
		currentDeviceMain.StringToType(req.FormValue(WasteLibrary.HTTP_DATA))
		var currentOldData WasteLibrary.RecyDeviceType
		currentOldData.DeviceId = currentDeviceMain.DeviceId
		currentOldData.GetByRedis("0")
		resultVal = currentDeviceMain.SaveToDb()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		resultVal = currentDeviceMain.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		if currentDeviceMain.CustomerId != currentOldData.DeviceMain.CustomerId {
			var customerDevices WasteLibrary.CustomerRecyDevicesType
			customerDevices.CustomerId = currentDeviceMain.CustomerId
			resultVal = customerDevices.GetByRedis("0")
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_GET
				w.Write(resultVal.ToByte())

				return
			}

			customerDevices.Devices[currentDeviceMain.ToIdString()] = currentDeviceMain.DeviceId
			resultVal = customerDevices.SaveToRedis()
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
				w.Write(resultVal.ToByte())

				return
			}

			var oldCustomerDevices WasteLibrary.CustomerRecyDevicesType
			oldCustomerDevices.CustomerId = currentOldData.DeviceMain.CustomerId
			resultVal = oldCustomerDevices.GetByRedis("0")
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_GET
				w.Write(resultVal.ToByte())

				return
			}

			delete(oldCustomerDevices.Devices, currentDeviceMain.ToIdString())
			resultVal = oldCustomerDevices.SaveToRedis()
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
				w.Write(resultVal.ToByte())

				return
			}
		}
		resultVal = WasteLibrary.PublishRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_CHANNEL+currentDeviceMain.ToCustomerIdString(), WasteLibrary.DATATYPE_RECY_MAIN, currentDeviceMain.ToString())

	} else {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_DEVICE_NOTFOUND
	}

	w.Write(resultVal.ToByte())

}

func getDevices(w http.ResponseWriter, req *http.Request) {

	if WasteLibrary.AllowCors {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,access-control-allow-origin, access-control-allow-headers")
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

	var linkCustomer WasteLibrary.CustomerType
	resultVal = linkCustomer.GetByRedisByLink(req.Host)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_NOTFOUND
		w.Write(resultVal.ToByte())

		return
	}

	resultVal = checkAuth(req.Header, linkCustomer.ToIdString())

	if resultVal.Result != WasteLibrary.RESULT_OK {
		w.Write(resultVal.ToByte())

		return
	}
	var currentHttpHeader WasteLibrary.HttpClientHeaderType
	currentHttpHeader.StringToType(req.FormValue(WasteLibrary.HTTP_HEADER))
	if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_RFID {
		var customers WasteLibrary.CustomersType
		resultVal = customers.GetByRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMERS_NOTFOUND
			w.Write(resultVal.ToByte())

			return
		}

		var customerDevicesList WasteLibrary.RfidDevicesListType
		customerDevicesList.Devices = make(map[string]WasteLibrary.RfidDeviceType)
		for _, customerId := range customers.Customers {
			var customerDevices WasteLibrary.CustomerRfidDevicesType
			customerDevices.CustomerId = customerId
			resultVal = customerDevices.GetByRedis("0")
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_DEVICES_NOTFOUND
				w.Write(resultVal.ToByte())

				return
			}

			for _, deviceId := range customerDevices.Devices {

				if deviceId != 0 {
					var currentDevice WasteLibrary.RfidDeviceType
					currentDevice.New()
					currentDevice.DeviceId = deviceId
					resultVal = currentDevice.GetByRedis("0")
					if resultVal.Result == WasteLibrary.RESULT_OK {
						customerDevicesList.Devices[currentDevice.ToIdString()] = currentDevice
					}
				}
			}

			resultVal.Result = WasteLibrary.RESULT_OK
			resultVal.Retval = customerDevicesList.ToString()
		}
	} else if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_ULT {
		var customers WasteLibrary.CustomersType
		resultVal = customers.GetByRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMERS_NOTFOUND
			w.Write(resultVal.ToByte())

			return
		}

		var customersList WasteLibrary.CustomersListType
		customersList.New()
		var customerDevicesList WasteLibrary.UltDevicesListType
		customerDevicesList.Devices = make(map[string]WasteLibrary.UltDeviceType)
		for _, customerId := range customers.Customers {
			var customerDevices WasteLibrary.CustomerUltDevicesType
			customerDevices.CustomerId = customerId
			resultVal = customerDevices.GetByRedis("0")
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_DEVICES_NOTFOUND
				w.Write(resultVal.ToByte())

				return
			}

			for _, deviceId := range customerDevices.Devices {

				if deviceId != 0 {
					var currentDevice WasteLibrary.UltDeviceType
					currentDevice.New()
					currentDevice.DeviceId = deviceId
					resultVal = currentDevice.GetByRedis("0")
					if resultVal.Result == WasteLibrary.RESULT_OK {
						customerDevicesList.Devices[currentDevice.ToIdString()] = currentDevice
					}
				}
			}

			resultVal.Result = WasteLibrary.RESULT_OK
			resultVal.Retval = customerDevicesList.ToString()
		}
	} else if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_RECY {
		var customers WasteLibrary.CustomersType
		resultVal = customers.GetByRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMERS_NOTFOUND
			w.Write(resultVal.ToByte())

			return
		}

		var customersList WasteLibrary.CustomersListType
		customersList.New()
		var customerDevicesList WasteLibrary.RecyDevicesListType
		customerDevicesList.Devices = make(map[string]WasteLibrary.RecyDeviceType)
		for _, customerId := range customers.Customers {
			var customerDevices WasteLibrary.CustomerRecyDevicesType
			customerDevices.CustomerId = customerId
			resultVal = customerDevices.GetByRedis("0")
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_DEVICES_NOTFOUND
				w.Write(resultVal.ToByte())

				return
			}

			for _, deviceId := range customerDevices.Devices {

				if deviceId != 0 {
					var currentDevice WasteLibrary.RecyDeviceType
					currentDevice.New()
					currentDevice.DeviceId = deviceId
					resultVal = currentDevice.GetByRedis("0")
					if resultVal.Result == WasteLibrary.RESULT_OK {
						customerDevicesList.Devices[currentDevice.ToIdString()] = currentDevice
					}
				}
			}

			resultVal.Result = WasteLibrary.RESULT_OK
			resultVal.Retval = customerDevicesList.ToString()
		}
	} else {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_DEVICES_NOTFOUND
	}
	w.Write(resultVal.ToByte())

}

func getDevice(w http.ResponseWriter, req *http.Request) {

	if WasteLibrary.AllowCors {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,access-control-allow-origin, access-control-allow-headers")
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

	var linkCustomer WasteLibrary.CustomerType
	resultVal = linkCustomer.GetByRedisByLink(req.Host)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_NOTFOUND
		w.Write(resultVal.ToByte())

		return
	}

	resultVal = checkAuth(req.Header, linkCustomer.ToIdString())

	if resultVal.Result != WasteLibrary.RESULT_OK {
		w.Write(resultVal.ToByte())

		return
	}

	var currentHttpHeader WasteLibrary.HttpClientHeaderType
	currentHttpHeader.StringToType(req.FormValue(WasteLibrary.HTTP_HEADER))
	if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_RFID {
		var currentDevice WasteLibrary.RfidDeviceType
		currentDevice.StringToType(req.FormValue(WasteLibrary.HTTP_DATA))
		resultVal = currentDevice.GetByRedis("0")
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DEVICES_NOTFOUND
			w.Write(resultVal.ToByte())

			return
		}

		resultVal.Result = WasteLibrary.RESULT_OK
		resultVal.Retval = currentDevice.ToString()
	} else if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_ULT {
		var currentDevice WasteLibrary.UltDeviceType
		currentDevice.StringToType(req.FormValue(WasteLibrary.HTTP_DATA))
		resultVal = currentDevice.GetByRedis("0")
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DEVICES_NOTFOUND
			w.Write(resultVal.ToByte())

			return
		}

		resultVal.Result = WasteLibrary.RESULT_OK
		resultVal.Retval = currentDevice.ToString()
	} else if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_RECY {
		var currentDevice WasteLibrary.RecyDeviceType
		currentDevice.StringToType(req.FormValue(WasteLibrary.HTTP_DATA))
		resultVal = currentDevice.GetByRedis("0")
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DEVICES_NOTFOUND
			w.Write(resultVal.ToByte())

			return
		}

		resultVal.Result = WasteLibrary.RESULT_OK
		resultVal.Retval = currentDevice.ToString()
	} else {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_DEVICES_NOTFOUND
	}
	w.Write(resultVal.ToByte())

}

func getConfig(w http.ResponseWriter, req *http.Request) {

	if WasteLibrary.AllowCors {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,access-control-allow-origin, access-control-allow-headers")
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
	var linkCustomer WasteLibrary.CustomerType
	resultVal = linkCustomer.GetByRedisByLink(req.Host)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_NOTFOUND
		w.Write(resultVal.ToByte())

		return
	}

	resultVal = checkAuth(req.Header, linkCustomer.ToIdString())

	if resultVal.Result != WasteLibrary.RESULT_OK {
		w.Write(resultVal.ToByte())

		return
	}

	var currentConfig WasteLibrary.CustomerConfigType
	currentConfig.StringToType(req.FormValue(WasteLibrary.HTTP_DATA))
	resultVal = currentConfig.GetByRedis()
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_CUSTOMERCONFIG_NOTFOUND
		w.Write(resultVal.ToByte())

		return
	}

	w.Write(resultVal.ToByte())

}

func setConfig(w http.ResponseWriter, req *http.Request) {

	if WasteLibrary.AllowCors {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,access-control-allow-origin, access-control-allow-headers")
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
	var linkCustomer WasteLibrary.CustomerType
	resultVal = linkCustomer.GetByRedisByLink(req.Host)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_NOTFOUND
		w.Write(resultVal.ToByte())

		return
	}

	resultVal = checkAuth(req.Header, linkCustomer.ToIdString())

	if resultVal.Result != WasteLibrary.RESULT_OK {
		w.Write(resultVal.ToByte())

		return
	}

	var currentData WasteLibrary.CustomerConfigType
	currentData.StringToType(req.FormValue(WasteLibrary.HTTP_DATA))
	resultVal = currentData.SaveToRedis()
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
		w.Write(resultVal.ToByte())

		return
	}

	w.Write(resultVal.ToByte())

}

func checkAuth(header http.Header, customerId string) WasteLibrary.ResultType {
	var resultVal WasteLibrary.ResultType
	resultVal = WasteLibrary.CheckAuth(header, customerId, []string{WasteLibrary.USER_ROLE_ADMIN})

	if resultVal.Result == WasteLibrary.RESULT_OK {
		var currentCustomer WasteLibrary.CustomerType
		currentCustomer.CustomerId = WasteLibrary.StringIdToFloat64(customerId)
		resultVal = currentCustomer.GetByRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_USER_AUTH
			return resultVal
		}
		if currentCustomer.CustomerName != "AFATEK" {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_USER_AUTH
			return resultVal
		}
	}
	return resultVal

}
