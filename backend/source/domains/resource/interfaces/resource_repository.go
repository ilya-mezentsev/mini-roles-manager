package interfaces

import (
	sharedModels "mini-roles-backend/source/domains/shared/models"
)

type (
	ResourceRepository interface {
		Create(
			accountId sharedModels.AccountId,
			resource sharedModels.Resource,
		) error

		List(accountId sharedModels.AccountId) ([]sharedModels.Resource, error)

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
