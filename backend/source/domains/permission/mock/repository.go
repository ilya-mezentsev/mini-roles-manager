package mock

import (
	"errors"
	sharedMock "mini-roles-backend/source/domains/shared/mock"
	shared "mini-roles-backend/source/domains/shared/models"
)

type PermissionRepository struct {
	permissions map[shared.AccountId]map[shared.RoleId][]shared.Permission
}

func (p *PermissionRepository) Reset() {
	p.permissions = map[shared.AccountId]map[shared.RoleId][]shared.Permission{
		sharedMock.ExistsAccountId: {
			sharedMock.ExistsRoleId: {
				{
					Id: "some-permission-id-1",
					Resource: shared.Resource{
						Id: LinkingResourceId,
						LinksTo: []shared.ResourceId{
							sharedMock.ExistsResourceId,
						},
					},
					Operation: DefinedOnLinkingOperation,
					Effect:    "permit",
				},
				{
					Id: "some-permission-id-2",
					Resource: shared.Resource{
						Id: sharedMock.ExistsResourceId,
					},
					Operation: PermittedOperation,
					Effect:    "permit",
				},
				{
					Id: "some-permission-id-3",
					Resource: shared.Resource{
						Id: sharedMock.ExistsResourceId,
					},
					Operation: DeniedOperation,
					Effect:    "deny",
				},
			},
		},
	}
}

func (p PermissionRepository) List(
	accountId shared.AccountId,
	roleId shared.RoleId,
) ([]shared.Permission, error) {
	if accountId == sharedMock.BadAccountId {
		return nil, errors.New("some-error")
	}

	rolePermissions, found := p.permissions[accountId]
	if found {
		return rolePermissions[roleId], nil
	} else {
		return nil, nil
	}
}
