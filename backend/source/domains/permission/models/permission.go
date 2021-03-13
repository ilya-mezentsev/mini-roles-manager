package models

import shared "mini-roles-backend/source/domains/shared/models"

type (
	Permission struct {
		Id         shared.PermissionId `json:"permission_id"`
		ResourceId shared.ResourceId   `json:"resource_id"`
		Operation  string              `json:"operation"`
		Effect     string              `json:"effect" db:"effect"` // permit|deny
	}
)
