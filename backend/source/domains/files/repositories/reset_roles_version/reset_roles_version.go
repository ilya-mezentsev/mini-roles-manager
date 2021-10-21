package reset_roles_version

import (
	"github.com/jmoiron/sqlx"
	sharedModels "mini-roles-backend/source/domains/shared/models"
	sharedRolesVersionRepository "mini-roles-backend/source/domains/shared/repositories/roles_version"
)

const (
	deleteRolesVersionsQuery = `delete from roles_version where account_hash = $1`
)

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repository {
	return Repository{db}
}

func (r Repository) Reset(
	accountId sharedModels.AccountId,
	defaultRolesVersion sharedModels.RolesVersion,
	rolesVersions []sharedModels.RolesVersion,
) error {
	_, err := r.db.Exec(deleteRolesVersionsQuery, accountId)
	if err != nil {
		return err
	}

	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}

	// "Default" now is just first created.
	// fixme. Should be rewritten when default roles version become managed.
	_, err = tx.NamedExec(
		sharedRolesVersionRepository.AddRolesVersionQuery,
		sharedRolesVersionRepository.MapFromRolesVersion(accountId, defaultRolesVersion),
	)
	if err != nil {
		return err
	}

	for _, rolesVersion := range rolesVersions {
		_, err = tx.NamedExec(
			sharedRolesVersionRepository.AddRolesVersionQuery,
			sharedRolesVersionRepository.MapFromRolesVersion(accountId, rolesVersion),
		)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}
