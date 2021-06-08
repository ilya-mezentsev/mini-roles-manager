package spec

import (
	"mini-roles-backend/source/domains/account/models"
	sharedModels "mini-roles-backend/source/domains/shared/models"
)

type (
	SessionWithCredentials struct {
		Credentials models.AccountCredentials
	}

	SessionWithId struct {
		Id sharedModels.AccountId
	}
)
