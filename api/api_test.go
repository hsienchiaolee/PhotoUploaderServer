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
	"github.com/hsienchiaolee/PhotoUploaderServer/service/servicefakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("API handler", func() {
	var (
		fileSystem *domainfakes.FakeFileSystem
		s3Service  *servicefakes.FakeS3Service
		server     *httptest.Server
	)

	BeforeEach(func() {
		fileSystem = new(domainfakes.FakeFileSystem)
		s3Service = new(servicefakes.FakeS3Service)
		server = httptest.NewServer(api.Handlers(fileSystem, s3Service))
	})

	Describe("uploading a file", func() {
		var (
			fileName     string
			fileContents []byte
			response     *http.Response
		)

		BeforeEach(func() {
			wd, _ := os.Getwd()
			filePath := filepath.Dir(wd) + "/someFile.txt"
			fileName, fileContents = openFile(filePath)
		})

		Context("to local server", func() {
			BeforeEach(func() {
				fileSystem.SaveReturns(0, nil)

				uploadUrl := server.URL + "/upload"
				request := multipartRequest(uploadUrl, fileName, fileContents)

				response, _ = http.DefaultClient.Do(request)
			})

			It("saves the file to the correct path", func() {
				Expect(fileSystem.SaveCallCount()).To(Equal(1))

				name, src := fileSystem.SaveArgsForCall(0)
				srcData, _ := ioutil.ReadAll(src)
				Expect(name).To(Equal("./files/" + fileName))
				Expect(srcData).To(Equal(fileContents))
			})

			It("returns 200 http status code", func() {
				Expect(response.StatusCode).To(Equal(200))
			})
		})

		Context("to aws s3", func() {
			BeforeEach(func() {
				uploadUrl := server.URL + "/s3"
				request := multipartRequest(uploadUrl, fileName, fileContents)

				response, _ = http.DefaultClient.Do(request)
			})

			It("uploads the file to s3", func() {
				Expect(s3Service.PutObjectCallCount()).To(Equal(1))
			})

			It("returns 200 http status code", func() {
				Expect(response.StatusCode).To(Equal(200))
			})
		})
	})
})

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
