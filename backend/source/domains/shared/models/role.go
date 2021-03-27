package models

type (
	Role struct {
		Id          RoleId         `json:"id" validate:"required"`
		Title       string         `json:"title"`
		Permissions []PermissionId `json:"permissions"`
		Extends     []RoleId       `json:"extends"`
	}
)
