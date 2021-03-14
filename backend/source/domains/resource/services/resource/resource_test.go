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
	"mini-roles-backend/source/domains/shared/services/response_factory"
	"os"
	"testing"
)

var (
	mockResourceRepository   = &mock.ResourceRepository{}
	mockPermissionRepository = &mock.PermissionRepository{}
	service                  = New(mockResourceRepository, mockPermissionRepository)
	expectedOkStatus         = response_factory.DefaultResponse().ApplicationStatus()
	expectedErrorStatus      = response_factory.ServerError(nil).ApplicationStatus()
)

func init() {
	resetRepositories()
}

func resetRepositories() {
	mockResourceRepository.Reset()
	mockPermissionRepository.Reset()
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

	response := service.CreateResource(request.CreateResourceRequest{
		AccountId: sharedMock.ExistsAccountId,
		Resource:  newResource,
	})

	assert.Equal(t, expectedOkStatus, response.ApplicationStatus())
	assert.False(t, response.HasData())
	assert.True(t, mockResourceRepository.Has(newResource))
	assert.NotEmpty(t, mockPermissionRepository.Get(sharedMock.ExistsAccountId, newResource.Id))
}

func TestService_CreateResourceValidationError(t *testing.T) {
	newResource := sharedModels.Resource{
		Id:    "new-resource",
		Title: "Some-New-Resource",
	}
	req := request.CreateResourceRequest{
		Resource: newResource,
	}

	response := service.CreateResource(req)

	assert.False(t, mockResourceRepository.Has(newResource))
	assert.Equal(t, expectedErrorStatus, response.ApplicationStatus())
	assert.True(t, response.HasData())
	assert.Equal(t, sharedError.ValidationErrorCode, response.GetData().(sharedError.ServiceError).Code)
	assert.Equal(t, validator.New().Struct(req).Error(), response.GetData().(sharedError.ServiceError).Description)
}

func TestService_CreateResourceDuplicateKeyError(t *testing.T) {
	newResource := sharedModels.Resource{
		Id:    sharedMock.ExistsResourceId,
		Title: "Some-New-Resource",
	}
	assert.True(t, mockResourceRepository.Has(newResource))

	response := service.CreateResource(request.CreateResourceRequest{
		AccountId: sharedMock.ExistsAccountId,
		Resource:  newResource,
	})

	assert.Equal(t, expectedErrorStatus, response.ApplicationStatus())
	assert.True(t, response.HasData())
	assert.Equal(t, resourceExistsCode, response.GetData().(sharedError.ServiceError).Code)
	assert.Equal(t, resourceExistsDescription, response.GetData().(sharedError.ServiceError).Description)
}

func TestService_CreateResourceDBError(t *testing.T) {
	newResource := sharedModels.Resource{
		Id:    "new-resource",
		Title: "Some-New-Resource",
	}

	response := service.CreateResource(request.CreateResourceRequest{
		AccountId: sharedMock.BadAccountId,
		Resource:  newResource,
	})

	assert.Equal(t, expectedErrorStatus, response.ApplicationStatus())
	assert.True(t, response.HasData())
	assert.Equal(t, sharedError.ServerErrorCode, response.GetData().(sharedError.ServiceError).Code)
	assert.Equal(t, sharedError.ServerErrorDescription, response.GetData().(sharedError.ServiceError).Description)
	assert.False(t, mockResourceRepository.Has(newResource))
	assert.Empty(t, mockPermissionRepository.Get(sharedMock.ExistsAccountId, newResource.Id))
}

func TestService_CreateResourcePermissionDBError(t *testing.T) {
	newResource := sharedModels.Resource{
		Id:    sharedMock.BadResourceId,
		Title: "Some-New-Resource",
	}

	response := service.CreateResource(request.CreateResourceRequest{
		AccountId: sharedMock.ExistsAccountId,
		Resource:  newResource,
	})

	assert.Equal(t, expectedErrorStatus, response.ApplicationStatus())
	assert.True(t, response.HasData())
	assert.Equal(t, sharedError.ServerErrorCode, response.GetData().(sharedError.ServiceError).Code)
	assert.Equal(t, sharedError.ServerErrorDescription, response.GetData().(sharedError.ServiceError).Description)
	assert.False(t, mockResourceRepository.Has(newResource))
	assert.Empty(t, mockPermissionRepository.Get(sharedMock.ExistsAccountId, newResource.Id))
}

func TestService_ResourcesListSuccess(t *testing.T) {
	response := service.ResourcesList(request.ResourcesListRequest{
		AccountId: sharedMock.ExistsAccountId,
	})

	assert.Equal(t, expectedOkStatus, response.ApplicationStatus())
	assert.True(t, response.HasData())
	assert.Equal(t, mockResourceRepository.Get(sharedMock.ExistsAccountId), response.GetData())
}

func TestService_ResourcesListValidationError(t *testing.T) {
	req := request.ResourcesListRequest{}
	response := service.ResourcesList(req)

	assert.Equal(t, expectedErrorStatus, response.ApplicationStatus())
	assert.True(t, response.HasData())
	assert.Equal(t, sharedError.ValidationErrorCode, response.GetData().(sharedError.ServiceError).Code)
	assert.Equal(t, validator.New().Struct(req).Error(), response.GetData().(sharedError.ServiceError).Description)
}

func TestService_ResourcesListDBError(t *testing.T) {
	response := service.ResourcesList(request.ResourcesListRequest{
		AccountId: sharedMock.BadAccountId,
	})

	assert.Equal(t, expectedErrorStatus, response.ApplicationStatus())
	assert.True(t, response.HasData())
	assert.Equal(t, sharedError.ServerErrorCode, response.GetData().(sharedError.ServiceError).Code)
	assert.Equal(t, sharedError.ServerErrorDescription, response.GetData().(sharedError.ServiceError).Description)
}

func TestService_UpdateResourceSuccess(t *testing.T) {
	defer resetRepositories()
	updatingResource := sharedModels.Resource{
		Id:    sharedMock.ExistsResourceId,
		Title: "Some-New-Title",
	}

	response := service.UpdateResource(request.UpdateResourceRequest{
		AccountId: sharedMock.ExistsAccountId,
		Resource:  updatingResource,
	})

	assert.Equal(t, expectedOkStatus, response.ApplicationStatus())
	assert.False(t, response.HasData())
	assert.Contains(t, mockResourceRepository.Get(sharedMock.ExistsAccountId), updatingResource)
}

func TestService_UpdateResourceValidationError(t *testing.T) {
	updatingResource := sharedModels.Resource{
		Id:    sharedMock.ExistsResourceId,
		Title: "Some-New-Title",
	}
	req := request.UpdateResourceRequest{
		Resource: updatingResource,
	}

	response := service.UpdateResource(req)

	assert.Equal(t, expectedErrorStatus, response.ApplicationStatus())
	assert.True(t, response.HasData())
	assert.Equal(t, sharedError.ValidationErrorCode, response.GetData().(sharedError.ServiceError).Code)
	assert.Equal(t, validator.New().Struct(req).Error(), response.GetData().(sharedError.ServiceError).Description)
	assert.NotContains(t, mockResourceRepository.Get(sharedMock.ExistsAccountId), updatingResource)
}

func TestService_UpdateResourceDBError(t *testing.T) {
	updatingResource := sharedModels.Resource{
		Id:    sharedMock.ExistsResourceId,
		Title: "Some-New-Title",
	}

	response := service.UpdateResource(request.UpdateResourceRequest{
		AccountId: sharedMock.BadAccountId,
		Resource:  updatingResource,
	})

	assert.Equal(t, expectedErrorStatus, response.ApplicationStatus())
	assert.True(t, response.HasData())
	assert.Equal(t, sharedError.ServerErrorCode, response.GetData().(sharedError.ServiceError).Code)
	assert.Equal(t, sharedError.ServerErrorDescription, response.GetData().(sharedError.ServiceError).Description)
}

func TestService_DeleteResourceSuccess(t *testing.T) {
	defer resetRepositories()

	response := service.DeleteResource(request.DeleteResourceRequest{
		AccountId:  sharedMock.ExistsAccountId,
		ResourceId: sharedMock.ExistsResourceId,
	})

	assert.Equal(t, expectedOkStatus, response.ApplicationStatus())
	assert.False(t, response.HasData())
	assert.False(t, mockResourceRepository.Has(sharedModels.Resource{
		Id: sharedMock.ExistsResourceId,
	}))
}

func TestService_DeleteResourceValidationError(t *testing.T) {
	req := request.DeleteResourceRequest{
		ResourceId: sharedMock.ExistsResourceId,
	}

	response := service.DeleteResource(req)

	assert.Equal(t, expectedErrorStatus, response.ApplicationStatus())
	assert.True(t, response.HasData())
	assert.Equal(t, sharedError.ValidationErrorCode, response.GetData().(sharedError.ServiceError).Code)
	assert.Equal(t, validator.New().Struct(req).Error(), response.GetData().(sharedError.ServiceError).Description)
	assert.True(t, mockResourceRepository.Has(sharedModels.Resource{
		Id: sharedMock.ExistsResourceId,
	}))
}

func TestService_DeleteResourceDBError(t *testing.T) {
	response := service.DeleteResource(request.DeleteResourceRequest{
		AccountId:  sharedMock.BadAccountId,
		ResourceId: sharedMock.ExistsResourceId,
	})

	assert.Equal(t, expectedErrorStatus, response.ApplicationStatus())
	assert.True(t, response.HasData())
	assert.Equal(t, sharedError.ServerErrorCode, response.GetData().(sharedError.ServiceError).Code)
	assert.Equal(t, sharedError.ServerErrorDescription, response.GetData().(sharedError.ServiceError).Description)
	assert.True(t, mockResourceRepository.Has(sharedModels.Resource{
		Id: sharedMock.ExistsResourceId,
	}))
}
