package api

import (
	"net/http"

	"github.com/hsienchiaolee/PhotoUploaderServer/domain"
)

func Handlers(fileSystem domain.FileSystem) *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/upload", NewUploadHandler(fileSystem).ServeHTTP)
	return router
}
