package mock

import sharedModels "mini-roles-backend/source/domains/shared/models"

const (
	ExistsResourceId1 = "resource-1"
	ExistsResourceId2 = "resource-2"
	ExistsResourceId3 = "resource-3"

	FlatRoleId1                  = "role-1"
	FlatRoleId2                  = "role-2"
	OneDepthLevelExtendingRoleId = "role-3"
	TwoDepthLevelExtendingRoleId = "role-4"

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
		Id: OneDepthLevelExtendingRoleId,
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
		Id: TwoDepthLevelExtendingRoleId,
		Extends: []sharedModels.RoleId{
			FlatRoleId2,
			OneDepthLevelExtendingRoleId,
		},
	}

	// flat - meaning that it extends no another role
	FlatRoles = []sharedModels.Role{
		{
			Id: FlatRoleId1,
			Permissions: []sharedModels.PermissionId{
				DenyCreatePermissionId1,
				PermitReadPermissionId1,
				DenyUpdatePermissionId1,
				DenyDeletePermissionId1,
			},
		},
		{
			Id: FlatRoleId2,
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
			Operation: "create",
			Effect:    "deny",
		},
		{
			Id: PermitReadPermissionId1,
			Resource: sharedModels.Resource{
				Id: ExistsResourceId1,
				LinksTo: []sharedModels.ResourceId{
					ExistsResourceId3,
				},
			},
			Operation: "read",
			Effect:    "permit",
		},
		{
			Id: DenyUpdatePermissionId1,
			Resource: sharedModels.Resource{
				Id: ExistsResourceId1,
				LinksTo: []sharedModels.ResourceId{
					ExistsResourceId3,
				},
			},
			Operation: "update",
			Effect:    "deny",
		},
		{
			Id: DenyDeletePermissionId1,
			Resource: sharedModels.Resource{
				Id: ExistsResourceId1,
				LinksTo: []sharedModels.ResourceId{
					ExistsResourceId3,
				},
			},
			Operation: "delete",
			Effect:    "deny",
		},

		{
			Id: PermitCreatePermissionId2,
			Resource: sharedModels.Resource{
				Id: ExistsResourceId2,
			},
			Operation: "create",
			Effect:    "permit",
		},
		{
			Id: PermitReadPermissionId2,
			Resource: sharedModels.Resource{
				Id: ExistsResourceId2,
			},
			Operation: "read",
			Effect:    "permit",
		},
		{
			Id: PermitUpdatePermissionId2,
			Resource: sharedModels.Resource{
				Id: ExistsResourceId2,
			},
			Operation: "update",
			Effect:    "permit",
		},
		{
			Id: DenyDeletePermissionId2,
			Resource: sharedModels.Resource{
				Id: ExistsResourceId2,
			},
			Operation: "delete",
			Effect:    "deny",
		},

		{
			Id: PermitCreatePermissionId3,
			Resource: sharedModels.Resource{
				Id: ExistsResourceId3,
			},
			Operation: "create",
			Effect:    "permit",
		},
		{
			Id: PermitReadPermissionId3,
			Resource: sharedModels.Resource{
				Id: ExistsResourceId3,
			},
			Operation: "read",
			Effect:    "permit",
		},
		{
			Id: PermitUpdatePermissionId3,
			Resource: sharedModels.Resource{
				Id: ExistsResourceId3,
			},
			Operation: "update",
			Effect:    "permit",
		},
		{
			Id: PermitDeletePermissionId3,
			Resource: sharedModels.Resource{
				Id: ExistsResourceId3,
			},
			Operation: "delete",
			Effect:    "deny",
		},
	}
)

func MakeRole1Permissions() []sharedModels.Permission {
	return Permissions[:4]
}

func MakeRole2Permissions() []sharedModels.Permission {
	return Permissions[4:8]
}

func MakeExtendingRolePermissions() []sharedModels.Permission {
	return Permissions[8:]
}
