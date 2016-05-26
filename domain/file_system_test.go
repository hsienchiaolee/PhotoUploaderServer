package domain_test

import (
	"bytes"
	"log"

	"github.com/hsienchiaolee/PhotoUploaderServer/domain"
	"github.com/hsienchiaolee/PhotoUploaderServer/domain/domainfakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("FileSystem", func() {
	var (
		subject domain.FileSystem
		fakeOS  *domainfakes.FakeOperatingSystem
		fakeIO  *domainfakes.FakeInputOutput
	)

	BeforeEach(func() {
		fakeOS = new(domainfakes.FakeOperatingSystem)
		fakeIO = new(domainfakes.FakeInputOutput)
		log.Printf("OS: %v; IO: %v;", fakeOS, fakeIO)
		subject = domain.NewFileSystem(fakeOS, fakeIO)
	})

	Describe("saving a file", func() {
		var (
			filePath string
			fakeFile *domainfakes.FakeFile
			reader   *bytes.Buffer
		)

		BeforeEach(func() {
			filePath = "filePath"
			fakeFile = new(domainfakes.FakeFile)
			reader = bytes.NewBufferString("some data to be stored")

			fakeOS.CreateReturns(fakeFile, nil)
			fakeIO.CopyReturns(0, nil)

			subject.Save(filePath, reader)
		})

		It("creates the destination file with os in the correct file path", func() {
			Expect(fakeOS.CreateCallCount()).To(Equal(1))
			Expect(fakeOS.CreateArgsForCall(0)).To(Equal(filePath))
		})

		It("copies the file into the destination file", func() {
			Expect(fakeIO.CopyCallCount()).To(Equal(1))

			dst, src := fakeIO.CopyArgsForCall(0)
			Expect(dst).To(Equal(fakeFile))
			Expect(src).To(Equal(reader))
		})

		It("closes the file upon completion", func() {
			Expect(fakeFile.CloseCallCount()).To(Equal(1))
		})
	})
})
