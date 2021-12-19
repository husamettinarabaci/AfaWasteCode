package main

import (
	"fmt"

	"github.com/devafatek/WasteLibrary"
)

func main() {

	var conf WasteLibrary.CustomerConfigType
	conf.New()
	conf.CustomerId = 3
	conf.ArventoApp = WasteLibrary.STATU_ACTIVE
	conf.ArventoUserName = "afatekbilisim"
	conf.ArventoPin1 = "Amca151200!Furkan"
	conf.ArventoPin2 = "Amca151200!Furkan"

	fmt.Println(conf.ToString())

}
