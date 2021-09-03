package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/AfatekDevelopers/result_lib_go/devafatekresult"
)

func main() {
	var resultVal devafatekresult.ResultType
	data := url.Values{
		"APPTYPE":      {"ADMIN"},
		"OPTYPE":       {"CUSTOMER"},
		"CUSTOMERID":   {"0"},
		"CUSTOMERNAME": {"Afatek 2"},
		"DOMAIN":       {"akilli.afatek.com.tr"},
		"RFIDAPP":      {"1"},
		"ULTAPP":       {"1"},
		"RECYAPP":      {"1"},
	}
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.PostForm("http://a579dddf21ea745a49c7237860760244-1808420299.eu-central-1.elb.amazonaws.com/setCustomer", data)

	if err != nil {
		logErr(err)

	} else {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logErr(err)
		}
		resultVal = devafatekresult.ByteToResultType(bodyBytes)
		logStr(resultVal.ToString())
	}

}

func logErr(err error) {
	if err != nil {
		logStr(err.Error())
	}
}

func logStr(value ...interface{}) {

	fmt.Println(value)

}
