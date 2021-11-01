package main

import (
	"fmt"
	"net/url"

	"github.com/devafatek/WasteLibrary"
)

func main() {

	var opType string = "USER"
	var opType2 string = "LOGIN_AUTH"
	var currentHeader WasteLibrary.HttpClientHeaderType
	currentHeader.New()
	var urlVal string = "afatek.aws.afatek.com.tr"
	var path1 string = "webapi"
	var path2 string = "getLink"
	data := url.Values{}
	var deviceId float64 = 0
	var customerId float64 = 1
	var userId float64 = 2
	var token = "MSMkMmEkMTAkMm9LY2dmSWlSRXc3TnhhSlplYmZNT1RrZmI3MmtMaVdXWnZ2eTczMlNCUVFlMGtjU1ZyWVc="

	if opType == "DEVICE" {
		if opType2 == "GET_RFIDDEVICE_WEB" {
			//TO DO
			//check
			path1 = "webapi"
			path2 = "getDevice"
			//currentHeader - HttpClientHeaderType *
			//currentHeader.DeviceType - DEVICE_TYPE_RFID *
			//HTTP_HEADER : currentHeader
			//
			//currentData - RfidDeviceType *
			//currentData.DevıceId - deviceId *
			//HTTP_DATA : currentData
			//
			//Retval : RfidDeviceType
			currentHeader.DeviceType = WasteLibrary.DEVICE_TYPE_RFID
			var currentData WasteLibrary.RfidDeviceType
			currentData.New()
			currentData.DeviceId = deviceId

			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
				WasteLibrary.HTTP_DATA:   {currentData.ToString()},
			}
		} else if opType2 == "GET_RECYDEVICE_WEB" {
			//TO DO
			//check
			path1 = "webapi"
			path2 = "getDevice"
			//currentHeader - HttpClientHeaderType *
			//currentHeader.DeviceType - DEVICE_TYPE_RECY *
			//HTTP_HEADER : currentHeader
			//
			//currentData - RecyDeviceType *
			//currentData.DevıceId - deviceId *
			//HTTP_DATA : currentData
			//
			//Retval : RecyDeviceType
			currentHeader.DeviceType = WasteLibrary.DEVICE_TYPE_RECY
			var currentData WasteLibrary.RecyDeviceType
			currentData.New()
			currentData.DeviceId = deviceId

			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
				WasteLibrary.HTTP_DATA:   {currentData.ToString()},
			}
		} else if opType2 == "GET_ULTDEVICE_WEB" {
			//TO DO
			//check
			path1 = "webapi"
			path2 = "getDevice"
			//currentHeader - HttpClientHeaderType *
			//currentHeader.DeviceType - DEVICE_TYPE_ULT *
			//HTTP_HEADER : currentHeader
			//
			//currentData - UltDeviceType *
			//currentData.DevıceId - deviceId *
			//HTTP_DATA : currentData
			//
			//Retval : UltDeviceType
			currentHeader.DeviceType = WasteLibrary.DEVICE_TYPE_ULT
			var currentData WasteLibrary.UltDeviceType
			currentData.New()
			currentData.DeviceId = deviceId

			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
				WasteLibrary.HTTP_DATA:   {currentData.ToString()},
			}
		} else if opType2 == "GET_RFIDDEVICE_AFATEK" {
			//TO DO
			//check
			path1 = "afatekapi"
			path2 = "getDevice"
			//currentHeader - HttpClientHeaderType *
			//currentHeader.DeviceType - DEVICE_TYPE_RFID *
			//currentHeader.Token - token *
			//HTTP_HEADER : currentHeader
			//
			//currentData - RfidDeviceType *
			//currentData.DevıceId - deviceId *
			//HTTP_DATA : currentData
			//
			//Retval : RfidDeviceType
			currentHeader.Token = token
			currentHeader.DeviceType = WasteLibrary.DEVICE_TYPE_RFID
			var currentData WasteLibrary.RfidDeviceType
			currentData.New()
			currentData.DeviceId = deviceId

			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
				WasteLibrary.HTTP_DATA:   {currentData.ToString()},
			}

		} else if opType2 == "GET_RECYDEVICE_AFATEK" {
			//TO DO
			//check
			path1 = "afatekapi"
			path2 = "getDevice"
			//currentHeader - HttpClientHeaderType *
			//currentHeader.DeviceType - DEVICE_TYPE_RECY *
			//currentHeader.Token - token *
			//HTTP_HEADER : currentHeader
			//
			//currentData - RecyDeviceType *
			//currentData.DevıceId - deviceId *
			//HTTP_DATA : currentData
			//
			//Retval : RecyDeviceType
			currentHeader.Token = token
			currentHeader.DeviceType = WasteLibrary.DEVICE_TYPE_RECY
			var currentData WasteLibrary.RecyDeviceType
			currentData.New()
			currentData.DeviceId = deviceId

			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
				WasteLibrary.HTTP_DATA:   {currentData.ToString()},
			}

		} else if opType2 == "GET_ULTDEVICE_AFATEK" {
			//TO DO
			//check
			path1 = "afatekapi"
			path2 = "getDevice"
			//currentHeader - HttpClientHeaderType *
			//currentHeader.DeviceType - DEVICE_TYPE_ULT *
			//currentHeader.Token - token *
			//HTTP_HEADER : currentHeader
			//
			//currentData - UltDeviceType *
			//currentData.DevıceId - deviceId *
			//HTTP_DATA : currentData
			//
			//Retval : UltDeviceType
			currentHeader.Token = token
			currentHeader.DeviceType = WasteLibrary.DEVICE_TYPE_ULT
			var currentData WasteLibrary.UltDeviceType
			currentData.New()
			currentData.DeviceId = deviceId

			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
				WasteLibrary.HTTP_DATA:   {currentData.ToString()},
			}
		} else if opType2 == "SET_RFIDDEVICE_AFATEK" {
			//TO DO
			//check
			path1 = "afatekapi"
			path2 = "setDevice"
			//currentHeader - HttpClientHeaderType *
			//currentHeader.DeviceType - DEVICE_TYPE_RFID *
			//currentHeader.Token - token *
			//HTTP_HEADER : currentHeader
			//
			//currentData - RfidDeviceType *
			//currentData.DevıceId - deviceId *
			//currentData.[All]
			//HTTP_DATA : currentData
			//
			//Retval : RfidDeviceType
			currentHeader.Token = token
			currentHeader.DeviceType = WasteLibrary.DEVICE_TYPE_RFID
			var currentData WasteLibrary.RfidDeviceType
			currentData.New()
			currentData.DeviceId = deviceId

			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
				WasteLibrary.HTTP_DATA:   {currentData.ToString()},
			}

		} else if opType2 == "SET_RECYDEVICE_AFATEK" {
			//TO DO
			//check
			path1 = "afatekapi"
			path2 = "setDevice"
			//currentHeader - HttpClientHeaderType *
			//currentHeader.DeviceType - DEVICE_TYPE_RECY *
			//currentHeader.Token - token *
			//HTTP_HEADER : currentHeader
			//
			//currentData - RecyDeviceType *
			//currentData.DevıceId - deviceId *
			//currentData.[All]
			//HTTP_DATA : currentData
			//
			//Retval : RecyDeviceType
			currentHeader.Token = token
			currentHeader.DeviceType = WasteLibrary.DEVICE_TYPE_RECY
			var currentData WasteLibrary.RecyDeviceType
			currentData.New()
			currentData.DeviceId = deviceId

			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
				WasteLibrary.HTTP_DATA:   {currentData.ToString()},
			}

		} else if opType2 == "SET_ULTDEVICE_AFATEK" {
			//TO DO
			//check
			path1 = "afatekapi"
			path2 = "setDevice"
			//currentHeader - HttpClientHeaderType *
			//currentHeader.DeviceType - DEVICE_TYPE_ULT *
			//currentHeader.Token - token *
			//HTTP_HEADER : currentHeader
			//
			//currentData - UltDeviceType *
			//currentData.DevıceId - deviceId *
			//currentData.[All]
			//HTTP_DATA : currentData
			//
			//Retval : UltDeviceType
			currentHeader.Token = token
			currentHeader.DeviceType = WasteLibrary.DEVICE_TYPE_ULT
			var currentData WasteLibrary.UltDeviceType
			currentData.New()
			currentData.DeviceId = deviceId

			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
				WasteLibrary.HTTP_DATA:   {currentData.ToString()},
			}
		} else if opType2 == "GET_RFIDDEVICES_WEB" {
			//TO DO
			//check
			path1 = "webapi"
			path2 = "getDevices"
			//currentHeader - HttpClientHeaderType *
			//currentHeader.DeviceType - DEVICE_TYPE_ULT *
			//HTTP_HEADER : currentHeader
			//
			//Retval : CustomerRfidDevicesListType
			currentHeader.DeviceType = WasteLibrary.DEVICE_TYPE_RFID

			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
			}
		} else if opType2 == "GET_RECYDEVICES_WEB" {
			//TO DO
			//check
			path1 = "webapi"
			path2 = "getDevices"
			//currentHeader - HttpClientHeaderType *
			//currentHeader.DeviceType - DEVICE_TYPE_RECY *
			//HTTP_HEADER : currentHeader
			//
			//Retval : CustomerRecyDevicesListType
			currentHeader.DeviceType = WasteLibrary.DEVICE_TYPE_RECY

			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
			}
		} else if opType2 == "GET_ULTDEVICES_WEB" {
			//TO DO
			//check
			path1 = "webapi"
			path2 = "getDevices"
			//currentHeader - HttpClientHeaderType *
			//currentHeader.DeviceType - DEVICE_TYPE_ULT *
			//HTTP_HEADER : currentHeader
			//
			//Retval : CustomerUltDevicesListType
			currentHeader.DeviceType = WasteLibrary.DEVICE_TYPE_ULT

			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
			}
		} else if opType2 == "GET_RFIDDEVICES_AFATEK" {
			//TO DO
			//check
			path1 = "afatekapi"
			path2 = "getDevices"
			//currentHeader - HttpClientHeaderType *
			//currentHeader.DeviceType - DEVICE_TYPE_RFID *
			//currentHeader.Token - token *
			//currentHeader.CustomerId - customerId *
			//HTTP_HEADER : currentHeader
			//
			//Retval : CustomerRfidDevicesListType
			currentHeader.Token = token
			currentHeader.DeviceType = WasteLibrary.DEVICE_TYPE_RFID
			currentHeader.CustomerId = customerId

			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
			}
		} else if opType2 == "GET_RECYDEVICES_AFATEK" {
			//TO DO
			//check
			path1 = "afatekapi"
			path2 = "getDevices"
			//currentHeader - HttpClientHeaderType *
			//currentHeader.DeviceType - DEVICE_TYPE_RECY *
			//currentHeader.Token - token *
			//currentHeader.CustomerId - customerId *
			//HTTP_HEADER : currentHeader
			//
			//Retval : CustomerRecyDevicesListType
			currentHeader.Token = token
			currentHeader.DeviceType = WasteLibrary.DEVICE_TYPE_RECY
			currentHeader.CustomerId = customerId

			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
			}
		} else if opType2 == "GET_ULTDEVICES_AFATEK" {
			//TO DO
			//check
			path1 = "afatekapi"
			path2 = "getDevices"
			//currentHeader - HttpClientHeaderType *
			//currentHeader.DeviceType - DEVICE_TYPE_ULT *
			//currentHeader.Token - token *
			//currentHeader.CustomerId - customerId *
			//HTTP_HEADER : currentHeader
			//
			//Retval : CustomerUltDevicesListType
			currentHeader.Token = token
			currentHeader.DeviceType = WasteLibrary.DEVICE_TYPE_ULT
			currentHeader.CustomerId = customerId

			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
			}
		} else {

		}
	} else if opType == "CUSTOMER" {
		if opType2 == "GET_CUSTOMER_WEB" {
			//TO DO
			//check
			//Retval : CustomerType
			path1 = "webapi"
			path2 = "getCustomer"

			data = url.Values{}
		} else if opType2 == "GET_CUSTOMER_ADMIN" {
			//TO DO
			//check
			path1 = "adminapi"
			path2 = "getCustomer"
			//currentHeader - HttpClientHeaderType *
			//currentHeader.Token - token *
			//HTTP_HEADER : currentHeader
			//
			//Retval : CustomerType
			currentHeader.Token = token

			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
			}
		} else if opType2 == "GET_CUSTOMER_AFATEK" {
			//TO DO
			//check
			path1 = "afatekapi"
			path2 = "getCustomer"
			//currentHeader - HttpClientHeaderType *
			//currentHeader.Token - token *
			//HTTP_HEADER : currentHeader
			//
			//currentData - CustomerType
			//currentData.CustomerId - customerId *
			//HTTP_DATA : currentData
			//
			//Retval : CustomerType
			currentHeader.Token = token

			var currentData WasteLibrary.CustomerType
			currentData.New()
			currentData.CustomerId = customerId

			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
				WasteLibrary.HTTP_DATA:   {currentData.ToString()},
			}
		} else if opType2 == "SET_CUSTOMER_AFATEK" {
			//TO DO
			//check
			path1 = "afatekapi"
			path2 = "setCustomer"
			//currentHeader - HttpClientHeaderType *
			//currentHeader.Token - token *
			//HTTP_HEADER : currentHeader
			//
			//currentData - CustomerType *
			//currentData.CustomerId - [0,customerId] *
			//currentData.[All]
			//HTTP_DATA : currentData
			//
			//Retval : CustomerType
			currentHeader.Token = token

			var currentData WasteLibrary.CustomerType
			currentData.New()
			currentData.CustomerId = customerId

			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
				WasteLibrary.HTTP_DATA:   {currentData.ToString()},
			}
		} else if opType2 == "GET_CUSTOMERS_AFATEK" {
			//TO DO
			//check
			path1 = "afatekapi"
			path2 = "getCustomers"
			//currentHeader - HttpClientHeaderType *
			//currentHeader.Token - token *
			//HTTP_HEADER : currentHeader
			//
			//Retval : CustomersListType
			currentHeader.Token = token

			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
			}
		} else {

		}
	} else if opType == "USER" {
		if opType2 == "GET_USER_ADMIN" {
			//TO DO
			//check
			path1 = "adminapi"
			path2 = "getUser"
			//currentHeader - HttpClientHeaderType *
			//currentHeader.Token - token *
			//HTTP_HEADER : currentHeader
			//
			//currentData - UserType *
			//currentData.UserId - userId *
			//HTTP_DATA : currentData
			//
			//Retval : UserType
			currentHeader.Token = token
			var currentData WasteLibrary.UserType
			currentData.New()
			currentData.UserId = userId

			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
				WasteLibrary.HTTP_DATA:   {currentData.ToString()},
			}
		} else if opType2 == "SET_USER_ADMIN" {
			//TO DO
			//check
			path1 = "adminapi"
			path2 = "setUser"
			//currentHeader - HttpClientHeaderType *
			//currentHeader.Token - token *
			//HTTP_HEADER : currentHeader
			//
			//currentData - UserType *
			//currentData.UserId - userId *
			//currentData.[All]
			//HTTP_DATA : currentData
			//
			//Retval : UserType
			currentHeader.Token = token
			var currentData WasteLibrary.UserType
			currentData.UserId = 2
			currentData.UserName = "devafatek3"
			currentData.UserRole = WasteLibrary.USER_ROLE_ADMIN
			currentData.Password = "Amca1512003"
			currentData.Email = "developer3@afatek.com.tr"
			currentData.Active = WasteLibrary.STATU_ACTIVE

			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
				WasteLibrary.HTTP_DATA:   {currentData.ToString()},
			}
		} else if opType2 == "LOGIN_AUTH" {
			//OK
			path1 = "authapi"
			path2 = "login"
			//currentData - UserType *
			//currentData.UserName - userName *
			//currentData.Password - password *
			//HTTP_DATA : currentData
			//
			//Retval : TOKEN
			var currentData WasteLibrary.UserType
			currentData.UserName = "devafatek"
			currentData.Password = "Amca151200"
			currentData.Active = WasteLibrary.STATU_ACTIVE
			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
				WasteLibrary.HTTP_DATA:   {currentData.ToString()},
			}
		} else if opType2 == "REGISTER_AUTH" {
			//OK
			path1 = "authapi"
			path2 = "register"
			//currentData - UserType *
			//currentData.UserName - userName *
			//currentData.Password - password *
			//currentData.[All]
			//HTTP_DATA : currentData
			var currentData WasteLibrary.UserType
			currentData.UserName = "devafatek"
			currentData.Password = "Amca151200"
			currentData.Email = "developer@afatek.com.tr"
			currentData.Active = WasteLibrary.STATU_ACTIVE

			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
				WasteLibrary.HTTP_DATA:   {currentData.ToString()},
			}

		} else if opType2 == "GET_USERS_ADMIN" {
			//TO DO
			//check
			path1 = "adminapi"
			path2 = "getUsers"
			//currentHeader - HttpClientHeaderType *
			//currentHeader.Token - token *
			//HTTP_HEADER : currentHeader
			//
			//Retval : CustomerUsersListType
			currentHeader.Token = token
			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
			}
		} else {

		}
	} else if opType == "CONFIG" {
		if opType2 == "GET_ADMINCONFIG_ADMIN" {
			//TO DO
			//check
			path1 = "adminapi"
			path2 = "getConfig"
			//currentHeader - HttpClientHeaderType *
			//currentHeader.Token - token *
			//currentHeader.OpType - OPTYPE_ADMINCONFIG *
			//HTTP_HEADER : currentHeader
			//
			//Retval : AdminConfigType
			currentHeader.Token = token
			currentHeader.OpType = WasteLibrary.OPTYPE_ADMINCONFIG
			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
			}
		} else if opType2 == "SET_ADMINCONFIG_ADMIN" {
			//TO DO
			//check
			path1 = "adminapi"
			path2 = "setConfig"
			//currentHeader - HttpClientHeaderType *
			//currentHeader.Token - token *
			//currentHeader.OpType - OPTYPE_ADMINCONFIG *
			//HTTP_HEADER : currentHeader
			//
			//currentData - AdminConfigType *
			//currentData.[All]
			//HTTP_DATA : currentData

			currentHeader.Token = token
			currentHeader.OpType = WasteLibrary.OPTYPE_ADMINCONFIG
			var currentData WasteLibrary.AdminConfigType
			currentData.New()

			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
				WasteLibrary.HTTP_DATA:   {currentData.ToString()},
			}
		} else if opType2 == "GET_LOCALCONFIG_ADMIN" {
			//TO DO
			//check
			path1 = "adminapi"
			path2 = "getConfig"
			//currentHeader - HttpClientHeaderType *
			//currentHeader.Token - token *
			//currentHeader.OpType - OPTYPE_LOCALCONFIG *
			//HTTP_HEADER : currentHeader
			//
			//Retval : LocalConfigType
			currentHeader.Token = token
			currentHeader.OpType = WasteLibrary.OPTYPE_LOCALCONFIG
			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
			}
		} else if opType2 == "SET_LOCALCONFIG_ADMIN" {
			//TO DO
			//check
			path1 = "adminapi"
			path2 = "setConfig"
			//currentHeader - HttpClientHeaderType *
			//currentHeader.Token - token *
			//currentHeader.OpType - OPTYPE_LOCALCONFIG *
			//HTTP_HEADER : currentHeader
			//
			//currentData - LocalConfigType *
			//currentData.[All]
			//HTTP_DATA : currentData

			currentHeader.Token = token
			currentHeader.OpType = WasteLibrary.OPTYPE_LOCALCONFIG
			var currentData WasteLibrary.LocalConfigType
			currentData.New()

			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
				WasteLibrary.HTTP_DATA:   {currentData.ToString()},
			}
		} else if opType2 == "GET_CUSTOMERCONFIG_ADMIN" {
			//TO DO
			//check
			path1 = "adminapi"
			path2 = "getConfig"
			//currentHeader - HttpClientHeaderType *
			//currentHeader.Token - token *
			//currentHeader.OpType - OPTYPE_CUSTOMERCONFIG *
			//HTTP_HEADER : currentHeader
			//
			//Retval : CustomerConfigType
			currentHeader.Token = token
			currentHeader.OpType = WasteLibrary.OPTYPE_CUSTOMERCONFIG
			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
			}
		} else if opType2 == "SET_CUSTOMERCONFIG_ADMIN" {
			//TO DO
			//check
			path1 = "adminapi"
			path2 = "setConfig"
			//currentHeader - HttpClientHeaderType *
			//currentHeader.Token - token *
			//currentHeader.OpType - OPTYPE_CUSTOMERCONFIG *
			//HTTP_HEADER : currentHeader
			//
			//currentData - CustomerConfigType *
			//currentData.[All]
			//HTTP_DATA : currentData
			currentHeader.Token = token
			currentHeader.OpType = WasteLibrary.OPTYPE_CUSTOMERCONFIG
			var currentData WasteLibrary.CustomerConfigType
			currentData.New()
			currentData.ArventoApp = WasteLibrary.STATU_ACTIVE
			currentData.ArventoUserName = "devafatekarvento"
			currentData.ArventoPin1 = "pin1"
			currentData.ArventoPin2 = "pin2"
			currentData.SystemProblem = WasteLibrary.STATU_PASSIVE
			currentData.TruckStopTrace = WasteLibrary.STATU_ACTIVE
			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
				WasteLibrary.HTTP_DATA:   {currentData.ToString()},
			}
		} else if opType2 == "GET_ADMINCONFIG_WEB" {
			//TO DO
			//check
			path1 = "webapi"
			path2 = "getConfig"
			//currentHeader - HttpClientHeaderType *
			//currentHeader.OpType - OPTYPE_ADMINCONFIG *
			//HTTP_HEADER : currentHeader
			//
			//Retval : AdminConfigType
			currentHeader.OpType = WasteLibrary.OPTYPE_ADMINCONFIG
			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
			}
		} else if opType2 == "GET_LOCALCONFIG_WEB" {
			//TO DO
			//check
			path1 = "webapi"
			path2 = "getConfig"
			//currentHeader - HttpClientHeaderType *
			//currentHeader.OpType - OPTYPE_LOCALCONFIG *
			//HTTP_HEADER : currentHeader
			//
			//Retval : LocalConfigType
			currentHeader.OpType = WasteLibrary.OPTYPE_LOCALCONFIG
			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
			}
		} else if opType2 == "GET_CUSTOMERCONFIG_WEB" {
			//TO DO
			//check
			path1 = "webapi"
			path2 = "getConfig"
			//currentHeader - HttpClientHeaderType *
			//currentHeader.OpType - OPTYPE_CUSTOMERCONFIG *
			//HTTP_HEADER : currentHeader
			//
			//Retval : CustomerConfigType
			currentHeader.OpType = WasteLibrary.OPTYPE_CUSTOMERCONFIG
			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
			}
		} else {

		}
	} else {
		var currentCustomerRfidDevices WasteLibrary.CustomerRfidDevicesType
		currentCustomerRfidDevices.New()
		currentCustomerRfidDevices.Devices["0"] = 0
		fmt.Println(currentCustomerRfidDevices)
		return
	}
	var urlFull string = "http://" + urlVal + "/" + path1 + "/" + path2
	fmt.Println(urlFull)
	resultVal := WasteLibrary.HttpPostReq(urlFull, data)
	fmt.Println(resultVal)
}
