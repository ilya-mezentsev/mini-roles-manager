package request

import "mini-roles-backend/source/domains/account/models"

type Registration struct {
	Credentials models.AccountCredentials `json:"credentials" validate:"required"`
}
