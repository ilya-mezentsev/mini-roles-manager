package mock

import (
	"errors"
	sharedError "mini-roles-backend/source/domains/shared/error"
	sharedModels "mini-roles-backend/source/domains/shared/models"
)

type RoleRepository struct {
	roles map[sharedModels.AccountId][]sharedModels.Role
}

func (r *RoleRepository) Reset() {
	r.roles = map[sharedModels.AccountId][]sharedModels.Role{
		ExistsAccountId: {
			{
				Id:    ExistsRoleId,
				Title: "Some-Role",
			},
		},
	}
}

func (r RoleRepository) Has(role sharedModels.Role) bool {
	for _, roles := range r.roles {
		for _, existsRole := range roles {
			if existsRole.Id == role.Id {
				return true
			}
		}
	}

	return false
}

func (r RoleRepository) Get(roleId sharedModels.RoleId) (role sharedModels.Role) {
	for _, roles := range r.roles {
		for _, existsRole := range roles {
			if existsRole.Id == roleId {
				role = existsRole
			}
		}
	}

	return
}

func (r *RoleRepository) Create(accountId sharedModels.AccountId, role sharedModels.Role) error {
	if accountId == BadAccountId {
		return errors.New("some-error")
	} else if r.Has(role) {
		return sharedError.DuplicateUniqueKey{}
	}

	r.roles[accountId] = append(r.roles[accountId], role)

	return nil
}

func (r RoleRepository) List(accountId sharedModels.AccountId) ([]sharedModels.Role, error) {
	if accountId == BadAccountId || accountId == BadAccountIdForRoleRepository {
		return nil, errors.New("some-error")
	}

	return r.roles[accountId], nil
}

func (r *RoleRepository) Update(accountId sharedModels.AccountId, role sharedModels.Role) error {
	if accountId == BadAccountId {
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
	if accountId == BadAccountId {
		return errors.New("some-error")
	}

	var newRoles []sharedModels.Role
	for _, role := range r.roles[accountId] {
		if role.Id != roleId {
			newRoles = append(newRoles, role)
		}
	}

	r.roles[accountId] = newRoles

	return nil
}
