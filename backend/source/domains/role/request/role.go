package request

import (
	"mini-roles-backend/source/domains/role/models"
	sharedModels "mini-roles-backend/source/domains/shared/models"
)

type (
	CreateRoleRequest struct {
		AccountId sharedModels.AccountId `json:"account_id" validate:"required"`
		Role      models.Role            `json:"role" validate:"required"`
	}

	RolesListRequest struct {
		AccountId sharedModels.AccountId `json:"account_id" validate:"required"`
	}

	UpdateRoleRequest struct {
		AccountId sharedModels.AccountId `json:"account_id" validate:"required"`
		Role      models.Role            `json:"role" validate:"required"`
	}

	DeleteRoleRequest struct {
		AccountId sharedModels.AccountId `json:"account_id" validate:"required"`
		RoleId    sharedModels.RoleId    `json:"role_id" validate:"required"`
	}
)
