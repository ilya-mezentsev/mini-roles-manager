package import_file

import (
	"github.com/stretchr/testify/assert"
	"io"
	"mini-roles-backend/source/domains/files/mock"
	"mini-roles-backend/source/domains/files/request"
	sharedError "mini-roles-backend/source/domains/shared/error"
	sharedMock "mini-roles-backend/source/domains/shared/mock"
	"os"
	"testing"
)

var (
	resetResourcesRepository     = &mock.ResetResourcesRepository{}
	resetRolesRepository         = &mock.ResetRolesRepository{}
	resetRolesVersionsRepository = &mock.ResetRolesVersionsRepository{}
)

func init() {
	cleanRepositories()
}

func TestService_ImportFromFileSuccess(t *testing.T) {
	defer cleanRepositories()

	assert.Empty(t, resetResourcesRepository.Resources)
	assert.Empty(t, resetRolesRepository.Roles)
	assert.Empty(t, resetRolesVersionsRepository.RolesVersions)
	assert.Empty(t, resetRolesVersionsRepository.DefaultRolesVersions)

	response := makeRegularService().ImportFromFile(request.ImportRequest{
		AccountId: sharedMock.ExistsAccountId,
		File:      mustGetValidPermissionFile(),
	})

	assert.True(t, response.IsOk())
	assert.False(t, response.HasData())

	assert.NotEmpty(t, resetResourcesRepository.Resources)
	assert.NotEmpty(t, resetRolesRepository.Roles)
	assert.NotEmpty(t, resetRolesVersionsRepository.RolesVersions)
	assert.NotEmpty(t, resetRolesVersionsRepository.DefaultRolesVersions)
}

func TestService_ImportFromFileValidationError(t *testing.T) {
	defer cleanRepositories()

	req := request.ImportRequest{
		AccountId: sharedMock.ExistsAccountId,
		File:      mustGetInvalidPermissionFile(),
	}

	assert.Empty(t, resetResourcesRepository.Resources)
	assert.Empty(t, resetRolesRepository.Roles)
	assert.Empty(t, resetRolesVersionsRepository.RolesVersions)
	assert.Empty(t, resetRolesVersionsRepository.DefaultRolesVersions)

	response := makeRegularService().ImportFromFile(req)

	assert.False(t, response.IsOk())
	assert.Equal(t, invalidImportFileCode, response.Data().(sharedError.ServiceError).Code)
	assert.Equal(
		t,
		"Role with id author has not exists permission id: 9e7b2cebc4f33d550cff0deee86b87e3, for version medium-load",
		response.Data().(sharedError.ServiceError).Description,
	)

	assert.Empty(t, resetResourcesRepository.Resources)
	assert.Empty(t, resetRolesRepository.Roles)
	assert.Empty(t, resetRolesVersionsRepository.RolesVersions)
	assert.Empty(t, resetRolesVersionsRepository.DefaultRolesVersions)
}

func TestService_ImportFromFileRolesVersionResetError(t *testing.T) {
	defer cleanRepositories()

	assert.Empty(t, resetResourcesRepository.Resources)
	assert.Empty(t, resetRolesRepository.Roles)
	assert.Empty(t, resetRolesVersionsRepository.RolesVersions)
	assert.Empty(t, resetRolesVersionsRepository.DefaultRolesVersions)

	response := makeServiceWithErroredRolesVersionsRepository().ImportFromFile(request.ImportRequest{
		AccountId: sharedMock.ExistsAccountId,
		File:      mustGetValidPermissionFile(),
	})

	assert.False(t, response.IsOk())
	assert.True(t, response.HasData())
	assert.Equal(t, sharedError.ServerErrorCode, response.Data().(sharedError.ServiceError).Code)
	assert.Equal(t, sharedError.ServerErrorDescription, response.Data().(sharedError.ServiceError).Description)

	assert.Empty(t, resetResourcesRepository.Resources)
	assert.Empty(t, resetRolesRepository.Roles)
	assert.Empty(t, resetRolesVersionsRepository.RolesVersions)
	assert.Empty(t, resetRolesVersionsRepository.DefaultRolesVersions)
}

func TestService_ImportFromFileResourcesResetError(t *testing.T) {
	defer cleanRepositories()

	assert.Empty(t, resetResourcesRepository.Resources)
	assert.Empty(t, resetRolesRepository.Roles)
	assert.Empty(t, resetRolesVersionsRepository.RolesVersions)
	assert.Empty(t, resetRolesVersionsRepository.DefaultRolesVersions)

	response := makeServiceWithErroredResourcesRepository().ImportFromFile(request.ImportRequest{
		AccountId: sharedMock.ExistsAccountId,
		File:      mustGetValidPermissionFile(),
	})

	assert.False(t, response.IsOk())
	assert.True(t, response.HasData())
	assert.Equal(t, sharedError.ServerErrorCode, response.Data().(sharedError.ServiceError).Code)
	assert.Equal(t, sharedError.ServerErrorDescription, response.Data().(sharedError.ServiceError).Description)

	assert.Empty(t, resetResourcesRepository.Resources)
	assert.Empty(t, resetRolesRepository.Roles)
	assert.NotEmpty(t, resetRolesVersionsRepository.RolesVersions)
	assert.NotEmpty(t, resetRolesVersionsRepository.DefaultRolesVersions)
}

func TestService_ImportFromFileRolesResetError(t *testing.T) {
	defer cleanRepositories()

	assert.Empty(t, resetResourcesRepository.Resources)
	assert.Empty(t, resetRolesRepository.Roles)
	assert.Empty(t, resetRolesVersionsRepository.RolesVersions)
	assert.Empty(t, resetRolesVersionsRepository.DefaultRolesVersions)

	response := makeServiceWithErroredRolesRepository().ImportFromFile(request.ImportRequest{
		AccountId: sharedMock.ExistsAccountId,
		File:      mustGetValidPermissionFile(),
	})

	assert.False(t, response.IsOk())
	assert.True(t, response.HasData())
	assert.Equal(t, sharedError.ServerErrorCode, response.Data().(sharedError.ServiceError).Code)
	assert.Equal(t, sharedError.ServerErrorDescription, response.Data().(sharedError.ServiceError).Description)

	assert.NotEmpty(t, resetResourcesRepository.Resources)
	assert.Empty(t, resetRolesRepository.Roles)
	assert.NotEmpty(t, resetRolesVersionsRepository.RolesVersions)
	assert.NotEmpty(t, resetRolesVersionsRepository.DefaultRolesVersions)
}

func TestService_ImportFromFileBadFile(t *testing.T) {
	defer cleanRepositories()

	assert.Empty(t, resetResourcesRepository.Resources)
	assert.Empty(t, resetRolesRepository.Roles)
	assert.Empty(t, resetRolesVersionsRepository.RolesVersions)
	assert.Empty(t, resetRolesVersionsRepository.DefaultRolesVersions)

	response := makeServiceWithErroredRolesRepository().ImportFromFile(request.ImportRequest{
		AccountId: sharedMock.ExistsAccountId,
		File:      mock.ErroredReader{},
	})

	assert.False(t, response.IsOk())
	assert.True(t, response.HasData())
	assert.Equal(t, sharedError.ServerErrorCode, response.Data().(sharedError.ServiceError).Code)
	assert.Equal(t, sharedError.ServerErrorDescription, response.Data().(sharedError.ServiceError).Description)

	assert.Empty(t, resetResourcesRepository.Resources)
	assert.Empty(t, resetRolesRepository.Roles)
	assert.Empty(t, resetRolesVersionsRepository.RolesVersions)
	assert.Empty(t, resetRolesVersionsRepository.DefaultRolesVersions)
}

func makeRegularService() Service {
	return New(
		resetRolesVersionsRepository,
		resetResourcesRepository,
		resetRolesRepository,
	)
}

func makeServiceWithErroredRolesRepository() Service {
	return New(
		resetRolesVersionsRepository,
		resetResourcesRepository,
		mock.ResetRolesErrorRepository{},
	)
}

func makeServiceWithErroredRolesVersionsRepository() Service {
	return New(
		mock.ResetRolesVersionsErrorRepository{},
		resetResourcesRepository,
		resetRolesRepository,
	)
}

func makeServiceWithErroredResourcesRepository() Service {
	return New(
		resetRolesVersionsRepository,
		mock.ResetResourcesErrorRepository{},
		resetRolesRepository,
	)
}

func mustGetValidPermissionFile() io.Reader {
	return mustGetFile(os.Getenv("VALID_PERMISSIONS_FILE"))
}

func mustGetInvalidPermissionFile() io.Reader {
	return mustGetFile(os.Getenv("INVALID_PERMISSIONS_FILE"))
}

func mustGetFile(filename string) io.Reader {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	return file
}

func cleanRepositories() {
	resetResourcesRepository.Clean()
	resetRolesRepository.Clean()
	resetRolesVersionsRepository.Clean()
}
