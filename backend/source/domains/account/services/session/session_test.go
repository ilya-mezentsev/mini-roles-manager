package session

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"mini-roles-backend/source/config"
	"mini-roles-backend/source/domains/account/mock"
	"mini-roles-backend/source/domains/account/models"
	"mini-roles-backend/source/domains/account/request"
	sharedError "mini-roles-backend/source/domains/shared/error"
	sharedMock "mini-roles-backend/source/domains/shared/mock"
	"mini-roles-backend/source/domains/shared/services/response_factory"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	configRepository      = config.Default()
	sessionRepositoryMock = &mock.SessionRepository{}
	expectedOkStatus      = response_factory.DefaultResponse().ApplicationStatus()
	expectedErrorStatus   = response_factory.ServerError(nil).ApplicationStatus()
	service               = New(sessionRepositoryMock, configRepository)
)

func init() {
	sessionRepositoryMock.Reset()
}

func TestService_CreateSessionSuccess(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = mock.CreateSimpleRequest()
	u := models.AccountCredentials{
		Login:    "some-login",
		Password: "some-password",
	}

	response := service.CreateSession(request.CreateSession{
		Context:     c,
		Credentials: u,
	})

	assert.Equal(t, expectedOkStatus, response.ApplicationStatus())
	assert.True(t, response.HasData())
	assert.Equal(t, sharedMock.ExistsAccountId, response.GetData().(models.AccountSession).Id)
	assert.True(t, strings.Contains(
		w.Header().Get("Set-Cookie"),
		fmt.Sprintf("%s=%s", cookieTokenKey, sharedMock.ExistsAccountId),
	))
}

func TestService_CreateSessionValidationError(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = mock.CreateSimpleRequest()
	req := request.CreateSession{
		Context: c,
	}

	response := service.CreateSession(req)

	assert.Equal(t, expectedErrorStatus, response.ApplicationStatus())
	assert.Equal(t, sharedError.ValidationErrorCode, response.GetData().(sharedError.ServiceError).Code)
	assert.Equal(t, validator.New().Struct(req).Error(), response.GetData().(sharedError.ServiceError).Description)
}

func TestService_CreateSessionMissedLoginError(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = mock.CreateSimpleRequest()
	u := models.AccountCredentials{
		Login:    mock.MissedLogin,
		Password: "some-password",
	}

	response := service.CreateSession(request.CreateSession{
		Context:     c,
		Credentials: u,
	})

	assert.Equal(t, expectedErrorStatus, response.ApplicationStatus())
	assert.Equal(t, credentialsNotFoundCode, response.GetData().(sharedError.ServiceError).Code)
	assert.Equal(t, credentialsNotFoundDescription, response.GetData().(sharedError.ServiceError).Description)
}

func TestService_CreateSessionServerError(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = mock.CreateSimpleRequest()
	u := models.AccountCredentials{
		Login:    mock.BadLogin,
		Password: "some-password",
	}

	response := service.CreateSession(request.CreateSession{
		Context:     c,
		Credentials: u,
	})

	assert.Equal(t, expectedErrorStatus, response.ApplicationStatus())
	assert.Equal(t, sharedError.ServerErrorCode, response.GetData().(sharedError.ServiceError).Code)
	assert.Equal(t, sharedError.ServerErrorDescription, response.GetData().(sharedError.ServiceError).Description)
}

func TestService_CreateSessionSuccessExists(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = mock.CreateRequestWithCookie(cookieTokenKey, string(sharedMock.ExistsAccountId))

	response := service.GetSession(request.GetSession{
		Context: c,
	})

	assert.Equal(t, expectedOkStatus, response.ApplicationStatus())
	assert.True(t, response.HasData())
	assert.Equal(t, sharedMock.ExistsAccountId, response.GetData().(models.AccountSession).Id)
	assert.True(t, strings.Contains(
		w.Header().Get("Set-Cookie"),
		fmt.Sprintf("%s=%s", cookieTokenKey, sharedMock.ExistsAccountId),
	))
}

func TestService_GetSessionNotExists(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = mock.CreateRequestWithCookie(cookieTokenKey, "some-token")

	response := service.GetSession(request.GetSession{
		Context: c,
	})

	assert.Equal(t, expectedOkStatus, response.ApplicationStatus())
	assert.False(t, response.HasData())
}

func TestService_GetSessionNoCookie(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = mock.CreateSimpleRequest()

	response := service.GetSession(request.GetSession{
		Context: c,
	})

	assert.Equal(t, expectedOkStatus, response.ApplicationStatus())
	assert.False(t, response.HasData())
}

func TestService_GetSessionServerError(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = mock.CreateRequestWithCookie(cookieTokenKey, string(sharedMock.BadAccountId))

	response := service.GetSession(request.GetSession{
		Context: c,
	})

	assert.Equal(t, expectedErrorStatus, response.ApplicationStatus())
	assert.Equal(t, sharedError.ServerErrorCode, response.GetData().(sharedError.ServiceError).Code)
	assert.Equal(t, sharedError.ServerErrorDescription, response.GetData().(sharedError.ServiceError).Description)
}

func TestService_DeleteSessionSuccess(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = mock.CreateRequestWithCookie(cookieTokenKey, string(sharedMock.ExistsAccountId))

	response := service.DeleteSession(request.DeleteSession{
		Context: c,
	})

	assert.Equal(t, expectedOkStatus, response.ApplicationStatus())
	assert.False(t, response.HasData())
	assert.True(t, strings.Contains(
		w.Header().Get("Set-Cookie"),
		fmt.Sprintf("%s=", cookieTokenKey),
	))
}
