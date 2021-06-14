package mock

import (
	"errors"
	sharedMock "mini-roles-backend/source/domains/shared/mock"
	sharedModels "mini-roles-backend/source/domains/shared/models"
	sharedResource "mini-roles-backend/source/domains/shared/resource"
)

type PermissionRepository struct {
	permissions map[sharedModels.AccountId]map[sharedModels.ResourceId][]sharedModels.Permission
}

func (p *PermissionRepository) Reset() {
	p.permissions = map[sharedModels.AccountId]map[sharedModels.ResourceId][]sharedModels.Permission{
		sharedMock.ExistsAccountId: {
			sharedMock.ExistsResourceId: []sharedModels.Permission{
				{
					Id: "some-permission-id-1",
					Resource: sharedModels.Resource{
						Id: sharedMock.ExistsResourceId,
					},
					Operation: sharedResource.CreateOperation,
					Effect:    sharedResource.DenyEffect,
				},
				{
					Id: "some-permission-id-2",
					Resource: sharedModels.Resource{
						Id: sharedMock.ExistsResourceId,
					},
					Operation: sharedResource.ReadOperation,
					Effect:    sharedResource.PermitEffect,
				},
				{
					Id: "some-permission-id-3",
					Resource: sharedModels.Resource{
						Id: sharedMock.ExistsResourceId,
					},
					Operation: sharedResource.UpdateOperation,
					Effect:    sharedResource.DenyEffect,
				},
				{
					Id: "some-permission-id-4",
					Resource: sharedModels.Resource{
						Id: sharedMock.ExistsResourceId,
					},
					Operation: sharedResource.DeleteOperation,
					Effect:    sharedResource.DenyEffect,
				},
			},
		},
	}
}

func (p PermissionRepository) Get(
	accountId sharedModels.AccountId,
	resourceId sharedModels.ResourceId,
) []sharedModels.Permission {
	return p.permissions[accountId][resourceId]
}

func (p *PermissionRepository) AddResourcePermissions(
	accountId sharedModels.AccountId,
	resourceId sharedModels.ResourceId,
	permissions []sharedModels.Permission,
) error {
	if resourceId == sharedMock.BadResourceId {
		return errors.New("some-error")
	}

	p.permissions[accountId][resourceId] = permissions
	return nil
}
