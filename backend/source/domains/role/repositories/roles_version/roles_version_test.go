package roles_version

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"mini-roles-backend/source/config"
	"mini-roles-backend/source/db/connection"
	sharedError "mini-roles-backend/source/domains/shared/error"
	sharedMock "mini-roles-backend/source/domains/shared/mock"
	sharedModels "mini-roles-backend/source/domains/shared/models"
	sharedSpec "mini-roles-backend/source/domains/shared/spec"
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
	rolesVersions, err := repository.List(sharedSpec.AccountWithId{
		AccountId: sharedMock.ExistsAccountId,
	})

	assert.Nil(t, err)
	assert.Contains(t, rolesVersions, sharedModels.RolesVersion{
		Id: sharedMock.ExistsRolesVersionId,
	})
}

func TestRepository_CreateSuccess(t *testing.T) {
	defer sharedMock.MustReinstall(db)

	newRolesVersion := sharedModels.RolesVersion{
		Id:    "Some-Id",
		Title: "Some-Title",
	}

	err := repository.Create(sharedMock.ExistsAccountId, newRolesVersion)
	assert.Nil(t, err)

	rolesVersions, _ := repository.List(sharedSpec.AccountWithId{
		AccountId: sharedMock.ExistsAccountId,
	})
	assert.Contains(t, rolesVersions, newRolesVersion)
}

func TestRepository_CreateDuplicateKeyError(t *testing.T) {
	newRolesVersion := sharedModels.RolesVersion{
		Id:    sharedMock.ExistsRolesVersionId,
		Title: "Some-Title",
	}

	err := repository.Create(sharedMock.ExistsAccountId, newRolesVersion)
	assert.True(t, errors.As(err, &sharedError.DuplicateUniqueKey{}))

	rolesVersions, _ := repository.List(sharedSpec.AccountWithId{
		AccountId: sharedMock.ExistsAccountId,
	})
	assert.NotContains(t, rolesVersions, newRolesVersion)
}

func TestRepository_UpdateSuccess(t *testing.T) {
	defer sharedMock.MustReinstall(db)

	updatingRolesVersion := sharedModels.RolesVersion{
		Id:    sharedMock.ExistsRolesVersionId,
		Title: "Some-Title",
	}

	err := repository.Update(sharedMock.ExistsAccountId, updatingRolesVersion)
	assert.Nil(t, err)

	rolesVersions, _ := repository.List(sharedSpec.AccountWithId{
		AccountId: sharedMock.ExistsAccountId,
	})
	assert.Contains(t, rolesVersions, updatingRolesVersion)
}

func TestRepository_DeleteSuccess(t *testing.T) {
	defer sharedMock.MustReinstall(db)

	err := repository.Delete(sharedMock.ExistsAccountId, sharedMock.ExistsRolesVersionId)
	assert.Nil(t, err)

	rolesVersions, _ := repository.List(sharedSpec.AccountWithId{
		AccountId: sharedMock.ExistsAccountId,
	})
	assert.Empty(t, rolesVersions)
}
