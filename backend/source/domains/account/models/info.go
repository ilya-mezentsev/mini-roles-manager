package models

import (
	sharedModels "mini-roles-backend/source/domains/shared/models"
	"time"
)

type AccountInfo struct {
	Created time.Time              `json:"created" db:"created"`
	ApiKey  sharedModels.AccountId `json:"apiKey" db:"api_key"`
	Login   string                 `json:"login" db:"login"`
}
