package main

import (
	"bytes"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
	AWS_S3_REGION = "eu-central-1"
	AWS_S3_BUCKET = "afatek-waste-videos-s3"
)

func main() {

	session, err := session.NewSession(&aws.Config{
		Region:      aws.String(AWS_S3_REGION),
		Credentials: credentials.NewStaticCredentials("AKIA2B2SBK7OQLT4GSVK", "n8UDTfpf+vyjLfXcyzeV0SmLZWtJMBHQlwSoiFK6", "TOKEN"),
	})
	if err != nil {
		log.Fatal(err)
	}

	err = uploadFile(session, "0cb932a0-06a5-11ec-b73c-b827ebb1d188.mp4")
	if err != nil {
		log.Fatal(err)
	}

}

func uploadFile(session *session.Session, uploadFileDir string) error {

	upFile, err := os.Open(uploadFileDir)
	if err != nil {
		return err
	}
	defer upFile.Close()

	upFileInfo, _ := upFile.Stat()
	var fileSize int64 = upFileInfo.Size()
	fileBuffer := make([]byte, fileSize)
	upFile.Read(fileBuffer)

	_, err = s3.New(session).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(AWS_S3_BUCKET),
		Key:                  aws.String(uploadFileDir),
		ACL:                  aws.String("private"),
		Body:                 bytes.NewReader(fileBuffer),
		ContentLength:        aws.Int64(fileSize),
		ContentType:          aws.String(http.DetectContentType(fileBuffer)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
	})
	return err
}

/*import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func main() {
	data := url.Values{
		"HASHKEY":  {"serial-customer"},
		"SUBKEY":   {"00000000c1b1d187"},
		"KEYVALUE": {"56"},
	}
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.PostForm("http://127.0.0.1:8080/setkey", data)
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

func logErr(err error) {
	if err != nil {
		logStr(err.Error())
	}
}

func logStr(value ...interface{}) {

	fmt.Println(value)

}*/
