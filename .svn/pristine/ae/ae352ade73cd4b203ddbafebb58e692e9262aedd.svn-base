package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"os/user"
	"strings"
	"sync"
	"time"
)

var debug bool = os.Getenv("DEBUG") == "1"
var opInterval time.Duration = 5 * 60
var wg sync.WaitGroup
var currentUser string
var currentTherm string = ""

func initStart() {

	currentTherm = ""
	if !debug {
		time.Sleep(60 * time.Second)
	}
	logStr("Successfully connected!")
	currentUser = getCurrentUser()
	logStr(currentUser)
}
func main() {

	initStart()

	time.Sleep(5 * time.Second)
	go thermCheck()
	wg.Add(1)
	time.Sleep(5 * time.Second)
	go sendTherm()
	wg.Add(1)

	http.HandleFunc("/status", status)
	http.ListenAndServe(":10004", nil)

	wg.Wait()
}

func status(w http.ResponseWriter, req *http.Request) {

	if err := req.ParseForm(); err != nil {
		logErr(err)
		return
	}
	opType := req.FormValue("OPTYPE")
	logStr(opType)

	if opType == "APP" {
		w.Write([]byte("OK"))
	} else {
		w.Write([]byte("FAIL"))
	}
}

func thermCheck() {
	if currentUser == "pi" {
		for {
			time.Sleep(opInterval * time.Second)
			cmd := exec.Command("vcgencmd", "measure_temp")
			var outb, errb bytes.Buffer
			cmd.Stdout = &outb
			cmd.Stderr = &errb
			err := cmd.Run()
			currentTherm = strings.TrimSuffix(outb.String(), "'C\n")
			logStr(currentTherm)
			if err != nil {
				logErr(err)

			}
		}
	}
	wg.Done()
}

func sendTherm() {

	for {
		time.Sleep(opInterval * time.Second)
		data := url.Values{
			"OPTYPE": {"THERM"},
			"THERM":  {string(currentTherm)},
		}
		client := http.Client{
			Timeout: 10 * time.Second,
		}
		resp, err := client.PostForm("http://127.0.0.1:10000/trans", data)
		if err != nil {
			logErr(err)

		} else {
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				logErr(err)
			}
			bodyString := string(bodyBytes)
			logStr(bodyString)
		}
	}
	wg.Done()
}

func getCurrentUser() string {
	user, err := user.Current()
	if err != nil {
		logErr(err)
	}

	username := user.Username
	return username
}

func logErr(err error) {
	if err != nil {
		logStr(err.Error())
	}
}

func logStr(value string) {
	if debug {
		fmt.Println(value)
	}
}
