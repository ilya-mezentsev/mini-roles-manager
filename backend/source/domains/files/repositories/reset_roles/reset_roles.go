package reset_roles

import (
	"github.com/jmoiron/sqlx"
	sharedModels "mini-roles-backend/source/domains/shared/models"
	sharedRoleRepository "mini-roles-backend/source/domains/shared/repositories/role"
)

const (
	deleteRolesQuery = `delete from role where account_hash = $1`
)

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repository {
	return Repository{db}
}

func (r Repository) Reset(accountId sharedModels.AccountId, roles []sharedModels.Role) error {
	_, err := r.db.Exec(deleteRolesQuery, accountId)
	if err != nil {
		return err
	}

	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}

	for _, role := range roles {
		_, err = tx.NamedExec(
			sharedRoleRepository.AddRoleQuery,
			sharedRoleRepository.MapFromRole(accountId, role),
		)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}
