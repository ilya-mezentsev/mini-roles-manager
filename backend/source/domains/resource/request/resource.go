package request

import sharedModels "mini-roles-backend/source/domains/shared/models"

type (
	CreateResource struct {
		AccountId sharedModels.AccountId `json:"account_id" validate:"required"`
		Resource  sharedModels.Resource  `json:"resource" validate:"required"`
	}

	ResourcesList struct {
		AccountId sharedModels.AccountId `json:"account_id" validate:"required"`
	}

	UpdateResource struct {
		AccountId sharedModels.AccountId `json:"account_id" validate:"required"`
		Resource  sharedModels.Resource  `json:"resource" validate:"required"`
	}

	DeleteResource struct {
		AccountId  sharedModels.AccountId  `json:"account_id" validate:"required"`
		ResourceId sharedModels.ResourceId `json:"resource_id" validate:"required"`
	}
)
