package request

import shared "mini-roles-backend/source/domains/shared/models"

type (
	PermissionAccess struct {
		RoleId     shared.RoleId     `json:"roleId" validate:"required"`
		ResourceId shared.ResourceId `json:"resourceId" validate:"required"`
		Operation  string            `json:"operation" validate:"required,oneof=create read update delete"`
	}
)

func (r PermissionAccess) WithResourceId(resourceId shared.ResourceId) PermissionAccess {
	r.ResourceId = resourceId
	return r
}
