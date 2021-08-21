package datastore

import (
	"log"
	"mime/multipart"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func CreateSessionInstace(awsProfile string) *session.Session {

	sessionInstance := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Profile:           awsProfile,
	}))

	return sessionInstance
}

func CreateS3UploaderInstance(sess *session.Session) *s3manager.Uploader {

	return s3manager.NewUploader(sess)

}

func UploadFileFromPath(filePath string, key string, bucketName string, uploader *s3manager.Uploader) error {

	f, err := os.Open(filePath)

	if err != nil {
		log.Panicf("failed to open file %q, %v", "", err)
	}

	_, err = uploader.S3.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
		Body:   f,
	})

	if err != nil {
		log.Panicf("failed to upload file %q, %v", key, err)
	}

	return err

}

func UploadFile(file multipart.File, key string, bucketName string, uploader *s3manager.Uploader) error {

	_, err := uploader.S3.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
		Body:   file,
	})

	if err != nil {
		log.Panicf("failed to upload file %q, %v", key, err)
	}

	return err

}

func DownloadFile(filePath string, key string, bucketName string, downloader *s3manager.Downloader) error {

	f, err := os.Create(filePath)

	if err != nil {
		log.Panicf("failed to create file %q, %v", "", err)
	}

	_, err = downloader.Download(f,
		&s3.GetObjectInput{
			Bucket: aws.String(bucketName),
			Key:    aws.String(key),
		})

	if err != nil {
		log.Panicf("failed to download file %q, %v", key, err)
	}

	return err

}

func DeleteFile(key string, bucketName string, sess *session.Session) error {

	svc := s3.New(sess)

	_, err := svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	})

	if err != nil {
		log.Panicf("failed to delete file %q, %v", key, err)
	}

	return err

}

func ListFiles(bucketName string, sess *session.Session) (*s3.ListObjectsOutput, error) {

	svc := s3.New(sess)

	result, err := svc.ListObjects(&s3.ListObjectsInput{
		Bucket: aws.String(bucketName),
	})

	if err != nil {
		log.Panicf("failed to list objects, %v", err)
	}

	return result, err

}
