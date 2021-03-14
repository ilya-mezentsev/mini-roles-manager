package models

import sharedModels "mini-roles-backend/source/domains/shared/models"

type (
	Role struct {
		Id          sharedModels.RoleId         `json:"role_id" db:"role_id" validate:"required"`
		Title       string                      `json:"title" db:"title"`
		Permissions []sharedModels.PermissionId `json:"permissions" db:"permissions"`
		Extends     []sharedModels.RoleId       `json:"extends" db:"extends"`
	}
)
