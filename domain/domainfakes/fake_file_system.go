// This file was generated by counterfeiter
package domainfakes

import (
	"io"
	"sync"

	"github.com/hsienchiaolee/PhotoUploaderServer/domain"
)

type FakeFileSystem struct {
	SaveStub        func(name string, src io.Reader) (written int64, error error)
	saveMutex       sync.RWMutex
	saveArgsForCall []struct {
		name string
		src  io.Reader
	}
	saveReturns struct {
		result1 int64
		result2 error
	}
	invocations map[string][][]interface{}
}

func (fake *FakeFileSystem) Save(name string, src io.Reader) (written int64, error error) {
	fake.saveMutex.Lock()
	fake.saveArgsForCall = append(fake.saveArgsForCall, struct {
		name string
		src  io.Reader
	}{name, src})
	fake.guard("Save")
	fake.invocations["Save"] = append(fake.invocations["Save"], []interface{}{name, src})
	fake.saveMutex.Unlock()
	if fake.SaveStub != nil {
		return fake.SaveStub(name, src)
	} else {
		return fake.saveReturns.result1, fake.saveReturns.result2
	}
}

func (fake *FakeFileSystem) SaveCallCount() int {
	fake.saveMutex.RLock()
	defer fake.saveMutex.RUnlock()
	return len(fake.saveArgsForCall)
}

func (fake *FakeFileSystem) SaveArgsForCall(i int) (string, io.Reader) {
	fake.saveMutex.RLock()
	defer fake.saveMutex.RUnlock()
	return fake.saveArgsForCall[i].name, fake.saveArgsForCall[i].src
}

func (fake *FakeFileSystem) SaveReturns(result1 int64, result2 error) {
	fake.SaveStub = nil
	fake.saveReturns = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *FakeFileSystem) Invocations() map[string][][]interface{} {
	return fake.invocations
}

func (fake *FakeFileSystem) guard(key string) {
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
}

var _ domain.FileSystem = new(FakeFileSystem)