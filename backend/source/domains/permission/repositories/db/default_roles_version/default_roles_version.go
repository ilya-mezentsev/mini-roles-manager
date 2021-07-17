package default_roles_version

import (
	"github.com/jmoiron/sqlx"
	sharedModels "mini-roles-backend/source/domains/shared/models"
	sharedSpec "mini-roles-backend/source/domains/shared/spec"
)

const (
	defaultRolesVersionQuery = `
	select trim(version_id) version_id, trim(title) title
	from roles_version where account_hash = $1 order by created_at limit 1`
)

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repository {
	return Repository{db}
}

func (r Repository) Fetch(spec sharedSpec.AccountWithId) (rolesVersion sharedModels.RolesVersion, err error) {
	err = r.db.Get(&rolesVersion, defaultRolesVersionQuery, spec.AccountId)
	return
}
