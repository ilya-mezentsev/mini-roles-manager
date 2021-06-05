package mock

import (
	"errors"
	sharedError "mini-roles-backend/source/domains/shared/error"
	sharedModels "mini-roles-backend/source/domains/shared/models"
)

type ResourceRepository struct {
	resources map[sharedModels.AccountId][]sharedModels.Resource
}

func (r *ResourceRepository) Reset() {
	r.resources = map[sharedModels.AccountId][]sharedModels.Resource{
		ExistsAccountId: {
			{
				Id:    ExistsResourceId,
				Title: "Some-Title",
			},
		},
	}
}

func (r ResourceRepository) Get(accountId sharedModels.AccountId) []sharedModels.Resource {
	return r.resources[accountId]
}

func (r ResourceRepository) Has(resource sharedModels.Resource) bool {
	for _, resources := range r.resources {
		for _, existsResource := range resources {
			if existsResource.Id == resource.Id {
				return true
			}
		}
	}

	return false
}

func (r *ResourceRepository) Create(accountId sharedModels.AccountId, resource sharedModels.Resource) error {
	if accountId == BadAccountId {
		return errors.New("some-error")
	} else if r.Has(resource) {
		return sharedError.DuplicateUniqueKey{}
	}

	r.resources[accountId] = append(r.resources[accountId], resource)
	return nil
}

func (r ResourceRepository) List(accountId sharedModels.AccountId) ([]sharedModels.Resource, error) {
	if accountId == BadAccountId {
		return nil, errors.New("some-error")
	}

	return r.resources[accountId], nil
}

func (r *ResourceRepository) Update(accountId sharedModels.AccountId, resource sharedModels.Resource) error {
	if accountId == BadAccountId {
		return errors.New("some-error")
	}

	for existsResourceIndex, existsResource := range r.resources[accountId] {
		if existsResource.Id == resource.Id {
			r.resources[accountId][existsResourceIndex] = resource
		}
	}

	return nil
}

func (r *ResourceRepository) Delete(accountId sharedModels.AccountId, resourceId sharedModels.ResourceId) error {
	if accountId == BadAccountId {
		return errors.New("some-error")
	}

	var newResources []sharedModels.Resource
	for _, resource := range r.resources[accountId] {
		if resource.Id != resourceId {
			newResources = append(newResources, resource)
		}
	}
	r.resources[accountId] = newResources

	return nil
}
