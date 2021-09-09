package validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"mini-roles-backend/source/domains/role/request"
	sharedError "mini-roles-backend/source/domains/shared/error"
	sharedMock "mini-roles-backend/source/domains/shared/mock"
	sharedModels "mini-roles-backend/source/domains/shared/models"
	"testing"
)

func TestMakeErrorResponseInvalidRequest(t *testing.T) {
	req := request.CreateRole{
		AccountId: sharedMock.ExistsAccountId,
		Role:      sharedModels.Role{},
	}

	response := MakeErrorResponse(req)

	assert.NotNil(t, response)
	assert.False(t, response.IsOk())
	assert.True(t, response.HasData())
	assert.Equal(t, sharedError.ValidationErrorCode, response.Data().(sharedError.ServiceError).Code)
	assert.Equal(
		t,
		MakeFailedValidationDescription(validator.New().Struct(req)),
		response.Data().(sharedError.ServiceError).Description,
	)
}

func TestMakeErrorResponseValidRequest(t *testing.T) {
	req := request.CreateRole{
		AccountId: sharedMock.ExistsAccountId,
		Role: sharedModels.Role{
			Id:        "some-id",
			VersionId: sharedMock.ExistsRolesVersionId,
		},
	}

	response := MakeErrorResponse(req)

	assert.Nil(t, response)
}
