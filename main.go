package main

import (
	"log"
	"net/http"
	
	"github.com/hsienchiaolee/PhotoUploaderServer/api"
	"github.com/hsienchiaolee/PhotoUploaderServer/domain"
)

func main() {
	log.Printf("Starting Server")
	fileSystem := domain.NewOsFileSystem()
	http.ListenAndServe(":8080", api.Handlers(fileSystem))
}
