package interfaces

import sharedModels "mini-roles-backend/source/domains/shared/models"

type InMemoryCacheInvalidator interface {
	Invalidate(accountId sharedModels.AccountId)
}
