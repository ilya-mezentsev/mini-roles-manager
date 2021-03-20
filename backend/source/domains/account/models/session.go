package models

import sharedModels "mini-roles-backend/source/domains/shared/models"

type AccountSession struct {
	Id sharedModels.AccountId `json:"id"`
}
