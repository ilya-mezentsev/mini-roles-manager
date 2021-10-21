package mock

import (
	"errors"
	sharedModels "mini-roles-backend/source/domains/shared/models"
)

type ResetRolesVersionsRepository struct {
	RolesVersions        map[sharedModels.AccountId][]sharedModels.RolesVersion
	DefaultRolesVersions map[sharedModels.AccountId]sharedModels.RolesVersion
}

func (r *ResetRolesVersionsRepository) Clean() {
	r.RolesVersions = map[sharedModels.AccountId][]sharedModels.RolesVersion{}
	r.DefaultRolesVersions = map[sharedModels.AccountId]sharedModels.RolesVersion{}
}

func (r *ResetRolesVersionsRepository) Reset(
	accountId sharedModels.AccountId,
	defaultRolesVersion sharedModels.RolesVersion,
	rolesVersions []sharedModels.RolesVersion,
) error {
	r.RolesVersions[accountId] = rolesVersions
	r.DefaultRolesVersions[accountId] = defaultRolesVersion

	return nil
}

type ResetRolesVersionsErrorRepository struct {
}

func (r ResetRolesVersionsErrorRepository) Reset(
	sharedModels.AccountId,
	sharedModels.RolesVersion,
	[]sharedModels.RolesVersion,
) error {
	return errors.New("some-error")
}
