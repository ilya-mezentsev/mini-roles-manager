package models

import shared "mini-roles-backend/source/domains/shared/models"

type (
	Permission struct {
		Id        shared.PermissionId `json:"permission_id"`
		Resource  shared.Resource     `json:"resource_id"`
		Operation string              `json:"operation"`          // create|read|update|delete
		Effect    string              `json:"effect" db:"effect"` // permit|deny
	}
)
