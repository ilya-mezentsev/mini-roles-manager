package interfaces

import (
	"mini-roles-backend/source/domains/permission/models"
	shared "mini-roles-backend/source/domains/shared/models"
)

type (
	PermissionRepository interface {
		List(
			accountId shared.AccountId,
			roleId shared.RoleId,
		) ([]models.Permission, error)
	}
)
