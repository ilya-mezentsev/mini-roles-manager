package permission

import (
	"mini-roles-backend/source/domains/permission/interfaces"
	"mini-roles-backend/source/domains/permission/spec"
	sharedModels "mini-roles-backend/source/domains/shared/models"
	"sync"
	"time"
)

type (
	CachedPermissions struct {
		permissions []sharedModels.Permission
		cachedTime  time.Time
	}

	Cache struct {
		sync.Mutex
		cachedPermissions    map[spec.PermissionWithAccountIdAndRoleId]CachedPermissions
		cacheLifetimeSeconds uint
		repository           interfaces.PermissionRepository
	}
)

func New(
	cacheLifetimeSeconds uint,
	repository interfaces.PermissionRepository,
) *Cache {
	c := &Cache{
		repository:           repository,
		cacheLifetimeSeconds: cacheLifetimeSeconds,
	}
	c.cachedPermissions = c.emptyCacheMap()

	return c
}

func (c *Cache) List(spec spec.PermissionWithAccountIdAndRoleId) ([]sharedModels.Permission, error) {
	c.Lock()
	defer c.Unlock()

	cachedPermissions, found := c.cachedPermissions[spec]
	if found && !c.isExpired(cachedPermissions.cachedTime) {
		return cachedPermissions.permissions, nil
	}

	permissions, err := c.repository.List(spec)
	if err != nil {
		return nil, err
	}

	c.cachedPermissions[spec] = CachedPermissions{
		permissions: permissions,
		cachedTime:  time.Now(),
	}

	return permissions, nil
}

func (c *Cache) isExpired(cachedTime time.Time) bool {
	expirationTime := cachedTime.Add(time.Second * time.Duration(c.cacheLifetimeSeconds))

	return expirationTime.Before(time.Now())
}

func (c *Cache) Invalidate(accountId sharedModels.AccountId) {
	c.Lock()
	defer c.Unlock()

	newCachedPermissions := c.emptyCacheMap()
	for cachedSpec, permissions := range c.cachedPermissions {
		if cachedSpec.AccountId != accountId {
			newCachedPermissions[cachedSpec] = permissions
		}
	}

	c.cachedPermissions = newCachedPermissions
}

func (c *Cache) emptyCacheMap() map[spec.PermissionWithAccountIdAndRoleId]CachedPermissions {
	return map[spec.PermissionWithAccountIdAndRoleId]CachedPermissions{}
}
