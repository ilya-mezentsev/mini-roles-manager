package interfaces

import "mini-roles-backend/source/domains/account/models"

type (
	SessionRepository interface {
		GetSession(credentials models.AccountCredentials) (models.AccountSession, error)
		SessionExists(session models.AccountSession) (bool, error)
	}
)
