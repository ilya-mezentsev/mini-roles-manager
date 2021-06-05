package request

import sharedModels "mini-roles-backend/source/domains/shared/models"

type ExportRequest struct {
	AccountId sharedModels.AccountId `validate:"required"`
}
