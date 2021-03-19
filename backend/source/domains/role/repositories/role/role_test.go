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

func TestRepository_ListSuccess(t *testing.T) {
	defer sharedMock.MustReinstall(db)
	someRole := sharedModels.Role{
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
	_, err := db.NamedExec(
		`insert into role(role_id, title, permissions, extends, account_hash) values(:role_id, :title, :permissions, :extends, :account_hash)`,
		repository.mapFromRole(sharedMock.ExistsAccountId, someRole),
	)
	assert.Nil(t, err)

	roles, err := repository.List(sharedMock.ExistsAccountId)

	assert.Nil(t, err)
	assert.Contains(t, roles, someRole)
}

func TestRepository_ListSuccessEmpty(t *testing.T) {
	roles, err := repository.List(sharedMock.ExistsAccountId)

	assert.Nil(t, err)
	assert.Empty(t, roles)
}

func TestRepository_CreateSuccess(t *testing.T) {
	defer sharedMock.MustReinstall(db)
	someRole := sharedModels.Role{
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

	err := repository.Create(sharedMock.ExistsAccountId, someRole)
	assert.Nil(t, err)

	var roleCreated bool
	_ = db.Get(&roleCreated, `select 1 from role where role_id = $1 and account_hash = $2`, someRole.Id, sharedMock.ExistsAccountId)

	assert.True(t, roleCreated)
}

func TestRepository_CreateDuplicateRoleId(t *testing.T) {
	defer sharedMock.MustReinstall(db)
	someRole := sharedModels.Role{
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

	_, err := db.NamedExec(
		`insert into role(role_id, title, permissions, extends, account_hash) values(:role_id, :title, :permissions, :extends, :account_hash)`,
		repository.mapFromRole(sharedMock.ExistsAccountId, someRole),
	)
	assert.Nil(t, err)

	err = repository.Create(sharedMock.ExistsAccountId, someRole)

	assert.True(t, errors.As(err, &sharedError.DuplicateUniqueKey{}))
}

func TestRepository_CreateDuplicateRoleIdButAnotherAccount(t *testing.T) {
	defer sharedMock.MustReinstall(db)
	someRole := sharedModels.Role{
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

	db.MustExec(`insert into account(hash) values($1)`, "some-account-id")
	_, err := db.NamedExec(
		`insert into role(role_id, title, permissions, extends, account_hash) values(:role_id, :title, :permissions, :extends, :account_hash)`,
		repository.mapFromRole("some-account-id", someRole),
	)
	assert.Nil(t, err)

	err = repository.Create(sharedMock.ExistsAccountId, someRole)
	assert.Nil(t, err)

	var roleCreated bool
	_ = db.Get(&roleCreated, `select 1 from role where role_id = $1 and account_hash = $2`, someRole.Id, sharedMock.ExistsAccountId)

	assert.True(t, roleCreated)
}

func TestRepository_UpdateSuccess(t *testing.T) {
	defer sharedMock.MustReinstall(db)
	someRole := sharedModels.Role{
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

	err := repository.Create(sharedMock.ExistsAccountId, someRole)
	assert.Nil(t, err)

	someRole.Title = "Some-New-Title"
	someRole.Extends = append(someRole.Extends, "preview")

	err = repository.Update(sharedMock.ExistsAccountId, someRole)
	assert.Nil(t, err)

	roles, _ := repository.List(sharedMock.ExistsAccountId)
	assert.Contains(t, roles, someRole)
}

func TestRepository_DeleteSuccess(t *testing.T) {
	defer sharedMock.MustReinstall(db)
	someRole := sharedModels.Role{
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

	err := repository.Create(sharedMock.ExistsAccountId, someRole)
	assert.Nil(t, err)

	err = repository.Delete(sharedMock.ExistsAccountId, someRole.Id)
	assert.Nil(t, err)

	var roleExists bool
	_ = db.Get(&roleExists, `select 1 from role where role_id = $1 and account_hash = $2`, someRole.Id, sharedMock.ExistsAccountId)

	assert.False(t, roleExists)
}
