package default_roles_version

import (
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"mini-roles-backend/source/config"
	"mini-roles-backend/source/db/connection"
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

func TestRepository_FetchFromOneVersion(t *testing.T) {
	rolesVersion, err := repository.Fetch(sharedSpec.AccountWithId{
		AccountId: sharedMock.ExistsAccountId,
	})

	assert.Nil(t, err)
	assert.Equal(t, rolesVersion, sharedModels.RolesVersion{
		Id: sharedMock.ExistsRolesVersionId,
	})
}

func TestRepository_FetchFromTwoVersions(t *testing.T) {
	defer sharedMock.MustReinstall(db)
	db.MustExec(
		`insert into roles_version(version_id, account_hash) values($1, $2)`,
		sharedMock.ExistsRolesVersionId2,
		sharedMock.ExistsAccountId,
	)

	rolesVersion, err := repository.Fetch(sharedSpec.AccountWithId{
		AccountId: sharedMock.ExistsAccountId,
	})

	assert.Nil(t, err)
	assert.Equal(t, rolesVersion, sharedModels.RolesVersion{
		Id: sharedMock.ExistsRolesVersionId,
	})
}
