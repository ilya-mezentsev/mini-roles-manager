package session_check

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"mini-roles-backend/source/domains/account/mock"
	"mini-roles-backend/source/domains/account/request"
	"mini-roles-backend/source/domains/account/services/shared"
	sharedError "mini-roles-backend/source/domains/shared/error"
	sharedMock "mini-roles-backend/source/domains/shared/mock"
	"mini-roles-backend/source/domains/shared/services/response_factory"
	sharedKeys "mini-roles-backend/source/shared/keys"
	"net/http/httptest"
	"testing"
)

var (
	sessionRepositoryMock = &mock.SessionRepository{}
	expectedErrorStatus   = response_factory.EmptyServerError().ApplicationStatus()
	service               = New(sessionRepositoryMock)
)

func init() {
	sessionRepositoryMock.Reset()
}

func TestService_CheckSessionFromCookieSuccess(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = mock.CreateRequestWithCookie(shared.CookieTokenKey, string(sharedMock.ExistsAccountId))

	response := service.CheckSessionFromCookie(request.SessionExists{
		Context: c,
	})

	assert.Nil(t, response)
	assert.Equal(t, string(sharedMock.ExistsAccountId), c.GetString(sharedKeys.ContextTokenKey))
}

func TestService_CheckSessionFromHeaderSuccess(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = mock.CreateRequestWithHeader(headerTokenKey, string(sharedMock.ExistsAccountId))

	response := service.CheckSessionFromHeader(request.SessionExists{
		Context: c,
	})

	assert.Nil(t, response)
	assert.Equal(t, string(sharedMock.ExistsAccountId), c.GetString(sharedKeys.ContextTokenKey))
}

func TestService_CheckSessionFromCookieNoCookieError(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = mock.CreateSimpleRequest()

	response := service.CheckSessionFromCookie(request.SessionExists{
		Context: c,
	})

	assert.Equal(t, expectedErrorStatus, response.ApplicationStatus())
	assert.Equal(t, missedTokenInCookiesCode, response.Data().(sharedError.ServiceError).Code)
	assert.Equal(t, missedTokenInCookiesDescription, response.Data().(sharedError.ServiceError).Description)
}

func TestService_CheckSessionFromHeaderNotTokenError(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = mock.CreateSimpleRequest()

	response := service.CheckSessionFromHeader(request.SessionExists{
		Context: c,
	})

	assert.Equal(t, expectedErrorStatus, response.ApplicationStatus())
	assert.Equal(t, missedTokenInHeadersCode, response.Data().(sharedError.ServiceError).Code)
	assert.Equal(t, missedTokenInHeadersDescription, response.Data().(sharedError.ServiceError).Description)
}

func TestService_CheckSessionFromCookieNoAccountByToken(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = mock.CreateRequestWithCookie(shared.CookieTokenKey, "some-token")

	response := service.CheckSessionFromCookie(request.SessionExists{
		Context: c,
	})

	assert.Equal(t, expectedErrorStatus, response.ApplicationStatus())
	assert.Equal(t, noAccountByTokenCode, response.Data().(sharedError.ServiceError).Code)
	assert.Equal(t, noAccountByTokenDescription, response.Data().(sharedError.ServiceError).Description)
}

func TestService_CheckSessionFromCookieServerError(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = mock.CreateRequestWithCookie(shared.CookieTokenKey, string(sharedMock.BadAccountId))

	response := service.CheckSessionFromCookie(request.SessionExists{
		Context: c,
	})

	assert.Equal(t, expectedErrorStatus, response.ApplicationStatus())
	assert.Equal(t, sharedError.ServerErrorCode, response.Data().(sharedError.ServiceError).Code)
	assert.Equal(t, sharedError.ServerErrorDescription, response.Data().(sharedError.ServiceError).Description)
}
