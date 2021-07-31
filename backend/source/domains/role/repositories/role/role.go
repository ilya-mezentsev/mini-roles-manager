package role

import (
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	sharedError "mini-roles-backend/source/domains/shared/error"
	sharedModels "mini-roles-backend/source/domains/shared/models"
	sharedSpec "mini-roles-backend/source/domains/shared/spec"
)

const (
	selectRolesQuery = `
	select trim(role_id) role_id, trim(title) title, trim(version_id) version_id, permissions, extends
	from role where account_hash = $1
	order by created_at`

	addRoleQuery = `
	insert into role(account_hash, role_id, version_id, title, permissions, extends)
	values(:account_hash, :role_id, :version_id, :title, :permissions, :extends)`

	updateRoleQuery = `
	update role
	set title = :title, permissions = :permissions, extends = :extends
	where account_hash = :account_hash and role_id = :role_id and version_id = :version_id`

	deleteRoleQuery       = `delete from role where account_hash = $1 and version_id = $2 and role_id = $3`
	deleteRoleFromExtends = `update role set extends = array_remove(extends, $3) where account_hash = $1 and version_id = $2`
)

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repository {
	return Repository{db}
}

func (r Repository) Create(accountId sharedModels.AccountId, role sharedModels.Role) error {
	_, err := r.db.NamedExec(addRoleQuery, r.mapFromRole(accountId, role))
	if sharedError.IsDuplicateKey(err) {
		err = sharedError.DuplicateUniqueKey{}
	}

	return err
}

func (r Repository) mapFromRole(accountId sharedModels.AccountId, role sharedModels.Role) map[string]interface{} {
	return map[string]interface{}{
		"role_id":      role.Id,
		"version_id":   role.VersionId,
		"title":        role.Title,
		"permissions":  pq.Array(role.Permissions),
		"extends":      pq.Array(role.Extends),
		"account_hash": accountId,
	}
}

func (r Repository) List(spec sharedSpec.AccountWithId) ([]sharedModels.Role, error) {
	var roles []roleProxy
	err := r.db.Select(&roles, selectRolesQuery, spec.AccountId)

	return r.makeRoles(roles), err
}

func (r Repository) makeRoles(proxies []roleProxy) []sharedModels.Role {
	var roles []sharedModels.Role
	for _, proxy := range proxies {
		roles = append(roles, sharedModels.Role{
			Id:          proxy.Id,
			VersionId:   proxy.VersionId,
			Title:       proxy.Title,
			Permissions: proxy.makePermissions(),
			Extends:     proxy.makeExtends(),
		})
	}

	return roles
}

func (r Repository) Update(accountId sharedModels.AccountId, role sharedModels.Role) error {
	_, err := r.db.NamedExec(updateRoleQuery, r.mapFromRole(accountId, role))

	return err
}

func (r Repository) Delete(
	accountId sharedModels.AccountId,
	rolesVersionId sharedModels.RolesVersionId,
	roleId sharedModels.RoleId,
) error {
	_, err := r.db.Exec(deleteRoleQuery, accountId, rolesVersionId, roleId)
	if err != nil {
		return err
	}

	_, err = r.db.Exec(deleteRoleFromExtends, accountId, rolesVersionId, roleId)

	return err
}
