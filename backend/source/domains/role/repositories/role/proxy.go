package role

import (
	sharedModels "mini-roles-backend/source/domains/shared/models"
)

type rolePermissionProxy struct {
	RoleId       sharedModels.RoleId       `db:"role_id"`
	PermissionId sharedModels.PermissionId `db:"permission_id"`
}

type roleExtendingProxy struct {
	RoleId      sharedModels.RoleId `db:"role_id"`
	ExtendsFrom sharedModels.RoleId `db:"extends_from"`
}

type roleProxy struct {
	Id    sharedModels.RoleId `db:"role_id"`
	Title string              `db:"title"`
}
