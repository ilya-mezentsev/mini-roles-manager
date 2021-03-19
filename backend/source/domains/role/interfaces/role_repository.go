package interfaces

import (
	shared "mini-roles-backend/source/domains/shared/models"
)

type (
	RoleRepository interface {
		Create(
			accountId shared.AccountId,
			role shared.Role,
		) error

		List(accountId shared.AccountId) ([]shared.Role, error)

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
