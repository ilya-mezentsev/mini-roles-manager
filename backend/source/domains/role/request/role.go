package request

import (
	sharedModels "mini-roles-backend/source/domains/shared/models"
)

type (
	CreateRole struct {
		AccountId sharedModels.AccountId `validate:"required"`
		Role      sharedModels.Role      `validate:"required"`
	}

	RolesList struct {
		AccountId sharedModels.AccountId `validate:"required"`
	}

	UpdateRole struct {
		AccountId sharedModels.AccountId `validate:"required"`
		Role      sharedModels.Role      `validate:"required"`
	}

	DeleteRole struct {
		AccountId sharedModels.AccountId `validate:"required"`
		RoleId    sharedModels.RoleId    `validate:"required"`
	}
)
