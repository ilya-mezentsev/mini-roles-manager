package response_factory

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestDefaultResponse(t *testing.T) {
	response := DefaultResponse()

	assert.Equal(t, statusOk, response.ApplicationStatus())
	assert.False(t, response.HasData())
	assert.Equal(t, http.StatusNoContent, response.HttpStatus())
	assert.Nil(t, response.GetData())
}

func TestSuccessResponse(t *testing.T) {
	someData := `data`
	response := SuccessResponse(someData)

	assert.Equal(t, statusOk, response.ApplicationStatus())
	assert.True(t, response.HasData())
	assert.Equal(t, http.StatusOK, response.HttpStatus())
	assert.Equal(t, someData, response.GetData())
}

func TestServerErrorResponse(t *testing.T) {
	someData := `data`
	response := ServerError(someData)

	assert.Equal(t, statusError, response.ApplicationStatus())
	assert.True(t, response.HasData())
	assert.Equal(t, http.StatusInternalServerError, response.HttpStatus())
	assert.Equal(t, someData, response.GetData())
}

func TestClientErrorResponse(t *testing.T) {
	someData := `data`
	response := ClientError(someData)

	assert.Equal(t, statusError, response.ApplicationStatus())
	assert.True(t, response.HasData())
	assert.Equal(t, http.StatusBadRequest, response.HttpStatus())
	assert.Equal(t, someData, response.GetData())
}

func TestForbiddenErrorResponse(t *testing.T) {
	someData := `data`
	response := ForbiddenError(someData)

	assert.Equal(t, statusError, response.ApplicationStatus())
	assert.True(t, response.HasData())
	assert.Equal(t, http.StatusForbidden, response.HttpStatus())
	assert.Equal(t, someData, response.GetData())
}

func TestUnauthorizedErrorResponse(t *testing.T) {
	someData := `data`
	response := UnauthorizedError(someData)

	assert.Equal(t, statusError, response.ApplicationStatus())
	assert.True(t, response.HasData())
	assert.Equal(t, http.StatusUnauthorized, response.HttpStatus())
	assert.Equal(t, someData, response.GetData())
}
