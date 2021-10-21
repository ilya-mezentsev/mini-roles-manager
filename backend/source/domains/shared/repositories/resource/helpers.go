package resource

import (
	"github.com/lib/pq"
	sharedModels "mini-roles-backend/source/domains/shared/models"
)

func MapFromResource(
	accountId sharedModels.AccountId,
	resource sharedModels.Resource,
) map[string]interface{} {
	return map[string]interface{}{
		"account_hash": accountId,
		"resource_id":  resource.Id,
		"title":        resource.Title,
		"links_to":     pq.Array(resource.LinksTo),
	}
}
