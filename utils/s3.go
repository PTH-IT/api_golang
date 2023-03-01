package utils

import (
	"PTH-IT/api_golang/config"
	"fmt"
	"mime/multipart"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func NewSession() *session.Session {
	awsEndpoint := fmt.Sprintf("%s:%s", config.Getconfig().Aws.Host, config.Getconfig().Aws.Port)
	awsRegion := config.Getconfig().Aws.Region
	id := config.Getconfig().Aws.Id
	secret := config.Getconfig().Aws.Secret
	token := config.Getconfig().Aws.Token
	sess := session.Must(session.NewSession(&aws.Config{
		Region:           aws.String(awsRegion),
		Endpoint:         aws.String(awsEndpoint),
		S3ForcePathStyle: aws.Bool(true),
		Credentials:      credentials.NewStaticCredentials(id, secret, token),
	}))
	return sess
}
func UpdateManager(myBucket string, myString string, filename string) {
	sess := NewSession()
	uploader := s3manager.NewUploader(sess)
	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("failed to open file %q, %v", filename, err)
	}
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(myBucket),
		Key:    aws.String(myString),
		Body:   f,
	})
	fmt.Printf("file uploaded to, %s\n", result.Location)
}
func DownloadManager(myBucket string, myString string, filename string) {
	// The session the S3 Downloader will use
	sess := NewSession()
	// Create a downloader with the session and default options
	downloader := s3manager.NewDownloader(sess)

	// Create a file to write the S3 Object contents to.
	f, err := os.Create(filename)
	if err != nil {
		fmt.Printf("failed to create file %q, %v", filename, err)
	}

	// Write the contents of S3 Object to the file
	n, err := downloader.Download(f, &s3.GetObjectInput{
		Bucket: aws.String(myBucket),
		Key:    aws.String(myString),
	})
	if err != nil {
		fmt.Printf("failed to download file, %v", err)
	}
	fmt.Printf("file downloaded, %d bytes\n", n)
}
func AddManager(myBucket string, myString string, f multipart.File) (string, error) {
	sess := NewSession()
	uploader := s3manager.NewUploader(sess, func(u *s3manager.Uploader) {
		// Define a strategy that will buffer 25 MiB in memory
		u.BufferProvider = s3manager.NewBufferedReadSeekerWriteToPool(2500 * 1024 * 1024)
	})
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(myBucket),
		Key:    aws.String(myString),
		Body:   f,
	})
	if err != nil {
		return "", err
	}
	return result.Location, err
}
