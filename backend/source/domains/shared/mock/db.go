package mock

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"mini-roles-backend/source/db/schema"
	sharedModels "mini-roles-backend/source/domains/shared/models"
)

const (
	accountTable            = "account"
	accountCredentialsTable = "account_credentials"
	resourceTable           = "resource"
	permissionTable         = "permission"
	versionTable            = "roles_version"
	roleTable               = "role"
)

func MustReinstall(db *sqlx.DB) {
	dropTables(db)
	createTables(db)
	addAccount(db)
	addRolesVersion(db)
}

func dropTables(db *sqlx.DB) {
	for _, tableName := range [...]string{
		accountTable,
		accountCredentialsTable,
		resourceTable,
		permissionTable,
		versionTable,
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

func addRolesVersion(db *sqlx.DB) {
	MustAddRolesVersion(db, ExistsRolesVersionId)
}

func MustAddRolesVersion(db *sqlx.DB, rolesVersionId sharedModels.RolesVersionId) {
	db.MustExec(`insert into roles_version(version_id, account_hash) values($1, $2)`, rolesVersionId, ExistsAccountId)
}

func MustAddResource(db *sqlx.DB) {
	_, err := db.NamedExec(
		`insert into resource(account_hash, resource_id, title) values(:account_hash, :resource_id, :title)`,
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

func MustDropPermissionTable(db *sqlx.DB) {
	db.MustExec(fmt.Sprintf("drop table if exists %s cascade", permissionTable))
}

func MustDropResourceTable(db *sqlx.DB) {
	db.MustExec(fmt.Sprintf("drop table if exists %s cascade", resourceTable))
}

func MustDropRoleTable(db *sqlx.DB) {
	db.MustExec(fmt.Sprintf("drop table if exists %s cascade", roleTable))
}

func MustDropRolesVersionTable(db *sqlx.DB) {
	db.MustExec(fmt.Sprintf("drop table if exists %s cascade", versionTable))
}
