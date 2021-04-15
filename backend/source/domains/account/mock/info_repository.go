package mock

import (
	"errors"
	"mini-roles-backend/source/domains/account/models"
	sharedError "mini-roles-backend/source/domains/shared/error"
	sharedMock "mini-roles-backend/source/domains/shared/mock"
	sharedModels "mini-roles-backend/source/domains/shared/models"
)

type InfoRepository struct {
	accounts    map[sharedModels.AccountId]models.AccountInfo
	credentials map[sharedModels.AccountId]models.UpdateAccountCredentials
}

func (r *InfoRepository) Reset() {
	r.accounts = map[sharedModels.AccountId]models.AccountInfo{
		sharedMock.ExistsAccountId: {
			Created: Created,
			ApiKey:  sharedMock.ExistsAccountId,
			Login:   ExistsLogin,
		},
	}
	r.credentials = map[sharedModels.AccountId]models.UpdateAccountCredentials{
		sharedMock.ExistsAccountId: {
			Login:    ExistsLogin,
			Password: ExistsPassword,
		},
	}
}

func (r InfoRepository) Credentials(accountId sharedModels.AccountId) models.UpdateAccountCredentials {
	return r.credentials[accountId]
}

func (r InfoRepository) Info(accountId sharedModels.AccountId) models.AccountInfo {
	return r.accounts[accountId]
}

func (r InfoRepository) FetchInfo(accountId sharedModels.AccountId) (models.AccountInfo, error) {
	if accountId == sharedMock.BadAccountId {
		return models.AccountInfo{}, errors.New("some-error")
	}

	return r.accounts[accountId], nil
}

func (r *InfoRepository) UpdateCredentials(accountId sharedModels.AccountId, credentials models.UpdateAccountCredentials) error {
	if accountId == sharedMock.BadAccountId {
		return errors.New("some-error")
	} else if credentials.Login == ExistsLogin {
		return sharedError.DuplicateUniqueKey{}
	}

	r.credentials[accountId] = credentials

	return nil
}

func (r *InfoRepository) UpdateLogin(accountId sharedModels.AccountId, newLogin string) error {
	if accountId == sharedMock.BadAccountId {
		return errors.New("some-error")
	} else if newLogin == ExistsLogin {
		return sharedError.DuplicateUniqueKey{}
	}

	credentials := r.credentials[accountId]
	credentials.Login = newLogin
	r.credentials[accountId] = credentials

	info := r.accounts[accountId]
	info.Login = newLogin
	r.accounts[accountId] = info

	return nil
}
