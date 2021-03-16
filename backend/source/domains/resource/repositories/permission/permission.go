package permission

import (
	"github.com/jmoiron/sqlx"
	sharedModels "mini-roles-backend/source/domains/shared/models"
)

const createPermissionQuery = `
insert into permission(resource_id, account_hash, permission_id, operation, effect)
values(:resource_id, :account_hash, :permission_id, :operation, :effect)`

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
		_, err = tx.NamedExec(createPermissionQuery, r.mapFromPermission(accountId, resourceId, permission))
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (r Repository) mapFromPermission(
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
