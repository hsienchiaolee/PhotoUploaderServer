package api_test

import (
	"net/http"
	"net/http/httptest"
	"os"

	"bytes"
	"io/ioutil"
	"mime/multipart"
	"path/filepath"

	"github.com/hsienchiaolee/PhotoUploaderServer/api"
	"github.com/hsienchiaolee/PhotoUploaderServer/domain/domainfakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func openFile(path string) (fileName string, content []byte) {
	file, _ := os.Open(path)
	defer file.Close()
	
	fileInfo, _ := file.Stat()
	fileContents, _ := ioutil.ReadAll(file)
	return fileInfo.Name(), fileContents
}

func multipartRequest(url string, fileName string, fileContents []byte) *http.Request {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", fileName)
	part.Write(fileContents)

	writer.Close()

	request, _ := http.NewRequest("POST", url, body)
	request.Header.Add("Content-Type", writer.FormDataContentType())
	return request
}

var _ = Describe("API handler", func() {
	var (
		fileSystem *domainfakes.FakeFileSystem
		server     *httptest.Server
		uploadUrl  string
	)

	BeforeEach(func() {
		fileSystem = new(domainfakes.FakeFileSystem)
		server = httptest.NewServer(api.Handlers(fileSystem))
		uploadUrl = server.URL + "/upload"
	})

	Describe("uploading a file", func() {
		var (
			fileName string
			fileContents []byte
			response *http.Response
		)

		BeforeEach(func() {
			fileSystem.SaveReturns(0, nil)

			wd, _ := os.Getwd()
			filePath := filepath.Dir(wd) + "/someFile.txt"
			fileName, fileContents = openFile(filePath)
			request := multipartRequest(uploadUrl, fileName, fileContents)

			response, _ = http.DefaultClient.Do(request)
		})

		It("saves the file to the correct path", func() {
			Expect(fileSystem.SaveCallCount()).To(Equal(1))

			name, src := fileSystem.SaveArgsForCall(0)
			Expect(name).To(Equal("./files/" + fileName))
			Expect(src).To(Equal(fileContents))
		})

		It("returns 200 http status code", func() {
			Expect(response.StatusCode).To(Equal(200))
		})
	})
})
