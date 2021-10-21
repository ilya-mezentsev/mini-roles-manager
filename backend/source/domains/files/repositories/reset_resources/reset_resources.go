package reset_resources

import (
	"github.com/jmoiron/sqlx"
	sharedModels "mini-roles-backend/source/domains/shared/models"
	sharedPermissionRepository "mini-roles-backend/source/domains/shared/repositories/permission"
	sharedResourceRepository "mini-roles-backend/source/domains/shared/repositories/resource"
)

const (
	deleteResourcesQuery = `delete from resource where account_hash = $1`
)

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repository {
	return Repository{db}
}

func (r Repository) Reset(accountId sharedModels.AccountId, resources []sharedModels.Resource) error {
	_, err := r.db.Exec(deleteResourcesQuery, accountId)
	if err != nil {
		return err
	}

	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}

	for _, resource := range resources {
		_, err = tx.NamedExec(
			sharedResourceRepository.CreateResourceQuery,
			sharedResourceRepository.MapFromResource(accountId, resource),
		)
		if err != nil {
			return err
		}

		for _, permission := range resource.Permissions {
			_, err = tx.NamedExec(
				sharedPermissionRepository.CreatePermissionQuery,
				sharedPermissionRepository.MapFromPermission(
					accountId,
					resource.Id,
					permission,
				),
			)
			if err != nil {
				return err
			}
		}
	}

	return tx.Commit()
}
