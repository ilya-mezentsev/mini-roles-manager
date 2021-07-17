package request

import sharedModels "mini-roles-backend/source/domains/shared/models"

type (
	PermissionAccess struct {
		RoleId         sharedModels.RoleId     `validate:"required"`
		ResourceId     sharedModels.ResourceId `validate:"required"`
		Operation      string                  `validate:"required,oneof=create read update delete"`
		RolesVersionId sharedModels.RolesVersionId
	}
)

func (r PermissionAccess) WithResourceId(resourceId sharedModels.ResourceId) PermissionAccess {
	r.ResourceId = resourceId
	return r
}
