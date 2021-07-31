package interfaces

import (
	"mini-roles-backend/source/domains/permission/spec"
	sharedModels "mini-roles-backend/source/domains/shared/models"
)

type (
	PermissionRepository interface {
		List(spec spec.PermissionWithAccountIdAndRoleId) ([]sharedModels.Permission, error)
	}
)
