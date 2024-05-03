// Code generated by counterfeiter. DO NOT EDIT.
package configfakes

import (
	"sync"

	"github.com/greenplum-db/gpdb/gp/internal/config"
)

type FakeSegmentHostsConfig struct {
	GetAuthenticationStub        func() config.AuthenticationConfig
	getAuthenticationMutex       sync.RWMutex
	getAuthenticationArgsForCall []struct {
	}
	getAuthenticationReturns struct {
		result1 config.AuthenticationConfig
	}
	getAuthenticationReturnsOnCall map[int]struct {
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
	GetHostnamePrefixStub        func() string
	getHostnamePrefixMutex       sync.RWMutex
	getHostnamePrefixArgsForCall []struct {
	}
	getHostnamePrefixReturns struct {
		result1 string
	}
	getHostnamePrefixReturnsOnCall map[int]struct {
		result1 string
	}
	GetNetworkStub        func() config.SegmentHostsNetworkConfig
	getNetworkMutex       sync.RWMutex
	getNetworkArgsForCall []struct {
	}
	getNetworkReturns struct {
		result1 config.SegmentHostsNetworkConfig
	}
	getNetworkReturnsOnCall map[int]struct {
		result1 config.SegmentHostsNetworkConfig
	}
	GetSegmentHostsCountStub        func() int
	getSegmentHostsCountMutex       sync.RWMutex
	getSegmentHostsCountArgsForCall []struct {
	}
	getSegmentHostsCountReturns struct {
		result1 int
	}
	getSegmentHostsCountReturnsOnCall map[int]struct {
		result1 int
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSegmentHostsConfig) GetAuthentication() config.AuthenticationConfig {
	fake.getAuthenticationMutex.Lock()
	ret, specificReturn := fake.getAuthenticationReturnsOnCall[len(fake.getAuthenticationArgsForCall)]
	fake.getAuthenticationArgsForCall = append(fake.getAuthenticationArgsForCall, struct {
	}{})
	stub := fake.GetAuthenticationStub
	fakeReturns := fake.getAuthenticationReturns
	fake.recordInvocation("GetAuthentication", []interface{}{})
	fake.getAuthenticationMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeSegmentHostsConfig) GetAuthenticationCallCount() int {
	fake.getAuthenticationMutex.RLock()
	defer fake.getAuthenticationMutex.RUnlock()
	return len(fake.getAuthenticationArgsForCall)
}

func (fake *FakeSegmentHostsConfig) GetAuthenticationCalls(stub func() config.AuthenticationConfig) {
	fake.getAuthenticationMutex.Lock()
	defer fake.getAuthenticationMutex.Unlock()
	fake.GetAuthenticationStub = stub
}

func (fake *FakeSegmentHostsConfig) GetAuthenticationReturns(result1 config.AuthenticationConfig) {
	fake.getAuthenticationMutex.Lock()
	defer fake.getAuthenticationMutex.Unlock()
	fake.GetAuthenticationStub = nil
	fake.getAuthenticationReturns = struct {
		result1 config.AuthenticationConfig
	}{result1}
}

func (fake *FakeSegmentHostsConfig) GetAuthenticationReturnsOnCall(i int, result1 config.AuthenticationConfig) {
	fake.getAuthenticationMutex.Lock()
	defer fake.getAuthenticationMutex.Unlock()
	fake.GetAuthenticationStub = nil
	if fake.getAuthenticationReturnsOnCall == nil {
		fake.getAuthenticationReturnsOnCall = make(map[int]struct {
			result1 config.AuthenticationConfig
		})
	}
	fake.getAuthenticationReturnsOnCall[i] = struct {
		result1 config.AuthenticationConfig
	}{result1}
}

func (fake *FakeSegmentHostsConfig) GetDomainName() string {
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

func (fake *FakeSegmentHostsConfig) GetDomainNameCallCount() int {
	fake.getDomainNameMutex.RLock()
	defer fake.getDomainNameMutex.RUnlock()
	return len(fake.getDomainNameArgsForCall)
}

func (fake *FakeSegmentHostsConfig) GetDomainNameCalls(stub func() string) {
	fake.getDomainNameMutex.Lock()
	defer fake.getDomainNameMutex.Unlock()
	fake.GetDomainNameStub = stub
}

func (fake *FakeSegmentHostsConfig) GetDomainNameReturns(result1 string) {
	fake.getDomainNameMutex.Lock()
	defer fake.getDomainNameMutex.Unlock()
	fake.GetDomainNameStub = nil
	fake.getDomainNameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeSegmentHostsConfig) GetDomainNameReturnsOnCall(i int, result1 string) {
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

func (fake *FakeSegmentHostsConfig) GetHostnamePrefix() string {
	fake.getHostnamePrefixMutex.Lock()
	ret, specificReturn := fake.getHostnamePrefixReturnsOnCall[len(fake.getHostnamePrefixArgsForCall)]
	fake.getHostnamePrefixArgsForCall = append(fake.getHostnamePrefixArgsForCall, struct {
	}{})
	stub := fake.GetHostnamePrefixStub
	fakeReturns := fake.getHostnamePrefixReturns
	fake.recordInvocation("GetHostnamePrefix", []interface{}{})
	fake.getHostnamePrefixMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeSegmentHostsConfig) GetHostnamePrefixCallCount() int {
	fake.getHostnamePrefixMutex.RLock()
	defer fake.getHostnamePrefixMutex.RUnlock()
	return len(fake.getHostnamePrefixArgsForCall)
}

func (fake *FakeSegmentHostsConfig) GetHostnamePrefixCalls(stub func() string) {
	fake.getHostnamePrefixMutex.Lock()
	defer fake.getHostnamePrefixMutex.Unlock()
	fake.GetHostnamePrefixStub = stub
}

func (fake *FakeSegmentHostsConfig) GetHostnamePrefixReturns(result1 string) {
	fake.getHostnamePrefixMutex.Lock()
	defer fake.getHostnamePrefixMutex.Unlock()
	fake.GetHostnamePrefixStub = nil
	fake.getHostnamePrefixReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeSegmentHostsConfig) GetHostnamePrefixReturnsOnCall(i int, result1 string) {
	fake.getHostnamePrefixMutex.Lock()
	defer fake.getHostnamePrefixMutex.Unlock()
	fake.GetHostnamePrefixStub = nil
	if fake.getHostnamePrefixReturnsOnCall == nil {
		fake.getHostnamePrefixReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.getHostnamePrefixReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeSegmentHostsConfig) GetNetwork() config.SegmentHostsNetworkConfig {
	fake.getNetworkMutex.Lock()
	ret, specificReturn := fake.getNetworkReturnsOnCall[len(fake.getNetworkArgsForCall)]
	fake.getNetworkArgsForCall = append(fake.getNetworkArgsForCall, struct {
	}{})
	stub := fake.GetNetworkStub
	fakeReturns := fake.getNetworkReturns
	fake.recordInvocation("GetNetwork", []interface{}{})
	fake.getNetworkMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeSegmentHostsConfig) GetNetworkCallCount() int {
	fake.getNetworkMutex.RLock()
	defer fake.getNetworkMutex.RUnlock()
	return len(fake.getNetworkArgsForCall)
}

func (fake *FakeSegmentHostsConfig) GetNetworkCalls(stub func() config.SegmentHostsNetworkConfig) {
	fake.getNetworkMutex.Lock()
	defer fake.getNetworkMutex.Unlock()
	fake.GetNetworkStub = stub
}

func (fake *FakeSegmentHostsConfig) GetNetworkReturns(result1 config.SegmentHostsNetworkConfig) {
	fake.getNetworkMutex.Lock()
	defer fake.getNetworkMutex.Unlock()
	fake.GetNetworkStub = nil
	fake.getNetworkReturns = struct {
		result1 config.SegmentHostsNetworkConfig
	}{result1}
}

func (fake *FakeSegmentHostsConfig) GetNetworkReturnsOnCall(i int, result1 config.SegmentHostsNetworkConfig) {
	fake.getNetworkMutex.Lock()
	defer fake.getNetworkMutex.Unlock()
	fake.GetNetworkStub = nil
	if fake.getNetworkReturnsOnCall == nil {
		fake.getNetworkReturnsOnCall = make(map[int]struct {
			result1 config.SegmentHostsNetworkConfig
		})
	}
	fake.getNetworkReturnsOnCall[i] = struct {
		result1 config.SegmentHostsNetworkConfig
	}{result1}
}

func (fake *FakeSegmentHostsConfig) GetSegmentHostsCount() int {
	fake.getSegmentHostsCountMutex.Lock()
	ret, specificReturn := fake.getSegmentHostsCountReturnsOnCall[len(fake.getSegmentHostsCountArgsForCall)]
	fake.getSegmentHostsCountArgsForCall = append(fake.getSegmentHostsCountArgsForCall, struct {
	}{})
	stub := fake.GetSegmentHostsCountStub
	fakeReturns := fake.getSegmentHostsCountReturns
	fake.recordInvocation("GetSegmentHostsCount", []interface{}{})
	fake.getSegmentHostsCountMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeSegmentHostsConfig) GetSegmentHostsCountCallCount() int {
	fake.getSegmentHostsCountMutex.RLock()
	defer fake.getSegmentHostsCountMutex.RUnlock()
	return len(fake.getSegmentHostsCountArgsForCall)
}

func (fake *FakeSegmentHostsConfig) GetSegmentHostsCountCalls(stub func() int) {
	fake.getSegmentHostsCountMutex.Lock()
	defer fake.getSegmentHostsCountMutex.Unlock()
	fake.GetSegmentHostsCountStub = stub
}

func (fake *FakeSegmentHostsConfig) GetSegmentHostsCountReturns(result1 int) {
	fake.getSegmentHostsCountMutex.Lock()
	defer fake.getSegmentHostsCountMutex.Unlock()
	fake.GetSegmentHostsCountStub = nil
	fake.getSegmentHostsCountReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeSegmentHostsConfig) GetSegmentHostsCountReturnsOnCall(i int, result1 int) {
	fake.getSegmentHostsCountMutex.Lock()
	defer fake.getSegmentHostsCountMutex.Unlock()
	fake.GetSegmentHostsCountStub = nil
	if fake.getSegmentHostsCountReturnsOnCall == nil {
		fake.getSegmentHostsCountReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.getSegmentHostsCountReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *FakeSegmentHostsConfig) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getAuthenticationMutex.RLock()
	defer fake.getAuthenticationMutex.RUnlock()
	fake.getDomainNameMutex.RLock()
	defer fake.getDomainNameMutex.RUnlock()
	fake.getHostnamePrefixMutex.RLock()
	defer fake.getHostnamePrefixMutex.RUnlock()
	fake.getNetworkMutex.RLock()
	defer fake.getNetworkMutex.RUnlock()
	fake.getSegmentHostsCountMutex.RLock()
	defer fake.getSegmentHostsCountMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSegmentHostsConfig) recordInvocation(key string, args []interface{}) {
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

var _ config.SegmentHostsConfig = new(FakeSegmentHostsConfig)