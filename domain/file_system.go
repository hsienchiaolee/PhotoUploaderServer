package domain

import (
	"io"
	"log"
)

type File interface {
	io.Closer
	io.Reader
	io.Writer
}

type OperatingSystem interface {
	Create(name string) (File, error)
}

type InputOutput interface {
	Copy(dst File, src io.Reader) (written int64, err error)
}

type FileSystem interface {
	Save(name string, src io.Reader) (written int64, error error)
}

type fileSystem struct {
	operatingSystem OperatingSystem
	inputOutput     InputOutput
}

func NewFileSystem(
	operatingSystem OperatingSystem,
	inputOutput InputOutput,
) FileSystem {
	return fileSystem{
		operatingSystem: operatingSystem,
		inputOutput:     inputOutput,
	}
}

func (fs fileSystem) Save(name string, src io.Reader) (written int64, error error) {
	out, error := fs.operatingSystem.Create(name)
	if error != nil {
		log.Printf("Unable to create file: %v", error)
	}
	defer out.Close()
	
	return fs.inputOutput.Copy(out, src)
}
