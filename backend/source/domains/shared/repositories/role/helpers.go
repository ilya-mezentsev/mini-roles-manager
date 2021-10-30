package role

import (
	"github.com/lib/pq"
	sharedModels "mini-roles-backend/source/domains/shared/models"
)

func MapFromRole(accountId sharedModels.AccountId, role sharedModels.Role) map[string]interface{} {
	return map[string]interface{}{
		"role_id":      role.Id,
		"version_id":   role.VersionId,
		"title":        role.Title,
		"permissions":  pq.Array(role.Permissions),
		"extends":      pq.Array(role.Extends),
		"account_hash": accountId,
	}
}
