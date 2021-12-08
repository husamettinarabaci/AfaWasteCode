package main

import (
	"fmt"

	"github.com/devafatek/WasteLibrary"
)

func main() {

	var customer WasteLibrary.CustomerType
	customer.New()
	customer.CustomerId = 5
	customer.CustomerName = "TestName"
	customer.CustomerLink = "TestLink"

	fmt.Println(customer.ToString())

	var customer2 WasteLibrary.CustomerType
	customer2 = WasteLibrary.StringToCustomerType(`{"CustomerId":5,"CustomerLink":"TestLink","UltApp":"PASSIVE","RecyApp":"PASSIVE","CreateTime":"2021-12-08T15:39:46+03:00"}`)
	fmt.Println(customer2.ToString())

}
