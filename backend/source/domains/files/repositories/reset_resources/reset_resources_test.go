package reset_resources

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"mini-roles-backend/source/config"
	"mini-roles-backend/source/db/connection"
	"mini-roles-backend/source/domains/resource/repositories/resource"
	sharedMock "mini-roles-backend/source/domains/shared/mock"
	sharedModels "mini-roles-backend/source/domains/shared/models"
	sharedResource "mini-roles-backend/source/domains/shared/resource"
	"mini-roles-backend/source/domains/shared/services/hash"
	"mini-roles-backend/source/domains/shared/spec"
	"testing"
)

var (
	db                 *sqlx.DB
	repository         Repository
	resourceRepository resource.Repository
)

func init() {
	db = sqlx.MustOpen(
		"postgres",
		connection.BuildPostgresString(config.Default()),
	)
	repository = New(db)
	resourceRepository = resource.New(db)

	sharedMock.MustReinstall(db)
}

func TestRepository_ResetEmptySuccess(t *testing.T) {
	defer sharedMock.MustReinstall(db)

	resources, _ := resourceRepository.List(spec.AccountWithId{
		AccountId: sharedMock.ExistsAccountId,
	})
	assert.Empty(t, resources)

	newResources := makeResources()
	err := repository.Reset(sharedMock.ExistsAccountId, newResources)
	assert.Nil(t, err)

	resources, _ = resourceRepository.List(spec.AccountWithId{
		AccountId: sharedMock.ExistsAccountId,
	})
	assert.Equal(t, newResources, resources)
}

func TestRepository_ResetReplaceSuccess(t *testing.T) {
	defer sharedMock.MustReinstall(db)

	// fill resource table with first version of generated resources
	newResources1 := makeResources()
	err := repository.Reset(sharedMock.ExistsAccountId, newResources1)
	assert.Nil(t, err)

	resources, _ := resourceRepository.List(spec.AccountWithId{
		AccountId: sharedMock.ExistsAccountId,
	})
	assert.Equal(t, newResources1, resources)

	// generate new version of resources (not equal to first one)
	newResources2 := makeResources()
	assert.NotEqual(t, newResources1, newResources2)

	// replace old version with new
	err = repository.Reset(sharedMock.ExistsAccountId, newResources2)
	assert.Nil(t, err)

	// check if permission of old resources is deleted
	var oldPermissionsExists bool
	err = db.Get(
		&oldPermissionsExists,
		`select 1 from permission where permission_id in ($1)`,
		pq.Array(makePermissionsIds(newResources1)),
	)
	assert.Equal(t, err, sql.ErrNoRows)
	assert.False(t, oldPermissionsExists)

	resources, _ = resourceRepository.List(spec.AccountWithId{
		AccountId: sharedMock.ExistsAccountId,
	})
	assert.Equal(t, newResources2, resources)
}

func TestRepository_ResetDuplicateResourceId(t *testing.T) {
	defer sharedMock.MustReinstall(db)

	newResources := makeResources()
	newResources[0].Id = newResources[1].Id

	err := repository.Reset(sharedMock.ExistsAccountId, newResources)
	assert.NotNil(t, err)
}

func TestRepository_ResetDuplicatePermissionId(t *testing.T) {
	defer sharedMock.MustReinstall(db)

	newResources := makeResources()
	newResources[0].Permissions[0].Id = newResources[0].Permissions[1].Id

	err := repository.Reset(sharedMock.ExistsAccountId, newResources)
	assert.NotNil(t, err)
}

func TestRepository_ResetNotTableError(t *testing.T) {
	defer sharedMock.MustReinstall(db)
	sharedMock.MustDropResourceTable(db)

	err := repository.Reset(sharedMock.ExistsAccountId, makeResources())
	assert.NotNil(t, err)
}

func makeResources() []sharedModels.Resource {
	var resources []sharedModels.Resource
	for i := 1; i < 4; i++ {
		var permissions []sharedModels.Permission
		for _, operation := range []string{
			sharedResource.CreateOperation,
			sharedResource.ReadOperation,
			sharedResource.UpdateOperation,
			sharedResource.DeleteOperation,
		} {
			for _, effect := range []string{
				sharedResource.DenyEffect,
				sharedResource.PermitEffect,
			} {
				permissions = append(permissions, sharedModels.Permission{
					Id:        sharedModels.PermissionId(hash.Md5WithTimeAsKey(effect + operation)),
					Operation: operation,
					Effect:    effect,
				})
			}
		}

		resources = append(resources, sharedModels.Resource{
			Id:          sharedModels.ResourceId(fmt.Sprintf("test-resource-%d", i)),
			Permissions: permissions,
		})
	}

	return resources
}

func makePermissionsIds(resources []sharedModels.Resource) []sharedModels.PermissionId {
	var permissionsIds []sharedModels.PermissionId
	for _, _resource := range resources {
		for _, permission := range _resource.Permissions {
			permissionsIds = append(permissionsIds, permission.Id)
		}
	}

	return permissionsIds
}
