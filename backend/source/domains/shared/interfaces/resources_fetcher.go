package interfaces

import (
	sharedModels "mini-roles-backend/source/domains/shared/models"
	sharedSpec "mini-roles-backend/source/domains/shared/spec"
)

type ResourceFetcherRepository interface {
	List(spec sharedSpec.AccountWithId) ([]sharedModels.Resource, error)
}
