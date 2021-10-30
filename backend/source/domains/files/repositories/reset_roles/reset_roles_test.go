package reset_roles

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"mini-roles-backend/source/config"
	"mini-roles-backend/source/db/connection"
	"mini-roles-backend/source/domains/role/repositories/role"
	sharedMock "mini-roles-backend/source/domains/shared/mock"
	sharedModels "mini-roles-backend/source/domains/shared/models"
	"mini-roles-backend/source/domains/shared/services/hash"
	"mini-roles-backend/source/domains/shared/spec"
	"testing"
)

var (
	db             *sqlx.DB
	repository     Repository
	roleRepository role.Repository
)

func init() {
	db = sqlx.MustOpen(
		"postgres",
		connection.BuildPostgresString(config.Default()),
	)
	repository = New(db)
	roleRepository = role.New(db)

	sharedMock.MustReinstall(db)
}

func TestRepository_ResetEmptySuccess(t *testing.T) {
	defer sharedMock.MustReinstall(db)

	roles, err := roleRepository.List(spec.AccountWithId{
		AccountId: sharedMock.ExistsAccountId,
	})
	assert.Nil(t, err)
	assert.Empty(t, roles)

	newRoles := makeRoles()
	err = repository.Reset(sharedMock.ExistsAccountId, newRoles)
	assert.Nil(t, err)

	roles, err = roleRepository.List(spec.AccountWithId{
		AccountId: sharedMock.ExistsAccountId,
	})
	assert.Nil(t, err)
	assert.Equal(t, newRoles, roles)
}

func TestRepository_ResetReplaceSuccess(t *testing.T) {
	defer sharedMock.MustReinstall(db)

	// fill resource table with first version of generated roles
	newRoles1 := makeRoles()
	err := repository.Reset(sharedMock.ExistsAccountId, newRoles1)
	assert.Nil(t, err)

	roles, err := roleRepository.List(spec.AccountWithId{
		AccountId: sharedMock.ExistsAccountId,
	})
	assert.Nil(t, err)
	assert.Equal(t, newRoles1, roles)

	// generate new version of roles (not equal to first one)
	newRoles2 := makeRoles()
	assert.NotEqual(t, newRoles2, newRoles1)

	// replace old version with new
	err = repository.Reset(sharedMock.ExistsAccountId, newRoles2)
	assert.Nil(t, err)

	roles, err = roleRepository.List(spec.AccountWithId{
		AccountId: sharedMock.ExistsAccountId,
	})
	assert.Nil(t, err)
	assert.Equal(t, newRoles2, roles)
}

func TestRepository_ResetDuplicateRoleId(t *testing.T) {
	defer sharedMock.MustReinstall(db)

	newRoles := makeRoles()
	newRoles[0].Id = newRoles[1].Id

	err := repository.Reset(sharedMock.ExistsAccountId, newRoles)
	assert.NotNil(t, err)
}

func TestRepository_ResetNoRoleTable(t *testing.T) {
	defer sharedMock.MustReinstall(db)
	sharedMock.MustDropRoleTable(db)

	err := repository.Reset(sharedMock.ExistsAccountId, makeRoles())
	assert.NotNil(t, err)
}

func makeRoles() []sharedModels.Role {
	var roles []sharedModels.Role
	for i := 0; i < 4; i++ {
		var permissionsIds []sharedModels.PermissionId
		for j := 0; j < 5; j++ {
			permissionsIds = append(
				permissionsIds,
				sharedModels.PermissionId(hash.Md5WithTimeAsKey("test-permission")),
			)
		}

		roles = append(roles, sharedModels.Role{
			Id:          sharedModels.RoleId(fmt.Sprintf("test-role-%d", i)),
			Permissions: permissionsIds,
			VersionId:   sharedMock.ExistsRolesVersionId,
		})
	}

	return roles
}
