package permission

import (
	"github.com/jmoiron/sqlx"
	sharedModels "mini-roles-backend/source/domains/shared/models"
	sharedPermissionRepository "mini-roles-backend/source/domains/shared/repositories/permission"
)

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repository {
	return Repository{db}
}

func (r Repository) AddResourcePermissions(
	accountId sharedModels.AccountId,
	resourceId sharedModels.ResourceId,
	permissions []sharedModels.Permission,
) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}

	for _, permission := range permissions {
		_, err = tx.NamedExec(
			sharedPermissionRepository.CreatePermissionQuery,
			sharedPermissionRepository.MapFromPermission(accountId, resourceId, permission),
		)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}
