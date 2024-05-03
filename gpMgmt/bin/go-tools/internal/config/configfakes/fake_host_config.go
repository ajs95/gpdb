// Code generated by counterfeiter. DO NOT EDIT.
package configfakes

import (
	"sync"

	"github.com/greenplum-db/gpdb/gp/internal/config"
)

type FakeHostConfig struct {
	GetAuthStub        func() config.AuthenticationConfig
	getAuthMutex       sync.RWMutex
	getAuthArgsForCall []struct {
	}
	getAuthReturns struct {
		result1 config.AuthenticationConfig
	}
	getAuthReturnsOnCall map[int]struct {
		result1 config.AuthenticationConfig
	}
	GetDomainNameStub        func() string
	getDomainNameMutex       sync.RWMutex
	getDomainNameArgsForCall []struct {
	}
	getDomainNameReturns struct {
		result1 string
	}
	getDomainNameReturnsOnCall map[int]struct {
		result1 string
	}
	GetHostnameStub        func() string
	getHostnameMutex       sync.RWMutex
	getHostnameArgsForCall []struct {
	}
	getHostnameReturns struct {
		result1 string
	}
	getHostnameReturnsOnCall map[int]struct {
		result1 string
	}
	GetIpStub        func() string
	getIpMutex       sync.RWMutex
	getIpArgsForCall []struct {
	}
	getIpReturns struct {
		result1 string
	}
	getIpReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeHostConfig) GetAuth() config.AuthenticationConfig {
	fake.getAuthMutex.Lock()
	ret, specificReturn := fake.getAuthReturnsOnCall[len(fake.getAuthArgsForCall)]
	fake.getAuthArgsForCall = append(fake.getAuthArgsForCall, struct {
	}{})
	stub := fake.GetAuthStub
	fakeReturns := fake.getAuthReturns
	fake.recordInvocation("GetAuth", []interface{}{})
	fake.getAuthMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeHostConfig) GetAuthCallCount() int {
	fake.getAuthMutex.RLock()
	defer fake.getAuthMutex.RUnlock()
	return len(fake.getAuthArgsForCall)
}

func (fake *FakeHostConfig) GetAuthCalls(stub func() config.AuthenticationConfig) {
	fake.getAuthMutex.Lock()
	defer fake.getAuthMutex.Unlock()
	fake.GetAuthStub = stub
}

func (fake *FakeHostConfig) GetAuthReturns(result1 config.AuthenticationConfig) {
	fake.getAuthMutex.Lock()
	defer fake.getAuthMutex.Unlock()
	fake.GetAuthStub = nil
	fake.getAuthReturns = struct {
		result1 config.AuthenticationConfig
	}{result1}
}

func (fake *FakeHostConfig) GetAuthReturnsOnCall(i int, result1 config.AuthenticationConfig) {
	fake.getAuthMutex.Lock()
	defer fake.getAuthMutex.Unlock()
	fake.GetAuthStub = nil
	if fake.getAuthReturnsOnCall == nil {
		fake.getAuthReturnsOnCall = make(map[int]struct {
			result1 config.AuthenticationConfig
		})
	}
	fake.getAuthReturnsOnCall[i] = struct {
		result1 config.AuthenticationConfig
	}{result1}
}

func (fake *FakeHostConfig) GetDomainName() string {
	fake.getDomainNameMutex.Lock()
	ret, specificReturn := fake.getDomainNameReturnsOnCall[len(fake.getDomainNameArgsForCall)]
	fake.getDomainNameArgsForCall = append(fake.getDomainNameArgsForCall, struct {
	}{})
	stub := fake.GetDomainNameStub
	fakeReturns := fake.getDomainNameReturns
	fake.recordInvocation("GetDomainName", []interface{}{})
	fake.getDomainNameMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeHostConfig) GetDomainNameCallCount() int {
	fake.getDomainNameMutex.RLock()
	defer fake.getDomainNameMutex.RUnlock()
	return len(fake.getDomainNameArgsForCall)
}

func (fake *FakeHostConfig) GetDomainNameCalls(stub func() string) {
	fake.getDomainNameMutex.Lock()
	defer fake.getDomainNameMutex.Unlock()
	fake.GetDomainNameStub = stub
}

func (fake *FakeHostConfig) GetDomainNameReturns(result1 string) {
	fake.getDomainNameMutex.Lock()
	defer fake.getDomainNameMutex.Unlock()
	fake.GetDomainNameStub = nil
	fake.getDomainNameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeHostConfig) GetDomainNameReturnsOnCall(i int, result1 string) {
	fake.getDomainNameMutex.Lock()
	defer fake.getDomainNameMutex.Unlock()
	fake.GetDomainNameStub = nil
	if fake.getDomainNameReturnsOnCall == nil {
		fake.getDomainNameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.getDomainNameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeHostConfig) GetHostname() string {
	fake.getHostnameMutex.Lock()
	ret, specificReturn := fake.getHostnameReturnsOnCall[len(fake.getHostnameArgsForCall)]
	fake.getHostnameArgsForCall = append(fake.getHostnameArgsForCall, struct {
	}{})
	stub := fake.GetHostnameStub
	fakeReturns := fake.getHostnameReturns
	fake.recordInvocation("GetHostname", []interface{}{})
	fake.getHostnameMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeHostConfig) GetHostnameCallCount() int {
	fake.getHostnameMutex.RLock()
	defer fake.getHostnameMutex.RUnlock()
	return len(fake.getHostnameArgsForCall)
}

func (fake *FakeHostConfig) GetHostnameCalls(stub func() string) {
	fake.getHostnameMutex.Lock()
	defer fake.getHostnameMutex.Unlock()
	fake.GetHostnameStub = stub
}

func (fake *FakeHostConfig) GetHostnameReturns(result1 string) {
	fake.getHostnameMutex.Lock()
	defer fake.getHostnameMutex.Unlock()
	fake.GetHostnameStub = nil
	fake.getHostnameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeHostConfig) GetHostnameReturnsOnCall(i int, result1 string) {
	fake.getHostnameMutex.Lock()
	defer fake.getHostnameMutex.Unlock()
	fake.GetHostnameStub = nil
	if fake.getHostnameReturnsOnCall == nil {
		fake.getHostnameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.getHostnameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeHostConfig) GetIp() string {
	fake.getIpMutex.Lock()
	ret, specificReturn := fake.getIpReturnsOnCall[len(fake.getIpArgsForCall)]
	fake.getIpArgsForCall = append(fake.getIpArgsForCall, struct {
	}{})
	stub := fake.GetIpStub
	fakeReturns := fake.getIpReturns
	fake.recordInvocation("GetIp", []interface{}{})
	fake.getIpMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeHostConfig) GetIpCallCount() int {
	fake.getIpMutex.RLock()
	defer fake.getIpMutex.RUnlock()
	return len(fake.getIpArgsForCall)
}

func (fake *FakeHostConfig) GetIpCalls(stub func() string) {
	fake.getIpMutex.Lock()
	defer fake.getIpMutex.Unlock()
	fake.GetIpStub = stub
}

func (fake *FakeHostConfig) GetIpReturns(result1 string) {
	fake.getIpMutex.Lock()
	defer fake.getIpMutex.Unlock()
	fake.GetIpStub = nil
	fake.getIpReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeHostConfig) GetIpReturnsOnCall(i int, result1 string) {
	fake.getIpMutex.Lock()
	defer fake.getIpMutex.Unlock()
	fake.GetIpStub = nil
	if fake.getIpReturnsOnCall == nil {
		fake.getIpReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.getIpReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeHostConfig) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getAuthMutex.RLock()
	defer fake.getAuthMutex.RUnlock()
	fake.getDomainNameMutex.RLock()
	defer fake.getDomainNameMutex.RUnlock()
	fake.getHostnameMutex.RLock()
	defer fake.getHostnameMutex.RUnlock()
	fake.getIpMutex.RLock()
	defer fake.getIpMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeHostConfig) recordInvocation(key string, args []interface{}) {
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

var _ config.HostConfig = new(FakeHostConfig)