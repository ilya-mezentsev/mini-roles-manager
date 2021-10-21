package interfaces

import sharedModels "mini-roles-backend/source/domains/shared/models"

type (
	ResetRolesVersionRepository interface {
		Reset(
			accountId sharedModels.AccountId,
			defaultRolesVersion sharedModels.RolesVersion,
			rolesVersions []sharedModels.RolesVersion,
		) error
	}

	ResetResourcesRepository interface {
		Reset(
			accountId sharedModels.AccountId,
			resources []sharedModels.Resource,
		) error
	}

	ResetRolesRepository interface {
		Reset(
			accountId sharedModels.AccountId,
			roles []sharedModels.Role,
		) error
	}
)
