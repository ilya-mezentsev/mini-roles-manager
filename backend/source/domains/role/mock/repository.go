package mock

import (
	"errors"
	"mini-roles-backend/source/domains/role/models"
	sharedError "mini-roles-backend/source/domains/shared/error"
	sharedMock "mini-roles-backend/source/domains/shared/mock"
	sharedModels "mini-roles-backend/source/domains/shared/models"
)

type RoleRepository struct {
	roles map[sharedModels.AccountId][]models.Role
}

func (r *RoleRepository) Reset() {
	r.roles = map[sharedModels.AccountId][]models.Role{
		sharedMock.ExistsAccountId: {
			{
				Id:    sharedMock.ExistsRoleId,
				Title: "Some-Role",
			},
		},
	}
}

func (r RoleRepository) Has(role models.Role) bool {
	for _, roles := range r.roles {
		for _, existsRole := range roles {
			if existsRole.Id == role.Id {
				return true
			}
		}
	}

	return false
}

func (r RoleRepository) Get(roleId sharedModels.RoleId) (role models.Role) {
	for _, roles := range r.roles {
		for _, existsRole := range roles {
			if existsRole.Id == roleId {
				role = existsRole
			}
		}
	}

	return
}

func (r *RoleRepository) Create(accountId sharedModels.AccountId, role models.Role) error {
	if accountId == sharedMock.BadAccountId {
		return errors.New("some-error")
	} else if r.Has(role) {
		return sharedError.DuplicateUniqueKey{}
	}

	r.roles[accountId] = append(r.roles[accountId], role)

	return nil
}

func (r RoleRepository) List(accountId sharedModels.AccountId) ([]models.Role, error) {
	if accountId == sharedMock.BadAccountId {
		return nil, errors.New("some-error")
	}

	return r.roles[accountId], nil
}

func (r *RoleRepository) Update(accountId sharedModels.AccountId, role models.Role) error {
	if accountId == sharedMock.BadAccountId {
		return errors.New("some-error")
	}

	for existsRoleIndex, existsRole := range r.roles[accountId] {
		if existsRole.Id == role.Id {
			r.roles[accountId][existsRoleIndex] = role
		}
	}

	return nil
}

func (r *RoleRepository) Delete(accountId sharedModels.AccountId, roleId sharedModels.RoleId) error {
	if accountId == sharedMock.BadAccountId {
		return errors.New("some-error")
	}

	var newRoles []models.Role
	for _, role := range r.roles[accountId] {
		if role.Id != roleId {
			newRoles = append(newRoles, role)
		}
	}

	r.roles[accountId] = newRoles

	return nil
}
