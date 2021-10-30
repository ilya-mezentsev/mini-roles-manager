package request

import (
	"io"
	sharedModels "mini-roles-backend/source/domains/shared/models"
)

type ImportRequest struct {
	AccountId sharedModels.AccountId `validate:"required"`
	File      io.Reader              `validate:"required"`
}
