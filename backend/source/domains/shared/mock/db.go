package mock

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"mini-roles-backend/source/db/schema"
	sharedModels "mini-roles-backend/source/domains/shared/models"
)

const (
	accountTable             = "account"
	accountCredentialsTable  = "account_credentials"
	resourceTable            = "resource"
	resourcePermissionsTable = "resource_permission"
	rolePermissionsTable     = "role_permission"
	roleExtendingTable       = "role_extending"
	roleTable                = "role"
)

func MustReinstall(db *sqlx.DB) {
	dropTables(db)
	createTables(db)
	addAccount(db)
	MustAddResource(db)
}

func dropTables(db *sqlx.DB) {
	for _, tableName := range [...]string{
		accountTable,
		accountCredentialsTable,
		resourceTable,
		resourcePermissionsTable,
		rolePermissionsTable,
		roleExtendingTable,
		roleTable,
	} {
		db.MustExec(fmt.Sprintf("drop table if exists %s cascade", tableName))
	}
}

func createTables(db *sqlx.DB) {
	db.MustExec(schema.Schema)
}

func addAccount(db *sqlx.DB) {
	db.MustExec(`insert into account(hash) values($1) on conflict do nothing`, ExistsAccountId)
	db.MustExec(
		`insert into account_credentials(login, password, account_hash) values($1, $2, $3)`,
		ExistsLogin,
		ExistsPassword,
		ExistsAccountId,
	)
}

func MustAddResource(db *sqlx.DB) {
	_, err := db.NamedExec(
		`insert into resource(account_hash, resource_id, title) values(:account_hash, :resource_id, :title) on conflict do nothing`,
		map[string]interface{}{
			"account_hash": ExistsAccountId,
			"resource_id":  ExistsResourceId,
			"title":        "Some-Resource-Title",
		},
	)
	if err != nil {
		panic(err)
	}
}

func MustDropResourcePermissionTable(db *sqlx.DB) {
	db.MustExec(fmt.Sprintf("drop table if exists %s cascade", resourcePermissionsTable))
}

func MustDropRolePermissionTable(db *sqlx.DB) {
	db.MustExec(fmt.Sprintf("drop table if exists %s cascade", rolePermissionsTable))
}

func MustDropRoleExtendingTable(db *sqlx.DB) {
	db.MustExec(fmt.Sprintf("drop table if exists %s cascade", roleExtendingTable))
}

func MustDropResourceTable(db *sqlx.DB) {
	db.MustExec(fmt.Sprintf("drop table if exists %s cascade", resourceTable))
}

func MustDropRoleTable(db *sqlx.DB) {
	db.MustExec(fmt.Sprintf("drop table if exists %s cascade", roleTable))
}

func MustAddPermissions(db *sqlx.DB, permissionsIds []sharedModels.PermissionId) {
	for _, permissionId := range permissionsIds {
		_, err := db.NamedExec(
			`insert into resource_permission(permission_id, operation, effect, resource_id, account_hash)
					values(:permission_id, :operation, :effect, :resource_id, :account_hash) on conflict do nothing`,
			map[string]interface{}{
				"permission_id": permissionId,
				"operation":     "create",
				"effect":        "deny",
				"resource_id":   ExistsResourceId,
				"account_hash":  ExistsAccountId,
			},
		)
		if err != nil {
			panic(err)
		}
	}
}

func MustAddRole(db *sqlx.DB, role sharedModels.Role) {
	mustAddRole(db, role, true)
}

func MustAddRoleWithoutExtends(db *sqlx.DB, role sharedModels.Role) {
	mustAddRole(db, role, false)
}

func MustAddExtendsFrom(db *sqlx.DB, role sharedModels.Role) {
	tx := db.MustBegin()
	for _, extendsFrom := range role.Extends {
		_, err := tx.NamedExec(
			`insert into role_extending(extends_from, role_id, account_hash) values(:extends_from, :role_id, :account_hash)`,
			map[string]interface{}{
				"extends_from": extendsFrom,
				"role_id":      role.Id,
				"account_hash": ExistsAccountId,
			},
		)
		if err != nil {
			panic(err)
		}
	}

	err := tx.Commit()
	if err != nil {
		panic(err)
	}
}

func mustAddRole(db *sqlx.DB, role sharedModels.Role, addExtends bool) {
	_, err := db.NamedExec(
		`insert into role(role_id, title, account_hash) values(:role_id, :title, :account_hash)`,
		map[string]interface{}{
			"role_id":      role.Id,
			"title":        role.Title,
			"account_hash": ExistsAccountId,
		},
	)
	if err != nil {
		panic(err)
	}

	tx := db.MustBegin()

	for _, permissionId := range role.Permissions {
		_, err = tx.NamedExec(
			`insert into role_permission(permission_id, role_id, account_hash) values(:permission_id, :role_id, :account_hash)`,
			map[string]interface{}{
				"permission_id": permissionId,
				"role_id":       role.Id,
				"account_hash":  ExistsAccountId,
			},
		)
		if err != nil {
			panic(err)
		}
	}

	err = tx.Commit()
	if err != nil {
		panic(err)
	}

	if addExtends {
		MustAddExtendsFrom(db, role)
	}
}
