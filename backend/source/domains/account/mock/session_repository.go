package mock

import (
	"errors"
	"mini-roles-backend/source/domains/account/models"
	"mini-roles-backend/source/domains/account/spec"
	sharedError "mini-roles-backend/source/domains/shared/error"
	sharedMock "mini-roles-backend/source/domains/shared/mock"
	sharedModels "mini-roles-backend/source/domains/shared/models"
)

type SessionRepository struct {
	sessions map[sharedModels.AccountId]models.AccountSession
}

func (s *SessionRepository) Reset() {
	s.sessions = map[sharedModels.AccountId]models.AccountSession{
		sharedMock.ExistsAccountId: {
			Id: sharedMock.ExistsAccountId,
		},
	}
}

func (s SessionRepository) GetSession(spec spec.SessionWithCredentials) (models.AccountSession, error) {
	if spec.Credentials.Login == BadLogin {
		return models.AccountSession{}, errors.New("some-error")
	} else if spec.Credentials.Login == MissedLogin {
		return models.AccountSession{}, sharedError.EntryNotFound{}
	}

	return s.sessions[sharedMock.ExistsAccountId], nil
}

func (s SessionRepository) SessionExists(spec spec.SessionWithId) (bool, error) {
	if spec.Id == sharedMock.BadAccountId {
		return false, errors.New("some-error")
	}

	_, found := s.sessions[spec.Id]
	return found, nil
}
