// Code generated by counterfeiter. DO NOT EDIT.
package filegroupfakes

import (
	"sync"

	"github.com/pivotal-cf/go-pivnet/v2"
	"github.com/pivotal-cf/pivnet-cli/commands/filegroup"
)

type FakePivnetClient struct {
	FileGroupsStub        func(productSlug string) ([]pivnet.FileGroup, error)
	fileGroupsMutex       sync.RWMutex
	fileGroupsArgsForCall []struct {
		productSlug string
	}
	fileGroupsReturns struct {
		result1 []pivnet.FileGroup
		result2 error
	}
	fileGroupsReturnsOnCall map[int]struct {
		result1 []pivnet.FileGroup
		result2 error
	}
	FileGroupsForReleaseStub        func(productSlug string, releaseID int) ([]pivnet.FileGroup, error)
	fileGroupsForReleaseMutex       sync.RWMutex
	fileGroupsForReleaseArgsForCall []struct {
		productSlug string
		releaseID   int
	}
	fileGroupsForReleaseReturns struct {
		result1 []pivnet.FileGroup
		result2 error
	}
	fileGroupsForReleaseReturnsOnCall map[int]struct {
		result1 []pivnet.FileGroup
		result2 error
	}
	ReleaseForVersionStub        func(productSlug string, releaseVersion string) (pivnet.Release, error)
	releaseForVersionMutex       sync.RWMutex
	releaseForVersionArgsForCall []struct {
		productSlug    string
		releaseVersion string
	}
	releaseForVersionReturns struct {
		result1 pivnet.Release
		result2 error
	}
	releaseForVersionReturnsOnCall map[int]struct {
		result1 pivnet.Release
		result2 error
	}
	FileGroupStub        func(productSlug string, fileGroupID int) (pivnet.FileGroup, error)
	fileGroupMutex       sync.RWMutex
	fileGroupArgsForCall []struct {
		productSlug string
		fileGroupID int
	}
	fileGroupReturns struct {
		result1 pivnet.FileGroup
		result2 error
	}
	fileGroupReturnsOnCall map[int]struct {
		result1 pivnet.FileGroup
		result2 error
	}
	CreateFileGroupStub        func(productSlug string, name string) (pivnet.FileGroup, error)
	createFileGroupMutex       sync.RWMutex
	createFileGroupArgsForCall []struct {
		productSlug string
		name        string
	}
	createFileGroupReturns struct {
		result1 pivnet.FileGroup
		result2 error
	}
	createFileGroupReturnsOnCall map[int]struct {
		result1 pivnet.FileGroup
		result2 error
	}
	UpdateFileGroupStub        func(productSlug string, fileGroup pivnet.FileGroup) (pivnet.FileGroup, error)
	updateFileGroupMutex       sync.RWMutex
	updateFileGroupArgsForCall []struct {
		productSlug string
		fileGroup   pivnet.FileGroup
	}
	updateFileGroupReturns struct {
		result1 pivnet.FileGroup
		result2 error
	}
	updateFileGroupReturnsOnCall map[int]struct {
		result1 pivnet.FileGroup
		result2 error
	}
	DeleteFileGroupStub        func(productSlug string, fileGroupID int) (pivnet.FileGroup, error)
	deleteFileGroupMutex       sync.RWMutex
	deleteFileGroupArgsForCall []struct {
		productSlug string
		fileGroupID int
	}
	deleteFileGroupReturns struct {
		result1 pivnet.FileGroup
		result2 error
	}
	deleteFileGroupReturnsOnCall map[int]struct {
		result1 pivnet.FileGroup
		result2 error
	}
	AddFileGroupToReleaseStub        func(productSlug string, fileGroupID int, releaseID int) error
	addFileGroupToReleaseMutex       sync.RWMutex
	addFileGroupToReleaseArgsForCall []struct {
		productSlug string
		fileGroupID int
		releaseID   int
	}
	addFileGroupToReleaseReturns struct {
		result1 error
	}
	addFileGroupToReleaseReturnsOnCall map[int]struct {
		result1 error
	}
	RemoveFileGroupFromReleaseStub        func(productSlug string, fileGroupID int, releaseID int) error
	removeFileGroupFromReleaseMutex       sync.RWMutex
	removeFileGroupFromReleaseArgsForCall []struct {
		productSlug string
		fileGroupID int
		releaseID   int
	}
	removeFileGroupFromReleaseReturns struct {
		result1 error
	}
	removeFileGroupFromReleaseReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePivnetClient) FileGroups(productSlug string) ([]pivnet.FileGroup, error) {
	fake.fileGroupsMutex.Lock()
	ret, specificReturn := fake.fileGroupsReturnsOnCall[len(fake.fileGroupsArgsForCall)]
	fake.fileGroupsArgsForCall = append(fake.fileGroupsArgsForCall, struct {
		productSlug string
	}{productSlug})
	fake.recordInvocation("FileGroups", []interface{}{productSlug})
	fake.fileGroupsMutex.Unlock()
	if fake.FileGroupsStub != nil {
		return fake.FileGroupsStub(productSlug)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.fileGroupsReturns.result1, fake.fileGroupsReturns.result2
}

func (fake *FakePivnetClient) FileGroupsCallCount() int {
	fake.fileGroupsMutex.RLock()
	defer fake.fileGroupsMutex.RUnlock()
	return len(fake.fileGroupsArgsForCall)
}

func (fake *FakePivnetClient) FileGroupsArgsForCall(i int) string {
	fake.fileGroupsMutex.RLock()
	defer fake.fileGroupsMutex.RUnlock()
	return fake.fileGroupsArgsForCall[i].productSlug
}

func (fake *FakePivnetClient) FileGroupsReturns(result1 []pivnet.FileGroup, result2 error) {
	fake.FileGroupsStub = nil
	fake.fileGroupsReturns = struct {
		result1 []pivnet.FileGroup
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) FileGroupsReturnsOnCall(i int, result1 []pivnet.FileGroup, result2 error) {
	fake.FileGroupsStub = nil
	if fake.fileGroupsReturnsOnCall == nil {
		fake.fileGroupsReturnsOnCall = make(map[int]struct {
			result1 []pivnet.FileGroup
			result2 error
		})
	}
	fake.fileGroupsReturnsOnCall[i] = struct {
		result1 []pivnet.FileGroup
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) FileGroupsForRelease(productSlug string, releaseID int) ([]pivnet.FileGroup, error) {
	fake.fileGroupsForReleaseMutex.Lock()
	ret, specificReturn := fake.fileGroupsForReleaseReturnsOnCall[len(fake.fileGroupsForReleaseArgsForCall)]
	fake.fileGroupsForReleaseArgsForCall = append(fake.fileGroupsForReleaseArgsForCall, struct {
		productSlug string
		releaseID   int
	}{productSlug, releaseID})
	fake.recordInvocation("FileGroupsForRelease", []interface{}{productSlug, releaseID})
	fake.fileGroupsForReleaseMutex.Unlock()
	if fake.FileGroupsForReleaseStub != nil {
		return fake.FileGroupsForReleaseStub(productSlug, releaseID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.fileGroupsForReleaseReturns.result1, fake.fileGroupsForReleaseReturns.result2
}

func (fake *FakePivnetClient) FileGroupsForReleaseCallCount() int {
	fake.fileGroupsForReleaseMutex.RLock()
	defer fake.fileGroupsForReleaseMutex.RUnlock()
	return len(fake.fileGroupsForReleaseArgsForCall)
}

func (fake *FakePivnetClient) FileGroupsForReleaseArgsForCall(i int) (string, int) {
	fake.fileGroupsForReleaseMutex.RLock()
	defer fake.fileGroupsForReleaseMutex.RUnlock()
	return fake.fileGroupsForReleaseArgsForCall[i].productSlug, fake.fileGroupsForReleaseArgsForCall[i].releaseID
}

func (fake *FakePivnetClient) FileGroupsForReleaseReturns(result1 []pivnet.FileGroup, result2 error) {
	fake.FileGroupsForReleaseStub = nil
	fake.fileGroupsForReleaseReturns = struct {
		result1 []pivnet.FileGroup
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) FileGroupsForReleaseReturnsOnCall(i int, result1 []pivnet.FileGroup, result2 error) {
	fake.FileGroupsForReleaseStub = nil
	if fake.fileGroupsForReleaseReturnsOnCall == nil {
		fake.fileGroupsForReleaseReturnsOnCall = make(map[int]struct {
			result1 []pivnet.FileGroup
			result2 error
		})
	}
	fake.fileGroupsForReleaseReturnsOnCall[i] = struct {
		result1 []pivnet.FileGroup
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) ReleaseForVersion(productSlug string, releaseVersion string) (pivnet.Release, error) {
	fake.releaseForVersionMutex.Lock()
	ret, specificReturn := fake.releaseForVersionReturnsOnCall[len(fake.releaseForVersionArgsForCall)]
	fake.releaseForVersionArgsForCall = append(fake.releaseForVersionArgsForCall, struct {
		productSlug    string
		releaseVersion string
	}{productSlug, releaseVersion})
	fake.recordInvocation("ReleaseForVersion", []interface{}{productSlug, releaseVersion})
	fake.releaseForVersionMutex.Unlock()
	if fake.ReleaseForVersionStub != nil {
		return fake.ReleaseForVersionStub(productSlug, releaseVersion)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.releaseForVersionReturns.result1, fake.releaseForVersionReturns.result2
}

func (fake *FakePivnetClient) ReleaseForVersionCallCount() int {
	fake.releaseForVersionMutex.RLock()
	defer fake.releaseForVersionMutex.RUnlock()
	return len(fake.releaseForVersionArgsForCall)
}

func (fake *FakePivnetClient) ReleaseForVersionArgsForCall(i int) (string, string) {
	fake.releaseForVersionMutex.RLock()
	defer fake.releaseForVersionMutex.RUnlock()
	return fake.releaseForVersionArgsForCall[i].productSlug, fake.releaseForVersionArgsForCall[i].releaseVersion
}

func (fake *FakePivnetClient) ReleaseForVersionReturns(result1 pivnet.Release, result2 error) {
	fake.ReleaseForVersionStub = nil
	fake.releaseForVersionReturns = struct {
		result1 pivnet.Release
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) ReleaseForVersionReturnsOnCall(i int, result1 pivnet.Release, result2 error) {
	fake.ReleaseForVersionStub = nil
	if fake.releaseForVersionReturnsOnCall == nil {
		fake.releaseForVersionReturnsOnCall = make(map[int]struct {
			result1 pivnet.Release
			result2 error
		})
	}
	fake.releaseForVersionReturnsOnCall[i] = struct {
		result1 pivnet.Release
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) FileGroup(productSlug string, fileGroupID int) (pivnet.FileGroup, error) {
	fake.fileGroupMutex.Lock()
	ret, specificReturn := fake.fileGroupReturnsOnCall[len(fake.fileGroupArgsForCall)]
	fake.fileGroupArgsForCall = append(fake.fileGroupArgsForCall, struct {
		productSlug string
		fileGroupID int
	}{productSlug, fileGroupID})
	fake.recordInvocation("FileGroup", []interface{}{productSlug, fileGroupID})
	fake.fileGroupMutex.Unlock()
	if fake.FileGroupStub != nil {
		return fake.FileGroupStub(productSlug, fileGroupID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.fileGroupReturns.result1, fake.fileGroupReturns.result2
}

func (fake *FakePivnetClient) FileGroupCallCount() int {
	fake.fileGroupMutex.RLock()
	defer fake.fileGroupMutex.RUnlock()
	return len(fake.fileGroupArgsForCall)
}

func (fake *FakePivnetClient) FileGroupArgsForCall(i int) (string, int) {
	fake.fileGroupMutex.RLock()
	defer fake.fileGroupMutex.RUnlock()
	return fake.fileGroupArgsForCall[i].productSlug, fake.fileGroupArgsForCall[i].fileGroupID
}

func (fake *FakePivnetClient) FileGroupReturns(result1 pivnet.FileGroup, result2 error) {
	fake.FileGroupStub = nil
	fake.fileGroupReturns = struct {
		result1 pivnet.FileGroup
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) FileGroupReturnsOnCall(i int, result1 pivnet.FileGroup, result2 error) {
	fake.FileGroupStub = nil
	if fake.fileGroupReturnsOnCall == nil {
		fake.fileGroupReturnsOnCall = make(map[int]struct {
			result1 pivnet.FileGroup
			result2 error
		})
	}
	fake.fileGroupReturnsOnCall[i] = struct {
		result1 pivnet.FileGroup
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) CreateFileGroup(productSlug string, name string) (pivnet.FileGroup, error) {
	fake.createFileGroupMutex.Lock()
	ret, specificReturn := fake.createFileGroupReturnsOnCall[len(fake.createFileGroupArgsForCall)]
	fake.createFileGroupArgsForCall = append(fake.createFileGroupArgsForCall, struct {
		productSlug string
		name        string
	}{productSlug, name})
	fake.recordInvocation("CreateFileGroup", []interface{}{productSlug, name})
	fake.createFileGroupMutex.Unlock()
	if fake.CreateFileGroupStub != nil {
		return fake.CreateFileGroupStub(productSlug, name)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createFileGroupReturns.result1, fake.createFileGroupReturns.result2
}

func (fake *FakePivnetClient) CreateFileGroupCallCount() int {
	fake.createFileGroupMutex.RLock()
	defer fake.createFileGroupMutex.RUnlock()
	return len(fake.createFileGroupArgsForCall)
}

func (fake *FakePivnetClient) CreateFileGroupArgsForCall(i int) (string, string) {
	fake.createFileGroupMutex.RLock()
	defer fake.createFileGroupMutex.RUnlock()
	return fake.createFileGroupArgsForCall[i].productSlug, fake.createFileGroupArgsForCall[i].name
}

func (fake *FakePivnetClient) CreateFileGroupReturns(result1 pivnet.FileGroup, result2 error) {
	fake.CreateFileGroupStub = nil
	fake.createFileGroupReturns = struct {
		result1 pivnet.FileGroup
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) CreateFileGroupReturnsOnCall(i int, result1 pivnet.FileGroup, result2 error) {
	fake.CreateFileGroupStub = nil
	if fake.createFileGroupReturnsOnCall == nil {
		fake.createFileGroupReturnsOnCall = make(map[int]struct {
			result1 pivnet.FileGroup
			result2 error
		})
	}
	fake.createFileGroupReturnsOnCall[i] = struct {
		result1 pivnet.FileGroup
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) UpdateFileGroup(productSlug string, fileGroup pivnet.FileGroup) (pivnet.FileGroup, error) {
	fake.updateFileGroupMutex.Lock()
	ret, specificReturn := fake.updateFileGroupReturnsOnCall[len(fake.updateFileGroupArgsForCall)]
	fake.updateFileGroupArgsForCall = append(fake.updateFileGroupArgsForCall, struct {
		productSlug string
		fileGroup   pivnet.FileGroup
	}{productSlug, fileGroup})
	fake.recordInvocation("UpdateFileGroup", []interface{}{productSlug, fileGroup})
	fake.updateFileGroupMutex.Unlock()
	if fake.UpdateFileGroupStub != nil {
		return fake.UpdateFileGroupStub(productSlug, fileGroup)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.updateFileGroupReturns.result1, fake.updateFileGroupReturns.result2
}

func (fake *FakePivnetClient) UpdateFileGroupCallCount() int {
	fake.updateFileGroupMutex.RLock()
	defer fake.updateFileGroupMutex.RUnlock()
	return len(fake.updateFileGroupArgsForCall)
}

func (fake *FakePivnetClient) UpdateFileGroupArgsForCall(i int) (string, pivnet.FileGroup) {
	fake.updateFileGroupMutex.RLock()
	defer fake.updateFileGroupMutex.RUnlock()
	return fake.updateFileGroupArgsForCall[i].productSlug, fake.updateFileGroupArgsForCall[i].fileGroup
}

func (fake *FakePivnetClient) UpdateFileGroupReturns(result1 pivnet.FileGroup, result2 error) {
	fake.UpdateFileGroupStub = nil
	fake.updateFileGroupReturns = struct {
		result1 pivnet.FileGroup
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) UpdateFileGroupReturnsOnCall(i int, result1 pivnet.FileGroup, result2 error) {
	fake.UpdateFileGroupStub = nil
	if fake.updateFileGroupReturnsOnCall == nil {
		fake.updateFileGroupReturnsOnCall = make(map[int]struct {
			result1 pivnet.FileGroup
			result2 error
		})
	}
	fake.updateFileGroupReturnsOnCall[i] = struct {
		result1 pivnet.FileGroup
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) DeleteFileGroup(productSlug string, fileGroupID int) (pivnet.FileGroup, error) {
	fake.deleteFileGroupMutex.Lock()
	ret, specificReturn := fake.deleteFileGroupReturnsOnCall[len(fake.deleteFileGroupArgsForCall)]
	fake.deleteFileGroupArgsForCall = append(fake.deleteFileGroupArgsForCall, struct {
		productSlug string
		fileGroupID int
	}{productSlug, fileGroupID})
	fake.recordInvocation("DeleteFileGroup", []interface{}{productSlug, fileGroupID})
	fake.deleteFileGroupMutex.Unlock()
	if fake.DeleteFileGroupStub != nil {
		return fake.DeleteFileGroupStub(productSlug, fileGroupID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.deleteFileGroupReturns.result1, fake.deleteFileGroupReturns.result2
}

func (fake *FakePivnetClient) DeleteFileGroupCallCount() int {
	fake.deleteFileGroupMutex.RLock()
	defer fake.deleteFileGroupMutex.RUnlock()
	return len(fake.deleteFileGroupArgsForCall)
}

func (fake *FakePivnetClient) DeleteFileGroupArgsForCall(i int) (string, int) {
	fake.deleteFileGroupMutex.RLock()
	defer fake.deleteFileGroupMutex.RUnlock()
	return fake.deleteFileGroupArgsForCall[i].productSlug, fake.deleteFileGroupArgsForCall[i].fileGroupID
}

func (fake *FakePivnetClient) DeleteFileGroupReturns(result1 pivnet.FileGroup, result2 error) {
	fake.DeleteFileGroupStub = nil
	fake.deleteFileGroupReturns = struct {
		result1 pivnet.FileGroup
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) DeleteFileGroupReturnsOnCall(i int, result1 pivnet.FileGroup, result2 error) {
	fake.DeleteFileGroupStub = nil
	if fake.deleteFileGroupReturnsOnCall == nil {
		fake.deleteFileGroupReturnsOnCall = make(map[int]struct {
			result1 pivnet.FileGroup
			result2 error
		})
	}
	fake.deleteFileGroupReturnsOnCall[i] = struct {
		result1 pivnet.FileGroup
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) AddFileGroupToRelease(productSlug string, fileGroupID int, releaseID int) error {
	fake.addFileGroupToReleaseMutex.Lock()
	ret, specificReturn := fake.addFileGroupToReleaseReturnsOnCall[len(fake.addFileGroupToReleaseArgsForCall)]
	fake.addFileGroupToReleaseArgsForCall = append(fake.addFileGroupToReleaseArgsForCall, struct {
		productSlug string
		fileGroupID int
		releaseID   int
	}{productSlug, fileGroupID, releaseID})
	fake.recordInvocation("AddFileGroupToRelease", []interface{}{productSlug, fileGroupID, releaseID})
	fake.addFileGroupToReleaseMutex.Unlock()
	if fake.AddFileGroupToReleaseStub != nil {
		return fake.AddFileGroupToReleaseStub(productSlug, fileGroupID, releaseID)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.addFileGroupToReleaseReturns.result1
}

func (fake *FakePivnetClient) AddFileGroupToReleaseCallCount() int {
	fake.addFileGroupToReleaseMutex.RLock()
	defer fake.addFileGroupToReleaseMutex.RUnlock()
	return len(fake.addFileGroupToReleaseArgsForCall)
}

func (fake *FakePivnetClient) AddFileGroupToReleaseArgsForCall(i int) (string, int, int) {
	fake.addFileGroupToReleaseMutex.RLock()
	defer fake.addFileGroupToReleaseMutex.RUnlock()
	return fake.addFileGroupToReleaseArgsForCall[i].productSlug, fake.addFileGroupToReleaseArgsForCall[i].fileGroupID, fake.addFileGroupToReleaseArgsForCall[i].releaseID
}

func (fake *FakePivnetClient) AddFileGroupToReleaseReturns(result1 error) {
	fake.AddFileGroupToReleaseStub = nil
	fake.addFileGroupToReleaseReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePivnetClient) AddFileGroupToReleaseReturnsOnCall(i int, result1 error) {
	fake.AddFileGroupToReleaseStub = nil
	if fake.addFileGroupToReleaseReturnsOnCall == nil {
		fake.addFileGroupToReleaseReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.addFileGroupToReleaseReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakePivnetClient) RemoveFileGroupFromRelease(productSlug string, fileGroupID int, releaseID int) error {
	fake.removeFileGroupFromReleaseMutex.Lock()
	ret, specificReturn := fake.removeFileGroupFromReleaseReturnsOnCall[len(fake.removeFileGroupFromReleaseArgsForCall)]
	fake.removeFileGroupFromReleaseArgsForCall = append(fake.removeFileGroupFromReleaseArgsForCall, struct {
		productSlug string
		fileGroupID int
		releaseID   int
	}{productSlug, fileGroupID, releaseID})
	fake.recordInvocation("RemoveFileGroupFromRelease", []interface{}{productSlug, fileGroupID, releaseID})
	fake.removeFileGroupFromReleaseMutex.Unlock()
	if fake.RemoveFileGroupFromReleaseStub != nil {
		return fake.RemoveFileGroupFromReleaseStub(productSlug, fileGroupID, releaseID)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.removeFileGroupFromReleaseReturns.result1
}

func (fake *FakePivnetClient) RemoveFileGroupFromReleaseCallCount() int {
	fake.removeFileGroupFromReleaseMutex.RLock()
	defer fake.removeFileGroupFromReleaseMutex.RUnlock()
	return len(fake.removeFileGroupFromReleaseArgsForCall)
}

func (fake *FakePivnetClient) RemoveFileGroupFromReleaseArgsForCall(i int) (string, int, int) {
	fake.removeFileGroupFromReleaseMutex.RLock()
	defer fake.removeFileGroupFromReleaseMutex.RUnlock()
	return fake.removeFileGroupFromReleaseArgsForCall[i].productSlug, fake.removeFileGroupFromReleaseArgsForCall[i].fileGroupID, fake.removeFileGroupFromReleaseArgsForCall[i].releaseID
}

func (fake *FakePivnetClient) RemoveFileGroupFromReleaseReturns(result1 error) {
	fake.RemoveFileGroupFromReleaseStub = nil
	fake.removeFileGroupFromReleaseReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePivnetClient) RemoveFileGroupFromReleaseReturnsOnCall(i int, result1 error) {
	fake.RemoveFileGroupFromReleaseStub = nil
	if fake.removeFileGroupFromReleaseReturnsOnCall == nil {
		fake.removeFileGroupFromReleaseReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.removeFileGroupFromReleaseReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakePivnetClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.fileGroupsMutex.RLock()
	defer fake.fileGroupsMutex.RUnlock()
	fake.fileGroupsForReleaseMutex.RLock()
	defer fake.fileGroupsForReleaseMutex.RUnlock()
	fake.releaseForVersionMutex.RLock()
	defer fake.releaseForVersionMutex.RUnlock()
	fake.fileGroupMutex.RLock()
	defer fake.fileGroupMutex.RUnlock()
	fake.createFileGroupMutex.RLock()
	defer fake.createFileGroupMutex.RUnlock()
	fake.updateFileGroupMutex.RLock()
	defer fake.updateFileGroupMutex.RUnlock()
	fake.deleteFileGroupMutex.RLock()
	defer fake.deleteFileGroupMutex.RUnlock()
	fake.addFileGroupToReleaseMutex.RLock()
	defer fake.addFileGroupToReleaseMutex.RUnlock()
	fake.removeFileGroupFromReleaseMutex.RLock()
	defer fake.removeFileGroupFromReleaseMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakePivnetClient) recordInvocation(key string, args []interface{}) {
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

var _ filegroup.PivnetClient = new(FakePivnetClient)
