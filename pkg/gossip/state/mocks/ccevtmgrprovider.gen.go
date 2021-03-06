// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	ccapi "github.com/hyperledger/fabric/extensions/chaincode/api"
)

type CCEventMgrProvider struct {
	GetMgrStub        func() ccapi.EventMgr
	getMgrMutex       sync.RWMutex
	getMgrArgsForCall []struct{}
	getMgrReturns     struct {
		result1 ccapi.EventMgr
	}
	getMgrReturnsOnCall map[int]struct {
		result1 ccapi.EventMgr
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *CCEventMgrProvider) GetMgr() ccapi.EventMgr {
	fake.getMgrMutex.Lock()
	ret, specificReturn := fake.getMgrReturnsOnCall[len(fake.getMgrArgsForCall)]
	fake.getMgrArgsForCall = append(fake.getMgrArgsForCall, struct{}{})
	fake.recordInvocation("GetMgr", []interface{}{})
	fake.getMgrMutex.Unlock()
	if fake.GetMgrStub != nil {
		return fake.GetMgrStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.getMgrReturns.result1
}

func (fake *CCEventMgrProvider) GetMgrCallCount() int {
	fake.getMgrMutex.RLock()
	defer fake.getMgrMutex.RUnlock()
	return len(fake.getMgrArgsForCall)
}

func (fake *CCEventMgrProvider) GetMgrReturns(result1 ccapi.EventMgr) {
	fake.GetMgrStub = nil
	fake.getMgrReturns = struct {
		result1 ccapi.EventMgr
	}{result1}
}

func (fake *CCEventMgrProvider) GetMgrReturnsOnCall(i int, result1 ccapi.EventMgr) {
	fake.GetMgrStub = nil
	if fake.getMgrReturnsOnCall == nil {
		fake.getMgrReturnsOnCall = make(map[int]struct {
			result1 ccapi.EventMgr
		})
	}
	fake.getMgrReturnsOnCall[i] = struct {
		result1 ccapi.EventMgr
	}{result1}
}

func (fake *CCEventMgrProvider) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getMgrMutex.RLock()
	defer fake.getMgrMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *CCEventMgrProvider) recordInvocation(key string, args []interface{}) {
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
