package main

import (
	"fmt"
	"net"
	"os"
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

func tcpServer() {
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		WasteLibrary.LogErr(err)
		os.Exit(1)
	}
	defer l.Close()
	WasteLibrary.LogStr("Listening on " + CONN_HOST + ":" + CONN_PORT)
	for {
		conn, err := l.Accept()
		if err != nil {
			WasteLibrary.LogErr(err)
			os.Exit(1)
		}
		connCount++
		fmt.Println(connCount)
		go handleTcpRequest(conn)
	}
	wg.Done()
}

func handleTcpRequest(conn net.Conn) {

	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		WasteLibrary.LogErr(err)
	}
	var strBuf = string(buf)
	fmt.Println(strBuf)
}
