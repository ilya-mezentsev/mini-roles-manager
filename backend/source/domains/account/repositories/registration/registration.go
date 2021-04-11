package registration

import (
	"github.com/jmoiron/sqlx"
	"mini-roles-backend/source/domains/account/models"
	sharedError "mini-roles-backend/source/domains/shared/error"
	sharedRepositories "mini-roles-backend/source/domains/shared/repositories"
)

const (
	createSessionQuery     = `insert into account(hash) values($1)`
	createCredentialsQuery = `
	insert into account_credentials(login, password, account_hash)
	values(:login, :password, :account_hash)`
)

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repository {
	return Repository{db}
}

func (r Repository) Register(
	session models.AccountSession,
	credentials models.AccountCredentials,
) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}

	_, err = tx.Exec(createSessionQuery, session.Id)
	if err != nil {
		return err
	}

	_, err = tx.NamedExec(createCredentialsQuery, map[string]interface{}{
		"login":        credentials.Login,
		"password":     credentials.Password,
		"account_hash": session.Id,
	})
	if err != nil {
		if sharedRepositories.IsDuplicateKey(err) {
			err = sharedError.DuplicateUniqueKey{}
		}

		_ = tx.Rollback()

		return err
	}

	return tx.Commit()
}
