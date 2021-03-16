package resource

import (
	"github.com/lib/pq"
	sharedModels "mini-roles-backend/source/domains/shared/models"
	"strings"
)

type resourceProxy struct {
	Id      sharedModels.ResourceId `db:"resource_id"`
	Title   string                  `db:"title"`
	LinksTo pq.StringArray          `db:"links_to"`
}

func (r resourceProxy) makeLinksTo() []sharedModels.ResourceId {
	var linksTo []sharedModels.ResourceId
	for _, id := range r.LinksTo {
		linksTo = append(
			linksTo,
			sharedModels.ResourceId(strings.TrimSpace(id)),
		)
	}

	return linksTo
}
