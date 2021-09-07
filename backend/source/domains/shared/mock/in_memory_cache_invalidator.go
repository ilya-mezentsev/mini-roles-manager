package mock

import (
	sharedModels "mini-roles-backend/source/domains/shared/models"
)

type InMemoryCacheInvalidator struct {
	accountId sharedModels.AccountId
}

func (i *InMemoryCacheInvalidator) Reset() {
	i.accountId = ""
}

func (i InMemoryCacheInvalidator) InvalidateCalledWith(accountId sharedModels.AccountId) bool {
	return i.accountId == accountId
}

func (i *InMemoryCacheInvalidator) Invalidate(accountId sharedModels.AccountId) {
	i.accountId = accountId
}
