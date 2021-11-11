package main

import (
	"fmt"
	"math"
	"net/url"
	"time"

	"github.com/devafatek/WasteLibrary"
)

func main() {

	var readerType string = "DEVICE"
	var readerType2 string = "SET_RFIDDEVICE_WEB"
	var currentHeader WasteLibrary.HttpClientHeaderType
	currentHeader.New()
	var urlVal string = "bodrum.aws.afatek.com.tr"
	var path1 string = "webapi"
	var path2 string = "getLink"
	data := url.Values{}
	var deviceId float64 = 14
	var customerId float64 = 1
	var userId float64 = 2
	var token = "MiMkMmEkMTAkdzJsRVgyUXBsdDU4MUNTTGRYWW5mZVZDVUd6TGpmbFMuUzZUQmhRRkx1bTdXRlQ3Yk9Ed1M="

	if readerType == "DEVICE" {
		if readerType2 == "GET_RFIDDEVICE_WEB" {
			//OK
			path1 = "webapi"
			path2 = "getDevice"
			//currentHeader - HttpClientHeaderType *
			//currentHeader.DeviceType - DEVICETYPE_RFID *
			//HTTP_HEADER : currentHeader
			//
			//currentData - RfidDeviceType *
			//currentData.DeviceId - deviceId *
			//HTTP_DATA : currentData
			//
			//Retval : RfidDeviceType
			currentHeader.DeviceType = WasteLibrary.DEVICETYPE_RFID
			var currentData WasteLibrary.RfidDeviceType
			currentData.New()
			currentData.DeviceId = deviceId

			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
				WasteLibrary.HTTP_DATA:   {currentData.ToString()},
			}
			sendData(readerType, urlVal, path1, path2, data)
		} else if readerType2 == "GET_RECYDEVICE_WEB" {
			//OK
			path1 = "webapi"
			path2 = "getDevice"
			//currentHeader - HttpClientHeaderType *
			//currentHeader.DeviceType - DEVICETYPE_RECY *
			//HTTP_HEADER : currentHeader
			//
			//currentData - RecyDeviceType *
			//currentData.DeviceId - deviceId *
			//HTTP_DATA : currentData
			//
			//Retval : RecyDeviceType
			currentHeader.DeviceType = WasteLibrary.DEVICETYPE_RECY
			var currentData WasteLibrary.RecyDeviceType
			currentData.New()
			currentData.DeviceId = deviceId

			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
				WasteLibrary.HTTP_DATA:   {currentData.ToString()},
			}
			sendData(readerType, urlVal, path1, path2, data)
		} else if readerType2 == "GET_ULTDEVICE_WEB" {
			//OK
			path1 = "webapi"
			path2 = "getDevice"
			//currentHeader - HttpClientHeaderType *
			//currentHeader.DeviceType - DEVICETYPE_ULT *
			//HTTP_HEADER : currentHeader
			//
			//currentData - UltDeviceType *
			//currentData.DeviceId - deviceId *
			//HTTP_DATA : currentData
			//
			//Retval : UltDeviceType
			currentHeader.DeviceType = WasteLibrary.DEVICETYPE_ULT
			var currentData WasteLibrary.UltDeviceType
			currentData.New()
			currentData.DeviceId = deviceId

			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
				WasteLibrary.HTTP_DATA:   {currentData.ToString()},
			}
			sendData(readerType, urlVal, path1, path2, data)
		} else if readerType2 == "GET_RFIDDEVICE_AFATEK" {
			//OK
			path1 = "afatekapi"
			path2 = "getDevice"
			//currentHeader - HttpClientHeaderType *
			//currentHeader.DeviceType - DEVICETYPE_RFID *
			//currentHeader.Token - token *
			//HTTP_HEADER : currentHeader
			//
			//currentData - RfidDeviceType *
			//currentData.DeviceId - deviceId *
			//HTTP_DATA : currentData
			//
			//Retval : RfidDeviceType
			currentHeader.Token = token
			currentHeader.DeviceType = WasteLibrary.DEVICETYPE_RFID
			var currentData WasteLibrary.RfidDeviceType
			currentData.New()
			currentData.DeviceId = deviceId

			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
				WasteLibrary.HTTP_DATA:   {currentData.ToString()},
			}
			sendData(readerType, urlVal, path1, path2, data)
		} else if readerType2 == "GET_RECYDEVICE_AFATEK" {
			//OK
			path1 = "afatekapi"
			path2 = "getDevice"
			//currentHeader - HttpClientHeaderType *
			//currentHeader.DeviceType - DEVICETYPE_RECY *
			//currentHeader.Token - token *
			//HTTP_HEADER : currentHeader
			//
			//currentData - RecyDeviceType *
			//currentData.DeviceId - deviceId *
			//HTTP_DATA : currentData
			//
			//Retval : RecyDeviceType
			currentHeader.Token = token
			currentHeader.DeviceType = WasteLibrary.DEVICETYPE_RECY
			var currentData WasteLibrary.RecyDeviceType
			currentData.New()
			currentData.DeviceId = deviceId

			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
				WasteLibrary.HTTP_DATA:   {currentData.ToString()},
			}
			sendData(readerType, urlVal, path1, path2, data)
		} else if readerType2 == "GET_ULTDEVICE_AFATEK" {
			//OK
			path1 = "afatekapi"
			path2 = "getDevice"
			//currentHeader - HttpClientHeaderType *
			//currentHeader.DeviceType - DEVICETYPE_ULT *
			//currentHeader.Token - token *
			//HTTP_HEADER : currentHeader
			//
			//currentData - UltDeviceType *
			//currentData.DeviceId - deviceId *
			//HTTP_DATA : currentData
			//
			//Retval : UltDeviceType
			currentHeader.Token = token
			currentHeader.DeviceType = WasteLibrary.DEVICETYPE_ULT
			var currentData WasteLibrary.UltDeviceType
			currentData.New()
			currentData.DeviceId = deviceId

			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
				WasteLibrary.HTTP_DATA:   {currentData.ToString()},
			}
			sendData(readerType, urlVal, path1, path2, data)
		} else if readerType2 == "SET_RFIDDEVICE_AFATEK" {
			//OK
			path1 = "afatekapi"
			path2 = "setDevice"
			//currentHeader - HttpClientHeaderType *
			//currentHeader.DeviceType - DEVICE_TYPE_RFID *
			//currentHeader.Token - token *
			//HTTP_HEADER : currentHeader
			//
			//currentData - RfidDeviceType *
			//currentData.DeviceId - deviceId *
			//currentData.[All]
			//HTTP_DATA : currentData
			//
			//Retval : RfidDeviceType

			currentHeader.Token = token
			currentHeader.DeviceType = WasteLibrary.DEVICETYPE_RFID
			var currentData WasteLibrary.RfidDeviceType
			currentData.New()
			currentData.DeviceId = 48
			currentData.DeviceMain.DeviceId = currentData.DeviceId
			currentData.DeviceMain.CustomerId = 3

			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
				WasteLibrary.HTTP_DATA:   {currentData.ToString()},
			}
			sendData(readerType, urlVal, path1, path2, data)
			time.Sleep(5 * time.Second)

		} else if readerType2 == "SET_RECYDEVICE_AFATEK" {
			//TO DO
			//check
			path1 = "afatekapi"
			path2 = "setDevice"
			//currentHeader - HttpClientHeaderType *
			//currentHeader.DeviceType - DEVICETYPE_RECY *
			//currentHeader.Token - token *
			//HTTP_HEADER : currentHeader
			//
			//currentData - RecyDeviceType *
			//currentData.DeviceId - deviceId *
			//currentData.[All]
			//HTTP_DATA : currentData
			//
			//Retval : RecyDeviceType
			currentHeader.Token = token
			currentHeader.DeviceType = WasteLibrary.DEVICETYPE_RECY
			var currentData WasteLibrary.RecyDeviceType
			currentData.New()
			currentData.DeviceId = deviceId

			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
				WasteLibrary.HTTP_DATA:   {currentData.ToString()},
			}
			sendData(readerType, urlVal, path1, path2, data)
		} else if readerType2 == "SET_ULTDEVICE_AFATEK" {
			//TO DO
			//check
			path1 = "afatekapi"
			path2 = "setDevice"
			//currentHeader - HttpClientHeaderType *
			//currentHeader.DeviceType - DEVICETYPE_ULT *
			//currentHeader.Token - token *
			//HTTP_HEADER : currentHeader
			//
			//currentData - UltDeviceType *
			//currentData.DeviceId - deviceId *
			//currentData.[All]
			//HTTP_DATA : currentData
			//
			//Retval : UltDeviceType
			currentHeader.Token = token
			currentHeader.DeviceType = WasteLibrary.DEVICETYPE_ULT
			var currentData WasteLibrary.UltDeviceType
			currentData.New()
			currentData.DeviceId = deviceId

			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
				WasteLibrary.HTTP_DATA:   {currentData.ToString()},
			}
			sendData(readerType, urlVal, path1, path2, data)
		} else if readerType2 == "SET_RFIDDEVICE_WEB" {
			//TO DO
			//check
			path1 = "webapi"
			path2 = "setDevice"
			//currentHeader - HttpClientHeaderType *
			//currentHeader.DeviceType - DEVICETYPE_RFID *
			//currentHeader.Token - token *
			//HTTP_HEADER : currentHeader
			//
			//currentData - RfidDeviceType *
			//currentData.DeviceId - deviceId *
			//currentData.[All]
			//HTTP_DATA : currentData
			//
			//Retval : RfidDeviceType
			currentHeader.Token = token
			currentHeader.DeviceType = WasteLibrary.DEVICETYPE_RFID
			var currentData WasteLibrary.RfidDeviceType
			currentData.New()
			currentData.DeviceId = 48
			currentData.DeviceDetail.PlateNo = "07 AAV 297"

			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
				WasteLibrary.HTTP_DATA:   {currentData.ToString()},
			}
			sendData(readerType, urlVal, path1, path2, data)
			time.Sleep(5 * time.Second)
		} else if readerType2 == "SET_RECYDEVICE_WEB" {
			//TO DO
			//check
			path1 = "webapi"
			path2 = "setDevice"
			//currentHeader - HttpClientHeaderType *
			//currentHeader.DeviceType - DEVICETYPE_RECY *
			//currentHeader.Token - token *
			//HTTP_HEADER : currentHeader
			//
			//currentData - RecyDeviceType *
			//currentData.DeviceId - deviceId *
			//currentData.[All]
			//HTTP_DATA : currentData
			//
			//Retval : RecyDeviceType
			currentHeader.Token = token
			currentHeader.DeviceType = WasteLibrary.DEVICETYPE_RECY
			var currentData WasteLibrary.RecyDeviceType
			currentData.New()
			currentData.DeviceId = deviceId

			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
				WasteLibrary.HTTP_DATA:   {currentData.ToString()},
			}
			sendData(readerType, urlVal, path1, path2, data)
		} else if readerType2 == "SET_ULTDEVICE_WEB" {
			//TO DO
			//check
			path1 = "webapi"
			path2 = "setDevice"
			//currentHeader - HttpClientHeaderType *
			//currentHeader.DeviceType - DEVICETYPE_ULT *
			//currentHeader.Token - token *
			//HTTP_HEADER : currentHeader
			//
			//currentData - UltDeviceType *
			//currentData.DeviceId - deviceId *
			//currentData.[All]
			//HTTP_DATA : currentData
			//
			//Retval : UltDeviceType
			currentHeader.Token = token
			currentHeader.DeviceType = WasteLibrary.DEVICETYPE_ULT
			var currentData WasteLibrary.UltDeviceType
			currentData.New()
			currentData.DeviceId = deviceId

			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
				WasteLibrary.HTTP_DATA:   {currentData.ToString()},
			}
			sendData(readerType, urlVal, path1, path2, data)
		} else if readerType2 == "GET_RFIDDEVICES_WEB" {
			//TO DO
			//check
			urlVal = "bodrum.aws.afatek.com.tr"
			path1 = "webapi"
			path2 = "getDevices"
			//currentHeader - HttpClientHeaderType *
			//currentHeader.DeviceType - DEVICETYPE_RFID *
			//HTTP_HEADER : currentHeader
			//
			//Retval : CustomerRfidDevicesListType
			currentHeader.DeviceType = WasteLibrary.DEVICETYPE_RFID

			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
			}
			sendData(readerType, urlVal, path1, path2, data)
		} else if readerType2 == "GET_RECYDEVICES_WEB" {
			//TO DO
			//check
			path1 = "webapi"
			path2 = "getDevices"
			//currentHeader - HttpClientHeaderType *
			//currentHeader.DeviceType - DEVICETYPE_RECY *
			//HTTP_HEADER : currentHeader
			//
			//Retval : CustomerRecyDevicesListType
			currentHeader.DeviceType = WasteLibrary.DEVICETYPE_RECY

			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
			}
			sendData(readerType, urlVal, path1, path2, data)
		} else if readerType2 == "GET_ULTDEVICES_WEB" {
			//TO DO
			//check
			path1 = "webapi"
			path2 = "getDevices"
			//currentHeader - HttpClientHeaderType *
			//currentHeader.DeviceType - DEVICETYPE_ULT *
			//HTTP_HEADER : currentHeader
			//
			//Retval : CustomerUltDevicesListType
			currentHeader.DeviceType = WasteLibrary.DEVICETYPE_ULT

			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
			}
			sendData(readerType, urlVal, path1, path2, data)
		} else if readerType2 == "GET_RFIDDEVICES_AFATEK" {
			//TO DO
			//check
			path1 = "afatekapi"
			path2 = "getDevices"
			//currentHeader - HttpClientHeaderType *
			//currentHeader.DeviceType - DEVICETYPE_RFID *
			//currentHeader.Token - token *
			//currentHeader.CustomerId - customerId *
			//HTTP_HEADER : currentHeader
			//
			//Retval : CustomerRfidDevicesListType
			currentHeader.Token = token
			currentHeader.DeviceType = WasteLibrary.DEVICETYPE_RFID
			currentHeader.CustomerId = customerId

			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
			}
			sendData(readerType, urlVal, path1, path2, data)
		} else if readerType2 == "GET_RECYDEVICES_AFATEK" {
			//TO DO
			//check
			path1 = "afatekapi"
			path2 = "getDevices"
			//currentHeader - HttpClientHeaderType *
			//currentHeader.DeviceType - DEVICETYPE_RECY *
			//currentHeader.Token - token *
			//currentHeader.CustomerId - customerId *
			//HTTP_HEADER : currentHeader
			//
			//Retval : CustomerRecyDevicesListType
			currentHeader.Token = token
			currentHeader.DeviceType = WasteLibrary.DEVICETYPE_RECY
			currentHeader.CustomerId = customerId

			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
			}
			sendData(readerType, urlVal, path1, path2, data)
		} else if readerType2 == "GET_ULTDEVICES_AFATEK" {
			//TO DO
			//check
			path1 = "afatekapi"
			path2 = "getDevices"
			//currentHeader - HttpClientHeaderType *
			//currentHeader.DeviceType - DEVICETYPE_ULT *
			//currentHeader.Token - token *
			//currentHeader.CustomerId - customerId *
			//HTTP_HEADER : currentHeader
			//
			//Retval : CustomerUltDevicesListType
			currentHeader.Token = token
			currentHeader.DeviceType = WasteLibrary.DEVICETYPE_ULT
			currentHeader.CustomerId = customerId

			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
			}
			sendData(readerType, urlVal, path1, path2, data)
		} else {

		}
	} else if readerType == "CUSTOMER" {
		if readerType2 == "GET_CUSTOMER_WEB" {
			//OK
			//Retval : CustomerType
			path1 = "webapi"
			path2 = "getCustomer"

			data = url.Values{}
			sendData(readerType, urlVal, path1, path2, data)
		} else if readerType2 == "GET_CUSTOMER_AFATEK" {
			//OK
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
			sendData(readerType, urlVal, path1, path2, data)
		} else if readerType2 == "SET_CUSTOMER_AFATEK" {
			//OK
			path1 = "afatekapi"
			path2 = "setCustomer"
			//currentHeader - HttpClientHeaderType *
			//currentHeader.Token - token *
			//HTTP_HEADER : currentHeader
			//
			//currentData - CustomerType *
			//currentData.CustomerId - [0,customerId] *
			//currentData.CustomerName - customerName *
			//currentData.CustomerLink - customerLink *
			//currentData.[All]
			//HTTP_DATA : currentData
			//
			//Retval : CustomerType
			currentHeader.Token = token

			var currentData WasteLibrary.CustomerType
			currentData.New()
			currentData.CustomerId = 0 //3 // customerId
			currentData.CustomerName = "BODRUM"
			currentData.CustomerLink = "bodrum.aws.afatek.com.tr"
			currentData.RfIdApp = WasteLibrary.STATU_ACTIVE
			currentData.UltApp = WasteLibrary.STATU_ACTIVE
			currentData.RecyApp = WasteLibrary.STATU_ACTIVE
			currentData.Active = WasteLibrary.STATU_ACTIVE
			currentData.CreateTime = WasteLibrary.GetTime()

			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
				WasteLibrary.HTTP_DATA:   {currentData.ToString()},
			}
			sendData(readerType, urlVal, path1, path2, data)
		} else if readerType2 == "GET_CUSTOMERS_AFATEK" {
			//OK
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
			sendData(readerType, urlVal, path1, path2, data)
		} else {

		}
	} else if readerType == "USER" {
		if readerType2 == "GET_USER_WEB" {
			//OK
			path1 = "webapi"
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
			sendData(readerType, urlVal, path1, path2, data)
		} else if readerType2 == "SET_USER_WEB" {
			//OK
			path1 = "webapi"
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
			currentData.UserId = 3
			currentData.UserName = "devafatek10"
			currentData.UserRole = WasteLibrary.USER_ROLE_ADMIN
			currentData.Password = "Amca1512003"
			currentData.Email = "developer3@afatek.com.tr"
			currentData.Active = WasteLibrary.STATU_ACTIVE

			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
				WasteLibrary.HTTP_DATA:   {currentData.ToString()},
			}
			sendData(readerType, urlVal, path1, path2, data)
		} else if readerType2 == "LOGIN_AUTH" {
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
			sendData(readerType, urlVal, path1, path2, data)
		} else if readerType2 == "REGISTER_AUTH" {
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

			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
				WasteLibrary.HTTP_DATA:   {currentData.ToString()},
			}
			sendData(readerType, urlVal, path1, path2, data)
		} else if readerType2 == "GET_USERS_WEB" {
			//OK
			path1 = "webapi"
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
			sendData(readerType, urlVal, path1, path2, data)
		} else {

		}
	} else if readerType == "CONFIG" {
		if readerType2 == "SET_ADMINCONFIG_WEB" {
			//OK
			path1 = "webapi"
			path2 = "setConfig"
			//currentHeader - HttpClientHeaderType *
			//currentHeader.Token - token *
			//currentHeader.DataType - DATATYPE_ADMINCONFIG *
			//HTTP_HEADER : currentHeader
			//
			//currentData - AdminConfigType *
			//currentData.[All]
			//HTTP_DATA : currentData

			currentHeader.Token = token
			currentHeader.DataType = WasteLibrary.DATATYPE_ADMINCONFIG
			var currentData WasteLibrary.AdminConfigType
			currentData.New()

			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
				WasteLibrary.HTTP_DATA:   {currentData.ToString()},
			}
			sendData(readerType, urlVal, path1, path2, data)
		} else if readerType2 == "SET_LOCALCONFIG_WEB" {
			//OK
			path1 = "webapi"
			path2 = "setConfig"
			//currentHeader - HttpClientHeaderType *
			//currentHeader.Token - token *
			//currentHeader.DataType - DATATYPE_LOCALCONFIG *
			//HTTP_HEADER : currentHeader
			//
			//currentData - LocalConfigType *
			//currentData.[All]
			//HTTP_DATA : currentData

			currentHeader.Token = token
			currentHeader.DataType = WasteLibrary.DATATYPE_LOCALCONFIG
			var currentData WasteLibrary.LocalConfigType
			currentData.New()

			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
				WasteLibrary.HTTP_DATA:   {currentData.ToString()},
			}
			sendData(readerType, urlVal, path1, path2, data)
		} else if readerType2 == "SET_CUSTOMERCONFIG_AFATEK" {
			//OK
			path1 = "afatekapi"
			path2 = "setConfig"
			//currentHeader - HttpClientHeaderType *
			//currentHeader.Token - token *
			//currentHeader.DataType - OPTYPE_CUSTOMERCONFIG *
			//HTTP_HEADER : currentHeader
			//
			//currentData - CustomerConfigType *
			//currentData.[All]
			//HTTP_DATA : currentData
			currentHeader.Token = token
			currentHeader.DataType = WasteLibrary.DATATYPE_CUSTOMERCONFIG
			var currentData WasteLibrary.CustomerConfigType
			currentData.New()
			currentData.CustomerId = 3
			currentData.ArventoApp = WasteLibrary.STATU_ACTIVE
			currentData.ArventoUserName = "afatekbilisim"
			currentData.ArventoPin1 = "Amca151200!Furkan"
			currentData.ArventoPin2 = "Amca151200!Furkan"
			currentData.SystemProblem = WasteLibrary.STATU_PASSIVE
			currentData.TruckStopTrace = WasteLibrary.STATU_ACTIVE
			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
				WasteLibrary.HTTP_DATA:   {currentData.ToString()},
			}
			sendData(readerType, urlVal, path1, path2, data)
		} else if readerType2 == "GET_ADMINCONFIG_WEB" {
			//OK
			path1 = "webapi"
			path2 = "getConfig"
			//currentHeader - HttpClientHeaderType *
			//currentHeader.DataType - DATATYPE_ADMINCONFIG *
			//HTTP_HEADER : currentHeader
			//
			//Retval : AdminConfigType
			currentHeader.DataType = WasteLibrary.DATATYPE_ADMINCONFIG
			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
			}
			sendData(readerType, urlVal, path1, path2, data)
		} else if readerType2 == "GET_LOCALCONFIG_WEB" {
			//OK
			path1 = "webapi"
			path2 = "getConfig"
			//currentHeader - HttpClientHeaderType *
			//currentHeader.DataType - DATATYPE_LOCALCONFIG *
			//HTTP_HEADER : currentHeader
			//
			//Retval : LocalConfigType
			currentHeader.DataType = WasteLibrary.DATATYPE_LOCALCONFIG
			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
			}
			sendData(readerType, urlVal, path1, path2, data)
		} else {

		}
	} else if readerType == "LISTENER" {
		urlVal = "listener.aws.afatek.com.tr"
		if readerType2 == "ULT" {
			//TO DO
			//check

		} else if readerType2 == "CAM_RFID" {
			//TO DO
			//check

		} else if readerType2 == "STATUS_RFID" {
			//OK
			urlVal = urlVal + "/data"
			currentHeader.New()
			currentHeader.AppType = WasteLibrary.APPTYPE_RFID
			currentHeader.DeviceNo = "12345678901234567"
			currentHeader.ReaderType = WasteLibrary.READERTYPE_STATUS
			currentHeader.Time = WasteLibrary.GetTime()
			currentHeader.Repeat = WasteLibrary.STATU_PASSIVE
			currentHeader.DeviceType = WasteLibrary.DEVICETYPE_RFID
			var currentData WasteLibrary.RfidDeviceType
			currentData.New()
			currentData.DeviceStatu.CamStatus = WasteLibrary.STATU_ACTIVE
			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
				WasteLibrary.HTTP_DATA:   {currentData.ToString()},
			}
			sendData(readerType, urlVal, path1, path2, data)
		} else if readerType2 == "READER_RFID" {
			//OK
			urlVal = urlVal + "/data"
			currentHeader.New()
			currentHeader.AppType = WasteLibrary.APPTYPE_RFID
			currentHeader.DeviceNo = "12345678901234567"
			currentHeader.ReaderType = WasteLibrary.READERTYPE_RF
			currentHeader.Time = WasteLibrary.GetTime()
			currentHeader.Repeat = WasteLibrary.STATU_PASSIVE
			currentHeader.DeviceType = WasteLibrary.DEVICETYPE_RFID
			var currentData WasteLibrary.TagType
			currentData.New()
			currentData.TagMain.Epc = "1234567"
			currentData.TagReader.UID = "89078"
			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
				WasteLibrary.HTTP_DATA:   {currentData.ToString()},
			}
			sendData(readerType, urlVal, path1, path2, data)
		} else if readerType2 == "THERM_RFID" {
			//OK
			urlVal = urlVal + "/data"
			currentHeader.New()
			currentHeader.AppType = WasteLibrary.APPTYPE_RFID
			currentHeader.DeviceNo = "12345678901234567"
			currentHeader.ReaderType = WasteLibrary.READERTYPE_THERM
			currentHeader.Time = WasteLibrary.GetTime()
			currentHeader.Repeat = WasteLibrary.STATU_PASSIVE
			currentHeader.DeviceType = WasteLibrary.DEVICETYPE_RFID
			var currentData WasteLibrary.RfidDeviceType
			currentData.New()
			currentData.DeviceTherm.Therm = "35"
			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
				WasteLibrary.HTTP_DATA:   {currentData.ToString()},
			}
			sendData(readerType, urlVal, path1, path2, data)
		} else if readerType2 == "GPS_RFID" {
			//OK
			urlVal = urlVal + "/data"
			currentHeader.New()
			currentHeader.AppType = WasteLibrary.APPTYPE_RFID
			currentHeader.DeviceNo = "12345678901234567"
			currentHeader.ReaderType = WasteLibrary.READERTYPE_GPS
			currentHeader.Time = WasteLibrary.GetTime()
			currentHeader.Repeat = WasteLibrary.STATU_PASSIVE
			currentHeader.DeviceType = WasteLibrary.DEVICETYPE_RFID
			var currentData WasteLibrary.RfidDeviceType
			currentData.New()
			currentData.DeviceGps.Latitude = 1
			currentData.DeviceGps.Longitude = 2
			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
				WasteLibrary.HTTP_DATA:   {currentData.ToString()},
			}
			sendData(readerType, urlVal, path1, path2, data)
		} else {

		}
	} else {

		distance := distanceInKmBetweenEarthCoordinates(39.983073, 32.814551, 39.981752, 32.769141)
		fmt.Println(distance)
	}

}

func sendData(readerType string, urlVal string, path1 string, path2 string, data url.Values) {
	var urlFull string = "http://" + urlVal + "/" + path1 + "/" + path2
	if readerType == "LISTENER" {
		urlFull = "http://" + urlVal
	}
	fmt.Println(urlFull)
	resultVal := WasteLibrary.HttpPostReq(urlFull, data)
	fmt.Println(resultVal)
}

func distanceInKmBetweenEarthCoordinates(lat1 float64, lon1 float64, lat2 float64, lon2 float64) float64 {
	var earthRadiusKm float64 = 6371

	var dLat = degreesToRadians(lat2 - lat1)
	var dLon = degreesToRadians(lon2 - lon1)

	lat1 = degreesToRadians(lat1)
	lat2 = degreesToRadians(lat2)

	var a = math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Sin(dLon/2)*math.Sin(dLon/2)*math.Cos(lat1)*math.Cos(lat2)
	var c = 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return earthRadiusKm * c * 1000
}

func degreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}
