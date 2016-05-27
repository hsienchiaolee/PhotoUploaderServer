package service

import "github.com/aws/aws-sdk-go/service/s3"

type S3Service interface {
	PutObject(input *s3.PutObjectInput) (*s3.PutObjectOutput, error)
}

type s3Service struct{}

func (service s3Service) PutObject(input *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	return service.PutObject(input)
}
