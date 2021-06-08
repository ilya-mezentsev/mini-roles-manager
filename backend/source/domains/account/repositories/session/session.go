package session

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"mini-roles-backend/source/domains/account/models"
	"mini-roles-backend/source/domains/account/spec"
	sharedError "mini-roles-backend/source/domains/shared/error"
	sharedModels "mini-roles-backend/source/domains/shared/models"
)

const (
	checkSessionExistenceQuery = `select 1 from account where hash = $1`
	getSessionQuery            = `
	select trim(account_hash) account_hash from account_credentials
	where login = $1 and password = $2`
)

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repository {
	return Repository{db}
}

func (r Repository) GetSession(spec spec.SessionWithCredentials) (models.AccountSession, error) {
	var accountId sharedModels.AccountId
	err := r.db.Get(&accountId, getSessionQuery, spec.Credentials.Login, spec.Credentials.Password)
	if err == sql.ErrNoRows {
		err = sharedError.EntryNotFound{}
	}

	return models.AccountSession{Id: accountId}, err
}

func (r Repository) SessionExists(spec spec.SessionWithId) (bool, error) {
	var sessionExists bool
	err := r.db.Get(&sessionExists, checkSessionExistenceQuery, spec.Id)
	if err == sql.ErrNoRows {
		sessionExists = false
		err = nil
	}

	return sessionExists, err
}
