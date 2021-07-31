package permission

import (
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"mini-roles-backend/source/config"
	"mini-roles-backend/source/db/connection"
	"mini-roles-backend/source/domains/permission/mock"
	"mini-roles-backend/source/domains/permission/spec"
	sharedMock "mini-roles-backend/source/domains/shared/mock"
	sharedModels "mini-roles-backend/source/domains/shared/models"
	"testing"
)

var (
	db         *sqlx.DB
	repository Repository
)

func init() {
	db = sqlx.MustOpen(
		"postgres",
		connection.BuildPostgresString(config.Default()),
	)
	repository = New(db)

	sharedMock.MustReinstall(db)
}

func initTestData() {
	addTestResources()
	addTestPermissions()
	sharedMock.MustAddRolesVersion(db, sharedMock.ExistsRolesVersionId2)
	addTestFlatRoles()
}

func TestRepository_ListEmpty(t *testing.T) {
	permissions, err := repository.List(spec.PermissionWithAccountIdAndRoleId{
		AccountId:      sharedMock.ExistsAccountId,
		RoleId:         sharedMock.ExistsRoleId,
		RolesVersionId: sharedMock.ExistsRolesVersionId,
	})

	assert.Nil(t, err)
	assert.Empty(t, permissions)
}

func TestRepository_ListEmptyDueVersion(t *testing.T) {
	initTestData()
	defer sharedMock.MustReinstall(db)

	permissions, err := repository.List(spec.PermissionWithAccountIdAndRoleId{
		AccountId:      sharedMock.ExistsAccountId,
		RoleId:         mock.FlatRoleId1,
		RolesVersionId: "foo-bar",
	})

	assert.Nil(t, err)
	assert.Empty(t, permissions)
}

func TestRepository_ListFlatRolePermissions(t *testing.T) {
	initTestData()
	defer sharedMock.MustReinstall(db)

	permissions, err := repository.List(spec.PermissionWithAccountIdAndRoleId{
		AccountId:      sharedMock.ExistsAccountId,
		RoleId:         mock.FlatRoleId1,
		RolesVersionId: sharedMock.ExistsRolesVersionId,
	})

	assert.Nil(t, err)
	for _, role1Permission := range mock.MakeRole1Permissions() {
		assert.Contains(t, permissions, role1Permission)
	}
}

func TestRepository_ListFlatRoleWithAnotherVersion(t *testing.T) {
	initTestData()
	defer sharedMock.MustReinstall(db)

	permissions, err := repository.List(spec.PermissionWithAccountIdAndRoleId{
		AccountId:      sharedMock.ExistsAccountId,
		RoleId:         mock.FlatRoleId1,
		RolesVersionId: sharedMock.ExistsRolesVersionId2,
	})

	assert.Nil(t, err)
	for _, role1Permission := range mock.MakeRole2Permissions() {
		assert.Contains(t, permissions, role1Permission)
	}
}

func TestRepository_ListOneDepthLevelExtending(t *testing.T) {
	initTestData()
	addTestOneDepthLevelExtendingRole()
	defer sharedMock.MustReinstall(db)

	permissions, err := repository.List(spec.PermissionWithAccountIdAndRoleId{
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
	initTestData()
	addTestOneDepthLevelExtendingRole()
	addTestTwoDepthLevelExtendingRole()
	defer sharedMock.MustReinstall(db)

	permissions, err := repository.List(spec.PermissionWithAccountIdAndRoleId{
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
	initTestData()
	addTestRecursiveExtendingRoles()
	defer sharedMock.MustReinstall(db)

	permissions, err := repository.List(spec.PermissionWithAccountIdAndRoleId{
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

func TestRepository_ListError(t *testing.T) {
	dropFunctionQuery := `
	drop function if exists recursive_permissions(roles_version_id character(32), entry_point_role_id character(32), _account_hash character(32), depth int, exclude character(32)[])`
	db.MustExec(dropFunctionQuery)
	db.MustExec(`
	create or replace function recursive_permissions(a character(32), b character(32), c character(32), d int, e character(32)[])
	returns void
	language plpgsql
	as $$
	    begin
	        raise log 'hello, % % % % %', a, b, c, d, e;
		end
	$$`)
	defer db.MustExec(dropFunctionQuery)

	_, err := repository.List(spec.PermissionWithAccountIdAndRoleId{
		AccountId:      sharedMock.ExistsAccountId,
		RoleId:         sharedMock.ExistsRoleId,
		RolesVersionId: sharedMock.ExistsRolesVersionId,
	})

	assert.NotNil(t, err)
}

func addTestResources() {
	tx := db.MustBegin()
	for _, resource := range mock.Resources {
		_, err := tx.NamedExec(
			`insert into resource(resource_id, links_to, account_hash) values(:resource_id, :links_to, :account_hash)`,
			map[string]interface{}{
				"resource_id":  resource.Id,
				"links_to":     pq.Array(resource.LinksTo),
				"account_hash": sharedMock.ExistsAccountId,
			},
		)
		if err != nil {
			panic(err)
		}
	}

	err := tx.Commit()
	if err != nil {
		panic(err)
	}
}

func addTestPermissions() {
	tx := db.MustBegin()
	for _, permission := range mock.Permissions {
		_, err := tx.NamedExec(
			`insert into permission(permission_id, operation, effect, resource_id, account_hash)
					values(:permission_id, :operation, :effect, :resource_id, :account_hash)`,
			map[string]interface{}{
				"permission_id": permission.Id,
				"operation":     permission.Operation,
				"effect":        permission.Effect,
				"resource_id":   permission.Resource.Id,
				"account_hash":  sharedMock.ExistsAccountId,
			},
		)
		if err != nil {
			panic(err)
		}
	}

	err := tx.Commit()
	if err != nil {
		panic(err)
	}
}

func addTestFlatRoles() {
	for _, flatRole := range mock.FlatRoles {
		mustAddRole(flatRole)
	}
}

func addTestOneDepthLevelExtendingRole() {
	mustAddRole(mock.OneDepthLevelExtendingRole)
}

func addTestTwoDepthLevelExtendingRole() {
	mustAddRole(mock.TwoDepthLevelExtendingRole)
}

func addTestRecursiveExtendingRoles() {
	for _, role := range []sharedModels.Role{
		mock.RecursiveExtendingRole1,
		mock.RecursiveExtendingRole2,
	} {
		mustAddRole(role)
	}
}

func mustAddRole(role sharedModels.Role) {
	_, err := db.NamedExec(
		`insert into role(role_id, permissions, extends, account_hash, version_id) values(:role_id, :permissions, :extends, :account_hash, :version_id)`,
		map[string]interface{}{
			"role_id":      role.Id,
			"version_id":   role.VersionId,
			"permissions":  pq.Array(role.Permissions),
			"extends":      pq.Array(role.Extends),
			"account_hash": sharedMock.ExistsAccountId,
		},
	)
	if err != nil {
		panic(err)
	}
}
