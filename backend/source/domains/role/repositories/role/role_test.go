package role

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"mini-roles-backend/source/config"
	"mini-roles-backend/source/db/connection"
	sharedError "mini-roles-backend/source/domains/shared/error"
	sharedMock "mini-roles-backend/source/domains/shared/mock"
	sharedModels "mini-roles-backend/source/domains/shared/models"
	sharedRoleRepository "mini-roles-backend/source/domains/shared/repositories/role"
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
	someRole := sharedModels.Role{
		Id:        "super-user",
		VersionId: sharedMock.ExistsRolesVersionId,
		Title:     "Damn Super User",
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
		`insert into role(role_id, version_id, title, permissions, extends, account_hash) values(:role_id, :version_id, :title, :permissions, :extends, :account_hash)`,
		sharedRoleRepository.MapFromRole(sharedMock.ExistsAccountId, someRole),
	)
	assert.Nil(t, err)

	roles, err := repository.List(sharedSpec.AccountWithId{
		AccountId: sharedMock.ExistsAccountId,
	})

	assert.Nil(t, err)
	assert.Contains(t, roles, someRole)
}

func TestRepository_ListSuccessEmpty(t *testing.T) {
	roles, err := repository.List(sharedSpec.AccountWithId{
		AccountId: sharedMock.ExistsAccountId,
	})

	assert.Nil(t, err)
	assert.Empty(t, roles)
}

func TestRepository_CreateSuccess(t *testing.T) {
	defer sharedMock.MustReinstall(db)
	someRole := sharedModels.Role{
		Id:        "super-user",
		VersionId: sharedMock.ExistsRolesVersionId,
		Title:     "Damn Super User",
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
		Id:        "super-user",
		VersionId: sharedMock.ExistsRolesVersionId,
		Title:     "Damn Super User",
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
		`insert into role(role_id, version_id, title, permissions, extends, account_hash) values(:role_id, :version_id, :title, :permissions, :extends, :account_hash)`,
		sharedRoleRepository.MapFromRole(sharedMock.ExistsAccountId, someRole),
	)
	assert.Nil(t, err)

	err = repository.Create(sharedMock.ExistsAccountId, someRole)

	assert.True(t, errors.As(err, &sharedError.DuplicateUniqueKey{}))
}

func TestRepository_CreateDuplicateRoleIdButAnotherAccount(t *testing.T) {
	defer sharedMock.MustReinstall(db)
	someRole := sharedModels.Role{
		Id:        "super-user",
		VersionId: sharedMock.ExistsRolesVersionId,
		Title:     "Damn Super User",
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
		sharedRoleRepository.MapFromRole("some-account-id", someRole),
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
		Id:        "super-user",
		VersionId: sharedMock.ExistsRolesVersionId,
		Title:     "Damn Super User",
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

	roles, _ := repository.List(sharedSpec.AccountWithId{
		AccountId: sharedMock.ExistsAccountId,
	})
	assert.Contains(t, roles, someRole)
}

func TestRepository_DeleteSuccess(t *testing.T) {
	sharedMock.MustAddRolesVersion(db, sharedMock.ExistsRolesVersionId2)
	defer sharedMock.MustReinstall(db)
	someRole := sharedModels.Role{
		Id:        "super-user",
		VersionId: sharedMock.ExistsRolesVersionId,
		Title:     "Damn Super User",
		Permissions: []sharedModels.PermissionId{
			"permission-id-1",
			"permission-id-2",
		},
		Extends: []sharedModels.RoleId{
			"user",
			"guest",
		},
	}
	roleWithAnotherVersionId := sharedModels.Role{
		Id:        someRole.Id,
		VersionId: sharedMock.ExistsRolesVersionId2,
	}

	assert.Nil(t, repository.Create(sharedMock.ExistsAccountId, someRole))
	assert.Nil(t, repository.Create(sharedMock.ExistsAccountId, roleWithAnotherVersionId))

	err := repository.Delete(sharedMock.ExistsAccountId, sharedMock.ExistsRolesVersionId, someRole.Id)
	assert.Nil(t, err)

	var roleExists bool
	_ = db.Get(
		&roleExists,
		`select 1 from role where role_id = $1 and account_hash = $2 and version_id = $3`,
		someRole.Id,
		sharedMock.ExistsAccountId,
		someRole.VersionId,
	)
	assert.False(t, roleExists)

	var anotherRolesExists bool
	_ = db.Get(
		&anotherRolesExists,
		`select 1 from role where role_id = $1 and account_hash = $2 and version_id = $3`,
		roleWithAnotherVersionId.Id,
		sharedMock.ExistsAccountId,
		roleWithAnotherVersionId.VersionId,
	)
	assert.True(t, anotherRolesExists)
}

func TestRepository_DeleteFromExtendsSuccess(t *testing.T) {
	sharedMock.MustAddRolesVersion(db, sharedMock.ExistsRolesVersionId2)
	defer sharedMock.MustReinstall(db)
	user := sharedModels.Role{
		Id:        "user",
		VersionId: sharedMock.ExistsRolesVersionId,
		Title:     "Damn User",
		Permissions: []sharedModels.PermissionId{
			"permission-id-3",
		},
	}
	superUser := sharedModels.Role{
		Id:        "super-user",
		VersionId: sharedMock.ExistsRolesVersionId,
		Title:     "Damn Super User",
		Permissions: []sharedModels.PermissionId{
			"permission-id-1",
			"permission-id-2",
		},
		Extends: []sharedModels.RoleId{
			user.Id,
			"guest",
		},
	}
	roleWithAnotherVersionId := sharedModels.Role{
		Id:        "some-user",
		VersionId: sharedMock.ExistsRolesVersionId2,
		Extends: []sharedModels.RoleId{
			user.Id,
		},
	}

	assert.Nil(t, repository.Create(sharedMock.ExistsAccountId, user))
	assert.Nil(t, repository.Create(sharedMock.ExistsAccountId, superUser))
	assert.Nil(t, repository.Create(sharedMock.ExistsAccountId, roleWithAnotherVersionId))

	err := repository.Delete(sharedMock.ExistsAccountId, sharedMock.ExistsRolesVersionId, user.Id)
	assert.Nil(t, err)

	var superUserExtends pq.StringArray
	_ = db.Get(&superUserExtends, `select extends from role where role_id = $1`, superUser.Id)
	for id, extends := range superUserExtends {
		superUserExtends[id] = strings.TrimSpace(extends)
	}
	assert.Contains(t, superUserExtends, "guest")
	assert.NotContains(t, superUserExtends, "user")

	var roleWithAnotherVersionIdExtends pq.StringArray
	_ = db.Get(&roleWithAnotherVersionIdExtends, `select extends from role where role_id = $1`, roleWithAnotherVersionId.Id)
	for id, extends := range roleWithAnotherVersionIdExtends {
		roleWithAnotherVersionIdExtends[id] = strings.TrimSpace(extends)
	}
	assert.Contains(t, roleWithAnotherVersionIdExtends, string(user.Id))
}

func TestRepository_DeleteError(t *testing.T) {
	sharedMock.MustDropRoleTable(db)
	defer sharedMock.MustReinstall(db)

	err := repository.Delete(sharedMock.ExistsAccountId, sharedMock.ExistsRolesVersionId, "user")
	assert.NotNil(t, err)
}
