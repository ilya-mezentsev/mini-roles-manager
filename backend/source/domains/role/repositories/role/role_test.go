package role

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"mini-roles-backend/source/config"
	"mini-roles-backend/source/db/connection"
	sharedError "mini-roles-backend/source/domains/shared/error"
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

func someRole() sharedModels.Role {
	return sharedModels.Role{
		Id:    "super-user",
		Title: "Damn Super User",
		Permissions: []sharedModels.PermissionId{
			"permission-id-1",
			"permission-id-2",
		},
		Extends: []sharedModels.RoleId{
			"user",
			"guest",
		},
	}
}

func someRoleWithoutExtends() sharedModels.Role {
	role := someRole()
	role.Extends = nil

	return role
}

func someRoleWithoutPermissions() sharedModels.Role {
	role := someRole()
	role.Permissions = nil

	return role
}

func TestRepository_ListSuccess(t *testing.T) {
	role := someRoleWithoutExtends()
	sharedMock.MustAddPermissions(db, role.Permissions)
	sharedMock.MustAddRole(db, role)
	defer sharedMock.MustReinstall(db)

	roles, err := repository.List(sharedMock.ExistsAccountId)

	assert.Nil(t, err)
	assert.Contains(t, roles, role)
}

func TestRepository_ListSuccessEmpty(t *testing.T) {
	roles, err := repository.List(sharedMock.ExistsAccountId)

	assert.Nil(t, err)
	assert.Empty(t, roles)
}

func TestRepository_CreateSuccess(t *testing.T) {
	role := someRoleWithoutExtends()
	sharedMock.MustAddPermissions(db, role.Permissions)
	defer sharedMock.MustReinstall(db)

	err := repository.Create(sharedMock.ExistsAccountId, role)
	assert.Nil(t, err)

	var roleCreated bool
	_ = db.Get(&roleCreated, `select 1 from role where role_id = $1 and account_hash = $2`, role.Id, sharedMock.ExistsAccountId)
	assert.True(t, roleCreated)
}

func TestRepository_CreateDuplicateRoleId(t *testing.T) {
	defer sharedMock.MustReinstall(db)

	_, err := db.NamedExec(
		`insert into role(role_id, title, account_hash) values(:role_id, :title, :account_hash)`,
		repository.mapFromRole(sharedMock.ExistsAccountId, someRole()),
	)
	assert.Nil(t, err)

	err = repository.Create(sharedMock.ExistsAccountId, someRole())

	assert.True(t, errors.As(err, &sharedError.DuplicateUniqueKey{}))
}

func TestRepository_CreateDuplicateRoleIdButAnotherAccount(t *testing.T) {
	role := someRoleWithoutExtends()
	sharedMock.MustAddPermissions(db, role.Permissions)
	defer sharedMock.MustReinstall(db)

	db.MustExec(`insert into account(hash) values($1)`, "some-account-id")
	_, err := db.NamedExec(
		`insert into role(role_id, title, account_hash) values(:role_id, :title, :account_hash)`,
		repository.mapFromRole("some-account-id", role),
	)
	assert.Nil(t, err)

	err = repository.Create(sharedMock.ExistsAccountId, role)
	assert.Nil(t, err)

	var roleCreated bool
	_ = db.Get(&roleCreated, `select 1 from role where role_id = $1 and account_hash = $2`, role.Id, sharedMock.ExistsAccountId)

	assert.True(t, roleCreated)
}

func TestRepository_CreateNoRolePermissionTable(t *testing.T) {
	sharedMock.MustDropRolePermissionTable(db)
	defer sharedMock.MustReinstall(db)

	err := repository.Create(sharedMock.ExistsAccountId, someRole())
	assert.NotNil(t, err)
}

func TestRepository_CreateNoRoleExtendingTable(t *testing.T) {
	sharedMock.MustDropRoleExtendingTable(db)
	defer sharedMock.MustReinstall(db)

	err := repository.Update(sharedMock.ExistsAccountId, someRoleWithoutPermissions())
	assert.NotNil(t, err)
}

func TestRepository_UpdateSuccess(t *testing.T) {
	role := someRoleWithoutExtends()
	role.Extends = append(role.Extends, "guest")
	sharedMock.MustAddPermissions(db, role.Permissions)
	defer sharedMock.MustReinstall(db)

	previewRole := someRoleWithoutExtends()
	previewRole.Id = "preview"
	assert.Nil(t, repository.Create(sharedMock.ExistsAccountId, previewRole))

	guestRole := someRoleWithoutExtends()
	guestRole.Id = "guest"
	assert.Nil(t, repository.Create(sharedMock.ExistsAccountId, guestRole))

	err := repository.Create(sharedMock.ExistsAccountId, role)
	assert.Nil(t, err)

	role.Title = "Some-New-Title"
	role.Extends = append(role.Extends[1:], "preview")
	role.Permissions = role.Permissions[1:]

	err = repository.Update(sharedMock.ExistsAccountId, role)
	assert.Nil(t, err)

	roles, _ := repository.List(sharedMock.ExistsAccountId)
	assert.Contains(t, roles, role)
}

func TestRepository_UpdateNoRolePermissionTable(t *testing.T) {
	sharedMock.MustDropRolePermissionTable(db)
	defer sharedMock.MustReinstall(db)

	err := repository.Update(sharedMock.ExistsAccountId, someRole())
	assert.NotNil(t, err)
}

func TestRepository_UpdateNoRoleExtendingTable(t *testing.T) {
	sharedMock.MustDropRoleExtendingTable(db)
	defer sharedMock.MustReinstall(db)

	err := repository.Update(sharedMock.ExistsAccountId, someRoleWithoutPermissions())
	assert.NotNil(t, err)
}

func TestRepository_DeleteSuccess(t *testing.T) {
	role := someRoleWithoutExtends()
	sharedMock.MustAddPermissions(db, role.Permissions)
	defer sharedMock.MustReinstall(db)

	err := repository.Create(sharedMock.ExistsAccountId, role)
	assert.Nil(t, err)

	err = repository.Delete(sharedMock.ExistsAccountId, role.Id)
	assert.Nil(t, err)

	var roleExists bool
	_ = db.Get(&roleExists, `select 1 from role where role_id = $1 and account_hash = $2`, role.Id, sharedMock.ExistsAccountId)

	assert.False(t, roleExists)
}

func TestRepository_DeleteFromExtendsSuccess(t *testing.T) {
	defer sharedMock.MustReinstall(db)
	user := sharedModels.Role{
		Id:    "user",
		Title: "Damn Super User",
		Permissions: []sharedModels.PermissionId{
			"permission-id-3",
		},
	}
	superUser := sharedModels.Role{
		Id:    "super-user",
		Title: "Damn Super User",
		Permissions: []sharedModels.PermissionId{
			"permission-id-1",
			"permission-id-2",
		},
		Extends: []sharedModels.RoleId{
			"user",
		},
	}
	sharedMock.MustAddPermissions(
		db,
		append(
			someRole().Permissions,
			append(superUser.Permissions, user.Permissions...)...,
		),
	)

	assert.Nil(t, repository.Create(sharedMock.ExistsAccountId, user))
	assert.Nil(t, repository.Create(sharedMock.ExistsAccountId, superUser))

	err := repository.Delete(sharedMock.ExistsAccountId, user.Id)
	assert.Nil(t, err)

	var superUserExtends []string
	_ = db.Select(&superUserExtends, `select trim(extends_from) extends_from from role where role_id = $1`, superUser.Id)

	assert.NotContains(t, superUserExtends, "user")
}

func TestRepository_DeleteError(t *testing.T) {
	sharedMock.MustDropRoleTable(db)
	defer sharedMock.MustReinstall(db)

	err := repository.Delete(sharedMock.ExistsAccountId, "user")
	assert.NotNil(t, err)
}
