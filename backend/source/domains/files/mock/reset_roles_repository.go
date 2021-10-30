package mock

import (
	"errors"
	sharedModels "mini-roles-backend/source/domains/shared/models"
)

type ResetRolesRepository struct {
	Roles map[sharedModels.AccountId][]sharedModels.Role
}

func (r *ResetRolesRepository) Clean() {
	r.Roles = map[sharedModels.AccountId][]sharedModels.Role{}
}

func (r *ResetRolesRepository) Reset(accountId sharedModels.AccountId, roles []sharedModels.Role) error {
	r.Roles[accountId] = roles

	return nil
}

type ResetRolesErrorRepository struct {
}

func (r ResetRolesErrorRepository) Reset(sharedModels.AccountId, []sharedModels.Role) error {
	return errors.New("some-error")
}
