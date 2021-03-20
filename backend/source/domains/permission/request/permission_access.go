package request

import shared "mini-roles-backend/source/domains/shared/models"

type (
	PermissionAccess struct {
		AccountId  shared.AccountId  `json:"account_id" validate:"required"`
		RoleId     shared.RoleId     `json:"role_id" validate:"required"`
		ResourceId shared.ResourceId `json:"resource_id" validate:"required"`
		Operation  string            `json:"operation" validate:"required,oneof=create read update delete"`
	}
)

func (r PermissionAccess) WithResourceId(resourceId shared.ResourceId) PermissionAccess {
	r.ResourceId = resourceId
	return r
}
