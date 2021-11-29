package main

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/devafatek/WasteLibrary"
)

const (
	AWS_S3_REGION = "eu-central-1"
	AWS_S3_BUCKET = "afatek-waste-webui-s3"
)

func initStart() {

	WasteLibrary.LogStr("Successfully connected!")
	go WasteLibrary.InitLog()
}
func main() {

	initStart()
	downloadFolder("WEB_BASE/dist/")
	http.HandleFunc("/health", WasteLibrary.HealthHandler)
	http.HandleFunc("/readiness", WasteLibrary.ReadinessHandler)
	http.HandleFunc("/status", WasteLibrary.StatusHandler)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.HandleFunc("/", root)

	http.ListenAndServe(":80", nil)
}

func root(w http.ResponseWriter, req *http.Request) {

	if WasteLibrary.AllowCors {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
	}
	http.ServeFile(w, req, req.URL.Path[1:])
}

func downloadFolder(folderName string) {
	sess, err := session.NewSession(
		&aws.Config{Region: aws.String(AWS_S3_REGION)},
	)
	if err != nil {
		WasteLibrary.LogErr(err)
		return
	}
	svc := s3.New(sess)

	res, err := svc.ListObjects(&s3.ListObjectsInput{
		Bucket: aws.String(AWS_S3_BUCKET),
		Prefix: aws.String(folderName),
	})
	if err != nil {
		WasteLibrary.LogErr(err)
		return
	}
	for _, object := range res.Contents {
		var filePath = *object.Key

		if *object.Key == folderName {
			continue
		}
		filePath = strings.Replace(filePath, folderName, "", -1)
		splPath := strings.Split(filePath, "/")
		folderPath := strings.Replace(filePath, splPath[len(splPath)-1], "", -1)
		WasteLibrary.LogStr(filePath)
		WasteLibrary.LogStr(folderPath)

		params := &s3.GetObjectInput{Bucket: aws.String(AWS_S3_BUCKET), Key: aws.String(folderName + filePath)}
		res, err := svc.GetObject(params)
		if err != nil {
			WasteLibrary.LogErr(err)
			return
		}

		defer res.Body.Close()

		if _, err := os.Stat(folderPath); os.IsNotExist(err) {
			os.MkdirAll(folderPath, 0700)
		}
		outFile, err := os.Create(filePath)
		if err != nil {
			WasteLibrary.LogErr(err)
			return
		}
		defer outFile.Close()
		io.Copy(outFile, res.Body)

	}

}
