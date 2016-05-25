package api_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"log"
	"io"
	"testing"

	"github.com/hsienchiaolee/PhotoUploaderServer/api"
)

var (
	server    *httptest.Server
	reader    io.Reader
	uploadUrl string
)

func init() {
	server = httptest.NewServer(api.Handlers())
	uploadUrl = server.URL + "/upload"

	log.Printf(uploadUrl)
}

func TestFileUpload(t *testing.T) {
	reader := strings.NewReader("")
	request, error := http.NewRequest("POST", uploadUrl, reader)

	response, error := http.DefaultClient.Do(request)

	if error != nil {
		t.Error(error)
	}

	if response.StatusCode != 200 {
		t.Errorf("Success expected: %d", response.StatusCode)
	}
}
