package info

import (
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"mini-roles-backend/source/domains/account/mock"
	"mini-roles-backend/source/domains/account/models"
	"mini-roles-backend/source/domains/account/request"
	"mini-roles-backend/source/domains/account/services/shared"
	sharedError "mini-roles-backend/source/domains/shared/error"
	sharedMock "mini-roles-backend/source/domains/shared/mock"
	"mini-roles-backend/source/domains/shared/services/response_factory"
	"testing"
)

var (
	mockInfoRepository  = &mock.InfoRepository{}
	expectedOkStatus    = response_factory.DefaultResponse().ApplicationStatus()
	expectedErrorStatus = response_factory.ServerError(nil).ApplicationStatus()
)

func init() {
	mockInfoRepository.Reset()
}

func TestService_GetInfoSuccess(t *testing.T) {
	response := New(mockInfoRepository).GetInfo(request.GetInfoRequest{
		AccountId: sharedMock.ExistsAccountId,
	})

	expectedInfo, _ := mockInfoRepository.FetchInfo(sharedMock.ExistsAccountId)

	assert.Equal(t, expectedOkStatus, response.ApplicationStatus())
	assert.True(t, response.HasData())
	assert.Equal(t, expectedInfo, response.Data())
}

func TestService_GetInfoValidationError(t *testing.T) {
	req := request.GetInfoRequest{}

	response := New(mockInfoRepository).GetInfo(req)

	assert.Equal(t, expectedErrorStatus, response.ApplicationStatus())
	assert.Equal(t, sharedError.ValidationErrorCode, response.Data().(sharedError.ServiceError).Code)
	assert.Equal(t, validator.New().Struct(req).Error(), response.Data().(sharedError.ServiceError).Description)
}

func TestService_GetInfoServerError(t *testing.T) {
	response := New(mockInfoRepository).GetInfo(request.GetInfoRequest{
		AccountId: sharedMock.BadAccountId,
	})

	assert.Equal(t, expectedErrorStatus, response.ApplicationStatus())
	assert.True(t, response.HasData())
	assert.Equal(t, sharedError.ServerErrorCode, response.Data().(sharedError.ServiceError).Code)
	assert.Equal(t, sharedError.ServerErrorDescription, response.Data().(sharedError.ServiceError).Description)
}

func TestService_UpdateCredentialsSuccess(t *testing.T) {
	defer mockInfoRepository.Reset()
	req := request.UpdateCredentialsRequest{
		AccountId: sharedMock.ExistsAccountId,
		Credentials: models.UpdateAccountCredentials{
			Login:    "some-login",
			Password: "some-password",
		},
	}

	response := New(mockInfoRepository).UpdateCredentials(req)

	req.Credentials.Password = shared.MakePassword(models.AccountCredentials(req.Credentials))
	assert.Equal(t, expectedOkStatus, response.ApplicationStatus())
	assert.False(t, response.HasData())
	assert.Equal(t, mockInfoRepository.Credentials(sharedMock.ExistsAccountId), req.Credentials)
}

func TestService_UpdateCredentialsWithoutPasswordSuccess(t *testing.T) {
	defer mockInfoRepository.Reset()
	req := request.UpdateCredentialsRequest{
		AccountId: sharedMock.ExistsAccountId,
		Credentials: models.UpdateAccountCredentials{
			Login: "some-login",
		},
	}

	response := New(mockInfoRepository).UpdateCredentials(req)

	req.Credentials.Password = shared.MakePassword(models.AccountCredentials(req.Credentials))
	assert.Equal(t, expectedOkStatus, response.ApplicationStatus())
	assert.False(t, response.HasData())
	assert.Equal(t, mockInfoRepository.Info(sharedMock.ExistsAccountId).Login, req.Credentials.Login)
	assert.Equal(t, mockInfoRepository.Credentials(sharedMock.ExistsAccountId).Password, mock.ExistsPassword)
}

func TestService_UpdateCredentialsValidationError(t *testing.T) {
	req := request.UpdateCredentialsRequest{
		AccountId: sharedMock.ExistsAccountId,
	}

	response := New(mockInfoRepository).UpdateCredentials(req)

	assert.Equal(t, expectedErrorStatus, response.ApplicationStatus())
	assert.Equal(t, sharedError.ValidationErrorCode, response.Data().(sharedError.ServiceError).Code)
	assert.Equal(t, validator.New().Struct(req).Error(), response.Data().(sharedError.ServiceError).Description)
}

func TestService_UpdateCredentialsLoginAlreadyExists(t *testing.T) {
	req := request.UpdateCredentialsRequest{
		AccountId: sharedMock.ExistsAccountId,
		Credentials: models.UpdateAccountCredentials{
			Login:    mock.ExistsLogin,
			Password: "some-password",
		},
	}

	response := New(mockInfoRepository).UpdateCredentials(req)

	assert.Equal(t, expectedErrorStatus, response.ApplicationStatus())
	assert.Equal(t, shared.LoginAlreadyExistsCode, response.Data().(sharedError.ServiceError).Code)
	assert.Equal(t, shared.LoginAlreadyExistsDescription, response.Data().(sharedError.ServiceError).Description)
}

func TestService_UpdateCredentialsServerError(t *testing.T) {
	req := request.UpdateCredentialsRequest{
		AccountId: sharedMock.BadAccountId,
		Credentials: models.UpdateAccountCredentials{
			Login:    "some-login",
			Password: "some-password",
		},
	}

	response := New(mockInfoRepository).UpdateCredentials(req)

	assert.Equal(t, expectedErrorStatus, response.ApplicationStatus())
	assert.True(t, response.HasData())
	assert.Equal(t, sharedError.ServerErrorCode, response.Data().(sharedError.ServiceError).Code)
	assert.Equal(t, sharedError.ServerErrorDescription, response.Data().(sharedError.ServiceError).Description)
}
