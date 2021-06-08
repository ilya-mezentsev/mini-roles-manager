package export

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"mini-roles-backend/source/domains/files/models"
	"mini-roles-backend/source/domains/files/request"
	sharedMock "mini-roles-backend/source/domains/shared/mock"
	"mini-roles-backend/source/domains/shared/services/response_factory"
	sharedSpec "mini-roles-backend/source/domains/shared/spec"
	"testing"
)

var (
	mockResourceRepository = &sharedMock.ResourceRepository{}
	mockRoleRepository     = &sharedMock.RoleRepository{}
	expectedOkStatus       = response_factory.DefaultResponse().ApplicationStatus()
	expectedErrorStatus    = response_factory.EmptyServerError().ApplicationStatus()

	service = New(mockRoleRepository, mockResourceRepository)
)

func init() {
	mockResourceRepository.Reset()
	mockRoleRepository.Reset()
}

func TestService_MakeExportFileSuccess(t *testing.T) {
	response := service.MakeExportFile(request.ExportRequest{
		AccountId: sharedMock.ExistsAccountId,
	})

	assert.Equal(t, expectedOkStatus, response.ApplicationStatus())
	assert.True(t, response.HasData())
	assert.FileExists(t, response.Data().(string))

	data, err := ioutil.ReadFile(response.Data().(string))
	if err != nil {
		t.Fatalf("unable to read tmp file: %v", err)
	}

	var exportData models.JSONRepresentation
	err = json.Unmarshal(data, &exportData)
	if err != nil {
		t.Fatalf("unable to unmarshal settings to struct: %v", err)
	}

	resources, _ := mockResourceRepository.List(sharedSpec.AccountWithId{
		AccountId: sharedMock.ExistsAccountId,
	})
	roles, _ := mockRoleRepository.List(sharedSpec.AccountWithId{
		AccountId: sharedMock.ExistsAccountId,
	})

	assert.Equal(
		t,
		models.JSONRepresentation{
			Resources: resources,
			Roles:     roles,
		},
		exportData,
	)
}

func TestService_MakeExportFileUnableToFetchResources(t *testing.T) {
	response := service.MakeExportFile(request.ExportRequest{
		AccountId: sharedMock.BadAccountId,
	})

	assert.Equal(t, expectedErrorStatus, response.ApplicationStatus())
	assert.False(t, response.HasData())
}

func TestService_MakeExportFileUnableToFetchRoles(t *testing.T) {
	response := service.MakeExportFile(request.ExportRequest{
		AccountId: sharedMock.BadAccountIdForRoleRepository,
	})

	assert.Equal(t, expectedErrorStatus, response.ApplicationStatus())
	assert.False(t, response.HasData())
}

func TestService_MakeExportFileValidationError(t *testing.T) {
	req := request.ExportRequest{}
	response := service.MakeExportFile(req)

	assert.Equal(t, expectedErrorStatus, response.ApplicationStatus())
	assert.False(t, response.HasData())
}
