package utils

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"io"
	"mime/multipart"
	"os"
)

func UploadFile(fileReader io.Reader, fileHeader *multipart.FileHeader) (string, error) {
	awsRegion := os.Getenv("AWS_REGION")
	awsAccessKey := os.Getenv("AWS_ACCESS_KEY")
	awsScretKey := os.Getenv("AWS_SECRET_KEY")
	bucketName := os.Getenv("AWS_BUCKET_NAME")

	awsSession, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region:      aws.String(awsRegion),
			Credentials: credentials.NewStaticCredentials(awsAccessKey, awsScretKey, ""),
		},
	})

	if err != nil {
		panic(err)
	}

	uploader := s3manager.NewUploader(awsSession)

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileHeader.Filename),
		Body:   fileReader,
	})
	if err != nil {
		return "", err
	}

	// Get the URL of the uploaded file
	url := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucketName, fileHeader.Filename)

	return url, nil
}
