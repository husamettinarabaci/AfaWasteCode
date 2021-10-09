package main

import (
	"fmt"
	"net/url"

	"github.com/devafatek/WasteLibrary"
)

func main() {

	var opType string = "CUSTOMER"
	var opType2 string = "SET_CUSTOMER"
	var currentHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.HttpClientHeaderType{}
	var urlVal string = "afatek.aws.afatek.com.tr"
	var path1 string = "webapi"
	var path2 string = "getLink"
	data := url.Values{}
	var deviceId float64 = 0
	var customerId float64 = 0
	var userId float64 = 0
	currentHeader.DeviceId = deviceId
	currentHeader.CustomerId = customerId
	currentHeader.Token = ""

	if opType == "DEVICE" {
		if opType2 == "GET_DEVICE_WEB" {
			path1 = "webapi"
			path2 = "getDevice"
			//currentData - DeviceType
			//currentData.DevıceId - deviceId
			//HTTP_DATA : currentData
			var currentData WasteLibrary.DeviceType = WasteLibrary.DeviceType{
				DeviceId: deviceId,
			}

			data = url.Values{
				WasteLibrary.HTTP_DATA: {currentData.ToString()},
			}
		} else if opType2 == "GET_DEVICE_AFATEK" {
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

			var currentData WasteLibrary.DeviceType = WasteLibrary.DeviceType{
				DeviceId: deviceId,
			}

			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
				WasteLibrary.HTTP_DATA:   {currentData.ToString()},
			}

		} else if opType2 == "SET_DEVICE" {
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
			path1 = "webapi"
			path2 = "getDevices"

			data = url.Values{}
		} else if opType2 == "GET_DEVICES_AFATEK" {
			path1 = "afatekapi"
			path2 = "getDevices"
			//currentHeader - HttpClientHeaderType
			//currentHeader.CustomerId - customerId
			//currentHeader.Token - token
			//HTTP_HEADER : currentHeader

			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
			}
		} else {

		}
	} else if opType == "CUSTOMER" {
		if opType2 == "GET_CUSTOMER_WEB" {
			path1 = "webapi"
			path2 = "getCustomer"

			data = url.Values{}
		} else if opType2 == "GET_CUSTOMER_ADMIN" {
			path1 = "adminapi"
			path2 = "getCustomer"
			//currentHeader - HttpClientHeaderType
			//currentHeader.Token - token
			//HTTP_HEADER : currentHeader

			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
			}
		} else if opType2 == "GET_CUSTOMER_AFATEK" {
			path1 = "afatekapi"
			path2 = "getCustomer"
			//currentHeader - HttpClientHeaderType
			//currentHeader.Token - token
			//HTTP_HEADER : currentHeader
			//
			//currentData - CustomerType
			//currentData.CustomerId - customerId
			//HTTP_DATA : currentData

			var currentData WasteLibrary.CustomerType = WasteLibrary.CustomerType{
				CustomerId: customerId,
			}

			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
				WasteLibrary.HTTP_DATA:   {currentData.ToString()},
			}
		} else if opType2 == "SET_CUSTOMER" {
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
			path1 = "afatekapi"
			path2 = "getCustomers"
			//currentHeader - HttpClientHeaderType
			//currentHeader.Token - token
			//HTTP_HEADER : currentHeader

			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
			}
		} else {

		}
	} else if opType == "USER" {
		if opType2 == "GET_USER" {
			path1 = "adminapi"
			path2 = "getUser"
			//currentHeader - HttpClientHeaderType
			//currentHeader.Token - token
			//HTTP_HEADER : currentHeader
			//
			//currentData - UserType
			//currentData.UserId - userId
			//HTTP_DATA : currentData

			var currentData WasteLibrary.UserType = WasteLibrary.UserType{
				UserId: userId,
			}

			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
				WasteLibrary.HTTP_DATA:   {currentData.ToString()},
			}
		} else if opType2 == "SET_USER" {
			path1 = "adminapi"
			path2 = "setUser"
			//currentHeader - HttpClientHeaderType
			//currentHeader.Token - token
			//HTTP_HEADER : currentHeader
			//
			//currentData - UserType
			//currentData.UserId - [0,userId]
			//currentData.[All]
			//HTTP_DATA : currentData

			var currentData WasteLibrary.UserType = WasteLibrary.UserType{
				UserId: userId,
			}

			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
				WasteLibrary.HTTP_DATA:   {currentData.ToString()},
			}
		} else if opType2 == "LOGIN" {
			path1 = "authapi"
			path2 = "login"
		} else if opType2 == "REGISTER" {
			path1 = "authapi"
			path2 = "register"
		} else if opType2 == "GET_USERS" {
			path1 = "adminapi"
			path2 = "getUsers"
			//currentHeader - HttpClientHeaderType
			//currentHeader.Token - token
			//HTTP_HEADER : currentHeader

			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
			}
		} else {

		}
	} else if opType == "CONFIG" {
		if opType2 == "GET_ADMINCONFIG_ADMIN" {
			path1 = "adminapi"
			path2 = "getConfig"
			//currentHeader - HttpClientHeaderType
			//currentHeader.Token - token
			//currentHeader.OpType - OPTYPE_ADMINCONFIG
			//HTTP_HEADER : currentHeader

			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
			}
		} else if opType2 == "SET_ADMINCONFIG_ADMIN" {
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
			path1 = "adminapi"
			path2 = "getConfig"
			//currentHeader - HttpClientHeaderType
			//currentHeader.Token - token
			//currentHeader.OpType - OPTYPE_LOCALCONFIG
			//HTTP_HEADER : currentHeader

			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
			}
		} else if opType2 == "SET_LOCALCONFIG_ADMIN" {
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
			path1 = "adminapi"
			path2 = "getConfig"
			//currentHeader - HttpClientHeaderType
			//currentHeader.Token - token
			//currentHeader.OpType - OPTYPE_CUSTOMERCONFIG
			//HTTP_HEADER : currentHeader

			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
			}
		} else if opType2 == "SET_CUSTOMERCONFIG_ADMIN" {
			path1 = "adminapi"
			path2 = "setConfig"
			//currentHeader - HttpClientHeaderType
			//currentHeader.Token - token
			//currentHeader.OpType - OPTYPE_CUSTOMERCONFIG
			//HTTP_HEADER : currentHeader
			//
			//currentData - CustomerConfigType
			//currentData.[All]
			//HTTP_DATA : currentData

			var currentData WasteLibrary.CustomerConfigType = WasteLibrary.CustomerConfigType{}

			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
				WasteLibrary.HTTP_DATA:   {currentData.ToString()},
			}
		} else if opType2 == "GET_ADMINCONFIG_WEB" {
			path1 = "webapi"
			path2 = "getConfig"
			//currentHeader - HttpClientHeaderType
			//currentHeader.OpType - OPTYPE_ADMINCONFIG
			//HTTP_HEADER : currentHeader

			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
			}
		} else if opType2 == "GET_LOCALCONFIG_WEB" {
			path1 = "webapi"
			path2 = "getConfig"
			//currentHeader - HttpClientHeaderType
			//currentHeader.OpType - OPTYPE_LOCALCONFIG
			//HTTP_HEADER : currentHeader

			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
			}
		} else if opType2 == "GET_CUSTOMERCONFIG_WEB" {
			path1 = "webapi"
			path2 = "getConfig"
			//currentHeader - HttpClientHeaderType
			//currentHeader.OpType - OPTYPE_CUSTOMERCONFIG
			//HTTP_HEADER : currentHeader

			data = url.Values{
				WasteLibrary.HTTP_HEADER: {currentHeader.ToString()},
			}
		} else {

		}
	} else {

	}

	resultVal := WasteLibrary.HttpPostReq(urlVal+"/"+path1+"/"+path2, data)
	fmt.Println(resultVal)

}
