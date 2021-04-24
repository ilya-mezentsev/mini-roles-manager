package permission

import (
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"mini-roles-backend/source/config"
	"mini-roles-backend/source/db/connection"
	sharedMock "mini-roles-backend/source/domains/shared/mock"
	sharedModels "mini-roles-backend/source/domains/shared/models"
	"mini-roles-backend/source/domains/shared/services/hash"
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

func TestRepository_AddResourcePermissionsSuccess(t *testing.T) {
	defer sharedMock.MustReinstall(db)
	sharedMock.MustAddResource(db)
	permissions := []sharedModels.Permission{
		{
			Id:        sharedModels.PermissionId(hash.Md5WithTimeAsKey("some-permission-id-1")),
			Operation: "create",
			Effect:    "permit",
		},
		{
			Id:        sharedModels.PermissionId(hash.Md5WithTimeAsKey("some-permission-id-2")),
			Operation: "delete",
			Effect:    "deny",
		},
	}

	err := repository.AddResourcePermissions(
		sharedMock.ExistsAccountId,
		sharedMock.ExistsResourceId,
		permissions,
	)
	assert.Nil(t, err)

	for _, permission := range permissions {
		var permissionExists bool
		err = db.Get(
			&permissionExists,
			`select 1 from resource_permission where account_hash = $1 and resource_id = $2 and permission_id = $3`,
			sharedMock.ExistsAccountId,
			sharedMock.ExistsResourceId,
			permission.Id,
		)

		assert.Nil(t, err)
		assert.True(t, permissionExists)
	}
}

func TestRepository_AddResourcePermissionsErrorNoTable(t *testing.T) {
	defer sharedMock.MustReinstall(db)
	sharedMock.MustDropResourcePermissionTable(db)

	permissions := []sharedModels.Permission{
		{
			Id:        "some-permission-id-1",
			Operation: "create",
			Effect:    "permit",
		},
		{
			Id:        "some-permission-id-2",
			Operation: "delete",
			Effect:    "deny",
		},
	}

	err := repository.AddResourcePermissions(
		sharedMock.ExistsAccountId,
		sharedMock.ExistsResourceId,
		permissions,
	)
	assert.NotNil(t, err)
}

func TestRepository_AddResourcePermissionsNoConnection(t *testing.T) {
	defer func() {
		db = sqlx.MustOpen(
			"postgres",
			connection.BuildPostgresString(config.Default()),
		)
		sharedMock.MustReinstall(db)
	}()

	_ = db.Close()

	permissions := []sharedModels.Permission{
		{
			Id:        "some-permission-id-1",
			Operation: "create",
			Effect:    "permit",
		},
		{
			Id:        "some-permission-id-2",
			Operation: "delete",
			Effect:    "deny",
		},
	}

	err := repository.AddResourcePermissions(
		sharedMock.ExistsAccountId,
		sharedMock.ExistsResourceId,
		permissions,
	)
	assert.NotNil(t, err)
}
