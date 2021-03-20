package request

import (
	sharedModels "mini-roles-backend/source/domains/shared/models"
)

type (
	CreateRole struct {
		AccountId sharedModels.AccountId `json:"account_id" validate:"required"`
		Role      sharedModels.Role      `json:"role" validate:"required"`
	}

	RolesList struct {
		AccountId sharedModels.AccountId `json:"account_id" validate:"required"`
	}

	UpdateRole struct {
		AccountId sharedModels.AccountId `json:"account_id" validate:"required"`
		Role      sharedModels.Role      `json:"role" validate:"required"`
	}

	DeleteRole struct {
		AccountId sharedModels.AccountId `json:"account_id" validate:"required"`
		RoleId    sharedModels.RoleId    `json:"role_id" validate:"required"`
	}
)
