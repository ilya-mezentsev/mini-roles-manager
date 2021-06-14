package mock

import (
	sharedModels "mini-roles-backend/source/domains/shared/models"
	"mini-roles-backend/source/domains/shared/services/hash"
)

var (
	BadAccountId    = sharedModels.AccountId(hash.Md5WithTimeAsKey("bad-account-id"))
	BadResourceId   = sharedModels.ResourceId("bad-resource-id")
	BadRoleId       = sharedModels.RoleId("bad-role-id")
	BadPermissionId = sharedModels.PermissionId("bad-permission-id")

	BadAccountIdForRoleRepository = sharedModels.AccountId(hash.Md5WithTimeAsKey("bad-account-id-for-role"))
)

var (
	ExistsAccountId  = sharedModels.AccountId(hash.Md5WithTimeAsKey("exists-account-id"))
	ExistsLogin      = "ExistsLogin"
	ExistsPassword   = "exists-password"
	ExistsRoleId     = sharedModels.RoleId(hash.Md5WithTimeAsKey("exists-role-id"))
	ExistsResourceId = sharedModels.ResourceId("exists-resource-id")
)
