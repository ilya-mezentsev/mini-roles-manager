package mock

import (
	"errors"
	"mini-roles-backend/source/domains/permission/spec"
	sharedMock "mini-roles-backend/source/domains/shared/mock"
	sharedModels "mini-roles-backend/source/domains/shared/models"
	sharedResource "mini-roles-backend/source/domains/shared/resource"
)

type PermissionRepository struct {
	permissions map[sharedModels.AccountId]map[sharedModels.RolesVersionId]map[sharedModels.RoleId][]sharedModels.Permission
}

func (p *PermissionRepository) Clean() {
	p.permissions = map[sharedModels.AccountId]map[sharedModels.RolesVersionId]map[sharedModels.RoleId][]sharedModels.Permission{}
}

func (p *PermissionRepository) Reset() {
	p.permissions = map[sharedModels.AccountId]map[sharedModels.RolesVersionId]map[sharedModels.RoleId][]sharedModels.Permission{
		sharedMock.ExistsAccountId: {
			sharedMock.ExistsRolesVersionId: {
				sharedMock.ExistsRoleId: {
					{
						Id: "some-permission-id-1",
						Resource: sharedModels.Resource{
							Id: LinkingResourceId,
							LinksTo: []sharedModels.ResourceId{
								sharedMock.ExistsResourceId,
							},
						},
						Operation: DefinedOnLinkingOperation,
						Effect:    sharedResource.PermitEffect,
					},
					{
						Id: "some-permission-id-2",
						Resource: sharedModels.Resource{
							Id: sharedMock.ExistsResourceId,
						},
						Operation: PermittedOperation,
						Effect:    sharedResource.PermitEffect,
					},
					{
						Id: "some-permission-id-3",
						Resource: sharedModels.Resource{
							Id: sharedMock.ExistsResourceId,
						},
						Operation: DeniedOperation,
						Effect:    sharedResource.DenyEffect,
					},
				},
			},
		},

		sharedMock.ExistsAccountId2: {
			sharedMock.ExistsRolesVersionId: {
				sharedMock.ExistsRoleId: {
					{
						Id: "some-permission-id-1",
						Resource: sharedModels.Resource{
							Id: LinkingResourceId,
							LinksTo: []sharedModels.ResourceId{
								sharedMock.ExistsResourceId,
							},
						},
						Operation: DefinedOnLinkingOperation,
						Effect:    sharedResource.PermitEffect,
					},
					{
						Id: "some-permission-id-2",
						Resource: sharedModels.Resource{
							Id: sharedMock.ExistsResourceId,
						},
						Operation: PermittedOperation,
						Effect:    sharedResource.PermitEffect,
					},
					{
						Id: "some-permission-id-3",
						Resource: sharedModels.Resource{
							Id: sharedMock.ExistsResourceId,
						},
						Operation: DeniedOperation,
						Effect:    sharedResource.DenyEffect,
					},
				},
			},
		},
	}
}

func (p PermissionRepository) List(spec spec.PermissionWithAccountIdAndRoleId) ([]sharedModels.Permission, error) {
	if spec.AccountId == sharedMock.BadAccountId {
		return nil, errors.New("some-error")
	}

	rolePermissions, found := p.permissions[spec.AccountId]
	if found {
		return rolePermissions[spec.RolesVersionId][spec.RoleId], nil
	} else {
		return nil, nil
	}
}
