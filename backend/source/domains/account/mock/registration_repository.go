package mock

import (
	"errors"
	"mini-roles-backend/source/domains/account/models"
	sharedError "mini-roles-backend/source/domains/shared/error"
	sharedMock "mini-roles-backend/source/domains/shared/mock"
	sharedModels "mini-roles-backend/source/domains/shared/models"
)

type RegistrationRepository struct {
	registrations map[sharedModels.AccountId]models.AccountCredentials
}

func (r RegistrationRepository) GetAll() []models.AccountCredentials {
	var allCredentials []models.AccountCredentials
	for _, credentials := range r.registrations {
		allCredentials = append(allCredentials, credentials)
	}

	return allCredentials
}

func (r *RegistrationRepository) Reset() {
	r.registrations = map[sharedModels.AccountId]models.AccountCredentials{
		sharedMock.ExistsAccountId: {
			Login:    ExistsLogin,
			Password: "some-password",
		},
	}
}

func (r *RegistrationRepository) Register(session models.AccountSession, credentials models.AccountCredentials) error {
	if credentials.Login == BadLogin {
		return errors.New("some error")
	} else if credentials.Login == ExistsLogin {
		return sharedError.DuplicateUniqueKey{}
	}

	r.registrations[session.Id] = credentials

	return nil
}

type FailingRolesVersionCreatorRepository struct {
}

func (f FailingRolesVersionCreatorRepository) Create(sharedModels.AccountId, sharedModels.RolesVersion) error {
	return errors.New("some-error")
}
