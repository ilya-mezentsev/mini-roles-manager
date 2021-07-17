package permission

import (
	"mini-roles-backend/source/domains/permission/spec"
	sharedModels "mini-roles-backend/source/domains/shared/models"
)

type Repository struct {
	appData sharedModels.AppData
}

func New(appData sharedModels.AppData) Repository {
	return Repository{appData}
}

func (r Repository) List(spec spec.PermissionWithAccountIdAndRoleId) ([]sharedModels.Permission, error) {
	return r.list(
		spec.RolesVersionId,
		spec.RoleId,
		make(map[sharedModels.RoleId]struct{}),
	), nil
}

func (r Repository) list(
	rolesVersionId sharedModels.RolesVersionId,
	entryPointRoleId sharedModels.RoleId,
	exclude map[sharedModels.RoleId]struct{},
) []sharedModels.Permission {
	for _, role := range r.appData.Roles {
		if role.Id == entryPointRoleId && role.VersionId == rolesVersionId {
			exclude[role.Id] = struct{}{}

			return r.findPermissions(
				rolesVersionId,
				role,
				exclude,
			)
		}
	}

	return nil
}

func (r Repository) findPermissions(
	rolesVersionId sharedModels.RolesVersionId,
	role sharedModels.Role,
	exclude map[sharedModels.RoleId]struct{},
) []sharedModels.Permission {
	var permissions []sharedModels.Permission
	for _, permissionId := range role.Permissions {
		for _, resource := range r.appData.Resources {
			for _, permission := range resource.Permissions {
				if permission.Id == permissionId {
					permissions = append(permissions, sharedModels.Permission{
						Id:        permission.Id,
						Operation: permission.Operation,
						Effect:    permission.Effect,
						Resource: sharedModels.Resource{
							Id:      resource.Id,
							LinksTo: resource.LinksTo,
						},
					})
				}
			}
		}
	}

	for _, roleId := range role.Extends {
		if _, excluded := exclude[roleId]; excluded {
			continue
		}

		permissions = append(
			permissions,
			r.list(
				rolesVersionId,
				roleId,
				exclude,
			)...,
		)
	}

	return permissions
}
