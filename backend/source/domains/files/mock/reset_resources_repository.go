package mock

import (
	"errors"
	sharedModels "mini-roles-backend/source/domains/shared/models"
)

type ResetResourcesRepository struct {
	Resources map[sharedModels.AccountId][]sharedModels.Resource
}

func (r *ResetResourcesRepository) Clean() {
	r.Resources = map[sharedModels.AccountId][]sharedModels.Resource{}
}

func (r *ResetResourcesRepository) Reset(accountId sharedModels.AccountId, resources []sharedModels.Resource) error {
	r.Resources[accountId] = resources

	return nil
}

type ResetResourcesErrorRepository struct {
}

func (r ResetResourcesErrorRepository) Reset(sharedModels.AccountId, []sharedModels.Resource) error {
	return errors.New("some-error")
}
