package permission

import (
	"github.com/stretchr/testify/assert"
	"mini-roles-backend/source/domains/permission/mock"
	"mini-roles-backend/source/domains/permission/spec"
	sharedMock "mini-roles-backend/source/domains/shared/mock"
	sharedModels "mini-roles-backend/source/domains/shared/models"
	"testing"
)

func TestRepository_ListEmpty(t *testing.T) {
	permissions, err := New(sharedModels.AppData{}).List(spec.PermissionWithAccountIdAndRoleId{
		AccountId: sharedMock.ExistsAccountId,
		RoleId:    sharedMock.ExistsRoleId,
	})

	assert.Nil(t, err)
	assert.Empty(t, permissions)
}

func TestRepository_ListEmptyDueVersion(t *testing.T) {
	permissions, err := New(mock.MakeAppDataForAllRoles(mock.FlatRoleId1)).List(spec.PermissionWithAccountIdAndRoleId{
		AccountId:      sharedMock.ExistsAccountId,
		RoleId:         mock.FlatRoleId1,
		RolesVersionId: "foo-bar",
	})

	assert.Nil(t, err)
	assert.Empty(t, permissions)
}

func TestRepository_ListFlatRolePermissions(t *testing.T) {
	permissions, err := New(mock.MakeAppDataForAllRoles(mock.FlatRoleId1)).List(spec.PermissionWithAccountIdAndRoleId{
		AccountId:      sharedMock.ExistsAccountId,
		RoleId:         mock.FlatRoleId1,
		RolesVersionId: sharedMock.ExistsRolesVersionId,
	})

	assert.Nil(t, err)
	for _, role1Permission := range mock.MakeRole1Permissions() {
		assert.Contains(t, permissions, role1Permission)
	}
}

func TestRepository_ListFlatRoleWithAnotherVersionPermissions(t *testing.T) {
	permissions, err := New(mock.MakeAppDataForAllRoles(mock.FlatRoleId1)).List(spec.PermissionWithAccountIdAndRoleId{
		AccountId:      sharedMock.ExistsAccountId,
		RoleId:         mock.FlatRoleId1,
		RolesVersionId: sharedMock.ExistsRolesVersionId2,
	})

	assert.Nil(t, err)
	for _, rolesWithAnotherVersionPermission := range mock.MakeRoles1WithAnotherVersionPermissions() {
		assert.Contains(t, permissions, rolesWithAnotherVersionPermission)
	}
}

func TestRepository_ListOneDepthLevelExtending(t *testing.T) {
	permissions, err := New(mock.MakeAppDataForAllRoles(mock.OneDepthLevelExtendingRoleId)).List(spec.PermissionWithAccountIdAndRoleId{
		AccountId:      sharedMock.ExistsAccountId,
		RoleId:         mock.OneDepthLevelExtendingRoleId,
		RolesVersionId: sharedMock.ExistsRolesVersionId,
	})

	assert.Nil(t, err)
	for _, role1Permission := range mock.MakeRole1Permissions() {
		assert.Contains(t, permissions, role1Permission)
	}

	for _, extendingRolePermission := range mock.MakeExtendingRolePermissions() {
		assert.Contains(t, permissions, extendingRolePermission)
	}

	for _, role2Permission := range mock.MakeRole2Permissions() {
		assert.NotContains(t, permissions, role2Permission)
	}
}

func TestRepository_ListTwoDepthLevelExtending(t *testing.T) {
	permissions, err := New(mock.MakeAppDataForAllRoles(mock.TwoDepthLevelExtendingRoleId)).List(spec.PermissionWithAccountIdAndRoleId{
		AccountId:      sharedMock.ExistsAccountId,
		RoleId:         mock.TwoDepthLevelExtendingRoleId,
		RolesVersionId: sharedMock.ExistsRolesVersionId,
	})

	assert.Nil(t, err)
	for _, extendingRolePermission := range mock.Permissions {
		assert.Contains(t, permissions, extendingRolePermission)
	}
}

func TestRepository_ListRecursiveRolesExtending(t *testing.T) {
	permissions, err := New(mock.MakeAppDataForAllRoles(mock.RecursiveExtendingRoleId1)).List(spec.PermissionWithAccountIdAndRoleId{
		AccountId:      sharedMock.ExistsAccountId,
		RoleId:         mock.RecursiveExtendingRoleId1,
		RolesVersionId: sharedMock.ExistsRolesVersionId,
	})

	assert.Nil(t, err)
	for _, expectedPermission := range append(
		mock.MakeRecursiveExtendingRole1Permissions(),
		mock.MakeRecursiveExtendingRole2Permissions()...,
	) {
		assert.Contains(t, permissions, expectedPermission)
	}
}
