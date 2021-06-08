package info

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"mini-roles-backend/source/config"
	"mini-roles-backend/source/db/connection"
	"mini-roles-backend/source/domains/account/mock"
	"mini-roles-backend/source/domains/account/models"
	sharedError "mini-roles-backend/source/domains/shared/error"
	sharedMock "mini-roles-backend/source/domains/shared/mock"
	sharedSpec "mini-roles-backend/source/domains/shared/spec"
	"testing"
	"time"
)

var (
	db         *sqlx.DB
	repository Repository
)

func init() {
	db = sqlx.MustOpen(
		"postgres",
		connection.BuildPostgresString(config.Default()),
	)
	repository = New(db)

	sharedMock.MustReinstall(db)
}

func TestRepository_FetchInfoSuccess(t *testing.T) {
	info, err := repository.FetchInfo(sharedSpec.AccountWithId{
		AccountId: sharedMock.ExistsAccountId,
	})

	assert.Nil(t, err)
	assert.Equal(t, mock.ExistsLogin, info.Login)
	assert.Equal(t, sharedMock.ExistsAccountId, info.ApiKey)
	assert.Equal(t, time.Now().Format(mock.Format), info.Created.Format(mock.Format))
}

func TestRepository_UpdateCredentialsSuccess(t *testing.T) {
	defer sharedMock.MustReinstall(db)
	credentials := models.UpdateAccountCredentials{
		Login:    "new-login",
		Password: "new-password",
	}

	err := repository.UpdateCredentials(sharedMock.ExistsAccountId, credentials)
	assert.Nil(t, err)

	var credentialsUpdated bool
	_ = db.Get(
		&credentialsUpdated,
		`select 1 from account_credentials where login = $1 and password = $2`,
		credentials.Login,
		credentials.Password,
	)
	assert.True(t, credentialsUpdated)
}

func TestRepository_UpdateLoginSuccess(t *testing.T) {
	defer sharedMock.MustReinstall(db)

	err := repository.UpdateLogin(sharedMock.ExistsAccountId, "some-new-login")
	assert.Nil(t, err)

	var newLogin string
	_ = db.Get(&newLogin, `select trim(login) login from account_credentials where account_hash = $1`, sharedMock.ExistsAccountId)
	assert.Equal(t, "some-new-login", newLogin)
}

func TestRepository_UpdateCredentialsDuplicateLogin(t *testing.T) {
	defer sharedMock.MustReinstall(db)
	db.MustExec(`insert into account_credentials(login, account_hash) values('foo-bar', $1)`, sharedMock.ExistsAccountId)

	credentials := models.UpdateAccountCredentials{
		Login:    "foo-bar",
		Password: "new-password",
	}

	err := repository.UpdateCredentials(sharedMock.ExistsAccountId, credentials)
	assert.True(t, errors.As(err, &sharedError.DuplicateUniqueKey{}))
}
