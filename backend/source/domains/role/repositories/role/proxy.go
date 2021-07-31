package role

import (
	"github.com/lib/pq"
	sharedModels "mini-roles-backend/source/domains/shared/models"
	"strings"
)

type roleProxy struct {
	Id          sharedModels.RoleId         `db:"role_id"`
	VersionId   sharedModels.RolesVersionId `db:"version_id"`
	Title       string                      `db:"title"`
	Permissions pq.StringArray              `db:"permissions"`
	Extends     pq.StringArray              `db:"extends"`
}

func (r roleProxy) makePermissions() []sharedModels.PermissionId {
	var permissions []sharedModels.PermissionId
	for _, permission := range r.Permissions {
		permissions = append(
			permissions,
			sharedModels.PermissionId(strings.TrimSpace(permission)),
		)
	}

	return permissions
}

func (r roleProxy) makeExtends() []sharedModels.RoleId {
	var extends []sharedModels.RoleId
	for _, roleId := range r.Extends {
		extends = append(
			extends,
			sharedModels.RoleId(strings.TrimSpace(roleId)),
		)
	}

	return extends
}
