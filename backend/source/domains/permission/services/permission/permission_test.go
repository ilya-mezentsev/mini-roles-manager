package permission

import (
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"mini-roles-backend/source/domains/permission/mock"
	"mini-roles-backend/source/domains/permission/models"
	"mini-roles-backend/source/domains/permission/request"
	sharedError "mini-roles-backend/source/domains/shared/error"
	sharedMock "mini-roles-backend/source/domains/shared/mock"
	"mini-roles-backend/source/domains/shared/services/response_factory"
	"os"
	"testing"
)

var (
	mockRepository      = mock.PermissionRepository{}
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

func TestService_HasPermissionPermit(t *testing.T) {
	response := New(mockRepository).HasPermission(
		sharedMock.ExistsAccountId,
		request.PermissionAccess{
			RoleId:     sharedMock.ExistsRoleId,
			ResourceId: sharedMock.ExistsResourceId,
			Operation:  mock.PermittedOperation,
		},
	)

	assert.Equal(t, expectedOkStatus, response.ApplicationStatus())
	assert.True(t, response.HasData())
	assert.Equal(t, permitEffectCode, response.Data().(models.EffectResponse).Effect)
}

func TestService_HasPermissionPermitByLinkingResource(t *testing.T) {
	response := New(mockRepository).HasPermission(
		sharedMock.ExistsAccountId,
		request.PermissionAccess{
			RoleId:     sharedMock.ExistsRoleId,
			ResourceId: sharedMock.ExistsResourceId,
			Operation:  mock.DefinedOnLinkingOperation,
		},
	)

	assert.Equal(t, expectedOkStatus, response.ApplicationStatus())
	assert.True(t, response.HasData())
	assert.Equal(t, permitEffectCode, response.Data().(models.EffectResponse).Effect)
}

func TestService_HasPermissionDeny(t *testing.T) {
	response := New(mockRepository).HasPermission(
		sharedMock.ExistsAccountId,
		request.PermissionAccess{
			RoleId:     sharedMock.ExistsRoleId,
			ResourceId: sharedMock.ExistsResourceId,
			Operation:  mock.DeniedOperation,
		},
	)

	assert.Equal(t, expectedOkStatus, response.ApplicationStatus())
	assert.True(t, response.HasData())
	assert.Equal(t, denyEffectCode, response.Data().(models.EffectResponse).Effect)
}

func TestService_HasPermissionDenyByUndefinedOperation(t *testing.T) {
	response := New(mockRepository).HasPermission(
		sharedMock.ExistsAccountId,
		request.PermissionAccess{
			RoleId:     sharedMock.ExistsRoleId,
			ResourceId: sharedMock.ExistsResourceId,
			Operation:  mock.UndefinedOperation,
		},
	)

	assert.Equal(t, expectedOkStatus, response.ApplicationStatus())
	assert.True(t, response.HasData())
	assert.Equal(t, denyEffectCode, response.Data().(models.EffectResponse).Effect)
}

func TestService_HasPermissionValidationError(t *testing.T) {
	req := request.PermissionAccess{
		Operation: "foo",
	}
	response := New(mockRepository).HasPermission(sharedMock.BadAccountId, req)

	assert.Equal(t, expectedErrorStatus, response.ApplicationStatus())
	assert.True(t, response.HasData())
	assert.Equal(t, sharedError.ValidationErrorCode, response.Data().(sharedError.ServiceError).Code)
	assert.Equal(t, validator.New().Struct(req).Error(), response.Data().(sharedError.ServiceError).Description)
}

func TestService_HasPermissionDBError(t *testing.T) {
	response := New(mockRepository).HasPermission(
		sharedMock.BadAccountId,
		request.PermissionAccess{
			RoleId:     sharedMock.ExistsRoleId,
			ResourceId: sharedMock.ExistsResourceId,
			Operation:  mock.DeniedOperation,
		},
	)

	assert.Equal(t, expectedErrorStatus, response.ApplicationStatus())
	assert.True(t, response.HasData())
	assert.Equal(t, sharedError.ServerErrorCode, response.Data().(sharedError.ServiceError).Code)
	assert.Equal(t, sharedError.ServerErrorDescription, response.Data().(sharedError.ServiceError).Description)
}
