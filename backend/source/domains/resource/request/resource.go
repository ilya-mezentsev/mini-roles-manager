package request

import sharedModels "mini-roles-backend/source/domains/shared/models"

type (
	CreateResource struct {
		AccountId sharedModels.AccountId `validate:"required"`
		Resource  sharedModels.Resource  `validate:"required"`
	}

	ResourcesList struct {
		AccountId sharedModels.AccountId `validate:"required"`
	}

	UpdateResource struct {
		AccountId sharedModels.AccountId `validate:"required"`
		Resource  sharedModels.Resource  `validate:"required"`
	}

	DeleteResource struct {
		AccountId  sharedModels.AccountId  `validate:"required"`
		ResourceId sharedModels.ResourceId `validate:"required"`
	}
)
