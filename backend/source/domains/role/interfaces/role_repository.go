package interfaces

import (
	"mini-roles-backend/source/domains/role/models"
	shared "mini-roles-backend/source/domains/shared/models"
)

type (
	RoleRepository interface {
		Create(
			accountId shared.AccountId,
			role models.Role,
		) error

		List(accountId shared.AccountId) ([]models.Role, error)

		Update(
			accountId shared.AccountId,
			role models.Role,
		) error

		Delete(
			accountId shared.AccountId,
			roleId shared.RoleId,
		) error
	}
)
