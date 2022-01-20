// Code generated by counterfeiter. DO NOT EDIT.
package repository

import (
	"sync"

	"github.com/u-shylianok/ad-service/svc-ads/model"
	"github.com/u-shylianok/ad-service/svc-ads/repository"
)

type TagMock struct {
	AttachToAdStub        func(int, int) error
	attachToAdMutex       sync.RWMutex
	attachToAdArgsForCall []struct {
		arg1 int
		arg2 int
	}
	attachToAdReturns struct {
		result1 error
	}
	attachToAdReturnsOnCall map[int]struct {
		result1 error
	}
	CreateStub        func(string) (int, error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 string
	}
	createReturns struct {
		result1 int
		result2 error
	}
	createReturnsOnCall map[int]struct {
		result1 int
		result2 error
	}
	DetachAllFromAdStub        func(int) error
	detachAllFromAdMutex       sync.RWMutex
	detachAllFromAdArgsForCall []struct {
		arg1 int
	}
	detachAllFromAdReturns struct {
		result1 error
	}
	detachAllFromAdReturnsOnCall map[int]struct {
		result1 error
	}
	DetachFromAdStub        func(int, int) error
	detachFromAdMutex       sync.RWMutex
	detachFromAdArgsForCall []struct {
		arg1 int
		arg2 int
	}
	detachFromAdReturns struct {
		result1 error
	}
	detachFromAdReturnsOnCall map[int]struct {
		result1 error
	}
	GetByNameStub        func(string) (model.Tag, error)
	getByNameMutex       sync.RWMutex
	getByNameArgsForCall []struct {
		arg1 string
	}
	getByNameReturns struct {
		result1 model.Tag
		result2 error
	}
	getByNameReturnsOnCall map[int]struct {
		result1 model.Tag
		result2 error
	}
	GetIDOrCreateIfNotExistsStub        func(string) (int, error)
	getIDOrCreateIfNotExistsMutex       sync.RWMutex
	getIDOrCreateIfNotExistsArgsForCall []struct {
		arg1 string
	}
	getIDOrCreateIfNotExistsReturns struct {
		result1 int
		result2 error
	}
	getIDOrCreateIfNotExistsReturnsOnCall map[int]struct {
		result1 int
		result2 error
	}
	ListNamesStub        func() ([]string, error)
	listNamesMutex       sync.RWMutex
	listNamesArgsForCall []struct {
	}
	listNamesReturns struct {
		result1 []string
		result2 error
	}
	listNamesReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	ListNamesByAdStub        func(int) ([]string, error)
	listNamesByAdMutex       sync.RWMutex
	listNamesByAdArgsForCall []struct {
		arg1 int
	}
	listNamesByAdReturns struct {
		result1 []string
		result2 error
	}
	listNamesByAdReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *TagMock) AttachToAd(arg1 int, arg2 int) error {
	fake.attachToAdMutex.Lock()
	ret, specificReturn := fake.attachToAdReturnsOnCall[len(fake.attachToAdArgsForCall)]
	fake.attachToAdArgsForCall = append(fake.attachToAdArgsForCall, struct {
		arg1 int
		arg2 int
	}{arg1, arg2})
	stub := fake.AttachToAdStub
	fakeReturns := fake.attachToAdReturns
	fake.recordInvocation("AttachToAd", []interface{}{arg1, arg2})
	fake.attachToAdMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *TagMock) AttachToAdCallCount() int {
	fake.attachToAdMutex.RLock()
	defer fake.attachToAdMutex.RUnlock()
	return len(fake.attachToAdArgsForCall)
}

func (fake *TagMock) AttachToAdCalls(stub func(int, int) error) {
	fake.attachToAdMutex.Lock()
	defer fake.attachToAdMutex.Unlock()
	fake.AttachToAdStub = stub
}

func (fake *TagMock) AttachToAdArgsForCall(i int) (int, int) {
	fake.attachToAdMutex.RLock()
	defer fake.attachToAdMutex.RUnlock()
	argsForCall := fake.attachToAdArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *TagMock) AttachToAdReturns(result1 error) {
	fake.attachToAdMutex.Lock()
	defer fake.attachToAdMutex.Unlock()
	fake.AttachToAdStub = nil
	fake.attachToAdReturns = struct {
		result1 error
	}{result1}
}

func (fake *TagMock) AttachToAdReturnsOnCall(i int, result1 error) {
	fake.attachToAdMutex.Lock()
	defer fake.attachToAdMutex.Unlock()
	fake.AttachToAdStub = nil
	if fake.attachToAdReturnsOnCall == nil {
		fake.attachToAdReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.attachToAdReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *TagMock) Create(arg1 string) (int, error) {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.CreateStub
	fakeReturns := fake.createReturns
	fake.recordInvocation("Create", []interface{}{arg1})
	fake.createMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *TagMock) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *TagMock) CreateCalls(stub func(string) (int, error)) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = stub
}

func (fake *TagMock) CreateArgsForCall(i int) string {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	argsForCall := fake.createArgsForCall[i]
	return argsForCall.arg1
}

func (fake *TagMock) CreateReturns(result1 int, result2 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *TagMock) CreateReturnsOnCall(i int, result1 int, result2 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 int
			result2 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *TagMock) DetachAllFromAd(arg1 int) error {
	fake.detachAllFromAdMutex.Lock()
	ret, specificReturn := fake.detachAllFromAdReturnsOnCall[len(fake.detachAllFromAdArgsForCall)]
	fake.detachAllFromAdArgsForCall = append(fake.detachAllFromAdArgsForCall, struct {
		arg1 int
	}{arg1})
	stub := fake.DetachAllFromAdStub
	fakeReturns := fake.detachAllFromAdReturns
	fake.recordInvocation("DetachAllFromAd", []interface{}{arg1})
	fake.detachAllFromAdMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *TagMock) DetachAllFromAdCallCount() int {
	fake.detachAllFromAdMutex.RLock()
	defer fake.detachAllFromAdMutex.RUnlock()
	return len(fake.detachAllFromAdArgsForCall)
}

func (fake *TagMock) DetachAllFromAdCalls(stub func(int) error) {
	fake.detachAllFromAdMutex.Lock()
	defer fake.detachAllFromAdMutex.Unlock()
	fake.DetachAllFromAdStub = stub
}

func (fake *TagMock) DetachAllFromAdArgsForCall(i int) int {
	fake.detachAllFromAdMutex.RLock()
	defer fake.detachAllFromAdMutex.RUnlock()
	argsForCall := fake.detachAllFromAdArgsForCall[i]
	return argsForCall.arg1
}

func (fake *TagMock) DetachAllFromAdReturns(result1 error) {
	fake.detachAllFromAdMutex.Lock()
	defer fake.detachAllFromAdMutex.Unlock()
	fake.DetachAllFromAdStub = nil
	fake.detachAllFromAdReturns = struct {
		result1 error
	}{result1}
}

func (fake *TagMock) DetachAllFromAdReturnsOnCall(i int, result1 error) {
	fake.detachAllFromAdMutex.Lock()
	defer fake.detachAllFromAdMutex.Unlock()
	fake.DetachAllFromAdStub = nil
	if fake.detachAllFromAdReturnsOnCall == nil {
		fake.detachAllFromAdReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.detachAllFromAdReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *TagMock) DetachFromAd(arg1 int, arg2 int) error {
	fake.detachFromAdMutex.Lock()
	ret, specificReturn := fake.detachFromAdReturnsOnCall[len(fake.detachFromAdArgsForCall)]
	fake.detachFromAdArgsForCall = append(fake.detachFromAdArgsForCall, struct {
		arg1 int
		arg2 int
	}{arg1, arg2})
	stub := fake.DetachFromAdStub
	fakeReturns := fake.detachFromAdReturns
	fake.recordInvocation("DetachFromAd", []interface{}{arg1, arg2})
	fake.detachFromAdMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *TagMock) DetachFromAdCallCount() int {
	fake.detachFromAdMutex.RLock()
	defer fake.detachFromAdMutex.RUnlock()
	return len(fake.detachFromAdArgsForCall)
}

func (fake *TagMock) DetachFromAdCalls(stub func(int, int) error) {
	fake.detachFromAdMutex.Lock()
	defer fake.detachFromAdMutex.Unlock()
	fake.DetachFromAdStub = stub
}

func (fake *TagMock) DetachFromAdArgsForCall(i int) (int, int) {
	fake.detachFromAdMutex.RLock()
	defer fake.detachFromAdMutex.RUnlock()
	argsForCall := fake.detachFromAdArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *TagMock) DetachFromAdReturns(result1 error) {
	fake.detachFromAdMutex.Lock()
	defer fake.detachFromAdMutex.Unlock()
	fake.DetachFromAdStub = nil
	fake.detachFromAdReturns = struct {
		result1 error
	}{result1}
}

func (fake *TagMock) DetachFromAdReturnsOnCall(i int, result1 error) {
	fake.detachFromAdMutex.Lock()
	defer fake.detachFromAdMutex.Unlock()
	fake.DetachFromAdStub = nil
	if fake.detachFromAdReturnsOnCall == nil {
		fake.detachFromAdReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.detachFromAdReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *TagMock) GetByName(arg1 string) (model.Tag, error) {
	fake.getByNameMutex.Lock()
	ret, specificReturn := fake.getByNameReturnsOnCall[len(fake.getByNameArgsForCall)]
	fake.getByNameArgsForCall = append(fake.getByNameArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetByNameStub
	fakeReturns := fake.getByNameReturns
	fake.recordInvocation("GetByName", []interface{}{arg1})
	fake.getByNameMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *TagMock) GetByNameCallCount() int {
	fake.getByNameMutex.RLock()
	defer fake.getByNameMutex.RUnlock()
	return len(fake.getByNameArgsForCall)
}

func (fake *TagMock) GetByNameCalls(stub func(string) (model.Tag, error)) {
	fake.getByNameMutex.Lock()
	defer fake.getByNameMutex.Unlock()
	fake.GetByNameStub = stub
}

func (fake *TagMock) GetByNameArgsForCall(i int) string {
	fake.getByNameMutex.RLock()
	defer fake.getByNameMutex.RUnlock()
	argsForCall := fake.getByNameArgsForCall[i]
	return argsForCall.arg1
}

func (fake *TagMock) GetByNameReturns(result1 model.Tag, result2 error) {
	fake.getByNameMutex.Lock()
	defer fake.getByNameMutex.Unlock()
	fake.GetByNameStub = nil
	fake.getByNameReturns = struct {
		result1 model.Tag
		result2 error
	}{result1, result2}
}

func (fake *TagMock) GetByNameReturnsOnCall(i int, result1 model.Tag, result2 error) {
	fake.getByNameMutex.Lock()
	defer fake.getByNameMutex.Unlock()
	fake.GetByNameStub = nil
	if fake.getByNameReturnsOnCall == nil {
		fake.getByNameReturnsOnCall = make(map[int]struct {
			result1 model.Tag
			result2 error
		})
	}
	fake.getByNameReturnsOnCall[i] = struct {
		result1 model.Tag
		result2 error
	}{result1, result2}
}

func (fake *TagMock) GetIDOrCreateIfNotExists(arg1 string) (int, error) {
	fake.getIDOrCreateIfNotExistsMutex.Lock()
	ret, specificReturn := fake.getIDOrCreateIfNotExistsReturnsOnCall[len(fake.getIDOrCreateIfNotExistsArgsForCall)]
	fake.getIDOrCreateIfNotExistsArgsForCall = append(fake.getIDOrCreateIfNotExistsArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetIDOrCreateIfNotExistsStub
	fakeReturns := fake.getIDOrCreateIfNotExistsReturns
	fake.recordInvocation("GetIDOrCreateIfNotExists", []interface{}{arg1})
	fake.getIDOrCreateIfNotExistsMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *TagMock) GetIDOrCreateIfNotExistsCallCount() int {
	fake.getIDOrCreateIfNotExistsMutex.RLock()
	defer fake.getIDOrCreateIfNotExistsMutex.RUnlock()
	return len(fake.getIDOrCreateIfNotExistsArgsForCall)
}

func (fake *TagMock) GetIDOrCreateIfNotExistsCalls(stub func(string) (int, error)) {
	fake.getIDOrCreateIfNotExistsMutex.Lock()
	defer fake.getIDOrCreateIfNotExistsMutex.Unlock()
	fake.GetIDOrCreateIfNotExistsStub = stub
}

func (fake *TagMock) GetIDOrCreateIfNotExistsArgsForCall(i int) string {
	fake.getIDOrCreateIfNotExistsMutex.RLock()
	defer fake.getIDOrCreateIfNotExistsMutex.RUnlock()
	argsForCall := fake.getIDOrCreateIfNotExistsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *TagMock) GetIDOrCreateIfNotExistsReturns(result1 int, result2 error) {
	fake.getIDOrCreateIfNotExistsMutex.Lock()
	defer fake.getIDOrCreateIfNotExistsMutex.Unlock()
	fake.GetIDOrCreateIfNotExistsStub = nil
	fake.getIDOrCreateIfNotExistsReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *TagMock) GetIDOrCreateIfNotExistsReturnsOnCall(i int, result1 int, result2 error) {
	fake.getIDOrCreateIfNotExistsMutex.Lock()
	defer fake.getIDOrCreateIfNotExistsMutex.Unlock()
	fake.GetIDOrCreateIfNotExistsStub = nil
	if fake.getIDOrCreateIfNotExistsReturnsOnCall == nil {
		fake.getIDOrCreateIfNotExistsReturnsOnCall = make(map[int]struct {
			result1 int
			result2 error
		})
	}
	fake.getIDOrCreateIfNotExistsReturnsOnCall[i] = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *TagMock) ListNames() ([]string, error) {
	fake.listNamesMutex.Lock()
	ret, specificReturn := fake.listNamesReturnsOnCall[len(fake.listNamesArgsForCall)]
	fake.listNamesArgsForCall = append(fake.listNamesArgsForCall, struct {
	}{})
	stub := fake.ListNamesStub
	fakeReturns := fake.listNamesReturns
	fake.recordInvocation("ListNames", []interface{}{})
	fake.listNamesMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *TagMock) ListNamesCallCount() int {
	fake.listNamesMutex.RLock()
	defer fake.listNamesMutex.RUnlock()
	return len(fake.listNamesArgsForCall)
}

func (fake *TagMock) ListNamesCalls(stub func() ([]string, error)) {
	fake.listNamesMutex.Lock()
	defer fake.listNamesMutex.Unlock()
	fake.ListNamesStub = stub
}

func (fake *TagMock) ListNamesReturns(result1 []string, result2 error) {
	fake.listNamesMutex.Lock()
	defer fake.listNamesMutex.Unlock()
	fake.ListNamesStub = nil
	fake.listNamesReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *TagMock) ListNamesReturnsOnCall(i int, result1 []string, result2 error) {
	fake.listNamesMutex.Lock()
	defer fake.listNamesMutex.Unlock()
	fake.ListNamesStub = nil
	if fake.listNamesReturnsOnCall == nil {
		fake.listNamesReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.listNamesReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *TagMock) ListNamesByAd(arg1 int) ([]string, error) {
	fake.listNamesByAdMutex.Lock()
	ret, specificReturn := fake.listNamesByAdReturnsOnCall[len(fake.listNamesByAdArgsForCall)]
	fake.listNamesByAdArgsForCall = append(fake.listNamesByAdArgsForCall, struct {
		arg1 int
	}{arg1})
	stub := fake.ListNamesByAdStub
	fakeReturns := fake.listNamesByAdReturns
	fake.recordInvocation("ListNamesByAd", []interface{}{arg1})
	fake.listNamesByAdMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *TagMock) ListNamesByAdCallCount() int {
	fake.listNamesByAdMutex.RLock()
	defer fake.listNamesByAdMutex.RUnlock()
	return len(fake.listNamesByAdArgsForCall)
}

func (fake *TagMock) ListNamesByAdCalls(stub func(int) ([]string, error)) {
	fake.listNamesByAdMutex.Lock()
	defer fake.listNamesByAdMutex.Unlock()
	fake.ListNamesByAdStub = stub
}

func (fake *TagMock) ListNamesByAdArgsForCall(i int) int {
	fake.listNamesByAdMutex.RLock()
	defer fake.listNamesByAdMutex.RUnlock()
	argsForCall := fake.listNamesByAdArgsForCall[i]
	return argsForCall.arg1
}

func (fake *TagMock) ListNamesByAdReturns(result1 []string, result2 error) {
	fake.listNamesByAdMutex.Lock()
	defer fake.listNamesByAdMutex.Unlock()
	fake.ListNamesByAdStub = nil
	fake.listNamesByAdReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *TagMock) ListNamesByAdReturnsOnCall(i int, result1 []string, result2 error) {
	fake.listNamesByAdMutex.Lock()
	defer fake.listNamesByAdMutex.Unlock()
	fake.ListNamesByAdStub = nil
	if fake.listNamesByAdReturnsOnCall == nil {
		fake.listNamesByAdReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.listNamesByAdReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *TagMock) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.attachToAdMutex.RLock()
	defer fake.attachToAdMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.detachAllFromAdMutex.RLock()
	defer fake.detachAllFromAdMutex.RUnlock()
	fake.detachFromAdMutex.RLock()
	defer fake.detachFromAdMutex.RUnlock()
	fake.getByNameMutex.RLock()
	defer fake.getByNameMutex.RUnlock()
	fake.getIDOrCreateIfNotExistsMutex.RLock()
	defer fake.getIDOrCreateIfNotExistsMutex.RUnlock()
	fake.listNamesMutex.RLock()
	defer fake.listNamesMutex.RUnlock()
	fake.listNamesByAdMutex.RLock()
	defer fake.listNamesByAdMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *TagMock) recordInvocation(key string, args []interface{}) {
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

var _ repository.Tag = new(TagMock)