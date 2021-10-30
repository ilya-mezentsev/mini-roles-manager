package roles_version

import sharedModels "mini-roles-backend/source/domains/shared/models"

func MapFromRolesVersion(
	accountId sharedModels.AccountId,
	rolesVersion sharedModels.RolesVersion,
) map[string]interface{} {
	return map[string]interface{}{
		"account_hash": accountId,
		"version_id":   rolesVersion.Id,
		"title":        rolesVersion.Title,
	}
}
