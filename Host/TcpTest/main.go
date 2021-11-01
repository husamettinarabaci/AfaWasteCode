package main

import (
	"sync"

	"github.com/devafatek/WasteLibrary"
)

var wg sync.WaitGroup

func initStart() {

	WasteLibrary.LogStr("Successfully connected!")
}

const (
	CONN_HOST = "0.0.0.0"
	CONN_PORT = "20000"
	CONN_TYPE = "tcp"
)

var connCount = 0

func main() {

	initStart()

	go tcpServer()
	wg.Add(1)

	wg.Wait()

}
