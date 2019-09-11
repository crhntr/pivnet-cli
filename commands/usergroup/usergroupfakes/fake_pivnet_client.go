// Code generated by counterfeiter. DO NOT EDIT.
package usergroupfakes

import (
	"sync"

	"github.com/pivotal-cf/go-pivnet/v2"
	"github.com/pivotal-cf/pivnet-cli/commands/usergroup"
)

type FakePivnetClient struct {
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
	UserGroupsStub        func() ([]pivnet.UserGroup, error)
	userGroupsMutex       sync.RWMutex
	userGroupsArgsForCall []struct{}
	userGroupsReturns     struct {
		result1 []pivnet.UserGroup
		result2 error
	}
	userGroupsReturnsOnCall map[int]struct {
		result1 []pivnet.UserGroup
		result2 error
	}
	UserGroupsForReleaseStub        func(productSlug string, releaseID int) ([]pivnet.UserGroup, error)
	userGroupsForReleaseMutex       sync.RWMutex
	userGroupsForReleaseArgsForCall []struct {
		productSlug string
		releaseID   int
	}
	userGroupsForReleaseReturns struct {
		result1 []pivnet.UserGroup
		result2 error
	}
	userGroupsForReleaseReturnsOnCall map[int]struct {
		result1 []pivnet.UserGroup
		result2 error
	}
	UserGroupStub        func(userGroupID int) (pivnet.UserGroup, error)
	userGroupMutex       sync.RWMutex
	userGroupArgsForCall []struct {
		userGroupID int
	}
	userGroupReturns struct {
		result1 pivnet.UserGroup
		result2 error
	}
	userGroupReturnsOnCall map[int]struct {
		result1 pivnet.UserGroup
		result2 error
	}
	CreateUserGroupStub        func(name string, description string, members []string) (pivnet.UserGroup, error)
	createUserGroupMutex       sync.RWMutex
	createUserGroupArgsForCall []struct {
		name        string
		description string
		members     []string
	}
	createUserGroupReturns struct {
		result1 pivnet.UserGroup
		result2 error
	}
	createUserGroupReturnsOnCall map[int]struct {
		result1 pivnet.UserGroup
		result2 error
	}
	UpdateUserGroupStub        func(userGroup pivnet.UserGroup) (pivnet.UserGroup, error)
	updateUserGroupMutex       sync.RWMutex
	updateUserGroupArgsForCall []struct {
		userGroup pivnet.UserGroup
	}
	updateUserGroupReturns struct {
		result1 pivnet.UserGroup
		result2 error
	}
	updateUserGroupReturnsOnCall map[int]struct {
		result1 pivnet.UserGroup
		result2 error
	}
	DeleteUserGroupStub        func(userGroupID int) error
	deleteUserGroupMutex       sync.RWMutex
	deleteUserGroupArgsForCall []struct {
		userGroupID int
	}
	deleteUserGroupReturns struct {
		result1 error
	}
	deleteUserGroupReturnsOnCall map[int]struct {
		result1 error
	}
	AddUserGroupStub        func(productSlug string, releaseID int, userGroupID int) error
	addUserGroupMutex       sync.RWMutex
	addUserGroupArgsForCall []struct {
		productSlug string
		releaseID   int
		userGroupID int
	}
	addUserGroupReturns struct {
		result1 error
	}
	addUserGroupReturnsOnCall map[int]struct {
		result1 error
	}
	RemoveUserGroupStub        func(productSlug string, releaseID int, userGroupID int) error
	removeUserGroupMutex       sync.RWMutex
	removeUserGroupArgsForCall []struct {
		productSlug string
		releaseID   int
		userGroupID int
	}
	removeUserGroupReturns struct {
		result1 error
	}
	removeUserGroupReturnsOnCall map[int]struct {
		result1 error
	}
	AddMemberToGroupStub        func(userGroupID int, emailAddress string, admin bool) (pivnet.UserGroup, error)
	addMemberToGroupMutex       sync.RWMutex
	addMemberToGroupArgsForCall []struct {
		userGroupID  int
		emailAddress string
		admin        bool
	}
	addMemberToGroupReturns struct {
		result1 pivnet.UserGroup
		result2 error
	}
	addMemberToGroupReturnsOnCall map[int]struct {
		result1 pivnet.UserGroup
		result2 error
	}
	RemoveMemberFromGroupStub        func(userGroupID int, emailAddress string) (pivnet.UserGroup, error)
	removeMemberFromGroupMutex       sync.RWMutex
	removeMemberFromGroupArgsForCall []struct {
		userGroupID  int
		emailAddress string
	}
	removeMemberFromGroupReturns struct {
		result1 pivnet.UserGroup
		result2 error
	}
	removeMemberFromGroupReturnsOnCall map[int]struct {
		result1 pivnet.UserGroup
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
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

func (fake *FakePivnetClient) UserGroups() ([]pivnet.UserGroup, error) {
	fake.userGroupsMutex.Lock()
	ret, specificReturn := fake.userGroupsReturnsOnCall[len(fake.userGroupsArgsForCall)]
	fake.userGroupsArgsForCall = append(fake.userGroupsArgsForCall, struct{}{})
	fake.recordInvocation("UserGroups", []interface{}{})
	fake.userGroupsMutex.Unlock()
	if fake.UserGroupsStub != nil {
		return fake.UserGroupsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.userGroupsReturns.result1, fake.userGroupsReturns.result2
}

func (fake *FakePivnetClient) UserGroupsCallCount() int {
	fake.userGroupsMutex.RLock()
	defer fake.userGroupsMutex.RUnlock()
	return len(fake.userGroupsArgsForCall)
}

func (fake *FakePivnetClient) UserGroupsReturns(result1 []pivnet.UserGroup, result2 error) {
	fake.UserGroupsStub = nil
	fake.userGroupsReturns = struct {
		result1 []pivnet.UserGroup
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) UserGroupsReturnsOnCall(i int, result1 []pivnet.UserGroup, result2 error) {
	fake.UserGroupsStub = nil
	if fake.userGroupsReturnsOnCall == nil {
		fake.userGroupsReturnsOnCall = make(map[int]struct {
			result1 []pivnet.UserGroup
			result2 error
		})
	}
	fake.userGroupsReturnsOnCall[i] = struct {
		result1 []pivnet.UserGroup
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) UserGroupsForRelease(productSlug string, releaseID int) ([]pivnet.UserGroup, error) {
	fake.userGroupsForReleaseMutex.Lock()
	ret, specificReturn := fake.userGroupsForReleaseReturnsOnCall[len(fake.userGroupsForReleaseArgsForCall)]
	fake.userGroupsForReleaseArgsForCall = append(fake.userGroupsForReleaseArgsForCall, struct {
		productSlug string
		releaseID   int
	}{productSlug, releaseID})
	fake.recordInvocation("UserGroupsForRelease", []interface{}{productSlug, releaseID})
	fake.userGroupsForReleaseMutex.Unlock()
	if fake.UserGroupsForReleaseStub != nil {
		return fake.UserGroupsForReleaseStub(productSlug, releaseID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.userGroupsForReleaseReturns.result1, fake.userGroupsForReleaseReturns.result2
}

func (fake *FakePivnetClient) UserGroupsForReleaseCallCount() int {
	fake.userGroupsForReleaseMutex.RLock()
	defer fake.userGroupsForReleaseMutex.RUnlock()
	return len(fake.userGroupsForReleaseArgsForCall)
}

func (fake *FakePivnetClient) UserGroupsForReleaseArgsForCall(i int) (string, int) {
	fake.userGroupsForReleaseMutex.RLock()
	defer fake.userGroupsForReleaseMutex.RUnlock()
	return fake.userGroupsForReleaseArgsForCall[i].productSlug, fake.userGroupsForReleaseArgsForCall[i].releaseID
}

func (fake *FakePivnetClient) UserGroupsForReleaseReturns(result1 []pivnet.UserGroup, result2 error) {
	fake.UserGroupsForReleaseStub = nil
	fake.userGroupsForReleaseReturns = struct {
		result1 []pivnet.UserGroup
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) UserGroupsForReleaseReturnsOnCall(i int, result1 []pivnet.UserGroup, result2 error) {
	fake.UserGroupsForReleaseStub = nil
	if fake.userGroupsForReleaseReturnsOnCall == nil {
		fake.userGroupsForReleaseReturnsOnCall = make(map[int]struct {
			result1 []pivnet.UserGroup
			result2 error
		})
	}
	fake.userGroupsForReleaseReturnsOnCall[i] = struct {
		result1 []pivnet.UserGroup
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) UserGroup(userGroupID int) (pivnet.UserGroup, error) {
	fake.userGroupMutex.Lock()
	ret, specificReturn := fake.userGroupReturnsOnCall[len(fake.userGroupArgsForCall)]
	fake.userGroupArgsForCall = append(fake.userGroupArgsForCall, struct {
		userGroupID int
	}{userGroupID})
	fake.recordInvocation("UserGroup", []interface{}{userGroupID})
	fake.userGroupMutex.Unlock()
	if fake.UserGroupStub != nil {
		return fake.UserGroupStub(userGroupID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.userGroupReturns.result1, fake.userGroupReturns.result2
}

func (fake *FakePivnetClient) UserGroupCallCount() int {
	fake.userGroupMutex.RLock()
	defer fake.userGroupMutex.RUnlock()
	return len(fake.userGroupArgsForCall)
}

func (fake *FakePivnetClient) UserGroupArgsForCall(i int) int {
	fake.userGroupMutex.RLock()
	defer fake.userGroupMutex.RUnlock()
	return fake.userGroupArgsForCall[i].userGroupID
}

func (fake *FakePivnetClient) UserGroupReturns(result1 pivnet.UserGroup, result2 error) {
	fake.UserGroupStub = nil
	fake.userGroupReturns = struct {
		result1 pivnet.UserGroup
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) UserGroupReturnsOnCall(i int, result1 pivnet.UserGroup, result2 error) {
	fake.UserGroupStub = nil
	if fake.userGroupReturnsOnCall == nil {
		fake.userGroupReturnsOnCall = make(map[int]struct {
			result1 pivnet.UserGroup
			result2 error
		})
	}
	fake.userGroupReturnsOnCall[i] = struct {
		result1 pivnet.UserGroup
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) CreateUserGroup(name string, description string, members []string) (pivnet.UserGroup, error) {
	var membersCopy []string
	if members != nil {
		membersCopy = make([]string, len(members))
		copy(membersCopy, members)
	}
	fake.createUserGroupMutex.Lock()
	ret, specificReturn := fake.createUserGroupReturnsOnCall[len(fake.createUserGroupArgsForCall)]
	fake.createUserGroupArgsForCall = append(fake.createUserGroupArgsForCall, struct {
		name        string
		description string
		members     []string
	}{name, description, membersCopy})
	fake.recordInvocation("CreateUserGroup", []interface{}{name, description, membersCopy})
	fake.createUserGroupMutex.Unlock()
	if fake.CreateUserGroupStub != nil {
		return fake.CreateUserGroupStub(name, description, members)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createUserGroupReturns.result1, fake.createUserGroupReturns.result2
}

func (fake *FakePivnetClient) CreateUserGroupCallCount() int {
	fake.createUserGroupMutex.RLock()
	defer fake.createUserGroupMutex.RUnlock()
	return len(fake.createUserGroupArgsForCall)
}

func (fake *FakePivnetClient) CreateUserGroupArgsForCall(i int) (string, string, []string) {
	fake.createUserGroupMutex.RLock()
	defer fake.createUserGroupMutex.RUnlock()
	return fake.createUserGroupArgsForCall[i].name, fake.createUserGroupArgsForCall[i].description, fake.createUserGroupArgsForCall[i].members
}

func (fake *FakePivnetClient) CreateUserGroupReturns(result1 pivnet.UserGroup, result2 error) {
	fake.CreateUserGroupStub = nil
	fake.createUserGroupReturns = struct {
		result1 pivnet.UserGroup
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) CreateUserGroupReturnsOnCall(i int, result1 pivnet.UserGroup, result2 error) {
	fake.CreateUserGroupStub = nil
	if fake.createUserGroupReturnsOnCall == nil {
		fake.createUserGroupReturnsOnCall = make(map[int]struct {
			result1 pivnet.UserGroup
			result2 error
		})
	}
	fake.createUserGroupReturnsOnCall[i] = struct {
		result1 pivnet.UserGroup
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) UpdateUserGroup(userGroup pivnet.UserGroup) (pivnet.UserGroup, error) {
	fake.updateUserGroupMutex.Lock()
	ret, specificReturn := fake.updateUserGroupReturnsOnCall[len(fake.updateUserGroupArgsForCall)]
	fake.updateUserGroupArgsForCall = append(fake.updateUserGroupArgsForCall, struct {
		userGroup pivnet.UserGroup
	}{userGroup})
	fake.recordInvocation("UpdateUserGroup", []interface{}{userGroup})
	fake.updateUserGroupMutex.Unlock()
	if fake.UpdateUserGroupStub != nil {
		return fake.UpdateUserGroupStub(userGroup)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.updateUserGroupReturns.result1, fake.updateUserGroupReturns.result2
}

func (fake *FakePivnetClient) UpdateUserGroupCallCount() int {
	fake.updateUserGroupMutex.RLock()
	defer fake.updateUserGroupMutex.RUnlock()
	return len(fake.updateUserGroupArgsForCall)
}

func (fake *FakePivnetClient) UpdateUserGroupArgsForCall(i int) pivnet.UserGroup {
	fake.updateUserGroupMutex.RLock()
	defer fake.updateUserGroupMutex.RUnlock()
	return fake.updateUserGroupArgsForCall[i].userGroup
}

func (fake *FakePivnetClient) UpdateUserGroupReturns(result1 pivnet.UserGroup, result2 error) {
	fake.UpdateUserGroupStub = nil
	fake.updateUserGroupReturns = struct {
		result1 pivnet.UserGroup
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) UpdateUserGroupReturnsOnCall(i int, result1 pivnet.UserGroup, result2 error) {
	fake.UpdateUserGroupStub = nil
	if fake.updateUserGroupReturnsOnCall == nil {
		fake.updateUserGroupReturnsOnCall = make(map[int]struct {
			result1 pivnet.UserGroup
			result2 error
		})
	}
	fake.updateUserGroupReturnsOnCall[i] = struct {
		result1 pivnet.UserGroup
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) DeleteUserGroup(userGroupID int) error {
	fake.deleteUserGroupMutex.Lock()
	ret, specificReturn := fake.deleteUserGroupReturnsOnCall[len(fake.deleteUserGroupArgsForCall)]
	fake.deleteUserGroupArgsForCall = append(fake.deleteUserGroupArgsForCall, struct {
		userGroupID int
	}{userGroupID})
	fake.recordInvocation("DeleteUserGroup", []interface{}{userGroupID})
	fake.deleteUserGroupMutex.Unlock()
	if fake.DeleteUserGroupStub != nil {
		return fake.DeleteUserGroupStub(userGroupID)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.deleteUserGroupReturns.result1
}

func (fake *FakePivnetClient) DeleteUserGroupCallCount() int {
	fake.deleteUserGroupMutex.RLock()
	defer fake.deleteUserGroupMutex.RUnlock()
	return len(fake.deleteUserGroupArgsForCall)
}

func (fake *FakePivnetClient) DeleteUserGroupArgsForCall(i int) int {
	fake.deleteUserGroupMutex.RLock()
	defer fake.deleteUserGroupMutex.RUnlock()
	return fake.deleteUserGroupArgsForCall[i].userGroupID
}

func (fake *FakePivnetClient) DeleteUserGroupReturns(result1 error) {
	fake.DeleteUserGroupStub = nil
	fake.deleteUserGroupReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePivnetClient) DeleteUserGroupReturnsOnCall(i int, result1 error) {
	fake.DeleteUserGroupStub = nil
	if fake.deleteUserGroupReturnsOnCall == nil {
		fake.deleteUserGroupReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteUserGroupReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakePivnetClient) AddUserGroup(productSlug string, releaseID int, userGroupID int) error {
	fake.addUserGroupMutex.Lock()
	ret, specificReturn := fake.addUserGroupReturnsOnCall[len(fake.addUserGroupArgsForCall)]
	fake.addUserGroupArgsForCall = append(fake.addUserGroupArgsForCall, struct {
		productSlug string
		releaseID   int
		userGroupID int
	}{productSlug, releaseID, userGroupID})
	fake.recordInvocation("AddUserGroup", []interface{}{productSlug, releaseID, userGroupID})
	fake.addUserGroupMutex.Unlock()
	if fake.AddUserGroupStub != nil {
		return fake.AddUserGroupStub(productSlug, releaseID, userGroupID)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.addUserGroupReturns.result1
}

func (fake *FakePivnetClient) AddUserGroupCallCount() int {
	fake.addUserGroupMutex.RLock()
	defer fake.addUserGroupMutex.RUnlock()
	return len(fake.addUserGroupArgsForCall)
}

func (fake *FakePivnetClient) AddUserGroupArgsForCall(i int) (string, int, int) {
	fake.addUserGroupMutex.RLock()
	defer fake.addUserGroupMutex.RUnlock()
	return fake.addUserGroupArgsForCall[i].productSlug, fake.addUserGroupArgsForCall[i].releaseID, fake.addUserGroupArgsForCall[i].userGroupID
}

func (fake *FakePivnetClient) AddUserGroupReturns(result1 error) {
	fake.AddUserGroupStub = nil
	fake.addUserGroupReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePivnetClient) AddUserGroupReturnsOnCall(i int, result1 error) {
	fake.AddUserGroupStub = nil
	if fake.addUserGroupReturnsOnCall == nil {
		fake.addUserGroupReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.addUserGroupReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakePivnetClient) RemoveUserGroup(productSlug string, releaseID int, userGroupID int) error {
	fake.removeUserGroupMutex.Lock()
	ret, specificReturn := fake.removeUserGroupReturnsOnCall[len(fake.removeUserGroupArgsForCall)]
	fake.removeUserGroupArgsForCall = append(fake.removeUserGroupArgsForCall, struct {
		productSlug string
		releaseID   int
		userGroupID int
	}{productSlug, releaseID, userGroupID})
	fake.recordInvocation("RemoveUserGroup", []interface{}{productSlug, releaseID, userGroupID})
	fake.removeUserGroupMutex.Unlock()
	if fake.RemoveUserGroupStub != nil {
		return fake.RemoveUserGroupStub(productSlug, releaseID, userGroupID)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.removeUserGroupReturns.result1
}

func (fake *FakePivnetClient) RemoveUserGroupCallCount() int {
	fake.removeUserGroupMutex.RLock()
	defer fake.removeUserGroupMutex.RUnlock()
	return len(fake.removeUserGroupArgsForCall)
}

func (fake *FakePivnetClient) RemoveUserGroupArgsForCall(i int) (string, int, int) {
	fake.removeUserGroupMutex.RLock()
	defer fake.removeUserGroupMutex.RUnlock()
	return fake.removeUserGroupArgsForCall[i].productSlug, fake.removeUserGroupArgsForCall[i].releaseID, fake.removeUserGroupArgsForCall[i].userGroupID
}

func (fake *FakePivnetClient) RemoveUserGroupReturns(result1 error) {
	fake.RemoveUserGroupStub = nil
	fake.removeUserGroupReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePivnetClient) RemoveUserGroupReturnsOnCall(i int, result1 error) {
	fake.RemoveUserGroupStub = nil
	if fake.removeUserGroupReturnsOnCall == nil {
		fake.removeUserGroupReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.removeUserGroupReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakePivnetClient) AddMemberToGroup(userGroupID int, emailAddress string, admin bool) (pivnet.UserGroup, error) {
	fake.addMemberToGroupMutex.Lock()
	ret, specificReturn := fake.addMemberToGroupReturnsOnCall[len(fake.addMemberToGroupArgsForCall)]
	fake.addMemberToGroupArgsForCall = append(fake.addMemberToGroupArgsForCall, struct {
		userGroupID  int
		emailAddress string
		admin        bool
	}{userGroupID, emailAddress, admin})
	fake.recordInvocation("AddMemberToGroup", []interface{}{userGroupID, emailAddress, admin})
	fake.addMemberToGroupMutex.Unlock()
	if fake.AddMemberToGroupStub != nil {
		return fake.AddMemberToGroupStub(userGroupID, emailAddress, admin)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.addMemberToGroupReturns.result1, fake.addMemberToGroupReturns.result2
}

func (fake *FakePivnetClient) AddMemberToGroupCallCount() int {
	fake.addMemberToGroupMutex.RLock()
	defer fake.addMemberToGroupMutex.RUnlock()
	return len(fake.addMemberToGroupArgsForCall)
}

func (fake *FakePivnetClient) AddMemberToGroupArgsForCall(i int) (int, string, bool) {
	fake.addMemberToGroupMutex.RLock()
	defer fake.addMemberToGroupMutex.RUnlock()
	return fake.addMemberToGroupArgsForCall[i].userGroupID, fake.addMemberToGroupArgsForCall[i].emailAddress, fake.addMemberToGroupArgsForCall[i].admin
}

func (fake *FakePivnetClient) AddMemberToGroupReturns(result1 pivnet.UserGroup, result2 error) {
	fake.AddMemberToGroupStub = nil
	fake.addMemberToGroupReturns = struct {
		result1 pivnet.UserGroup
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) AddMemberToGroupReturnsOnCall(i int, result1 pivnet.UserGroup, result2 error) {
	fake.AddMemberToGroupStub = nil
	if fake.addMemberToGroupReturnsOnCall == nil {
		fake.addMemberToGroupReturnsOnCall = make(map[int]struct {
			result1 pivnet.UserGroup
			result2 error
		})
	}
	fake.addMemberToGroupReturnsOnCall[i] = struct {
		result1 pivnet.UserGroup
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) RemoveMemberFromGroup(userGroupID int, emailAddress string) (pivnet.UserGroup, error) {
	fake.removeMemberFromGroupMutex.Lock()
	ret, specificReturn := fake.removeMemberFromGroupReturnsOnCall[len(fake.removeMemberFromGroupArgsForCall)]
	fake.removeMemberFromGroupArgsForCall = append(fake.removeMemberFromGroupArgsForCall, struct {
		userGroupID  int
		emailAddress string
	}{userGroupID, emailAddress})
	fake.recordInvocation("RemoveMemberFromGroup", []interface{}{userGroupID, emailAddress})
	fake.removeMemberFromGroupMutex.Unlock()
	if fake.RemoveMemberFromGroupStub != nil {
		return fake.RemoveMemberFromGroupStub(userGroupID, emailAddress)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.removeMemberFromGroupReturns.result1, fake.removeMemberFromGroupReturns.result2
}

func (fake *FakePivnetClient) RemoveMemberFromGroupCallCount() int {
	fake.removeMemberFromGroupMutex.RLock()
	defer fake.removeMemberFromGroupMutex.RUnlock()
	return len(fake.removeMemberFromGroupArgsForCall)
}

func (fake *FakePivnetClient) RemoveMemberFromGroupArgsForCall(i int) (int, string) {
	fake.removeMemberFromGroupMutex.RLock()
	defer fake.removeMemberFromGroupMutex.RUnlock()
	return fake.removeMemberFromGroupArgsForCall[i].userGroupID, fake.removeMemberFromGroupArgsForCall[i].emailAddress
}

func (fake *FakePivnetClient) RemoveMemberFromGroupReturns(result1 pivnet.UserGroup, result2 error) {
	fake.RemoveMemberFromGroupStub = nil
	fake.removeMemberFromGroupReturns = struct {
		result1 pivnet.UserGroup
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) RemoveMemberFromGroupReturnsOnCall(i int, result1 pivnet.UserGroup, result2 error) {
	fake.RemoveMemberFromGroupStub = nil
	if fake.removeMemberFromGroupReturnsOnCall == nil {
		fake.removeMemberFromGroupReturnsOnCall = make(map[int]struct {
			result1 pivnet.UserGroup
			result2 error
		})
	}
	fake.removeMemberFromGroupReturnsOnCall[i] = struct {
		result1 pivnet.UserGroup
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.releaseForVersionMutex.RLock()
	defer fake.releaseForVersionMutex.RUnlock()
	fake.userGroupsMutex.RLock()
	defer fake.userGroupsMutex.RUnlock()
	fake.userGroupsForReleaseMutex.RLock()
	defer fake.userGroupsForReleaseMutex.RUnlock()
	fake.userGroupMutex.RLock()
	defer fake.userGroupMutex.RUnlock()
	fake.createUserGroupMutex.RLock()
	defer fake.createUserGroupMutex.RUnlock()
	fake.updateUserGroupMutex.RLock()
	defer fake.updateUserGroupMutex.RUnlock()
	fake.deleteUserGroupMutex.RLock()
	defer fake.deleteUserGroupMutex.RUnlock()
	fake.addUserGroupMutex.RLock()
	defer fake.addUserGroupMutex.RUnlock()
	fake.removeUserGroupMutex.RLock()
	defer fake.removeUserGroupMutex.RUnlock()
	fake.addMemberToGroupMutex.RLock()
	defer fake.addMemberToGroupMutex.RUnlock()
	fake.removeMemberFromGroupMutex.RLock()
	defer fake.removeMemberFromGroupMutex.RUnlock()
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

var _ usergroup.PivnetClient = new(FakePivnetClient)
