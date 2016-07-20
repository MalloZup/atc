// This file was generated by counterfeiter
package metricfakes

import (
	"database/sql/driver"
	"sync"

	"github.com/concourse/atc/metric"
)

type FakeConn struct {
	PrepareStub        func(query string) (driver.Stmt, error)
	prepareMutex       sync.RWMutex
	prepareArgsForCall []struct {
		query string
	}
	prepareReturns struct {
		result1 driver.Stmt
		result2 error
	}
	CloseStub        func() error
	closeMutex       sync.RWMutex
	closeArgsForCall []struct{}
	closeReturns     struct {
		result1 error
	}
	BeginStub        func() (driver.Tx, error)
	beginMutex       sync.RWMutex
	beginArgsForCall []struct{}
	beginReturns     struct {
		result1 driver.Tx
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeConn) Prepare(query string) (driver.Stmt, error) {
	fake.prepareMutex.Lock()
	fake.prepareArgsForCall = append(fake.prepareArgsForCall, struct {
		query string
	}{query})
	fake.recordInvocation("Prepare", []interface{}{query})
	fake.prepareMutex.Unlock()
	if fake.PrepareStub != nil {
		return fake.PrepareStub(query)
	} else {
		return fake.prepareReturns.result1, fake.prepareReturns.result2
	}
}

func (fake *FakeConn) PrepareCallCount() int {
	fake.prepareMutex.RLock()
	defer fake.prepareMutex.RUnlock()
	return len(fake.prepareArgsForCall)
}

func (fake *FakeConn) PrepareArgsForCall(i int) string {
	fake.prepareMutex.RLock()
	defer fake.prepareMutex.RUnlock()
	return fake.prepareArgsForCall[i].query
}

func (fake *FakeConn) PrepareReturns(result1 driver.Stmt, result2 error) {
	fake.PrepareStub = nil
	fake.prepareReturns = struct {
		result1 driver.Stmt
		result2 error
	}{result1, result2}
}

func (fake *FakeConn) Close() error {
	fake.closeMutex.Lock()
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct{}{})
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if fake.CloseStub != nil {
		return fake.CloseStub()
	} else {
		return fake.closeReturns.result1
	}
}

func (fake *FakeConn) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeConn) CloseReturns(result1 error) {
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeConn) Begin() (driver.Tx, error) {
	fake.beginMutex.Lock()
	fake.beginArgsForCall = append(fake.beginArgsForCall, struct{}{})
	fake.recordInvocation("Begin", []interface{}{})
	fake.beginMutex.Unlock()
	if fake.BeginStub != nil {
		return fake.BeginStub()
	} else {
		return fake.beginReturns.result1, fake.beginReturns.result2
	}
}

func (fake *FakeConn) BeginCallCount() int {
	fake.beginMutex.RLock()
	defer fake.beginMutex.RUnlock()
	return len(fake.beginArgsForCall)
}

func (fake *FakeConn) BeginReturns(result1 driver.Tx, result2 error) {
	fake.BeginStub = nil
	fake.beginReturns = struct {
		result1 driver.Tx
		result2 error
	}{result1, result2}
}

func (fake *FakeConn) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.prepareMutex.RLock()
	defer fake.prepareMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.beginMutex.RLock()
	defer fake.beginMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeConn) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ metric.Conn = new(FakeConn)