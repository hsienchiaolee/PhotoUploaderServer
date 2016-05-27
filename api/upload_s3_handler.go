package api

import (
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/hsienchiaolee/PhotoUploaderServer/service"
)

func NewUploadS3Handler(s3Service service.S3Service) http.Handler {
	return uploadS3Handler{s3Service: s3Service}
}

type uploadS3Handler struct {
	s3Service service.S3Service
}

func (handler uploadS3Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	file, header, error := request.FormFile("file")
	if error != nil {
		log.Printf("request: %v; Error: %v;", request, error)
		return
	}

	defer file.Close()

	params := &s3.PutObjectInput{
		Bucket: aws.String("kaisky"),
		Key:    aws.String("/go/" + header.Filename),
		Body:   file,
	}
	_, error = handler.s3Service.PutObject(params)
	if error != nil {
		log.Printf("Uploading to s3 failed: %v", error)
	}
}
