// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/silk/config"
)

type Common struct {
	BasicSetupStub        func(deviceName string, local, peer config.DualAddress) error
	basicSetupMutex       sync.RWMutex
	basicSetupArgsForCall []struct {
		deviceName string
		local      config.DualAddress
		peer       config.DualAddress
	}
	basicSetupReturns struct {
		result1 error
	}
	basicSetupReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Common) BasicSetup(deviceName string, local config.DualAddress, peer config.DualAddress) error {
	fake.basicSetupMutex.Lock()
	ret, specificReturn := fake.basicSetupReturnsOnCall[len(fake.basicSetupArgsForCall)]
	fake.basicSetupArgsForCall = append(fake.basicSetupArgsForCall, struct {
		deviceName string
		local      config.DualAddress
		peer       config.DualAddress
	}{deviceName, local, peer})
	fake.recordInvocation("BasicSetup", []interface{}{deviceName, local, peer})
	fake.basicSetupMutex.Unlock()
	if fake.BasicSetupStub != nil {
		return fake.BasicSetupStub(deviceName, local, peer)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.basicSetupReturns.result1
}

func (fake *Common) BasicSetupCallCount() int {
	fake.basicSetupMutex.RLock()
	defer fake.basicSetupMutex.RUnlock()
	return len(fake.basicSetupArgsForCall)
}

func (fake *Common) BasicSetupArgsForCall(i int) (string, config.DualAddress, config.DualAddress) {
	fake.basicSetupMutex.RLock()
	defer fake.basicSetupMutex.RUnlock()
	return fake.basicSetupArgsForCall[i].deviceName, fake.basicSetupArgsForCall[i].local, fake.basicSetupArgsForCall[i].peer
}

func (fake *Common) BasicSetupReturns(result1 error) {
	fake.BasicSetupStub = nil
	fake.basicSetupReturns = struct {
		result1 error
	}{result1}
}

func (fake *Common) BasicSetupReturnsOnCall(i int, result1 error) {
	fake.BasicSetupStub = nil
	if fake.basicSetupReturnsOnCall == nil {
		fake.basicSetupReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.basicSetupReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Common) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.basicSetupMutex.RLock()
	defer fake.basicSetupMutex.RUnlock()
	return fake.invocations
}

func (fake *Common) recordInvocation(key string, args []interface{}) {
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
