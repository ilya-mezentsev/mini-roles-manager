package request

import (
	"mini-roles-backend/source/domains/account/models"
	sharedModels "mini-roles-backend/source/domains/shared/models"
)

type (
	GetInfoRequest struct {
		AccountId sharedModels.AccountId `validate:"required"`
	}

	UpdateCredentialsRequest struct {
		AccountId   sharedModels.AccountId          `validate:"required"`
		Credentials models.UpdateAccountCredentials `validate:"required"`
	}
)
