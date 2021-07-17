package mock

import (
	"errors"
	sharedError "mini-roles-backend/source/domains/shared/error"
	sharedModels "mini-roles-backend/source/domains/shared/models"
	sharedSpec "mini-roles-backend/source/domains/shared/spec"
)

type RolesVersionRepository struct {
	rolesVersions map[sharedModels.AccountId][]sharedModels.RolesVersion
}

func (r *RolesVersionRepository) Reset() {
	r.rolesVersions = map[sharedModels.AccountId][]sharedModels.RolesVersion{
		ExistsAccountId: {
			{
				Id: ExistsRolesVersionId,
			},
			{
				Id: ExistsRolesVersionId2,
			},
		},
	}
}

func (r RolesVersionRepository) Has(rolesVersion sharedModels.RolesVersion) bool {
	for _, rolesVersions := range r.rolesVersions {
		for _, rv := range rolesVersions {
			if rv == rolesVersion {
				return true
			}
		}
	}

	return false
}

func (r *RolesVersionRepository) Create(
	accountId sharedModels.AccountId,
	rolesVersion sharedModels.RolesVersion,
) error {
	if accountId == BadAccountId {
		return errors.New("some-error")
	} else if r.Has(rolesVersion) {
		return sharedError.DuplicateUniqueKey{}
	}

	r.rolesVersions[accountId] = append(r.rolesVersions[accountId], rolesVersion)

	return nil
}

func (r *RolesVersionRepository) List(spec sharedSpec.AccountWithId) ([]sharedModels.RolesVersion, error) {
	if spec.AccountId == BadAccountId {
		return nil, errors.New("some-error")
	}

	return r.rolesVersions[spec.AccountId], nil
}

func (r *RolesVersionRepository) Update(
	accountId sharedModels.AccountId,
	rolesVersion sharedModels.RolesVersion,
) error {
	if accountId == BadAccountId {
		return errors.New("some-error")
	}

	for existsRolesVersionIndex, existsRolesVersion := range r.rolesVersions[accountId] {
		if existsRolesVersion.Id == rolesVersion.Id {
			r.rolesVersions[accountId][existsRolesVersionIndex] = rolesVersion
		}
	}

	return nil
}

func (r *RolesVersionRepository) Delete(
	accountId sharedModels.AccountId,
	rolesVersionId sharedModels.RolesVersionId,
) error {
	if accountId == BadAccountId || rolesVersionId == BadRolesVersionId {
		return errors.New("some-error")
	}

	var newRolesVersions []sharedModels.RolesVersion
	for _, existsRolesVersion := range r.rolesVersions[accountId] {
		if existsRolesVersion.Id != rolesVersionId {
			newRolesVersions = append(newRolesVersions, existsRolesVersion)
		}
	}

	r.rolesVersions[accountId] = newRolesVersions

	return nil
}
