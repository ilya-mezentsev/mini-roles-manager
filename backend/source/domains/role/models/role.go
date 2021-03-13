package models

import shared "mini-roles-backend/source/domains/shared/models"

type (
	Role struct {
		Id          shared.RoleId         `json:"role_id" db:"role_id"`
		Title       string                `json:"title" db:"title"`
		Permissions []shared.PermissionId `json:"permissions"`
		Extends     []shared.RoleId       `json:"extends"`
	}
)
