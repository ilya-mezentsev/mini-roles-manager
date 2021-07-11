package permission_db

import (
	"github.com/lib/pq"
	sharedModels "mini-roles-backend/source/domains/shared/models"
	"strings"
)

type permissionProxy struct {
	PermissionId sharedModels.PermissionId `db:"permission_id"`
	Operation    string                    `db:"operation"`
	Effect       string                    `db:"effect"`
	ResourceId   sharedModels.ResourceId   `db:"resource_id"`
	LinksTo      pq.StringArray            `db:"links_to"`
}

func (p permissionProxy) makeLinksTo() []sharedModels.ResourceId {
	var linksTo []sharedModels.ResourceId
	for _, proxyLinksTo := range p.LinksTo {
		linksTo = append(
			linksTo,
			sharedModels.ResourceId(strings.TrimSpace(proxyLinksTo)),
		)
	}

	return linksTo
}
