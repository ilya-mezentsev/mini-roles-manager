package info

import (
	"github.com/jmoiron/sqlx"
	"mini-roles-backend/source/domains/account/models"
	sharedError "mini-roles-backend/source/domains/shared/error"
	sharedModels "mini-roles-backend/source/domains/shared/models"
	sharedSpec "mini-roles-backend/source/domains/shared/spec"
)

const (
	fetchAccountInfoQuery = `
	select
		a.hash api_key,
		a.created_at created,
		trim(ac.login) login
	from account a
	join account_credentials ac on a.hash = ac.account_hash
	where a.hash = $1`

	updateCredentialsQuery = `update account_credentials set login = $2, password = $3 where account_hash = $1`

	updateLoginQuery = `update account_credentials set login = $2 where account_hash = $1`
)

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repository {
	return Repository{db}
}

func (r Repository) FetchInfo(spec sharedSpec.AccountWithId) (models.AccountInfo, error) {
	var info models.AccountInfo
	err := r.db.Get(&info, fetchAccountInfoQuery, spec.AccountId)

	return info, err
}

func (r Repository) UpdateCredentials(accountId sharedModels.AccountId, credentials models.UpdateAccountCredentials) error {
	_, err := r.db.Exec(updateCredentialsQuery, accountId, credentials.Login, credentials.Password)

	return r.parseUpdateError(err)
}

func (r Repository) parseUpdateError(err error) error {
	if sharedError.IsDuplicateKey(err) {
		err = sharedError.DuplicateUniqueKey{}
	}

	return err
}

func (r Repository) UpdateLogin(accountId sharedModels.AccountId, newLogin string) error {
	_, err := r.db.Exec(updateLoginQuery, accountId, newLogin)

	return r.parseUpdateError(err)
}
