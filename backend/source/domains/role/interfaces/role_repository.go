package interfaces

import (
	sharedInterfaces "mini-roles-backend/source/domains/shared/interfaces"
	sharedModels "mini-roles-backend/source/domains/shared/models"
)

type (
	RoleRepository interface {
		sharedInterfaces.RolesFetcherRepository

		Create(
			accountId sharedModels.AccountId,
			role sharedModels.Role,
		) error

		Update(
			accountId sharedModels.AccountId,
			role sharedModels.Role,
		) error

		Delete(
			accountId sharedModels.AccountId,
			rolesVersionId sharedModels.RolesVersionId,
			roleId sharedModels.RoleId,
		) error
	}
)
