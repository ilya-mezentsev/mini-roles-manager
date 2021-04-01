package permission

import (
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"mini-roles-backend/source/config"
	"mini-roles-backend/source/db/connection"
	"mini-roles-backend/source/domains/permission/mock"
	sharedMock "mini-roles-backend/source/domains/shared/mock"
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
	addTestFlatRoles()
}

func TestRepository_ListEmpty(t *testing.T) {
	permissions, err := repository.List(sharedMock.ExistsAccountId, sharedMock.ExistsRoleId)

	assert.Nil(t, err)
	assert.Empty(t, permissions)
}

func TestRepository_ListFlatRolePermissions(t *testing.T) {
	initTestData()
	defer sharedMock.MustReinstall(db)

	permissions, err := repository.List(sharedMock.ExistsAccountId, mock.FlatRoleId1)

	assert.Nil(t, err)
	for _, role1Permission := range mock.MakeRole1Permissions() {
		assert.Contains(t, permissions, role1Permission)
	}
}

func TestRepository_ListOneDepthLevelExtending(t *testing.T) {
	initTestData()
	addTestOneLevelDepthExtendingRole()
	defer sharedMock.MustReinstall(db)

	permissions, err := repository.List(sharedMock.ExistsAccountId, mock.OneDepthLevelExtendingRoleId)

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
	addTestOneLevelDepthExtendingRole()
	addTestTwoLevelDepthExtendingRole()
	defer sharedMock.MustReinstall(db)

	permissions, err := repository.List(sharedMock.ExistsAccountId, mock.TwoDepthLevelExtendingRoleId)

	assert.Nil(t, err)
	for _, extendingRolePermission := range mock.Permissions {
		assert.Contains(t, permissions, extendingRolePermission)
	}
}

func TestRepository_ListError(t *testing.T) {
	dropFunctionQuery := `
	drop function recursive_permissions(entry_point_role_id character(32), _account_hash character(32), depth int)`
	db.MustExec(dropFunctionQuery)
	db.MustExec(`
	create or replace function recursive_permissions(a character(32), b character(32), c int)
	returns void
	language plpgsql
	as $$
	    begin
	        raise log 'hello, % % %', a, b, c;
		end
	$$`)
	defer db.MustExec(dropFunctionQuery)

	_, err := repository.List(sharedMock.ExistsAccountId, sharedMock.ExistsRoleId)

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
	tx := db.MustBegin()
	for _, flatRole := range mock.FlatRoles {
		_, err := tx.NamedExec(
			`insert into role(role_id, permissions, account_hash) values(:role_id, :permissions, :account_hash)`,
			map[string]interface{}{
				"role_id":      flatRole.Id,
				"permissions":  pq.Array(flatRole.Permissions),
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

func addTestOneLevelDepthExtendingRole() {
	_, err := db.NamedExec(
		`insert into role(role_id, permissions, extends, account_hash) values(:role_id, :permissions, :extends, :account_hash)`,
		map[string]interface{}{
			"role_id":      mock.OneDepthLevelExtendingRole.Id,
			"permissions":  pq.Array(mock.OneDepthLevelExtendingRole.Permissions),
			"extends":      pq.Array(mock.OneDepthLevelExtendingRole.Extends),
			"account_hash": sharedMock.ExistsAccountId,
		},
	)
	if err != nil {
		panic(err)
	}
}

func addTestTwoLevelDepthExtendingRole() {
	_, err := db.NamedExec(
		`insert into role(role_id, extends, account_hash) values(:role_id, :extends, :account_hash)`,
		map[string]interface{}{
			"role_id":      mock.TwoDepthLevelExtendingRole.Id,
			"extends":      pq.Array(mock.TwoDepthLevelExtendingRole.Extends),
			"account_hash": sharedMock.ExistsAccountId,
		},
	)
	if err != nil {
		panic(err)
	}
}
