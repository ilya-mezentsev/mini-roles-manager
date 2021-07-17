package interfaces

import (
	sharedModels "mini-roles-backend/source/domains/shared/models"
	sharedSpec "mini-roles-backend/source/domains/shared/spec"
)

type (
	RolesVersionCreatorRepository interface {
		Create(
			accountId sharedModels.AccountId,
			rolesVersion sharedModels.RolesVersion,
		) error
	}

	DefaultRolesVersionFetcherRepository interface {
		Fetch(spec sharedSpec.AccountWithId) (sharedModels.RolesVersion, error)
	}
)
