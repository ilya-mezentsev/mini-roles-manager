package registration

import (
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"mini-roles-backend/source/domains/account/mock"
	"mini-roles-backend/source/domains/account/models"
	"mini-roles-backend/source/domains/account/request"
	"mini-roles-backend/source/domains/account/services/shared"
	sharedError "mini-roles-backend/source/domains/shared/error"
	"mini-roles-backend/source/domains/shared/services/response_factory"
	"mini-roles-backend/source/domains/shared/services/validation"
	"testing"
)

var (
	mockRegistrationRepository = &mock.RegistrationRepository{}
	expectedOkStatus           = response_factory.DefaultResponse().ApplicationStatus()
	expectedErrorStatus        = response_factory.ServerError(nil).ApplicationStatus()
)

func init() {
	mockRegistrationRepository.Reset()
}

func TestService_RegisterSuccess(t *testing.T) {
	defer mockRegistrationRepository.Reset()
	credentials := models.AccountCredentials{
		Login:    "SomeLogin",
		Password: "SomePassword",
	}

	response := New(mockRegistrationRepository).Register(request.Registration{
		Credentials: credentials,
	})
	credentials.Password = shared.MakePassword(credentials)

	assert.Contains(t, mockRegistrationRepository.GetAll(), credentials)
	assert.Equal(t, expectedOkStatus, response.ApplicationStatus())
	assert.False(t, response.HasData())
}

func TestService_RegisterValidationError(t *testing.T) {
	req := request.Registration{}

	response := New(mockRegistrationRepository).Register(req)

	assert.Equal(t, expectedErrorStatus, response.ApplicationStatus())
	assert.True(t, response.HasData())
	assert.Equal(t, sharedError.ValidationErrorCode, response.Data().(sharedError.ServiceError).Code)
	assert.Equal(
		t,
		validation.MakeFailedValidationDescription(validator.New().Struct(req)),
		response.Data().(sharedError.ServiceError).Description,
	)
}

func TestService_RegisterLoginExistsError(t *testing.T) {
	response := New(mockRegistrationRepository).Register(request.Registration{
		Credentials: models.AccountCredentials{
			Login:    mock.ExistsLogin,
			Password: "SomePassword",
		},
	})

	assert.Equal(t, expectedErrorStatus, response.ApplicationStatus())
	assert.True(t, response.HasData())
	assert.Equal(t, shared.LoginAlreadyExistsCode, response.Data().(sharedError.ServiceError).Code)
	assert.Equal(t, shared.LoginAlreadyExistsDescription, response.Data().(sharedError.ServiceError).Description)
}

func TestService_RegisterServerError(t *testing.T) {
	response := New(mockRegistrationRepository).Register(request.Registration{
		Credentials: models.AccountCredentials{
			Login:    mock.BadLogin,
			Password: "SomePassword",
		},
	})

	assert.Equal(t, expectedErrorStatus, response.ApplicationStatus())
	assert.True(t, response.HasData())
	assert.Equal(t, sharedError.ServerErrorCode, response.Data().(sharedError.ServiceError).Code)
	assert.Equal(t, sharedError.ServerErrorDescription, response.Data().(sharedError.ServiceError).Description)
}
