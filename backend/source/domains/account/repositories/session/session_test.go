package session

import (
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

func TestRepository_GetSessionSuccess(t *testing.T) {
	accountSession, err := repository.GetSession(models.AccountCredentials{
		Login:    sharedMock.ExistsLogin,
		Password: sharedMock.ExistsPassword,
	})

	assert.Nil(t, err)
	assert.Equal(t, sharedMock.ExistsAccountId, accountSession.Id)
}

func TestRepository_GetSessionNoCredentialsFound(t *testing.T) {
	_, err := repository.GetSession(models.AccountCredentials{
		Login:    "foo-bar",
		Password: sharedMock.ExistsPassword,
	})

	assert.True(t, errors.As(err, &sharedError.EntryNotFound{}))
}

func TestRepository_SessionExistsTrue(t *testing.T) {
	sessionExists, err := repository.SessionExists(models.AccountSession{
		Id: sharedMock.ExistsAccountId,
	})

	assert.Nil(t, err)
	assert.True(t, sessionExists)
}

func TestRepository_SessionExistsFalse(t *testing.T) {
	sessionExists, err := repository.SessionExists(models.AccountSession{
		Id: "foo-bar",
	})

	assert.Nil(t, err)
	assert.False(t, sessionExists)
}
