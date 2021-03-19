package role

import (
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"mini-roles-backend/source/domains/role/mock"
	"mini-roles-backend/source/domains/role/request"
	sharedError "mini-roles-backend/source/domains/shared/error"
	sharedMock "mini-roles-backend/source/domains/shared/mock"
	sharedModels "mini-roles-backend/source/domains/shared/models"
	"mini-roles-backend/source/domains/shared/services/response_factory"
	"os"
	"testing"
)

var (
	mockRepository      = &mock.RoleRepository{}
	expectedOkStatus    = response_factory.DefaultResponse().ApplicationStatus()
	expectedErrorStatus = response_factory.ServerError(nil).ApplicationStatus()
)

func init() {
	mockRepository.Reset()
}

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}

func TestService_CreateSuccess(t *testing.T) {
	defer mockRepository.Reset()
	newRole := sharedModels.Role{
		Id:    "some-new-role",
		Title: "Some New Role Title",
	}

	response := New(mockRepository).Create(request.CreateRoleRequest{
		AccountId: sharedMock.ExistsAccountId,
		Role:      newRole,
	})

	assert.True(t, mockRepository.Has(newRole))
	assert.Equal(t, expectedOkStatus, response.ApplicationStatus())
	assert.False(t, response.HasData())
}

func TestService_CreateDuplicateKeyError(t *testing.T) {
	defer mockRepository.Reset()
	newRole := sharedModels.Role{
		Id: sharedMock.ExistsRoleId,
	}
	assert.True(t, mockRepository.Has(newRole))

	response := New(mockRepository).Create(request.CreateRoleRequest{
		AccountId: sharedMock.ExistsAccountId,
		Role:      newRole,
	})

	assert.Equal(t, expectedErrorStatus, response.ApplicationStatus())
	assert.True(t, response.HasData())
	assert.Equal(t, roleExistsCode, response.GetData().(sharedError.ServiceError).Code)
	assert.Equal(t, roleExistsDescription, response.GetData().(sharedError.ServiceError).Description)
}

func TestService_CreateValidationError(t *testing.T) {
	defer mockRepository.Reset()
	newRole := sharedModels.Role{}
	req := request.CreateRoleRequest{
		AccountId: sharedMock.ExistsAccountId,
		Role:      newRole,
	}

	response := New(mockRepository).Create(req)

	assert.False(t, mockRepository.Has(newRole))
	assert.Equal(t, expectedErrorStatus, response.ApplicationStatus())
	assert.True(t, response.HasData())
	assert.Equal(t, sharedError.ValidationErrorCode, response.GetData().(sharedError.ServiceError).Code)
	assert.Equal(t, validator.New().Struct(req).Error(), response.GetData().(sharedError.ServiceError).Description)
}

func TestService_CreateDBError(t *testing.T) {
	defer mockRepository.Reset()
	newRole := sharedModels.Role{
		Id:    "some-new-role",
		Title: "Some New Role Title",
	}

	response := New(mockRepository).Create(request.CreateRoleRequest{
		AccountId: sharedMock.BadAccountId,
		Role:      newRole,
	})

	assert.False(t, mockRepository.Has(newRole))
	assert.Equal(t, expectedErrorStatus, response.ApplicationStatus())
	assert.Equal(t, sharedError.ServerErrorCode, response.GetData().(sharedError.ServiceError).Code)
	assert.Equal(t, sharedError.ServerErrorDescription, response.GetData().(sharedError.ServiceError).Description)
}

func TestService_RolesListSuccess(t *testing.T) {
	expectedRoles, err := mockRepository.List(sharedMock.ExistsAccountId)
	assert.Nil(t, err)
	assert.NotEmpty(t, expectedRoles)

	response := New(mockRepository).RolesList(request.RolesListRequest{
		AccountId: sharedMock.ExistsAccountId,
	})

	assert.Equal(t, expectedOkStatus, response.ApplicationStatus())
	assert.True(t, response.HasData())
	assert.Equal(t, expectedRoles, response.GetData())
}

func TestService_RolesListEmpty(t *testing.T) {
	response := New(mockRepository).RolesList(request.RolesListRequest{
		AccountId: "some-id",
	})

	assert.Equal(t, expectedOkStatus, response.ApplicationStatus())
	assert.True(t, response.HasData())
	assert.Empty(t, response.GetData())
}

func TestService_RolesListValidationError(t *testing.T) {
	req := request.RolesListRequest{}

	response := New(mockRepository).RolesList(req)

	assert.Equal(t, expectedErrorStatus, response.ApplicationStatus())
	assert.True(t, response.HasData())
	assert.Equal(t, sharedError.ValidationErrorCode, response.GetData().(sharedError.ServiceError).Code)
	assert.Equal(t, validator.New().Struct(req).Error(), response.GetData().(sharedError.ServiceError).Description)
}

func TestService_RolesListDBError(t *testing.T) {
	response := New(mockRepository).RolesList(request.RolesListRequest{
		AccountId: sharedMock.BadAccountId,
	})

	assert.Equal(t, expectedErrorStatus, response.ApplicationStatus())
	assert.Equal(t, sharedError.ServerErrorCode, response.GetData().(sharedError.ServiceError).Code)
	assert.Equal(t, sharedError.ServerErrorDescription, response.GetData().(sharedError.ServiceError).Description)
}

func TestService_UpdateRoleSuccess(t *testing.T) {
	defer mockRepository.Reset()
	updatingRole := sharedModels.Role{
		Id:    sharedMock.ExistsRoleId,
		Title: "some-title",
	}

	response := New(mockRepository).UpdateRole(request.UpdateRoleRequest{
		AccountId: sharedMock.ExistsAccountId,
		Role:      updatingRole,
	})

	assert.Equal(t, expectedOkStatus, response.ApplicationStatus())
	assert.False(t, response.HasData())
	assert.Equal(t, mockRepository.Get(updatingRole.Id), updatingRole)
}

func TestService_UpdateRoleValidationError(t *testing.T) {
	updatingRole := sharedModels.Role{
		Id:    sharedMock.ExistsRoleId,
		Title: "some-title",
	}
	req := request.UpdateRoleRequest{
		Role: updatingRole,
	}

	response := New(mockRepository).UpdateRole(req)

	assert.Equal(t, expectedErrorStatus, response.ApplicationStatus())
	assert.True(t, response.HasData())
	assert.Equal(t, sharedError.ValidationErrorCode, response.GetData().(sharedError.ServiceError).Code)
	assert.Equal(t, validator.New().Struct(req).Error(), response.GetData().(sharedError.ServiceError).Description)
}

func TestService_UpdateRoleDBError(t *testing.T) {
	updatingRole := sharedModels.Role{
		Id:    sharedMock.ExistsRoleId,
		Title: "some-title",
	}

	response := New(mockRepository).UpdateRole(request.UpdateRoleRequest{
		AccountId: sharedMock.BadAccountId,
		Role:      updatingRole,
	})

	assert.Equal(t, expectedErrorStatus, response.ApplicationStatus())
	assert.Equal(t, sharedError.ServerErrorCode, response.GetData().(sharedError.ServiceError).Code)
	assert.Equal(t, sharedError.ServerErrorDescription, response.GetData().(sharedError.ServiceError).Description)
}

func TestService_DeleteRoleSuccess(t *testing.T) {
	defer mockRepository.Reset()

	response := New(mockRepository).DeleteRole(request.DeleteRoleRequest{
		AccountId: sharedMock.ExistsAccountId,
		RoleId:    sharedMock.ExistsRoleId,
	})

	assert.Equal(t, expectedOkStatus, response.ApplicationStatus())
	assert.False(t, response.HasData())
	assert.False(t, mockRepository.Has(sharedModels.Role{Id: sharedMock.ExistsRoleId}))
}

func TestService_DeleteRoleValidationError(t *testing.T) {
	req := request.DeleteRoleRequest{}
	response := New(mockRepository).DeleteRole(req)

	assert.True(t, mockRepository.Has(sharedModels.Role{Id: sharedMock.ExistsRoleId}))
	assert.Equal(t, expectedErrorStatus, response.ApplicationStatus())
	assert.True(t, response.HasData())
	assert.Equal(t, sharedError.ValidationErrorCode, response.GetData().(sharedError.ServiceError).Code)
	assert.Equal(t, validator.New().Struct(req).Error(), response.GetData().(sharedError.ServiceError).Description)
}

func TestService_DeleteRoleDBError(t *testing.T) {
	response := New(mockRepository).DeleteRole(request.DeleteRoleRequest{
		AccountId: sharedMock.BadAccountId,
		RoleId:    sharedMock.ExistsRoleId,
	})

	assert.True(t, mockRepository.Has(sharedModels.Role{Id: sharedMock.ExistsRoleId}))
	assert.Equal(t, expectedErrorStatus, response.ApplicationStatus())
	assert.True(t, response.HasData())
	assert.Equal(t, sharedError.ServerErrorCode, response.GetData().(sharedError.ServiceError).Code)
	assert.Equal(t, sharedError.ServerErrorDescription, response.GetData().(sharedError.ServiceError).Description)
}
