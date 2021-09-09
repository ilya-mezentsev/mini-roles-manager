package export

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"mini-roles-backend/source/domains/files/request"
	sharedMock "mini-roles-backend/source/domains/shared/mock"
	sharedModels "mini-roles-backend/source/domains/shared/models"
	sharedSpec "mini-roles-backend/source/domains/shared/spec"
	"testing"
)

var (
	mockResourceRepository            = &sharedMock.ResourceRepository{}
	mockDefaultRolesVersionRepository = &sharedMock.DefaultRolesVersionRepository{}
	mockRoleRepository                = &sharedMock.RoleRepository{}

	service = New(
		mockRoleRepository,
		mockResourceRepository,
		mockDefaultRolesVersionRepository,
	)
)

func init() {
	mockResourceRepository.Reset()
	mockRoleRepository.Reset()
	mockDefaultRolesVersionRepository.Reset()
}

func TestService_MakeExportFileSuccess(t *testing.T) {
	response := service.MakeExportFile(request.ExportRequest{
		AccountId: sharedMock.ExistsAccountId,
	})

	assert.True(t, response.IsOk())
	assert.True(t, response.HasData())
	assert.FileExists(t, response.Data().(string))

	data, err := ioutil.ReadFile(response.Data().(string))
	if err != nil {
		t.Fatalf("unable to read tmp file: %v", err)
	}

	var exportData sharedModels.AppData
	err = json.Unmarshal(data, &exportData)
	if err != nil {
		t.Fatalf("unable to unmarshal settings to struct: %v", err)
	}

	spec := sharedSpec.AccountWithId{
		AccountId: sharedMock.ExistsAccountId,
	}
	resources, _ := mockResourceRepository.List(spec)
	roles, _ := mockRoleRepository.List(spec)
	defaultRolesVersion, _ := mockDefaultRolesVersionRepository.Fetch(spec)

	assert.Equal(
		t,
		sharedModels.AppData{
			Resources:             resources,
			Roles:                 roles,
			DefaultRolesVersionId: defaultRolesVersion.Id,
		},
		exportData,
	)
}

func TestService_MakeExportFileUnableToFetchResources(t *testing.T) {
	response := service.MakeExportFile(request.ExportRequest{
		AccountId: sharedMock.BadAccountId,
	})

	assert.False(t, response.IsOk())
	assert.False(t, response.HasData())
}

func TestService_MakeExportFileUnableToFetchRoles(t *testing.T) {
	response := service.MakeExportFile(request.ExportRequest{
		AccountId: sharedMock.BadAccountIdForRoleRepository,
	})

	assert.False(t, response.IsOk())
	assert.False(t, response.HasData())
}

func TestService_MakeExportFileUnableToFetchDefaultRolesVersion(t *testing.T) {
	response := service.MakeExportFile(request.ExportRequest{
		AccountId: sharedMock.BadAccountIdForDefaultRolesVersionRepository,
	})

	assert.False(t, response.IsOk())
	assert.False(t, response.HasData())
}

func TestService_MakeExportFileValidationError(t *testing.T) {
	req := request.ExportRequest{}
	response := service.MakeExportFile(req)

	assert.False(t, response.IsOk())
	assert.False(t, response.HasData())
}
