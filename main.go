package main

import (
	"log"
	"net/http"
	"io"
	"os"
)

func uploadHandler(writer http.ResponseWriter, request *http.Request) {
	file, header, error := request.FormFile("file")
	if error != nil {
		log.Printf("request: %v; Error: %v;", error, request)
		return
	}

	defer file.Close()

	out, error := os.Create("./files/" + header.Filename)
	if error != nil {
		log.Printf("Unable to create the file: %v", error)
		return
	}

	defer out.Close()

	_, error = io.Copy(out, file)
	if error != nil {
		log.Printf("Unable to copy file: %v", error)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/upload", uploadHandler)
	http.ListenAndServe(":8080", mux)
}
