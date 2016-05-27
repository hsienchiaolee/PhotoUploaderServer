package api

import (
	"net/http"

	"github.com/hsienchiaolee/PhotoUploaderServer/domain"
	"github.com/hsienchiaolee/PhotoUploaderServer/service"
)

func Handlers(fileSystem domain.FileSystem, s3Service service.S3Service) *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/upload", NewUploadHandler(fileSystem).ServeHTTP)
	router.HandleFunc("/s3", NewUploadS3Handler(s3Service).ServeHTTP)
	return router
}
