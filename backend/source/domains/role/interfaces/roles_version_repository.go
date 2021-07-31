package interfaces

import (
	sharedInterfaces "mini-roles-backend/source/domains/shared/interfaces"
	sharedModels "mini-roles-backend/source/domains/shared/models"
	sharedSpec "mini-roles-backend/source/domains/shared/spec"
)

type (
	RolesVersionRepository interface {
		sharedInterfaces.RolesVersionCreatorRepository

		List(spec sharedSpec.AccountWithId) ([]sharedModels.RolesVersion, error)

		Update(
			accountId sharedModels.AccountId,
			rolesVersion sharedModels.RolesVersion,
		) error

		Delete(
			accountId sharedModels.AccountId,
			rolesVersionId sharedModels.RolesVersionId,
		) error
	}
)
