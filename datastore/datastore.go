package datastore

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type awsSession *session.Session

type s3Uploader *s3manager.Uploader

var (
	sessionInstance    awsSession
	s3UploaderInstance s3Uploader
)

func getSessionInstance() *session.Session {

	if sessionInstance == nil {
		sessionInstance = session.Must(session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
		}))
	}

	return sessionInstance
}

func getS3UploaderInstance() *s3manager.Uploader {

	if s3UploaderInstance == nil {
		s3UploaderInstance = s3manager.NewUploader(getSessionInstance())
	}

	return s3UploaderInstance

}

func UploadFile(filePath string, bucket string, key string) {

	f, err := os.Open(filePath)

	if err != nil {
		log.Fatalf("failed to open file %q, %v", "", err)
	}

	_, err = getS3UploaderInstance().S3.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   f,
	})

	if err != nil {
		log.Fatalf("failed to upload file %q, %v", "", err)
	}

}
