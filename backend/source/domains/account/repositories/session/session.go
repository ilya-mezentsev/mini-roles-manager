package session

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"mini-roles-backend/source/domains/account/models"
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

func (r Repository) GetSession(credentials models.AccountCredentials) (models.AccountSession, error) {
	var accountId sharedModels.AccountId
	err := r.db.Get(&accountId, getSessionQuery, credentials.Login, credentials.Password)
	if err == sql.ErrNoRows {
		err = sharedError.EntryNotFound{}
	}

	return models.AccountSession{Id: accountId}, err
}

func (r Repository) SessionExists(session models.AccountSession) (bool, error) {
	var sessionExists bool
	err := r.db.Get(&sessionExists, checkSessionExistenceQuery, session.Id)
	if err == sql.ErrNoRows {
		sessionExists = false
		err = nil
	}

	return sessionExists, err
}
