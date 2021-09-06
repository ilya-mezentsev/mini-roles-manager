package roles_version

import (
	"github.com/go-playground/validator/v10"
	responseFactory "github.com/ilya-mezentsev/response-factory"
	"github.com/stretchr/testify/assert"
	"mini-roles-backend/source/domains/role/request"
	sharedError "mini-roles-backend/source/domains/shared/error"
	sharedMock "mini-roles-backend/source/domains/shared/mock"
	sharedModels "mini-roles-backend/source/domains/shared/models"
	"mini-roles-backend/source/domains/shared/services/validation"
	sharedSpec "mini-roles-backend/source/domains/shared/spec"
	"testing"
)

var (
	mockRepository                 = &sharedMock.RolesVersionRepository{}
	mockPermissionCacheInvalidator = &sharedMock.InMemoryCacheInvalidator{}
	service                        = New(mockRepository, mockPermissionCacheInvalidator)
	expectedOkStatus               = responseFactory.DefaultResponse().ApplicationStatus()
	expectedErrorStatus            = responseFactory.EmptyServerError().ApplicationStatus()
)

func init() {
	reset()
}

func reset() {
	mockRepository.Reset()
	mockPermissionCacheInvalidator.Reset()
}

func TestService_CreateRolesVersionSuccess(t *testing.T) {
	defer reset()

	newRolesVersion := sharedModels.RolesVersion{
		Id: "some-new-roles-version",
	}

	response := service.CreateRolesVersion(request.CreateRolesVersion{
		AccountId:    sharedMock.ExistsAccountId,
		RolesVersion: newRolesVersion,
	})

	assert.True(t, mockRepository.Has(newRolesVersion))
	assert.Equal(t, expectedOkStatus, response.ApplicationStatus())
	assert.False(t, response.HasData())
	assert.True(t, mockPermissionCacheInvalidator.InvalidateCalledWith(sharedMock.ExistsAccountId))
}

func TestService_CreateRolesVersionDuplicateKeyError(t *testing.T) {
	defer reset()

	newRolesVersion := sharedModels.RolesVersion{
		Id: sharedMock.ExistsRolesVersionId,
	}

	response := service.CreateRolesVersion(request.CreateRolesVersion{
		AccountId:    sharedMock.ExistsAccountId,
		RolesVersion: newRolesVersion,
	})

	assert.Equal(t, expectedErrorStatus, response.ApplicationStatus())
	assert.True(t, response.HasData())
	assert.Equal(t, rolesVersionExistsCode, response.Data().(sharedError.ServiceError).Code)
	assert.Equal(t, rolesVersionExistsDescription, response.Data().(sharedError.ServiceError).Description)
}

func TestService_CreateRolesVersionValidationError(t *testing.T) {
	defer reset()

	newRolesVersion := sharedModels.RolesVersion{}
	req := request.CreateRolesVersion{
		AccountId:    sharedMock.ExistsAccountId,
		RolesVersion: newRolesVersion,
	}

	response := service.CreateRolesVersion(req)

	assert.False(t, mockRepository.Has(newRolesVersion))
	assert.Equal(t, expectedErrorStatus, response.ApplicationStatus())
	assert.True(t, response.HasData())
	assert.Equal(t, sharedError.ValidationErrorCode, response.Data().(sharedError.ServiceError).Code)
	assert.Equal(
		t,
		validation.MakeFailedValidationDescription(validator.New().Struct(req)),
		response.Data().(sharedError.ServiceError).Description,
	)
}

func TestService_CreateRolesVersionDBError(t *testing.T) {
	defer reset()

	newRolesVersion := sharedModels.RolesVersion{
		Id: "some-new-roles-version",
	}

	response := service.CreateRolesVersion(request.CreateRolesVersion{
		AccountId:    sharedMock.BadAccountId,
		RolesVersion: newRolesVersion,
	})

	assert.False(t, mockRepository.Has(newRolesVersion))
	assert.Equal(t, expectedErrorStatus, response.ApplicationStatus())
	assert.Equal(t, sharedError.ServerErrorCode, response.Data().(sharedError.ServiceError).Code)
	assert.Equal(t, sharedError.ServerErrorDescription, response.Data().(sharedError.ServiceError).Description)
}

func TestService_RolesVersionListSuccess(t *testing.T) {
	expectedRolesVersionList, err := mockRepository.List(sharedSpec.AccountWithId{
		AccountId: sharedMock.ExistsAccountId,
	})
	assert.Nil(t, err)
	assert.NotEmpty(t, expectedRolesVersionList)

	response := service.RolesVersionList(request.RolesVersionList{
		AccountId: sharedMock.ExistsAccountId,
	})

	assert.Equal(t, expectedOkStatus, response.ApplicationStatus())
	assert.True(t, response.HasData())
	assert.Equal(t, expectedRolesVersionList, response.Data())
}

func TestService_RolesVersionListValidationError(t *testing.T) {
	req := request.RolesVersionList{}

	response := service.RolesVersionList(req)

	assert.Equal(t, expectedErrorStatus, response.ApplicationStatus())
	assert.True(t, response.HasData())
	assert.Equal(t, sharedError.ValidationErrorCode, response.Data().(sharedError.ServiceError).Code)
	assert.Equal(
		t,
		validation.MakeFailedValidationDescription(validator.New().Struct(req)),
		response.Data().(sharedError.ServiceError).Description,
	)
}

func TestService_RolesVersionListDBError(t *testing.T) {
	response := service.RolesVersionList(request.RolesVersionList{
		AccountId: sharedMock.BadAccountId,
	})

	assert.Equal(t, expectedErrorStatus, response.ApplicationStatus())
	assert.Equal(t, sharedError.ServerErrorCode, response.Data().(sharedError.ServiceError).Code)
	assert.Equal(t, sharedError.ServerErrorDescription, response.Data().(sharedError.ServiceError).Description)
}

func TestService_UpdateRolesVersionSuccess(t *testing.T) {
	defer reset()

	updatingRolesVersion := sharedModels.RolesVersion{
		Id:    sharedMock.ExistsRolesVersionId,
		Title: "Some-Title",
	}

	response := service.UpdateRolesVersion(request.UpdateRolesVersion{
		AccountId:    sharedMock.ExistsAccountId,
		RolesVersion: updatingRolesVersion,
	})

	assert.Equal(t, expectedOkStatus, response.ApplicationStatus())
	assert.False(t, response.HasData())
	assert.True(t, mockRepository.Has(updatingRolesVersion))
	assert.True(t, mockPermissionCacheInvalidator.InvalidateCalledWith(sharedMock.ExistsAccountId))
}

func TestService_UpdateRolesVersionValidationError(t *testing.T) {
	req := request.UpdateRolesVersion{}

	response := service.UpdateRolesVersion(req)

	assert.Equal(t, expectedErrorStatus, response.ApplicationStatus())
	assert.True(t, response.HasData())
	assert.Equal(t, sharedError.ValidationErrorCode, response.Data().(sharedError.ServiceError).Code)
	assert.Equal(
		t,
		validation.MakeFailedValidationDescription(validator.New().Struct(req)),
		response.Data().(sharedError.ServiceError).Description,
	)
}

func TestService_UpdateRolesVersionDBError(t *testing.T) {
	updatingRolesVersion := sharedModels.RolesVersion{
		Id:    sharedMock.ExistsRolesVersionId,
		Title: "Some-Title",
	}

	response := service.UpdateRolesVersion(request.UpdateRolesVersion{
		AccountId:    sharedMock.BadAccountId,
		RolesVersion: updatingRolesVersion,
	})

	assert.Equal(t, expectedErrorStatus, response.ApplicationStatus())
	assert.Equal(t, sharedError.ServerErrorCode, response.Data().(sharedError.ServiceError).Code)
	assert.Equal(t, sharedError.ServerErrorDescription, response.Data().(sharedError.ServiceError).Description)
}

func TestService_DeleteRolesVersionSuccess(t *testing.T) {
	defer reset()

	deletingRolesVersion := sharedModels.RolesVersion{
		Id: sharedMock.ExistsRolesVersionId,
	}

	response := service.DeleteRolesVersion(request.DeleteRolesVersion{
		AccountId:      sharedMock.ExistsAccountId,
		RolesVersionId: deletingRolesVersion.Id,
	})

	assert.Equal(t, expectedOkStatus, response.ApplicationStatus())
	assert.False(t, response.HasData())
	assert.False(t, mockRepository.Has(deletingRolesVersion))
	assert.True(t, mockPermissionCacheInvalidator.InvalidateCalledWith(sharedMock.ExistsAccountId))
}

func TestService_DeleteRolesVersionCannotDeleteLast(t *testing.T) {
	defer reset()
	_ = mockRepository.Delete(sharedMock.ExistsAccountId, sharedMock.ExistsRolesVersionId2)

	deletingRolesVersion := sharedModels.RolesVersion{
		Id: sharedMock.ExistsRolesVersionId,
	}

	response := service.DeleteRolesVersion(request.DeleteRolesVersion{
		AccountId:      sharedMock.ExistsAccountId,
		RolesVersionId: deletingRolesVersion.Id,
	})

	assert.Equal(t, expectedErrorStatus, response.ApplicationStatus())
	assert.True(t, response.HasData())
	assert.Equal(t, cannotDeleteLastRolesVersionCode, response.Data().(sharedError.ServiceError).Code)
	assert.Equal(
		t,
		cannotDeleteLastRolesVersionDescription,
		response.Data().(sharedError.ServiceError).Description,
	)
	assert.True(t, mockRepository.Has(deletingRolesVersion))
}

func TestService_DeleteRolesVersionValidationError(t *testing.T) {
	req := request.DeleteRolesVersion{}

	response := service.DeleteRolesVersion(req)

	assert.Equal(t, expectedErrorStatus, response.ApplicationStatus())
	assert.True(t, response.HasData())
	assert.Equal(t, sharedError.ValidationErrorCode, response.Data().(sharedError.ServiceError).Code)
	assert.Equal(
		t,
		validation.MakeFailedValidationDescription(validator.New().Struct(req)),
		response.Data().(sharedError.ServiceError).Description,
	)
}

func TestService_DeleteRolesVersionDBErrorInDeleting(t *testing.T) {
	response := service.DeleteRolesVersion(request.DeleteRolesVersion{
		AccountId:      sharedMock.ExistsAccountId,
		RolesVersionId: sharedMock.BadRolesVersionId,
	})

	assert.Equal(t, expectedErrorStatus, response.ApplicationStatus())
	assert.Equal(t, sharedError.ServerErrorCode, response.Data().(sharedError.ServiceError).Code)
	assert.Equal(t, sharedError.ServerErrorDescription, response.Data().(sharedError.ServiceError).Description)
}

func TestService_DeleteRolesVersionDBErrorInFetching(t *testing.T) {
	deletingRolesVersion := sharedModels.RolesVersion{
		Id: sharedMock.ExistsRolesVersionId,
	}

	response := service.DeleteRolesVersion(request.DeleteRolesVersion{
		AccountId:      sharedMock.BadAccountId,
		RolesVersionId: deletingRolesVersion.Id,
	})

	assert.Equal(t, expectedErrorStatus, response.ApplicationStatus())
	assert.Equal(t, sharedError.ServerErrorCode, response.Data().(sharedError.ServiceError).Code)
	assert.Equal(t, sharedError.ServerErrorDescription, response.Data().(sharedError.ServiceError).Description)
	assert.True(t, mockRepository.Has(deletingRolesVersion))
}
