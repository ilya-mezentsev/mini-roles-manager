package interfaces

import (
	"mini-roles-backend/source/domains/account/models"
	sharedModels "mini-roles-backend/source/domains/shared/models"
)

type (
	AccountInfoRepository interface {
		FetchInfo(accountId sharedModels.AccountId) (models.AccountInfo, error)

		UpdateCredentials(
			accountId sharedModels.AccountId,
			credentials models.UpdateAccountCredentials,
		) error

		UpdateLogin(
			accountId sharedModels.AccountId,
			newLogin string,
		) error
	}
)
