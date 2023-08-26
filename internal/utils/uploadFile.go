package utils

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"io"
	"mime/multipart"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/config"
)

func UploadFile(fileReader io.Reader, fileHeader *multipart.FileHeader) (string, error) {
	config, err := config.LoadConfig(".")
	if err != nil {
		return "", err
	}

	awsSession, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region:      aws.String(config.AWSRegion),
			Credentials: credentials.NewStaticCredentials(config.AWSAccessKey, config.AWSSecretKey, ""),
		},
	})

	if err != nil {
		panic(err)
	}

	uploader := s3manager.NewUploader(awsSession)

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(config.AWSBucketName),
		Key:    aws.String(fileHeader.Filename),
		Body:   fileReader,
	})
	if err != nil {
		return "", err
	}

	// Get the URL of the uploaded file
	url := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", config.AWSBucketName, fileHeader.Filename)

	return url, nil
}
