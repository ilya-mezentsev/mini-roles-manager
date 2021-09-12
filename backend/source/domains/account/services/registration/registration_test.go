package registration

import (
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"mini-roles-backend/source/domains/account/mock"
	"mini-roles-backend/source/domains/account/models"
	"mini-roles-backend/source/domains/account/request"
	"mini-roles-backend/source/domains/account/services/shared"
	sharedError "mini-roles-backend/source/domains/shared/error"
	sharedMock "mini-roles-backend/source/domains/shared/mock"
	sharedModels "mini-roles-backend/source/domains/shared/models"
	"mini-roles-backend/source/domains/shared/services/validation"
	"testing"
)

var (
	mockRegistrationRepository = &mock.RegistrationRepository{}
	mockRolesVersionRepository = &sharedMock.RolesVersionRepository{}
	service                    = New(mockRegistrationRepository, mockRolesVersionRepository)
)

func init() {
	reset()
}

func reset() {
	mockRegistrationRepository.Reset()
	mockRolesVersionRepository.Reset()
}

func TestService_RegisterSuccess(t *testing.T) {
	defer reset()
	credentials := models.AccountCredentials{
		Login:    "SomeLogin",
		Password: "SomePassword",
	}

	response := service.Register(request.Registration{
		Credentials: credentials,
	})
	credentials.Password = shared.MakePassword(credentials)

	assert.True(t, mockRolesVersionRepository.Has(sharedModels.RolesVersion{
		Id: defaultRolesVersionId,
	}))
	assert.Contains(t, mockRegistrationRepository.GetAll(), credentials)
	assert.True(t, response.IsOk())
	assert.False(t, response.HasData())
}

func TestService_RegisterValidationError(t *testing.T) {
	req := request.Registration{}

	response := service.Register(req)

	assert.False(t, response.IsOk())
	assert.True(t, response.HasData())
	assert.Equal(t, sharedError.ValidationErrorCode, response.Data().(sharedError.ServiceError).Code)
	assert.Equal(
		t,
		validation.MakeFailedValidationDescription(validator.New().Struct(req)),
		response.Data().(sharedError.ServiceError).Description,
	)
}

func TestService_RegisterLoginExistsError(t *testing.T) {
	response := service.Register(request.Registration{
		Credentials: models.AccountCredentials{
			Login:    mock.ExistsLogin,
			Password: "SomePassword",
		},
	})

	assert.False(t, response.IsOk())
	assert.True(t, response.HasData())
	assert.Equal(t, shared.LoginAlreadyExistsCode, response.Data().(sharedError.ServiceError).Code)
	assert.Equal(t, shared.LoginAlreadyExistsDescription, response.Data().(sharedError.ServiceError).Description)
}

func TestService_RegisterServerError(t *testing.T) {
	response := service.Register(request.Registration{
		Credentials: models.AccountCredentials{
			Login:    mock.BadLogin,
			Password: "SomePassword",
		},
	})

	assert.False(t, response.IsOk())
	assert.True(t, response.HasData())
	assert.Equal(t, sharedError.ServerErrorCode, response.Data().(sharedError.ServiceError).Code)
	assert.Equal(t, sharedError.ServerErrorDescription, response.Data().(sharedError.ServiceError).Description)
}

func TestService_RegisterServerError2(t *testing.T) {
	defer reset()
	credentials := models.AccountCredentials{
		Login:    "SomeLogin",
		Password: "SomePassword",
	}

	response := New(mockRegistrationRepository, mock.FailingRolesVersionCreatorRepository{}).Register(request.Registration{
		Credentials: credentials,
	})
	credentials.Password = shared.MakePassword(credentials)

	assert.False(t, response.IsOk())
	assert.True(t, response.HasData())
	assert.Equal(t, sharedError.ServerErrorCode, response.Data().(sharedError.ServiceError).Code)
	assert.Equal(t, sharedError.ServerErrorDescription, response.Data().(sharedError.ServiceError).Description)
}
