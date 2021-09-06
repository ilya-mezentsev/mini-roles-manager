package permission

import (
	"github.com/stretchr/testify/assert"
	"mini-roles-backend/source/domains/permission/mock"
	"mini-roles-backend/source/domains/permission/spec"
	sharedMock "mini-roles-backend/source/domains/shared/mock"
	"testing"
	"time"
)

var (
	cacheLifetimeSeconds     = uint(1)
	mockPermissionRepository = &mock.PermissionRepository{}
	cache                    = New(cacheLifetimeSeconds, mockPermissionRepository)
)

func init() {
	mockPermissionRepository.Reset()
}

func TestCache_List(t *testing.T) {
	defer mockPermissionRepository.Reset()
	defer resetCache()

	s := defaultSpec()
	permissions, err := cache.List(s)

	assert.Nil(t, err)
	assert.NotEmpty(t, permissions)
	assert.Equal(t, permissions, cache.cachedPermissions[s].permissions)

	mockPermissionRepository.Clean()

	cachedPermissions, err := cache.List(s)
	assert.Nil(t, err)
	// taken from cache
	assert.Equal(t, permissions, cachedPermissions)
}

func TestCache_ListExpired(t *testing.T) {
	defer resetCache()

	cacheLifetimeSeconds_ := uint(0)
	cache_ := New(cacheLifetimeSeconds_, mockPermissionRepository)

	s := defaultSpec()
	permissions, err := cache_.List(s)

	assert.Nil(t, err)
	assert.NotEmpty(t, permissions)
	assert.Equal(t, permissions, cache_.cachedPermissions[s].permissions)

	time.Sleep(time.Second * time.Duration(cacheLifetimeSeconds_))

	assert.True(t, cache_.isExpired(cache_.cachedPermissions[s].cachedTime))
}

func TestCache_ListError(t *testing.T) {
	defer resetCache()

	s := spec.PermissionWithAccountIdAndRoleId{
		AccountId:      sharedMock.BadAccountId,
		RoleId:         sharedMock.ExistsRoleId,
		RolesVersionId: sharedMock.ExistsRolesVersionId,
	}
	_, err := cache.List(s)

	assert.NotNil(t, err)
	assert.Empty(t, cache.cachedPermissions[s])
}

func TestCache_Invalidate(t *testing.T) {
	defer resetCache()

	s := defaultSpec()
	permissions, err := cache.List(s)

	assert.Nil(t, err)
	assert.NotEmpty(t, permissions)
	assert.Equal(t, permissions, cache.cachedPermissions[s].permissions)

	s2 := defaultSpec()
	s2.AccountId = sharedMock.ExistsAccountId2
	permissions, err = cache.List(s2)

	assert.Nil(t, err)
	assert.NotEmpty(t, permissions)
	assert.Equal(t, permissions, cache.cachedPermissions[s2].permissions)

	cache.Invalidate(s.AccountId)
	assert.Empty(t, cache.cachedPermissions[s])
	assert.NotEmpty(t, cache.cachedPermissions[s2])
}

func resetCache() {
	cache.cachedPermissions = cache.emptyCacheMap()
}

func defaultSpec() spec.PermissionWithAccountIdAndRoleId {
	return spec.PermissionWithAccountIdAndRoleId{
		AccountId:      sharedMock.ExistsAccountId,
		RoleId:         sharedMock.ExistsRoleId,
		RolesVersionId: sharedMock.ExistsRolesVersionId,
	}
}
