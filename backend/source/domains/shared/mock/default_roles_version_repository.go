package mock

import (
	"errors"
	sharedModels "mini-roles-backend/source/domains/shared/models"
	sharedSpec "mini-roles-backend/source/domains/shared/spec"
)

type DefaultRolesVersionRepository struct {
	rolesVersions map[sharedModels.AccountId][]sharedModels.RolesVersion
}

func (r *DefaultRolesVersionRepository) Reset() {
	r.rolesVersions = map[sharedModels.AccountId][]sharedModels.RolesVersion{
		ExistsAccountId: {
			{
				Id: ExistsRolesVersionId,
			},
		},
	}
}

func (r *DefaultRolesVersionRepository) Add(
	accountId sharedModels.AccountId,
	rolesVersion sharedModels.RolesVersion,
) {
	r.rolesVersions[accountId] = append(r.rolesVersions[accountId], rolesVersion)
}

func (r DefaultRolesVersionRepository) Fetch(spec sharedSpec.AccountWithId) (sharedModels.RolesVersion, error) {
	if spec.AccountId == BadAccountIdForDefaultRolesVersionRepository {
		return sharedModels.RolesVersion{}, errors.New("some-error")
	}

	return r.rolesVersions[spec.AccountId][0], nil
}
