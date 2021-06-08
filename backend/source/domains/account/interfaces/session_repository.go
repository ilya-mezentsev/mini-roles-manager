package interfaces

import (
	"mini-roles-backend/source/domains/account/models"
	"mini-roles-backend/source/domains/account/spec"
)

type (
	SessionRepository interface {
		GetSession(spec spec.SessionWithCredentials) (models.AccountSession, error)
		SessionExists(spec spec.SessionWithId) (bool, error)
	}
)
