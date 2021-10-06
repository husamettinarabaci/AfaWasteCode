package main

import (
	"net/http"

	"github.com/devafatek/WasteLibrary"
)

func initStart() {

	WasteLibrary.LogStr("Successfully connected!")

}

func main() {

	initStart()

	WasteLibrary.LogStr("Start")
	http.HandleFunc("/health", WasteLibrary.HealthHandler)
	http.HandleFunc("/readiness", WasteLibrary.ReadinessHandler)
	http.HandleFunc("/status", WasteLibrary.StatusHandler)
	http.HandleFunc("/getkey", getkey)
	http.HandleFunc("/setkey", setkey)
	http.HandleFunc("/saveStaticDbMain", saveStaticDbMain)
	http.HandleFunc("/getStaticDbMain", getStaticDbMain)
	http.HandleFunc("/saveConfigDbMain", saveConfigDbMain)
	http.HandleFunc("/getConfigDbMain", getConfigDbMain)
	http.HandleFunc("/saveReaderDbMain", saveReaderDbMain)
	http.HandleFunc("/getReaderDbMain", getReaderDbMain)
	http.HandleFunc("/saveBulkDbMain", saveBulkDbMain)
	http.HandleFunc("/getBulkDbMain", getBulkDbMain)
	http.ListenAndServe(":80", nil)

}

func saveBulkDbMain(w http.ResponseWriter, req *http.Request) {
	resultVal := WasteLibrary.HttpPostReq("http://waste-storeapiforbulkdb-cluster-ip/saveBulkDbMain", req.Form)
	w.Write(resultVal.ToByte())
}

func getBulkDbMain(w http.ResponseWriter, req *http.Request) {
	resultVal := WasteLibrary.HttpPostReq("http://waste-storeapiforbulkdb-cluster-ip/getBulkDbMain", req.Form)
	w.Write(resultVal.ToByte())
}

func saveReaderDbMain(w http.ResponseWriter, req *http.Request) {
	resultVal := WasteLibrary.HttpPostReq("http://waste-storeapiforreaderdb-cluster-ip/saveReaderDbMain", req.Form)
	w.Write(resultVal.ToByte())
}

func getReaderDbMain(w http.ResponseWriter, req *http.Request) {
	resultVal := WasteLibrary.HttpPostReq("http://waste-storeapiforreaderdb-cluster-ip/getReaderDbMain", req.Form)
	w.Write(resultVal.ToByte())
}

func saveConfigDbMain(w http.ResponseWriter, req *http.Request) {
	resultVal := WasteLibrary.HttpPostReq("http://waste-storeapiforconfigdb-cluster-ip/saveConfigDbMain", req.Form)
	w.Write(resultVal.ToByte())
}

func getConfigDbMain(w http.ResponseWriter, req *http.Request) {
	resultVal := WasteLibrary.HttpPostReq("http://waste-storeapiforconfigdb-cluster-ip/getConfigDbMain", req.Form)
	w.Write(resultVal.ToByte())
}

func saveStaticDbMain(w http.ResponseWriter, req *http.Request) {
	resultVal := WasteLibrary.HttpPostReq("http://waste-storeapiforstaticdb-cluster-ip/saveStaticDbMain", req.Form)
	w.Write(resultVal.ToByte())
}

func getStaticDbMain(w http.ResponseWriter, req *http.Request) {
	resultVal := WasteLibrary.HttpPostReq("http://waste-storeapiforstaticdb-cluster-ip/getStaticDbMain", req.Form)
	w.Write(resultVal.ToByte())
}

func getkey(w http.ResponseWriter, req *http.Request) {
	resultVal := WasteLibrary.HttpPostReq("http://waste-storeapiforredis-cluster-ip/getkey", req.Form)
	w.Write(resultVal.ToByte())
}

func setkey(w http.ResponseWriter, req *http.Request) {
	resultVal := WasteLibrary.HttpPostReq("http://waste-storeapiforredis-cluster-ip/setkey", req.Form)
	w.Write(resultVal.ToByte())
}
