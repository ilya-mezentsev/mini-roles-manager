package resource

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"mini-roles-backend/source/config"
	"mini-roles-backend/source/db/connection"
	"mini-roles-backend/source/domains/permission/mock"
	sharedError "mini-roles-backend/source/domains/shared/error"
	sharedMock "mini-roles-backend/source/domains/shared/mock"
	sharedModels "mini-roles-backend/source/domains/shared/models"
	sharedResource "mini-roles-backend/source/domains/shared/resource"
	sharedSpec "mini-roles-backend/source/domains/shared/spec"
	"strings"
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

	resources, err := repository.List(sharedSpec.AccountWithId{
		AccountId: sharedMock.ExistsAccountId,
	})

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
				Operation: sharedResource.CreateOperation,
				Effect:    sharedResource.DenyEffect,
			},
			{
				Id:        "permission-id-2",
				Operation: sharedResource.ReadOperation,
				Effect:    sharedResource.PermitEffect,
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
			insert into permission(account_hash, resource_id, permission_id, operation, effect)
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

	resources, err := repository.List(sharedSpec.AccountWithId{
		AccountId: sharedMock.ExistsAccountId,
	})

	assert.Nil(t, err)
	assert.Contains(t, resources, someResource)
}

func TestRepository_ListEmpty(t *testing.T) {
	resources, err := repository.List(sharedSpec.AccountWithId{
		AccountId: sharedMock.ExistsAccountId,
	})

	assert.Nil(t, err)
	assert.Empty(t, resources)
}

func TestRepository_ListNoResourceTable(t *testing.T) {
	defer sharedMock.MustReinstall(db)
	sharedMock.MustDropResourceTable(db)

	_, err := repository.List(sharedSpec.AccountWithId{
		AccountId: sharedMock.ExistsAccountId,
	})

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

	resources, err := repository.List(sharedSpec.AccountWithId{
		AccountId: sharedMock.ExistsAccountId,
	})
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
	db.MustExec(
		`insert into permission(permission_id, operation, effect, resource_id, account_hash) values($1, $2, $3, $4, $5)`,
		mock.PermitReadPermissionId1,
		mock.PermittedOperation,
		sharedResource.PermitEffect,
		someResource.Id,
		sharedMock.ExistsAccountId,
	)
	db.MustExec(
		`insert into role(role_id, permissions, account_hash) values($1, $2, $3)`,
		mock.FlatRoleId1,
		pq.Array([]string{
			mock.PermitReadPermissionId1,
			mock.PermitDeletePermissionId3,
		}),
		sharedMock.ExistsAccountId,
	)

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

	var rolePermissionsRow pq.StringArray
	_ = db.Get(&rolePermissionsRow, `select permissions from role where role_id = $1`, mock.FlatRoleId1)
	var rolePermissions []string
	for _, row := range rolePermissionsRow {
		rolePermissions = append(rolePermissions, strings.TrimSpace(row))
	}
	assert.Contains(t, rolePermissions, mock.PermitDeletePermissionId3)
	assert.NotContains(t, rolePermissions, mock.PermitReadPermissionId1)
}
