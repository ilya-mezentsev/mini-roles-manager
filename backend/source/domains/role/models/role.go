package models

import sharedModels "mini-roles-backend/source/domains/shared/models"

type (
	Role struct {
		Id          sharedModels.RoleId         `json:"role_id" validate:"required"`
		Title       string                      `json:"title"`
		Permissions []sharedModels.PermissionId `json:"permissions"`
		Extends     []sharedModels.RoleId       `json:"extends"`
	}
)
