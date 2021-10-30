package permission

import sharedModels "mini-roles-backend/source/domains/shared/models"

func MapFromPermission(
	accountId sharedModels.AccountId,
	resourceId sharedModels.ResourceId,
	permission sharedModels.Permission,
) map[string]interface{} {
	return map[string]interface{}{
		"resource_id":   resourceId,
		"account_hash":  accountId,
		"permission_id": permission.Id,
		"operation":     permission.Operation,
		"effect":        permission.Effect,
	}
}
