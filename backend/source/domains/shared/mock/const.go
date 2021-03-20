package mock

import (
	shared "mini-roles-backend/source/domains/shared/models"
	"mini-roles-backend/source/domains/shared/services/hash"
)

var (
	BadAccountId  = shared.AccountId(hash.Md5WithTimeAsKey("bad-account-id"))
	BadResourceId = shared.ResourceId("bad-resource-id")
)

var (
	ExistsAccountId  = shared.AccountId(hash.Md5WithTimeAsKey("exists-account-id"))
	ExistsLogin      = "ExistsLogin"
	ExistsPassword   = "exists-password"
	ExistsRoleId     = shared.RoleId(hash.Md5WithTimeAsKey("exists-role-id"))
	ExistsResourceId = shared.ResourceId("exists-resource-id")
)
