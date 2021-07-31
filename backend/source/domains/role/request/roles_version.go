package request

import sharedModels "mini-roles-backend/source/domains/shared/models"

type (
	CreateRolesVersion struct {
		AccountId    sharedModels.AccountId    `validate:"required"`
		RolesVersion sharedModels.RolesVersion `validate:"required"`
	}

	RolesVersionList struct {
		AccountId sharedModels.AccountId `validate:"required"`
	}

	UpdateRolesVersion struct {
		AccountId    sharedModels.AccountId    `validate:"required"`
		RolesVersion sharedModels.RolesVersion `validate:"required"`
	}

	DeleteRolesVersion struct {
		AccountId      sharedModels.AccountId      `validate:"required"`
		RolesVersionId sharedModels.RolesVersionId `validate:"required"`
	}
)
