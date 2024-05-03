// Code generated by counterfeiter. DO NOT EDIT.
package configfakes

import (
	"sync"

	"github.com/greenplum-db/gpdb/gp/internal/config"
	"github.com/greenplum-db/gpdb/gp/internal/enums"
)

type FakeDatabaseConfig struct {
	GetAdminStub        func() config.UserConfig
	getAdminMutex       sync.RWMutex
	getAdminArgsForCall []struct {
	}
	getAdminReturns struct {
		result1 config.UserConfig
	}
	getAdminReturnsOnCall map[int]struct {
		result1 config.UserConfig
	}
	GetDeploymentTypeStub        func() enums.DeploymentType
	getDeploymentTypeMutex       sync.RWMutex
	getDeploymentTypeArgsForCall []struct {
	}
	getDeploymentTypeReturns struct {
		result1 enums.DeploymentType
	}
	getDeploymentTypeReturnsOnCall map[int]struct {
		result1 enums.DeploymentType
	}
	GetSegmentsPerSegmentHostStub        func() int
	getSegmentsPerSegmentHostMutex       sync.RWMutex
	getSegmentsPerSegmentHostArgsForCall []struct {
	}
	getSegmentsPerSegmentHostReturns struct {
		result1 int
	}
	getSegmentsPerSegmentHostReturnsOnCall map[int]struct {
		result1 int
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDatabaseConfig) GetAdmin() config.UserConfig {
	fake.getAdminMutex.Lock()
	ret, specificReturn := fake.getAdminReturnsOnCall[len(fake.getAdminArgsForCall)]
	fake.getAdminArgsForCall = append(fake.getAdminArgsForCall, struct {
	}{})
	stub := fake.GetAdminStub
	fakeReturns := fake.getAdminReturns
	fake.recordInvocation("GetAdmin", []interface{}{})
	fake.getAdminMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeDatabaseConfig) GetAdminCallCount() int {
	fake.getAdminMutex.RLock()
	defer fake.getAdminMutex.RUnlock()
	return len(fake.getAdminArgsForCall)
}

func (fake *FakeDatabaseConfig) GetAdminCalls(stub func() config.UserConfig) {
	fake.getAdminMutex.Lock()
	defer fake.getAdminMutex.Unlock()
	fake.GetAdminStub = stub
}

func (fake *FakeDatabaseConfig) GetAdminReturns(result1 config.UserConfig) {
	fake.getAdminMutex.Lock()
	defer fake.getAdminMutex.Unlock()
	fake.GetAdminStub = nil
	fake.getAdminReturns = struct {
		result1 config.UserConfig
	}{result1}
}

func (fake *FakeDatabaseConfig) GetAdminReturnsOnCall(i int, result1 config.UserConfig) {
	fake.getAdminMutex.Lock()
	defer fake.getAdminMutex.Unlock()
	fake.GetAdminStub = nil
	if fake.getAdminReturnsOnCall == nil {
		fake.getAdminReturnsOnCall = make(map[int]struct {
			result1 config.UserConfig
		})
	}
	fake.getAdminReturnsOnCall[i] = struct {
		result1 config.UserConfig
	}{result1}
}

func (fake *FakeDatabaseConfig) GetDeploymentType() enums.DeploymentType {
	fake.getDeploymentTypeMutex.Lock()
	ret, specificReturn := fake.getDeploymentTypeReturnsOnCall[len(fake.getDeploymentTypeArgsForCall)]
	fake.getDeploymentTypeArgsForCall = append(fake.getDeploymentTypeArgsForCall, struct {
	}{})
	stub := fake.GetDeploymentTypeStub
	fakeReturns := fake.getDeploymentTypeReturns
	fake.recordInvocation("GetDeploymentType", []interface{}{})
	fake.getDeploymentTypeMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeDatabaseConfig) GetDeploymentTypeCallCount() int {
	fake.getDeploymentTypeMutex.RLock()
	defer fake.getDeploymentTypeMutex.RUnlock()
	return len(fake.getDeploymentTypeArgsForCall)
}

func (fake *FakeDatabaseConfig) GetDeploymentTypeCalls(stub func() enums.DeploymentType) {
	fake.getDeploymentTypeMutex.Lock()
	defer fake.getDeploymentTypeMutex.Unlock()
	fake.GetDeploymentTypeStub = stub
}

func (fake *FakeDatabaseConfig) GetDeploymentTypeReturns(result1 enums.DeploymentType) {
	fake.getDeploymentTypeMutex.Lock()
	defer fake.getDeploymentTypeMutex.Unlock()
	fake.GetDeploymentTypeStub = nil
	fake.getDeploymentTypeReturns = struct {
		result1 enums.DeploymentType
	}{result1}
}

func (fake *FakeDatabaseConfig) GetDeploymentTypeReturnsOnCall(i int, result1 enums.DeploymentType) {
	fake.getDeploymentTypeMutex.Lock()
	defer fake.getDeploymentTypeMutex.Unlock()
	fake.GetDeploymentTypeStub = nil
	if fake.getDeploymentTypeReturnsOnCall == nil {
		fake.getDeploymentTypeReturnsOnCall = make(map[int]struct {
			result1 enums.DeploymentType
		})
	}
	fake.getDeploymentTypeReturnsOnCall[i] = struct {
		result1 enums.DeploymentType
	}{result1}
}

func (fake *FakeDatabaseConfig) GetSegmentsPerSegmentHost() int {
	fake.getSegmentsPerSegmentHostMutex.Lock()
	ret, specificReturn := fake.getSegmentsPerSegmentHostReturnsOnCall[len(fake.getSegmentsPerSegmentHostArgsForCall)]
	fake.getSegmentsPerSegmentHostArgsForCall = append(fake.getSegmentsPerSegmentHostArgsForCall, struct {
	}{})
	stub := fake.GetSegmentsPerSegmentHostStub
	fakeReturns := fake.getSegmentsPerSegmentHostReturns
	fake.recordInvocation("GetSegmentsPerSegmentHost", []interface{}{})
	fake.getSegmentsPerSegmentHostMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeDatabaseConfig) GetSegmentsPerSegmentHostCallCount() int {
	fake.getSegmentsPerSegmentHostMutex.RLock()
	defer fake.getSegmentsPerSegmentHostMutex.RUnlock()
	return len(fake.getSegmentsPerSegmentHostArgsForCall)
}

func (fake *FakeDatabaseConfig) GetSegmentsPerSegmentHostCalls(stub func() int) {
	fake.getSegmentsPerSegmentHostMutex.Lock()
	defer fake.getSegmentsPerSegmentHostMutex.Unlock()
	fake.GetSegmentsPerSegmentHostStub = stub
}

func (fake *FakeDatabaseConfig) GetSegmentsPerSegmentHostReturns(result1 int) {
	fake.getSegmentsPerSegmentHostMutex.Lock()
	defer fake.getSegmentsPerSegmentHostMutex.Unlock()
	fake.GetSegmentsPerSegmentHostStub = nil
	fake.getSegmentsPerSegmentHostReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeDatabaseConfig) GetSegmentsPerSegmentHostReturnsOnCall(i int, result1 int) {
	fake.getSegmentsPerSegmentHostMutex.Lock()
	defer fake.getSegmentsPerSegmentHostMutex.Unlock()
	fake.GetSegmentsPerSegmentHostStub = nil
	if fake.getSegmentsPerSegmentHostReturnsOnCall == nil {
		fake.getSegmentsPerSegmentHostReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.getSegmentsPerSegmentHostReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *FakeDatabaseConfig) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getAdminMutex.RLock()
	defer fake.getAdminMutex.RUnlock()
	fake.getDeploymentTypeMutex.RLock()
	defer fake.getDeploymentTypeMutex.RUnlock()
	fake.getSegmentsPerSegmentHostMutex.RLock()
	defer fake.getSegmentsPerSegmentHostMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDatabaseConfig) recordInvocation(key string, args []interface{}) {
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

var _ config.DatabaseConfig = new(FakeDatabaseConfig)
