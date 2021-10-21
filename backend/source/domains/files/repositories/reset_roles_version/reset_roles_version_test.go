package reset_roles_version

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"mini-roles-backend/source/config"
	"mini-roles-backend/source/db/connection"
	"mini-roles-backend/source/domains/permission/repositories/db/default_roles_version"
	"mini-roles-backend/source/domains/role/repositories/roles_version"
	sharedMock "mini-roles-backend/source/domains/shared/mock"
	sharedModels "mini-roles-backend/source/domains/shared/models"
	"mini-roles-backend/source/domains/shared/spec"
	"testing"
)

var (
	db                            *sqlx.DB
	repository                    Repository
	roleVersionRepository         roles_version.Repository
	defaultRolesVersionRepository default_roles_version.Repository
)

func init() {
	db = sqlx.MustOpen(
		"postgres",
		connection.BuildPostgresString(config.Default()),
	)
	repository = New(db)
	roleVersionRepository = roles_version.New(db)
	defaultRolesVersionRepository = default_roles_version.New(db)

	sharedMock.MustReinstall(db)
}

func TestRepository_ResetSuccess(t *testing.T) {
	defer sharedMock.MustReinstall(db)

	newRolesVersions, defaultRolesVersion := makeRolesVersions()

	rolesVersions, err := roleVersionRepository.List(spec.AccountWithId{
		AccountId: sharedMock.ExistsAccountId,
	})
	assert.Nil(t, err)
	assert.Contains(t, rolesVersions, sharedModels.RolesVersion{Id: sharedMock.ExistsRolesVersionId})

	err = repository.Reset(sharedMock.ExistsAccountId, defaultRolesVersion, newRolesVersions)
	assert.Nil(t, err)

	rolesVersions, err = roleVersionRepository.List(spec.AccountWithId{
		AccountId: sharedMock.ExistsAccountId,
	})
	assert.Nil(t, err)
	assert.NotContains(t, rolesVersions, sharedModels.RolesVersion{Id: sharedMock.ExistsRolesVersionId})
	for _, rolesVersion := range newRolesVersions {
		assert.Contains(t, rolesVersions, rolesVersion)
	}

	defaultRolesVersionFromDB, _ := defaultRolesVersionRepository.Fetch(spec.AccountWithId{
		AccountId: sharedMock.ExistsAccountId,
	})
	assert.Equal(t, defaultRolesVersion, defaultRolesVersionFromDB)
}

func TestRepository_ResetDuplicateVersionId(t *testing.T) {
	defer sharedMock.MustReinstall(db)

	newRolesVersions, defaultRolesVersion := makeRolesVersions()
	newRolesVersions[0].Id = newRolesVersions[1].Id

	err := repository.Reset(sharedMock.ExistsAccountId, defaultRolesVersion, newRolesVersions)
	assert.NotNil(t, err)
}

func TestRepository_ResetNotVersionTable(t *testing.T) {
	defer sharedMock.MustReinstall(db)
	sharedMock.MustDropRolesVersionTable(db)

	newRolesVersions, defaultRolesVersion := makeRolesVersions()

	err := repository.Reset(sharedMock.ExistsAccountId, defaultRolesVersion, newRolesVersions)
	assert.NotNil(t, err)
}

func makeRolesVersions() (rolesVersions []sharedModels.RolesVersion, defaultRolesVersion sharedModels.RolesVersion) {
	for i := 0; i < 4; i++ {
		rolesVersions = append(rolesVersions, sharedModels.RolesVersion{
			Id: sharedModels.RolesVersionId(fmt.Sprintf("test-roles-version-%d", i)),
		})
	}

	return rolesVersions, sharedModels.RolesVersion{Id: "default-roles-version"}
}
