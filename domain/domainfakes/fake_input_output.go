// This file was generated by counterfeiter
package domainfakes

import (
	"io"
	"sync"

	"github.com/hsienchiaolee/PhotoUploaderServer/domain"
)

type FakeInputOutput struct {
	CopyStub        func(dst domain.File, src io.Reader) (written int64, err error)
	copyMutex       sync.RWMutex
	copyArgsForCall []struct {
		dst domain.File
		src io.Reader
	}
	copyReturns struct {
		result1 int64
		result2 error
	}
	invocations map[string][][]interface{}
}

func (fake *FakeInputOutput) Copy(dst domain.File, src io.Reader) (written int64, err error) {
	fake.copyMutex.Lock()
	fake.copyArgsForCall = append(fake.copyArgsForCall, struct {
		dst domain.File
		src io.Reader
	}{dst, src})
	fake.guard("Copy")
	fake.invocations["Copy"] = append(fake.invocations["Copy"], []interface{}{dst, src})
	fake.copyMutex.Unlock()
	if fake.CopyStub != nil {
		return fake.CopyStub(dst, src)
	} else {
		return fake.copyReturns.result1, fake.copyReturns.result2
	}
}

func (fake *FakeInputOutput) CopyCallCount() int {
	fake.copyMutex.RLock()
	defer fake.copyMutex.RUnlock()
	return len(fake.copyArgsForCall)
}

func (fake *FakeInputOutput) CopyArgsForCall(i int) (domain.File, io.Reader) {
	fake.copyMutex.RLock()
	defer fake.copyMutex.RUnlock()
	
	return fake.copyArgsForCall[i].dst, fake.copyArgsForCall[i].src
}

func (fake *FakeInputOutput) CopyReturns(result1 int64, result2 error) {
	fake.CopyStub = nil
	fake.copyReturns = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *FakeInputOutput) Invocations() map[string][][]interface{} {
	return fake.invocations
}

func (fake *FakeInputOutput) guard(key string) {
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
}

var _ domain.InputOutput = new(FakeInputOutput)
