package api

import (
	"log"
	"net/http"

	"github.com/hsienchiaolee/PhotoUploaderServer/domain"
)

func NewUploadHandler(fileSystem domain.FileSystem) http.Handler {
	return uploadHandler{fileSystem: fileSystem}
}

type uploadHandler struct {
	fileSystem domain.FileSystem
}

func (handler uploadHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	file, header, error := request.FormFile("file")
	if error != nil {
		log.Printf("request: %v; Error: %v;", request, error)
		return
	}

	defer file.Close()

	_, error =	handler.fileSystem.Save("./files/" + header.Filename, file)
	if error != nil {
		log.Printf("Unable to save file: %v", error)
	}
}
