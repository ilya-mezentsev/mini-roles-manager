package registration

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"mini-roles-backend/source/config"
	"mini-roles-backend/source/db/connection"
	"mini-roles-backend/source/domains/account/models"
	sharedError "mini-roles-backend/source/domains/shared/error"
	sharedMock "mini-roles-backend/source/domains/shared/mock"
	"testing"
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

func TestRepository_RegisterSuccess(t *testing.T) {
	defer sharedMock.MustReinstall(db)

	err := repository.Register(
		models.AccountSession{Id: "some-id"},
		models.AccountCredentials{
			Login:    "some-login",
			Password: "some-password",
		},
	)
	assert.Nil(t, err)

	var accountHashCreated bool
	_ = db.Get(&accountHashCreated, `select 1 from account where hash = 'some-id'`)
	assert.True(t, accountHashCreated)

	var credentialsCreated bool
	_ = db.Get(
		&credentialsCreated,
		`select 1 from account_credentials where login = 'some-login' and password = 'some-password'`,
	)
	assert.True(t, credentialsCreated)
}

func TestRepository_RegisterDuplicateAccountHash(t *testing.T) {
	defer sharedMock.MustReinstall(db)

	err := repository.Register(
		models.AccountSession{Id: sharedMock.ExistsAccountId},
		models.AccountCredentials{
			Login:    "some-login",
			Password: "some-password",
		},
	)
	assert.NotNil(t, err)

	var credentialsCreated bool
	err = db.Get(
		&credentialsCreated,
		`select 1 from account_credentials where login = 'some-login' and password = 'some-password'`,
	)
	assert.Equal(t, sql.ErrNoRows, err)
	assert.False(t, credentialsCreated)
}

func TestRepository_RegisterDuplicateLogin(t *testing.T) {
	defer sharedMock.MustReinstall(db)

	err := repository.Register(
		models.AccountSession{Id: "some-id"},
		models.AccountCredentials{
			Login:    sharedMock.ExistsLogin,
			Password: "some-password",
		},
	)
	assert.True(t, errors.As(err, &sharedError.DuplicateUniqueKey{}))

	var accountHashCreated bool
	err = db.Get(&accountHashCreated, `select 1 from account where hash = 'some-id'`)
	assert.Equal(t, sql.ErrNoRows, err)
	assert.False(t, accountHashCreated)
}

func TestRepository_RegisterNoConnection(t *testing.T) {
	defer func() {
		db = sqlx.MustOpen(
			"postgres",
			connection.BuildPostgresString(config.Default()),
		)
		sharedMock.MustReinstall(db)
	}()

	_ = db.Close()

	err := repository.Register(
		models.AccountSession{Id: "some-id"},
		models.AccountCredentials{
			Login:    sharedMock.ExistsLogin,
			Password: "some-password",
		},
	)

	assert.NotNil(t, err)
}
