package interfaces

import "mini-roles-backend/source/domains/account/models"

type (
	RegistrationRepository interface {
		Register(
			session models.AccountSession,
			credentials models.AccountCredentials,
		) error
	}
)
