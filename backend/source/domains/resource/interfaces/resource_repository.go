package interfaces

import (
	sharedInterfaces "mini-roles-backend/source/domains/shared/interfaces"
	sharedModels "mini-roles-backend/source/domains/shared/models"
)

type (
	ResourceRepository interface {
		sharedInterfaces.ResourceFetcherRepository

		Create(
			accountId sharedModels.AccountId,
			resource sharedModels.Resource,
		) error

		Update(
			accountId sharedModels.AccountId,
			resource sharedModels.Resource,
		) error

		Delete(
			accountId sharedModels.AccountId,
			resourceId sharedModels.ResourceId,
		) error
	}
)
