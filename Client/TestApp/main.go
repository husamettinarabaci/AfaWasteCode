package main

import (
	"fmt"
	"net/url"

	"github.com/devafatek/WasteLibrary"
)

func main() {

	var opType string = "CONFIG"
	var opType2 string = "SET_CUSTOMERCONFIG_ADMIN"
	var currentHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.HttpClientHeaderType{}
	var urlVal string = "afatek.aws.afatek.com.tr"
	var path1 string = "webapi"
	var path2 string = "getLink"
	data := url.Values{}
	var deviceId float64 = 0
	var customerId float64 = 1
	var userId float64 = 2
	var token = "MSMkMmEkMTAkOGVMTmlVM1Nqc0hockI1alJCcG5ZdWdQdFU1a2FpMDluazJtUXZxUEtXSVp3SUxVc2FyZks="
	currentHeader.DeviceId = deviceId
	currentHeader.CustomerId = customerId
	currentHeader.Token = token

	if opType == "DEVICE" {
		if opType2 == "GET_DEVICE_WEB" {
			//TO DO
			//check
			path1 = "webapi"
			path2 = "getDevice"
			//currentData - DeviceType
			//currentData.DevıceId - deviceId
			//HTTP_DATA : currentData
			//
			//Retval :
			var currentData WasteLibrary.DeviceType = WasteLibrary.DeviceType{
				DeviceId: deviceId,
			}

			data = url.Values{
				WasteLibrary.HTTP_DATA: {currentData.ToString()},
			}
		} else if opType2 == "GET_DEVICE_AFATEK" {
			//TO DO
			//check
			path1 = "afatekapi"
			path2 = "getDevice"
			//currentHeader - HttpClientHeaderType
			//currentHeader.DevıceId - deviceId
			//currentHeader.CustomerId - customerId
			//currentHeader.Token - token
			//HTTP_HEADER : currentHeader
			//
			//currentData - DeviceType
			//currentData.DevıceId - deviceId
			//HTTP_DATA : currentData
			//
			//Retval :
			var currentData WasteLibrary.DeviceType = WasteLibrary.DeviceType{
				DeviceId: deviceId,
			}

			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
				WasteLibrary.HTTP_DATA:   {currentData.ToString()},
			}

		} else if opType2 == "SET_DEVICE" {
			//TO DO
			//check
			path1 = "afatekapi"
			path2 = "setDevice"
			//currentHeader - HttpClientHeaderType
			//currentHeader.DevıceId - deviceId
			//currentHeader.CustomerId - customerId
			//currentHeader.Token - token
			//HTTP_HEADER : currentHeader
			//
			//currentData - DeviceType
			//currentData.DevıceId - [0,deviceId]
			//currentData.[All]
			//HTTP_DATA : currentData

			var currentData WasteLibrary.DeviceType = WasteLibrary.DeviceType{
				DeviceId: deviceId,
			}

			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
				WasteLibrary.HTTP_DATA:   {currentData.ToString()},
			}

		} else if opType2 == "GET_DEVICES_WEB" {
			//TO DO
			//check
			path1 = "webapi"
			path2 = "getDevices"

			data = url.Values{}
		} else if opType2 == "GET_DEVICES_AFATEK" {
			//TO DO
			//check
			path1 = "afatekapi"
			path2 = "getDevices"
			//currentHeader - HttpClientHeaderType
			//currentHeader.CustomerId - customerId
			//currentHeader.Token - token
			//HTTP_HEADER : currentHeader
			//
			//Retval :
			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
			}
		} else {

		}
	} else if opType == "CUSTOMER" {
		if opType2 == "GET_CUSTOMER_WEB" {
			//OK
			//Retval : CustomerType
			path1 = "webapi"
			path2 = "getCustomer"

			data = url.Values{}
		} else if opType2 == "GET_CUSTOMER_ADMIN" {
			//TO DO
			//check
			path1 = "adminapi"
			path2 = "getCustomer"
			//currentHeader - HttpClientHeaderType
			//currentHeader.Token - token
			//HTTP_HEADER : currentHeader
			//
			//Retval :
			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
			}
		} else if opType2 == "GET_CUSTOMER_AFATEK" {
			//TO DO
			//check
			path1 = "afatekapi"
			path2 = "getCustomer"
			//currentHeader - HttpClientHeaderType
			//currentHeader.Token - token
			//HTTP_HEADER : currentHeader
			//
			//currentData - CustomerType
			//currentData.CustomerId - customerId
			//HTTP_DATA : currentData
			//
			//Retval :
			var currentData WasteLibrary.CustomerType = WasteLibrary.CustomerType{
				CustomerId: customerId,
			}

			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
				WasteLibrary.HTTP_DATA:   {currentData.ToString()},
			}
		} else if opType2 == "SET_CUSTOMER" {
			//TO DO
			//check
			path1 = "afatekapi"
			path2 = "setCustomer"
			//currentHeader - HttpClientHeaderType
			//currentHeader.Token - token
			//HTTP_HEADER : currentHeader
			//
			//currentData - CustomerType
			//currentData.CustomerId - [0,customerId]
			//currentData.[All]
			//HTTP_DATA : currentData

			var currentData WasteLibrary.CustomerType = WasteLibrary.CustomerType{
				CustomerId: customerId,
			}

			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
				WasteLibrary.HTTP_DATA:   {currentData.ToString()},
			}
		} else if opType2 == "GET_CUSTOMERS" {
			//OK
			path1 = "afatekapi"
			path2 = "getCustomers"
			//currentHeader - HttpClientHeaderType
			//currentHeader.Token - token *
			//HTTP_HEADER : currentHeader
			//
			//Retval : CustomersListType
			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
			}
		} else {

		}
	} else if opType == "USER" {
		if opType2 == "GET_USER" {
			//OK
			path1 = "adminapi"
			path2 = "getUser"
			//currentHeader - HttpClientHeaderType
			//currentHeader.Token - token *
			//HTTP_HEADER : currentHeader
			//
			//currentData - UserType
			//currentData.UserId - userId *
			//HTTP_DATA : currentData
			//
			//Retval : UserType
			var currentData WasteLibrary.UserType = WasteLibrary.UserType{
				UserId: userId,
			}

			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
				WasteLibrary.HTTP_DATA:   {currentData.ToString()},
			}
		} else if opType2 == "SET_USER" {
			//OK
			path1 = "adminapi"
			path2 = "setUser"
			//currentHeader - HttpClientHeaderType
			//currentHeader.Token - token *
			//HTTP_HEADER : currentHeader
			//
			//currentData - UserType
			//currentData.UserId - userId *
			//currentData.[All]
			//HTTP_DATA : currentData

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
		} else if opType2 == "LOGIN" {
			//OK
			path1 = "authapi"
			path2 = "login"
			//currentData - UserType
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
		} else if opType2 == "REGISTER" {
			//OK
			path1 = "authapi"
			path2 = "register"
			//currentData - UserType
			//currentData.UserName - userName *
			//currentData.Password - password *
			//currentData.[All]
			//HTTP_DATA : currentData
			var currentData WasteLibrary.UserType
			currentData.UserName = "devafatek2"
			currentData.Password = "Amca1512002"
			currentData.Email = "developer@afatek.com.tr"
			currentData.Active = WasteLibrary.STATU_ACTIVE

			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
				WasteLibrary.HTTP_DATA:   {currentData.ToString()},
			}

		} else if opType2 == "GET_USERS" {
			//OK
			path1 = "adminapi"
			path2 = "getUsers"
			//currentHeader - HttpClientHeaderType
			//currentHeader.Token - token *
			//HTTP_HEADER : currentHeader
			//
			//Retval : CustomerUsersListType

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
			//currentHeader - HttpClientHeaderType
			//currentHeader.Token - token
			//currentHeader.OpType - OPTYPE_ADMINCONFIG
			//HTTP_HEADER : currentHeader
			//
			//Retval :
			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
			}
		} else if opType2 == "SET_ADMINCONFIG_ADMIN" {
			//TO DO
			//check
			path1 = "adminapi"
			path2 = "setConfig"
			//currentHeader - HttpClientHeaderType
			//currentHeader.Token - token
			//currentHeader.OpType - OPTYPE_ADMINCONFIG
			//HTTP_HEADER : currentHeader
			//
			//currentData - AdminConfigType
			//currentData.[All]
			//HTTP_DATA : currentData

			var currentData WasteLibrary.AdminConfigType = WasteLibrary.AdminConfigType{}

			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
				WasteLibrary.HTTP_DATA:   {currentData.ToString()},
			}
		} else if opType2 == "GET_LOCALCONFIG_ADMIN" {
			//TO DO
			//check
			path1 = "adminapi"
			path2 = "getConfig"
			//currentHeader - HttpClientHeaderType
			//currentHeader.Token - token
			//currentHeader.OpType - OPTYPE_LOCALCONFIG
			//HTTP_HEADER : currentHeader
			//
			//Retval :
			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
			}
		} else if opType2 == "SET_LOCALCONFIG_ADMIN" {
			//TO DO
			//check
			path1 = "adminapi"
			path2 = "setConfig"
			//currentHeader - HttpClientHeaderType
			//currentHeader.Token - token
			//currentHeader.OpType - OPTYPE_LOCALCONFIG
			//HTTP_HEADER : currentHeader
			//
			//currentData - LocalConfigType
			//currentData.[All]
			//HTTP_DATA : currentData

			var currentData WasteLibrary.LocalConfigType = WasteLibrary.LocalConfigType{}

			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
				WasteLibrary.HTTP_DATA:   {currentData.ToString()},
			}
		} else if opType2 == "GET_CUSTOMERCONFIG_ADMIN" {
			//TO DO
			//check
			path1 = "adminapi"
			path2 = "getConfig"
			//currentHeader - HttpClientHeaderType
			//currentHeader.Token - token
			//currentHeader.OpType - OPTYPE_CUSTOMERCONFIG
			//HTTP_HEADER : currentHeader
			//
			//Retval :
			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
			}
		} else if opType2 == "SET_CUSTOMERCONFIG_ADMIN" {
			//TO DO
			//check
			path1 = "adminapi"
			path2 = "setConfig"
			//currentHeader - HttpClientHeaderType
			//currentHeader.Token - token *
			//currentHeader.OpType - OPTYPE_CUSTOMERCONFIG *
			//HTTP_HEADER : currentHeader
			//
			//currentData - CustomerConfigType
			//currentData.[All]
			//HTTP_DATA : currentData

			currentHeader.OpType = WasteLibrary.OPTYPE_CUSTOMERCONFIG
			var currentData WasteLibrary.CustomerConfigType = WasteLibrary.CustomerConfigType{}
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
			//currentHeader - HttpClientHeaderType
			//currentHeader.OpType - OPTYPE_ADMINCONFIG
			//HTTP_HEADER : currentHeader
			//
			//Retval :
			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
			}
		} else if opType2 == "GET_LOCALCONFIG_WEB" {
			//TO DO
			//check
			path1 = "webapi"
			path2 = "getConfig"
			//currentHeader - HttpClientHeaderType
			//currentHeader.OpType - OPTYPE_LOCALCONFIG
			//HTTP_HEADER : currentHeader
			//
			//Retval :
			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
			}
		} else if opType2 == "GET_CUSTOMERCONFIG_WEB" {
			//TO DO
			//check
			path1 = "webapi"
			path2 = "getConfig"
			//currentHeader - HttpClientHeaderType
			//currentHeader.OpType - OPTYPE_CUSTOMERCONFIG
			//HTTP_HEADER : currentHeader
			//
			//Retval :
			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
			}
		} else {

		}
	} else {
		fmt.Println(WasteLibrary.GetUserIdByToken("MSMkMmEkMTAkOGVMTmlVM1Nqc0hockI1alJCcG5ZdWdQdFU1a2FpMDluazJtUXZxUEtXSVp3SUxVc2FyZks="))
		return
	}
	var urlFull string = "http://" + urlVal + "/" + path1 + "/" + path2
	fmt.Println(urlFull)
	resultVal := WasteLibrary.HttpPostReq(urlFull, data)
	fmt.Println(resultVal)
}
