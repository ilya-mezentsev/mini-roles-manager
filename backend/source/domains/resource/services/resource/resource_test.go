package resource

import (
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"mini-roles-backend/source/domains/resource/mock"
	"mini-roles-backend/source/domains/resource/request"
	sharedError "mini-roles-backend/source/domains/shared/error"
	sharedMock "mini-roles-backend/source/domains/shared/mock"
	sharedModels "mini-roles-backend/source/domains/shared/models"
	"mini-roles-backend/source/domains/shared/services/validation"
	"os"
	"testing"
)

var (
	mockResourceRepository         = &sharedMock.ResourceRepository{}
	mockPermissionRepository       = &mock.PermissionRepository{}
	mockPermissionCacheInvalidator = &sharedMock.InMemoryCacheInvalidator{}
	service                        = New(
		mockResourceRepository,
		mockPermissionRepository,
		mockPermissionCacheInvalidator,
	)
)

func init() {
	resetRepositories()
}

func resetRepositories() {
	mockResourceRepository.Reset()
	mockPermissionRepository.Reset()
	mockPermissionCacheInvalidator.Reset()
}

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}

func TestService_CreateResourceSuccess(t *testing.T) {
	defer resetRepositories()
	newResource := sharedModels.Resource{
		Id:    "new-resource",
		Title: "Some-New-Resource",
	}

	response := service.CreateResource(request.CreateResource{
		AccountId: sharedMock.ExistsAccountId,
		Resource:  newResource,
	})

	assert.True(t, response.IsOk())
	assert.False(t, response.HasData())
	assert.True(t, mockResourceRepository.Has(newResource))
	assert.NotEmpty(t, mockPermissionRepository.Get(sharedMock.ExistsAccountId, newResource.Id))
	assert.True(t, mockPermissionCacheInvalidator.InvalidateCalledWith(sharedMock.ExistsAccountId))
}

func TestService_CreateResourceValidationError(t *testing.T) {
	newResource := sharedModels.Resource{
		Id:    "new-resource",
		Title: "Some-New-Resource",
	}
	req := request.CreateResource{
		Resource: newResource,
	}

	response := service.CreateResource(req)

	assert.False(t, mockResourceRepository.Has(newResource))
	assert.False(t, response.IsOk())
	assert.True(t, response.HasData())
	assert.Equal(t, sharedError.ValidationErrorCode, response.Data().(sharedError.ServiceError).Code)
	assert.Equal(
		t,
		validation.MakeFailedValidationDescription(validator.New().Struct(req)),
		response.Data().(sharedError.ServiceError).Description,
	)
}

func TestService_CreateResourceDuplicateKeyError(t *testing.T) {
	newResource := sharedModels.Resource{
		Id:    sharedMock.ExistsResourceId,
		Title: "Some-New-Resource",
	}
	assert.True(t, mockResourceRepository.Has(newResource))

	response := service.CreateResource(request.CreateResource{
		AccountId: sharedMock.ExistsAccountId,
		Resource:  newResource,
	})

	assert.False(t, response.IsOk())
	assert.True(t, response.HasData())
	assert.Equal(t, resourceExistsCode, response.Data().(sharedError.ServiceError).Code)
	assert.Equal(t, resourceExistsDescription, response.Data().(sharedError.ServiceError).Description)
}

func TestService_CreateResourceDBError(t *testing.T) {
	newResource := sharedModels.Resource{
		Id:    "new-resource",
		Title: "Some-New-Resource",
	}

	response := service.CreateResource(request.CreateResource{
		AccountId: sharedMock.BadAccountId,
		Resource:  newResource,
	})

	assert.False(t, response.IsOk())
	assert.True(t, response.HasData())
	assert.Equal(t, sharedError.ServerErrorCode, response.Data().(sharedError.ServiceError).Code)
	assert.Equal(t, sharedError.ServerErrorDescription, response.Data().(sharedError.ServiceError).Description)
	assert.False(t, mockResourceRepository.Has(newResource))
	assert.Empty(t, mockPermissionRepository.Get(sharedMock.ExistsAccountId, newResource.Id))
}

func TestService_CreateResourcePermissionDBError(t *testing.T) {
	newResource := sharedModels.Resource{
		Id:    sharedMock.BadResourceId,
		Title: "Some-New-Resource",
	}

	response := service.CreateResource(request.CreateResource{
		AccountId: sharedMock.ExistsAccountId,
		Resource:  newResource,
	})

	assert.False(t, response.IsOk())
	assert.True(t, response.HasData())
	assert.Equal(t, sharedError.ServerErrorCode, response.Data().(sharedError.ServiceError).Code)
	assert.Equal(t, sharedError.ServerErrorDescription, response.Data().(sharedError.ServiceError).Description)
	assert.False(t, mockResourceRepository.Has(newResource))
	assert.Empty(t, mockPermissionRepository.Get(sharedMock.ExistsAccountId, newResource.Id))
}

func TestService_ResourcesListSuccess(t *testing.T) {
	response := service.ResourcesList(request.ResourcesList{
		AccountId: sharedMock.ExistsAccountId,
	})

	assert.True(t, response.IsOk())
	assert.True(t, response.HasData())
	assert.Equal(t, mockResourceRepository.Get(sharedMock.ExistsAccountId), response.Data())
}

func TestService_ResourcesListValidationError(t *testing.T) {
	req := request.ResourcesList{}
	response := service.ResourcesList(req)

	assert.False(t, response.IsOk())
	assert.True(t, response.HasData())
	assert.Equal(t, sharedError.ValidationErrorCode, response.Data().(sharedError.ServiceError).Code)
	assert.Equal(
		t,
		validation.MakeFailedValidationDescription(validator.New().Struct(req)),
		response.Data().(sharedError.ServiceError).Description,
	)
}

func TestService_ResourcesListDBError(t *testing.T) {
	response := service.ResourcesList(request.ResourcesList{
		AccountId: sharedMock.BadAccountId,
	})

	assert.False(t, response.IsOk())
	assert.True(t, response.HasData())
	assert.Equal(t, sharedError.ServerErrorCode, response.Data().(sharedError.ServiceError).Code)
	assert.Equal(t, sharedError.ServerErrorDescription, response.Data().(sharedError.ServiceError).Description)
}

func TestService_UpdateResourceSuccess(t *testing.T) {
	defer resetRepositories()
	updatingResource := sharedModels.Resource{
		Id:    sharedMock.ExistsResourceId,
		Title: "Some-New-Title",
	}

	response := service.UpdateResource(request.UpdateResource{
		AccountId: sharedMock.ExistsAccountId,
		Resource:  updatingResource,
	})

	assert.True(t, response.IsOk())
	assert.False(t, response.HasData())
	assert.Contains(t, mockResourceRepository.Get(sharedMock.ExistsAccountId), updatingResource)
	assert.True(t, mockPermissionCacheInvalidator.InvalidateCalledWith(sharedMock.ExistsAccountId))
}

func TestService_UpdateResourceValidationError(t *testing.T) {
	updatingResource := sharedModels.Resource{
		Id:    sharedMock.ExistsResourceId,
		Title: "Some-New-Title",
	}
	req := request.UpdateResource{
		Resource: updatingResource,
	}

	response := service.UpdateResource(req)

	assert.False(t, response.IsOk())
	assert.True(t, response.HasData())
	assert.Equal(t, sharedError.ValidationErrorCode, response.Data().(sharedError.ServiceError).Code)
	assert.Equal(
		t,
		validation.MakeFailedValidationDescription(validator.New().Struct(req)),
		response.Data().(sharedError.ServiceError).Description,
	)
	assert.NotContains(t, mockResourceRepository.Get(sharedMock.ExistsAccountId), updatingResource)
}

func TestService_UpdateResourceDBError(t *testing.T) {
	updatingResource := sharedModels.Resource{
		Id:    sharedMock.ExistsResourceId,
		Title: "Some-New-Title",
	}

	response := service.UpdateResource(request.UpdateResource{
		AccountId: sharedMock.BadAccountId,
		Resource:  updatingResource,
	})

	assert.False(t, response.IsOk())
	assert.True(t, response.HasData())
	assert.Equal(t, sharedError.ServerErrorCode, response.Data().(sharedError.ServiceError).Code)
	assert.Equal(t, sharedError.ServerErrorDescription, response.Data().(sharedError.ServiceError).Description)
}

func TestService_DeleteResourceSuccess(t *testing.T) {
	defer resetRepositories()

	response := service.DeleteResource(request.DeleteResource{
		AccountId:  sharedMock.ExistsAccountId,
		ResourceId: sharedMock.ExistsResourceId,
	})

	assert.True(t, response.IsOk())
	assert.False(t, response.HasData())
	assert.False(t, mockResourceRepository.Has(sharedModels.Resource{
		Id: sharedMock.ExistsResourceId,
	}))
	assert.True(t, mockPermissionCacheInvalidator.InvalidateCalledWith(sharedMock.ExistsAccountId))
}

func TestService_DeleteResourceValidationError(t *testing.T) {
	req := request.DeleteResource{
		ResourceId: sharedMock.ExistsResourceId,
	}

	response := service.DeleteResource(req)

	assert.False(t, response.IsOk())
	assert.True(t, response.HasData())
	assert.Equal(t, sharedError.ValidationErrorCode, response.Data().(sharedError.ServiceError).Code)
	assert.Equal(
		t,
		validation.MakeFailedValidationDescription(validator.New().Struct(req)),
		response.Data().(sharedError.ServiceError).Description,
	)
	assert.True(t, mockResourceRepository.Has(sharedModels.Resource{
		Id: sharedMock.ExistsResourceId,
	}))
}

func TestService_DeleteResourceDBError(t *testing.T) {
	response := service.DeleteResource(request.DeleteResource{
		AccountId:  sharedMock.BadAccountId,
		ResourceId: sharedMock.ExistsResourceId,
	})

	assert.False(t, response.IsOk())
	assert.True(t, response.HasData())
	assert.Equal(t, sharedError.ServerErrorCode, response.Data().(sharedError.ServiceError).Code)
	assert.Equal(t, sharedError.ServerErrorDescription, response.Data().(sharedError.ServiceError).Description)
	assert.True(t, mockResourceRepository.Has(sharedModels.Resource{
		Id: sharedMock.ExistsResourceId,
	}))
}
