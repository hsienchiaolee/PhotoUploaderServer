package main

import (
	"log"
	"net/http"
	
	"github.com/hsienchiaolee/PhotoUploaderServer/api"
)

func main() {
	log.Printf("Starting Server")
	http.ListenAndServe(":8080", api.Handlers())
}
