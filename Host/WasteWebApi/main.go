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
	http.HandleFunc("/getCustomer", getCustomer)
	http.HandleFunc("/getDevice", getDevice)
	http.HandleFunc("/getDevices", getDevices)
	http.HandleFunc("/getConfig", getConfig)
	http.HandleFunc("/getTags", getTags)
	http.HandleFunc("/getTag", getTag)
	http.HandleFunc("/setDevice", setDevice)
	http.HandleFunc("/setUser", setUser)
	http.HandleFunc("/setConfig", setConfig)
	http.HandleFunc("/getUser", getUser)
	http.HandleFunc("/getUsers", getUsers)

	http.ListenAndServe(":80", nil)
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

	var currentAdminConfig WasteLibrary.AdminConfigType
	currentAdminConfig.CustomerId = linkCustomer.CustomerId
	currentAdminConfig.GetByRedis()
	if currentAdminConfig.WebUIPrivate == WasteLibrary.STATU_ACTIVE {
		resultVal = checkAuth(req.Header, linkCustomer.ToIdString())

		if resultVal.Result != WasteLibrary.RESULT_OK {
			w.Write(resultVal.ToByte())

			return
		}
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

	var currentAdminConfig WasteLibrary.AdminConfigType
	currentAdminConfig.CustomerId = linkCustomer.CustomerId
	currentAdminConfig.GetByRedis()
	if currentAdminConfig.WebUIPrivate == WasteLibrary.STATU_ACTIVE {
		resultVal = checkAuth(req.Header, linkCustomer.ToIdString())

		if resultVal.Result != WasteLibrary.RESULT_OK {
			w.Write(resultVal.ToByte())

			return
		}
	}

	var currentHttpHeader WasteLibrary.HttpClientHeaderType
	currentHttpHeader.StringToType(req.FormValue(WasteLibrary.HTTP_HEADER))
	dbIndex := WasteLibrary.GetDbIndexByDate(currentHttpHeader.Time)
	if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_RFID {
		var currentData WasteLibrary.RfidDeviceType
		currentData.StringToType(req.FormValue(WasteLibrary.HTTP_DATA))
		currentData.GetByRedis(dbIndex)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DEVICE_NOTFOUND
			w.Write(resultVal.ToByte())

			return
		}
		if currentData.DeviceMain.CustomerId == linkCustomer.CustomerId {
			resultVal.Result = WasteLibrary.RESULT_OK
			resultVal.Retval = currentData.ToString()
			w.Write(resultVal.ToByte())

		} else {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DEVICE_NOTFOUND
			w.Write(resultVal.ToByte())

		}
	} else if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_ULT {
		var currentData WasteLibrary.UltDeviceType
		currentData.StringToType(req.FormValue(WasteLibrary.HTTP_DATA))
		currentData.GetByRedis(dbIndex)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DEVICE_NOTFOUND
			w.Write(resultVal.ToByte())

			return
		}
		if currentData.DeviceMain.CustomerId == linkCustomer.CustomerId {
			resultVal.Result = WasteLibrary.RESULT_OK
			resultVal.Retval = currentData.ToString()
			w.Write(resultVal.ToByte())

		} else {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DEVICE_NOTFOUND
			w.Write(resultVal.ToByte())

		}
	} else if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_RECY {
		var currentData WasteLibrary.RecyDeviceType
		currentData.StringToType(req.FormValue(WasteLibrary.HTTP_DATA))
		currentData.GetByRedis(dbIndex)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DEVICE_NOTFOUND
			w.Write(resultVal.ToByte())

			return
		}
		if currentData.DeviceMain.CustomerId == linkCustomer.CustomerId {
			resultVal.Result = WasteLibrary.RESULT_OK
			resultVal.Retval = currentData.ToString()
			w.Write(resultVal.ToByte())

		} else {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DEVICE_NOTFOUND
			w.Write(resultVal.ToByte())

		}
	} else {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_DEVICE_NOTFOUND
		w.Write(resultVal.ToByte())

	}

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

	var currentAdminConfig WasteLibrary.AdminConfigType
	currentAdminConfig.CustomerId = linkCustomer.CustomerId
	currentAdminConfig.GetByRedis()
	if currentAdminConfig.WebUIPrivate == WasteLibrary.STATU_ACTIVE {
		resultVal = checkAuth(req.Header, linkCustomer.ToIdString())

		if resultVal.Result != WasteLibrary.RESULT_OK {
			w.Write(resultVal.ToByte())

			return
		}
	}

	var currentHttpHeader WasteLibrary.HttpClientHeaderType
	currentHttpHeader.StringToType(req.FormValue(WasteLibrary.HTTP_HEADER))
	dbIndex := WasteLibrary.GetDbIndexByDate(currentHttpHeader.Time)
	if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_RFID {
		var customerDevicesList WasteLibrary.CustomerRfidDevicesViewListType
		customerDevicesList.CustomerId = linkCustomer.CustomerId
		resultVal = customerDevicesList.GetByRedisByReel(dbIndex)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_DEVICES_NOTFOUND
			w.Write(resultVal.ToByte())

			return
		}

		resultVal.Result = WasteLibrary.RESULT_OK
		resultVal.Retval = customerDevicesList.ToString()
		w.Write(resultVal.ToByte())
		return
	} else if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_ULT {
		var customerDevicesList WasteLibrary.CustomerUltDevicesViewListType
		customerDevicesList.CustomerId = linkCustomer.CustomerId
		resultVal = customerDevicesList.GetByRedisByReel(dbIndex)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_DEVICES_NOTFOUND
			w.Write(resultVal.ToByte())

			return
		}

		resultVal.Result = WasteLibrary.RESULT_OK
		resultVal.Retval = customerDevicesList.ToString()
		w.Write(resultVal.ToByte())
		return
	} else if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_RECY {
		var customerDevicesList WasteLibrary.CustomerRecyDevicesViewListType
		customerDevicesList.CustomerId = linkCustomer.CustomerId
		resultVal = customerDevicesList.GetByRedisByReel(dbIndex)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_DEVICES_NOTFOUND
			w.Write(resultVal.ToByte())

			return
		}

		resultVal.Result = WasteLibrary.RESULT_OK
		resultVal.Retval = customerDevicesList.ToString()
		w.Write(resultVal.ToByte())
		return
	} else {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_DEVICE_NOTFOUND
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

	var currentAdminConfig WasteLibrary.AdminConfigType
	currentAdminConfig.CustomerId = linkCustomer.CustomerId
	currentAdminConfig.GetByRedis()
	if currentAdminConfig.WebUIPrivate == WasteLibrary.STATU_ACTIVE {
		resultVal = checkAuth(req.Header, linkCustomer.ToIdString())

		if resultVal.Result != WasteLibrary.RESULT_OK {
			w.Write(resultVal.ToByte())

			return
		}
	}

	var currentHttpHeader WasteLibrary.HttpClientHeaderType
	currentHttpHeader.StringToType(req.FormValue(WasteLibrary.HTTP_HEADER))
	if currentHttpHeader.DataType == WasteLibrary.DATATYPE_ADMINCONFIG {
		var customerConfig WasteLibrary.AdminConfigType
		customerConfig.CustomerId = linkCustomer.CustomerId
		resultVal = customerConfig.GetByRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_ADMINCONFIG_NOTFOUND
			w.Write(resultVal.ToByte())

			return
		}
	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_LOCALCONFIG {
		var customerConfig WasteLibrary.LocalConfigType
		customerConfig.CustomerId = linkCustomer.CustomerId
		resultVal = customerConfig.GetByRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_LOCALCONFIG_NOTFOUND
			w.Write(resultVal.ToByte())

			return
		}
	} else {
		resultVal.Result = WasteLibrary.RESULT_FAIL
	}

	w.Write(resultVal.ToByte())

}

func getTags(w http.ResponseWriter, req *http.Request) {

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

	var currentAdminConfig WasteLibrary.AdminConfigType
	currentAdminConfig.CustomerId = linkCustomer.CustomerId
	currentAdminConfig.GetByRedis()
	if currentAdminConfig.WebUIPrivate == WasteLibrary.STATU_ACTIVE {
		resultVal = checkAuth(req.Header, linkCustomer.ToIdString())

		if resultVal.Result != WasteLibrary.RESULT_OK {
			w.Write(resultVal.ToByte())

			return
		}
	}

	var currentHttpHeader WasteLibrary.HttpClientHeaderType
	currentHttpHeader.StringToType(req.FormValue(WasteLibrary.HTTP_HEADER))
	dbIndex := WasteLibrary.GetDbIndexByDate(currentHttpHeader.Time)
	var customerTagsList WasteLibrary.CustomerTagsViewListType
	customerTagsList.CustomerId = linkCustomer.CustomerId
	resultVal = customerTagsList.GetByRedisByReel(dbIndex)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_TAGS_NOTFOUND
		w.Write(resultVal.ToByte())

		return
	}

	resultVal.Result = WasteLibrary.RESULT_OK
	resultVal.Retval = customerTagsList.ToString()
	w.Write(resultVal.ToByte())

}

func getTag(w http.ResponseWriter, req *http.Request) {

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

	var currentAdminConfig WasteLibrary.AdminConfigType
	currentAdminConfig.CustomerId = linkCustomer.CustomerId
	currentAdminConfig.GetByRedis()
	if currentAdminConfig.WebUIPrivate == WasteLibrary.STATU_ACTIVE {
		resultVal = checkAuth(req.Header, linkCustomer.ToIdString())

		if resultVal.Result != WasteLibrary.RESULT_OK {
			w.Write(resultVal.ToByte())

			return
		}
	}

	var currentHttpHeader WasteLibrary.HttpClientHeaderType
	currentHttpHeader.StringToType(req.FormValue(WasteLibrary.HTTP_HEADER))
	dbIndex := WasteLibrary.GetDbIndexByDate(currentHttpHeader.Time)
	var currentData WasteLibrary.TagType
	currentData.StringToType(req.FormValue(WasteLibrary.HTTP_DATA))
	resultVal = currentData.GetByRedis(dbIndex)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_TAG_NOTFOUND
		w.Write(resultVal.ToByte())

		return
	}
	if currentData.TagMain.CustomerId == linkCustomer.CustomerId {
		resultVal.Result = WasteLibrary.RESULT_OK
		resultVal.Retval = currentData.ToString()
		w.Write(resultVal.ToByte())

	} else {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_TAG_NOTFOUND
		w.Write(resultVal.ToByte())

	}
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
		var currentData WasteLibrary.RfidDeviceType
		currentData.StringToType(req.FormValue(WasteLibrary.HTTP_DATA))
		var currentOldData WasteLibrary.RfidDeviceType
		currentOldData.DeviceId = currentData.DeviceId
		currentOldData.GetByRedis("0")
		currentData.DeviceMain.CustomerId = currentOldData.DeviceMain.CustomerId
		if currentData.DeviceMain.CustomerId != linkCustomer.CustomerId {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DEVICE_NOTFOUND
			w.Write(resultVal.ToByte())
			return
		}

		//DeviceBase
		currentData.DeviceBase.DeviceId = currentData.DeviceId
		resultVal = currentData.DeviceBase.SaveToDb()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		resultVal = currentData.DeviceBase.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		resultVal = WasteLibrary.PublishRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_CHANNEL+currentData.DeviceMain.ToCustomerIdString(), WasteLibrary.DATATYPE_RFID_BASE, currentData.DeviceBase.ToString())

		//DeviceDetail
		currentData.DeviceDetail.DeviceId = currentData.DeviceId
		resultVal = currentData.DeviceDetail.SaveToDb()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		resultVal = currentData.DeviceDetail.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		resultVal = WasteLibrary.PublishRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_CHANNEL+currentData.DeviceMain.ToCustomerIdString(), WasteLibrary.DATATYPE_RFID_DETAIL, currentData.DeviceDetail.ToString())

	} else if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_ULT {
		var currentData WasteLibrary.UltDeviceType
		currentData.StringToType(req.FormValue(WasteLibrary.HTTP_DATA))
		var currentOldData WasteLibrary.UltDeviceType
		currentOldData.DeviceId = currentData.DeviceId
		currentOldData.GetByRedis("0")
		currentData.DeviceMain.CustomerId = currentOldData.DeviceMain.CustomerId
		if currentData.DeviceMain.CustomerId != linkCustomer.CustomerId {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DEVICE_NOTFOUND
			w.Write(resultVal.ToByte())

			return
		}

		//DeviceBase
		currentData.DeviceBase.DeviceId = currentData.DeviceId
		resultVal = currentData.DeviceBase.SaveToDb()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		resultVal = currentData.DeviceBase.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		resultVal = WasteLibrary.PublishRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_CHANNEL+currentData.DeviceMain.ToCustomerIdString(), WasteLibrary.DATATYPE_ULT_BASE, currentData.DeviceBase.ToString())

	} else if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_RECY {
		var currentData WasteLibrary.RecyDeviceType
		currentData.StringToType(req.FormValue(WasteLibrary.HTTP_DATA))
		var currentOldData WasteLibrary.RecyDeviceType
		currentOldData.DeviceId = currentData.DeviceId
		currentOldData.GetByRedis("0")
		currentData.DeviceMain.CustomerId = currentOldData.DeviceMain.CustomerId
		if currentData.DeviceMain.CustomerId != linkCustomer.CustomerId {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DEVICE_NOTFOUND
			w.Write(resultVal.ToByte())

			return
		}

		//DeviceBase
		currentData.DeviceBase.DeviceId = currentData.DeviceId
		resultVal = currentData.DeviceBase.SaveToDb()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		resultVal = currentData.DeviceBase.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		resultVal = WasteLibrary.PublishRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_CHANNEL+currentData.DeviceMain.ToCustomerIdString(), WasteLibrary.DATATYPE_RECY_BASE, currentData.DeviceBase.ToString())

	} else {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_DEVICE_NOTFOUND
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

	var currentHttpHeader WasteLibrary.HttpClientHeaderType
	currentHttpHeader.StringToType(req.FormValue(WasteLibrary.HTTP_HEADER))
	if currentHttpHeader.DataType == WasteLibrary.DATATYPE_ADMINCONFIG {
		var currentData WasteLibrary.AdminConfigType
		currentData.StringToType(req.FormValue(WasteLibrary.HTTP_DATA))
		currentData.CustomerId = linkCustomer.CustomerId
		resultVal = currentData.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		resultVal = WasteLibrary.PublishRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_CHANNEL+currentData.ToIdString(), WasteLibrary.DATATYPE_ADMINCONFIG, currentData.ToString())
	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_LOCALCONFIG {
		var currentData WasteLibrary.LocalConfigType
		currentData.StringToType(req.FormValue(WasteLibrary.HTTP_DATA))
		currentData.CustomerId = linkCustomer.CustomerId
		resultVal = currentData.SaveToRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())
			return
		}

		resultVal = WasteLibrary.PublishRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_CHANNEL+currentData.ToIdString(), WasteLibrary.DATATYPE_LOCALCONFIG, currentData.ToString())

	} else {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_DATATYPE
		w.Write(resultVal.ToByte())

		return
	}

	w.Write(resultVal.ToByte())

}

func setUser(w http.ResponseWriter, req *http.Request) {

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

	var currentData WasteLibrary.UserType
	currentData.StringToType(req.FormValue(WasteLibrary.HTTP_DATA))
	var currentDbData WasteLibrary.UserType
	currentDbData.UserId = currentData.UserId
	resultVal = currentDbData.GetByRedis()
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_USER_AUTH
		w.Write(resultVal.ToByte())

		return
	}
	if currentDbData.CustomerId != linkCustomer.CustomerId {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_USER_AUTH
		w.Write(resultVal.ToByte())

		return
	}

	var currentCustomerUsers WasteLibrary.CustomerUsersType
	currentCustomerUsers.CustomerId = linkCustomer.CustomerId
	resultVal = currentCustomerUsers.GetByRedis()
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_NOTFOUND
		w.Write(resultVal.ToByte())

		return
	}

	for _, userId := range currentCustomerUsers.Users {
		if userId != 0 && userId != currentData.UserId {
			var inRedisUser WasteLibrary.UserType
			inRedisUser.UserId = userId
			resultVal = inRedisUser.GetByRedis()
			if resultVal.Result == WasteLibrary.RESULT_OK {

				if inRedisUser.Email == currentData.Email && inRedisUser.Email != "" {
					resultVal.Result = WasteLibrary.RESULT_FAIL
					resultVal.Retval = WasteLibrary.RESULT_ERROR_USER_AUTH
					w.Write(resultVal.ToByte())

					return
				}
			}
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
	w.Write(resultVal.ToByte())

}

func getUser(w http.ResponseWriter, req *http.Request) {

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

	var currentData WasteLibrary.UserType
	currentData.StringToType(req.FormValue(WasteLibrary.HTTP_DATA))
	resultVal = currentData.GetByRedis()
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_USER_AUTH
		w.Write(resultVal.ToByte())

		return
	}
	if currentData.CustomerId == linkCustomer.CustomerId {
		w.Write(resultVal.ToByte())

	} else {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_USER_AUTH
		w.Write(resultVal.ToByte())

	}
}

func getUsers(w http.ResponseWriter, req *http.Request) {

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

	var customerUsers WasteLibrary.CustomerUsersType
	customerUsers.CustomerId = linkCustomer.CustomerId
	resultVal = customerUsers.GetByRedis()
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_NOTFOUND
		w.Write(resultVal.ToByte())

		return
	}

	var customerUsersList WasteLibrary.CustomerUsersListType = WasteLibrary.CustomerUsersListType{
		CustomerId: linkCustomer.CustomerId,
		Users:      make(map[string]WasteLibrary.UserType),
	}
	for _, userId := range customerUsers.Users {

		if userId != 0 {
			var currentUser WasteLibrary.UserType
			currentUser.UserId = userId
			resultVal = currentUser.GetByRedis()
			if resultVal.Result == WasteLibrary.RESULT_OK {

				customerUsersList.Users[currentUser.ToIdString()] = currentUser
			}

		}
	}
	resultVal.Result = WasteLibrary.RESULT_OK
	resultVal.Retval = customerUsersList.ToString()

	w.Write(resultVal.ToByte())

}

func checkAuth(header http.Header, customerId string) WasteLibrary.ResultType {
	return WasteLibrary.CheckAuth(header, customerId, "ADMIN")

}
