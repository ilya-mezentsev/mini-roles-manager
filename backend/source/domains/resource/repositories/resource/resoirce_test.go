package resource

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

func TestRepository_ListEmpty(t *testing.T) {
	resources, err := repository.List(sharedMock.ExistsAccountId)

	assert.Nil(t, err)
	assert.Empty(t, resources)
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
