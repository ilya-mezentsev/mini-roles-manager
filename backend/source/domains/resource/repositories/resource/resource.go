package resource

import (
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	sharedError "mini-roles-backend/source/domains/shared/error"
	sharedModels "mini-roles-backend/source/domains/shared/models"
	"strings"
)

const (
	selectResourcesQuery = `
	select trim(resource_id) resource_id, trim(title) title, links_to from resource
	where account_hash = $1`

	createResourceQuery = `
	insert into resource(account_hash, resource_id, title, links_to)
	values(:account_hash, :resource_id, :title, :links_to)`

	updateResourceQuery = `
	update resource
	set title = :title, links_to = :links_to
	where account_hash = :account_hash and resource_id = :resource_id`

	deleteResourceQuery = `delete from resource where account_hash = $1 and resource_id = $2`
)

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repository {
	return Repository{db}
}

func (r Repository) Create(accountId sharedModels.AccountId, resource sharedModels.Resource) error {
	_, err := r.db.NamedExec(createResourceQuery, r.mapFromResource(accountId, resource))
	if err != nil && strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
		err = sharedError.DuplicateUniqueKey{}
	}

	return err
}

func (r Repository) mapFromResource(
	accountId sharedModels.AccountId,
	resource sharedModels.Resource,
) map[string]interface{} {
	return map[string]interface{}{
		"account_hash": accountId,
		"resource_id":  resource.Id,
		"title":        resource.Title,
		"links_to":     pq.Array(resource.LinksTo),
	}
}

func (r Repository) List(accountId sharedModels.AccountId) ([]sharedModels.Resource, error) {
	var resources []resourceProxy
	err := r.db.Select(&resources, selectResourcesQuery, accountId)

	return r.makeResources(resources), err
}

func (r Repository) makeResources(resourcesProxies []resourceProxy) []sharedModels.Resource {
	var resources []sharedModels.Resource
	for _, proxy := range resourcesProxies {
		resources = append(resources, sharedModels.Resource{
			Id:      proxy.Id,
			Title:   proxy.Title,
			LinksTo: proxy.makeLinksTo(),
		})
	}

	return resources
}

func (r Repository) Update(accountId sharedModels.AccountId, resource sharedModels.Resource) error {
	_, err := r.db.NamedExec(updateResourceQuery, r.mapFromResource(accountId, resource))

	return err
}

func (r Repository) Delete(accountId sharedModels.AccountId, resourceId sharedModels.ResourceId) error {
	_, err := r.db.Exec(deleteResourceQuery, accountId, resourceId)

	return err
}
