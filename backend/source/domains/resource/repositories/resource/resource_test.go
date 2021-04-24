package resource

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"mini-roles-backend/source/config"
	"mini-roles-backend/source/db/connection"
	"mini-roles-backend/source/domains/permission/mock"
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
	someResource := sharedModels.Resource{
		Id:      "some-resource-id-1",
		Title:   "Some-Resource-Title",
		LinksTo: nil,
	}
	_, err := db.NamedExec(
		`insert into resource(account_hash, resource_id, title, links_to) values(:account_hash, :resource_id, :title, :links_to)`,
		repository.mapFromResource(sharedMock.ExistsAccountId, someResource),
	)
	assert.Nil(t, err)

	resources, err := repository.List(sharedMock.ExistsAccountId)

	assert.Nil(t, err)
	assert.Contains(t, resources, someResource)
}

func TestRepository_ListSuccessWithPermissions(t *testing.T) {
	defer sharedMock.MustReinstall(db)
	someResource := sharedModels.Resource{
		Id:      "some-resource-id-1",
		Title:   "Some-Resource-Title",
		LinksTo: nil,
		Permissions: []sharedModels.Permission{
			{
				Id:        "permission-id-1",
				Operation: "create",
				Effect:    "deny",
			},
			{
				Id:        "permission-id-2",
				Operation: "read",
				Effect:    "permit",
			},
		},
	}
	_, err := db.NamedExec(
		`insert into resource(account_hash, resource_id, title, links_to) values(:account_hash, :resource_id, :title, :links_to)`,
		repository.mapFromResource(sharedMock.ExistsAccountId, someResource),
	)
	assert.Nil(t, err)
	for _, permission := range someResource.Permissions {
		_, err = db.NamedExec(
			`
			insert into resource_permission(account_hash, resource_id, permission_id, operation, effect)
			values(:account_hash, :resource_id, :permission_id, :operation, :effect)`,
			map[string]interface{}{
				"account_hash":  sharedMock.ExistsAccountId,
				"resource_id":   someResource.Id,
				"permission_id": permission.Id,
				"operation":     permission.Operation,
				"effect":        permission.Effect,
			},
		)
		assert.Nil(t, err)
	}

	resources, err := repository.List(sharedMock.ExistsAccountId)

	assert.Nil(t, err)
	assert.Contains(t, resources, someResource)
}

func TestRepository_ListNoResourceTable(t *testing.T) {
	defer sharedMock.MustReinstall(db)
	sharedMock.MustDropResourceTable(db)

	_, err := repository.List(sharedMock.ExistsAccountId)

	assert.NotNil(t, err)
}

func TestRepository_CreateSuccess(t *testing.T) {
	defer sharedMock.MustReinstall(db)
	someResource := sharedModels.Resource{
		Id:      "some-resource-id-1",
		Title:   "Some-Resource-Title",
		LinksTo: nil,
	}

	err := repository.Create(sharedMock.ExistsAccountId, someResource)
	assert.Nil(t, err)

	var resourceExists bool
	_ = db.Get(
		&resourceExists,
		`select 1 from resource where account_hash = $1 and resource_id = $2`,
		sharedMock.ExistsAccountId,
		someResource.Id,
	)
	assert.True(t, resourceExists)
}

func TestRepository_CreateDuplicateResourceId(t *testing.T) {
	defer sharedMock.MustReinstall(db)
	someResource := sharedModels.Resource{
		Id:      "some-resource-id-1",
		Title:   "Some-Resource-Title",
		LinksTo: nil,
	}
	_, err := db.NamedExec(
		`insert into resource(account_hash, resource_id, title, links_to) values(:account_hash, :resource_id, :title, :links_to)`,
		repository.mapFromResource(sharedMock.ExistsAccountId, someResource),
	)
	assert.Nil(t, err)

	err = repository.Create(sharedMock.ExistsAccountId, someResource)
	assert.True(t, errors.As(err, &sharedError.DuplicateUniqueKey{}))
}

func TestRepository_UpdateSuccess(t *testing.T) {
	defer sharedMock.MustReinstall(db)
	someResource := sharedModels.Resource{
		Id:      "some-resource-id-1",
		Title:   "Some-Resource-Title",
		LinksTo: nil,
	}
	_, err := db.NamedExec(
		`insert into resource(account_hash, resource_id, title, links_to) values(:account_hash, :resource_id, :title, :links_to)`,
		repository.mapFromResource(sharedMock.ExistsAccountId, someResource),
	)
	assert.Nil(t, err)

	someResource.LinksTo = append(someResource.LinksTo, "some-resource-id-2")
	someResource.Title = "Some-New-Title"

	err = repository.Update(sharedMock.ExistsAccountId, someResource)
	assert.Nil(t, err)

	resources, err := repository.List(sharedMock.ExistsAccountId)
	assert.Nil(t, err)
	assert.Contains(t, resources, someResource)
}

func TestRepository_DeleteSuccess(t *testing.T) {
	defer sharedMock.MustReinstall(db)
	someResource := sharedModels.Resource{
		Id:      "some-resource-id-1",
		Title:   "Some-Resource-Title",
		LinksTo: nil,
	}
	_, err := db.NamedExec(
		`insert into resource(account_hash, resource_id, title, links_to) values(:account_hash, :resource_id, :title, :links_to)`,
		repository.mapFromResource(sharedMock.ExistsAccountId, someResource),
	)
	assert.Nil(t, err)

	err = repository.Delete(sharedMock.ExistsAccountId, someResource.Id)
	assert.Nil(t, err)

	var resourceExists bool
	_ = db.Get(
		&resourceExists,
		`select 1 from resource where account_hash = $1 and resource_id = $2`,
		sharedMock.ExistsAccountId,
		someResource.Id,
	)
	assert.False(t, resourceExists)
}

func TestRepository_DeleteSuccessFilterRolePermissions(t *testing.T) {
	defer sharedMock.MustReinstall(db)
	someResource := sharedModels.Resource{
		Id:      "some-resource-id-1",
		Title:   "Some-Resource-Title",
		LinksTo: nil,
	}
	_, err := db.NamedExec(
		`insert into resource(account_hash, resource_id, title, links_to) values(:account_hash, :resource_id, :title, :links_to)`,
		repository.mapFromResource(sharedMock.ExistsAccountId, someResource),
	)
	assert.Nil(t, err)
	for _, permission := range mock.MakeRole1Permissions() {
		db.MustExec(
			`insert into resource_permission(permission_id, operation, effect, resource_id, account_hash) values($1, $2, $3, $4, $5)`,
			permission.Id,
			mock.PermittedOperation,
			"permit",
			someResource.Id,
			sharedMock.ExistsAccountId,
		)
	}
	sharedMock.MustAddRole(db, mock.FlatRoles[0])

	err = repository.Delete(sharedMock.ExistsAccountId, someResource.Id)
	assert.Nil(t, err)

	var resourceExists bool
	_ = db.Get(
		&resourceExists,
		`select 1 from resource where account_hash = $1 and resource_id = $2`,
		sharedMock.ExistsAccountId,
		someResource.Id,
	)
	assert.False(t, resourceExists)

	var rolePermissions []sharedModels.PermissionId
	_ = db.Select(
		&rolePermissions,
		`select trim(permission_id) permission_id from role_permission where role_id = $1`,
		mock.FlatRoleId1,
	)
	for _, permission := range mock.MakeRole1Permissions() {
		assert.NotContains(t, rolePermissions, permission.Id)
	}
}
