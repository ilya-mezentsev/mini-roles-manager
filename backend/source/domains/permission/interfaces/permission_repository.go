package interfaces

import (
	shared "mini-roles-backend/source/domains/shared/models"
)

type (
	PermissionRepository interface {
		List(
			accountId shared.AccountId,
			roleId shared.RoleId,
		) ([]shared.Permission, error)
	}
)
