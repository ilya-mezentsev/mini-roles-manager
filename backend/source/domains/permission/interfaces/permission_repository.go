package interfaces

import (
	"mini-roles-backend/source/domains/permission/spec"
	shared "mini-roles-backend/source/domains/shared/models"
)

type (
	PermissionRepository interface {
		List(spec spec.PermissionWithAccountIdAndRoleId) ([]shared.Permission, error)
	}
)
