package spec

import shared "mini-roles-backend/source/domains/shared/models"

type PermissionWithAccountIdAndRoleId struct {
	AccountId shared.AccountId
	RoleId    shared.RoleId
}
