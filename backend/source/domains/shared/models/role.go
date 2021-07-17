package models

type (
	Role struct {
		Id          RoleId         `json:"id" validate:"required"`
		VersionId   RolesVersionId `json:"versionId" validate:"required"`
		Title       string         `json:"title"`
		Permissions []PermissionId `json:"permissions"`
		Extends     []RoleId       `json:"extends"`
	}
)
