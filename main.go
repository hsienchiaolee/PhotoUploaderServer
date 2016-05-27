package main

import (
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/hsienchiaolee/PhotoUploaderServer/api"
	"github.com/hsienchiaolee/PhotoUploaderServer/domain"
)

func main() {
	log.Printf("Starting Server")

	fileSystem := domain.NewOsFileSystem()
	s3Service := s3.New(session.New(), &aws.Config{
		Region:      aws.String("us-west-1"),
		Credentials: credentials.NewSharedCredentials("", "go"),
	})

	http.ListenAndServe(":8080", api.Handlers(fileSystem, s3Service))
}
