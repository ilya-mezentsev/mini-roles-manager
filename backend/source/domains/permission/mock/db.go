package mock

import (
	"fmt"
	sharedMock "mini-roles-backend/source/domains/shared/mock"
	sharedModels "mini-roles-backend/source/domains/shared/models"
	sharedResource "mini-roles-backend/source/domains/shared/resource"
)

const (
	ExistsResourceId1 = "resource-1"
	ExistsResourceId2 = "resource-2"
	ExistsResourceId3 = "resource-3"

	FlatRoleId1                  = "role-1"
	FlatRoleId2                  = "role-2"
	OneDepthLevelExtendingRoleId = "role-3"
	TwoDepthLevelExtendingRoleId = "role-4"
	RecursiveExtendingRoleId1    = "role-5"
	RecursiveExtendingRoleId2    = "role-6"

	DenyCreatePermissionId1 = "deny-create-resource-1"
	PermitReadPermissionId1 = "permit-read-resource-1"
	DenyUpdatePermissionId1 = "deny-update-resource-1"
	DenyDeletePermissionId1 = "deny-delete-resource-1"

	PermitCreatePermissionId2 = "permit-create-resource-2"
	PermitReadPermissionId2   = "permit-read-resource-2"
	PermitUpdatePermissionId2 = "permit-update-resource-2"
	DenyDeletePermissionId2   = "deny-delete-resource-2"

	PermitCreatePermissionId3 = "permit-create-resource-3"
	PermitReadPermissionId3   = "permit-read-resource-3"
	PermitUpdatePermissionId3 = "permit-update-resource-3"
	PermitDeletePermissionId3 = "permit-delete-resource-3"
)

var (
	Resources = []sharedModels.Resource{
		{
			Id: ExistsResourceId1,
			LinksTo: []sharedModels.ResourceId{
				ExistsResourceId3,
			},
		},
		{
			Id: ExistsResourceId2,
		},
		{
			Id: ExistsResourceId3,
		},
	}

	OneDepthLevelExtendingRole = sharedModels.Role{
		Id:        OneDepthLevelExtendingRoleId,
		VersionId: sharedMock.ExistsRolesVersionId,
		Permissions: []sharedModels.PermissionId{
			PermitCreatePermissionId3,
			PermitReadPermissionId3,
			PermitUpdatePermissionId3,
			PermitDeletePermissionId3,
		},
		Extends: []sharedModels.RoleId{
			FlatRoleId1,
		},
	}

	TwoDepthLevelExtendingRole = sharedModels.Role{
		Id:        TwoDepthLevelExtendingRoleId,
		VersionId: sharedMock.ExistsRolesVersionId,
		Extends: []sharedModels.RoleId{
			FlatRoleId2,
			OneDepthLevelExtendingRoleId,
		},
	}

	RecursiveExtendingRole1 = sharedModels.Role{
		Id:          RecursiveExtendingRoleId1,
		VersionId:   sharedMock.ExistsRolesVersionId,
		Permissions: FlatRoles[0].Permissions,
		Extends: []sharedModels.RoleId{
			RecursiveExtendingRoleId2,
		},
	}

	RecursiveExtendingRole2 = sharedModels.Role{
		Id:          RecursiveExtendingRoleId2,
		VersionId:   sharedMock.ExistsRolesVersionId,
		Permissions: FlatRoles[1].Permissions,
		Extends: []sharedModels.RoleId{
			RecursiveExtendingRoleId1,
		},
	}

	// FlatRoles flat - meaning that it extends no another role
	FlatRoles = []sharedModels.Role{
		{
			Id:        FlatRoleId1,
			VersionId: sharedMock.ExistsRolesVersionId,
			Permissions: []sharedModels.PermissionId{
				DenyCreatePermissionId1,
				PermitReadPermissionId1,
				DenyUpdatePermissionId1,
				DenyDeletePermissionId1,
			},
		},
		{
			Id:        FlatRoleId2,
			VersionId: sharedMock.ExistsRolesVersionId,
			Permissions: []sharedModels.PermissionId{
				PermitCreatePermissionId2,
				PermitReadPermissionId2,
				PermitUpdatePermissionId2,
				DenyDeletePermissionId2,
			},
		},
		{
			Id:        FlatRoleId1,
			VersionId: sharedMock.ExistsRolesVersionId2,
			Permissions: []sharedModels.PermissionId{
				PermitCreatePermissionId2,
				PermitReadPermissionId2,
				PermitUpdatePermissionId2,
				DenyDeletePermissionId2,
			},
		},
	}

	Permissions = []sharedModels.Permission{
		{
			Id: DenyCreatePermissionId1,
			Resource: sharedModels.Resource{
				Id: ExistsResourceId1,
				LinksTo: []sharedModels.ResourceId{
					ExistsResourceId3,
				},
			},
			Operation: sharedResource.CreateOperation,
			Effect:    sharedResource.DenyEffect,
		},
		{
			Id: PermitReadPermissionId1,
			Resource: sharedModels.Resource{
				Id: ExistsResourceId1,
				LinksTo: []sharedModels.ResourceId{
					ExistsResourceId3,
				},
			},
			Operation: sharedResource.ReadOperation,
			Effect:    sharedResource.PermitEffect,
		},
		{
			Id: DenyUpdatePermissionId1,
			Resource: sharedModels.Resource{
				Id: ExistsResourceId1,
				LinksTo: []sharedModels.ResourceId{
					ExistsResourceId3,
				},
			},
			Operation: sharedResource.UpdateOperation,
			Effect:    sharedResource.DenyEffect,
		},
		{
			Id: DenyDeletePermissionId1,
			Resource: sharedModels.Resource{
				Id: ExistsResourceId1,
				LinksTo: []sharedModels.ResourceId{
					ExistsResourceId3,
				},
			},
			Operation: sharedResource.DeleteOperation,
			Effect:    sharedResource.DenyEffect,
		},

		{
			Id: PermitCreatePermissionId2,
			Resource: sharedModels.Resource{
				Id: ExistsResourceId2,
			},
			Operation: sharedResource.CreateOperation,
			Effect:    sharedResource.PermitEffect,
		},
		{
			Id: PermitReadPermissionId2,
			Resource: sharedModels.Resource{
				Id: ExistsResourceId2,
			},
			Operation: sharedResource.ReadOperation,
			Effect:    sharedResource.PermitEffect,
		},
		{
			Id: PermitUpdatePermissionId2,
			Resource: sharedModels.Resource{
				Id: ExistsResourceId2,
			},
			Operation: sharedResource.UpdateOperation,
			Effect:    sharedResource.PermitEffect,
		},
		{
			Id: DenyDeletePermissionId2,
			Resource: sharedModels.Resource{
				Id: ExistsResourceId2,
			},
			Operation: sharedResource.DeleteOperation,
			Effect:    sharedResource.DenyEffect,
		},

		{
			Id: PermitCreatePermissionId3,
			Resource: sharedModels.Resource{
				Id: ExistsResourceId3,
			},
			Operation: sharedResource.CreateOperation,
			Effect:    sharedResource.PermitEffect,
		},
		{
			Id: PermitReadPermissionId3,
			Resource: sharedModels.Resource{
				Id: ExistsResourceId3,
			},
			Operation: sharedResource.ReadOperation,
			Effect:    sharedResource.PermitEffect,
		},
		{
			Id: PermitUpdatePermissionId3,
			Resource: sharedModels.Resource{
				Id: ExistsResourceId3,
			},
			Operation: sharedResource.UpdateOperation,
			Effect:    sharedResource.PermitEffect,
		},
		{
			Id: PermitDeletePermissionId3,
			Resource: sharedModels.Resource{
				Id: ExistsResourceId3,
			},
			Operation: sharedResource.DeleteOperation,
			Effect:    sharedResource.DenyEffect,
		},
	}
)

func MakeRole1Permissions() []sharedModels.Permission {
	return makeCopy(Permissions[:4])
}

func MakeRole2Permissions() []sharedModels.Permission {
	return makeCopy(Permissions[4:8])
}

func MakeRoles1WithAnotherVersionPermissions() []sharedModels.Permission {
	return MakeRole2Permissions()
}

func MakeExtendingRolePermissions() []sharedModels.Permission {
	return makeCopy(Permissions[8:])
}

func MakeRecursiveExtendingRole1Permissions() []sharedModels.Permission {
	return MakeRole1Permissions()
}

func MakeRecursiveExtendingRole2Permissions() []sharedModels.Permission {
	return MakeRole2Permissions()
}

func MakeAppDataForAllRoles(roleId sharedModels.RoleId) sharedModels.AppData {
	return sharedModels.AppData{
		Resources: makeResourcesByRoleId(roleId),
		Roles:     allRoles(),
	}
}

func makeResourcesByRoleId(roleId sharedModels.RoleId) []sharedModels.Resource {
	var resources []sharedModels.Resource
	var resourcesMap = make(map[sharedModels.ResourceId]sharedModels.Resource)
	var rolePermissionsIds = makeRolePermissions(roleId)

	for _, permission := range Permissions {
		for _, rolePermissionId := range rolePermissionsIds {
			if permission.Id == rolePermissionId {
				resource, resourceFound := resourcesMap[permission.Resource.Id]
				if resourceFound {
					resource.Permissions = append(resource.Permissions, sharedModels.Permission{
						Id:        permission.Id,
						Operation: permission.Operation,
						Effect:    permission.Effect,
						Resource:  permission.Resource,
					})
					resourcesMap[permission.Resource.Id] = resource
				} else {
					resourcesMap[permission.Resource.Id] = sharedModels.Resource{
						Id:      permission.Resource.Id,
						Title:   permission.Resource.Title,
						LinksTo: permission.Resource.LinksTo,
						Permissions: []sharedModels.Permission{
							{
								Id:        permission.Id,
								Operation: permission.Operation,
								Effect:    permission.Effect,
								Resource:  permission.Resource,
							},
						},
					}
				}
			}
		}
	}

	for _, resource := range resourcesMap {
		resources = append(resources, resource)
	}

	return resources
}

func makeRolePermissions(roleId sharedModels.RoleId) []sharedModels.PermissionId {
	var (
		permissions    []sharedModels.Permission
		permissionsIds []sharedModels.PermissionId
	)

	switch roleId {
	case OneDepthLevelExtendingRoleId:
		permissions = append(
			MakeRole1Permissions(),
			MakeExtendingRolePermissions()...,
		)

	case TwoDepthLevelExtendingRoleId:
		permissions = Permissions

	case RecursiveExtendingRoleId1:
		permissions = append(
			MakeRecursiveExtendingRole1Permissions(),
			MakeRecursiveExtendingRole2Permissions()...,
		)

	case FlatRoleId1:
		permissions = MakeRole1Permissions()

	case RecursiveExtendingRoleId2:
	case FlatRoleId2:
	default:
		panic(fmt.Sprintf("Unknown role id for making permissions: %s", roleId))
	}

	for _, permission := range permissions {
		permissionsIds = append(permissionsIds, permission.Id)
	}

	return permissionsIds
}

func allRoles() []sharedModels.Role {
	return append(
		FlatRoles,
		OneDepthLevelExtendingRole,
		TwoDepthLevelExtendingRole,
		RecursiveExtendingRole1,
		RecursiveExtendingRole2,
	)
}

func makeCopy(permissions []sharedModels.Permission) []sharedModels.Permission {
	var dist []sharedModels.Permission
	copy(dist, permissions)
	return dist
}
