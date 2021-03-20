package mock

import (
	"errors"
	"mini-roles-backend/source/domains/account/models"
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

func (s SessionRepository) GetSession(credentials models.AccountCredentials) (models.AccountSession, error) {
	if credentials.Login == BadLogin {
		return models.AccountSession{}, errors.New("some-error")
	} else if credentials.Login == MissedLogin {
		return models.AccountSession{}, sharedError.EntryNotFound{}
	}

	return s.sessions[sharedMock.ExistsAccountId], nil
}

func (s SessionRepository) SessionExists(session models.AccountSession) (bool, error) {
	if session.Id == sharedMock.BadAccountId {
		return false, errors.New("some-error")
	}

	_, found := s.sessions[session.Id]
	return found, nil
}
