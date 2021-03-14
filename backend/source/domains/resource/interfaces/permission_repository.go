package interfaces

import sharedModels "mini-roles-backend/source/domains/shared/models"

type (
	PermissionRepository interface {
		AddResourcePermissions(
			accountId sharedModels.AccountId,
			resourceId sharedModels.ResourceId,
			permissions []sharedModels.Permission,
		) error
	}
)
