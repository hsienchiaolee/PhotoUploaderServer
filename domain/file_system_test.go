package domain_test

import (
	"io"
	"bytes"
	"testing"

	"github.com/hsienchiaolee/PhotoUploaderServer/domain"
	"github.com/hsienchiaolee/PhotoUploaderServer/domain/domainfakes"
)

var (
	fakeOS  *domainfakes.FakeOperatingSystem
	fakeIO  *domainfakes.FakeInputOutput
	subject domain.FileSystem
)

func init() {
	fakeOS = new(domainfakes.FakeOperatingSystem)
	fakeIO = new(domainfakes.FakeInputOutput)
	subject = domain.NewFileSystem(fakeOS, fakeIO)
}

func TestSavingFile(t *testing.T) {
	filePath := "filePath"
	fakeFile := new(domainfakes.FakeFile)
	fakeOS.CreateStub = func(name string) (domain.File, error) {
		if name != filePath {
			t.Errorf("Expect fakeOS.create to be called with %s, but instead got %s", filePath, name)
		}
		return fakeFile, nil
	}

	fakeIO.CopyStub = func(dst domain.File, src io.Reader) (written int64, err error) {
		if dst != fakeFile {
			t.Errorf("Expect fakeOS.copy to be called with destination file %s, but instead got %s", fakeFile, dst)
		}

		buf := new(bytes.Buffer)
		buf.ReadFrom(src)
		sourceData := buf.String()
		if sourceData != "some data to be stored" {
			t.Errorf("Expect fakeOS.copy to be called with source data 'some data to be stored', but instead got %s", sourceData)
		}
		return 0, nil
	}

	reader := bytes.NewBufferString("some data to be stored")
	subject.Save(filePath, reader)

	if fakeOS.CreateCallCount() != 1 {
		t.Error("Expect os.Create to be called")
	}

	if fakeIO.CopyCallCount() != 1 {
		t.Error("Expect io.Copy to be called")
	}

	if fakeFile.CloseCallCount() != 1 {
		t.Error("Expect file.Close to be called")
	}
}
