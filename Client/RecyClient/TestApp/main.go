package main

import (
	"bytes"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/devafatek/WasteLibrary"
)

const (
	AWS_S3_REGION = "eu-central-1"
	AWS_S3_BUCKET = "afatek-waste-images-s3"
)

func main() {
	session, err := session.NewSession(&aws.Config{Region: aws.String(AWS_S3_REGION)})
	WasteLibrary.LogErr(err)
	err = uploadFile(session, "WAIT_CAM/1.png")
	if err != nil {
		WasteLibrary.LogErr(err)
	} else {
		WasteLibrary.RemoveFile("WAIT_CAM/1.png")
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
