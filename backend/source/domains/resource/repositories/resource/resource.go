package resource

import (
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	sharedError "mini-roles-backend/source/domains/shared/error"
	sharedModels "mini-roles-backend/source/domains/shared/models"
	sharedResourceRepository "mini-roles-backend/source/domains/shared/repositories/resource"
	sharedSpec "mini-roles-backend/source/domains/shared/spec"
)

const (
	selectResourcesQuery = `
	select trim(resource_id) resource_id, trim(title) title, links_to from resource
	where account_hash = $1
	order by created_at`
	selectResourcePermissionsQuery = `
	select
		trim(permission_id) permission_id,
		trim(operation) operation,
		trim(effect) effect,
		trim(resource_id) resource_id
	from permission
	where resource_id = any($1) and account_hash = $2`

	updateResourceQuery = `
	update resource
	set title = :title, links_to = :links_to
	where account_hash = :account_hash and resource_id = :resource_id`

	deleteResourceQuery       = `delete from resource where account_hash = $1 and resource_id = $2`
	removeResourcePermissions = `
	update role set permissions = array(
		select unnest(permissions) from role
		except select permission_id from permission where account_hash = $1 and resource_id = $2
	)
	where account_hash = $1`
)

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repository {
	return Repository{db}
}

func (r Repository) Create(accountId sharedModels.AccountId, resource sharedModels.Resource) error {
	_, err := r.db.NamedExec(sharedResourceRepository.CreateResourceQuery, sharedResourceRepository.MapFromResource(accountId, resource))
	if sharedError.IsDuplicateKey(err) {
		err = sharedError.DuplicateUniqueKey{}
	}

	return err
}

func (r Repository) List(spec sharedSpec.AccountWithId) ([]sharedModels.Resource, error) {
	var resources []resourceProxy
	err := r.db.Select(&resources, selectResourcesQuery, spec.AccountId)
	if err != nil {
		return nil, err
	}

	var permissions []permissionProxy
	err = r.db.Select(
		&permissions,
		selectResourcePermissionsQuery,
		pq.Array(r.makeResourceIds(resources)),
		spec.AccountId,
	)

	return r.makeResources(resources, permissions), err
}

func (r Repository) makeResourceIds(resourcesProxies []resourceProxy) []sharedModels.ResourceId {
	var ids []sharedModels.ResourceId
	for _, proxy := range resourcesProxies {
		ids = append(ids, proxy.Id)
	}

	return ids
}

func (r Repository) makeResources(
	resourcesProxies []resourceProxy,
	permissionsProxies []permissionProxy,
) []sharedModels.Resource {
	permissionsMap := r.makePermissionsMap(permissionsProxies)
	var resources []sharedModels.Resource
	for _, proxy := range resourcesProxies {
		resources = append(resources, sharedModels.Resource{
			Id:          proxy.Id,
			Title:       proxy.Title,
			LinksTo:     proxy.makeLinksTo(),
			Permissions: permissionsMap[proxy.Id],
		})
	}

	return resources
}

func (r Repository) makePermissionsMap(
	permissionsProxies []permissionProxy,
) map[sharedModels.ResourceId][]sharedModels.Permission {
	permissionsMap := make(map[sharedModels.ResourceId][]sharedModels.Permission)
	for _, permission := range permissionsProxies {
		permissionsMap[permission.ResourceId] = append(permissionsMap[permission.ResourceId], sharedModels.Permission{
			Id:        permission.Id,
			Operation: permission.Operation,
			Effect:    permission.Effect,
		})
	}

	return permissionsMap
}

func (r Repository) Update(accountId sharedModels.AccountId, resource sharedModels.Resource) error {
	_, err := r.db.NamedExec(updateResourceQuery, sharedResourceRepository.MapFromResource(accountId, resource))

	return err
}

func (r Repository) Delete(accountId sharedModels.AccountId, resourceId sharedModels.ResourceId) error {
	_, err := r.db.Exec(removeResourcePermissions, accountId, resourceId)
	if err != nil {
		return err
	}

	_, err = r.db.Exec(deleteResourceQuery, accountId, resourceId)

	return err
}
