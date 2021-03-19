package models

type (
	Role struct {
		Id          RoleId         `json:"role_id" validate:"required"`
		Title       string         `json:"title"`
		Permissions []PermissionId `json:"permissions"`
		Extends     []RoleId       `json:"extends"`
	}
)
