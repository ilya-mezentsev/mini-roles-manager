package request

import sharedModels "mini-roles-backend/source/domains/shared/models"

type (
	CreateResourceRequest struct {
		AccountId sharedModels.AccountId `json:"account_id" validate:"required"`
		Resource  sharedModels.Resource  `json:"resource" validate:"required"`
	}

	ResourcesListRequest struct {
		AccountId sharedModels.AccountId `json:"account_id" validate:"required"`
	}

	UpdateResourceRequest struct {
		AccountId sharedModels.AccountId `json:"account_id" validate:"required"`
		Resource  sharedModels.Resource  `json:"resource" validate:"required"`
	}

	DeleteResourceRequest struct {
		AccountId  sharedModels.AccountId  `json:"account_id" validate:"required"`
		ResourceId sharedModels.ResourceId `json:"resource_id" validate:"required"`
	}
)
