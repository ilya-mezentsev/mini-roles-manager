package mock

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"mini-roles-backend/source/db/schema"
)

const (
	accountTable            = "account"
	accountCredentialsTable = "account_credentials"
	resourceTable           = "resource"
	permissionTable         = "permission"
	roleTable               = "role"
)

func MustReinstall(db *sqlx.DB) {
	dropTables(db)
	createTables(db)
	addAccount(db)
}

func dropTables(db *sqlx.DB) {
	for _, tableName := range [...]string{
		accountTable,
		accountCredentialsTable,
		resourceTable,
		permissionTable,
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
}
