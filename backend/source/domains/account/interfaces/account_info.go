package interfaces

import (
	"mini-roles-backend/source/domains/account/models"
	sharedModels "mini-roles-backend/source/domains/shared/models"
	sharedSpec "mini-roles-backend/source/domains/shared/spec"
)

type (
	AccountInfoRepository interface {
		FetchInfo(spec sharedSpec.AccountWithId) (models.AccountInfo, error)

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
