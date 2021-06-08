package models

import sharedModels "mini-roles-backend/source/domains/shared/models"

type JSONRepresentation struct {
	Resources []sharedModels.Resource `json:"resources"`
	Roles     []sharedModels.Role     `json:"roles"`
}
