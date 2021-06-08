package role

import (
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"mini-roles-backend/source/domains/role/request"
	sharedError "mini-roles-backend/source/domains/shared/error"
	sharedMock "mini-roles-backend/source/domains/shared/mock"
	sharedModels "mini-roles-backend/source/domains/shared/models"
	"mini-roles-backend/source/domains/shared/services/response_factory"
	"mini-roles-backend/source/domains/shared/services/validation"
	sharedSpec "mini-roles-backend/source/domains/shared/spec"
	"os"
	"testing"
)

var (
	mockRepository      = &sharedMock.RoleRepository{}
	expectedOkStatus    = response_factory.DefaultResponse().ApplicationStatus()
	expectedErrorStatus = response_factory.EmptyServerError().ApplicationStatus()
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

	response := New(mockRepository).Create(request.CreateRole{
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

	response := New(mockRepository).Create(request.CreateRole{
		AccountId: sharedMock.ExistsAccountId,
		Role:      newRole,
	})

	assert.Equal(t, expectedErrorStatus, response.ApplicationStatus())
	assert.True(t, response.HasData())
	assert.Equal(t, roleExistsCode, response.Data().(sharedError.ServiceError).Code)
	assert.Equal(t, roleExistsDescription, response.Data().(sharedError.ServiceError).Description)
}

func TestService_CreateValidationError(t *testing.T) {
	defer mockRepository.Reset()
	newRole := sharedModels.Role{}
	req := request.CreateRole{
		AccountId: sharedMock.ExistsAccountId,
		Role:      newRole,
	}

	response := New(mockRepository).Create(req)

	assert.False(t, mockRepository.Has(newRole))
	assert.Equal(t, expectedErrorStatus, response.ApplicationStatus())
	assert.True(t, response.HasData())
	assert.Equal(t, sharedError.ValidationErrorCode, response.Data().(sharedError.ServiceError).Code)
	assert.Equal(
		t,
		validation.MakeFailedValidationDescription(validator.New().Struct(req)),
		response.Data().(sharedError.ServiceError).Description,
	)
}

func TestService_CreateDBError(t *testing.T) {
	defer mockRepository.Reset()
	newRole := sharedModels.Role{
		Id:    "some-new-role",
		Title: "Some New Role Title",
	}

	response := New(mockRepository).Create(request.CreateRole{
		AccountId: sharedMock.BadAccountId,
		Role:      newRole,
	})

	assert.False(t, mockRepository.Has(newRole))
	assert.Equal(t, expectedErrorStatus, response.ApplicationStatus())
	assert.Equal(t, sharedError.ServerErrorCode, response.Data().(sharedError.ServiceError).Code)
	assert.Equal(t, sharedError.ServerErrorDescription, response.Data().(sharedError.ServiceError).Description)
}

func TestService_RolesListSuccess(t *testing.T) {
	expectedRoles, err := mockRepository.List(sharedSpec.AccountWithId{
		AccountId: sharedMock.ExistsAccountId,
	})
	assert.Nil(t, err)
	assert.NotEmpty(t, expectedRoles)

	response := New(mockRepository).RolesList(request.RolesList{
		AccountId: sharedMock.ExistsAccountId,
	})

	assert.Equal(t, expectedOkStatus, response.ApplicationStatus())
	assert.True(t, response.HasData())
	assert.Equal(t, expectedRoles, response.Data())
}

func TestService_RolesListEmpty(t *testing.T) {
	response := New(mockRepository).RolesList(request.RolesList{
		AccountId: "some-id",
	})

	assert.Equal(t, expectedOkStatus, response.ApplicationStatus())
	assert.True(t, response.HasData())
	assert.Empty(t, response.Data())
}

func TestService_RolesListValidationError(t *testing.T) {
	req := request.RolesList{}

	response := New(mockRepository).RolesList(req)

	assert.Equal(t, expectedErrorStatus, response.ApplicationStatus())
	assert.True(t, response.HasData())
	assert.Equal(t, sharedError.ValidationErrorCode, response.Data().(sharedError.ServiceError).Code)
	assert.Equal(
		t,
		validation.MakeFailedValidationDescription(validator.New().Struct(req)),
		response.Data().(sharedError.ServiceError).Description,
	)
}

func TestService_RolesListDBError(t *testing.T) {
	response := New(mockRepository).RolesList(request.RolesList{
		AccountId: sharedMock.BadAccountId,
	})

	assert.Equal(t, expectedErrorStatus, response.ApplicationStatus())
	assert.Equal(t, sharedError.ServerErrorCode, response.Data().(sharedError.ServiceError).Code)
	assert.Equal(t, sharedError.ServerErrorDescription, response.Data().(sharedError.ServiceError).Description)
}

func TestService_UpdateRoleSuccess(t *testing.T) {
	defer mockRepository.Reset()
	updatingRole := sharedModels.Role{
		Id:    sharedMock.ExistsRoleId,
		Title: "some-title",
	}

	response := New(mockRepository).UpdateRole(request.UpdateRole{
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
	req := request.UpdateRole{
		Role: updatingRole,
	}

	response := New(mockRepository).UpdateRole(req)

	assert.Equal(t, expectedErrorStatus, response.ApplicationStatus())
	assert.True(t, response.HasData())
	assert.Equal(t, sharedError.ValidationErrorCode, response.Data().(sharedError.ServiceError).Code)
	assert.Equal(
		t,
		validation.MakeFailedValidationDescription(validator.New().Struct(req)),
		response.Data().(sharedError.ServiceError).Description,
	)
}

func TestService_UpdateRoleDBError(t *testing.T) {
	updatingRole := sharedModels.Role{
		Id:    sharedMock.ExistsRoleId,
		Title: "some-title",
	}

	response := New(mockRepository).UpdateRole(request.UpdateRole{
		AccountId: sharedMock.BadAccountId,
		Role:      updatingRole,
	})

	assert.Equal(t, expectedErrorStatus, response.ApplicationStatus())
	assert.Equal(t, sharedError.ServerErrorCode, response.Data().(sharedError.ServiceError).Code)
	assert.Equal(t, sharedError.ServerErrorDescription, response.Data().(sharedError.ServiceError).Description)
}

func TestService_DeleteRoleSuccess(t *testing.T) {
	defer mockRepository.Reset()

	response := New(mockRepository).DeleteRole(request.DeleteRole{
		AccountId: sharedMock.ExistsAccountId,
		RoleId:    sharedMock.ExistsRoleId,
	})

	assert.Equal(t, expectedOkStatus, response.ApplicationStatus())
	assert.False(t, response.HasData())
	assert.False(t, mockRepository.Has(sharedModels.Role{Id: sharedMock.ExistsRoleId}))
}

func TestService_DeleteRoleValidationError(t *testing.T) {
	req := request.DeleteRole{}
	response := New(mockRepository).DeleteRole(req)

	assert.True(t, mockRepository.Has(sharedModels.Role{Id: sharedMock.ExistsRoleId}))
	assert.Equal(t, expectedErrorStatus, response.ApplicationStatus())
	assert.True(t, response.HasData())
	assert.Equal(t, sharedError.ValidationErrorCode, response.Data().(sharedError.ServiceError).Code)
	assert.Equal(
		t,
		validation.MakeFailedValidationDescription(validator.New().Struct(req)),
		response.Data().(sharedError.ServiceError).Description,
	)
}

func TestService_DeleteRoleDBError(t *testing.T) {
	response := New(mockRepository).DeleteRole(request.DeleteRole{
		AccountId: sharedMock.BadAccountId,
		RoleId:    sharedMock.ExistsRoleId,
	})

	assert.True(t, mockRepository.Has(sharedModels.Role{Id: sharedMock.ExistsRoleId}))
	assert.Equal(t, expectedErrorStatus, response.ApplicationStatus())
	assert.True(t, response.HasData())
	assert.Equal(t, sharedError.ServerErrorCode, response.Data().(sharedError.ServiceError).Code)
	assert.Equal(t, sharedError.ServerErrorDescription, response.Data().(sharedError.ServiceError).Description)
}
