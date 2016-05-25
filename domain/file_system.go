package domain

import (
	"os"
	"io"
	"log"
)

type File interface {
	io.Closer
	io.Reader
	io.Writer
}

// os interface
type OperatingSystem interface {
	Create(name string) (File, error)
}

type osOperatingSystem struct {}

func (osOperatingSystem) Create(name string) (File, error) {
	return os.Create(name)
}

// io interface
type InputOutput interface {
	Copy(dst File, src io.Reader) (written int64, err error)
}

type ioInputOutput struct {}

func (ioInputOutput) Copy(dst File, src io.Reader) (written int64, err error) {
	return io.Copy(dst, src)
}

// fileSystem
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

func NewOsFileSystem() FileSystem {
		return fileSystem{
			operatingSystem: osOperatingSystem{},
			inputOutput:     ioInputOutput{},
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
