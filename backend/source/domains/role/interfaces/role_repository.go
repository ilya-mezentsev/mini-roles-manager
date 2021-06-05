package interfaces

import (
	sharedInterfaces "mini-roles-backend/source/domains/shared/interfaces"
	shared "mini-roles-backend/source/domains/shared/models"
)

type (
	RoleRepository interface {
		sharedInterfaces.RolesFetcherRepository

		Create(
			accountId shared.AccountId,
			role shared.Role,
		) error

		Update(
			accountId shared.AccountId,
			role shared.Role,
		) error

		Delete(
			accountId shared.AccountId,
			roleId shared.RoleId,
		) error
	}
)
