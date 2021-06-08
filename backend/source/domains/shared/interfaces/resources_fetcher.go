package interfaces

import sharedModels "mini-roles-backend/source/domains/shared/models"

type ResourceFetcherRepository interface {
	List(accountId sharedModels.AccountId) ([]sharedModels.Resource, error)
}
