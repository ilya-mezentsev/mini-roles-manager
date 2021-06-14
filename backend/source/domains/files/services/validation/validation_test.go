package validation

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"mini-roles-backend/source/domains/files/mock"
	sharedMock "mini-roles-backend/source/domains/shared/mock"
	sharedModels "mini-roles-backend/source/domains/shared/models"
	"testing"
)

func TestValidateSuccess(t *testing.T) {
	representation := mock.MakeValidAppData()
	validatedRepresentation, errorsMessages := Validate(mustMarshal(representation))

	assert.Nil(t, errorsMessages)
	assert.Equal(t, representation, validatedRepresentation)
}

func TestValidateDuplicateResourceId(t *testing.T) {
	representation := mock.MakeValidAppData()
	representation.Resources = append(representation.Resources, representation.Resources[0])

	_, errorsMessages := Validate(mustMarshal(representation))

	assert.Contains(
		t,
		errorsMessages,
		fmt.Sprintf("Resource id %s found 2 times", representation.Resources[0].Id),
	)
}

func TestValidateNotExistsLinkedResource(t *testing.T) {
	representation := mock.MakeValidAppData()
	representation.Resources[0].LinksTo = append(representation.Resources[0].LinksTo, sharedMock.BadResourceId)

	_, errorsMessages := Validate(mustMarshal(representation))

	assert.Contains(
		t,
		errorsMessages,
		fmt.Sprintf(
			"Resource with id %s has not exists link with id: %s",
			representation.Resources[0].Id,
			sharedMock.BadResourceId,
		),
	)
}

func TestValidateDuplicatePermissionId(t *testing.T) {
	representation := mock.MakeValidAppData()
	representation.Resources[0].Permissions[0].Id = representation.Resources[0].Permissions[1].Id

	_, errorsMessages := Validate(mustMarshal(representation))

	assert.Contains(
		t,
		errorsMessages,
		fmt.Sprintf(
			"Resource with id %s has not unique permission id - %s",
			representation.Resources[0].Id,
			representation.Resources[0].Permissions[1].Id,
		),
	)
}

func TestValidateMissedOperation(t *testing.T) {
	representation := mock.MakeValidAppData()
	invalidRepresentation := mock.MakeValidAppData()
	invalidRepresentation.Resources[0].Id = representation.Resources[0].Id
	invalidRepresentation.Resources[0].Permissions = invalidRepresentation.Resources[0].Permissions[2:]

	_, errorsMessages := Validate(mustMarshal(invalidRepresentation))

	assert.Contains(
		t,
		errorsMessages,
		fmt.Sprintf(
			"Resource with id %s has no %s operation",
			representation.Resources[0].Id,
			representation.Resources[0].Permissions[0].Operation,
		),
	)
}

func TestValidateMissedEffect(t *testing.T) {
	representation := mock.MakeValidAppData()
	invalidRepresentation := mock.MakeValidAppData()
	invalidRepresentation.Resources[0].Id = representation.Resources[0].Id
	invalidRepresentation.Resources[0].Permissions = invalidRepresentation.Resources[0].Permissions[1:]

	_, errorsMessages := Validate(mustMarshal(invalidRepresentation))

	assert.Contains(
		t,
		errorsMessages,
		fmt.Sprintf(
			"Resource with id %s has no %s effect for %s operation",
			representation.Resources[0].Id,
			representation.Resources[0].Permissions[0].Effect,
			representation.Resources[0].Permissions[0].Operation,
		),
	)
}

func TestValidateDuplicateRoleId(t *testing.T) {
	representation := mock.MakeValidAppData()
	representation.Roles = append(representation.Roles, representation.Roles[0])

	_, errorsMessages := Validate(mustMarshal(representation))

	assert.Contains(
		t,
		errorsMessages,
		fmt.Sprintf("Role id %s found %d times", representation.Roles[0].Id, 2),
	)
}

func TestValidateMissedExtendsId(t *testing.T) {
	representation := mock.MakeValidAppData()
	representation.Roles[0].Extends = append(representation.Roles[0].Extends, sharedMock.BadRoleId)

	_, errorsMessages := Validate(mustMarshal(representation))

	assert.Contains(
		t,
		errorsMessages,
		fmt.Sprintf(
			"Role with id %s has not exists extends id: %s",
			representation.Roles[0].Id,
			sharedMock.BadRoleId,
		),
	)
}

func TestValidateMissedPermissionId(t *testing.T) {
	representation := mock.MakeValidAppData()
	representation.Roles[0].Permissions = append(representation.Roles[0].Permissions, sharedMock.BadPermissionId)

	_, errorsMessages := Validate(mustMarshal(representation))

	assert.Contains(
		t,
		errorsMessages,
		fmt.Sprintf(
			"Role with id %s has not exists permission id: %s",
			representation.Roles[0].Id,
			sharedMock.BadPermissionId,
		),
	)
}

func TestValidateConflictPermission(t *testing.T) {
	representation := mock.MakeValidAppData()
	representation.Roles[0].Permissions = append(
		representation.Roles[0].Permissions,
		representation.Resources[0].Permissions[0].Id,
	)

	_, errorsMessages := Validate(mustMarshal(representation))

	assert.Contains(
		t,
		errorsMessages,
		fmt.Sprintf(
			"Role with id %s has conflict permissions for resource %s on %s operation",
			representation.Roles[0].Id,
			representation.Resources[0].Id,
			representation.Resources[0].Permissions[0].Operation,
		),
	)
}

func TestValidateBadJSON(t *testing.T) {
	_, errorsMessages := Validate([]byte(`1`))

	assert.Contains(t, errorsMessages, "Unable to unmarshal json data")
}

func BenchmarkValidateParallelOneRole(b *testing.B) {
	representation := mock.MakeValidAppData()
	for n := 0; n < b.N; n++ {
		_, _ = Validate(mustMarshal(representation))
	}
}

func BenchmarkValidateParallelCoupleRoles(b *testing.B) {
	representation := mock.MakeLoadsValidAppData()
	for n := 0; n < b.N; n++ {
		_, _ = Validate(mustMarshal(representation))
	}
}

func BenchmarkValidateSingleGoroutineOneRole(b *testing.B) {
	representation := mock.MakeValidAppData()
	for n := 0; n < b.N; n++ {
		_, _ = singleGoroutineValidate(mustMarshal(representation))
	}
}

func BenchmarkValidateSingleGoroutineCoupleRoles(b *testing.B) {
	representation := mock.MakeLoadsValidAppData()
	for n := 0; n < b.N; n++ {
		_, _ = singleGoroutineValidate(mustMarshal(representation))
	}
}

func mustMarshal(v sharedModels.AppData) []byte {
	bytes, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	return bytes
}

func singleGoroutineValidate(fileDataBytes []byte) (sharedModels.AppData, []string) {
	var fileData sharedModels.AppData
	err := json.Unmarshal(fileDataBytes, &fileData)
	if err != nil {
		log.Errorf("Unable to unmarshal json data: %v", err)

		return fileData, []string{
			"Unable to unmarshal json data",
		}
	}

	var errorMessages []string
	for _, resourcesCheckFn := range []CheckFn{
		resourcesIdsAreUnique,
		allLinksAreExist,
		permissionsIdsAreUnique,
		allOperationsAreExist,
		rolesIdsAreUnique,
		extendsIdsAreExist,
		permissionsAreExist,
		permissionsDoNotConflict,
	} {
		errorMessages = append(errorMessages, resourcesCheckFn(fileData)...)
	}

	return fileData, errorMessages
}
