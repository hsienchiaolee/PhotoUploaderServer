package api_test

import (
	"io"
	"os"
	"net/http"
	"net/http/httptest"
	
	"io/ioutil"
	"bytes"
	"mime/multipart"
	"path/filepath"

	"testing"

	"github.com/hsienchiaolee/PhotoUploaderServer/api"
	"github.com/hsienchiaolee/PhotoUploaderServer/domain/domainfakes"
)

var (
	fileSystem *domainfakes.FakeFileSystem
	server     *httptest.Server
	uploadUrl  string
)

func init() {
	fileSystem = new(domainfakes.FakeFileSystem)
	server = httptest.NewServer(api.Handlers(fileSystem))
	uploadUrl = server.URL + "/upload"
}

func TestFileUpload(t *testing.T) {
	fileName := "someFile.txt"
	fileSystem.SaveStub = func(name string, src io.Reader) (written int64, error error) {
		if name != "./files/" + fileName {
			t.Errorf("Expected fileSystem.Save to be called with ./files/%s, but instead got %s", fileName, name)
		}

		if src == nil {
			t.Error("Expected fileSystem.Save to be called with non-nil source file")
		}
		return 0, nil
	}
	
	wd, _ := os.Getwd()
	filePath := filepath.Dir(wd) + "/" + fileName
	request := multipartRequest(uploadUrl, filePath)	

	response, error := http.DefaultClient.Do(request)

	if error != nil {
		t.Error(error)
	}

	if response.StatusCode != 200 {
		t.Errorf("Success expected: %d", response.StatusCode)
	}

	if fileSystem.SaveCallCount() != 1 {
		t.Error("Expected fileSystem.Save to be called")
	}
}

func multipartRequest(url string, filePath string) (*http.Request) {
	file, _ := os.Open(filePath)
	fileContents, _ := ioutil.ReadAll(file)
	fileInfo, _ := file.Stat()
	file.Close()
	
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", fileInfo.Name())
	part.Write(fileContents)

	writer.Close()
	
	request, _ := http.NewRequest("POST", url, body)
	request.Header.Add("Content-Type", writer.FormDataContentType())
	return request
}
