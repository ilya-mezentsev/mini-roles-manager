package spec

import sharedModels "mini-roles-backend/source/domains/shared/models"

type PermissionWithAccountIdAndRoleId struct {
	AccountId      sharedModels.AccountId
	RoleId         sharedModels.RoleId
	RolesVersionId sharedModels.RolesVersionId
}
