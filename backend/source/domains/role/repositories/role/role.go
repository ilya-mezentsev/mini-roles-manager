package role

import (
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	sharedError "mini-roles-backend/source/domains/shared/error"
	sharedModels "mini-roles-backend/source/domains/shared/models"
	sharedRepositories "mini-roles-backend/source/domains/shared/repositories"
)

const (
	selectRolesQuery = `
	select trim(role_id) role_id, trim(title) title
	from role where account_hash = $1
	order by created_at`
	selectPermissionsQuery = `
	select trim(permission_id) permission_id, trim(role_id) role_id
	from role_permission
	where role_id = any($2) and account_hash = $1`
	selectExtendingQuery = `
	select trim(extends_from) extends_from, trim(role_id) role_id
	from role_extending
	where role_id = any($2) and account_hash = $1`

	addRoleQuery = `
	insert into role(account_hash, role_id, title)
	values(:account_hash, :role_id, :title)`
	addRolePermissionQuery = `
	insert into role_permission(permission_id, role_id, account_hash)
	values(:permission_id, :role_id, :account_hash)
	on conflict do nothing`
	addRoleExtendingQuery = `
	insert into role_extending(extends_from, role_id, account_hash)
	values(:extends_from, :role_id, :account_hash)
	on conflict do nothing`

	updateRoleQuery = `
	update role
	set title = :title
	where account_hash = :account_hash and role_id = :role_id`
	deleteDeletedRolePermissionsQuery = `
	delete from role_permission
	where
		not permission_id = any(:exist_permissions_ids) and
		role_id = :role_id and
		account_hash = :account_hash`
	deleteDeletedRoleExtendingQuery = `
	delete from role_extending
	where
		not extends_from = any(:exist_extending_ids) and
		role_id = :role_id and
		account_hash = :account_hash`

	deleteRoleQuery = `delete from role where account_hash = $1 and role_id = $2`
)

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repository {
	return Repository{db}
}

func (r Repository) Create(accountId sharedModels.AccountId, role sharedModels.Role) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}

	_, err = tx.NamedExec(addRoleQuery, r.mapFromRole(accountId, role))
	if err != nil {
		if sharedRepositories.IsDuplicateKey(err) {
			err = sharedError.DuplicateUniqueKey{}
		}

		return err
	}

	for _, permissionMap := range r.mapFromPermission(accountId, role) {
		_, err = tx.NamedExec(addRolePermissionQuery, permissionMap)
		if err != nil {
			_ = tx.Rollback()

			return err
		}
	}

	for _, extendsMap := range r.mapFromExtends(accountId, role) {
		_, err = tx.NamedExec(addRoleExtendingQuery, extendsMap)
		if err != nil {
			_ = tx.Rollback()

			return err
		}
	}

	return tx.Commit()
}

func (r Repository) mapFromRole(accountId sharedModels.AccountId, role sharedModels.Role) map[string]interface{} {
	return map[string]interface{}{
		"role_id":      role.Id,
		"title":        role.Title,
		"account_hash": accountId,
	}
}

func (r Repository) mapFromPermission(accountId sharedModels.AccountId, role sharedModels.Role) []map[string]interface{} {
	var permissionsMap []map[string]interface{}
	for _, permissionId := range role.Permissions {
		permissionsMap = append(permissionsMap, map[string]interface{}{
			"permission_id": permissionId,
			"role_id":       role.Id,
			"account_hash":  accountId,
		})
	}

	return permissionsMap
}

func (r Repository) mapFromExtends(accountId sharedModels.AccountId, role sharedModels.Role) []map[string]interface{} {
	var extendsMap []map[string]interface{}
	for _, extendsFromRoleId := range role.Extends {
		extendsMap = append(extendsMap, map[string]interface{}{
			"extends_from": extendsFromRoleId,
			"role_id":      role.Id,
			"account_hash": accountId,
		})
	}

	return extendsMap
}

func (r Repository) List(accountId sharedModels.AccountId) ([]sharedModels.Role, error) {
	var roles []roleProxy
	err := r.db.Select(&roles, selectRolesQuery, accountId)
	if err != nil {
		return nil, err
	}

	rolesIds := pq.Array(r.makeRolesIds(roles))
	var rolesPermissions []rolePermissionProxy
	err = r.db.Select(&rolesPermissions, selectPermissionsQuery, accountId, rolesIds)
	if err != nil {
		return nil, err
	}

	var rolesExtendsFrom []roleExtendingProxy
	err = r.db.Select(&rolesExtendsFrom, selectExtendingQuery, accountId, rolesIds)

	return r.makeRoles(roles, rolesPermissions, rolesExtendsFrom), err
}

func (r Repository) makeRolesIds(proxies []roleProxy) []sharedModels.RoleId {
	var rolesIds []sharedModels.RoleId
	for _, proxy := range proxies {
		rolesIds = append(rolesIds, proxy.Id)
	}

	return rolesIds
}

func (r Repository) makeRoles(
	roleProxies []roleProxy,
	rolesPermissionsProxy []rolePermissionProxy,
	rolesExtendsFrom []roleExtendingProxy,
) []sharedModels.Role {
	var roles []sharedModels.Role
	for _, proxy := range roleProxies {
		roles = append(roles, sharedModels.Role{
			Id:          proxy.Id,
			Title:       proxy.Title,
			Permissions: r.makeRolePermissions(proxy.Id, rolesPermissionsProxy),
			Extends:     r.makeRoleExtendsFrom(proxy.Id, rolesExtendsFrom),
		})
	}

	return roles
}

func (r Repository) makeRolePermissions(roleId sharedModels.RoleId, rolesPermissionsProxy []rolePermissionProxy) []sharedModels.PermissionId {
	var rolesPermissions []sharedModels.PermissionId
	for _, rolePermissionsProxy := range rolesPermissionsProxy {
		if rolePermissionsProxy.RoleId == roleId {
			rolesPermissions = append(rolesPermissions, rolePermissionsProxy.PermissionId)
		}
	}

	return rolesPermissions
}

func (r Repository) makeRoleExtendsFrom(roleId sharedModels.RoleId, rolesExtendsFrom []roleExtendingProxy) []sharedModels.RoleId {
	var roleExtendsFrom []sharedModels.RoleId
	for _, roleExtendsFromProxy := range rolesExtendsFrom {
		if roleExtendsFromProxy.RoleId == roleId {
			roleExtendsFrom = append(roleExtendsFrom, roleExtendsFromProxy.ExtendsFrom)
		}
	}

	return roleExtendsFrom
}

func (r Repository) Update(accountId sharedModels.AccountId, role sharedModels.Role) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}

	_, err = tx.NamedExec(updateRoleQuery, r.mapFromRole(accountId, role))
	if err != nil {
		return err
	}

	err = r.updatePermissions(tx, accountId, role)
	if err != nil {
		return err
	}

	err = r.updateExtends(tx, accountId, role)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r Repository) updatePermissions(
	tx *sqlx.Tx,
	accountId sharedModels.AccountId,
	role sharedModels.Role,
) error {
	for _, permissionMap := range r.mapFromPermission(accountId, role) {
		_, err := tx.NamedExec(addRolePermissionQuery, permissionMap)
		if err != nil {
			return err
		}
	}
	_, err := tx.NamedExec(deleteDeletedRolePermissionsQuery, map[string]interface{}{
		"exist_permissions_ids": pq.Array(role.Permissions),
		"role_id":               role.Id,
		"account_hash":          accountId,
	})

	return err
}

func (r Repository) updateExtends(
	tx *sqlx.Tx,
	accountId sharedModels.AccountId,
	role sharedModels.Role,
) error {
	for _, extendsMap := range r.mapFromExtends(accountId, role) {
		_, err := tx.NamedExec(addRoleExtendingQuery, extendsMap)
		if err != nil {
			return err
		}
	}
	_, err := tx.NamedExec(deleteDeletedRoleExtendingQuery, map[string]interface{}{
		"exist_extending_ids": pq.Array(role.Extends),
		"role_id":             role.Id,
		"account_hash":        accountId,
	})

	return err
}

func (r Repository) Delete(accountId sharedModels.AccountId, roleId sharedModels.RoleId) error {
	_, err := r.db.Exec(deleteRoleQuery, accountId, roleId)

	return err
}
